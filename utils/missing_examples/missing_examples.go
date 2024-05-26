// UTIL for finding missing examples in the project
// USAGE:
// 	go run utils/missing_examples/missing_examples.go shapes

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const RAYLIB_EXAMPLE_PATH = "raylib/examples"
const GO_EXAMPLE_PATH = "src/examples"

const RAYLIB_EXAMPLES_EXT = ".c"
const GO_EXAMPLES_EXT = ".go"

type Example struct {
	Name string
	Ext  string
	Path string
}

type exampleDomain int

const (
	GO_RAYLIB exampleDomain = iota
	C_RAYLIB
)

func (e Example) domain() exampleDomain {
	// Assume all C Raylib Domains have .c Extension
	switch e.Ext {
	case RAYLIB_EXAMPLES_EXT:
		return C_RAYLIB
	case GO_EXAMPLES_EXT:
		return GO_RAYLIB
	}
	panic("Unreachable")
}

func (e Example) to_create() Example {
	var new_example Example

	switch e.domain() {
	case C_RAYLIB:
		path, found := strings.CutPrefix(e.Path, RAYLIB_EXAMPLE_PATH)
		if !found {
			fmt.Printf("ERROR: Raylib path not in example %s\n", e)
			panic("Path not in example")
		}
		module, found := strings.CutSuffix(path, e.Name+e.Ext)
		if !found {
			fmt.Printf("ERROR: Raylib file not in example %s\n", e)
			panic("File not in Example")
		}
		return Example{
			Name: e.Name,
			Ext:  ".go",
			Path: GO_EXAMPLE_PATH + module + e.Name + "/" + e.Name + ".go",
		}
	}
	return new_example

}

func main() {
	var program string
	var args []string
	var module string

	args = os.Args
	program, args = args[0], args[1:]
	if len(args) == 0 {
		fmt.Printf("USAGE: %s <module>\n", program)
		os.Exit(1)
	}

	module, _ = args[0], args[1:]

	raylib_example_path := fmt.Sprintf("%s/%s", RAYLIB_EXAMPLE_PATH, module)
	go_example_path := fmt.Sprintf("%s/%s", GO_EXAMPLE_PATH, module)

	raylib_files, err := get_raylib_example_files(raylib_example_path)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	go_files, err := get_go_example_files(go_example_path)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Missing Examples for module: %s\n", module)
	fmt.Println("In Raylib but not in Go:")
	not_in_go := difference(raylib_files, go_files)
	for _, file := range not_in_go {
		fmt.Printf("\t%s\n", file.Path)
		to_create := file.to_create()
		fmt.Printf("\t\tCreate: %s\n", to_create.Path)
	}

	fmt.Println("In Go but not in Raylib:")
	not_in_raylib := difference(go_files, raylib_files)
	for _, file := range not_in_raylib {
		fmt.Printf("\t%s\n", file.Path)
	}

}

func difference(a, b []Example) []Example {
	m := make(map[string]bool)
	var diff []Example

	for _, item := range b {
		m[item.Name] = true
	}

	for _, item := range a {
		if _, ok := m[item.Name]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}

func get_go_example_files(path string) ([]Example, error) {
	go_files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []Example
	for _, file := range go_files {
		if file.IsDir() {
			sub_files, err := get_go_example_files(
				fmt.Sprintf("%s/%s", path, file.Name()))
			if err != nil {
				return nil, err
			}
			files = append(files, sub_files...)
			continue
		}
		if filepath.Ext(file.Name()) != GO_EXAMPLES_EXT {
			continue
		}

		files = append(files, Example{
			Name: file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))],
			Ext:  filepath.Ext(file.Name()),
			Path: fmt.Sprintf("%s/%s", path, file.Name()),
		})
	}
	return files, nil
}

func get_raylib_example_files(path string) ([]Example, error) {
	raylib_files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []Example
	for _, file := range raylib_files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != RAYLIB_EXAMPLES_EXT {
			continue
		}

		files = append(files, Example{
			Name: file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))],
			Ext:  filepath.Ext(file.Name()),
			Path: fmt.Sprintf("%s/%s", path, file.Name()),
		})
	}
	return files, nil
}
