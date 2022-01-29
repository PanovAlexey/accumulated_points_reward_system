package logging

import (
	"github.com/PanovAlexey/accumulated_points_reward_system/config"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"log"
	"time"
)

type logger struct {
	zap *zap.SugaredLogger
}

func GetLogger(config config.Config) LoggerInterface {
	zapLogger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	initSentry(config.GetAppEnvironment(), config.GetAppLoggerDsn(), config.IsAppDebugMode())

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

func initSentry(environment, dsn string, isDebug bool) {
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: environment,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: isDebug,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
