package good

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

func zapLog() {
	zapLog, _ := zap.NewProduction()
	strHello := "hello"
	strWorld := "world"
	strSome := "So123MeSTR"

	str := testStruct{str: "hello"}

	// Test defaul
	zapLog.Info("hello World")
	zapLog.Debug("hELLO WORLD")

	// Test with numbers
	zapLog.Info("123")
	zapLog.Warn("890" + "HELLO")

	// Test with +
	zapLog.Warn("hello" + "World")
	zapLog.Error("hello" + "world")
	zapLog.Error("hello" + "WORLD")
	zapLog.Debug("hEllo" + "World")

	// Test with brackets
	zapLog.Info(("hello World"))

	// Test with struct
	zapLog.Info(str.str)

	// Test with var
	zapLog.Debug(strHello)
	zapLog.Error(strSome)

	// Test with func
	zapLog.Warn(hello())

	// Test general
	zapLog.Warn(strHello + "World")
	zapLog.Debug(strHello + strWorld)
	zapLog.Info("hello" + str.str)
	zapLog.Warn((str.str) + "hello")
	zapLog.Debug(("hello") + "world")
	zapLog.Error(hello() + "world")
	zapLog.Info("world" + hello())
}

func slogLog() {
	slogLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	strHello := "hello"
	strWorld := "world"
	strSome := "So123MeSTR"

	str := testStruct{str: "hello"}

	// Test defaul
	slogLog.Info("hello World")
	slogLog.Debug("hELLO WORLD")

	// Test with numbers
	slogLog.Info("123")
	slogLog.Warn("890" + "HELLO")

	// Test with +
	slogLog.Warn("hello" + "World")
	slogLog.Error("hello" + "world")
	slogLog.Error("hello" + "WORLD")
	slogLog.Debug("hEllo" + "World")

	// Test with brackets
	slogLog.Info(("hello World"))

	// Test with struct
	slogLog.Info(str.str)

	// Test with var
	slogLog.Debug(strHello)
	slogLog.Error(strSome)

	// Test with func
	slogLog.Warn(hello())

	// Test enum
	slogLog.Info("hello", "world")
	slogLog.Error("he", "llo", "wor", "ld")

	// Test general
	slogLog.Warn(strHello + "World")
	slogLog.Debug(strHello + strWorld)
	slogLog.Info("hello", str.str)
	slogLog.Warn((str.str) + "hello")
	slogLog.Debug(("hello") + "world")
	slogLog.Error(hello(), "world")
	slogLog.Info("world" + hello())
}
