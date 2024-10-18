package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	folders := f.folders

	res := []Folder{}
	// check for invalid folder, invalid orgID
	folder_found := false
	for _, f := range folders {
		if f.Name == name {
			folder_found = true
			if f.OrgId != orgID {
				fmt.Println("Error: Folder does not exist in the specified organization")
				return nil
			}
			break
		}
	}

	if !folder_found {
		fmt.Println("Error: Folder does not exist")
		return nil
	}

	// get child folders
	for _, f := range folders {
		if f.OrgId == orgID && strings.Contains(f.Paths, name) {
			if f.Name != name { // only include child folders
				res = append(res, f)
			}
		}
	}
	return res
}
