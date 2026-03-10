package bad

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
)

type testStruct struct {
	str string
}

func ZapLog() {
	zapLog, _ := zap.NewProduction()
	strWorld := "world"

	str := testStruct{str: "world"}

	zapLog.Info("Привет Мир")          // want "the log must contain only English characters"
	zapLog.Info("Сәлем Әлем")          // want "the log must contain only English characters"
	zapLog.Warn("你好世界")                // want "the log must contain only English characters"
	zapLog.Warn("Hello" + "Wооrld")    // want "the log must contain only English characters"
	zapLog.Debug("ХELLO WORLD")        // want "the log must contain only English characters"
	zapLog.Error("Heллo!!" + strWorld) // want "the log must contain only English characters"
	zapLog.Error("Hello" + "WOрLD")    // want "the log must contain only English characters"
	zapLog.Debug("HEllo" + "Worlд")    // want "the log must contain only English characters"
	zapLog.Info(("Hello Мир"))         // want "the log must contain only English characters"
	zapLog.Info("Hеllо" + str.str)     // want "the log must contain only English characters"
	zapLog.Debug(("Сәлем") + "world")  // want "the log must contain only English characters"
}

func SlogLog() {
	slogLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	strWorld := "world"

	str := testStruct{str: "world"}

	slogLog.Info("Привет Мир")         // want "the log must contain only English characters"
	slogLog.Info("Сәлем Әлем")         // want "the log must contain only English characters"
	slogLog.Warn("你好世界")               // want "the log must contain only English characters"
	slogLog.Warn("Hello" + "Wооrld")   // want "the log must contain only English characters"
	slogLog.Debug("ХELLO WORLD")       // want "the log must contain only English characters"
	slogLog.Error("Heллo" + strWorld)  // want "the log must contain only English characters"
	slogLog.Error("Hello" + "WOрLD")   // want "the log must contain only English characters"
	slogLog.Debug("HEllo" + "Worlд")   // want "the log must contain only English characters"
	slogLog.Info(("Hello Мир"))        // want "the log must contain only English characters"
	slogLog.Info("Hеllо" + str.str)    // want "the log must contain only English characters"
	slogLog.Debug(("Сәлем") + "world") // want "the log must contain only English characters"
}
