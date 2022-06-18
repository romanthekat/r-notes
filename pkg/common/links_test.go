package common

import (
	"reflect"
	"testing"
)

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
				content: []string{
					"text [[202202261908]] and",
					"[[202202261909]] another link",
					"duplicate [[202202261909]] ",
					"[[202202261910]]",
					"[[text link]]",
					"[[text link three spaces]]",
					"[[202204101811]] and then [[202204101812]]",
					"[[ broken link]]",
					"[[202206182033 link with title]]",
					"[[202012122011 Σ programming]]",
				},
			},
			want: []string{"202012122011", "202202261908", "202202261909", "202202261910",
				"202204101811", "202204101812", "202206182033", "text link", "text link three spaces"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWikilinks(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWikilinks() = %v, want %v", got, tt.want)
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
				note: NewNote("202202261406", "A name", "", []string{}),
			},
			want: "[[202202261406 A name]]",
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
		NewNote("202202261406", "Note one", "", []string{"[[202202261407]]", "[[202202261408]]"}),
		NewNote("202202261407", "Note two", "", []string{"[[202202261406]]"}),
		NewNote("202202261408", "Note three", "", []string{"[[202202261407]]"}),
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

func Test_generateContentWithBacklinks(t *testing.T) {
	type args struct {
		note *Note
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "note without backlinks",
			args: args{
				note: NewNoteWithLinks("202202261746",
					"Note without backlinks",
					"",
					[]string{
						"# Note without backlinks",
						"line one",
						"line two"},
					nil,
					[]*Note{
						NewNote("202202261747", "The backlink", "", nil),
					},
				),
			},
			want: []string{
				"# Note without backlinks",
				"line one",
				"line two",
				BacklinksHeader,
				"- [[202202261747 The backlink]]",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateContentWithBacklinks(tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateContentWithBacklinks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateContentWithBacklinks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
