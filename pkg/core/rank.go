package core

import "sort"

var tagsRanks = map[string]int{
	"flag":    42,
	"index":   32,
	"project": 32,
	"howto":   16,
}

func GetNoteRank(note *Note) int {
	rank := 0

	rank += len(note.Links)
	rank += len(note.Backlinks) * 4

	//assumption is: higher level notes are higher in priority/rank
	rank += (5 - note.Level) * 2

	for tag := range note.Tags {
		if tagRank, ok := tagsRanks[tag]; ok {
			rank += tagRank
		} else {
			rank += 2
		}
	}

	return rank
}

func SortByRank(notes []*Note) []*Note {
	sort.SliceStable(notes, func(i, j int) bool {
		return GetNoteRank(notes[i]) > GetNoteRank(notes[j])
	})

	return notes
}
