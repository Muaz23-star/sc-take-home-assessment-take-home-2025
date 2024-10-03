package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"reflect"
)

func TestGetAllChildFolders(t *testing.T) {
	t.Parallel()

	// I used a custom sample data for testing as it was easier to understand
	folders := folder.GetSampleData("sample2.json")

	// Create a UUID for testing
	orgID, _ := uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")

	// Initialize the Driver with the folders
	driver := folder.NewDriver(folders)

	// Test cases for the GetAllChildFolders function
	tests := []struct {
		name       string
		orgID      uuid.UUID
		parentName string
		want       []folder.Folder
		wantErr    string // Expected error message, if any
	}{
		{
			name:       "Test case 1: Parent with multiple children",
			orgID:      orgID,
			parentName: "file3",
			want: []folder.Folder{
				{Name: "file4", Paths: "file1.file2.file3.file4", OrgId: orgID},
				{Name: "file5", Paths: "file1.file2.file3.file4.file5", OrgId: orgID},
			},
			wantErr: "", // No error expected
		},
		{
			name:       "Test case 2: Parent with no children",
			orgID:      orgID,
			parentName: "file5",
			want:       []folder.Folder{}, // No children expected
			wantErr:    "",                 // No error expected
		},
		{
			name:       "Test case 3: Parent not found",
			orgID:      orgID,
			parentName: "file6",
			want:       nil, // Should return nil
			wantErr:    "parent folder not found", // Error expected
		},
		{
			name:       "Test case 4: Organization ID not found",
			orgID:      uuid.Must(uuid.NewV4()),
			parentName: "file3",
			want:       nil, // Should return nil
			wantErr:    "organization ID not found", // Error expected
		},
		{
			name:       "Test case 5: Organization ID and parent not found",
			orgID:      uuid.Must(uuid.NewV4()),
			parentName: "file6",
			want:       nil, // Should return nil
			wantErr:    "organization ID not found", // Organization not found error should take precedence
		},
	}

	// Running through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Getting the child folders using the function 
			got, err := driver.GetAllChildFolders(tt.orgID, tt.parentName)

			// Check if the error matches the expected error
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("GetAllChildFolders() error = %v, wantErr %v", err.Error(), tt.wantErr)
				}
			} else if tt.wantErr != "" {
				t.Errorf("GetAllChildFolders() expected error = %v, got no error", tt.wantErr)
			}

			// Comparing the output using reflect.DeepEqual
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllChildFolders() = %v, want %v", got, tt.want)
			}
		})
	}
}


