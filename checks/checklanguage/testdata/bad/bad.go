package bad

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
)

type testStruct struct {
	str string
}

func world() string {
	return "world"
}

func ZapLog() {
	zapLog, _ := zap.NewProduction()
	strWorld := "world"

	str := testStruct{str: "world"}

	// Test defaul
	zapLog.Info("Привет Мир")   // want "the log must contain only English characters"
	zapLog.Info("Сәлем Әлем")   // want "the log must contain only English characters"
	zapLog.Debug("ХELLO WORLD") // want "the log must contain only English characters"
	zapLog.Warn("你好世界")         // want "the log must contain only English characters"

	// Test with +
	zapLog.Warn("Hello" + "Wооrld") // want "the log must contain only English characters"
	zapLog.Error("Hello" + "WOрLD") // want "the log must contain only English characters"
	zapLog.Debug("HEllo" + "Worlд") // want "the log must contain only English characters"

	// Test with brackets
	zapLog.Info(("Hello Мир")) // want "the log must contain only English characters"

	// Test general
	zapLog.Error("Heллo!!" + strWorld) // want "the log must contain only English characters"
	zapLog.Info("Hеllо" + str.str)     // want "the log must contain only English characters"
	zapLog.Debug(("Сәлем") + "world")  // want "the log must contain only English characters"
	zapLog.Warn("磷ривет" + world())    // want "the log must contain only English characters"
}

func SlogLog() {
	slogLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	strWorld := "world"

	str := testStruct{str: "world"}

	// Test defaul
	slogLog.Info("Привет Мир")   // want "the log must contain only English characters"
	slogLog.Info("Сәлем Әлем")   // want "the log must contain only English characters"
	slogLog.Debug("ХELLO WORLD") // want "the log must contain only English characters"
	slogLog.Warn("你好世界")         // want "the log must contain only English characters"

	// Test with +
	slogLog.Warn("Hello" + "Wооrld") // want "the log must contain only English characters"
	slogLog.Error("Hello" + "WOрLD") // want "the log must contain only English characters"
	slogLog.Debug("HEllo" + "Worlд") // want "the log must contain only English characters"

	// Test with brackets
	slogLog.Info(("Hello Мир")) // want "the log must contain only English characters"

	// Test with enum
	slogLog.Debug("He我我o", "World")                                // want "the log must contain only English characters"
	slogLog.Warn("H", "e", "我", "我", "o", "W", "o", "r", "l", "d") // want "the log must contain only English characters"

	// Test general
	slogLog.Error("Heллo!!" + strWorld) // want "the log must contain only English characters"
	slogLog.Info("Hеllо" + str.str)     // want "the log must contain only English characters"
	slogLog.Debug(("Сәлем") + "world")  // want "the log must contain only English characters"
	slogLog.Warn("磷ривет" + world())    // want "the log must contain only English characters"
}
