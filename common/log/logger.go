package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const callerSkip = 2

type Logger struct {
	zap *zap.Logger
}

var loggerInstance *Logger

func NewLogger() error {
	logger, err := initZapLogger()
	if err != nil {
		return err
	}
	loggerInstance = &Logger{zap: logger}
	return nil
}

func initZapLogger() (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  CustomLevelEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	level := zap.InfoLevel

	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), level)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(callerSkip)), nil
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// Info logs a message at InfoLevel.
func (l *Logger) Info(msg string, args ...interface{}) {
	l.zap.Sugar().Infow(msg, args...)
}

// Debug logs a message at DebugLevel.
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.zap.Sugar().Debugw(msg, args...)
}

// Warn logs a message at WarnLevel.
func (l *Logger) Warn(msg string, args ...interface{}) {
	l.zap.Sugar().Warnw(msg, args...)
}

// Error logs a message at ErrorLevel.
func (l *Logger) Error(msg string, args ...interface{}) {
	l.zap.Sugar().Errorw(msg, args...)
}

// Fatal logs a message at FatalLevel and then exits.
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.zap.Sugar().Fatalw(msg, args...)
}
