package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		line, err := readLine(filepath.Join(dir, entry.Name()))
		if err != nil {
			return nil, err
		}
		line = bytes.ReplaceAll(line, []byte("\x00"), []byte("\n"))
		v := strings.TrimRight(string(line), " ")
		if len(v) == 0 {
			envs[entry.Name()] = EnvValue{"", true}
		} else {
			envs[entry.Name()] = EnvValue{v, false}
		}
	}
	return envs, nil
}

func readLine(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	line, _, err := scanner.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, nil
		}
		return nil, err
	}
	return line, nil
}
