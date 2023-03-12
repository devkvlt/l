package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	args := os.Args[1:]

	// Default arg to cwd
	base := "."
	if len(args) > 0 {
		base = args[0]
	}

	list, err := os.ReadDir(base)
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
		icon := dirIcon
		fullPath := filepath.Join(base, name)
		if isEmpty(fullPath) {
			icon = emptyDirIcon
		}
		fmt.Printf("\033[34m%s %s\033[0m\n", icon, name)
	}

	// Print files
	for _, file := range files {
		name := file.Name()
		ext := filepath.Ext(name)
		icon, ok := icons[ext]
		if !ok {
			icon = fileIcon
		}
		fmt.Printf("%s %s\n", icon, name)
	}
}

// isEmpty reports whether a directory is empty.
func isEmpty(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == nil {
		return false
	} else {
		return true
	}
}
