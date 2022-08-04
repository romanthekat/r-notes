package common

func JoinContent(notes []*Note) []string {
	var result []string

	notesBreak := []string{"  \t", "  \t"}
	for _, note := range notes {
		result = append(result, note.Content...)
		result = append(result, notesBreak...)
	}

	return result
}
