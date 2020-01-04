// Package executil provides utilites for creating exec.Cmd objects.
package executil

import (
	"errors"
	"os/exec"
	"strings"
)

// ErrEmptyCommandString indicates that an empty command string was provided
// when a string containing a command was expected.
var ErrEmptyCommandString error = errors.New("command string was empty")

// ParseCmd parses a string containing an execution command (e.g.,
// `find . -name *.txt`) into an exec.Cmd object.
func ParseCmd(cmdStr string) (*exec.Cmd, error) {
	if cmdStr == "" {
		return nil, ErrEmptyCommandString
	}

	process, args := parseCmdIntoProcessAndArgs(cmdStr)

	cmd := exec.Command(process, args...)
	return cmd, nil
}

func parseCmdIntoProcessAndArgs(cmdStr string) (string, []string) {
	parts := strings.Split(cmdStr, " ")
	process := parts[0]
	args := []string{}

	if len(parts) > 1 {
		args = parts[1:]
	}

	return process, args
}

// CloneCmd clones the provided cmd into a new cmd.
func CloneCmd(cmd *exec.Cmd) *exec.Cmd {
	newCmd := &exec.Cmd{
		Path:        cmd.Path,
		Args:        cmd.Args,
		Env:         cmd.Env,
		Dir:         cmd.Dir,
		ExtraFiles:  cmd.ExtraFiles,
		SysProcAttr: cmd.SysProcAttr,
	}

	return newCmd
}
