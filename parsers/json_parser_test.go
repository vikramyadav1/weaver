package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testJson string = `
{
  "name": "hobby",
	"pluralName": "hobbies",
  "columns": [
    {
      "name": "title",
      "type": "string"
    },
    {
      "name": "age",
      "type": "int"
    }
  ],
	"references": {
		"belongsTo": [
			{
				"name": "person"
			}	
		]
	}
}
`

func TestComputeResourceDefinition(t *testing.T) {
	assert := assert.New(t)
	rd := NewJsonParser(testJson).Parse()

	assert.Equal("hobby", rd.Name)
	assert.Equal("hobbies", rd.PluralName)
	assert.Equal("Hobby", rd.PublicName)
	assert.Equal("Hobbies", rd.PublicPluralName)
	assert.Equal("H", rd.PublicVarPrefix)
	assert.Equal("h", rd.VarPrefix)
	assert.Equal("title,age", rd.ColumnsCSV)
	assert.Equal(":title,:age", rd.ColumnsColonCSV)
	assert.Equal("title=:title,age=:age", rd.NamedColumns)
}

func TestParseResourceColumns(t *testing.T) {
	assert := assert.New(t)

	resourceDefinition := NewJsonParser(testJson).Parse()
	rc := resourceDefinition.Columns
	expectedResourceColumns := []ResourceColumn{
		{
			Name:       "title",
			PublicName: "Title",
			Type:       "string",
		},
		{
			Name:       "age",
			PublicName: "Age",
			Type:       "int",
		},
	}

	assert.Equal(expectedResourceColumns, rc)
}

func TestParseResourceReferences(t *testing.T) {
	assert := assert.New(t)

	rd := NewJsonParser(testJson).Parse()
	rr := rd.References
	rc := make([]ResourceColumn, 0)
	rc = append(rc, ResourceColumn{Name: "person"})

	expected := ResourceReferences{
		BelongsTo: rc,
	}

	assert.Equal(expected, rr)
}
