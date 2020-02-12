package dockerfile // import "github.com/sequix/moby/builder/dockerfile"

import (
	"os"
	"runtime"
	"testing"

	"github.com/sequix/moby/builder/remotecontext"
	"github.com/sequix/moby/pkg/archive"
	"github.com/sequix/moby/pkg/reexec"
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"gotest.tools/skip"
)

type dispatchTestCase struct {
	name, expectedError string
	cmd                 instructions.Command
	files               map[string]string
}

func init() {
	reexec.Init()
}

func TestDispatch(t *testing.T) {
	if runtime.GOOS != "windows" {
		skip.If(t, os.Getuid() != 0, "skipping test that requires root")
	}
	testCases := []dispatchTestCase{
		{
			name: "ADD multiple files to file",
			cmd: &instructions.AddCommand{SourcesAndDest: instructions.SourcesAndDest{
				"file1.txt",
				"file2.txt",
				"test",
			}},
			expectedError: "When using ADD with more than one source file, the destination must be a directory and end with a /",
			files:         map[string]string{"file1.txt": "test1", "file2.txt": "test2"},
		},
		{
			name: "Wildcard ADD multiple files to file",
			cmd: &instructions.AddCommand{SourcesAndDest: instructions.SourcesAndDest{
				"file*.txt",
				"test",
			}},
			expectedError: "When using ADD with more than one source file, the destination must be a directory and end with a /",
			files:         map[string]string{"file1.txt": "test1", "file2.txt": "test2"},
		},
		{
			name: "COPY multiple files to file",
			cmd: &instructions.CopyCommand{SourcesAndDest: instructions.SourcesAndDest{
				"file1.txt",
				"file2.txt",
				"test",
			}},
			expectedError: "When using COPY with more than one source file, the destination must be a directory and end with a /",
			files:         map[string]string{"file1.txt": "test1", "file2.txt": "test2"},
		},
		{
			name: "ADD multiple files to file with whitespace",
			cmd: &instructions.AddCommand{SourcesAndDest: instructions.SourcesAndDest{
				"test file1.txt",
				"test file2.txt",
				"test",
			}},
			expectedError: "When using ADD with more than one source file, the destination must be a directory and end with a /",
			files:         map[string]string{"test file1.txt": "test1", "test file2.txt": "test2"},
		},
		{
			name: "COPY multiple files to file with whitespace",
			cmd: &instructions.CopyCommand{SourcesAndDest: instructions.SourcesAndDest{
				"test file1.txt",
				"test file2.txt",
				"test",
			}},
			expectedError: "When using COPY with more than one source file, the destination must be a directory and end with a /",
			files:         map[string]string{"test file1.txt": "test1", "test file2.txt": "test2"},
		},
		{
			name: "COPY wildcard no files",
			cmd: &instructions.CopyCommand{SourcesAndDest: instructions.SourcesAndDest{
				"file*.txt",
				"/tmp/",
			}},
			expectedError: "COPY failed: no source files were specified",
			files:         nil,
		},
		{
			name: "COPY url",
			cmd: &instructions.CopyCommand{SourcesAndDest: instructions.SourcesAndDest{
				"https://index.docker.io/robots.txt",
				"/",
			}},
			expectedError: "source can't be a URL for COPY",
			files:         nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contextDir, cleanup := createTestTempDir(t, "", "builder-dockerfile-test")
			defer cleanup()

			for filename, content := range tc.files {
				createTestTempFile(t, contextDir, filename, content, 0777)
			}

			tarStream, err := archive.Tar(contextDir, archive.Uncompressed)

			if err != nil {
				t.Fatalf("Error when creating tar stream: %s", err)
			}

			defer func() {
				if err = tarStream.Close(); err != nil {
					t.Fatalf("Error when closing tar stream: %s", err)
				}
			}()

			context, err := remotecontext.FromArchive(tarStream)

			if err != nil {
				t.Fatalf("Error when creating tar context: %s", err)
			}

			defer func() {
				if err = context.Close(); err != nil {
					t.Fatalf("Error when closing tar context: %s", err)
				}
			}()

			b := newBuilderWithMockBackend()
			sb := newDispatchRequest(b, '`', context, NewBuildArgs(make(map[string]*string)), newStagesBuildResults())
			err = dispatch(sb, tc.cmd)
			assert.Check(t, is.ErrorContains(err, tc.expectedError))
		})
	}
}
