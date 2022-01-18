package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
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
		Required bool
	}

	Internal struct {
		Matches []string
		Merged  []byte
	}

	Plugin struct {
		Repo     Repo
		Build    Build
		Commit   Commit
		Config   Config
		Internal Internal
	}
)

func (p *Plugin) Exec() error {
	if p.Config.Token == "" {
		return errors.New("you must provide a token")
	}

	if p.Commit.Sha == "" {
		return errors.New("you must provide a commit")
	}

	cmd := p.command()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("CODECOV_TOKEN=%s", p.Config.Token),
	)

	return cmd.Run()
}
