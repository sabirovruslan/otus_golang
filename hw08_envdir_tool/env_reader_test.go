package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type envSuite struct {
	suite.Suite
}

func TestRunEnvSuite(t *testing.T) {
	suite.Run(t, new(envSuite))
}

func (s *envSuite) TestRead() {
	_, err := ReadDir("")
	s.EqualError(err, "open : no such file or directory")

	envs, err := ReadDir("testdata/env")
	s.Equal(err, nil)
	s.Equal(len(envs), 5)

	envValue := envs["BAR"]
	s.Equal(envValue.Value, "bar")
	s.Equal(envValue.NeedRemove, false)

	envValue = envs["FOO"]
	s.Equal(envValue.Value, "   foo\nwith new line")
	s.Equal(envValue.NeedRemove, false)

	envValue = envs["UNSET"]
	s.Equal(envValue.Value, "")
	s.Equal(envValue.NeedRemove, true)

	envValue = envs["EMPTY"]
	s.Equal(envValue.Value, "")
	s.Equal(envValue.NeedRemove, true)

	envValue = envs["HELLO"]
	s.Equal(envValue.Value, "\"hello\"")
	s.Equal(envValue.NeedRemove, false)
}
