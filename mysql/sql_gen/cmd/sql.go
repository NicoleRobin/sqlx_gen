package cmd

import (
	"github.com/nicolerobin/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"

	"learngo/mysql/sql_gen/visitor"
)

var sql = &cobra.Command{
	Use:   "sql",
	Short: "generate code from sql",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		sqlContent, err := os.ReadFile(sqlFile)
		if err != nil {
			log.Error("os.ReadFile() failed", zap.Error(err))
			return
		}

		stmtNode, err := parse(string(sqlContent))
		if err != nil {
			log.Error("parse() failed, err:%s", err)
			return
		}

		log.Info("stmtNode:%+v", *stmtNode)
		tables := visitor.Extract(stmtNode)
		log.Info("tables:%+v", tables)
	},
}

var sqlFile string

func init() {
	sql.PersistentFlags().StringVarP(&sqlFile, "sql", "s", "", "sql file path")

	rootCmd.AddCommand(sql)
}
