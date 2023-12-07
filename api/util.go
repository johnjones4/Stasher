package api

import (
	"encoding/json"
	"main/core"
	"net/http"
	"net/url"
)

func (a *API) jsonResponse(w http.ResponseWriter, status int, info any) {
	bytes, err := json.Marshal(info)
	if err != nil {
		a.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}

func (a *API) errorResponse(w http.ResponseWriter, status int, err error) {
	a.log.Error(err)
	msg := map[string]any{
		"ok":      false,
		"message": err.Error(),
	}
	a.jsonResponse(w, status, msg)
}

func (a *API) process(urlStr string) (core.Item, error) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return core.Item{}, err
	}

	item := core.Item{
		URL:  parsedUrl,
		Info: make(map[string]any),
	}

	for _, processor := range a.runtime.Processors {
		err = processor(&item)
		if err != nil {
			return core.Item{}, err
		}
	}

	return item, nil
}
