package main

import (
	"context"

	"github.com/adharmonics/llog"
	"google.golang.org/grpc/metadata"
)

func init() {
	llog.ConfigLlog(llog.Options{
		Enabled:    true,
		AppVersion: "v1.0",
		LogLevel:   "DebugLevel",
	})
}

func main() {
	// logging a simple string
	llog.Debug("this line was triggered")
	// output:
	// {"msg":"this line was triggered","severity":"debug","time":"2018-06-15T13:39:59.188969404-04:00","v":"v1.0"}

	llog.Debugf("this is line %d", 2)
	// output:
	// {"msg":"this is line 2","severity":"debug","time":"2018-06-15T13:39:59.189151721-04:00","v":"v1.0"}

	llog.Infof("this is info message %s, attempt %d", "foo", 1)
	// output:
	// {"msg":"this is info message foo, attempt 1","severity":"info","time":"2018-06-29T16:27:05.981212855-04:00","v":"v1.0"}

	// verbose logging for maps, logrus-style
	llog.WithFields(llog.Fields{
		"category": "example",
		"type":     "with multiple fields",
	}).Info("logged at info level for multiple fields")
	// output:
	// {"category":"example","msg":"logged at info level for multiple fields","severity":"info","time":"2018-06-15T13:39:59.18916998-04:00","type":"with multiple fields","v":"v1.0"}

	// verbose logging for maps with formatting directive
	msg := "foo"
	llog.WithFields(llog.Fields{
		"category": "example",
		"type":     "with multiple fields",
	}).Infof("my formatting directive %d: %v", 1, msg)
	// output:
	// {"category":"example","msg":"my formatting directive 1: foo","severity":"info","time":"2018-06-29T16:27:05.981234987-04:00","type":"with multiple fields","v":"v1.0"}

	// shorthand logging of maps
	llog.Infow(llog.Fields{
		"category": "wrapper example",
		"type":     "shorthand, with multiple fields",
	})
	// output:
	// {"category":"wrapper example","severity":"info","time":"2018-06-15T13:39:59.189193092-04:00","type":"shorthand, with multiple fields","v":"v1.0"}

	// logging with tracing information
	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-request-id", "fake-request-id")
	llog.InfoWithTracing(ctx, llog.Fields{
		"category": "tracing example",
		"type":     "with correlation id",
	})
	// output:
	// {"category":"tracing example","request_id":"fake-request-id","severity":"info","time":"2018-07-05T14:27:58.263165042-04:00","type":"with correlation id","v":"v1.0"}

	// logging with tracing information, logrus-style
	llog.WithTracing(ctx, llog.Fields{
		"category": "logrus-style tracing example",
		"type":     "with correlation id",
	}).Info("logging with tracing information")
	// output:
	// {"category":"logrus-style tracing example","msg":"logging with tracing information","request_id":"fake-request-id","severity":"info","time":"2018-07-16T07:12:37.825498301-04:00","type":"with correlation id","v":"v1.0"}

	// cumulative logging
	cLogger := llog.NewLogger()

	cLogger.Infow(llog.Fields{
		"myField": "foo",
	})
	// output:
	// {"myField":"foo","severity":"info","time":"2018-06-15T13:39:59.189201154-04:00","v":"v1.0"}

	cLogger.Infow(llog.Fields{
		"additionalField": "bar",
	})
	// output:
	// {"additionalField":"bar","myField":"foo","severity":"info","time":"2018-06-15T13:39:59.189206055-04:00","v":"v1.0"}
}
