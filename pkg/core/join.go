package core

const JoinedNotesSeparator = "\n\n"

func JoinContent(notes []*Note) []string {
	var result []string

	for _, note := range notes {
		result = append(result, note.GetContent()...)
		result = append(result, JoinedNotesSeparator)
	}

	return result
}
