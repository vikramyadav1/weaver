package renderer

import "bytes"
import "fmt"
import "github.com/vikramyadav1/weaver/parsers"
import "io/ioutil"
import "path/filepath"
import "text/template"

var ModelFileName string = "model.go.template"
var RepositoryFileName string = "repository.go.template"
var ServerFileName string = "server.go.template"
var MainFileName string = "main.go.template"
var MainPartialFileName = "main.partial.go.template"

type goRenderer struct {
	templateDir string
	rd          parsers.ResourceDefinition
}

func NewGoRenderer(templateDir string, rd parsers.ResourceDefinition) goRenderer {
	return goRenderer{templateDir, rd}
}

func (gr goRenderer) Render() Renderings {

	return goRenderings{
		rd:          gr.rd,
		templateDir: gr.templateDir,
	}
}

type goRenderings struct {
	rd          parsers.ResourceDefinition
	templateDir string
}

func (gr goRenderings) UpMigration() []byte {
	return make([]byte, 0)
}

func (gr goRenderings) DownMigration() []byte {
	return make([]byte, 0)
}

func (gr goRenderings) Model() []byte {
	filepath := filepath.Join(gr.templateDir, "go", ModelFileName)
	b, _ := ioutil.ReadFile(filepath)
	t, _ := template.New("model-tmpl").Parse(string(b))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the model.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) Server() []byte {
	filepath := filepath.Join(gr.templateDir, "go", ServerFileName)
	b, _ := ioutil.ReadFile(filepath)
	t := template.Must(template.New("server-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the server.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) Repository() []byte {
	filepath := filepath.Join(gr.templateDir, "go", RepositoryFileName)
	b, _ := ioutil.ReadFile(filepath)
	t := template.Must(template.New("repository-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the repository.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) Main() []byte {
	filepath := filepath.Join(gr.templateDir, "go", MainFileName)
	b, _ := ioutil.ReadFile(filepath)
	t := template.Must(template.New("main-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the main file.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) PartialMain() []byte {
	filepath := filepath.Join(gr.templateDir, "go", MainPartialFileName)
	b, _ := ioutil.ReadFile(filepath)
	t := template.Must(template.New("partial-main-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the partial main file.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}
