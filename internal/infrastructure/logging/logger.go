package logging

import (
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"log"
	"time"
)

type logger struct {
	zap *zap.SugaredLogger
}

func GetLogger() LoggerInterface {
	zapLogger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	initSentry()

	return logger{zap: zapLogger.Sugar()}
}

func (l logger) Debug(msg string, fields ...interface{}) {
	sendToSentry(sentry.LevelDebug, msg, fields)

	l.zap.Debug(msg, fields)
}

func (l logger) Info(msg string, fields ...interface{}) {
	sendToSentry(sentry.LevelInfo, msg, fields)

	l.zap.Info(msg, fields)
}

func (l logger) Warn(msg string, fields ...interface{}) {
	sendToSentry(sentry.LevelWarning, msg, fields)

	l.zap.Warn(msg, fields)
}

func (l logger) Error(msg string, fields ...interface{}) {
	sendToSentry(sentry.LevelError, msg, fields)

	l.zap.Error(msg, fields)
}

func (l logger) Panic(msg string, fields ...interface{}) {
	sendToSentry(sentry.LevelFatal, msg, fields)

	l.zap.Panic(msg, fields)
}

func (l logger) Fatal(msg string, fields ...interface{}) {
	sendToSentry(sentry.LevelFatal, msg, fields)

	l.zap.Fatal(msg, fields)
}

func (l logger) Fatalf(text string, v ...interface{}) {
	sendToSentry(sentry.LevelFatal, text, v)

	l.zap.Fatalf(text, v)
}

func sendToSentry(errorLevel sentry.Level, text string, v ...interface{}) {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetLevel(errorLevel)
		scope.SetContext("additional", v)
	})

	sentry.CaptureMessage(text)
	sentry.Flush(time.Minute)
}

func (l logger) Close() {
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	sentry.Flush(time.Second * 2)
	l.zap.Sync()
}

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "https://1e8c898aac7c45259639d9a6eae5a926@o1210124.ingest.sentry.io/6345772", // @TODO
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: "", // @TODO
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true, // @TODO
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
