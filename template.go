package table_doc

import (
	"errors"
	"fmt"
)

const (
	formatHTML      = "html"
	formatBootstrap = "bootstrap"
	formatMarkdown  = "markdown"
)

type Template interface {
	Write(schemas []*Schema) error
	Read() ([]*Schema, error)
}

type View struct {
	template Template
}

func (v *View) Render(schemas []*Schema) error {
	return v.template.Write(schemas)
}

func Render(dataSource string, format string, path string, schemas []*Schema) error {
	var template Template
	switch format {
	case formatHTML:
		template = NewSimpleHTML(dataSource, path)
	case formatBootstrap:
		template = NewBootstrapHTML(dataSource, path)
	case formatMarkdown:
		template = NewMarkdownTemplate(dataSource, path)
	}
	if template == nil {
		return errors.New(fmt.Sprintf("Undefined format %s", format))
	}
	view := View{template: template}
	return view.Render(schemas)
}
