// +build !linux,!freebsd

package logger // import "github.com/sequix/moby/daemon/logger"

import (
	"errors"
	"io"
)

func openPluginStream(a *pluginAdapter) (io.WriteCloser, error) {
	return nil, errors.New("log plugin not supported")
}
