package visitor

import (
	"github.com/nicolerobin/log"
	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/types"
)

type ColumnDef struct {
	Name string
	Type types.EvalType
}

type TableDef struct {
	Name    string
	Columns []ColumnDef
}

type CreateTableVisitor struct {
	tables []TableDef
}

func (v *CreateTableVisitor) Enter(in ast.Node) (ast.Node, bool) {
	// log.Info("Enter(), in.Text():%s", in.Text())
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		tableDef := TableDef{
			Name: node.Table.Name.O,
		}

		log.Info("ast.CreateTableStmt, table name:%s", node.Table.Name.O)
		for _, col := range node.Cols {
			tableDef.Columns = append(tableDef.Columns, ColumnDef{
				Name: col.Name.Name.O,
				Type: col.Tp.EvalType(),
			})
		}
		v.tables = append(v.tables, tableDef)
	}

	return in, false
}

func (v *CreateTableVisitor) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

// Extract generate TableDef array from sql
func Extract(rootNode *ast.StmtNode) []TableDef {
	v := &CreateTableVisitor{}
	(*rootNode).Accept(v)
	return v.tables
}
