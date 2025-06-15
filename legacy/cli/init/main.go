package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var folders = []string{
	"linkedlist",
	"stack",
	"queue",
	"tree",
	"graph",
	"sort",
	"dp",
	"hash",
	"common",
}

func initProject() {
	for _, folder := range folders {
		path := folder
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatalf("Failed to create folder %s: %v", path, err)
		}

		// Create go file
		goFile := fmt.Sprintf("%s/%s.go", folder, folder)
		if err := os.WriteFile(goFile, []byte(fmt.Sprintf("package %s\n\n", folder)), 0644); err != nil {
			log.Fatalf("Failed to create %s: %v", goFile, err)
		}

		// Create test file
		testFile := fmt.Sprintf("%s/%s_test.go", folder, folder)
		if err := os.WriteFile(testFile, []byte(fmt.Sprintf("package %s\n\nimport \"testing\"\n\nfunc TestExample(t *testing.T) {}\n", folder)), 0644); err != nil {
			log.Fatalf("Failed to create %s: %v", testFile, err)
		}
	}

	// Create README
	readme := `# Go Data Structures & Algorithms

This is a starter scaffold for implementing all common data structures and algorithms in Go.`
	if err := os.WriteFile("README.md", []byte(readme), 0644); err != nil {
		log.Fatalf("Failed to create README.md: %v", err)
	}

	fmt.Println("✅ Project initialized.")
}

func addModule(moduleName string) {
	if moduleName == "" {
		log.Fatal("Module name is required")
	}

	if err := os.MkdirAll(moduleName, os.ModePerm); err != nil {
		log.Fatalf("Failed to create folder %s: %v", moduleName, err)
	}

	goFile := fmt.Sprintf("%s/%s.go", moduleName, moduleName)
	testFile := fmt.Sprintf("%s/%s_test.go", moduleName, moduleName)

	os.WriteFile(goFile, []byte(fmt.Sprintf("package %s\n\n", moduleName)), 0644)
	os.WriteFile(testFile, []byte(fmt.Sprintf("package %s\n\nimport \"testing\"\n\nfunc TestExample(t *testing.T) {}\n", moduleName)), 0644)

	fmt.Printf("✅ Module '%s' added.\n", moduleName)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-dsa init | add <module>")
		return
	}

	switch os.Args[1] {
	case "init":
		initProject()
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a module name, e.g. go-dsa add trie")
		}
		addModule(strings.ToLower(os.Args[2]))
	default:
		fmt.Println("Unknown command")
	}
}
