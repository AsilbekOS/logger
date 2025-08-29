package logger

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger - umumiy log instance ni o'rab turuvchi struct.
// Bu orqali Info(), Error(), Warn(), Debug() kabi metodlar chaqiriladi.
type Logger struct {
	instance *zap.Logger
}

var once sync.Once // sync.Once loggerni faqat bir marta yaratish uchun ishlatiladi
var logger *Logger // global logger instance

// UseLogger - Loyihada yagona (singleton) logger yaratadi va qaytaradi.
//
// Parametrlar:
//
//	encFormat (string) - log formatini tanlash: "json" yoki "console".
//	    "json"   -> loglar JSON formatida yoziladi (odatda production uchun)
//	    "console"-> loglar plain text formatda yoziladi (odatda development uchun)
//
//	logFile (string) - log yoziladigan fayl nomi (masalan: "logger.log").
//
//	outputTerminal (bool) - log terminalga chiqsinmi yoki faqat faylga yozilsinmi.
//	    true  -> loglar terminal (stdout) va faylga yoziladi
//	    false -> loglar faqat faylga yoziladi
//
//	level (string) - log darajasi.
//	    "debug" -> barcha loglar (Debug ham) chiqadi
//	    "info"  -> faqat Info va undan yuqori loglar chiqadi
//	    "warn"  -> faqat Warning va Error loglar chiqadi
//	    "error" -> faqat Error loglar chiqadi
//
// Qaytaradi:
//
//	*Logger - global logger instance, undan Info(), Error(), Warn(), Debug() metodlarini ishlatish mumkin.
//
// Misol:
//
//	log := logger.UseLogger("json", "logger.log", true, "debug")
//	log.Info("Server started")
//	log.Debug("Config loaded", zap.String("file", "config.yaml"))
func UseLogger(encFormat, logFile string, outputTerminal bool, level string) *Logger {
	// log formatni tanlash
	var encoding string
	switch encFormat {
	case "console":
		encoding = "console"
	case "json":
		encoding = "json"
	default:
		encoding = "json" // default format
	}

	// log darajasini tanlash
	var logLevel zapcore.Level
	switch level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel // default daraja
	}

	// chiqish manzillarini sozlash
	var outputs []string
	if outputTerminal {
		// stdout va faylga yozish
		outputs = []string{"stdout", logFile}
	} else {
		// faqat faylga yozish
		outputs = []string{logFile}
	}

	// loggerni faqat bir marta yaratish
	once.Do(func() {
		cfg := zap.Config{
			Encoding:         encoding,
			Level:            zap.NewAtomicLevelAt(logLevel),
			OutputPaths:      outputs,
			ErrorOutputPaths: []string{"stderr"},
			Development:      true, // development rejimi yoqilgan
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:     "time",    // vaqt maydoni
				LevelKey:    "level",   // log darajasi
				MessageKey:  "message", // asosiy xabar
				CallerKey:   "caller",  // fayl va qator raqami
				EncodeLevel: zapcore.CapitalLevelEncoder,
				EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
					enc.AppendString(t.Format("2006-01-02 15:04:05"))
				},
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		// config asosida logger yaratish
		l, err := cfg.Build(zap.AddCaller())
		if err != nil {
			panic(err)
		}
		logger = &Logger{instance: l}
	})
	return logger
}

// Close - loggerni yopadi va xotirani tozalaydi.
// Asosan main() ichida defer bilan ishlatiladi.
//
//	Misol: defer logger.Close()
func (l *Logger) Close() {
	_ = l.instance.Sync()
}

// Info - Info darajadagi log yozadi.
// Qo'shimcha maydonlarni ham qo'shish mumkin.
//
//	Misol:
//	  logger.Info("User logged in",
//	    zap.String("username", "johndoe"),
//	    zap.Int("user_id", 42),
//	  )
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.instance.Info(msg, fields...)
}

// Error - Error darajadagi log yozadi.
// Odatda xatoliklar uchun ishlatiladi.
//
//	Misol:
//	  logger.Error("Failed to process request",
//	    zap.String("request_id", "12345"),
//	    zap.Error(err),
//	  )
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.instance.Error(msg, fields...)
}

// Warn - Warning darajadagi log yozadi.
// Ogohlantirishlar uchun ishlatiladi.
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.instance.Warn(msg, fields...)
}

// Debug - Debug darajadagi log yozadi.
// Faqat Debug level yoqilganda ishlaydi.
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.instance.Debug(msg, fields...)
}
