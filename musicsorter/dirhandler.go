package main

import (
	"log"
	"os"
	"path/filepath"
	// "github.com/dhowden/tag"
)

func visit(path_m string, info os.DirEntry, err error) error {
	if getsupportsuffix(filepath.Ext(path_m)) {
		files = append(files, path_m)
	}
	return nil
}

func Genlist(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, visit)
	if err != nil {
		log.Fatal(err)
	}
	return files, err
}
