package tools

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func CreateDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return errors.New(fmt.Sprintf("failed to create dir %s: %s", path, err))
	}
	return nil
}

func CreateFile(path string) (*os.File, error) {
	// dir
	err := CreateDir(filepath.Dir(path))
	if err != nil {
		return nil, err
	}
	// file
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open file %s: %s", path, err))
	}
	return file, nil
}
