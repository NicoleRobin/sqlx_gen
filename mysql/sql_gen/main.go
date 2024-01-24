package main

import (
	_ "github.com/pingcap/tidb/pkg/parser/test_driver"

	"learngo/mysql/sql_gen/cmd"
)

func main() {
	cmd.Execute()
}
