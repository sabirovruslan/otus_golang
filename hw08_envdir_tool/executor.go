package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Stdout = os.Stdout

	for k, v := range env {
		_, ok := os.LookupEnv(k)
		if v.NeedRemove {
			if ok {
				os.Unsetenv(k)
			}
			continue
		}
		os.Setenv(k, v.Value)
	}
	command.Env = os.Environ()

	if err := command.Start(); err != nil {
		log.Fatalf("Start command: %v", err)
	}

	if err := command.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		log.Fatalf("Wait run: %v", err)
	}
	return
}
