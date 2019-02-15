package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "codecov plugin"
	app.Usage = "codecov plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "token for authentication",
			EnvVar: "PLUGIN_TOKEN,CODECOV_TOKEN",
		},
		cli.StringFlag{
			Name:   "name",
			Usage:  "name for coverage upload",
			EnvVar: "PLUGIN_NAME",
		},
		cli.StringSliceFlag{
			Name:   "path",
			Usage:  "paths for searching for coverage files",
			EnvVar: "PLUGIN_PATHS",
		},
		cli.StringSliceFlag{
			Name:   "file",
			Usage:  "files for coverage upload",
			EnvVar: "PLUGIN_FILES",
		},
		cli.StringSliceFlag{
			Name:   "flag",
			Usage:  "flags for coverage upload",
			EnvVar: "PLUGIN_FLAGS",
		},
		cli.StringSliceFlag{
			Name:   "env",
			Usage:  "inject environment",
			EnvVar: "PLUGIN_ENV",
		},
		cli.BoolFlag{
			Name:   "dump",
			Usage:  "dump instead of upload",
			EnvVar: "PLUGIN_DUMP",
		},
		cli.BoolFlag{
			Name:   "verbose",
			Usage:  "print verbose output",
			EnvVar: "PLUGIN_VERBOSE",
		},
		cli.BoolTFlag{
			Name:   "required",
			Usage:  "errors on failed upload",
			EnvVar: "PLUGIN_REQUIRED",
		},
		cli.StringFlag{
			Name:   "repo.fullname",
			Usage:  "repository full name",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "commit branch",
			EnvVar: "DRONE_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "commit sha",
			EnvVar: "DRONE_COMMIT",
		},
		cli.StringFlag{
			Name:   "commit.tag",
			Usage:  "commit tag",
			EnvVar: "DRONE_TAG",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.IntFlag{
			Name:   "pull.request",
			Usage:  "pull request",
			EnvVar: "DRONE_PULL_REQUEST",
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
			Dump:     c.Bool("dump"),
			Verbose:  c.Bool("verbose"),
			Required: c.Bool("required"),
		},
	}

	return plugin.Exec()
}
