package gen

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genTypes(table Table) (string, error) {
	log.Info("genTypes(), table:%+v", table)
	fieldsString, err := genFields(table, table.Fields)
	if err != nil {
		return "", err
	}
	output, err := template.With("types").Parse(template.Types).Execute(map[string]any{
		"fields": fieldsString,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
