package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Check if the current user is root
	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root.")
		os.Exit(1)
	}

	// Define the Compose file names to search for
	composeFiles := []string{"compose.yaml", "docker-compose.yaml", "docker-compose.yml", "compose.yml"}

	// Walk the root directory to find Compose files
	err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		// Skip the /proc directory
		if info.IsDir() && info.Name() == "proc" {
			return filepath.SkipDir
		}

		// Check if the file name matches any of the Compose file names
		for _, file := range composeFiles {
			if info.Name() == file {
				// Found a Compose file, perform update operations
				dir := filepath.Dir(path)
				fmt.Println("Found Compose file in directory", dir)
				err := updateCompose(dir)
				if err != nil {
					fmt.Printf("Error updating compose in %s: %v\n", dir, err)
				} else {
					fmt.Printf("Updated compose in %s\n", dir)
				}
				break
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
	}
}

func updateCompose(dir string) error {
	// Change to the directory containing the Compose file
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	// Check the Docker version
	out, err := exec.Command("docker", "--version").Output()
	if err != nil {
		return fmt.Errorf("error checking Docker version: %v", err)
	}

	// Determine the Compose command based on the Docker version
	composeCmd := "/usr/bin/docker compose"
	if strings.Contains(string(out), "version 1.") {
		composeCmd = "docker-compose"
	}

	// Execute Docker Compose commands in sequence
	for _, cmd := range []string{"stop", "pull", "up -d"} {
		out, err := exec.Command(composeCmd, strings.Split(cmd, " ")...).CombinedOutput()
		if err != nil {
			return fmt.Errorf("error running '%s %s': %v, output: %s", composeCmd, cmd, err, out)
		}
	}

	return nil
}
