// +build windows

package daemon // import "github.com/sequix/moby/daemon"

import "github.com/sequix/moby/pkg/plugingetter"

func registerMetricsPluginCallback(getter plugingetter.PluginGetter, sockPath string) {
}

func (daemon *Daemon) listenMetricsSock() (string, error) {
	return "", nil
}
