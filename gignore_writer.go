package main

import (
	"os"
)

func WriteResponseBodyToFile(b []byte) error {
	const op Operation = "Write Response Body To File"

	if len(b) == 0 {
		return NewError(op, ResponseBodyIsEmpty)
	}

	file, err := os.Create(".gitignore")
	if err != nil {
		return NewError(op, UnableToCreateGitignoreFile, err)
	}

	_, err = file.Write(b)
	if err != nil {
		return NewError(op, UnableToWriteGitignoreFile, err)
	}

	return nil
}
