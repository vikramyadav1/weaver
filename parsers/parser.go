package parsers

type ResourceDefinition struct {
	Name             string
	PluralName       string
	PublicName       string
	PublicPluralName string
	PublicVarPrefix  string
	VarPrefix        string
	ColumnsCSV       string
	ColumnsColonCSV  string
	NamedColumns     string
	Columns          []ResourceColumn
	References       ResourceReferences
}

type ResourceColumn struct {
	Name       string
	PublicName string
	Type       string
}

type ResourceReferences struct {
	BelongsTo []ResourceColumn
}

type Parser interface {
	Parse() ResourceDefinition
}
