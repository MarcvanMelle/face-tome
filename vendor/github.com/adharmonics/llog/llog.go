// Package llog is a structured logger
package llog

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
)

// Fields is a map containing the fields to be logged
type Fields map[string]interface{}

const requestIDHeader string = "x-request-id"
const requestIDKey string = "request_id"

func (f *Fields) addBaseFields() {
	(*f)["time"] = time.Now().Format(time.RFC3339Nano)
	(*f)["v"] = appVersion
}

// addTrading extracts tracing information from context and adds it to the log, if available
// N.B. for llog to successfully retrieve the `x-request-id` key, users of llog must set the key through metadata (as in the test cases)
func (f *Fields) addTracing(ctx context.Context) {
	requestID := RequestIDFromContext(ctx)

	if requestID != "" {
		(*f)[requestIDKey] = requestID
	}
}

func (f *Fields) checkSeverityLevel() (Level, error) {
	severityLevel, ok := (*f)["severity"].(string)
	if !ok {
		return 0, fmt.Errorf("severity level not set for %v", f)
	}

	return ParseLevel(severityLevel)
}

// NewLogger returns a map used for cumulative logging
func NewLogger() Fields {
	f := make(Fields)
	return f
}

// Debug logs at the debug severity level (staging and development)
func (f *Fields) Debug(msg string) {
	(*f)["severity"] = DebugLevel.String()
	(*f)["msg"] = msg

	structuredWrap(f)
}

// Debugf logs at the debug severity level (staging and development) with a formatting directive
func (f *Fields) Debugf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)

	(*f)["severity"] = DebugLevel.String()
	(*f)["msg"] = message
	structuredWrap(f)
}

// Debugw is a cumulative logger method for the Fields map that logs at the debug level
func (f *Fields) Debugw(details Fields) {
	for k, v := range details {
		(*f)[k] = v
	}

	(*f)["severity"] = DebugLevel.String()
	structuredWrap(f)
}

func (f *Fields) fire() {
	err := json.NewEncoder(destination).Encode(f)

	if err != nil {
		fmt.Printf("logging through llog: %v", err)
	}
}

// Info logs at the info level
func (f *Fields) Info(msg string) {
	(*f)["severity"] = InfoLevel.String()
	(*f)["msg"] = msg
	structuredWrap(f)
}

// Infof logs at the info level with a formatting directive
func (f *Fields) Infof(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)

	(*f)["severity"] = InfoLevel.String()
	(*f)["msg"] = message
	structuredWrap(f)
}

// Infow is a cumulative logger method for the Fields map that logs at the info level
func (f *Fields) Infow(details Fields) {
	for k, v := range details {
		(*f)[k] = v
	}

	(*f)["severity"] = InfoLevel.String()
	structuredWrap(f)
}

func (f *Fields) resetWrapper() {
	for k := range *f {
		delete(*f, k)
	}
}

// Debugf logs a message with a formatting directive at the debug severity level
func Debugf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)

	textWrap(message, DebugLevel.String())
}

// Debug logs a message at the debug severity level
func Debug(v ...interface{}) {
	message := make([]string, len(v))
	for i, value := range v {
		message[i] = fmt.Sprint(value)
	}

	textWrap(strings.Join(message, ", "), DebugLevel.String())
}

// Debugw creates a debug-level log from a map
func Debugw(f Fields) {
	f["severity"] = DebugLevel.String()
	structuredWrap(&f)
}

// DebugWithTracing creates a debug-level log from a map, and extracts tracing information from context
func DebugWithTracing(ctx context.Context, f Fields) {
	f["severity"] = DebugLevel.String()
	f.addTracing(ctx)
	structuredWrap(&f)
}

// Infof logs a message with a formatting directive at the info severity level
func Infof(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)

	textWrap(message, InfoLevel.String())
}

// Info logs a message at the info severity level
func Info(v ...interface{}) {
	message := make([]string, len(v))
	for i, value := range v {
		message[i] = fmt.Sprint(value)
	}

	textWrap(strings.Join(message, ", "), InfoLevel.String())
}

// Infow creates an info-level log from a map
func Infow(f Fields) {
	f["severity"] = InfoLevel.String()
	structuredWrap(&f)
}

// InfoWithTracing creates an info-level log from a map, and extracts tracing information from context
func InfoWithTracing(ctx context.Context, f Fields) {
	f["severity"] = InfoLevel.String()
	f.addTracing(ctx)
	structuredWrap(&f)
}

// RequestIDFromContext retrieves the request_id from the context struct
func RequestIDFromContext(ctx context.Context) string {
	// check if request ID is stored in the incomingKey
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md.Get(requestIDHeader)) > 0 {
		return md.Get(requestIDHeader)[0]
	}

	// check if request ID is stored in the outgoingKey
	md, ok = metadata.FromOutgoingContext(ctx)
	if ok && len(md.Get(requestIDHeader)) > 0 {
		return md.Get(requestIDHeader)[0]
	}

	return ""
}

func structuredWrap(msgMap *Fields) {
	msgMap.addBaseFields()

	messageLevel, err := msgMap.checkSeverityLevel()
	if err != nil {
		fmt.Printf("checking log message severity level: %v", err)
		messageLevel = InfoLevel // default to InfoLevel in case of error
	}

	// only write the log if the severity level rises to the specified threshold
	if messageLevel >= appLogLevel {
		msgMap.fire()
	}
}

func textWrap(msg string, level string) {
	data := fieldsPool.Get().(*Fields)
	defer fieldsPool.Put(data)
	defer data.resetWrapper() // defers are executed as LIFO per https://blog.golang.org/defer-panic-and-recover
	(*data)["severity"] = level
	(*data)["msg"] = msg
	data.addBaseFields()

	messageLevel, err := data.checkSeverityLevel()
	if err != nil {
		fmt.Printf("checking log message severity level: %v", err)
		messageLevel = InfoLevel // default to InfoLevel in case of error
	}

	// only write the log if the severity level rises to the specified threshold
	if messageLevel >= appLogLevel {
		data.fire()
	}
}

// WithFields preserves the signature of our previous logging package and returns a Fields map
// The log will be serialized and written when a level is called, e.g. llog.WithFields({}).Info("message")
func WithFields(details Fields) *Fields {
	// make a copy of the map values to prevent a data race during concurrent calls
	data := Fields{}
	for k, v := range details {
		data[k] = v
	}
	return &data
}

// WithTracing behaved like WithFields, but in addition, will extract tracing information from the supplied context struct and add it to the fields map to be logged
// The log will be serialized and written when a level is called, e.g. llog.WithFields({}).Info("message")
func WithTracing(ctx context.Context, details Fields) *Fields {
	// make a copy of the map values to prevent a data race during concurrent calls
	data := Fields{}
	for k, v := range details {
		data[k] = v
	}

	data.addTracing(ctx)
	return &data
}
