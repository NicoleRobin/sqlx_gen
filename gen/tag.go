package gen

import "github.com/nicolerobin/sqlx_gen/template"

func genTag(table Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	output, err := template.With("tag").Parse(template.Tag).Execute(map[string]any{
		"field": in,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
