package parsers

import "encoding/json"
import "fmt"
import "strings"

type jsonParser struct {
	jsonString string
}

func NewJsonParser(jsonText string) jsonParser {
	return jsonParser{jsonString: jsonText}
}

func (j jsonParser) Parse() ResourceDefinition {
	var rd ResourceDefinition
	err := json.Unmarshal([]byte(j.jsonString), &rd)

	if err != nil {
		fmt.Printf("Error occured while parsing json: %v", err)
	}

	fullRd := computeDefinitions(rd)
	return fullRd
}

func computeDefinitions(rd ResourceDefinition) ResourceDefinition {
	nameSlice := strings.Split(rd.Name, "")
	var columnNames []string = make([]string, 0)
	var columnColonNames []string = make([]string, 0)
	var namedColonColumns []string = make([]string, 0)

	rd.Name = strings.ToLower(rd.Name)
	rd.PublicName = strings.Title(rd.Name)
	if rd.PluralName == "" {
		rd.PluralName = strings.ToLower(rd.Name) + "s"
	}
	rd.PublicPluralName = strings.Title(rd.PluralName)
	rd.PublicVarPrefix = strings.ToUpper(nameSlice[0])
	rd.VarPrefix = strings.ToLower(nameSlice[0])

	for i, column := range rd.Columns {
		column.PublicName = strings.Title(column.Name)
		columnNames = append(columnNames, column.Name)
		columnColonNames = append(columnColonNames, ":"+column.Name)
		namedColonColumns = append(namedColonColumns, column.Name+"="+":"+column.Name)

		rd.Columns[i] = column
	}

	rd.ColumnsCSV = strings.Join(columnNames, ",")
	rd.ColumnsColonCSV = strings.Join(columnColonNames, ",")
	rd.NamedColumns = strings.Join(namedColonColumns, ",")

	return rd
}
