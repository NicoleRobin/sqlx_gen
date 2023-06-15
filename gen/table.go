package gen

import (
	"github.com/nicolerobin/sqlx_gen/parser"
)

// Table
type Table struct {
	parser.Table
	PrimaryCacheKey        Key
	UniqueCacheKey         []Key
	ContainsUniqueCacheKey bool
	ignoreColumns          []string
}

func (t *Table) isIgnoreColumns(col string) bool {
	for _, column := range t.ignoreColumns {
		if col == column {
			return true
		}
	}
	return false
}
