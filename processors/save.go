package processors

import (
	"main/core"
	"os"
	"path"
	"strings"

	"github.com/flytam/filenamify"
)

func NewSave(outputDir string) core.Processor {
	return func(req *core.Item) error {
		urlStr := req.URL.String()
		if len(req.Body) == 0 || urlStr == "" {
			return nil
		}

		cleanName, err := filenamify.Filenamify(req.Name, filenamify.Options{})
		if err != nil {
			return err
		}

		outPath := path.Join(outputDir, cleanName+"."+extensionForContentType(req.ContentType))
		err = os.WriteFile(outPath, req.Body, 0755)
		if err != nil {
			return err
		}

		return nil
	}
}

func extensionForContentType(ct string) string {
	switch ct {
	case core.ContentTypeText:
		return "txt"
	case core.ContentTypeMarkdown:
		return "md"
	default:
		return ct[strings.Index(ct, "/")+1:]
	}
}
