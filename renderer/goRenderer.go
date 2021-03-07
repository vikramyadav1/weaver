package renderer

import "bytes"
import "fmt"
import "github.com/vikramyadav1/weaver/parsers"
import "io/ioutil"
import "path/filepath"
import "text/template"

var UpMigrationFilename string = "UpMigration.go.template"
var DownMigrationFilename string = "DownMigration.go.template"
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
	upMigrationFilepath := filepath.Join(gr.templateDir, "go", UpMigrationFilename)

	b, _ := ioutil.ReadFile(upMigrationFilepath)
	t, _ := template.New("up-migration-tmpl").Parse(string(b))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the up migration.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) DownMigration() []byte {
	downMigrationFilename := filepath.Join(gr.templateDir, "go", DownMigrationFilename)

	b, _ := ioutil.ReadFile(downMigrationFilename)
	t, _ := template.New("down-migration-tmpl").Parse(string(b))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the down migration.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) Model() []byte {
	modelFilepath := filepath.Join(gr.templateDir, "go", ModelFileName)
	b, _ := ioutil.ReadFile(modelFilepath)
	t, _ := template.New("model-tmpl").Parse(string(b))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the model.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) Server() []byte {
	serverFilepath := filepath.Join(gr.templateDir, "go", ServerFileName)
	b, _ := ioutil.ReadFile(serverFilepath)
	t := template.Must(template.New("server-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the server.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) Repository() []byte {
	repositoryFilepath := filepath.Join(gr.templateDir, "go", RepositoryFileName)
	b, _ := ioutil.ReadFile(repositoryFilepath)
	t := template.Must(template.New("repository-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the repository.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) Main() []byte {
	mainFilepath := filepath.Join(gr.templateDir, "go", MainFileName)
	b, _ := ioutil.ReadFile(mainFilepath)
	t := template.Must(template.New("main-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the main file.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}

func (gr goRenderings) PartialMain() []byte {
	partialMainFilepath := filepath.Join(gr.templateDir, "go", MainPartialFileName)
	b, _ := ioutil.ReadFile(partialMainFilepath)
	t := template.Must(template.New("partial-main-tmpl").Parse(string(b)))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, gr.rd); err != nil {
		fmt.Printf("\nAn error occured while rendering the partial main file.\n Error: %v\n", err)
	}

	return tpl.Bytes()
}
