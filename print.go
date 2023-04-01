package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// printDir prints a directory's name along with its icon, color.
func printDir(name string, mod bool) {
	icon := dirIcon
	path := filepath.Join(baseDir, name)
	if isEmpty(path) {
		icon = emptyDirIcon
	}
	modTime := formatModTime(name, mod)
	fmt.Printf("\033[34m%s %s\033[0m%s\n", icon, name, modTime)
}

// printFile print a filename along with its icon and color.
func printFile(name string, mod bool) {
	ext := filepath.Ext(name)
	icon := defaultFileIcon

	if data, ok := specialFileIcons[name]; ok {
		icon = color(data.icon, data.color)
	} else if data, ok := filetypeIcons[ext]; ok {
		icon = color(data.icon, data.color)
	}
	modified := formatModTime(name, mod)
	fmt.Printf("%s %s%s\n", icon, name, modified)
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

func formatModTime(name string, mod bool) string {
	if !mod {
		return ""
	}

	info, err := os.Stat(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}
	return fmt.Sprintf(" (%s)", timeSince(info.ModTime()))
}

// timeSince returns the time elapsed since the given date in a human-friendly
// format.
func timeSince(t time.Time) string {
	diff := time.Since(t)
	days := int(diff.Hours() / 24)

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
