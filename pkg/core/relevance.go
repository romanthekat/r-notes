package core

import "strings"

func GetRelevantNotes(mainNote *Note, notes []*Note) []*Note {
	result := append([]*Note{}, mainNote.Links...)

	//neighbors, f.e. 54a.1 is neighbor/child to 54a
	for _, note := range notes {
		if strings.HasPrefix(note.Id, mainNote.Id) {
			result = append(result, note)
		}
	}

	result = append(result, mainNote.Backlinks...)

	return result
}
