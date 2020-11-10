package log

import (
	"fmt"
	"log"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var logger *zap.Logger

// Field zap field
type Field = zap.Field

// Initialize the logger object to avoid the problem of unit test failure when the logger is empty.
func init() {
	logger = zap.NewNop()
}

// SetUp set up log
func SetUp(isProduction bool, fs ...Field) {
	var err error

	if isProduction {
		encoderConf := zap.NewProductionEncoderConfig()
		encoderConf.MessageKey = "message"

		conf := zap.NewProductionConfig()
		conf.EncoderConfig = encoderConf

		logger, err = conf.Build(zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.WarnLevel),
			zap.Fields(fs...))
	} else {
		conf := zap.NewDevelopmentConfig()
		logger, err = conf.Build(zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.WarnLevel))
	}

	if err != nil {
		log.Panicf("can't initialize zap logger: %v", err)
	}
}

// Sync flushes buffer
func Sync() {
	if err := logger.Sync(); err != nil {
		log.Panicf("logger sync failed: %v", err)
	}
}

// Any takes a key and an arbitrary value and chooses the best way to represent
// them as a field, falling back to a reflection-based approach only if
// necessary.
//
// Since byte/uint8 and rune/int32 are aliases, Any can't differentiate between
// them. To minimize surprises, []byte values are treated as binary blobs, byte
// values are treated as uint8, and runes are always treated as integers.
func Any(key string, value interface{}) Field {
	return zap.Any(key, value)
}

// Print calls Output to print to the standard logger.
func Print(v ...interface{}) {
	logger.Debug(fmt.Sprint(v...))
}

// Println calls Output to print to the standard logger.
func Println(v ...interface{}) {
	logger.Debug(fmt.Sprintln(v...))
}

// Printf calls Output to print to the standard logger.
func Printf(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

// Debug logs a message at DebugLevel. The message includes any fields passed
func Debug(msg string, fields ...Field) {
	logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
func Warn(msg string, fields ...Field) {
	logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
func Error(msg string, fields ...Field) {
	logger.Error(msg, fields...)
}

// ErrorE logs a message at ErrorLevel. The message includes any fields passed
func ErrorE(msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

// DPanic logs a message at DPanicLevel. The message includes any fields
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func DPanic(msg string, fields ...Field) {
	logger.DPanic(msg, fields...)
}

// DPanicE logs a message at DPanicLevel. The message includes any fields
func DPanicE(msg string, err error) {
	logger.DPanic(msg, zap.Error(err))
}

// Panic logs a message at PanicLevel. The message includes any fields passed
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...Field) {
	logger.Panic(msg, fields...)
}

// PanicE logs a message at PanicLevel. The message includes any fields passed
//
// The logger then panics, even if logging at PanicLevel is disabled.
func PanicE(msg string, err error) {
	logger.Panic(msg, zap.Error(err))
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...Field) {
	logger.Fatal(msg, fields...)
}

// FatalE logs a message at FatalLevel. The message includes any fields passed
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func FatalE(msg string, err error) {
	logger.Fatal(msg, zap.Error(err))
}
