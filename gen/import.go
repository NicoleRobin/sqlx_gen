package gen

import "github.com/nicolerobin/sqlx_gen/template"

func genImports(table Table, timeImport bool) (string, error) {
	buffer, err := template.With("import").Parse(template.ImportNoCache).Execute(map[string]any{
		"time":       timeImport,
		"containsPQ": table.ContainsPQ,
		"data":       table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
