package folder

import "github.com/gofrs/uuid"

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	// GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder,error)

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}



type Driver struct {  // Capitalized to make it accessible from other packages
    Folders []Folder
}
func NewDriver(folders []Folder) IDriver {
	return &Driver{
		// initialize attributes here
		Folders: folders,
	}
}
