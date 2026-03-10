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

	zapLog.Info("Hello World")        // want "log must start with a lowercase letter"
	zapLog.Warn("Hello" + "World")    // want "log must start with a lowercase letter"
	zapLog.Debug("HELLO WORLD")       // want "log must start with a lowercase letter"
	zapLog.Error("Hello" + strWorld)  // want "log must start with a lowercase letter"
	zapLog.Error("Hello" + "WORLD")   // want "log must start with a lowercase letter"
	zapLog.Debug("HEllo" + "World")   // want "log must start with a lowercase letter"
	zapLog.Info(("Hello World"))      // want "log must start with a lowercase letter"
	zapLog.Info("Hello" + str.str)    // want "log must start with a lowercase letter"
	zapLog.Debug(("Hello") + "world") // want "log must start with a lowercase letter"
}

func SlogLog() {
	slogLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	str := testStruct{str: "world"}

	slogLog.Info("Hello World")        // want "log must start with a lowercase letter"
	slogLog.Warn("Hello", "World")     // want "log must start with a lowercase letter"
	slogLog.Debug("HELLO WORLD")       // want "log must start with a lowercase letter"
	slogLog.Error("Hello" + "world")   // want "log must start with a lowercase letter"
	slogLog.Error("Hello" + "WORLD")   // want "log must start with a lowercase letter"
	slogLog.Debug("HEllo" + "World")   // want "log must start with a lowercase letter"
	slogLog.Info(("Hello World"))      // want "log must start with a lowercase letter"
	slogLog.Info("Hello" + str.str)    // want "log must start with a lowercase letter"
	slogLog.Debug(("Hello") + "world") // want "log must start with a lowercase letter"
}
