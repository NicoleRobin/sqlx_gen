package gen

import (
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/log"
	"strings"

	"github.com/nicolerobin/sqlx_gen/parser"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genFields(table Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genField(table Table, field *parser.Field) (string, error) {
	log.Info("genField(), table:%+v, field:%+v", table, *field)
	tag, err := genTag(table, field.NameOriginal)
	if err != nil {
		return "", err
	}

	output, err := template.With("types").
		Parse(template.Types).
		Execute(map[string]any{
			"name":       strcase.ToCamel(field.Name),
			"type":       field.DataType,
			"tag":        tag,
			"hasComment": field.Comment != "",
			"comment":    field.Comment,
			"data":       table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
