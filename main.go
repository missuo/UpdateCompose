package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
				err := updateCompose(filepath.Dir(path))
				if err != nil {
					fmt.Printf("Error updating compose at %s: %v\n", path, err)
				} else {
					fmt.Printf("Updated compose at %s\n", path)
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

	// Execute docker compose commands in sequence
	for _, cmd := range []string{"stop", "pull", "up -d"} {
		out, err := exec.Command("docker", "compose", cmd).CombinedOutput()
		if err != nil {
			return fmt.Errorf("error running 'docker compose %s': %v, output: %s", cmd, err, out)
		}
	}

	return nil
}
