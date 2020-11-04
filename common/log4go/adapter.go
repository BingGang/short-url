package log4go

import (
	"strings"
)

var adapter = strings.NewReplacer(
	"%T", "{{if .TIME}}{{.TIME.Format \"15:04:05\"}}{{end}}",
	"%D", "{{if .TIME}}{{.TIME.Format \"2006-01-02\"}}{{end}}",
	"%L", "{{.LEVEL}}",
	"%S", "{{.SOURCE}}",
	"%M", "{{.MESSAGE}}",
)
