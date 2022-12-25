package main

import (
	"log"
	"os"
)

func main() {
	pathEnv := os.Args[1]
	cmd := os.Args[2:]

	env, err := ReadDir(pathEnv)
	if err != nil {
		log.Fatal(err)
	}

	RunCmd(cmd, env)
}
