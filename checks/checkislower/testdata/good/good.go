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

	zapLog.Info("hello World")
	zapLog.Warn("hello" + "World")
	zapLog.Debug("hELLO WORLD")
	zapLog.Error("hello" + "world")
	zapLog.Info("123")
	zapLog.Warn("890" + "HELLO")
	zapLog.Error("hello" + "WORLD")
	zapLog.Debug("hEllo" + "World")
	zapLog.Info(("hello World"))
	zapLog.Warn(strHello + "World")
	zapLog.Debug(strHello + strWorld)
	zapLog.Error(strSome)
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

	slogLog.Info("hello World")
	slogLog.Warn("hello", "World")
	slogLog.Debug("hELLO WORLD")
	slogLog.Error("hello" + "world")
	slogLog.Info("123")
	slogLog.Warn("890", "HELLO")
	slogLog.Error("hello" + "WORLD")
	slogLog.Debug("hEllo" + "World")
	slogLog.Info(("hello World"))
	slogLog.Warn(strHello, "World")
	slogLog.Debug(strHello + strWorld)
	slogLog.Error(strSome)
	slogLog.Info("hello" + str.str)
	slogLog.Warn((str.str) + "hello")
	slogLog.Debug(("hello") + "world")
	slogLog.Error(hello() + "world")
	slogLog.Info("world", hello())
}
