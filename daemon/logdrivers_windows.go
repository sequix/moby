package daemon // import "github.com/sequix/moby/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/sequix/moby/daemon/logger/awslogs"
	_ "github.com/sequix/moby/daemon/logger/etwlogs"
	_ "github.com/sequix/moby/daemon/logger/fluentd"
	_ "github.com/sequix/moby/daemon/logger/gcplogs"
	_ "github.com/sequix/moby/daemon/logger/gelf"
	_ "github.com/sequix/moby/daemon/logger/jsonfilelog"
	_ "github.com/sequix/moby/daemon/logger/logentries"
	_ "github.com/sequix/moby/daemon/logger/splunk"
	_ "github.com/sequix/moby/daemon/logger/syslog"
)
