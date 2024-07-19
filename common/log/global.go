package log

// Info logs a message at InfoLevel.
func Info(msg string, args ...interface{}) {
	loggerInstance.zap.Sugar().Infow(msg, args...)
}

// Debug logs a message at DebugLevel.
func Debug(msg string, args ...interface{}) {
	loggerInstance.zap.Sugar().Debugw(msg, args...)

}

// Warn logs a message at WarnLevel.
func Warn(msg string, args ...interface{}) {
	loggerInstance.zap.Sugar().Warnw(msg, args...)

}

// Error logs a message at ErrorLevel.
func Error(msg string, args ...interface{}) {
	loggerInstance.zap.Sugar().Errorw(msg, args...)

}

// Fatal logs a message at FatalLevel and then exits.
func Fatal(msg string, args ...interface{}) {
	loggerInstance.zap.Sugar().Fatalw(msg, args...)

}

// GetLogger returns the singleton Logger instance.
func GetLogger() *Logger {
	return loggerInstance
}
