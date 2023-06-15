package gen

import (
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genTableName(table Table) (string, error) {
	output, err := template.With("tablename").Parse(template.TableName).Execute(map[string]any{
		"tableName":             table.Name,
		"upperStartCamelObject": strcase.ToCamel(table.Name),
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
