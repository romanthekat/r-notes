package core

import (
	"github.com/romanthekat/r-notes/pkg/sys"
	"reflect"
	"sync"
	"testing"
)

func TestGetLevel(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple folgezettel id",
			args: args{id: "1.2.3a.4"},
			want: 4,
		},
		{
			name: "plain id without levels",
			args: args{id: "123"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLevel(tt.args.id); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNoteById(t *testing.T) {
	type args struct {
		notes []*Note
		id    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Note
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNoteById(tt.args.notes, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNoteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNoteById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNoteByPath(t *testing.T) {
	type args struct {
		notes    []*Note
		notePath sys.Path
	}
	tests := []struct {
		name    string
		args    args
		want    *Note
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNoteByPath(tt.args.notes, tt.args.notePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNoteByPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNoteByPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNotes(t *testing.T) {
	type args struct {
		folder sys.Path
	}
	tests := []struct {
		name    string
		args    args
		want    []*Note
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNotes(tt.args.folder)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNotes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNotesDetailed(t *testing.T) {
	type args struct {
		folder    sys.Path
		fillLinks bool
		fillTabs  bool
	}
	tests := []struct {
		name    string
		args    args
		want    []*Note
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNotesDetailed(tt.args.folder, tt.args.fillLinks, tt.args.fillTabs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotesDetailed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNotesDetailed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNote(t *testing.T) {
	type args struct {
		id      string
		name    string
		path    sys.Path
		content []string
	}
	tests := []struct {
		name string
		args args
		want *Note
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNote(tt.args.id, tt.args.name, tt.args.path, tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNoteByPath(t *testing.T) {
	type args struct {
		path sys.Path
	}
	tests := []struct {
		name string
		args args
		want *Note
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteByPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteByPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNoteWithLinks(t *testing.T) {
	type args struct {
		id        string
		name      string
		path      sys.Path
		content   []string
		links     []*Note
		backlinks []*Note
	}
	tests := []struct {
		name string
		args args
		want *Note
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteWithLinks(tt.args.id, tt.args.name, tt.args.path, tt.args.content, tt.args.links, tt.args.backlinks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteWithLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNoteWithTags(t *testing.T) {
	type args struct {
		id      string
		name    string
		path    sys.Path
		content []string
		tags    map[string]any
	}
	tests := []struct {
		name string
		args args
		want *Note
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteWithTags(tt.args.id, tt.args.name, tt.args.path, tt.args.content, tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteWithTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNotesByPaths(t *testing.T) {
	type args struct {
		paths []sys.Path
	}
	tests := []struct {
		name string
		args args
		want []*Note
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotesByPaths(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotesByPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNote_GetContent(t *testing.T) {
	type fields struct {
		Id          string
		Name        string
		Content     []string
		Level       int
		Links       []*Note
		Backlinks   []*Note
		Tags        map[string]any
		Path        sys.Path
		loadContent *sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Note{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Content:     tt.fields.Content,
				Level:       tt.fields.Level,
				Links:       tt.fields.Links,
				Backlinks:   tt.fields.Backlinks,
				Tags:        tt.fields.Tags,
				Path:        tt.fields.Path,
				loadContent: tt.fields.loadContent,
			}
			if got := n.GetContent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNote_HasId(t *testing.T) {
	type fields struct {
		Id          string
		Name        string
		Content     []string
		Level       int
		Links       []*Note
		Backlinks   []*Note
		Tags        map[string]any
		Path        sys.Path
		loadContent *sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Note{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Content:     tt.fields.Content,
				Level:       tt.fields.Level,
				Links:       tt.fields.Links,
				Backlinks:   tt.fields.Backlinks,
				Tags:        tt.fields.Tags,
				Path:        tt.fields.Path,
				loadContent: tt.fields.loadContent,
			}
			if got := n.HasId(); got != tt.want {
				t.Errorf("HasId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNote_String(t *testing.T) {
	type fields struct {
		Id          string
		Name        string
		Content     []string
		Level       int
		Links       []*Note
		Backlinks   []*Note
		Tags        map[string]any
		Path        sys.Path
		loadContent *sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Note{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Content:     tt.fields.Content,
				Level:       tt.fields.Level,
				Links:       tt.fields.Links,
				Backlinks:   tt.fields.Backlinks,
				Tags:        tt.fields.Tags,
				Path:        tt.fields.Path,
				loadContent: tt.fields.loadContent,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
