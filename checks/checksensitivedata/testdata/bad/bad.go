package bad

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
)

type testStruct struct {
	token string
}

func testFn(s string) string {
	return s
}

func ZapLog() {
	zapLog, _ := zap.NewProduction()

	token := "token"

	tokenStr := testStruct{token: "token"}

	// Test with +
	zapLog.Info("token:" + token)      // want "the log contains sensitive data"
	zapLog.Debug(token + ":token")     // want "the log contains sensitive data"
	zapLog.Warn("token" + ":" + token) // want "the log contains sensitive data"

	// Test with brackets
	zapLog.Debug((token)) // want "the log contains sensitive data"

	// Test with func
	zapLog.Error(testFn(token)) // want "the log contains sensitive data"

	// Test with struct
	zapLog.Info(tokenStr.token) // want "the log contains sensitive data"

	// Test general
	zapLog.Error(("token:" + token))              // want "the log contains sensitive data"
	zapLog.Error("token:" + testFn(token))        // want "the log contains sensitive data"
	zapLog.Debug((testFn(token)) + ":" + "token") // want "the log contains sensitive data"
	zapLog.Warn("token:" + (tokenStr.token))      // want "the log contains sensitive data"
}

func SlogLog() {
	slogLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	token := "token"

	tokenStr := testStruct{token: "token"}

	// Test with +
	slogLog.Info("token:" + token)      // want "the log contains sensitive data"
	slogLog.Debug(token + ":token")     // want "the log contains sensitive data"
	slogLog.Warn("token" + ":" + token) // want "the log contains sensitive data"

	// Test with brackets
	slogLog.Debug((token)) // want "the log contains sensitive data"

	// Test with func
	slogLog.Error(testFn(token)) // want "the log contains sensitive data"

	// Test with struct
	slogLog.Info(tokenStr.token) // want "the log contains sensitive data"

	// Test with enum
	slogLog.Warn("token:", token)  // want "the log contains sensitive data"
	slogLog.Debug(token, ":token") // want "the log contains sensitive data"

	// Test general
	slogLog.Error(("token:" + token))              // want "the log contains sensitive data"
	slogLog.Error("token:" + testFn(token))        // want "the log contains sensitive data"
	slogLog.Debug((testFn(token)) + ":" + "token") // want "the log contains sensitive data"
	slogLog.Warn("token:" + (tokenStr.token))      // want "the log contains sensitive data"
	slogLog.Info("token:", ("token")+token)        // want "the log contains sensitive data"
}
