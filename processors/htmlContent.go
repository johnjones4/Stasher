package processors

import (
	"bytes"
	"main/core"

	"github.com/PuerkitoBio/goquery"
)

const (
	htmlContentKey = "html"
)

var (
	selectors = []string{
		"[role=\"main\"]",
		"article",
		"[id=\"content\"]",
		"[id=\"article\"]",
		"body",
	}
)

func HTMLContent(req *core.Item) error {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(req.Body))
	if err != nil {
		return err
	}

	for _, selector := range selectors {
		sel := doc.Find(selector)
		if sel != nil && len(sel.Nodes) > 0 {
			req.Info[htmlContentKey], err = sel.Html()
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}
