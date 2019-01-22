package table_doc

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"
)

type MarkdownTemplate struct {
	markdown     string
	menu         string
	path         string
	dataSource   string
	htmlTemplate *HTMLTemplate
}

func NewMarkdownTemplate(dataSource string, path string) *MarkdownTemplate {
	t := MarkdownTemplate{
		dataSource: dataSource,
		path: path,
		htmlTemplate: NewBootstrapHTML(dataSource, path),
	}
	// language=markdown
	t.menu = `
{{- $activeSchemaName := .Schema.Name -}}
{{- $lastIndex := .Last -}}
{{- $last := len .Schemas -}}
{{- range $i, $schema := .Schemas -}}
{{- if eq $schema.Name $activeSchemaName -}}
{{ $schema.Name }}
{{- else -}}
<a href="{{ $schema.Name }}.md">{{ $schema.Name }}</a>
{{- end -}}
{{- if ne $i $lastIndex }} | {{ else }} {{ end -}}
{{- end -}}

`

	// language=markdown
	t.markdown = `
# {{ html .Schema.Name }} <a id="table-doc-top"></a>

<div style="text-align: right">
{{ .Menu }}
</div>

{{ .Schema.Memo }}

{{ range $i, $table := .Schema.Tables }}
* [{{ $table.Name }}](#table-{{ $table.Name }}) {{ $table.Comment }}
{{- end }}

{{ range $i, $table := .Schema.Tables }}

## {{ html $table.Name }} <a id="table-{{ html $table.Name }}"></a>

{{ html $table.Type }}

> {{ html $table.Comment }} 

{{ $table.Memo }}

No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
{{- range $k, $column := $table.Columns }}
{{ $column.Position }} | {{ html $column.Name }} | {{ html $column.Type }} | {{ html $column.Default }} | {{ html $column.Nullable }} | {{ html $column.Comment }} | {{ $column.Memo }}  
{{- end }}

` + "```" + `sql
{{ $table.Definition }}
` + "```" + `

[▲ top](#table-doc-top)

{{- end }}
`
	return &t
}

func (t *MarkdownTemplate) Read() ([]*Schema, error) {
	return t.htmlTemplate.Read()
}

func (t *MarkdownTemplate) Write(schemas []*Schema) error {
	if err := t.htmlTemplate.Write(schemas); err != nil {
		return err
	}

	newSchemas, err := t.Read()
	if err != nil {
		return err
	}

	menu := map[string]string{}
	for _, schema := range newSchemas {
		tmpl, err := template.New(schema.Name).Parse(t.menu)
		if err != nil {
			return errors.New(fmt.Sprintf("MarkdownTemplate#Write: fail to parse template. %s", err.Error()))
		}
		buf := &bytes.Buffer{}
		data := map[string]interface{}{
			"Schemas": newSchemas,
			"Schema": schema,
			"Last": len(newSchemas) - 1,
		}
		if err := tmpl.Execute(buf, data); err != nil {
			return errors.New(fmt.Sprintf("MarkdownTemplate#Write: fail to execute template. %s", err.Error()))
		}
		menu[schema.Name] = string(buf.Bytes())
	}

	for _, schema := range newSchemas {
		fpath := fmt.Sprintf("%s/%s.md", t.path, schema.Name)
		file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return errors.New(fmt.Sprintf("MarkdownTemplate#Write: fail to open file %s. %s", fpath, err.Error()))
		}

		tmpl, err := template.New(schema.Name).Parse(t.markdown)
		if err != nil {
			return errors.New(fmt.Sprintf("MarkdownTemplate#Write: fail to parse template. %s", err.Error()))
		}
		if err := tmpl.Execute(file, map[string]interface{}{ "Schema": schema, "Menu": menu[schema.Name]}); err != nil {
			return errors.New(fmt.Sprintf("MarkdownTemplate#Write: fail to execute template. %s", err.Error()))
		}
	}

	return nil
}

