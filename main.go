package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var baseDir = "."

// Length of the longest file/dir name
var maxLen = 0

// Entry represents a simplified version of fs.DirEntry.
type Entry struct {
	name              string
	daysSinceModified int
}

// add converts an entry of type fs.DirEntry into an entry of type Entry and
// adds it to the provided slice of entries.
func add(entries []Entry, fsEntry fs.DirEntry) []Entry {
	path := filepath.Join(baseDir, fsEntry.Name())
	stat, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	modTime := stat.ModTime()

	if fsEntry.IsDir() {
		dirEntries, err := os.ReadDir(path)
		if err == nil {
			for _, e := range dirEntries {
				if e.IsDir() && e.Name() == ".git" {
					gitPath := filepath.Join(baseDir, fsEntry.Name(), ".git")
					gitStat, err := os.Stat(gitPath)
					if err == nil {
						modTime = gitStat.ModTime()
					}
				}
			}
		}
	}

	return append(entries, Entry{
		name:              fsEntry.Name(),
		daysSinceModified: int(time.Since(modTime).Hours() / 24),
	})
}

func main() {
	fFlag := flag.Bool("f", false, "List files only")
	dFlag := flag.Bool("d", false, "List directories only")
	mFlag := flag.Bool("m", false, "Show last time modified")
	sFlag := flag.String("s", "name", "Sort method")
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
	var dirs []Entry
	var files []Entry

	// Populate the slices depending on the flag
	switch {
	case *dFlag && *fFlag:
		fmt.Fprintln(os.Stderr, "l: flags -f and -d cannot be used together")
		os.Exit(1)
	case *fFlag:
		for _, item := range list {
			if !item.IsDir() {
				files = add(files, item)
			}
		}
	case *dFlag:
		for _, item := range list {
			if item.IsDir() {
				dirs = add(dirs, item)
			}
		}
	default:
		for _, item := range list {
			if item.IsDir() {
				dirs = add(dirs, item)
			} else {
				files = add(files, item)
			}
		}
	}

	// Determine the length of the longest name
	// TODO: truncate vert long names
	if *mFlag {
		for _, dir := range files {
			l := len(dir.name)
			if l > maxLen {
				maxLen = l
			}
		}
		for _, dir := range dirs {
			l := len(dir.name)
			if l > maxLen {
				maxLen = l
			}
		}
	}

	// Sort
	switch *sFlag {
	case "name":
		sort.Slice(dirs, func(i, j int) bool { return dirs[i].name < dirs[j].name })
		sort.Slice(files, func(i, j int) bool { return files[i].name < files[j].name })
	case "time":
		sort.Slice(dirs, func(i, j int) bool { return dirs[i].daysSinceModified < dirs[j].daysSinceModified })
		sort.Slice(files, func(i, j int) bool { return files[i].daysSinceModified < files[j].daysSinceModified })
	default:
		fmt.Fprintln(os.Stderr, "l: unknown value for flag -s, possible values are \"name\" (default) and \"time\"")
		os.Exit(1)
	}

	// Print dirs
	for _, dir := range dirs {
		printDir(dir, *mFlag)
	}

	// Print files
	for _, file := range files {
		printFile(file, *mFlag)
	}
}
