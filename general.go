// +build !windows

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func (p *Plugin) command() *exec.Cmd {
	args := []string{
		"/bin/codecov",
		"-K",
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

	if p.Config.Verbose {
		args = append(args, "-v")
	}

	if p.Config.Dump {
		args = append(args, "-d")
	}

	if p.Config.Required {
		args = append(args, "-Z")
	}

	for _, file := range p.Config.Files {
		args = append(args, fmt.Sprintf("-f '%s'", file))
	}

	fmt.Println("$", strings.Join(args, " "))

	return exec.Command(
		"bash",
		"-c",
		strings.Join(args, " "),
	)
}
