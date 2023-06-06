package template

import (
	_ "embed"
)

const (
	regularPerm = 0o666
)

// Insert defines a template for insert code in model
//
//go:embed tpl/insert.tpl
var Insert string

// TableName defines a template for table name
//
//go:embed tpl/table-name.tpl
var TableName string
