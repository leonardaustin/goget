package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: goget <repo-url>")
	}

	repoURL := os.Args[1]
	destPath := getDestPath(repoURL)

	err := gitClone(repoURL, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository cloned and stored successfully!")
}

func getDestPath(repoURL string) string {
	// Remove the protocol from the URL
	repoURL = strings.TrimPrefix(repoURL, "https://")
	repoURL = strings.TrimPrefix(repoURL, "http://")
	repoURL = strings.TrimPrefix(repoURL, "git://")

	// Split the remaining URL into parts
	parts := strings.Split(repoURL, "/")

	// Ensure the URL has at least 3 parts (e.g., github.com/orgname/repo)
	if len(parts) < 3 {
		log.Fatal("Invalid repository URL")
	}

	// Extract the domain and organization parts
	domain := parts[0]
	organization := parts[1]

	// Create the destination path
	destPath := filepath.Join(os.Getenv("HOME"), "src", domain, organization)

	return destPath
}

func gitClone(repoURL, destPath string) error {
	// Create the parent directory if it doesn't exist
	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Set the command to execute
	cmd := exec.Command("git", "clone", "https://"+repoURL, destPath)

	// Execute the command
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
