// +build windows

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func (p *Plugin) command() *exec.Cmd {
	args := []string{
		"--no-color",
		fmt.Sprintf("-c %s", p.Commit.Sha),
	}

	if p.Commit.Branch != "" {
		args = append(args, fmt.Sprintf("--branch %s", p.Commit.Branch))
	}

	if p.Commit.Tag != "" {
		args = append(args, fmt.Sprintf("--tag %s", p.Commit.Tag))
	}

	if p.Build.PullRequest != 0 {
		args = append(args, fmt.Sprintf("--pr %s", strconv.Itoa(p.Build.PullRequest)))
	}

	if p.Build.Number != 0 {
		args = append(args, fmt.Sprintf("-b %s", strconv.Itoa(p.Build.Number)))
	}

	if p.Config.Name != "" {
		args = append(args, fmt.Sprintf("-n %s", p.Config.Name))
	}

	if len(p.Config.Flags) != 0 {
		args = append(args, fmt.Sprintf("--flag %s", strings.Join(p.Config.Flags, ",")))
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
		args = append(args, "--required")
	}

	for _, file := range p.Config.Files {
		args = append(args, fmt.Sprintf("-f '%s'", file))
	}

	fmt.Println("$ codecov.exe", strings.Join(args, " "))

	return exec.Command(
		"codecov.exe",
		args...,
	)
}
