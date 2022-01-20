package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	version = "next"
)

func main() {
	app := cli.NewApp()
	app.Name = "codecov plugin"
	app.Usage = "codecov plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Usage:   "token for authentication",
			EnvVars: []string{"PLUGIN_TOKEN", "CODECOV_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "name",
			Usage:   "name for coverage upload",
			EnvVars: []string{"PLUGIN_NAME"},
		},
		&cli.StringSliceFlag{
			Name:    "path",
			Usage:   "paths for searching for coverage files",
			EnvVars: []string{"PLUGIN_PATHS"},
		},
		&cli.StringSliceFlag{
			Name:    "file",
			Usage:   "files for coverage upload",
			EnvVars: []string{"PLUGIN_FILES"},
		},
		&cli.StringSliceFlag{
			Name:    "flag",
			Usage:   "flags for coverage upload",
			EnvVars: []string{"PLUGIN_FLAGS"},
		},
		&cli.StringSliceFlag{
			Name:    "env",
			Usage:   "inject environment",
			EnvVars: []string{"PLUGIN_ENV"},
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Usage:   "print verbose output",
			EnvVars: []string{"PLUGIN_VERBOSE"},
		},
		&cli.BoolFlag{
			Name:    "dry_run",
			Usage:   "dont upload files",
			EnvVars: []string{"PLUGIN_DRY_RUN"},
		},
		&cli.BoolFlag{
			Name:    "required",
			Usage:   "errors on failed upload",
			EnvVars: []string{"PLUGIN_REQUIRED"},
			Value:   true,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Token:    c.String("token"),
		Name:     c.String("name"),
		Paths:    c.StringSlice("path"),
		Files:    c.StringSlice("file"),
		Flags:    c.StringSlice("flag"),
		Env:      c.StringSlice("env"),
		Verbose:  c.Bool("verbose"),
		DryRun:   c.Bool("dry_run"),
		Required: c.Bool("required"),
	}

	return plugin.Exec()
}
