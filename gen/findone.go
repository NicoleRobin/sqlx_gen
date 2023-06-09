package gen

import (
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genFindOne(table Table, withCache, postgreSql bool) (string, string, error) {
	output, err := template.With("fineOne").
		Parse(template.FindOne).
		Execute(map[string]any{
			"withCache":                 withCache,
			"upperStartCamelObject":     strcase.ToCamel(table.Name),
			"lowerStartCamelObject":     Untitle(strcase.ToCamel(table.Name)),
			"originalPrimaryKey":        wrapWithRawString(table.PrimaryKey.Name, postgreSql),
			"lowerStartCamelPrimaryKey": EscapeGolangKeyword(strcase.ToLowerCamel(table.PrimaryKey.Name)),
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
			"upperStartCamelObject":     strcase.ToCamel(table.Name),
			"lowerStartCamelPrimaryKey": EscapeGolangKeyword(strcase.ToLowerCamel(table.PrimaryKey.Name)),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return output.String(), findOneMethod.String(), nil
}

func genFineOneByField()
