package template

import (
	_ "embed"
)

const (
	regularPerm = 0o666
)

// Err
//
//go:embed tpl/err.tpl
var Err string

// Field
//
//go:embed tpl/field.tpl
var Field string

// FindOne
//
//go:embed tpl/find-one.tpl
var FindOne string

// FindOneByField
//
//go:embed tpl/find-one-by-field.tpl
var FindOneByField string

// FindOneByField
//
//go:embed tpl/interface-find-one-by-field.tpl
var FindOneByFieldMethod string

// FindOneByFieldExtraMethod
//
//go:embed tpl/find-one-by-field-extra-method.tpl
var FindOneByFieldExtraMethod string

// Import
//
//go:embed tpl/import.tpl
var Import string

// ImportNoCache
//
//go:embed tpl/import-no-cache.tpl
var ImportNoCache string

// Insert defines a template for insert code in model
//
//go:embed tpl/insert.tpl
var Insert string

// InterfaceDelete
//
//go:embed tpl/interface-delete.tpl
var InterfaceDelete string

// InterfaceFindOne
//
//go:embed tpl/interface-find-one.tpl
var InterfaceFindOne string

// InterfaceFineOneByField
//
//go:embed tpl/interface-find-one-by-field.tpl
var InterfaceFindOneByField string

// InterfaceInsert
//
//go:embed tpl/interface-insert.tpl
var InterfaceInster string

// InterfaceUpdate
//
//go:embed tpl/interface-update.tpl
var InterfaceUpdate string

// Model defines a template for table name
//
//go:embed tpl/model.tpl
var Model string

// ModelNew defines a template for table name
//
//go:embed tpl/model-new.tpl
var ModelNew string

// TableName defines a template for table name
//
//go:embed tpl/table-name.tpl
var TableName string

// Tag defines a template for table name
//
//go:embed tpl/tag.tpl
var Tag string

// Types defines a template for table name
//
//go:embed tpl/types.tpl
var Types string

// Update defines a template for table name
//
//go:embed tpl/update.tpl
var Update string

// Var defines a template for table name
//
//go:embed tpl/var.tpl
var Var string
