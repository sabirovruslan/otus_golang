package main

import (
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type copySuite struct {
	suite.Suite
}

func TestRunCopySuite(t *testing.T) {
	suite.Run(t, new(copySuite))
}

func (s *copySuite) TestCopy() {
	s.Error(Copy("", "", 0, 0))
	s.EqualError(Copy("testdata/input.txt", "", 0, 0), "open : no such file or directory")
	s.EqualError(Copy("not_exist.txt", "/tmp/out.txt", 0, 0), "stat not_exist.txt: no such file or directory")
	s.EqualError(Copy("testdata/input.txt", "/tmp/out.txt", math.MaxInt64, 0), "offset exceeds file size")
	s.EqualError(Copy("/dev/urandom", "/tmp/out.txt", 0, 0), "unsupported file")

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", 0, 0), nil)
	_, err := os.Stat("/tmp/out.txt")
	s.Equal(err, nil)

	buf := make([]byte, 20)
	infoIn, _ := os.Stat("testdata/input.txt")

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", 0, 1), nil)
	infoOut, _ := os.Stat("/tmp/out.txt")
	s.Equal(infoOut.Size(), int64(1))
	file, _ := os.Open("/tmp/out.txt")
	n, _ := file.Read(buf)
	s.Equal(string(buf[:n]), "G")
	file.Close()

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", 0, 2), nil)
	infoOut, _ = os.Stat("/tmp/out.txt")
	s.Equal(infoOut.Size(), int64(2))
	file, _ = os.Open("/tmp/out.txt")
	n, _ = file.Read(buf)
	s.Equal(string(buf[:n]), "Go")
	file.Close()

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", 0, 5), nil)
	infoOut, _ = os.Stat("/tmp/out.txt")
	s.Equal(infoOut.Size(), int64(5))
	file, _ = os.Open("/tmp/out.txt")
	n, _ = file.Read(buf)
	s.Equal(string(buf[:n]), "Go\nDo")
	file.Close()

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", 3, 2), nil)
	infoOut, _ = os.Stat("/tmp/out.txt")
	s.Equal(infoOut.Size(), int64(2))
	file, _ = os.Open("/tmp/out.txt")
	n, _ = file.Read(buf)
	s.Equal(string(buf[:n]), "Do")
	file.Close()

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", 0, 0), nil)
	infoOut, _ = os.Stat("/tmp/out.txt")
	s.Equal(infoOut.Size(), infoIn.Size())

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", -1, -1), nil)
	infoOut, _ = os.Stat("/tmp/out.txt")
	s.Equal(infoOut.Size(), infoIn.Size())

	s.Equal(Copy("testdata/input.txt", "/tmp/out.txt", infoIn.Size(), 0), nil)
	infoOut, _ = os.Stat("/tmp/out.txt")
	s.Equal(infoOut.Size(), int64(0))

}
