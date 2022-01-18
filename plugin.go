package main

import (
	"errors"
	"fmt"
	"os"
)

type Plugin struct {
	Token    string
	Name     string
	Files    []string
	Paths    []string
	Flags    []string
	Env      []string
	Verbose  bool
	DryRun   bool
	Required bool
}

func (p *Plugin) Exec() error {
	if p.Token == "" {
		return errors.New("you must provide a token")
	}

	args := p.generateArgs()
	cmd := p.command(args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("CODECOV_TOKEN=%s", p.Token),
	)

	return cmd.Run()
}
