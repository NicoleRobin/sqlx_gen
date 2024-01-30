package gen

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genTableName(table Table) (string, error) {
	log.Info("genTableName(), table:%+v", table)
	output, err := template.With("tablename").Parse(template.TableName).Execute(map[string]any{
		"tableName": table.Name,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
