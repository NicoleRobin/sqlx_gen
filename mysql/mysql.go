package mysql

import (
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/gen"
	"github.com/spf13/cobra"
)

var (
	VarStrSrc string
	VarStrDir string
)

func Ddl(cmd *cobra.Command, args []string) {
	g, err := gen.NewGenerator(VarStrDir)
	if err != nil {
		log.Error("gen.NewGenerator() failed, err:%s", err)
		return
	}

	err = g.StartFromDDL(VarStrSrc, VarStrDir, "test")
	if err != nil {
		log.Error("g.StartFromDDL() failed, err:%s", err)
		return
	}
}
