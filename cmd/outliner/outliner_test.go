package main

import (
	"github.com/romanthekat/r-notes/pkg/core"
	"reflect"
	"testing"
)

func Test_getNotesOutline(t *testing.T) {
	type args struct {
		note    *core.Note
		padding string
		result  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "main",
			args: args{
				note: core.NewNoteWithLinks("202105122138", "note", "/path/to/202105122138 note.md", []string{""},
					[]*core.Note{
						core.NewNote("202105122139", "child", "", []string{""}),
					},
					nil),
				padding: "",
				result:  []string{},
			},
			want: []string{"- [[202105122138 note]]  ", "    - [[202105122139 child]]  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNotesOutline(tt.args.note, tt.args.padding, 3, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNotesOutline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNoteName(t *testing.T) {
	type args struct {
		path core.Path
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "main",
			args: args{
				path: "/somewhere/path/note.md",
			},
			want: "note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := core.GetFilename(tt.args.path); got != tt.want {
				t.Errorf("GetFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
