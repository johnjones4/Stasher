package processors

import (
	"encoding/json"
	"fmt"
	"main/core"
	"strings"
	"time"

	"github.com/johnjones4/duration"
)

func Markdown(req *core.Item) error {
	if len(req.Body) == 0 {
		return nil
	}

	switch req.ContentType {
	case core.ContentTypeJson:
		var data []core.StructuredDataProperty
		err := json.Unmarshal(req.Body, &data)
		if err != nil {
			return err
		}

		if recipe, ok := findData(data, []string{"http://schema.org/Recipe"}); ok {
			req.Name, req.Body, err = formatRecipe(recipe, req)
			if err != nil {
				return err
			}
			req.ContentType = core.ContentTypeMarkdown
		}
	}

	return nil
}

func contains(arr []string, i string) bool {
	for _, s := range arr {
		if s == i {
			return true
		}
	}
	return false
}

func findData(data []core.StructuredDataProperty, propTypes []string) (core.StructuredDataProperty, bool) {
	for _, prop := range data {
		for _, propType := range propTypes {
			if contains(prop.Type, propType) {
				return prop, true
			}
		}
	}
	return core.StructuredDataProperty{}, false
}

func findDataMulti(data []core.StructuredDataProperty, propTypes []string) []core.StructuredDataProperty {
	matches := make([]core.StructuredDataProperty, 0)
	for _, prop := range data {
		for _, propType := range propTypes {
			if contains(prop.Type, propType) {
				matches = append(matches, prop)
				break
			}
		}
	}
	return matches
}

func formatRecipe(data core.StructuredDataProperty, req *core.Item) (string, []byte, error) {
	builder := new(strings.Builder)
	title := ""

	if prop, ok := findData(data.Properties, []string{"http://schema.org/headline", "http://schema.org/name"}); ok {
		builder.WriteString(fmt.Sprintf("# %s\n\n", prop.String))
		title = fmt.Sprintf("%s [%s %s]", prop.String, req.URL.Host, time.Now().Format(time.ANSIC))
	}

	if prop, ok := findData(data.Properties, []string{"http://schema.org/description"}); ok {
		builder.WriteString(fmt.Sprintf("%s\n\n", prop.String))
	}

	type metaProp struct {
		key   string
		value string
	}
	metaProps := make([]metaProp, 0)

	if prop, ok := findData(data.Properties, []string{"http://schema.org/recipeYield"}); ok {
		metaProps = append(metaProps, metaProp{
			key:   "Yield",
			value: prop.String,
		})
	}

	if prop, ok := findData(data.Properties, []string{"http://schema.org/totalTime"}); ok {
		durStr := prop.String
		if strings.Index(durStr, "PT") == 0 {
			dur, err := duration.ParseISO8601(durStr)
			if err != nil {
				return "", nil, err
			}
			durStr = dur.TimeDuration().String()
		}
		metaProps = append(metaProps, metaProp{
			key:   "Total Time",
			value: durStr,
		})
	}

	if prop, ok := findData(data.Properties, []string{"http://schema.org/prepTime"}); ok {
		durStr := prop.String
		if strings.Index(durStr, "PT") == 0 {
			dur, err := duration.ParseISO8601(durStr)
			if err != nil {
				return "", nil, err
			}
			durStr = dur.TimeDuration().String()
		}
		metaProps = append(metaProps, metaProp{
			key:   "Prep Time",
			value: durStr,
		})
	}

	if prop, ok := findData(data.Properties, []string{"http://schema.org/cookTime"}); ok {
		durStr := prop.String
		if strings.Index(durStr, "PT") == 0 {
			dur, err := duration.ParseISO8601(durStr)
			if err != nil {
				return "", nil, err
			}
			durStr = dur.TimeDuration().String()
		}
		metaProps = append(metaProps, metaProp{
			key:   "Cook Time",
			value: durStr,
		})
	}

	if len(metaProps) > 0 {
		for _, prop := range metaProps {
			builder.WriteString(fmt.Sprintf("**%s:** %s\n", prop.key, prop.value))
		}
		builder.WriteString("\n")
	}

	ingredients := findDataMulti(data.Properties, []string{"http://schema.org/recipeIngredient"})
	if len(ingredients) > 0 {
		builder.WriteString("## Ingredients:\n\n")

		for _, prop := range ingredients {
			builder.WriteString(fmt.Sprintf("* %s\n", prop.String))
		}

		builder.WriteString("\n")
	}

	steps := findDataMulti(data.Properties, []string{"http://schema.org/HowToStep"})
	if len(steps) > 0 {
		builder.WriteString("## Steps:\n\n")
		for i, prop := range steps {
			if prop.String != "" {
				builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, prop.String))
			} else if text, ok := findData(prop.Properties, []string{"http://schema.org/text"}); ok {
				builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, text.String))
			}
		}
	}

	return title, []byte(builder.String()), nil
}
