package gen

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/sqlx_gen/template"
	"sort"
	"strings"
)

func genVars(table Table, withCache, postgreSql bool) (string, error) {
	keys := make([]string, 0)
	keys = append(keys, table.PrimaryCacheKey.VarExpression)
	for _, v := range table.UniqueCacheKey {
		keys = append(keys, v.VarExpression)
	}

	camel := strcase.ToCamel(table.Name)
	output, err := template.With("var").
		Parse(template.Var).
		GoFmt(true).Execute(map[string]any{
		"lowerStartCamelObject": strcase.ToLowerCamel(table.Name),
		"upperStartCamelObject": camel,
		"cacheKeys":             strings.Join(keys, "\n"),
		"autoIncrement":         table.PrimaryKey.AutoIncrement,
		"originalPrimaryKey":    wrapWithRawString(table.PrimaryKey.Name, postgreSql),
		"withCache":             withCache,
		"postgreSql":            postgreSql,
		"data":                  table,
		"ignoreColumns": func() string {
			var set = make(map[string]interface{})
			for _, c := range table.ignoreColumns {
				if postgreSql {
					set[fmt.Sprintf(`"%s"`, c)] = struct{}{}
				} else {
					set[fmt.Sprintf("\"`%s`\"", c)] = struct{}{}
				}
			}
			var list []string
			for key, _ := range set {
				list = append(list, key)
			}
			sort.Strings(list)
			return strings.Join(list, ", ")
		}(),
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
