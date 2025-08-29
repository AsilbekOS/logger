# Logger (Zap Wrapper)

Go (Golang) uchun oddiy va qulay `zap` kutubxonasiga o‘rab yozilgan **logger**.  
Loyihada yagona (singleton) instance yaratadi va `Info()`, `Error()`, `Warn()`, `Debug()` kabi metodlar orqali log yozishni osonlashtiradi.

---

## 🚀 O‘rnatish

```bash
go get go.uber.org/zap
```

Logger paketini o‘z loyihangga import qil:
```go
import "your_project/logger"
```

---

## ⚙️ Foydalanish

### 1. Loggerni ishga tushirish

```go
log := logger.UseLogger(
    "json",        // log format: "json" yoki "console"
    "logger.log",  // log fayl nomi
    true,          // outputTerminal: true -> terminal + fayl, false -> faqat fayl
    "debug",       // level: "debug", "info", "warn", "error"
)
defer log.Close()
```

---

### 2. Log yozish

```go
log.Info("Server started")

log.Debug("Config loaded",
    zap.String("file", "config.yaml"),
    zap.Int("version", 2),
)

log.Warn("Memory usage is high",
    zap.Int("percent", 85),
)

log.Error("Failed to connect to database",
    zap.String("db", "postgres"),
    zap.Error(err),
)
```

---

## 📖 Parametrlar

| Parametr         | Turi    | Tavsifi |
|------------------|---------|---------|
| `encFormat`      | string  | `"json"` yoki `"console"` format |
| `logFile`        | string  | log yoziladigan fayl nomi, masalan: `logger.log` |
| `outputTerminal` | bool    | `true` → log terminal (stdout) va faylga yoziladi <br>`false` → faqat faylga yoziladi |
| `level`          | string  | `"debug"`, `"info"`, `"warn"`, `"error"` log darajasi |

---

## 📊 Log darajalari

- `debug` → barcha loglar (Debug ham) chiqadi  
- `info` → faqat Info va undan yuqori loglar chiqadi  
- `warn` → faqat Warning va Error loglar chiqadi  
- `error` → faqat Error loglar chiqadi  

---

## 🔥 Misol natija (console format)

```text
2025-08-29 18:55:22  INFO    Server started          caller=main.go:15
2025-08-29 18:55:22  DEBUG   Config loaded           caller=main.go:16 file=config.yaml version=2
2025-08-29 18:55:22  WARN    Memory usage is high    caller=main.go:17 percent=85
2025-08-29 18:55:22  ERROR   Failed to connect to database caller=main.go:18 db=postgres error=dial tcp 127.0.0.1:5432: connect: connection refused
```

---

## 📌 Eslatma

- **`defer log.Close()`** chaqirish muhim, aks holda log fayl to‘liq yozilmasligi mumkin.  
- Default format `"json"`, default level esa `"info"`.  
- Bitta projectda faqat **bitta logger instance** yaratiladi (`sync.Once`).  

---

👨‍💻 Author: [Asilbek Xolmatov](https://github.com/AsilbekOS)  