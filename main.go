package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "codecov plugin"
	app.Usage = "codecov plugin"
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Action = run
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
		cli.StringFlag{
			Name:   "pattern",
			Usage:  "coverage file pattern",
			Value:  "**/coverage.out",
			EnvVar: "PLUGIN_PATTERN",
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
			Tag:    c.String("commit.branch"),
		},
		Config: Config{
			Token:   c.String("token"),
			Name:    c.String("name"),
			Pattern: c.String("pattern"),
			Files:   c.StringSlice("file"),
			Flags:   c.StringSlice("flag"),
		},
	}

	return plugin.Exec()
}
