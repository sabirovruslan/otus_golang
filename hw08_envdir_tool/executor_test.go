package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type runCmdSuite struct {
	suite.Suite
}

func TestRunCmdSuite(t *testing.T) {
	suite.Run(t, new(runCmdSuite))
}

func (s *runCmdSuite) TestRunCmd() {
	command := []string{"ls", "-la"}
	s.Equal(RunCmd(command, nil), 0)

	command = []string{"ls"}
	s.Equal(RunCmd(command, nil), 0)

	command = []string{"ls", "la"}
	s.NotEqual(RunCmd(command, nil), 0)

	command = []string{"echo", "ENV"}
	envs, err := ReadDir("testdata/env")
	s.Equal(err, nil)
	s.Equal(RunCmd(command, envs), 0)
}
