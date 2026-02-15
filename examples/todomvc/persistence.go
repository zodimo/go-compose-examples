package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// persistedState is the structure saved to disk.
// We don't persist EditingID since that's ephemeral UI state.
type persistedState struct {
	Todos  []Todo `json:"todos"`
	Filter Filter `json:"filter"`
	NextID int    `json:"nextId"`
}

// getDataFilePath returns the path to the data file.
func getDataFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}
	return filepath.Join(configDir, "todomvc-gocompose.json")
}

// SaveToFile persists the todo state to a JSON file.
func (s *TodoState) SaveToFile() error {
	data := persistedState{
		Todos:  s.Todos,
		Filter: s.Filter,
		NextID: s.NextID,
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(getDataFilePath(), bytes, 0644)
}

// LoadFromFile loads the todo state from a JSON file.
// Returns a new state if the file doesn't exist or can't be read.
func LoadFromFile() *TodoState {
	path := getDataFilePath()
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading state: %v\n", err)
		return NewTodoState()
	}

	var data persistedState
	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Printf("Error unmarshalling state: %v\n", err)
		return NewTodoState()
	}

	return &TodoState{
		Todos:     data.Todos,
		Filter:    data.Filter,
		EditingID: -1,
		NextID:    data.NextID,
	}
}
