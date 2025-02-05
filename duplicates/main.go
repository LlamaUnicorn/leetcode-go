package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var filePattern = regexp.MustCompile(`^(\d{4})_`)

type FileInfo struct {
	Files []string
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	filesMap := make(map[string]map[string]*FileInfo)

	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		matches := filePattern.FindStringSubmatch(info.Name())
		if matches != nil {
			fileKey := matches[1]
			dir := filepath.Dir(path)

			if _, exists := filesMap[dir]; !exists {
				filesMap[dir] = make(map[string]*FileInfo)
			}

			if fileInfo, exists := filesMap[dir][fileKey]; exists {
				fileInfo.Files = append(fileInfo.Files, info.Name())
			} else {
				filesMap[dir][fileKey] = &FileInfo{
					Files: []string{info.Name()},
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
		return
	}

	// Print the results
	for dir, filesByPrefix := range filesMap {
		for key, fileInfo := range filesByPrefix {
			if len(fileInfo.Files) > 1 {
				fmt.Printf("Folder Path: %s\n", dir)
				fmt.Printf("Duplicate Files with prefix %s:\n", key)
				for _, file := range fileInfo.Files {
					fmt.Printf("  %s\n", file)
				}
				fmt.Println()
			}
		}
	}
}
