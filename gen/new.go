package gen

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genNew(table Table, withCache, postgreSql bool) (string, error) {
	log.Info("genNew(), table:%+v, withCache:%t, postgreSql:%t", table, withCache, postgreSql)
	t := fmt.Sprintf(`"%s"`, wrapWithRawString(table.Name, postgreSql))
	if postgreSql {
		t = "`" + fmt.Sprintf(`"%s"."%s"`, table.Db, table.Name) + "`"
	}

	output, err := template.With("new").
		Parse(template.ModelNew).
		Execute(map[string]any{
			"table":                 t,
			"withCache":             withCache,
			"upperStartCamelObject": strcase.ToCamel(table.Name),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
