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
		&cli.StringFlag{
			Name:    "repo.fullname",
			Usage:   "repository full name",
			EnvVars: []string{"CI_REPO"},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Value:   "master",
			Usage:   "commit branch",
			EnvVars: []string{"CI_COMMIT_SOURCE_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "commit.sha",
			Usage:   "commit sha",
			EnvVars: []string{"CI_COMMIT_SHA"},
		},
		&cli.StringFlag{
			Name:    "commit.tag",
			Usage:   "commit tag",
			EnvVars: []string{"CI_COMMIT_TAG"},
		},
		&cli.IntFlag{
			Name:    "build.number",
			Usage:   "build number",
			EnvVars: []string{"CI_BUILD_NUMBER"},
		},
		&cli.StringFlag{
			Name:    "build.link",
			Usage:   "build link",
			EnvVars: []string{"CI_BUILD_LINK"},
		},
		&cli.IntFlag{
			Name:    "pull.request",
			Usage:   "pull request",
			EnvVars: []string{"CI_COMMIT_PULL_REQUEST"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Fullname: c.String("repo.fullname"),
		},
		Build: Build{
			Number:      c.Int("build.number"),
			Link:        c.String("build.link"),
			PullRequest: c.Int("pull.request"),
		},
		Commit: Commit{
			Sha:    c.String("commit.sha"),
			Branch: c.String("commit.branch"),
			Tag:    c.String("commit.tag"),
		},
		Config: Config{
			Token:    c.String("token"),
			Name:     c.String("name"),
			Pattern:  c.String("pattern"),
			Paths:    c.StringSlice("path"),
			Files:    c.StringSlice("file"),
			Flags:    c.StringSlice("flag"),
			Env:      c.StringSlice("env"),
			Verbose:  c.Bool("verbose"),
			DryRun:   c.Bool("dry_run"),
			Required: c.Bool("required"),
		},
	}

	return plugin.Exec()
}
