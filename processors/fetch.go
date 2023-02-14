package processors

import (
	"fmt"
	"io"
	"main/core"
	"net/http"
	"time"
)

func Fetch(req *core.Item) error {
	urlStr := req.URL.String()
	if len(req.Body) > 0 || urlStr == "" {
		return nil
	}

	res, err := http.Get(urlStr)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	req.ContentType = res.Header.Get("Content-Type")
	req.Body = body
	req.Name = fmt.Sprintf("%s%s [%s]", req.URL.Host, req.URL.Path, time.Now().Format(time.Stamp))

	return nil
}
