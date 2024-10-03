package folder

import (
    // "encoding/json"
    // "io/ioutil"
    // "log"
    // "os"
    "errors"
    "strings"
    "github.com/gofrs/uuid"
   
    
)




// Get all child folders of a given parent folder
func (d *Driver) GetAllChildFolders(orgID uuid.UUID, parentName string) ([]Folder,error) {
    childFolders := []Folder{}
    var IDfound bool = false
    var parentExists bool = false

    // Iterate through all folders to find children
    for _, folder := range d.Folders {
        // Check if folder belongs to the specified organization
        if folder.OrgId == orgID {
            IDfound = true
            // Check if folder is the parent folder
            if folder.Name == parentName {
                parentExists = true
            }    
            // Check if folder's path starts with parentName and is a direct child
            if folder.Paths != parentName && isChildFolder(parentName, folder.Paths) {
                childFolders = append(childFolders, folder)
            }
        }
    }

        
    // Check if the organization ID was found
    if !IDfound {
       return nil, errors.New("organization ID not found")

    }
    // Check parent folder existence
    if !parentExists {
        return nil, errors.New("parent folder not found")
    }
    // Return the list of child folders
    return childFolders,nil
}

// Helper function to determine if a folder is a child of the given parent folder
func isChildFolder(parentName, folderPath string) bool {
    // Use strings.HasPrefix to check if folderPath starts with "parentName."
    return strings.Contains(folderPath, parentName+".")
}

