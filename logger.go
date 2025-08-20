package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log - global o'zgaruvchi (umumiy logger instance)
// Bitta marta yaratiladi va butun loyihada foydalaniladi
var log *zap.Logger

// UseLogger - logger konfiguratsiyasini sozlaydi va ishga tushiradi
func UseLogger() {
	// zap konfiguratsiyasi
	cfg := zap.Config{
		Encoding:         "json",                                  // log format: json yoki console
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel), // minimal level: Info
		OutputPaths:      []string{"stdout", "logger.log"},        // loglar qayerga yoziladi: terminal (stdout) va fayl (logger.log)
		ErrorOutputPaths: []string{"stderr"},                      // error loglar (stderr) ga yoziladi
		Development:      true,                                    // development rejimida qo‘shimcha ma’lumotlar chiqadi

		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "time",    // vaqt maydoni nomi
			LevelKey:   "level",   // log darajasi (INFO, ERROR, WARN)
			MessageKey: "message", // asosiy xabar
			CallerKey:  "caller",  // log yozilgan fayl va qator raqami

			EncodeLevel: zapcore.CapitalLevelEncoder, // INFO, ERROR, WARN kabi katta harflarda
			// vaqtni "odam tushunadigan" formatda chiqarish
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			EncodeCaller: zapcore.FullCallerEncoder, // fayl nomi + qator raqamini qisqa formatda chiqarish
		},
	}

	// config asosida logger yaratish
	logger, err := cfg.Build()
	if err != nil {
		panic(err) // logger yaratishda xato bo‘lsa dastur to‘xtaydi
	}

	// global log o'zgaruvchisini ishga tushirish
	log = logger
}

// Close - log faylni yopadi va xotirani tozalaydi
// Asosan main() ichida defer bilan ishlatiladi
func Close() {
	_ = log.Sync()
}

// Wrapper funksiyalar - log.Info(), log.Error() kabi qisqa qilib ishlatish uchun

// Info darajadagi log yozish
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Error darajadagi log yozish
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// Warn darajadagi log yozish
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

// Debug darajadagi log yozish (faqat Debug mode uchun)
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}
