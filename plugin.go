package main

import (
	"errors"
	"fmt"
	"os"
)

type (
	Repo struct {
		Fullname string
	}

	Build struct {
		Number      int
		Link        string
		PullRequest int
	}

	Commit struct {
		Sha    string
		Branch string
		Tag    string
	}

	Config struct {
		Token    string
		Name     string
		Pattern  string
		Files    []string
		Paths    []string
		Flags    []string
		Env      []string
		Verbose  bool
		DryRun   bool
		Required bool
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
	}
)

func (p *Plugin) Exec() error {
	if p.Config.Token == "" {
		return errors.New("you must provide a token")
	}

	if p.Commit.Sha == "" {
		return errors.New("you must provide a commit")
	}

	args := p.generateArgs()
	cmd := p.command(args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("CODECOV_TOKEN=%s", p.Config.Token),
	)

	return cmd.Run()
}
