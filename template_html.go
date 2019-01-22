package table_doc

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

const (
	htmlTemplateSchemaContainer = "*[data-table-doc='schema-container']"
	htmlTemplateSchemaName      = "*[data-table-doc='schema-name']"
	htmlTemplateSchemaMemo      = "*[data-table-doc='schema-memo']"

	htmlTemplateTableContainer  = "*[data-table-doc='table-container']"
	htmlTemplateTableName       = "*[data-table-doc='table-name']"
	htmlTemplateTableType       = "*[data-table-doc='table-type']"
	htmlTemplateTableComment    = "*[data-table-doc='table-comment']"
	htmlTemplateTableDefinition = "*[data-table-doc='table-definition']"
	htmlTemplateTableMemo       = "*[data-table-doc='table-memo']"

	htmlTemplateColumnContainer = "*[data-table-doc='column-container']"
	htmlTemplateColumnPosition  = "*[data-table-doc='column-position']"
	htmlTemplateColumnName      = "*[data-table-doc='column-name']"
	htmlTemplateColumnType      = "*[data-table-doc='column-type']"
	htmlTemplateColumnDefault   = "*[data-table-doc='column-default']"
	htmlTemplateColumnNullable  = "*[data-table-doc='column-nullable']"
	htmlTemplateColumnComment   = "*[data-table-doc='column-comment']"
	htmlTemplateColumnMemo      = "*[data-table-doc='column-memo']"
)

type HTMLTemplate struct {
	path       string
	html       string
	menu       string
	dataSource string
	assets     map[string]string
}

func (t *HTMLTemplate) makeTitle(schemas []*Schema) string {
	var names []string
	for _, schema := range schemas {
		names = append(names, schema.Name)
	}
	return fmt.Sprintf("Schema %s | Table Doc", strings.Join(names, " "))
}

func (t *HTMLTemplate) lang() string {
	var lang string
	var ok bool

	lang, ok = os.LookupEnv("LANG")
	if !ok {
		lang, ok = os.LookupEnv("LC_CTYPE")
		if !ok {
			lang = "en"
		}
	}

	l := strings.Split(lang, ".")
	if len(l) > 0 {
		return l[0]
	} else {
		return "en"
	}
}

func (t *HTMLTemplate) Read() ([]*Schema, error) {
	if _, err := os.Stat(t.path); os.IsNotExist(err) {
		return []*Schema{}, nil
	} else if info, _ := os.Stat(t.path); !info.IsDir() {
		return []*Schema{}, nil
	}

	dir, err := ioutil.ReadDir(t.path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("HTMLTemplate#Read: fail to read dir %s. %s", t.path, err.Error()))
	}

	var schemas []*Schema
	for _, info := range dir {
		if path.Ext(info.Name()) == ".html" {
			fpath := fmt.Sprintf("%s/%s", t.path, info.Name())
			file, err := os.Open(fpath)
			defer file.Close()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("HTMLTemplate#Read: fail to read file %s. %s", fpath, err.Error()))
			}
			schema, err := t.parseSchema(file)
			if err != nil {
				return nil, err
			}
			schemas = append(schemas, schema)
		}
	}
	return schemas, nil
}

