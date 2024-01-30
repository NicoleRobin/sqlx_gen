package gen

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genCount(table Table) (string, error) {
	log.Info("genCount(), table:%+v", table)
	output, err := template.With("count").Parse(template.Count).Execute(map[string]any{})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
