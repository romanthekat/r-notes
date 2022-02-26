package common

import (
	"reflect"
	"testing"
)

func Test_getFilesByWikilinks(t *testing.T) {
	type args struct {
		currentFile Path
		files       []Path
		wikiLinks   []string
	}
	tests := []struct {
		name string
		args args
		want []Path
	}{
		{
			name: "main",
			args: args{
				currentFile: "path.md",
				files:       []Path{"path.md", "first.md", "second.md", "third.md"},
				wikiLinks:   []string{"first", "third"}},
			want: []Path{"first.md", "third.md"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilesByWikiLinks(tt.args.currentFile, tt.args.files, tt.args.wikiLinks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFilesByWikiLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWikiLinks(t *testing.T) {
	type args struct {
		content []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				content: []string{"text [[link1]] and", "[[link2]] another link", "duplicate [[link1]] "},
			},
			want: []string{"link1", "link2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWikiLinks(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWikiLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNoteLink(t *testing.T) {
	type args struct {
		note *Note
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "wikilinks based link",
			args: args{
				note: NewNote("202202261406", "A name", "", []string{}, nil),
			},
			want: "A name [[202202261406]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNoteLink(tt.args.note); got != tt.want {
				t.Errorf("GetNoteLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

//TODO cover more cases
func TestFillLinks(t *testing.T) {
	notes := []*Note{
		NewNote("202202261406", "Note one", "", []string{"[[202202261407]]", "[[202202261408]]"}, nil),
		NewNote("202202261407", "Note two", "", []string{"[[202202261406]]"}, nil),
		NewNote("202202261408", "Note three", "", []string{"[[202202261407]]"}, nil),
	}

	notes = FillLinks(notes)

	note1 := notes[0]
	note2 := notes[1]
	note3 := notes[2]

	okLinks := hasLinks(note1.Links, note2.Id, note3.Id) &&
		hasLinks(note2.Links, note1.Id) &&
		hasLinks(note3.Links, note2.Id)
	okBacklinks := hasLinks(note1.Backlinks, note2.Id) &&
		hasLinks(note2.Backlinks, note1.Id, note3.Id) &&
		hasLinks(note3.Backlinks, note1.Id)

	if !(okLinks && okBacklinks) {
		t.Errorf("filling (back)links doesn't work, links cycle broken:\n%+v\n%+v\n%+v\n", note1, note2, note3)
	}
}

func hasLinks(notes []*Note, ids ...string) bool {
	if len(notes) != len(ids) {
		return false
	}

	notesSet := make(map[string]struct{})

	for _, note := range notes {
		notesSet[note.Id] = struct{}{}
	}

	for _, id := range ids {
		if _, ok := notesSet[id]; !ok {
			return false
		}

	}

	return true
}
