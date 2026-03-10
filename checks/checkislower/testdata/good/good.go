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
	log, _ := zap.NewProduction()
	strHello := "hello"
	strWorld := "world"
	strSome := "So123MeSTR"

	str := testStruct{str: "hello"}

	log.Info("hello World")
	log.Warn("hello" + "World")
	log.Debug("hELLO WORLD")
	log.Error("hello" + "world")
	log.Info("123")
	log.Warn("890" + "HELLO")
	log.Error("hello" + "WORLD")
	log.Debug("hEllo" + "World")
	log.Info(("hello World"))
	log.Warn(strHello + "World")
	log.Debug(strHello + strWorld)
	log.Error(strSome)
	log.Info("hello" + str.str)
	log.Warn((str.str) + "hello")
	log.Debug(("hello") + "world")
	log.Error(hello() + "world")
	log.Info("world" + hello())
}

func slogLog() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	strHello := "hello"
	strWorld := "world"
	strSome := "So123MeSTR"

	str := testStruct{str: "hello"}

	log.Info("hello World")
	log.Warn("hello", "World")
	log.Debug("hELLO WORLD")
	log.Error("hello" + "world")
	log.Info("123")
	log.Warn("890", "HELLO")
	log.Error("hello" + "WORLD")
	log.Debug("hEllo" + "World")
	log.Info(("hello World"))
	log.Warn(strHello, "World")
	log.Debug(strHello + strWorld)
	log.Error(strSome)
	log.Info("hello" + str.str)
	log.Warn((str.str) + "hello")
	log.Debug(("hello") + "world")
	log.Error(hello() + "world")
	log.Info("world", hello())
}
