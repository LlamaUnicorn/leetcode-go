package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// filePattern matches files that start with a four-digit prefix (e.g. "0016_").
var filePattern = regexp.MustCompile(`^(\d{4})_`)

// FileInfo holds the list of file names with the same prefix.
type FileInfo struct {
	Files []string `json:"files"`
}

// DuplicateState stores duplicates per folder and prefix.
type DuplicateState map[string]map[string]*FileInfo

// scanDirectory walks through root and returns the state of duplicate files.
// If 'since' is non-nil, only files newer than the specified date (by
// parsing the file name or using the mod time) are included.
func scanDirectory(root string, since *time.Time) (DuplicateState, error) {
	state := make(DuplicateState)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip __pycache__ directories.
		if info.IsDir() {
			if info.Name() == "__pycache__" {
				return filepath.SkipDir
			}
			return nil
		}

		// If filtering by date, skip files that do not meet the threshold.
		if since != nil {
			if !fileMeetsDateThreshold(info, info.Name(), *since) {
				return nil
			}
		}

		// Check if the file name starts with the four-digit prefix.
		matches := filePattern.FindStringSubmatch(info.Name())
		if matches != nil {
			fileKey := matches[1]
			dir := filepath.Dir(path)

			if _, exists := state[dir]; !exists {
				state[dir] = make(map[string]*FileInfo)
			}
			if group, exists := state[dir][fileKey]; exists {
				group.Files = append(group.Files, info.Name())
			} else {
				state[dir][fileKey] = &FileInfo{
					Files: []string{info.Name()},
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return state, nil
}

// fileMeetsDateThreshold returns true if the file's date is at or after 'since'.
// It first attempts to extract a date from the file name using the pattern "auto_(YYYYMMDD)".
// If that fails, it falls back to the file's modification time.
func fileMeetsDateThreshold(info os.FileInfo, fileName string, since time.Time) bool {
	re := regexp.MustCompile(`auto_(\d{8})`)
	m := re.FindStringSubmatch(fileName)
	if len(m) == 2 {
		t, err := time.Parse("20060102", m[1])
		if err == nil {
			// Return true if t is equal to or after the threshold.
			return !t.Before(since)
		}
	}
	// Fallback to the file modification time.
	modTime := info.ModTime()
	return !modTime.Before(since)
}

// saveState writes the current DuplicateState to the given filename in JSON format.
func saveState(filename string, state DuplicateState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// loadState reads and unmarshals the DuplicateState from the JSON file.
func loadState(filename string) (DuplicateState, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var state DuplicateState
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

// diffStates compares the old and new states and returns two states:
// one for added duplicates and one for removed duplicates.
func diffStates(oldState, newState DuplicateState) (added DuplicateState, removed DuplicateState) {
	added = make(DuplicateState)
	removed = make(DuplicateState)

	// Process newState to detect new or updated duplicate groups.
	for dir, newGroups := range newState {
		for prefix, newGroup := range newGroups {
			// Only consider groups that actually have duplicates.
			if len(newGroup.Files) <= 1 {
				continue
			}
			oldGroups, exists := oldState[dir]
			if !exists {
				// Entire group is new.
				if _, ok := added[dir]; !ok {
					added[dir] = make(map[string]*FileInfo)
				}
				added[dir][prefix] = newGroup
			} else {
				oldGroup, exists2 := oldGroups[prefix]
				if !exists2 {
					// New group for this directory.
					if _, ok := added[dir]; !ok {
						added[dir] = make(map[string]*FileInfo)
					}
					added[dir][prefix] = newGroup
				} else {
					// Compare file lists to determine added and removed files.
					addFiles := difference(newGroup.Files, oldGroup.Files)
					removeFiles := difference(oldGroup.Files, newGroup.Files)
					if len(addFiles) > 0 {
						if _, ok := added[dir]; !ok {
							added[dir] = make(map[string]*FileInfo)
						}
						added[dir][prefix] = &FileInfo{Files: addFiles}
					}
					if len(removeFiles) > 0 {
						if _, ok := removed[dir]; !ok {
							removed[dir] = make(map[string]*FileInfo)
						}
						removed[dir][prefix] = &FileInfo{Files: removeFiles}
					}
				}
			}
		}
	}

	// Process oldState to detect groups that no longer exist.
	for dir, oldGroups := range oldState {
		for prefix, oldGroup := range oldGroups {
			newGroups, exists := newState[dir]
			if !exists || newGroups[prefix] == nil {
				if _, ok := removed[dir]; !ok {
					removed[dir] = make(map[string]*FileInfo)
				}
				removed[dir][prefix] = oldGroup
			}
		}
	}
	return
}

// difference returns the strings in 'a' that are not present in 'b'.
func difference(a, b []string) []string {
	m := make(map[string]bool)
	for _, s := range b {
		m[s] = true
	}
	var diff []string
	for _, s := range a {
		if !m[s] {
			diff = append(diff, s)
		}
	}
	return diff
}

// printState displays the duplicate groups in a user-friendly way.
func printState(state DuplicateState) {
	for dir, groups := range state {
		for prefix, group := range groups {
			if len(group.Files) > 1 {
				fmt.Printf("Folder Path: %s\n", dir)
				fmt.Printf("Duplicate Files with prefix %s:\n", prefix)
				for _, file := range group.Files {
					fmt.Printf("  %s\n", file)
				}
				fmt.Println()
			}
		}
	}
}

// printDiff displays the differences found between states.
func printDiff(added, removed DuplicateState) {
	if len(added) == 0 && len(removed) == 0 {
		fmt.Println("No differences found.")
		return
	}
	if len(added) > 0 {
		fmt.Println("Added Duplicates:")
		printState(added)
	}
	if len(removed) > 0 {
		fmt.Println("Removed Duplicates:")
		printState(removed)
	}
}

func main() {
	// Define command-line flags.
	diffFlag := flag.Bool("diff", false, "Show differences compared to stored state")
	sinceStr := flag.String("since-date", "", "Filter files by date (YYYY-MM-DD)")
	stateFile := flag.String("state-file", "duplicates_state.json", "Path to state file")
	flag.Parse()

	// Parse the since-date flag if provided.
	var since *time.Time = nil
	if *sinceStr != "" {
		t, err := time.Parse("2006-01-02", *sinceStr)
		if err != nil {
			fmt.Println("Invalid date format. Expected YYYY-MM-DD.")
			return
		}
		since = &t
	}

	// Get the current working directory.
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Scan the directory to build the current duplicate state.
	newState, err := scanDirectory(currentDir, since)
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	// Run in one of three modes:
	// 1. --diff: load previous state, compare to current, and display differences.
	// 2. --since-date only: display filtered duplicates (stateless run).
	// 3. Default (no flags): display duplicates and save the current state.
	if *diffFlag {
		oldState, err := loadState(*stateFile)
		if err != nil {
			fmt.Println("Error loading state file:", err)
			return
		}
		added, removed := diffStates(oldState, newState)
		printDiff(added, removed)
	} else if since != nil {
		// Only filtering by date.
		printState(newState)
	} else {
		// Default mode: print duplicates and update state file.
		printState(newState)
		err = saveState(*stateFile, newState)
		if err != nil {
			fmt.Println("Error saving state:", err)
		} else {
			fmt.Println("State saved to", *stateFile)
		}
	}
}
