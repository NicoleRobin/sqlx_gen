package gen

import (
	"bytes"
	"fmt"
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/parser"
	"github.com/nicolerobin/sqlx_gen/template"
	"github.com/zeromicro/ddl-parser/console"
	"os"
	"path/filepath"
	"strings"
)

type (
	Generator interface {
	}
	// Generator generator
	defaultGenerator struct {
		console.Console
		dir           string
		pkg           string
		isPostgreSql  bool
		ignoreColumns []string
	}

	// Option defines a function with argument defaultGenerator
	Option func(generator *defaultGenerator)

	Code struct {
		importsCode string
		varsCode    string
		typesCode   string
		newCode     string
		insertCode  string
		findCode    []string
		updateCode  string
		deleteCode  string
		cacheExtra  string
		tableName   string
	}

	CodeTuple struct {
		modelCode       string
		modelCustomCode string
	}
)

// # TODO: i think it is not good to return an error in constructor func
// NewGenerator create generator for generate Code
func NewGenerator(dir string) (*defaultGenerator, error) {
	if dir == "" {
		dir = "."
	}
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	return &defaultGenerator{
		dir: absDir,
		pkg: filepath.Base(absDir),
	}, nil
}

func (g *defaultGenerator) StartFromDDL(filename, outputDir string, database string) error {
	log.Info("StartFromDDL(), filename:%s, database:%s", filename, database)

	modelList, err := g.genFromDDL(filename, database)
	if err != nil {
		return err
	}
	return g.createFile(outputDir, modelList)
}

func (g *defaultGenerator) genFromDDL(filename, database string) (map[string]*CodeTuple, error) {
	log.Info("genFromDDL(), filename:%s, database:%s", filename, database)

	tables, err := parser.Parse(filename, database, true)
	if err != nil {
		log.Error("parser.Parse() failed, err:%s", err)
		return nil, err
	}

	m := make(map[string]*CodeTuple)
	for i, table := range tables {
		log.Info("i:%d, table:%+v", i, *table)
		code, err := g.genModel(table, false)
		if err != nil {
			log.Error("g.genModel() failed, err:%s", err)
			continue
		}

		m[table.Name] = &CodeTuple{
			modelCode: code,
		}
	}
	return m, nil
}

func (g *defaultGenerator) createFile(outputDir string, modelList map[string]*CodeTuple) error {
	log.Info("createFile, outputDir:%s, modelList:%+v", outputDir, modelList)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			log.Error("os.MkdirAll() failed, err:%s", err)
			return err
		}
	}

	for tableName, code := range modelList {
		// log.Info("tableName:%s, modelCode:\n%s\n, modelCustomCode:\n%s", tableName, code.modelCode, code.modelCustomCode)

		codeFileName := fmt.Sprintf("%s.go", tableName)
		tableFilePath := filepath.Join(outputDir, codeFileName)
		file, err := os.Create(tableFilePath)
		if err != nil {
			log.Error("os.Create() failed, err:%s, path:%s", err, tableFilePath)
			continue
		}
		_, err = file.WriteString(code.modelCode)
		if err != nil {
			log.Error("file.WriteString() failed, err:%s", err)
		}
		file.Close()
	}
	return nil
}

func (g *defaultGenerator) executeModel(table Table, code *Code) (*bytes.Buffer, error) {
	t := template.With("model").Parse(template.ModelGen).GoFmt(true)
	output, err := t.Execute(map[string]any{
		"pkg":         g.pkg,
		"imports":     code.importsCode,
		"vars":        code.varsCode,
		"types":       code.typesCode,
		"new":         code.newCode,
		"insert":      code.insertCode,
		"find":        strings.Join(code.findCode, "\n"),
		"update":      code.updateCode,
		"delete":      code.deleteCode,
		"extraMethod": code.cacheExtra,
		"tableName":   code.tableName,
		"data":        table,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func genUpdate(table string, withCache, isPostgreSql bool) (string, string, error) {
	return "", "", nil
}

func genDelete(table string, withCache, isPostgreSql bool) (string, string, error) {
	return "", "", nil
}

func (g *defaultGenerator) genModel(in *parser.Table, withCache bool) (string, error) {
	log.Info("genModel(), in:%+v", *in)
	if len(in.PrimaryKey.Name) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name)
	}

	var table Table
	table.Name = in.Name

	importsCode, err := genImports(table, true)
	if err != nil {
		log.Error("genImports() failed, err:%s", err)
		return "", err
	}

	varsCode, err := genVars(table, withCache, g.isPostgreSql)
	if err != nil {
		log.Error("genVars() failed, err:%s", err)
		return "", err
	}

	insertCode, interfaceInsert, err := genInsert(table, withCache, g.isPostgreSql)
	if err != nil {
		log.Error("genInsert() failed, err:%s", err)
		return "", err
	}

	findCode := make([]string, 0)
	findOneCode, interfaceFindOne, err := genFindOne(table, withCache, g.isPostgreSql)
	if err != nil {
		log.Error("genFindOne() failed, err:%s", err)
		return "", err
	}

	ret, err := genFindOneByField(table, withCache, g.isPostgreSql)
	if err != nil {
		log.Error("genFindOneByField() failed, err:%s", err)
		return "", err
	}

	findCode = append(findCode, findOneCode, ret.findOneMethod)
	updateCode, interfaceUpdate, err := genUpdate(in.Name, withCache, g.isPostgreSql)
	if err != nil {
		return "", err
	}

	deleteCode, interfaceDelete, err := genDelete(in.Name, withCache, g.isPostgreSql)

	var list []string
	list = append(list, interfaceInsert, interfaceFindOne, ret.findOneInterfaceMethod,
		interfaceUpdate, interfaceDelete)
	typesCode, err := genTypes(table, strings.Join(list, NL), withCache)
	if err != nil {
		return "", err
	}

	newCode, err := genNew(table, withCache, g.isPostgreSql)
	if err != nil {
		return "", err
	}

	tableName, err := genTableName(table)
	if err != nil {
		return "", err
	}

	code := &Code{
		importsCode: importsCode,
		varsCode:    varsCode,
		typesCode:   typesCode,
		newCode:     newCode,
		insertCode:  insertCode,
		findCode:    findCode,
		updateCode:  updateCode,
		deleteCode:  deleteCode,
		cacheExtra:  ret.cacheExtra,
		tableName:   tableName,
	}

	output, err := g.executeModel(table, code)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
