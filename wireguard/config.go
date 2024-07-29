package wireguard

import (
	"bytes"
	"embed"
	"strings"
	"text/template"
)

//go:embed *.tmpl
var fs embed.FS

// Template function map for reusability.
var tmplFuncMap = template.FuncMap{
	"join": func(v []string, sep string) string {
		return strings.Join(v, sep)
	},
}

// ToConfig generates the WireGuard client configuration as a string.
func (co *ClientOptions) ToConfig() (string, error) {
	return "", nil
}

// ToConfig generates the WireGuard server configuration as a string.
func (so *ServerOptions) ToConfig() (string, error) {
	text, err := fs.ReadFile("server.conf.tmpl")
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
