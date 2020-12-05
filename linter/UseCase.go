package linter

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type UseCase interface {
	ReadAndParseFile(ctx context.Context) (*EnvFile, error)
	CheckFile(ctx context.Context, f *EnvFile)
	FixFile(ctx context.Context, f *EnvFile)
	BackupFile(ctx context.Context, f *EnvFile)
}

type EnvFile struct {
	file *os.File
	path string
}

// https://github.com/ilyakaznacheev/cleanenv/blob/master/cleanenv.go
func parseFile(path string, cfg interface{}) error {
	// open  file
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	// parse the file depending on the file type
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".env":
		err = parseENV(f, cfg)
	default:
		return fmt.Errorf("file format '%s' doesn't supported by the parser", ext)
	}
	if err != nil {
		return fmt.Errorf("parsing error: %s", err.Error())
	}
	return nil
}

func parseENV(f *os.File, cfg interface{}) error {
	return nil
}