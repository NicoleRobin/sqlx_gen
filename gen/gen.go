package gen

import (
	"fmt"
	"github.com/nicolerobin/log"
	"github.com/nicolerobin/sqlx_gen/parser"
	"path/filepath"
	"strings"
)

type (
	// Generator generator
	Generator struct {
		dir          string
		isPostgreSql bool
	}

	// Table
	Table struct {
		parser.Table
		PrimaryCacheKey        Key
		UniqueCacheKey         []Key
		ContainsUniqueCacheKey bool
		ignoreColumns          []string
	}

	// Option defines a function with argument defaultGenerator
	Option func(generator *defaultGenerator)

	code struct {
		importCode string
		varsCode   string
		typesCode  string
		newCode    string
		insertCode string
		findCode   string
		updateCode string
		deleteCode string
		cacheExtra string
		tableName  string
	}

	codeTuple struct {
		modelCode       string
		modelCustomCode string
	}
)

// # TODO: i think it is not good to return an error in constructor func
// NewGenerator create generator for generate code
func NewGenerator(dir string) (*Generator, error) {
	if dir == "" {
		dir = "."
	}
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	return &Generator{
		dir: absDir,
	}, nil
}

func (g *Generator) StartFromDDL(filename string, database string) error {
	log.Info("StartFromDDL(), filename:%s, database:%s", filename, database)

	modelList, err := g.genFromDDL(filename, database)
	if err != nil {
		return err
	}
	return g.createFile(modelList)
}

func (g *Generator) genFromDDL(filename, database string) (map[string]*codeTuple, error) {
	log.Info("genFromDDL(), filename:%s, database:%s", filename, database)

	tables, err := parser.Parse(filename, database, true)
	if err != nil {
		log.Error("parser.Parse() failed, err:%s", err)
		return nil, err
	}

	m := make(map[string]*codeTuple)
	for i, table := range tables {
		log.Info("i:%d, table:%+v", i, *table)
		code, err := g.genModel(table, false)
		if err != nil {
			log.Error("g.genModel() failed, err:%s", err)
			continue
		}

		m[table.Name] = &codeTuple{
			modelCode: code,
		}
	}
	return m, nil
}

func (g *Generator) createFile(modelList map[string]*codeTuple) error {
	log.Info("createFile, modelList:%+v", modelList)
	for tableName, code := range modelList {
		log.Info("tableName:%s, modelCode:%s, modelCustomCode:%s", tableName, code.modelCode,
			code.modelCustomCode)
	}
	return nil
}

func (g *Generator) genModel(in *parser.Table, withCache bool) (string, error) {
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
		return "", err
	}

	insertCode, insertCodeMethod, err := genInsert(table, withCache, g.isPostgreSql)
	if err != nil {
		return "", err
	}

	findCode := make([]string, 0)
	findOneCode, findOneCodeMethod, err := genFindOne(table, withCache, g.isPostgreSql)
	if err != nil {
		return "", err
	}

	ret, err := genFindOneByField(table, withCache, g.isPostgreSql)
	if err != nil {
		return "", err
	}

	findCode = append(findCode, findOneCode, ret.findOneMethod)
	updateCode, updateCodeMethod, err := genUpdate(table, withCache, g.isPostgreSql)
	if err != nil {
		return "", err
	}

	var list []string
	list = append(list, insertCodeMethod, findOneCodeMethod, ret.findOneInterfaceMethod,
		updateCodeMethod, deleteCodeMethod)
	typesCode, err := genTypes(table, strings.Join(list, "\n"), withCache)
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

	code := &code{}

	return importsCode, nil
}
