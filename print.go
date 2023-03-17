package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// printDir prints a directory's name along with its icon and color.
func printDir(name string) {
	icon := dirIcon
	path := filepath.Join(baseDir, name)
	if isEmpty(path) {
		icon = emptyDirIcon
	}
	fmt.Printf("\033[34m%s %s\033[0m\n", icon, name)
}

// printFile print a filename along with its icon and color.
func printFile(name string) {
	ext := filepath.Ext(name)
	icon := defaultFileIcon

	if data, ok := specialFileIcons[name]; ok {
		icon = color(data.icon, data.color)
	} else if data, ok := filetypeIcons[ext]; ok {
		icon = color(data.icon, data.color)
	}

	fmt.Printf("%s %s\n", icon, name)
}

// color colors a string with the given color.
func color(s string, c int8) string {
	switch c {
	case red:
		return "\033[31m" + s + "\033[0m"
	case green:
		return "\033[32m" + s + "\033[0m"
	case yellow:
		return "\033[33m" + s + "\033[0m"
	case blue:
		return "\033[34m" + s + "\033[0m"
	case magenta:
		return "\033[35m" + s + "\033[0m"
	case cyan:
		return "\033[36m" + s + "\033[0m"
	}
	return s
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
	return err != nil
}
