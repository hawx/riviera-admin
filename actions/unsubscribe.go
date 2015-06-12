package actions

import (
	"os"

	"hawx.me/code/riviera/subscriptions/opml"
)

func Unsubscribe(opmlPath, page string) error {
	outline, err := opml.Load(opmlPath)
	if err != nil {
		return err
	}

	idx := -1
	for i, o := range outline.Body.Outline {
		if o.XmlUrl == page {
			idx = i
			break
		}
	}

	if idx >= 0 {
		outline.Body.Outline = append(outline.Body.Outline[:idx], outline.Body.Outline[idx+1:]...)
	}

	file, err := os.OpenFile(opmlPath, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	return outline.WriteTo(file)
}
