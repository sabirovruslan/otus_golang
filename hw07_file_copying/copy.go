package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	infoInFile, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if infoInFile.Size() < offset {
		return ErrOffsetExceedsFileSize
	}
	if infoInFile.Size() == 0 {
		return ErrUnsupportedFile
	}

	if limit <= 0 {
		limit = infoInFile.Size()
	}
	if offset < 0 {
		offset = int64(0)
	}

	inFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	bar := pb.Start64(limit)
	defer bar.Finish()

	bar.Start()
	buf := make([]byte, 1)
	var currentSize int64
	for currentSize < limit {
		n, err := inFile.ReadAt(buf, offset)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		_, err = outFile.Write(buf[:n])
		if err != nil {
			return err
		}
		currentSize += int64(n)
		offset += int64(n)
		bar.Increment()
	}

	return nil
}
