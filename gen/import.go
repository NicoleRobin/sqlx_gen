package gen

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/template"
)

func genImports(table Table, timeImport bool) (string, error) {
	log.Info("table:%+v, timeImport:%t", table, timeImport)
	buffer, err := template.With("import").Parse(template.ImportNoCache).Execute(map[string]any{
		"time":       timeImport,
		"containsPQ": table.ContainsPQ,
		"data":       table,
	})
	if err != nil {
		log.Error("template.Parse() failed, err:%s", err)
		return "", err
	}

	return buffer.String(), nil
}
