package gen

import "path/filepath"

type Generator struct {
	dir string
}

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
	modelList, err := g.genFromDDL(filename, database)
	if err != nil {
		return err
	}
	return g.createFile(modelList)
}

func (g *Generator) genFromDDL(filename, database) ([]model, error) {
	return nil
}

func (g *Generator) createFile(models []Model) error {
	return nil
}
