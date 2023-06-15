package gen

import (
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genTypes(table Table, methods string, withCache bool) (string, error) {
	fieldsString, err := genFields(table, table.Fields)
	if err != nil {
		return "", err
	}
	output, err := template.With("types").Parse(template.Types).Execute(map[string]any{
		"withCache":             withCache,
		"method":                methods,
		"upperStartCamelObject": strcase.ToCamel(table.Name),
		"lowerStartCamelObject": strcase.ToLowerCamel(table.Name),
		"fields":                fieldsString,
		"data":                  table,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
