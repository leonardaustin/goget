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
	destPath := "src" // Change this to your desired destination path

	err := cloneAndStore(repoURL, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository cloned and stored successfully!")
}

func cloneAndStore(repoURL, destPath string) error {
	// Create the destination directory if it doesn't exist
	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Retrieve the repository name from the URL
	repoName := filepath.Base(repoURL)

	// Set the command to execute
	cmd := exec.Command("git", "clone", repoURL, filepath.Join(destPath, repoName))

	// Set the command's working directory
	cmd.Dir = filepath.Dir(destPath)

	// Execute the command
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
