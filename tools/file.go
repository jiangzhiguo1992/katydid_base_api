package tools

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func DirCreate(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return errors.New(fmt.Sprintf("failed to create dir %s: %s", path, err))
	}
	return nil
}

func FileCreate(path string) (*os.File, error) {
	// dir
	err := DirCreate(filepath.Dir(path))
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

func FileSizeFormat(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}
