package isthere

import (
	"os"
	"path/filepath"
)

const (
	version = "0.0.1"
)

type Info struct {
	Version string
}

func GetInfo() *Info {
	return &Info{
		Version: version,
	}
}

func IsThere(name string) string {
	pathEnv := os.Getenv("PATH")
	paths := filepath.SplitList(pathEnv)
	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.Name() == name {
				return path + "/" + name
			}
		}
	}

	return "Unable find " + name + " in: " + pathEnv
}
