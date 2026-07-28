package main

import (
	"strings"

	"charm.land/glamour/v2"
)

func previewMd(editorLines []string) ([]string, error) {
	in := strings.Join(editorLines, "\n")

	out, err := glamour.Render(in, "dracula")

	if err != nil {
		return nil, err
	}

	outLines := strings.Split(out, "\n")

	return outLines, nil
}
