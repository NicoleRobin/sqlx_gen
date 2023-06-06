package parser

import (
	"github.com/nicolerobin/log"
	"github.com/zeromicro/ddl-parser/parser"
)

type (
	// Table table
	Table struct {
		Name        string
		Db          string
		PrimaryKey  Primary
		UniqueIndex map[string][]*Field
		Fields      []*Field
		ContainsPQ  bool
	}

	// Field field
	Field struct {
		NameOriginal string
		Name         string
		DataType     string
	}

	// Primary
	Primary struct {
		Field
		AutoIncrement bool
	}
)

func Parse(filename, database string) ([]*Table, error) {
	log.Info("Parse(), filename:%s, database:%s", filename, database)
	p := parser.NewParser()
	tables, err := p.From(filename)
	if err != nil {
		log.Error("p.From() failed, err:%s", err)
		return nil, err
	}
	for _, table := range tables {

	}
	return p.From(filename)
}
