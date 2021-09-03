package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

//CreateZapLogger creates a new zap logger
func CreateZapLogger() (*zap.Logger, func()) {
	level := zap.DebugLevel

	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,

			TimeKey:    "@timestamp",
			EncodeTime: zapcore.RFC3339TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,

			StacktraceKey: "stacktrace",
		},
	}
	logger, err := cfg.Build(zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	logger.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel))

	return logger, func() {
		// it is here for not to forget to flush logs on server shutdown
		_ = logger.Sync()
	}
}
