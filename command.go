//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func (_ *Plugin) command(args []string) *exec.Cmd {
	fmt.Println("$ /bin/codecov", strings.Join(args, " "))

	return exec.Command("/bin/codecov",
		args...,
	)
}
