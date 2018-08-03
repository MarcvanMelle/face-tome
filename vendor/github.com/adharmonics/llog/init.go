package llog

import (
	"io"
	"io/ioutil"
	"os"
	"sync"
)

var appVersion string
var destination io.Writer
var appLogLevel Level

// Options is a struct containing initialization options for llog
// Disabled controls whether or not the logs will be output to os.Stdout or disposed (useful for test environments)
// AppVersion is the version of the current application.  It will be attached to all logs for troubleshooting purposes.
type Options struct {
	Enabled    bool
	AppVersion string
	LogLevel   string
}

// fieldsPool caches allocated but unused items for later reuse,
// relieving pressure on the garbage collector.
var fieldsPool *sync.Pool

func init() {
	destination = os.Stdout
	appLogLevel = InfoLevel

	fieldsPool = &sync.Pool{
		New: func() interface{} {
			return &Fields{}
		},
	}
}

// ConfigLlog overrides the default llog initialization with custom options
func ConfigLlog(options Options) {
	if !options.Enabled {
		destination = ioutil.Discard
	}

	appVersion = options.AppVersion

	switch options.LogLevel {
	case "DebugLevel":
		appLogLevel = DebugLevel
	case "InfoLevel":
		appLogLevel = InfoLevel
	case "": // already defaulted to Info
	default:
		panic("Unexpected log level specified")
	}
}
