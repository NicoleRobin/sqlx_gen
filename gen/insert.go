package gen

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genInsert(table Table) (string, error) {
	log.Info("genInsert(), table:%+v", table)
	output, err := template.With("insert").Parse(template.Insert).Execute(map[string]any{})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
