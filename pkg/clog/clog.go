package clog

import (
	"context"
	"log"
	"path"
	"runtime"
	"strconv"
)

type contextKey string

const (
	traceIDKey = contextKey("trace_id")
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

var LevelString = map[int]string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO",
	LevelWarning: "WARNING",
	LevelError:   "ERROR",
	LevelFatal:   "FATAL",
}

var (
	minLevel int
)

func SetMinLogLevel(level int) {
	minLevel = level
}

func ContextWithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func Info(ctx context.Context, msg string) {
	if minLevel > LevelInfo {
		return
	}

	log.Println(formatLogger(ctx, LevelInfo, msg))
}

func Error(ctx context.Context, msg string) {
	if minLevel > LevelError {
		return
	}

	log.Println(formatLogger(ctx, LevelError, msg))
}

func Fatal(ctx context.Context, msg string) {
	if minLevel < LevelFatal {
		return
	}

	log.Println(formatLogger(ctx, LevelFatal, msg))
}

func formatLogger(ctx context.Context, logLevel int, msg string) string {
	traceID := getTraceIDFromContext(ctx)

	fileLine := getActualCaller(3)

	logMsg := traceID + "|" + fileLine + "|" + msg

	level, ok := LevelString[logLevel]
	if ok {
		logMsg = level + "|" + logMsg
	}

	return logMsg
}

func getTraceIDFromContext(ctx context.Context) string {
	traceID := ctx.Value(traceIDKey)
	if traceID == nil {
		return ""
	}

	return traceID.(string)
}

func getActualCaller(skip int) (fileLine string) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	fileLine = path.Base(file) + ":" + strconv.Itoa(line)

	return fileLine
}
