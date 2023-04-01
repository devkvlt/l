package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var baseDir = "."

func main() {
	fFlag := flag.Bool("f", false, "List files only")
	dFlag := flag.Bool("d", false, "List directories only")
	mFlag := flag.Bool("m", false, "Show last time modified")
	flag.Parse()

	argc := flag.NArg()
	if argc == 1 {
		baseDir = flag.Arg(0)
	}
	if argc > 1 {
		fmt.Fprintln(os.Stderr, "l: too many arguments")
		os.Exit(1)
	}

	list, err := os.ReadDir(baseDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Slices of dirs and files
	var dirs []os.DirEntry
	var files []os.DirEntry

	// Populate the slices depending on the flag
	switch {
	case *dFlag && *fFlag:
		fmt.Fprintln(os.Stderr, "l: flags -f and -d can't be both set")
		os.Exit(1)
	case !*dFlag && *fFlag:
		for _, item := range list {
			if !item.IsDir() {
				files = append(files, item)
			}
		}
	case *dFlag && !*fFlag:
		for _, item := range list {
			if item.IsDir() {
				dirs = append(dirs, item)
			}
		}
	default:
		for _, item := range list {
			if item.IsDir() {
				dirs = append(dirs, item)
			} else {
				files = append(files, item)
			}
		}
	}

	// Sort by name
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	// Print dirs
	for _, dir := range dirs {
		name := dir.Name()
		printDir(name, *mFlag)
	}

	// Print files
	for _, file := range files {
		name := file.Name()
		printFile(name, *mFlag)
	}
}
