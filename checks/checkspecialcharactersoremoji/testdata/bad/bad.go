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

func testFunc(s string) string {
	return s
}

func ZapLog() {
	zapLog, _ := zap.NewProduction()
	strWorld := "world"

	str := testStruct{str: "world"}

	// TEST SPECIAL CHARACTERS

	// Test defaul
	zapLog.Info("Hello World!")  // want "the log contains special character"
	zapLog.Debug("HELLO, WORLD") // want "the log contains special character"

	// Test with +
	zapLog.Warn("Hello" + "," + "World")  // want "the log contains special character"
	zapLog.Error("Hello" + "WORLD" + "!") // want "the log contains special character"
	zapLog.Debug("!HEllo" + "World")      // want "the log contains special character"

	// Test with brackets
	zapLog.Info(("Hello World...")) // want "the log contains special character"

	// Test with func
	zapLog.Warn(testFunc("He!!o")) // want "the log contains special character"

	// Test general
	zapLog.Error("Hello," + strWorld) // want "the log contains special character"
	zapLog.Info("Hello!!!" + str.str) // want "the log contains special character"
	zapLog.Debug(("Hello") + "w#rld") // want "the log contains special character"
	zapLog.Warn("$eLLo" + hello())    // want "the log contains special character"

	// TEST EMOJI

	// Test defaul
	zapLog.Info("Hello World🔥")  // want "the log contains emoji"
	zapLog.Debug("HELLO🫣 WORLD") // want "the log contains emoji"

	// Test with +
	zapLog.Warn("Hello" + "😬" + "World")     // want "the log contains emoji"
	zapLog.Error("Hello" + "WORLD" + "😶‍🌫️") // want "the log contains emoji"
	zapLog.Debug("👾HEllo" + "World")         // want "the log contains emoji"

	// Test with brackets
	zapLog.Info(("Hello World😈😈😈")) // want "the log contains emoji"

	// Test with func
	zapLog.Warn(testFunc("He🧐🧐o")) // want "the log contains emoji"

	// Test general
	zapLog.Error("Hello🥳" + strWorld) // want "the log contains emoji"
	zapLog.Info("Hello🤯🤯🤯" + str.str) // want "the log contains emoji"
	zapLog.Debug(("Hello") + "w🥴rld") // want "the log contains emoji"
	zapLog.Warn("😸eLLo" + hello())    // want "the log contains emoji"
}

func SlogLog() {
	slogLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	strWorld := "world"

	str := testStruct{str: "world"}

	// TEST SPECIAL CHARACTERS

	// Test defaul
	slogLog.Info("Hello World!")  // want "the log contains special character"
	slogLog.Debug("HELLO, WORLD") // want "the log contains special character"

	// Test with +
	slogLog.Warn("Hello" + "," + "World")  // want "the log contains special character"
	slogLog.Error("Hello" + "WORLD" + "!") // want "the log contains special character"
	slogLog.Debug("!HEllo" + "World")      // want "the log contains special character"

	// Test with brackets
	slogLog.Info(("Hello World...")) // want "the log contains special character"

	// Test with func
	slogLog.Warn(testFunc("He!!o")) // want "the log contains special character"

	// Test with enum
	slogLog.Debug("Hello", ",", "World") // want "the log contains special character"

	// Test general
	slogLog.Error("Hello," + strWorld) // want "the log contains special character"
	slogLog.Info("Hello!!!" + str.str) // want "the log contains special character"
	slogLog.Debug(("Hello") + "w#rld") // want "the log contains special character"
	slogLog.Warn("$eLLo" + hello())    // want "the log contains special character"

	// TEST EMOJI

	// Test defaul
	slogLog.Info("Hello World🔥")  // want "the log contains emoji"
	slogLog.Debug("HELLO🫣 WORLD") // want "the log contains emoji"

	// Test with +
	slogLog.Warn("Hello" + "😬" + "World")     // want "the log contains emoji"
	slogLog.Error("Hello" + "WORLD" + "😶‍🌫️") // want "the log contains emoji"
	slogLog.Debug("👾HEllo" + "World")         // want "the log contains emoji"

	// Test with brackets
	slogLog.Info(("Hello World😈😈😈")) // want "the log contains emoji"

	// Test with func
	slogLog.Warn(testFunc("He🧐🧐o")) // want "the log contains emoji"

	// Test with enum
	slogLog.Debug("Hello", "😡", "World") // want "the log contains emoji"

	// Test general
	slogLog.Error("Hello🥳" + strWorld) // want "the log contains emoji"
	slogLog.Info("Hello🤯🤯🤯" + str.str) // want "the log contains emoji"
	slogLog.Debug(("Hello") + "w🥴rld") // want "the log contains emoji"
	slogLog.Warn("😸eLLo" + hello())    // want "the log contains emoji"
}
