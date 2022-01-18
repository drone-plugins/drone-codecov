//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func (_ *Plugin) command(args []string) *exec.Cmd {
	fmt.Println("$", strings.Join(args, " "))

	cmd := fmt.Sprintf("/bin/codecov %s", strings.Join(args, " "))
	fmt.Printf("$ %s", cmd)

	return exec.Command("/bin/codecov",
		args...,
	)
}
