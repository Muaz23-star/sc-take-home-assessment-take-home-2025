package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"reflect"
)

// TestMoveFolder tests the MoveFolder function.
func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	// Load sample folder data from a JSON file
	folders := folder.GetSampleData("sample2.json")

	


	// Initialize the Driver with the folders
	driver := folder.NewDriver(folders)

	// Define test cases
	tests := []struct {
		name        string
		source      string
		destination string
		orgID       uuid.UUID
		wantError   string
	}{
		{
			name:        "Test case 1: Successfully move folder",
			source:      "file3",
			destination: "file2",
			wantError:   "", // No error expected
		},
		{
			name:        "Test case 2: Source folder does not exist",
			source:      "nonExistentFolder",
			destination: "file2",
			wantError:   "source folder does not exist",
		},
		{
			name:        "Test case 3: Destination folder does not exist",
			source:      "file3",
			destination: "nonExistentFolder",
			wantError:   "destination folder does not exist",
		},
		{
			name:        "Test case 4: Cannot move folder to itself",
			source:      "file3",
			destination: "file3",
			wantError:   "cannot move a folder to itself",
		},
		{
			name:        "Test case 5: Cannot move folder to a child of itself",
			source:      "file2",
			destination: "file3",
			wantError:   "cannot move a folder to a child of itself",
		},
		{
			name:        "Test case 6: Cannot move folder to a different organization",
			source:      "file6",
			destination: "file5",
			wantError:   "cannot move a folder to a different organization",
		},
	}

	// Run through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get the source and destination folder paths before moving
			updatedFolders, err := driver.MoveFolder(tt.source, tt.destination)

			// Check if the error matches the expected error
			if err != nil {
				if err.Error() != tt.wantError {
					t.Errorf("MoveFolder() error = %v, want %v", err.Error(), tt.wantError)
				}
			} else if tt.wantError != "" {
				t.Errorf("MoveFolder() expected error = %v, got no error", tt.wantError)
			}

			// If there's no error, verify the folder structure
			if err == nil {
				sourcePath := ""
				destPath := ""

				// Find the new path for source and destination folders
				for _, folder := range updatedFolders {
					if folder.Name == tt.source {
						sourcePath = folder.Paths
					}
					if folder.Name == tt.destination {
						destPath = folder.Paths
					}
				}

				// Verify if the move was successful by comparing paths
				if !reflect.DeepEqual(sourcePath, destPath+"."+tt.source) {
					t.Errorf("MoveFolder() moved folder has incorrect path: got %v, want %v", sourcePath, destPath+"."+tt.source)
				}
			}
		})
	}
}
