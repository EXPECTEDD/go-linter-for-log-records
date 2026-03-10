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
	log, _ := zap.NewProduction()
	strWorld := "world"

	str := testStruct{str: "world"}

	log.Info("Hello World")        // want "log must start with a lowercase letter"
	log.Warn("Hello" + "World")    // want "log must start with a lowercase letter"
	log.Debug("HELLO WORLD")       // want "log must start with a lowercase letter"
	log.Error("Hello" + strWorld)  // want "log must start with a lowercase letter"
	log.Error("Hello" + "WORLD")   // want "log must start with a lowercase letter"
	log.Debug("HEllo" + "World")   // want "log must start with a lowercase letter"
	log.Info(("Hello World"))      // want "log must start with a lowercase letter"
	log.Info("Hello" + str.str)    // want "log must start with a lowercase letter"
	log.Debug(("Hello") + "world") // want "log must start with a lowercase letter"
}

func SlogLog() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	log.Info("Hello World")      // want "log must start with a lowercase letter"
	log.Warn("Hello", "World")   // want "log must start with a lowercase letter"
	log.Debug("HELLO WORLD")     // want "log must start with a lowercase letter"
	log.Error("Hello" + "world") // want "log must start with a lowercase letter"
	log.Error("Hello" + "WORLD") // want "log must start with a lowercase letter"
	log.Debug("HEllo" + "World") // want "log must start with a lowercase letter"
	log.Info(("Hello World"))    // want "log must start with a lowercase letter"
}
