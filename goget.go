package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: goget <repo-url>")
	}

	repoURL := os.Args[1]
	destPath := filepath.Join(os.Getenv("HOME"), "src", filepath.Base(repoURL))

	err := gitClone(repoURL, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository cloned and stored successfully!")
}

func gitClone(repoURL, destPath string) error {
	// Create the parent directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
	if err != nil {
		return err
	}

	// Set the command to execute
	cmd := exec.Command("git", "clone", repoURL, destPath)

	// Execute the command
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
