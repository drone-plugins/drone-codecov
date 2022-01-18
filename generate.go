package main

import (
	"os"
	"strings"
)

func (p *Plugin) generateArgs() []string {
	args := []string{"-Q", "'Woodpecker Plugin'"}

	if path, err := os.Getwd(); err == nil {
		args = append(args, "--rootDir", path)
	}

	if p.Name != "" {
		args = append(args, "--name", p.Name)
	}

	if len(p.Flags) != 0 {
		args = append(args, "--flags", strings.Join(p.Flags, ","))
	}

	if len(p.Env) != 0 {
		args = append(args, "--env", strings.Join(p.Env, ","))
	}

	if p.DryRun {
		args = append(args, "--dryRun")
	}

	if p.Required {
		args = append(args, "--nonZero")
	}

	for _, file := range p.Files {
		args = append(args, "--file", file)
	}

	for _, path := range p.Paths {
		args = append(args, "--dir", path)
	}

	return args
}
