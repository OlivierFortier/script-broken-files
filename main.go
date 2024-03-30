package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type FileData struct {
	Path             string `json:"path"`
	ModifiedDate     int64  `json:"modified_date"`
	Size             int64  `json:"size"`
	CurrentExtension string `json:"current_extension"`
	ProperExtensions string `json:"proper_extensions"`
}

func main() {
	// Read the JSON file
	jsonFile, err := os.ReadFile("data.json")
	if err != nil {
		log.Print("Failed to read data.json")
		log.Fatal(err)
	}

	// Unmarshal the JSON data
	var files []FileData
	err = json.Unmarshal(jsonFile, &files)
	if err != nil {
		log.Print("Failed to unmarshal data.json")
		log.Fatal(err)
	}

	// Create a slice to store the files that were not renamed
	var notRenamedFiles []FileData

	// Rename the files
	for _, file := range files {
		log.Print("===================================== \n")
		log.Printf("Processing file: %s\n", file.Path)
		// Extract the extension from the parentheses
		extension := strings.Trim(strings.Split(file.ProperExtensions, ")")[0], "(")
		log.Printf("Extension: %s\n", extension)
		newPath := file.Path + "." + extension
		log.Printf("New path: %s\n", newPath)
		err = os.Rename(file.Path, newPath)
		if err != nil {
			// Add the file to the notRenamedFiles slice
			notRenamedFiles = append(notRenamedFiles, file)
			log.Printf("Failed to rename file: %s\n", file.Path)
		} else {
			log.Printf("Renamed file: %s\n", file.Path)
		}
	}

	// Write the notRenamedFiles slice to not_renamed.json
	notRenamedJSON, err := json.Marshal(notRenamedFiles)
	if err != nil {
		log.Print("Failed to marshal notRenamedFiles")
		log.Fatal(err)
	}
	err = os.WriteFile("not_renamed.json", notRenamedJSON, 0644)
	if err != nil {
		log.Print("Failed to write not_renamed.json")
		log.Fatal(err)
	}

	log.Println("Process completed successfully.")
}
