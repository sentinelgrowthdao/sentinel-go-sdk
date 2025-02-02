package v2ray

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed *.tmpl
var fs embed.FS

// Template function map for reusability.
var tmplFuncMap = template.FuncMap{
	"sum": func(a, b int) int {
		return a + b
	},
}

// ToConfig generates the V2Ray client configuration as a string.
func (co *ClientOptions) ToConfig() (string, error) {
	return "", nil
}

// ToConfig generates the V2Ray server configuration as a string.
func (so *ServerOptions) ToConfig() (string, error) {
	text, err := fs.ReadFile("server.json.tmpl")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("config").
		Funcs(tmplFuncMap).
		Parse(string(text))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, so); err != nil {
		return "", err
	}

	return buf.String(), nil
}
