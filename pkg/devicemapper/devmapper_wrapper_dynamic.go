// +build linux,cgo,!static_build

package devicemapper // import "github.com/sequix/moby/pkg/devicemapper"

// #cgo pkg-config: devmapper
import "C"
