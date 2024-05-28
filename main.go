/*
 * @Author: Vincent Yang
 * @Date: 2024-04-27 09:35:37
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2024-05-28 11:52:06
 * @FilePath: /UpdateCompose/main.go
 * @Telegram: https://t.me/missuo
 * @GitHub: https://github.com/missuo
 *
 * Copyright © 2024 by Vincent, All Rights Reserved.
 */

package main

import (
	"flag"
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

	// Parse command line arguments for include and exclude
	includeDirs := flag.String("include", "", "Comma-separated list of directories to include")
	excludeDirs := flag.String("exclude", "", "Comma-separated list of directories to exclude")
	flag.Parse()

	includeList := strings.Split(*includeDirs, ",")
	excludeList := strings.Split(*excludeDirs, ",")

	// Walk the root directory to find Compose files
	err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		// Skip the /proc directory
		if info.IsDir() && info.Name() == "proc" {
			return filepath.SkipDir
		}

		// Skip excluded directories
		for _, exclude := range excludeList {
			if strings.HasPrefix(path, exclude) {
				return filepath.SkipDir
			}
		}

		// Check if the file name matches any of the Compose file names
		for _, file := range composeFiles {
			if info.Name() == file {
				dir := filepath.Dir(path)

				// Check if the directory is in the include list, if specified
				if len(includeList) > 0 {
					included := false
					for _, include := range includeList {
						if strings.HasPrefix(dir, include) {
							included = true
							break
						}
					}
					if !included {
						return nil
					}
				}

				// Found a Compose file, perform update operations
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
		return fmt.Errorf("error changing directory to %s: %v", dir, err)
	}

	// Use the absolute path of the docker compose command
	composeCmd := "docker"
	composeCmdArgs := []string{"compose"}

	// Execute Docker Compose commands in sequence
	for _, cmd := range []string{"stop", "pull", "up -d"} {
		fmt.Printf("Running command: %s %s %s\n", composeCmd, strings.Join(composeCmdArgs, " "), cmd)
		out, err := exec.Command(composeCmd, append(composeCmdArgs, strings.Split(cmd, " ")...)...).CombinedOutput()
		if err != nil {
			return fmt.Errorf("error running '%s %s %s': %v, output: %s", composeCmd, strings.Join(composeCmdArgs, " "), cmd, err, string(out))
		}
	}

	return nil
}
