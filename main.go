package main

import (
	"main/api"
	"main/core"
	"main/processors"
	"net/http"

	env "github.com/Netflix/go-env"
)

func main() {
	var environment core.Env
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		panic(err)
	}

	rtCtx := core.RuntimeContext{
		Processors: []core.Processor{
			processors.Fetch,
			processors.StructuredData,
			processors.Markdown,
			processors.NewSave(environment.DataDir),
		},
		Env: environment,
	}

	r := api.New(&rtCtx)

	err = http.ListenAndServe(environment.HttpHost, r)
	panic(err)
}
