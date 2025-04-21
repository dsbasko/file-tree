# FileTree

A Go library for recursively traversing a project's file structure and displaying it as a
tree in the console.

## Features

- Recursive traversal of directories and files
- Visualization of the structure as an ASCII tree
- Optional display of hidden files (starting with a dot)
- Configurable maximum traversal depth
- Sorting of elements (directories first, then files)

## Example Output

```
project
├─ cmd
│　└─ main.go
├─ filetree
│　└─ filetree.go
├─ internal
│　├─ controller
│　└─ service
│　　　└─ some_service.go
└─ .gitignore
```

## Installation

```bash
go get github.com/dsbasko/file-tree
```

## Usage

### Simple Example

```go
package main

import (
	"fmt"
	"github.com/dsbasko/file-tree"
)

func main() {
	err := filetree.PrintTree(".")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

### Example with Options

```go
package main

import (
	"fmt"
	"github.com/dsbasko/file-tree"
)

func main() {
	options := filetree.Options{
		ShowHidden: true,   // show hidden files
		MaxDepth:   2,      // limit depth to 2 levels
	}

	err := filetree.PrintTreeWithOptions("/path/to/project", options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

## Using the CLI

The project includes a simple CLI utility for quick use:

```bash
# Install the CLI
go install github.com/dsbasko/file-tree

# Usage
filetree -path=/path/to/project -hidden=true -depth=3
```

### Available Flags

| Flag      | Description                                | Default Value           |
| --------- | ------------------------------------------ | ----------------------- |
| `-path`   | Path to the directory for analysis         | `.` (current directory) |
| `-hidden` | Show hidden files                          | `false`                 |
| `-depth`  | Maximum traversal depth (-1 for unlimited) | `-1`                    |

## API

### Types

```go
// Options contains settings for printing the file tree
type Options struct {
	ShowHidden bool // whether to show hidden files and directories
	MaxDepth   int  // maximum traversal depth (-1 for unlimited)
}
```

### Functions

```go
// PrintTree outputs the directory structure to the console with default settings
func PrintTree(rootPath string) error

// PrintTreeWithOptions outputs the directory structure with specified settings
func PrintTreeWithOptions(rootPath string, options Options) error

// DefaultOptions returns the default settings
func DefaultOptions() Options
```
