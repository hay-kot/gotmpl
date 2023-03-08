// Package commands contains the CLI commands for the application
package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hay-kot/gotmpl/app/commands/engine"
	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type Vars map[string]interface{}

type Controller struct {
	LogLevel string

	Template string
	DataFile string
}

func (c *Controller) Render(ctx *cli.Context) error {
	ext := filepath.Ext(c.DataFile)

	var rawData any

	f, err := os.Open(c.DataFile)
	if err != nil {
		return err
	}

	switch ext {
	case ".json":
		err = json.NewDecoder(f).Decode(&rawData)
	case ".yaml", ".yml":
		err = yaml.NewDecoder(f).Decode(&rawData)
	case ".toml":
		err = toml.NewDecoder(f).Decode(&rawData)
	default:
		return fmt.Errorf("unsupported data file format: %s", ext)
	}

	if err != nil {
		return err
	}

	vars := Vars{"Data": rawData}

	e := engine.New()

	tmplFile, err := os.Open(c.Template)
	if err != nil {
		return err
	}

	tmpl, err := e.Factory(tmplFile)
	if err != nil {
		return err
	}

	err = e.Render(os.Stdout, tmpl, vars)
	if err != nil {
		return err
	}

	return nil
}
