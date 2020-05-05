package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/CuberL/spar/src/parser"
	"github.com/CuberL/spar/src/types"
)

func getColumnType(column types.Column) string {
	switch column.Type.TypeTag {
	case types.Bool:
		return "bool"
	case types.Int64:
		return "number"
	case types.Float64:
		return "number"
	case types.String:
		return "string"
	case types.Timestamp:
		return "string"
	case types.Date:
		return "string"
	case types.Bytes:
		return "string"
	}
	return ""
}

func getColumnFormat(column types.Column) string {
	switch column.Type.TypeTag {
	case types.Bool:
		return "bool"
	case types.Int64:
		return "number"
	case types.Float64:
		return "number"
	case types.String:
		return "string"
	case types.Timestamp:
		return "timestamp"
	case types.Date:
		return "date"
	case types.Bytes:
		return "bytes"
	}
	return ""
}

func genFieldsFromStatement(createTableStatement types.CreateTableStatement) (Field, error) {
	ret := &ObjectField{
		BaseField: BaseField{
			Required: true,
			Type:     "object",
			Format:   createTableStatement.TableName,
		},
		Properties: map[string]Field{},
	}

	for _, column := range createTableStatement.Columns {
		ret.Properties[column.Name] = &BaseField{
			Type:        getColumnType(column),
			Format:      getColumnFormat(column),
			Required:    column.NotNull,
			Description: column.Comment,
		}
	}

	return ret, nil
}

func main() {
	parseResult, err := parser.Parse(os.Stdin)
	if err != nil {
		log.Panicln(err.Error())
	}

	if len(parseResult.CreateTables) < 1 {
		log.Panicln("create table statement not found.")
	}

	createTableStatement := parseResult.CreateTables[0]
	fields, _ := genFieldsFromStatement(createTableStatement)

	fieldsJson, err := json.Marshal(fields)
	if err != nil {
		log.Panicln(err.Error())
	}

	os.Stdout.Write(fieldsJson)
}
