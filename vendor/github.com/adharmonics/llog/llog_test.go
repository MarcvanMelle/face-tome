package llog

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"

	"google.golang.org/grpc/metadata"
)

type logTestCase struct {
	name           string
	data           interface{}
	context        context.Context
	message        string
	style          string
	messageLevel   Level
	appLogLevel    Level
	template       string
	expectedResult []string
}

var logTestCases = []logTestCase{
	logTestCase{
		name: "info field with string value",
		data: Fields{
			"Test": "Foo",
		},
		message:        "info message",
		style:          "WithFields",
		messageLevel:   InfoLevel,
		expectedResult: []string{`"Test":"Foo"`, `"severity":"info"`},
	},
	logTestCase{
		name: "info field with int value",
		data: Fields{
			"Test": 1,
		},
		message:        "info message",
		style:          "WithFields",
		messageLevel:   InfoLevel,
		expectedResult: []string{`"Test":1`, `"severity":"info"`},
	},
	logTestCase{
		name:    "info field with tracing (request_id)",
		context: metadata.AppendToOutgoingContext(context.Background(), requestIDHeader, "fake-request-id"),
		data: Fields{
			"Test": 1,
		},
		message:        "info message",
		style:          "WithTracing",
		messageLevel:   InfoLevel,
		expectedResult: []string{`"Test":1`, `"request_id":"fake-request-id"`, `"severity":"info"`},
	},
	logTestCase{
		name: "debug field with string value",
		data: Fields{
			"Test": "Foo",
		},
		message:        "debug message",
		style:          "WithFields",
		messageLevel:   DebugLevel,
		expectedResult: []string{`"Test":"Foo"`, `"severity":"debug"`},
	},
	logTestCase{
		name: "debug field with int value",
		data: Fields{
			"Test": 1,
		},
		message:        "debug message",
		style:          "WithFields",
		messageLevel:   DebugLevel,
		expectedResult: []string{`"Test":1`, `"severity":"debug"`},
	},
	logTestCase{
		name:           "Debug",
		data:           "Debug Test",
		messageLevel:   DebugLevel,
		expectedResult: []string{`"msg":"Debug Test"`, `"severity":"debug"`},
	},
	logTestCase{
		name:           "Debugf",
		data:           "Debugf Test",
		template:       "Sending: %v",
		messageLevel:   DebugLevel,
		expectedResult: []string{`"msg":"Sending: Debugf Test"`, `"severity":"debug"`},
	},
	logTestCase{
		name:           "Info",
		data:           "Info Test",
		messageLevel:   InfoLevel,
		expectedResult: []string{`"msg":"Info Test"`, `"severity":"info"`},
	},
	logTestCase{
		name:           "Infof",
		data:           "Infof Test",
		messageLevel:   InfoLevel,
		template:       "Sending: %v",
		expectedResult: []string{`"msg":"Sending: Infof Test"`, `"severity":"info"`},
	},
	logTestCase{
		name:           "Infof with multiple arguments",
		data:           []string{"Infof Test", "foo"},
		messageLevel:   InfoLevel,
		template:       "Sending: %v, Result: %v",
		expectedResult: []string{`"msg":"Sending: Infof Test, Result: foo"`, `"severity":"info"`},
	},
	logTestCase{
		name:           "for simple message, do not log debug message when info level logging specified",
		data:           "debug message",
		messageLevel:   DebugLevel,
		appLogLevel:    InfoLevel,
		expectedResult: []string{},
	},
	logTestCase{
		name: "for multiple fields, do not log debug message when info level logging specified",
		data: Fields{
			"Test": "Foo",
		},
		message:        "debug message",
		style:          "WithFields",
		messageLevel:   DebugLevel,
		appLogLevel:    InfoLevel,
		expectedResult: []string{},
	},
	logTestCase{
		name:           "for simple message, log info message when debug level logging specified",
		data:           "info log",
		messageLevel:   InfoLevel,
		appLogLevel:    DebugLevel,
		expectedResult: []string{`"msg":"info log"`, `"severity":"info"`},
	},
	logTestCase{
		name: "for multiple fields, log info message when debug level logging specified",
		data: Fields{
			"Test": "Foo",
		},
		message:        "info message",
		style:          "WithFields",
		messageLevel:   InfoLevel,
		appLogLevel:    DebugLevel,
		expectedResult: []string{`"Test":"Foo"`, `"severity":"info"`},
	},
	logTestCase{
		name: "for multiple fields, log info message with formatting directive",
		data: Fields{
			"Test": "Foo",
		},
		message:        "template info message",
		template:       "Sending: %v",
		style:          "WithFields template",
		messageLevel:   InfoLevel,
		appLogLevel:    InfoLevel,
		expectedResult: []string{`"Test":"Foo"`, `"msg":"Sending: template info message"`, `"severity":"info"`},
	},
	logTestCase{
		name: "for multiple fields, log at info level when using .Infow",
		data: Fields{
			"Test": "Foo",
		},
		style:          "wrapper",
		messageLevel:   InfoLevel,
		appLogLevel:    InfoLevel,
		expectedResult: []string{`"Test":"Foo"`, `"severity":"info"`},
	},
	logTestCase{
		name: "for multiple fields, log at debug level when using .Debugw",
		data: Fields{
			"Test": "Foo",
		},
		style:          "wrapper",
		messageLevel:   DebugLevel,
		appLogLevel:    DebugLevel,
		expectedResult: []string{`"Test":"Foo"`, `"severity":"debug"`},
	},
	logTestCase{
		name: "log tracing information with DebugWithTracing",
		data: Fields{
			"Test": "Foo",
		},
		context:        metadata.AppendToOutgoingContext(context.Background(), requestIDHeader, "fake-request-id"),
		style:          "withTrace",
		messageLevel:   DebugLevel,
		appLogLevel:    DebugLevel,
		expectedResult: []string{`"Test":"Foo"`, `"request_id":"fake-request-id"`, `"severity":"debug"`},
	},
	logTestCase{
		name: "log tracing information with InfoWithTracing",
		data: Fields{
			"Test": "Foo",
		},
		context:        metadata.AppendToOutgoingContext(context.Background(), requestIDHeader, "fake-request-id"),
		style:          "withTrace",
		messageLevel:   InfoLevel,
		appLogLevel:    InfoLevel,
		expectedResult: []string{`"Test":"Foo"`, `"request_id":"fake-request-id"`, `"severity":"info"`},
	},
}

