package mysql

import (
	"github.com/nicolerobin/log"
	"github.com/spf13/cobra"
	"github.com/zeromicro/ddl-parser/parser"
	"path/filepath"
)

var (
	VarStrSrc string
	VarStrDir string
)

func Ddl(cmd *cobra.Command, args []string) {
	var err error
	p := parser.NewParser()
	absPath := VarStrSrc
	if !filepath.IsAbs(VarStrSrc) {
		absPath, err = filepath.Abs(VarStrSrc)
		if err != nil {
			log.Error("filepath.Abs() failed, err:%s", err)
			return
		}
	}

	tables, err := p.From(absPath)
	if err != nil {
		log.Error("p.From() failed, err:%s", err)
		return
	}

	for _, table := range tables {
		log.Info("table name:%s", table.Name)
		for _, col := range table.Columns {
			log.Info("col name:%s, type id:%d, type value:%s", col.Name, col.DataType.Type(), col.DataType.Value())
		}
	}
}
