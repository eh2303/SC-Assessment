package folder_test

import (
	"testing"
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	org1, err := uuid.NewV4()
	if err != nil {
        fmt.Println("Error generating UUID for org1:", err)
        return
    }
	org2, err := uuid.NewV4()
	if err != nil {
        fmt.Println("Error generating UUID for org2:", err)
        return
    }
	invalid_org, err := uuid.NewV4()
	if err != nil {
        fmt.Println("Error generating UUID for invalid_org:", err)
        return
    }

	a := folder.Folder{Name: "a", OrgId: org1, Paths: "a"}
    b := folder.Folder{Name: "b", OrgId: org1, Paths: "a.b"}
    c := folder.Folder{Name: "c", OrgId: org2, Paths: "c"}

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{ // get folders in org1
			name: "org1 folders",
			orgID: org1,
			folders: []folder.Folder{a, b, c},
			want: []folder.Folder{a, b},
		},
		{ // no org
			name: "Invalid org folders",
			orgID: invalid_org,
			folders: []folder.Folder{a, b, c},
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 f := folder.NewDriver(tt.folders)
			 get := f.GetFoldersByOrgID(tt.orgID)
			if len(get) != len(tt.want) {
				t.Errorf("Expected %d folders, got %d", len(tt.want), len(get))
			}
			for i := range get {
				if get[i] != tt.want[i] {
					t.Errorf("Folders don't match")
				}
			}
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	org1, err := uuid.NewV4()
	if err != nil {
        fmt.Println("Error generating UUID for org1:", err)
        return
    }
	org2, err := uuid.NewV4()
	if err != nil {
        fmt.Println("Error generating UUID for org2:", err)
        return
    }
	

	a := folder.Folder{Name: "a", OrgId: org1, Paths: "a"}
    b := folder.Folder{Name: "b", OrgId: org1, Paths: "a.b"}
    c := folder.Folder{Name: "c", OrgId: org2, Paths: "c"}
	d := folder.Folder{Name: "d", OrgId: org2, Paths: "c.d"}
	e := folder.Folder{Name: "e", OrgId: org1, Paths: "a.b.e"}
	f := folder.Folder{Name: "f", OrgId: org1, Paths: "a.b.e.f"}
	g := folder.Folder{Name: "g", OrgId: org1, Paths: "a.g"}

	tests := [...]struct {
		name    string
		source	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{ 
			name: "get children of a",
			source: "a",
			orgID: org1,
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want: []folder.Folder{b, e, f, g},
		},
		{ 
			name: "get children of b",
			source: "b",
			orgID: org1,
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want: []folder.Folder{e, f},
		},
		{ 
			name: "no folder in org",
			source: "none",
			orgID: org1,
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want: nil,
		},
		{ 
			name: "wrong orgID",
			source: "a",
			orgID: org2,
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want: nil,
		},
		{ 
			name: "folder has no children", 
			source: "g",
			orgID: org1,
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 f := folder.NewDriver(tt.folders)
			 get := f.GetAllChildFolders(tt.orgID, tt.source)
			if len(get) != len(tt.want) {
				t.Errorf("Expected %d folders, got %d", len(tt.want), len(get))
			}
			for i := range get {
				if get[i] != tt.want[i] {
					t.Errorf("Folders don't match")
				}
			}
		})
	}
}