func TestLog(t *testing.T) {
	for _, testCase := range logTestCases {
		ConfigLlog(Options{
			Enabled:    false,
			AppVersion: "test",
			LogLevel:   strings.Title(testCase.messageLevel.String()) + "Level",
		})

		// Pipe creates a synchronous in-memory pipe. It can be used to connect code expecting an io.Reader with code expecting an io.Writer.
		r, w := io.Pipe()
		actualResult := new(bytes.Buffer)
		previousDestination := destination
		destination = w // send log output to the pipe instead of os.Stdout

		defer func() {
			destination = previousDestination
		}()

		t.Run(testCase.name, func(t *testing.T) {
			go func() { // each Write to the PipeWriter blocks until one or more Reads from the PipeReader fully consume the written data
				switch payload := testCase.data.(type) {
				case Fields:
					switch testCase.style {
					case "WithFields":
						switch testCase.messageLevel {
						case DebugLevel:
							WithFields(payload).Debug(testCase.message)
						case InfoLevel:
							WithFields(payload).Info(testCase.message)
						}
					case "WithFields template":
						switch testCase.messageLevel {
						case DebugLevel:
							WithFields(payload).Debugf(testCase.template, testCase.message)
						case InfoLevel:
							WithFields(payload).Infof(testCase.template, testCase.message)
						}
					case "wrapper":
						switch testCase.messageLevel {
						case DebugLevel:
							Debugw(payload)
						case InfoLevel:
							Infow(payload)
						}
					case "WithTracing":
						switch testCase.messageLevel {
						case DebugLevel:
							WithTracing(testCase.context, payload).Debug(testCase.message)
						case InfoLevel:
							WithTracing(testCase.context, payload).Info(testCase.message)
						}
					case "withTrace":
						switch testCase.messageLevel {
						case DebugLevel:
							DebugWithTracing(testCase.context, payload)
						case InfoLevel:
							InfoWithTracing(testCase.context, payload)
						}
					}

				case string:
					if testCase.template != "" {
						switch testCase.messageLevel {
						case DebugLevel:
							Debugf(testCase.template, payload)
						case InfoLevel:
							Infof(testCase.template, payload)
						}
					} else {
						switch testCase.messageLevel {
						case DebugLevel:
							Debug(payload)
						case InfoLevel:
							Info(payload)
						}
					}
				case []string:
					if testCase.template != "" {
						switch testCase.messageLevel {
						case DebugLevel:
							Debugf(testCase.template, payload[0], payload[1])
						case InfoLevel:
							Infof(testCase.template, payload[0], payload[1])
						}
					} else {
						switch testCase.messageLevel {
						case DebugLevel:
							Debug(payload[0])
						case InfoLevel:
							Info(payload[0])
						}
					}
				}
				w.Close()
			}()

			actualResult.ReadFrom(r) // blocks until data is put into the pipe
			for _, result := range testCase.expectedResult {
				if !strings.Contains(actualResult.String(), string(result)) {
					t.Fatalf("expected %v to contain %v", actualResult, result)
				}
			}
		})
	}
}

