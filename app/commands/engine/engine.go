// Package engine provides a simple template engine for Go Templates.
package engine

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	"github.com/rs/zerolog/log"
)

// ErrTemplateIsEmpty is returned when a provided reader is empty.
var ErrTemplateIsEmpty = fmt.Errorf("template is empty")

type Vars map[string]any

type Engine struct {
	baseTemplate *template.Template
}

func New() *Engine {
	fm := sprig.FuncMap()

	fm["wraptmpl"] = wraptmpl

	return &Engine{
		baseTemplate: template.New("template").Funcs(fm),
	}
}

func isTemplate(s string) bool {
	return strings.Contains(s, "{{")
}

func (e *Engine) TmplString(str string, vars any) (string, error) {
	if !isTemplate(str) {
		return str, nil
	}

	tmpl, err := e.baseTemplate.Parse(str)
	if err != nil {
		log.Err(err).Msg("failed to parse template")
		return "", err
	}

	out := &strings.Builder{}

	err = e.Render(out, tmpl, vars)
	if err != nil {
		log.Err(err).Msg("failed to render template")
		return "", err
	}

	return out.String(), nil
}

// Factory returns a new template from the provided reader.
// if the reader is empty, an ErrTemplateIsEmpty is returned.
func (e *Engine) Factory(reader io.Reader) (*template.Template, error) {
	if reader == nil {
		return nil, fmt.Errorf("reader is nil")
	}

	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, ErrTemplateIsEmpty
	}

	return e.baseTemplate.Parse(string(out))
}

func (e *Engine) Render(w io.Writer, tmpl *template.Template, vars any) error {
	err := tmpl.Execute(w, vars)
	if err != nil {
		return err
	}

	return nil
}

func IsValidIdentifier(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return true
}
