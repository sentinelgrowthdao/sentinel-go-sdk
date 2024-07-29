package v2ray

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed *.tmpl
var fs embed.FS

func (co *ClientOptions) ToConfig() (string, error) {
	return "", nil
}

func (so *ServerOptions) ToConfig() (string, error) {
	text, err := fs.ReadFile("server_config.json.tmpl")
	if err != nil {
		return "", err
	}

	tmplFuncMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"tag": func(v *InboundServerOptions) string {
			return v.Tag()
		},
	}

	tmpl, err := template.New("server_config").
		Funcs(tmplFuncMap).
		Parse(string(text))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, so); err != nil {
		return "", err
	}

	return buf.String(), nil
}
