package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora"
	"github.com/nicolerobin/sqlx_gen/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	BuildVersion = "v0.0.1"
	AppName      = "sqlx_gen"
)

var (
	rootCmd = cobra.Command{
		Use:   AppName,
		Short: "sqlx_gen is a tool to generate code for sqlx library",
		Run:   mysql.Ddl,
	}
)

func init() {
	rootCmd.Version = fmt.Sprintf("%s %s/%s", BuildVersion, runtime.GOOS, runtime.GOARCH)
	rootCmd.PersistentFlags().StringVarP(&mysql.VarStrSrc, "src", "s", "", "input sql file")
	_ = rootCmd.MarkPersistentFlagRequired("src")
	rootCmd.PersistentFlags().StringVarP(&mysql.VarStrDir, "dir", "d", "", "output generate code directory")
	_ = rootCmd.MarkPersistentFlagRequired("dir")
	_ = viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
	_ = viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(-1)
	}
}
