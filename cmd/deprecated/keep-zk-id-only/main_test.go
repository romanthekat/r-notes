package main

import (
	"github.com/romanthekat/r-notes/pkg/sys"
	"testing"
)

func Test_getFilepathOnlyId(t *testing.T) {
	type args struct {
		path sys.Path
		id   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "main",
			args: args{
				path: "/somewhere/zkId/202105091600 a note.md",
				id:   "202105091600",
			},
			want: "/somewhere/zkId/202105091600.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFilepathOnlyId(tt.args.path, tt.args.id); got != tt.want {
				t.Errorf("getFilepathOnlyId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseNoteName(t *testing.T) {
	type args struct {
		path sys.Path
	}
	tests := []struct {
		name string
		args args
		want struct {
			flag     bool
			id, name string
		}
	}{
		{
			name: "common zettelkasten note",
			args: args{
				path: "/somewhere/zkId/202105091600 note.md",
			},
			want: struct {
				flag     bool
				id, name string
			}{flag: true, id: "202105091600", name: "note"},
		},
		{
			name: "multi word name zettelkasten note",
			args: args{
				path: "/somewhere/zkId/202105091600 multi word.md",
			},
			want: struct {
				flag     bool
				id, name string
			}{flag: true, id: "202105091600", name: "multi word"},
		},
		{
			name: "not a zettelkasten formatted name",
			args: args{
				path: "/somewhere/zkId/that's not a note you are searching for.md",
			},
			want: struct {
				flag     bool
				id, name string
			}{flag: false, id: "", name: ""},
		},
		{
			name: "already formatted name",
			args: args{
				path: "/somewhere/zkId/202105091600.md",
			},
			want: struct {
				flag     bool
				id, name string
			}{flag: true, id: "202105091600", name: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, id, name := parseNoteNameByPath(tt.args.path)

			if flag != tt.want.flag || id != tt.want.id || name != tt.want.name {
				t.Errorf("parseNoteNameByPath() = %t %s %s, want %v", flag, id, name, tt.want)
			}
		})
	}
}
