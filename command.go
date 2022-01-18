//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func (_ *Plugin) command(args []string) *exec.Cmd {
	cmd := fmt.Sprintf("/bin/codecov %s", strings.Join(args, " "))
	fmt.Printf("$ %s\n", cmd)

	return exec.Command("/bin/codecov",
		args...,
	)
}
