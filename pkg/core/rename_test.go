package core

import (
	"github.com/romanthekat/r-notes/pkg/sys"
	"reflect"
	"sync"
	"testing"
)

func Test_getNewPath(t *testing.T) {
	type args struct {
		oldPath sys.Path
		newName string
	}
	tests := []struct {
		name string
		args args
		want sys.Path
	}{
		{
			name: "simple",
			args: args{
				oldPath: "/somewhere/202012051855 zettelkasten.md",
				newName: "42 change_zettelkasten.md",
			},
			want: "/somewhere/42 change_zettelkasten.md",
		},
		{
			name: "local path case",
			args: args{
				oldPath: "./somewhere/202012051855 zettelkasten.md",
				newName: "142 change_zettelkasten.md",
			},
			want: "somewhere/142 change_zettelkasten.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNewPath(tt.args.oldPath, tt.args.newName); got != tt.want {
				t.Errorf("getNewPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateNoteContent(t *testing.T) {
	type args struct {
		note    *Note
		newPath sys.Path
		newName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				note: &Note{
					Id:          "202012051855",
					Name:        "zettelkasten",
					Content:     []string{"# 202012051855 zettelkasten", "some line"},
					Path:        "/somewhere/202012051855 zettelkasten.md",
					loadContent: &sync.Once{},
				},
				newPath: "/somewhere/42 updated_zettelkasten.md",
				newName: "updated_zettelkasten",
			},
			wantErr: false,
		},
		{
			name: "not zettel -> error",
			args: args{
				note: &Note{
					Id:      "202012051855",
					Name:    "zettelkasten",
					Content: []string{"# 202012051855 zettelkasten", "some line"},
					Path:    "/somewhere/202012051855 zettelkasten.md",
				},
				newPath: "/somewhere/NOT_ZETTEL updated_zettelkasten.md",
				newName: "updated_zettelkasten",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := updateNoteContent(tt.args.note, tt.args.newPath, tt.args.newName); (err != nil) != tt.wantErr {
				t.Errorf("updateNoteContent() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && (tt.args.note.Path != tt.args.newPath || tt.args.note.Name != tt.args.newName) {
				t.Errorf("path or name did not match: path: %v, want: %v; name: %v, want: %v",
					tt.args.note.Path, tt.args.newPath, tt.args.note.Name, tt.args.newName)
			}
		})
	}
}

func Test_updateLinks(t *testing.T) {
	type args struct {
		note    *Note
		oldLink string
		newLink string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple",
			args: args{
				note: NewNote("42", "some note", "", []string{
					"# 42 some note",
					"a link -> [[2 some link]], then some text",
				}),
				oldLink: "[[2 some link]]",
				newLink: "[[4 updated link]]",
			},
			want: []string{
				"# 42 some note",
				"a link -> [[4 updated link]], then some text",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateLinks(tt.args.note, tt.args.oldLink, tt.args.newLink); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_syncNoteHeader(t *testing.T) {
	type args struct {
		note *Note
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple markdown header #",
			args: args{
				note: NewNote("42", "note", "", []string{
					"# 4200 old_note",
					"some content",
				}),
			},
			want: "# 42 note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := syncNoteHeader(tt.args.note); got != tt.want {
				t.Errorf("syncNoteHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
