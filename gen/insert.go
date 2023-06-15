package gen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/golang-collections/collections/set"
	"github.com/iancoleman/strcase"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genInsert(table Table, withCache, isPostgreSql bool) (string, string, error) {
	keySet := set.New()
	keyVariableSet := set.New()
	keySet.Insert(table.PrimaryCacheKey.DataKeyExpression)
	keyVariableSet.Insert(table.PrimaryCacheKey.KeyLeft)
	for _, key := range table.UniqueCacheKey {
		keySet.Insert(key.DataKeyExpression)
		keyVariableSet.Insert(key.KeyLeft)
	}
	var keys []string
	keySet.Do(func(i interface{}) {
		keys = append(keys, i.(string))
	})
	sort.Strings(keys)
	var keyVars []string
	keyVariableSet.Do(func(i interface{}) {
		keyVars = append(keyVars, i.(string))
	})
	sort.Strings(keyVars)

	expressions := make([]string, 0)
	expressionValues := make([]string, 0)
	var count int
	for _, field := range table.Fields {
		camel := strcase.ToCamel(field.Name)
		if table.isIgnoreColumns(field.Name) {
			continue
		}

		if field.Name == table.PrimaryKey.Name {
			if table.PrimaryKey.AutoIncrement {
				continue
			}
		}

		count += 1
		if isPostgreSql {
			expressions = append(expressions, fmt.Sprintf("$%d", count))
		} else {
			expressions = append(expressions, "?")
		}
		expressionValues = append(expressionValues, "data."+camel)
	}

	camel := strcase.ToCamel(table.Name)
	output, err := template.With("insert").Parse(template.Insert).Execute(map[string]any{
		"withCache":             withCache,
		"upperStartCamelObject": camel,
		"lowerStartCamelObject": strcase.ToLowerCamel(camel),
		"expression":            strings.Join(expressions, ", "),
		"expressionValues":      strings.Join(expressionValues, ", "),
		"keys":                  strings.Join(keys, "\n"),
		"keyValues":             strings.Join(keyVars, ", "),
		"data":                  table,
	})
	if err != nil {
		return "", "", err
	}

	insertMethodOutput, err := template.With("insertMethod").Parse(template.InterfaceInsert).Execute(map[string]any{
		"upperStartCamelObject": camel,
		"data":                  table,
	})
	if err != nil {
		return "", "", err
	}
	return output.String(), insertMethodOutput.String(), nil
}