func (t *HTMLTemplate) Write(schemas []*Schema) error {
	readSchemas, err := t.Read()
	if err != nil {
		return err
	}

	if _, err := os.Stat(t.path); os.IsNotExist(err) {
		if err := os.MkdirAll(t.path, 0755); err != nil {
			return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to make dir %s. %s", t.path, err.Error()))
		}
	} else if info, _ := os.Stat(t.path); !info.IsDir() {
		return errors.New(fmt.Sprintf("HTMLTemplate#Write: %s is not directory", t.path))
	}

	newSchemas := UpdateSchemas(readSchemas, schemas)

	// create menu
	menu := map[string]string{}
	for _, schema := range newSchemas {
		tmpl, err := template.New(schema.Name).Parse(t.menu)
		if err != nil {
			return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to parse template. %s", err.Error()))
		}
		buf := &bytes.Buffer{}
		data := map[string]interface{}{
			"Schemas": newSchemas,
			"Schema": schema,
		}
		if err := tmpl.Execute(buf, data); err != nil {
			return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to execute template. %s", err.Error()))
		}
		menu[schema.Name] = string(buf.Bytes())
	}

	// assets
	for name, url := range t.assets {
		assetsDir := fmt.Sprintf("%s/assets", t.path)
		if err := os.MkdirAll(assetsDir, 0755); err != nil {
			return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to make dir %s. %s", assetsDir, err.Error()))
		}
		assetPath := fmt.Sprintf("%s/%s", assetsDir, name)
		if _, err := os.Stat(assetPath); os.IsNotExist(err) {
			if err := t.downloadAsset(assetPath, url); err != nil {
				return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to download asset %s. %s", assetPath, err.Error()))
			}
		}
	}

	lang := t.lang()
	for _, schema := range newSchemas {
		fpath := fmt.Sprintf("%s/%s.html", t.path, schema.Name)
		file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to open file %s. %s", fpath, err.Error()))
		}

		tmpl, err := template.New(schema.Name).Parse(t.html)
		if err != nil {
			return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to parse template. %s", err.Error()))
		}

		title := fmt.Sprintf("Schema %s | Table Doc", schema.Name)
		data := map[string]interface{}{
			"Lang": lang,
			"Title": title,
			"Schema": schema,
			"Menu": menu[schema.Name],
		}
		if err := tmpl.Execute(file, data); err != nil {
			return errors.New(fmt.Sprintf("HTMLTemplate#Write: fail to execute template. %s", err.Error()))
		}
	}
	return nil
}

func (t *HTMLTemplate) downloadAsset(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (t *HTMLTemplate) parseSchema(file io.Reader) (*Schema, error) {
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("HTMLTemplate#parseSchema: fail to read document. %s", err.Error()))
	}
	selection := doc.Find(htmlTemplateSchemaContainer).First()

	schema := Schema{}
	schema.Name = html.UnescapeString(selection.Find(htmlTemplateSchemaName).First().Text())
	schema.Memo, err = selection.Find(htmlTemplateSchemaMemo).First().Html()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("HTMLTemplate#parseSchema: fail to parse element. %s", err.Error()))
	}
	selection.Find(htmlTemplateTableContainer).Each(func(j int, t *goquery.Selection) {
		var table Table
		table.Name = html.UnescapeString(t.Find(htmlTemplateTableName).First().Text())
		table.Type = html.UnescapeString(t.Find(htmlTemplateTableType).First().Text())
		table.Comment = html.UnescapeString(t.Find(htmlTemplateTableComment).First().Text())
		table.Definition = html.UnescapeString(t.Find(htmlTemplateTableDefinition).First().Text())
		table.Memo, _ = t.Find(htmlTemplateTableMemo).First().Html()
		t.Find(htmlTemplateColumnContainer).Each(func(k int, c *goquery.Selection) {
			var column Column
			pos, _ := strconv.Atoi(c.Find(htmlTemplateColumnPosition).First().Text())
			column.Position = uint(pos)
			column.Name = html.UnescapeString(c.Find(htmlTemplateColumnName).First().Text())
			column.Type = html.UnescapeString(c.Find(htmlTemplateColumnType).First().Text())
			column.Default = html.UnescapeString(c.Find(htmlTemplateColumnDefault).First().Text())
			column.Nullable = html.UnescapeString(c.Find(htmlTemplateColumnNullable).First().Text())
			column.Comment = html.UnescapeString(c.Find(htmlTemplateColumnComment).First().Text())
			column.Memo, _ = c.Find(htmlTemplateColumnMemo).First().Html()
			table.Columns = append(table.Columns, &column)
		})
		schema.Tables = append(schema.Tables, &table)
	})

	return &schema, nil
}

