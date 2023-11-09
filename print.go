package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// printDir prints a directory's name along with its icon, color.
func printDir(entry Entry, mod bool) {
	icon := dirIcon
	path := filepath.Join(baseDir, entry.name)
	if isEmpty(path) {
		icon = emptyDirIcon
	}
	modTime := formatModTime(entry, mod)
	name := entry.name
	if entry.isSymlink {
		name = colorize(name, cyan)
		icon = symlinkIcon
	}
	fmt.Printf("\033[34m%s %s\033[0m%s\n", icon, name, modTime)
}

// printFile prints a filename along with its icon and color.
func printFile(entry Entry, mod bool) {
	ext := filepath.Ext(entry.name)
	icon := defaultFileIcon

	if data, ok := specialFileIcons[entry.name]; ok {
		icon = colorize(data.icon, data.color)
	} else if data, ok := filetypeIcons[ext]; ok {
		icon = colorize(data.icon, data.color)
	}
	modTime := formatModTime(entry, mod)
	name := entry.name
	if entry.isSymlink {
		name = colorize(name, cyan)
	}
	fmt.Printf("%s %s%s\n", icon, name, modTime)
}

// colorize colorizes a string with the given colorize.
func colorize(s string, c int8) string {
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

// formatModTime formats modified time by adding appropriate padding and
// coloring.
func formatModTime(entry Entry, mod bool) string {
	if !mod {
		return ""
	}

	modColor := none
	days := entry.daysSinceModified

	switch {
	case days > 120:
		modColor = red
	case days > 30:
		modColor = yellow
	case days <= 7:
		modColor = green
	}

	pad := strings.Repeat(" ", maxLen-len(entry.name)+2)
	modTime := colorize(timeSince(days), modColor)

	return fmt.Sprintf("%s%s", pad, modTime)
}

// timeSince returns the time elapsed in the given number of days in a
// human-friendly format.
func timeSince(days int) string {
	switch {
	case days >= 365:
		years := days / 365
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%v years ago", years)
	case days >= 30:
		months := days / 30
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%v months ago", months)
	case days >= 7:
		weeks := days / 7
		if weeks == 1 {
			return "1 week ago"
		}
		return fmt.Sprintf("%v weeks ago", weeks)
	case days > 0:
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%v days ago", days)
	}
	return "Today"
}
