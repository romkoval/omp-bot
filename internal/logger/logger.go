package logger

import (
	"context"
	"os"
	"sync"

	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/omp-bot/internal/config"
	gelf "github.com/snovichkov/zap-gelf"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey struct{}

var attachedLoggerKey = &ctxKey{}

var globalLogger *zap.SugaredLogger
var gloggerMux sync.Mutex

func fromContext(ctx context.Context) *zap.SugaredLogger {
	var logger *zap.SugaredLogger
	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		logger = attachedLogger
	} else {
		// for testing
		gloggerMux.Lock()
		defer gloggerMux.Unlock()
		if globalLogger == nil {
			logger, _ := zap.NewDevelopment()
			globalLogger = logger.Sugar()
		}

		logger = globalLogger
	}

	jaegerSpan := opentracing.SpanFromContext(ctx)
	if jaegerSpan != nil {
		if spanCtx, ok := opentracing.SpanFromContext(ctx).Context().(jaeger.SpanContext); ok {
			logger = logger.With("trace-id", spanCtx.TraceID())
		}
	}

	return logger

}

func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Errorw(message, kvs...)
}

func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Warnw(message, kvs...)
}

func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Infow(message, kvs...)
}

func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Debugw(message, kvs...)
}

func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Fatalw(message, kvs...)
}

func AttachLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, logger)
}

func CloneWithLevel(ctx context.Context, level zapcore.Level) *zap.SugaredLogger {
	return fromContext(ctx).Desugar().WithOptions(WithLevel(level)).Sugar()
}

func SetLogger(newLogger *zap.SugaredLogger) {
	globalLogger = newLogger
}

func InitLogger(ctx context.Context, cfg config.Config, serviceName string) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		FatalKV(ctx, "gelfCore setup error", "err", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore), zap.AddCaller(), zap.AddCallerSkip(1))

	sugaredLogger := notSugaredLogger.Sugar()
	SetLogger(sugaredLogger.With(
		"service", serviceName,
	))

	return func() {
		notSugaredLogger.Sync()
	}
}
