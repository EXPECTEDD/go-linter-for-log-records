package bad

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
)

type testStruct struct {
	str string
}

func hello() string {
	return "hello"
}

func testFn(s string) string {
	return s
}

func ZapLog() {
	zapLog, _ := zap.NewProduction()
	strWorld := "world"

	str := testStruct{str: "world"}

	// Test defaul
	zapLog.Info("Hello World")  // want "log must start with a lowercase letter"
	zapLog.Debug("HELLO WORLD") // want "log must start with a lowercase letter"

	// Test with +
	zapLog.Warn("Hello" + "World")  // want "log must start with a lowercase letter"
	zapLog.Error("Hello" + "WORLD") // want "log must start with a lowercase letter"
	zapLog.Debug("HEllo" + "World") // want "log must start with a lowercase letter"

	// Test with brackets
	zapLog.Info(("Hello World")) // want "log must start with a lowercase letter"

	// Test general
	zapLog.Error("Hello" + strWorld)  // want "log must start with a lowercase letter"
	zapLog.Info("Hello" + str.str)    // want "log must start with a lowercase letter"
	zapLog.Debug(("Hello") + "world") // want "log must start with a lowercase letter"
	zapLog.Warn("HeLLo" + hello())    // want "log must start with a lowercase letter"
}

func SlogLog() {
	slogLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	strWorld := "world"

	str := testStruct{str: "world"}

	// Test defaul
	slogLog.Info("Hello World")  // want "log must start with a lowercase letter"
	slogLog.Debug("HELLO WORLD") // want "log must start with a lowercase letter"

	// Test with +
	slogLog.Warn("Hello" + "World")  // want "log must start with a lowercase letter"
	slogLog.Error("Hello" + "WORLD") // want "log must start with a lowercase letter"
	slogLog.Debug("HEllo" + "World") // want "log must start with a lowercase letter"

	// Test with brackets
	slogLog.Info(("Hello World")) // want "log must start with a lowercase letter"

	// Test with func
	slogLog.Debug(testFn("Hello")) // want "log must start with a lowercase letter"

	// Test with enum
	slogLog.Debug("Hello", "world")                                 // want "log must start with a lowercase letter"
	slogLog.Error("H", "e", "l", "l", "o", "w", "o", "r", "l", "d") // want "log must start with a lowercase letter"

	// Test general
	slogLog.Error("Hello" + strWorld)  // want "log must start with a lowercase letter"
	slogLog.Info("Hello" + str.str)    // want "log must start with a lowercase letter"
	slogLog.Debug(("Hello") + "world") // want "log must start with a lowercase letter"
	slogLog.Warn("HeLLo" + hello())    // want "log must start with a lowercase letter"
}
