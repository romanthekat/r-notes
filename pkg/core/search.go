package core

import "strings"

func GetTopNotes(notes []*Note, level int) []*Note {
	var result []*Note

	for _, note := range notes {
		if note.Level <= level {
			result = append(result, note)
		}
	}

	return result
}

func GetRelevantNotes(mainNote *Note, notes []*Note) []*Note {
	result := append([]*Note{}, mainNote.Links...)

	result = append(result, mainNote.Links...)
	result = append(result, mainNote.Backlinks...)
	//neighbors, f.e. 54a.1 is neighbor/child to 54a
	for _, note := range notes {
		if strings.HasPrefix(note.Id, mainNote.Id) {
			result = append(result, note)
		}
	}

	return result
}
