package core

const JoinedNotesSeparator = "  \t\n  \t\n"

func JoinContent(notes []*Note) []string {
	var result []string

	for _, note := range notes {
		result = append(result, note.Content...)
		result = append(result, JoinedNotesSeparator)
	}

	return result
}
