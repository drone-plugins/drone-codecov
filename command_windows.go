//go:build windows
// +build windows

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func (_ *Plugin) command(args []string) *exec.Cmd {
	fmt.Println("$ codecov.exe", strings.Join(args, " "))

	return exec.Command(
		"codecov.exe",
		args...,
	)
}
