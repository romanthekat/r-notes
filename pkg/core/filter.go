package core

import "strings"

func FilterNotesBySubstring(notes []*Note, filterSubstring string) []*Note {
	if len(notes) == 0 {
		return notes
	}

	var result []*Note
	for _, note := range notes {
		if strings.Contains(note.Id, filterSubstring) || strings.Contains(note.Name, filterSubstring) {
			result = append(result, note)
		}
	}
	return result
}
