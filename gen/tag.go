package gen

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genTag(table Table, filedName string) (string, error) {
	log.Info("genTag(), table:%+v, in:%s", table, filedName)
	if filedName == "" {
		return "", nil
	}

	output, err := template.With("tag").Parse(template.Tag).Execute(map[string]any{
		"field": filedName,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
