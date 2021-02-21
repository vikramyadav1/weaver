package templates

type Migration struct{}

func (m Migration) Up() string {
	return `CREATE TABLE {{.TableName}} {
	{{range .TableColumns}}
		{{.Name}} {{.Type}},
	{{end}}
	};`
}

func (m Migration) Down() string {
	return `DROP TABLE IF EXISTS {{.TableName}}`
}
