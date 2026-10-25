package logger

import (
	"log/slog"
	"os"
)

// Logger é o logger estruturado global da aplicação
var Logger *slog.Logger

// Init inicializa o logger com o nível especificado.
// Suporta os níveis: "debug", "warn", "error". Qualquer outro valor usa "info".
func Init(level string) {
	var logLevel slog.Level

	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
		// Adiciona a linha do arquivo de origem nos logs de debug
		AddSource: logLevel == slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	Logger = slog.New(handler)

	// Define como logger padrão do pacote log
	slog.SetDefault(Logger)
}

// Info loga uma mensagem no nível Info com atributos opcionais
func Info(msg string, args ...any) {
	Logger.Info(msg, args...)
}

// Warn loga uma mensagem no nível Warn com atributos opcionais
func Warn(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

// Error loga uma mensagem no nível Error com atributos opcionais
func Error(msg string, args ...any) {
	Logger.Error(msg, args...)
}

// Debug loga uma mensagem no nível Debug com atributos opcionais
func Debug(msg string, args ...any) {
	Logger.Debug(msg, args...)
}

// Fatal loga uma mensagem de erro e encerra a aplicação
func Fatal(msg string, args ...any) {
	Logger.Error(msg, args...)
	os.Exit(1)
}
