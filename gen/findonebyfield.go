package gen

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/sqlx_gen/template"
	"strings"
)

type findOneCode struct {
	findOneMethod          string
	findOneInterfaceMethod string
	cacheExtra             string
}

func genFindOneByField(table Table, withCache, postgreSql bool) (*findOneCode, error) {
	t := template.With("findOneByField").Parse(template.FindOneByField)
	var list []string
	camelTableName := strcase.ToCamel(table.Name)
	for _, key := range table.UniqueCacheKey {
		in, paramJoinString, originalFieldString := convertJoin(key, postgreSql)

		output, err := t.Execute(map[string]any{
			"upperStartCamelObject":     camelTableName,
			"upperField":                key.FieldNameJoin,
			"in":                        in,
			"withCache":                 withCache,
			"cacheKey":                  key.KeyExpression,
			"cacheKeyVariable":          key.KeyLeft,
			"lowerStartCamelObject":     strcase.ToLowerCamel(camelTableName),
			"lowerStartCamelField":      paramJoinString,
			"upperStartCamelPrimaryKey": strcase.ToCamel(table.PrimaryKey.Name),
			"originalField":             originalFieldString,
			"postgreSql":                postgreSql,
			"data":                      table,
		})
		if err != nil {
			return nil, err
		}

		list = append(list, output.String())
	}

	t = template.With("findOneByFieldMethod").Parse(template.FindOneByFieldMethod)
	var listMethod []string
	for _, key := range table.UniqueCacheKey {
		var inJoin, paramJoin Join
		for _, f := range key.Fields {
			param := EscapeGolangKeyword(strcase.ToLowerCamel(f.Name))
			inJoin = append(inJoin, fmt.Sprintf("%s %s", param, f.DataType))
			paramJoin = append(paramJoin, param)
		}

		var in string
		if len(inJoin) > 0 {
			in = inJoin.With(", ").Source()
		}
		output, err := t.Execute(map[string]any{
			"upperStartCamelObject": camelTableName,
			"upperField":            key.FieldNameJoin.Camel().With("").Source(),
			"in":                    in,
			"data":                  table,
		})
		if err != nil {
			return nil, err
		}

		listMethod = append(listMethod, output.String())
	}

	if withCache {
		text, err := pathx.LoadTemplate(category, findOneByFieldExtraMethodTemplateFile,
			template.FindOneByFieldExtraMethod)
		if err != nil {
			return nil, err
		}

		out, err := util.With("findOneByFieldExtraMethod").Parse(text).Execute(map[string]any{
			"upperStartCamelObject": camelTableName,
			"primaryKeyLeft":        table.PrimaryCacheKey.VarLeft,
			"lowerStartCamelObject": stringx.From(camelTableName).Untitle(),
			"originalPrimaryField":  wrapWithRawString(table.PrimaryKey.Name.Source(), postgreSql),
			"postgreSql":            postgreSql,
			"data":                  table,
		})
		if err != nil {
			return nil, err
		}

		return &findOneCode{
			findOneMethod:          strings.Join(list, pathx.NL),
			findOneInterfaceMethod: strings.Join(listMethod, pathx.NL),
			cacheExtra:             out.String(),
		}, nil
	}

	return &findOneCode{
		findOneMethod:          strings.Join(list, pathx.NL),
		findOneInterfaceMethod: strings.Join(listMethod, pathx.NL),
	}, nil
}

func convertJoin(key Key, postgreSql bool) (in, paramJoinString, originalFieldString string) {
	var inJoin, paramJoin, argJoin Join
	for index, f := range key.Fields {
		param := util.EscapeGolangKeyword(stringx.From(f.Name.ToCamel()).Untitle())
		inJoin = append(inJoin, fmt.Sprintf("%s %s", param, f.DataType))
		paramJoin = append(paramJoin, param)
		if postgreSql {
			argJoin = append(argJoin, fmt.Sprintf("%s = $%d", wrapWithRawString(f.Name.Source(), postgreSql), index+1))
		} else {
			argJoin = append(argJoin, fmt.Sprintf("%s = ?", wrapWithRawString(f.Name.Source(), postgreSql)))
		}
	}
	if len(inJoin) > 0 {
		in = inJoin.With(", ").Source()
	}

	if len(paramJoin) > 0 {
		paramJoinString = paramJoin.With(",").Source()
	}

	if len(argJoin) > 0 {
		originalFieldString = argJoin.With(" and ").Source()
	}
	return in, paramJoinString, originalFieldString
}