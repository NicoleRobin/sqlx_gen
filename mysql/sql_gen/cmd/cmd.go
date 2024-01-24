package cmd

import (
	"os"

	"github.com/nicolerobin/log"
	"github.com/pingcap/tidb/pkg/parser"
	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sql_gen",
	Short: "sql_gen is a code generate tool",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute the entry of all commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error("rootCmd.Execute() failed, err:%s", err)
		os.Exit(1)
	}
}

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()
	stmtNodes, _, err := p.ParseSQL(sql)
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}
