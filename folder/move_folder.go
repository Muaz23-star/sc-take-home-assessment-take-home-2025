package folder

import (
	"errors"
	//"fmt"
	"strings"


)

// MoveFolder moves a folder and its subtree to a new destination folder
func (d *Driver) MoveFolder(name string, dst string) ([]Folder, error) {
	var sourceFolder *Folder
	var destinationFolder *Folder
	var sourceIndex int

	// 1. Locate the source folder
	for i, folder := range d.Folders {
		if folder.Name == name {
			sourceFolder = &folder
			sourceIndex = i
			break
		}
	}
	if sourceFolder == nil {
		return nil, errors.New("source folder does not exist")
	}

	// 2. Locate the destination folder
	for _, folder := range d.Folders {
		if folder.Name == dst {
			destinationFolder = &folder
			break
		}
	}
	if destinationFolder == nil {
		return nil, errors.New("destination folder does not exist")
	}

	// 3. Error handling
	if sourceFolder.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}
	if sourceFolder.Paths == destinationFolder.Paths {
		return nil, errors.New("cannot move a folder to itself")
	}
	if strings.HasPrefix(destinationFolder.Paths, sourceFolder.Paths) {
		return nil, errors.New("cannot move a folder to a child of itself")
	}

	// 4. Move the subtree
	// Get the new path prefix for the moved folder
	oldPathPrefix := sourceFolder.Paths
	newPathPrefix := destinationFolder.Paths + "." + name
	

	

	// Update the source folder's path
	d.Folders[sourceIndex].Paths = newPathPrefix

	// Update the paths of all child folders in the subtree
	for i, folder := range d.Folders {
		// Check if the folder is a child of the source folder (by checking the path prefix)
		if strings.Contains(folder.Paths, oldPathPrefix+".") {
			// Update the folder's path by replacing the old path prefix with the new one
			newPath := strings.Replace(folder.Paths, oldPathPrefix, newPathPrefix, 1)
			d.Folders[i].Paths = newPath
		}
	}

	return d.Folders, nil
}
