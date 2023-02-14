package main

import (
	"main/api"
	"main/core"
	"main/processors"
	"net/http"
	"os"
)

func main() {
	rtCtx := core.RuntimeContext{
		Processors: []core.Processor{
			processors.Fetch,
			processors.StructuredData,
			processors.Markdown,
			processors.NewSave(os.Getenv("DATA_DIR")),
		},
		StashKey: os.Getenv("STASH_KEY"),
	}

	r := api.New(&rtCtx)

	err := http.ListenAndServe(os.Getenv("HTTP_HOST"), r)
	panic(err)
}
