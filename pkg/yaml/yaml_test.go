package yaml

import (
	"github.com/romanthekat/r-notes/pkg/sys"
	"reflect"
	"testing"
)

func TestParseForYamlHeader(t *testing.T) {
	type args struct {
		content []string
	}
	tests := []struct {
		name string
		args args
		want *Header
	}{
		{
			name: "top yaml header",
			args: struct{ content []string }{content: []string{
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
				"# header",
				"some value",
			}},
			want: &Header{
				Content: []string{
					"---",
					"title: ",
					"date: ",
					"tags: ",
					"---"},
				From: 0,
				To:   4,
			},
		},
		{
			name: "bottom yaml header",
			args: struct{ content []string }{content: []string{
				"# header",
				"some value",
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
			}},
			want: &Header{
				Content: []string{
					"---",
					"title: ",
					"date: ",
					"tags: ",
					"---"},
				From: 2,
				To:   6,
			},
		},
		{
			name: "no yaml header",
			args: struct{ content []string }{content: []string{
				"# header",
				"some value",
			}},
			want: &Header{
				Content: nil,
				From:    -1,
				To:      -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractYamlHeader(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractYamlHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlHeader_Exists(t *testing.T) {
	type fields struct {
		Content []string
		From    int
		To      int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "existing yaml header",
			fields: struct {
				Content []string
				From    int
				To      int
			}{
				Content: []string{}, From: 0, To: 10},
			want: true,
		},
		{
			name: "non-existing yaml header",
			fields: struct {
				Content []string
				From    int
				To      int
			}{
				Content: []string{}, From: -1, To: -1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := Header{
				Content: tt.fields.Content,
				From:    tt.fields.From,
				To:      tt.fields.To,
			}
			if got := y.Exists(); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveHeaderFromTopToBottom(t *testing.T) {
	type args struct {
		path    sys.Path
		content []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{
			name: "header moved from top to bottom",
			args: struct {
				path    sys.Path
				content []string
			}{path: "/some/path/file.md", content: []string{
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
				"# header",
				"some value",
			}},
			want: []string{
				"# header",
				"some value",
				" ",
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
			},
			want1: true,
		},
		{
			name: "header not on top - not moved",
			args: struct {
				path    sys.Path
				content []string
			}{path: "/some/path/file.md", content: []string{
				"# header",
				"some value",
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
			}},
			want: []string{
				"# header",
				"some value",
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
			},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MoveHeaderFromTopToBottom(tt.args.path, tt.args.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveHeaderFromTopToBottom() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MoveHeaderFromTopToBottom() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRemoveHeader(t *testing.T) {
	type args struct {
		path    sys.Path
		content []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{
			name: "header removed from top",
			args: struct {
				path    sys.Path
				content []string
			}{path: "/some/path/file.md", content: []string{
				"---",
				"title: ",
				"date: ",
				"tags: #meow, #verymeow",
				"---",
				"# header",
				"some value",
			}},
			want: []string{
				"# header",
				"#meow, #verymeow",
				"some value",
			},
			want1: true,
		},
		{
			name: "header tags empty",
			args: struct {
				path    sys.Path
				content []string
			}{path: "/some/path/file.md", content: []string{
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
				"# header",
				"some value",
			}},
			want: []string{
				"# header",
				"some value",
			},
			want1: true,
		},
		{
			name: "header not on top - not moved",
			args: struct {
				path    sys.Path
				content []string
			}{path: "/some/path/file.md", content: []string{
				"# header",
				"some value",
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
			}},
			want: []string{
				"# header",
				"some value",
				"---",
				"title: ",
				"date: ",
				"tags: ",
				"---",
			},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RemoveHeader(tt.args.path, tt.args.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveHeader() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RemoveHeader() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
