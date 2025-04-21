package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	path := flag.String("path", ".", "Путь к директории для отображения структуры")
	showHidden := flag.Bool("hidden", false, "Показывать скрытые файлы и директории")
	maxDepth := flag.Int("depth", -1, "Максимальная глубина обхода (-1 для неограниченной)")
	flag.Parse()

	options := Options{
		ShowHidden: *showHidden,
		MaxDepth:   *maxDepth,
	}

	err := PrintTreeWithOptions(*path, options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}

type Options struct {
	ShowHidden bool
	MaxDepth   int
}

func DefaultOptions() Options {
	return Options{
		ShowHidden: false,
		MaxDepth:   -1,
	}
}

func PrintTree(rootPath string) error {
	return PrintTreeWithOptions(rootPath, DefaultOptions())
}

func PrintTreeWithOptions(rootPath string, options Options) error {
	info, err := os.Stat(rootPath)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s не является директорией", rootPath)
	}

	fmt.Println(filepath.Base(rootPath))
	err = printDir(rootPath, "", 0, options)
	if err != nil {
		return err
	}

	return nil
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func printDir(path string, prefix string, depth int, options Options) error {
	if options.MaxDepth >= 0 && depth > options.MaxDepth {
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		name := entry.Name()
		if !options.ShowHidden && isHidden(name) {
			continue
		}
		filteredEntries = append(filteredEntries, entry)
	}
	entries = filteredEntries

	sort.Slice(entries, func(i, j int) bool {
		iIsDir := entries[i].IsDir()
		jIsDir := entries[j].IsDir()

		if iIsDir != jIsDir {
			return iIsDir
		}
		return entries[i].Name() < entries[j].Name()
	})

	for i, entry := range entries {
		isLast := i == len(entries)-1

		var currentPrefix string
		if isLast {
			currentPrefix = prefix + "└─ "
		} else {
			currentPrefix = prefix + "├─ "
		}

		fmt.Printf("%s%s\n", currentPrefix, entry.Name())

		if entry.IsDir() {
			var newPrefix string
			if isLast {
				newPrefix = prefix + "　　"
			} else {
				newPrefix = prefix + "│　"
			}

			err = printDir(filepath.Join(path, entry.Name()), newPrefix, depth+1, options)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
