package logger

import (
	"context"
	"log/slog"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Logger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)

	With(...any) Logger
	WithLevel(LogLevel) Logger
}

type logger struct {
	slog *slog.Logger
}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) WithLevel(level LogLevel) Logger {
	var lvl slog.Level
	switch level {
	case DebugLevel:
		lvl = slog.LevelDebug
	case InfoLevel:
		lvl = slog.LevelInfo
	case WarnLevel:
		lvl = slog.LevelWarn
	case ErrorLevel:
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	l.slog.WithGroup("level").With("level", lvl.String())

	return l
}

func (l *logger) Info(ctx context.Context, msg string, args ...any) {
	l.slog.InfoContext(ctx, msg, args...)
}

func (l *logger) Error(ctx context.Context, msg string, args ...any) {
	l.slog.ErrorContext(ctx, msg, args...)
}

func (l *logger) Debug(ctx context.Context, msg string, args ...any) {
	l.slog.DebugContext(ctx, msg, args...)
}

func (l *logger) Warn(ctx context.Context, msg string, args ...any) {
	l.slog.WarnContext(ctx, msg, args...)
}

func (l *logger) With(args ...any) Logger {
	return &logger{
		slog: l.slog.With(args...),
	}
}
