// +build !linux

package archive // import "github.com/sequix/moby/pkg/archive"

func getWhiteoutConverter(format WhiteoutFormat, inUserNS bool) tarWhiteoutConverter {
	return nil
}