func TestCumulativeLogger(t *testing.T) {

	want := []string{`"myField":"foo"`, `"additionalField":"bar"`, `"severity":"info"`}

	ConfigLlog(Options{
		Enabled:    false,
		AppVersion: "test",
		LogLevel:   strings.Title(InfoLevel.String()) + "Level",
	})

	// Pipe creates a synchronous in-memory pipe. It can be used to connect code expecting an io.Reader with code expecting an io.Writer.
	r, w := io.Pipe()
	actualResult := new(bytes.Buffer)
	previousDestination := destination
	destination = w // send log output to the pipe instead of os.Stdout

	defer func() {
		destination = previousDestination
	}()

	go func() { // each Write to the PipeWriter blocks until one or more Reads from the PipeReader fully consume the written data

		cLogger := NewLogger()

		cLogger.Infow(Fields{
			"myField": "foo",
		})

		cLogger.Infow(Fields{
			"additionalField": "bar",
		})

		w.Close()
	}()

	actualResult.ReadFrom(r) // blocks until data is put into the pipe

	for _, result := range want {
		if !strings.Contains(actualResult.String(), string(result)) {
			t.Fatalf("expected %v to contain %v", actualResult, result)
		}
	}
}

func TestLogConcurrently(t *testing.T) {
	ConfigLlog(Options{
		Enabled:    false,
		AppVersion: "test",
		LogLevel:   "InfoLevel",
	})

	payload := Fields{
		"Test": "Foo",
	}

	numTests := 100

	stringChan := make(chan string, numTests)

	t.Run("concurrency test", func(t *testing.T) {
		// test 100 concurrent logs to see if we get a panic
		wg := sync.WaitGroup{}
		for i := 0; i < numTests; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				msg := fmt.Sprintf("log number %d", i)
				WithFields(payload).Info(msg)
				stringChan <- msg
			}(i)

		}
		wg.Wait()

		msgMap := make(map[string]struct{})

		for i := 0; i < numTests; i++ {
			msg := <-stringChan
			msgMap[msg] = struct{}{}
		}

		mapLen := len(msgMap)

		// confirm that no logs have collided and overwritten each other
		if mapLen != numTests {
			t.Fatalf("expected %d logs, but got %d", numTests, mapLen)
		}
	})
}

func BenchmarkLlog(b *testing.B) {
	ConfigLlog(Options{
		Enabled:    false,
		AppVersion: "test",
		LogLevel:   "InfoLevel",
	})

	payload := Fields{
		"Test": "Foo",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		WithFields(payload).Info("benchmark log")
	}
}
