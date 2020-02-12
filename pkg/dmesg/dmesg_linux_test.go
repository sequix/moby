package dmesg // import "github.com/sequix/moby/pkg/dmesg"

import (
	"testing"
)

func TestDmesg(t *testing.T) {
	t.Logf("dmesg output follows:\n%v", string(Dmesg(512)))
}
