package main

import (
	"fmt"
	"os"
	"sort"
)

// Default baseDir directory
var baseDir = "."

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		baseDir = args[0]
	}

	list, err := os.ReadDir(baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Slices of dirs and files
	var dirs []os.DirEntry
	var files []os.DirEntry

	// Populate the slices
	for _, item := range list {
		if item.IsDir() {
			dirs = append(dirs, item)
		} else {
			files = append(files, item)
		}
	}

	// Sort by name
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	// Print dirs
	for _, dir := range dirs {
		name := dir.Name()
		printDir(name)
	}

	// Print files
	for _, file := range files {
		name := file.Name()
		printFile(name)
	}
}
