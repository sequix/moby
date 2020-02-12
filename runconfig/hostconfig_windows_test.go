// +build windows

package runconfig // import "github.com/sequix/moby/runconfig"

import (
	"testing"

	"github.com/sequix/moby/api/types/container"
)

func TestValidatePrivileged(t *testing.T) {
	expected := "Windows does not support privileged mode"
	err := validatePrivileged(&container.HostConfig{Privileged: true})
	if err == nil || err.Error() != expected {
		t.Fatalf("Expected %s", expected)
	}
}
