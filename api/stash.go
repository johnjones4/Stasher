package api

import (
	_ "embed"
	"errors"
	"html/template"
	"main/core"
	"net/http"
)

//go:embed response.html
var templateStr string

func (a *API) stash(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("key") != a.runtime.Env.StashKey {
		a.errorResponse(w, http.StatusForbidden, errors.New("access denied"))
		return
	}

	item, err := a.process(r.URL.Query().Get("url"))
	if err != nil {
		a.errorResponse(w, http.StatusBadRequest, err)
		return
	}

	temp, err := template.New("t").Parse(templateStr)
	if err != nil {
		a.errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", core.ContentTypeHTML)

	type templateVars struct {
		Title string
		Body  string
	}

	temp.Execute(w, templateVars{
		Title: item.Name,
		Body:  string(item.Body),
	})
}
