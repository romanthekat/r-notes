package main

import (
	"github.com/romanthekat/r-notes/pkg/common"
	"reflect"
	"testing"
)

func Test_getNotesOutline(t *testing.T) {
	type args struct {
		note    *common.Note
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
				note: common.NewNote(
					"202105122138",
					"note",
					"/path/to/202105122138 note.md",
					[]string{""},
					[]*common.Note{
						{
							Id:   "202105122139",
							Name: "child",
							Path: "path/to/202105122139 child.md",
						},
					}),
				padding: "",
				result:  []string{},
			},
			want: []string{"- note [[202105122138]]  ", "    - child [[202105122139]]  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNotesOutline(tt.args.note, tt.args.padding, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNotesOutline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNoteName(t *testing.T) {
	type args struct {
		path common.Path
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
			if got := common.GetFilename(tt.args.path); got != tt.want {
				t.Errorf("GetFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
