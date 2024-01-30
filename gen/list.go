package gen

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genList(table Table) (string, error) {
	log.Info("genList(), table:%+v", table)
	output, err := template.With("list").Parse(template.List).Execute(map[string]any{})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
