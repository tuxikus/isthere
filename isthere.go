package isthere

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	version = "v1.0"
)

func GetVersion() string {
	return version
}

func IsThere(name string) (string, error) {
	pathEnv := os.Getenv("PATH")
	paths := filepath.SplitList(pathEnv)
	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.Name() == name {
				return path + "/" + name, nil
			}
		}
	}

	return "Unable find " + name + " in: " + pathEnv, errors.New("cannot find command")
}
