package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var ignore = []string{".git", ".DS_Store", ".gitignore"}

func fileInIgnore(f string, ignore []string) bool {

	_, file := filepath.Split(f)

	for _, elem := range ignore {
		if file == elem {
			return true
		}
	}
	return false
}

func visit(path string, f os.FileInfo, err error) error {
	if fileInIgnore(path, ignore) {
		if f.IsDir() {
			return filepath.SkipDir
		}

		return nil
	}

	fmt.Printf("Visited: %s\n", path)
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