func NewSimpleHTML(dataSource string, path string) *HTMLTemplate {
	t := HTMLTemplate{dataSource: dataSource, path: path}

	// language=html
	t.menu = `
{{- $activeSchemaName := .Schema.Name -}}
<ul>
{{- range $i, $schema := .Schemas }}
    <li>{{ if eq $schema.Name $activeSchemaName }}{{ html $schema.Name }}{{ else }}<a href="{{ $schema.Name }}.html">{{ html $schema.Name }}</a>{{ end }}</li>
{{- end }}
</ul>
`

	// language=html
	t.html = `<!DOCTYPE html>
<html lang="{{ .Lang }}">
<head><meta charset="utf-8"><title>{{ .Title }}</title></head>
<body>
{{ .Menu }}

<section data-table-doc="schema-container">
    <h2 data-table-doc="schema-name">{{ html .Schema.Name }}</h2>
    <div data-table-doc="schema-memo">{{ .Schema.Memo }}</div>
    <ul>
    {{- range $j, $table := .Schema.Tables }}
        <li><a href="#{{ $table.Name }}">{{ html $table.Name }}</a></li>
    {{- end }}
    </ul>
    {{- range $j, $table := .Schema.Tables }}
    <div data-table-doc="table-container" id="{{ $table.Name }}">
        <h3 data-table-doc="table-name">{{ html $table.Name }}</h3>
        <dl>
            <dt>Type</dt><dd data-table-doc="table-type">{{ html $table.Type }}</dd>
            <dt>Comment</dt><dd data-table-doc="table-comment">{{ html $table.Comment }}</dd>
            <dt>Memo</dt><dd data-table-doc="table-memo">{{ $table.Memo }}</dd>
            <dt>Definition</dt><dd><pre data-table-doc="table-definition">{{ html $table.Definition }}</pre></dd>
        </dl>
        <table>
            <thead>
            <tr>
                <th>&nbsp;</th>
                <th>Name</th>
                <th>Type</th>
                <th>Default</th>
                <th>Nullable</th>
                <th>Comment</th>
                <th>Memo</th>
            </tr>
            </thead>
            <tbody>
            {{- range $k, $column := $table.Columns }}
            <tr data-table-doc="column-container">
                <td data-table-doc="column-position">{{ html $column.Position }}</td>
                <td data-table-doc="column-name">{{ html $column.Name }}</td>
                <td data-table-doc="column-type">{{ html $column.Type }}</td>
                <td data-table-doc="column-default">{{ html $column.Default }}</td>
                <td data-table-doc="column-nullable">{{ html $column.Nullable }}</td>
                <td data-table-doc="column-comment">{{ html $column.Comment }}</td>
                <td data-table-doc="column-memo">{{ $column.Memo }}</td>
            </tr>
            {{- end }}
            </tbody>
        </table>
    </div>
    {{- end }}
</section>
</body>
</html>
`
	return &t
}

