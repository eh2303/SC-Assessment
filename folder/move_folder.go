package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	folders := f.folders

	res := []Folder{}

	var new_path string
	var dst_orgid uuid.UUID

	if dst == name {
		return res, errors.New("Error: Cannot move a folder to itself")
	}

	// get path of destination folder
	dst_found := false
	for _, f := range folders {
		if f.Name == dst {
			dst_found = true
			if strings.Contains(f.Paths, name) {
				return res, errors.New("Error: Cannot move a folder to a child of itself")
			}
			new_path = f.Paths
			dst_orgid = f.OrgId
			break
		}
	}
	if !dst_found {
		return res, errors.New("Error: Destination folder does not exist")
	}

	// update path of moving folder and its children
	source_found := false
	for _, f := range folders {
		if f.Name == name {
			source_found = true
			if f.OrgId != dst_orgid {
				return []Folder{}, errors.New("Error: Cannot move a folder to a different organization")
			}
			f.Paths = strings.Join([]string{new_path, name}, ".")
		} else if strings.Contains(f.Paths, name) {
			i := strings.Index(f.Paths, name)
			keep_path := f.Paths[i:]
			f.Paths = strings.Join([]string{new_path, keep_path}, ".")
		}
		res = append(res, f)
	}

	if !source_found {
		return []Folder{}, errors.New("Error: Source folder does not exist")
	}

	return res, nil
}
