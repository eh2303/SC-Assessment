package folder_test

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"testing"
	"errors"
)

func Test_folder_MoveFolder(t *testing.T) {
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
		dst     string
		folders []folder.Folder
		want    []folder.Folder
		exp_err error
	}{
		{ 
			name:    "move g to b",
			source: "g",
			dst:     "b",
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want:    []folder.Folder{
				{"a", org1, "a"},
				{"b", org1, "a.b"},
				{"c", org2, "c"},
				{"d", org2, "c.d"},
				{"e", org1, "a.b.e"},
				{"f", org1, "a.b.e.f"},
				{"g", org1, "a.b.g"},
			},
			exp_err: nil,
		}, 
		{ 
			name:    "move b to g, and check b children paths updated",
			source:	"b",
			dst:     "g",
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want:    []folder.Folder{
				{"a", org1, "a"},
				{"b", org1, "a.g.b"},
				{"c", org2, "c"},
				{"d", org2, "c.d"},
				{"e", org1, "a.g.b.e"},
				{"f", org1, "a.g.b.e.f"},
				{"g", org1, "a.g"},
			},
			exp_err: nil,
		}, 
		{ 
			name:    "move folder to itself",
			source: "b",
			dst:     "b",
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want:    []folder.Folder{},
			exp_err: errors.New("Error: Cannot move a folder to itself"),
		}, 
		{ 
			name:    "move folder to own child",
			source: "b",
			dst:     "e",
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want:    []folder.Folder{},
			exp_err: errors.New("Error: Cannot move a folder to a child of itself"),
		}, 
		{ 
			name:    "dst folder does not exist",
			source: "b",
			dst:     "non",
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want:    []folder.Folder{},
			exp_err: errors.New("Error: Destination folder does not exist"),
		}, 
		{ 
			name:    "move to different org",
			source:	"b",
			dst:     "c",
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want:    []folder.Folder{},
			exp_err: errors.New("Error: Cannot move a folder to a different organization"),
		},
		{ 
			name:    "source folder does not exist",
			source: "none",
			dst:     "c",
			folders: []folder.Folder{a, b, c, d, e, f, g},
			want:    []folder.Folder{},
			exp_err: errors.New("Error: Source folder does not exist"),
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.source, tt.dst)
			if len(get) != len(tt.want) {
				t.Errorf("Expected %d folders, got %d", len(tt.want), len(get))
			}
			if err != nil && tt.exp_err != nil && err.Error() != tt.exp_err.Error() {
				t.Errorf("Error does not match")
			}
			if err == nil && tt.exp_err == nil {
				for i := range get {
					if get[i] != tt.want[i] {
						t.Errorf("Folders don't match")
					}
				}
			}

		})
	}
}