func NewBootstrapHTML(dataSource string, path string) *HTMLTemplate {
	t := HTMLTemplate{dataSource: dataSource, path: path}
	t.assets = map[string]string{
		"bootstrap.4.2.1.css": "https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css",
		"bootstrap.4.2.1.js":  "https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js",
		"jquery.3.3.1.js":     "https://code.jquery.com/jquery-3.3.1.slim.min.js",
		"popper.1.14.3.js":    "https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js",
	}

	// language=html
	t.menu = `
{{ $activeSchemaName := .Schema.Name }}
<nav class="navbar navbar-dark bg-dark navbar-expand-sm" id="top">
    <div class="container-fluid" >
        <span class="navbar-brand mb-0 h1">Table-doc</span>
        <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
                {{- range $i, $schema := .Schemas }}
                <li class="nav-item{{ if eq $schema.Name $activeSchemaName }} active{{ end }}">
                    <a class="nav-link" href="{{ $schema.Name }}.html">{{ $schema.Name }}</a>
                </li>
                {{- end }}
            </ul>
        </div>
    </div>
</nav>
`

	// language=html
	t.html = `<!DOCTYPE html>
<html lang="{{ .Lang }}">
<head>
    <title>{{ .Title }}</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="assets/bootstrap.4.2.1.css">
    <script src="assets/jquery.3.3.1.js"></script>
    <script src="assets/popper.1.14.3.js"></script>
    <script src="assets/bootstrap.4.2.1.js"></script>
</head>
<body>

{{ .Menu }}

<div class="container-fluid py-3 bg-white rounded shadow-sm">
    <section id="schema-{{ .Schema.Name }}" data-table-doc="schema-container">
        <div class="container-fluid my-4">
            <div class="card">
                <div class="card-body">
                    <h3 class="card-title" data-table-doc="schema-name">{{ html .Schema.Name }}</h3>
                    <div class="py-2" data-table-doc="schema-memo">{{ .Schema.Memo }}</div>
                </div>
            </div>
        </div>
        <div class="container-fluid">
            <div class="row">
                <div class="col-xl-2 mb-4">
                    <ul class="nav flex-column sticky-top">
                        <li class="text-muted ">Tables</li>
                        {{- range $j, $table := .Schema.Tables }}
                        <li class="nav-item"><a href="#{{ $table.Name }}" class="nav-link p-0 text-truncate font-weight-light">{{ html $table.Name }}</a></li>
                        {{- end }}
                    </ul>
                </div>
                <div class="col-xl-10">
                    {{- range $j, $table := .Schema.Tables }}
                    <div class="card px-3 mb-5" id="{{ $table.Name }}" data-table-doc="table-container">
                        <div class="card-body px-0 py-3">
                            <h3 class="card-title" data-table-doc="table-name">{{ html $table.Name }}</h3>
                            <div class="text-muted text-right" data-table-doc="table-type">{{ html $table.Type }}</div>
                            <div class="text-muted" data-table-doc="table-comment">{{ html $table.Comment }}</div>
                            <div class="py-2 edit" data-table-doc="table-memo">{{ $table.Memo }}</div>
                            <div class="text-right"><a href="#{{ $table.Name }}-definition" data-toggle="collapse" class="btn btn-outline-secondary btn-sm">definition</a></div>
                            <pre class="collapse border p-4 mt-2 rounded bg-light" id="{{ $table.Name }}-definition" data-table-doc="table-definition">{{ html $table.Definition }}</pre>
                        </div>
                        <table class="table thead-light table-bordered table-sm">
                            <thead class="thead-dark">
                            <tr>
                                <th>&nbsp;</th>
                                <th>Name</th>
                                <th>Type</th>
                                <th>Default</th>
                                <th>Nullable</th>
                                <th>Comment</th>
                                <th>Memo</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{- range $k, $column := $table.Columns }}
                                <tr data-table-doc="column-container">
                                    <td data-table-doc="column-position">{{ html $column.Position }}</td>
                                    <td data-table-doc="column-name">{{ html $column.Name }}</td>
                                    <td data-table-doc="column-type">{{ html $column.Type }}</td>
                                    <td data-table-doc="column-default">{{ html $column.Default }}</td>
                                    <td data-table-doc="column-nullable">{{ html $column.Nullable }}</td>
                                    <td data-table-doc="column-comment">{{ html $column.Comment }}</td>
                                    <td data-table-doc="column-memo">{{ $column.Memo }}</td>
                                </tr>
                            {{- end }}
                            </tbody>
                        </table>
                    </div>
                    <div class="text-right mb-3">
                        <a href="#top">top</a>
                    </div>
                    {{- end }}
                </div>
            </div>
        </div>
    </section>
</div>
</body>
</html>
`
	return &t
}
