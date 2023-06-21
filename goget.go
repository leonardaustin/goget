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
	destPath := filepath.Join(os.Getenv("HOME"), "src", repoURL)

	err := gitClone(repoURL, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository cloned and stored successfully!")
}

func gitClone(repoURL, destPath string) error {
	// Create the destination directory if it doesn't exist
	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Retrieve the repository name from the URL
	repoName := filepath.Base(repoURL)

	// Set the command to execute
	cmd := exec.Command("git", "clone", "https://"+repoURL, filepath.Join(destPath, repoName))

	// Execute the command
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
