package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (p *Plugin) generateArgs() []string {
	args := []string{
		"/bin/codecov",
		fmt.Sprintf("-C %s", p.Commit.Sha),
	}

	if p.Commit.Branch != "" {
		args = append(args, fmt.Sprintf("-B %s", p.Commit.Branch))
	}

	if p.Commit.Tag != "" {
		args = append(args, fmt.Sprintf("-T %s", p.Commit.Tag))
	}

	if p.Build.PullRequest != 0 {
		args = append(args, fmt.Sprintf("-P %s", strconv.Itoa(p.Build.PullRequest)))
	}

	if p.Build.Number != 0 {
		args = append(args, fmt.Sprintf("-b %s", strconv.Itoa(p.Build.Number)))
	}

	if p.Config.Name != "" {
		args = append(args, fmt.Sprintf("-n %s", p.Config.Name))
	}

	if len(p.Config.Flags) != 0 {
		args = append(args, fmt.Sprintf("-F %s", strings.Join(p.Config.Flags, ",")))
	}

	if len(p.Config.Env) != 0 {
		args = append(args, fmt.Sprintf("-e %s", strings.Join(p.Config.Env, ",")))
	}

	if p.Config.DryRun {
		args = append(args, "--dryRun")
	}

	if p.Config.Required {
		args = append(args, "--nonZero")
	}

	for _, file := range p.Config.Files {
		args = append(args, fmt.Sprintf("-f '%s'", file))
	}

	for _, path := range p.Config.Paths {
		args = append(args, fmt.Sprintf("-s '%s'", path))
	}

	return args
}
