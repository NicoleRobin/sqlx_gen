package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/converter"
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
		Comment      string
	}

	// Primary
	Primary struct {
		Field
		AutoIncrement bool
	}
)

// Parse 解析
func Parse(filename, database string, strict bool) ([]*Table, error) {
	log.Info("Parse(), filename:%s, database:%s", filename, database)
	p := parser.NewParser()
	tables, err := p.From(filename)
	if err != nil {
		log.Error("p.From() failed, err:%s", err)
		return nil, err
	}

	var list []*Table
	for i, t := range tables {
		log.Info("i:%d, table:%+v", i, *t)
		var (
			primaryColumn    string
			primaryColumnSet = make(map[string]struct{})
			uniqueKeyMap     = make(map[string][]string)
			normalKeyMap     = make(map[string][]string)
			columns          = t.Columns
		)

		// 查看各个列的约束
		for _, column := range columns {
			log.Info("column:%+v", *column)
			if column.Constraint != nil {
				if column.Constraint.Primary {
					primaryColumn = column.Name
					primaryColumnSet[column.Name] = struct{}{}
				}

				if column.Constraint.Unique {
					indexName := fmt.Sprintf("%s_%s", column.Name, "unique")
					uniqueKeyMap[indexName] = []string{column.Name}
				}

				if column.Constraint.Key {
					indexName := fmt.Sprintf("%s_%s", column.Name, "idx")
					normalKeyMap[indexName] = []string{column.Name}
				}
			}
		}

		// 查看表的约束
		for _, constraint := range t.Constraints {
			log.Info("constraint:%+v", constraint)
			if len(constraint.ColumnPrimaryKey) > 1 {
				return nil, fmt.Errorf("%s: unexpected join primary key", filepath.Base(filename))
			}

			if len(constraint.ColumnPrimaryKey) == 1 {
				primaryColumn = constraint.ColumnPrimaryKey[0]
			}

			// 为什么没有普通索引的信息呢？
			if len(constraint.ColumnUniqueKey) > 0 {
				// 使用nil的方式清空list
				list := append([]string(nil), constraint.ColumnUniqueKey...)
				list = append(list, "unique")
				indexName := strings.Join(list, "_")
				uniqueKeyMap[indexName] = constraint.ColumnUniqueKey
			}
		}

		// 为什么不支持多主键呢？
		if len(primaryColumnSet) > 1 {
			return nil, fmt.Errorf("%s: unexpected join primary key", filepath.Base(filename))
		}

		primaryKey, fieldM, err := convertColumns(columns, primaryColumn, strict)
		if err != nil {
			return nil, err
		}

		var fields []*Field
		for _, column := range columns {
			field, ok := fieldM[column.Name]
			if ok {
				fields = append(fields, field)
			}
		}

		var (
			uniqueIndex = make(map[string][]*Field)
			normalIndex = make(map[string][]*Field)
		)

		for indexName, each := range uniqueKeyMap {
			for _, columnName := range each {
				uniqueIndex[indexName] = append(uniqueIndex[indexName], fieldM[columnName])
			}
		}

		for indexName, each := range normalKeyMap {
			for _, columnName := range each {
				normalIndex[indexName] = append(normalIndex[indexName], fieldM[columnName])
			}
		}

		list = append(list, &Table{
			Name:        t.Name,
			Db:          database,
			PrimaryKey:  primaryKey,
			UniqueIndex: uniqueIndex,
			Fields:      fields,
		})
	}
	return list, nil
}

// convertColumns convert parser.Column to self defined field
func convertColumns(columns []*parser.Column, primaryColumn string, strict bool) (Primary, map[string]*Field, error) {
	var (
		primaryKey Primary
		fieldM     = make(map[string]*Field)
	)

	for _, column := range columns {
		if column == nil {
			continue
		}

		var (
			comment       string
			isDefaultNull bool
		)

		if column.Constraint != nil {
			comment = column.Constraint.Comment
			isDefaultNull = !column.Constraint.NotNull
			if !column.Constraint.NotNull && column.Constraint.HasDefaultValue {
				isDefaultNull = false
			}
			if column.Name == primaryColumn {
				isDefaultNull = false
			}
		}

		dataType, err := converter.ConvertDataType(column.DataType.Type(), isDefaultNull, column.DataType.Unsigned(), strict)
		if err != nil {
			return Primary{}, nil, err
		}

		if column.Constraint != nil {
			if column.Name == primaryColumn {
				if !column.Constraint.AutoIncrement && dataType == "int64" {
					log.Warn("The column %q is recommended to add constraint `auto_increment`", column.Name)
				}
			} else if column.Constraint.NotNull && !column.Constraint.HasDefaultValue {
				log.Warn("The column %q is recommended to add constraint `default`", column.Name)
			}
		}

		var field Field
		field.Name = column.Name
		field.DataType = dataType
		field.Comment = comment

		if field.Name == primaryColumn {
			primaryKey = Primary{
				Field: field,
			}

			if column.Constraint != nil {
				primaryKey.AutoIncrement = column.Constraint.AutoIncrement
			}
		}

		fieldM[field.Name] = &field
	}

	return primaryKey, fieldM, nil
}
