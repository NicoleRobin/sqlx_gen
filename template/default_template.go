package template

import (
	"bytes"
	"fmt"
	goformat "go/format"
	"os"
	"text/template"
)

// DefaultTemplate
type DefaultTemplate struct {
	name  string
	text  string
	goFmt bool
}

func With(name string) *DefaultTemplate {
	return &DefaultTemplate{
		name: name,
	}
}

func (t *DefaultTemplate) Parse(text string) *DefaultTemplate {
	t.text = text
	return t
}

func (t *DefaultTemplate) GoFmt(format bool) *DefaultTemplate {
	t.goFmt = format
	return t
}

// SaveTo
func (t *DefaultTemplate) SaveTo(data any, path string, forceUpdate bool) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("path:%s not exist, please check", path)
	}

	output, err := t.Execute(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, output.Bytes(), regularPerm)
}

// Execute
func (t *DefaultTemplate) Execute(data any) (*bytes.Buffer, error) {
	tem, err := template.New(t.name).Parse(t.text)
	if err != nil {
		return nil, fmt.Errorf("template parse failed, err:%s", err)
	}

	var buf bytes.Buffer
	if err = tem.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("template execute failed, err:%s", err)
	}
	if !t.goFmt {
		return &buf, nil
	}

	formatOutput, err := goformat.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("goformat.Source() failed, err:%s", err)
	}
	buf.Reset()
	buf.Write(formatOutput)
	return &buf, nil
}
