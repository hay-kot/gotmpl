package main

import (
	"fmt"

	"os"

	"github.com/hay-kot/gotmpl/app/commands"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	// Build information. Populated at build-time via -ldflags flag.
	version = "dev"
	commit  = "HEAD"
	date    = "now"
)

func build() string {
	short := commit
	if len(commit) > 7 {
		short = commit[:7]
	}

	return fmt.Sprintf("%s (%s) %s", version, short, date)
}

func main() {
	ctrl := &commands.Controller{}

	app := &cli.App{
		Name:    "gotmpl",
		Usage:   "Tiny CLI template engine for generating files quickly",
		Version: build(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Usage: "log level (debug, info, warn, error, fatal, panic)",
				Value: "panic",
			},
		},
		Before: func(ctx *cli.Context) error {
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

			switch ctx.String("log-level") {
			case "debug":
				log.Level(zerolog.DebugLevel)
			case "info":
				log.Level(zerolog.InfoLevel)
			case "warn":
				log.Level(zerolog.WarnLevel)
			case "error":
				log.Level(zerolog.ErrorLevel)
			case "fatal":
				log.Level(zerolog.FatalLevel)
			default:
				log.Level(zerolog.PanicLevel)
			}

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:   "render",
				Usage:  "render template",
				Action: ctrl.Render,
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:        "template",
						Aliases:     []string{"tmpl", "t"},
						Usage:       "template file",
						Required:    true,
						Destination: &ctrl.Template,
					},
					&cli.PathFlag{
						Name:        "data",
						Aliases:     []string{"d"},
						Usage:       "data file",
						Required:    true,
						Destination: &ctrl.DataFile,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run gotmpl")
	}
}
