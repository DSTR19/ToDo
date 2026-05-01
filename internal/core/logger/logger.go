package core_logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	file *os.File
}

func NewLogger(config Config) (*Logger, error) {
	// Create a zap logger with a file output
	zaplvl := zap.NewAtomicLevel()
	if err := zaplvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}
	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05.000000")
	logFilePath := filepath.Join(
		config.Folder,
		fmt.Sprintf("%s.log", timestamp),
	)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000")

	zapEncoder := zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zaplvl),
		zapcore.NewCore(zapEncoder, zapcore.AddSync(logFile), zaplvl),
	)

	zapLogger := zap.New(core, zap.AddCaller())
	return &Logger{
		Logger: zapLogger,
		file:   logFile,
	}, nil
}
func FromContext(ctx context.Context) *Logger {
	log, ok := ctx.Value("log").(*Logger)
	if !ok {
		panic("Logger not found in context")
	}
	return log
}
func (l *Logger) With(field ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(field...),
		file:   l.file,
	}
}
func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		fmt.Println("Failed to close application logger:", err)
	}
}
