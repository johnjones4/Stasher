package main

import (
	"main/api"
	"main/core"
	"main/processors"
	"net/http"
	"os"
	"strconv"

	env "github.com/Netflix/go-env"
	"github.com/johnjones4/errorbot"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	chatId, err := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID"))
	if err != nil {
		panic(err)
	}
	bot := errorbot.New(
		"stasher",
		os.Getenv("TELEGRAM_TOKEN"),
		chatId,
	)

	config := zap.NewDevelopmentConfig()
	l, err := config.Build(zap.Hooks(bot.ZapHook([]zapcore.Level{
		zapcore.FatalLevel,
		zapcore.PanicLevel,
		zapcore.DPanicLevel,
		zapcore.ErrorLevel,
		zapcore.WarnLevel,
	})))
	if err != nil {
		panic(err)
	}

	defer l.Sync()
	log := l.Sugar()

	var environment core.Env
	_, err = env.UnmarshalFromEnviron(&environment)
	if err != nil {
		panic(err)
	}

	rtCtx := core.RuntimeContext{
		Processors: []core.Processor{
			processors.Fetch,
			processors.StructuredData,
			processors.HTMLContent,
			processors.Markdown,
			processors.NewSave(environment.DataDir),
		},
		Env: environment,
	}

	r := api.New(&rtCtx, log)

	err = http.ListenAndServe(environment.HttpHost, r)
	panic(err)
}
