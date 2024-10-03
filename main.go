package main

import (
    // "encoding/json"
    "fmt"
    // "io/ioutil"
    "log"
    // "os"

	"sc-take-home-assessment-take-home-2025/folder" // Import the folder package

    "github.com/gofrs/uuid"
)




func main() {
    // Load folder data using GetSampleData()
    folders := folder.GetSampleData("sample2.json")

    // Initialize a driver with folder data
    driver := folder.Driver{Folders: folders}

    // A sample organization ID and parent folder name (sample2 file i made for testing) easier to understand
    orgID := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
    parentName := "file1"

    // Component 1: GetAllChildFolders
    childFolders,err := driver.GetAllChildFolders(orgID, parentName)
    if err != nil {
        log.Fatalf("Failed to get child folders: %v", err)
    }

    // Print results for GetAllChildFolders
    fmt.Printf("Child folders of %s:\n", parentName)
    for _, folder := range childFolders {
        fmt.Printf("Folder: %s, Path: %s\n", folder.Name, folder.Paths)
    }
    
    fmt.Println("Component 2: MoveFolder")
    
    // Component 2: MoveFolder
	movedFolders, err := driver.MoveFolder("file1", "file5")
	if err != nil {
		log.Fatalf("Failed to move folder: %v", err)
	}
    // Print results for MoveFolder
	fmt.Println("Folders after moving:")
	for _, folder := range movedFolders {
		fmt.Printf("Folder: %s, Path: %s\n", folder.Name, folder.Paths)
	}

}

