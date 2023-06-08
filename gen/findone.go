package gen

import (
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genFindOne(table Table, withCache, postgreSql bool) (string, string, error) {
	buffer, err := template.With("fineOne").
		Parse(template.FindOne).
		Execute(map[string]any{
			"withCache":                 withCache,
			"upperStartCamelObject":     strcase.ToCamel(table.Name),
			"lowerStartCamelObject":     Untitle(strcase.ToCamel(table.Name)),
			"originalPrimaryKey":        wrapWithRawString(table.PrimaryKey.Name, postgreSql),
			"lowerStartCamelPrimaryKey": EscapeGolangKeyword(Untitle(ToCamel(table.PrimaryKey.Name))),
			"dataType":                  table.PrimaryKey.DataType,
			"cacheKey":                  table.PrimaryCacheKey.KeyExpression,
			"cacheKeyVariable":          table.PrimaryCacheKey.KeyLeft,
			"postgreSql":                postgreSql,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	findOneMethod, err := template.With("findOneMethod").
		Parse(template.InterfaceFindOne).
		Execute(map[string]any{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": util.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return buffer.String(), nil
}
