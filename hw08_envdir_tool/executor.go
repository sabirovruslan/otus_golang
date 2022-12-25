package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)
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
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		log.Fatalf("Wait run: %v", err)
	}
	return
}
