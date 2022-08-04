package zk

import (
	"github.com/romanthekat/r-notes/pkg/render"
	"testing"
)

func Test_formatIdAsDate(t *testing.T) {
	type args struct {
		zkId string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple example",
			args: args{
				zkId: "202105091600",
			},
			want: "2021-05-09 16:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := render.FormatIdAsIsoDate(tt.args.zkId); got != tt.want {
				t.Errorf("FormatIdAsIsoDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZkId(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "correct zk id",
			args: args{
				id: "202202222020",
			},
			want: true,
		},
		{
			name: "incorrect zk id",
			args: args{
				id: "not exactly 202202222020",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZkId(tt.args.id); got != tt.want {
				t.Errorf("IsZkId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNoteNameByNoteContent(t *testing.T) {
	type args struct {
		content []string
	}
	tests := []struct {
		name     string
		args     args
		wantName string
		wantErr  bool
	}{
		{
			name: "by markdown header",
			args: args{
				content: []string{"# markdown header", "some content", "another line of content"},
			},
			wantName: "markdown header",
			wantErr:  false,
		},
		{
			name: "by yaml header",
			args: args{
				content: []string{"---", "title: yaml header", "---",
					"# markdown header", "some content", "another line of content"},
			},
			wantName: "yaml header",
			wantErr:  false,
		},
		{
			name: "mixed header - pick yaml",
			args: args{
				content: []string{"---", "title: yaml header", "---",
					"# markdown header", "some content", "another line of content"},
			},
			wantName: "yaml header",
			wantErr:  false,
		},
		{
			name: "no header -> error",
			args: args{
				content: []string{"no header", "just some plain text"},
			},
			wantName: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, err := GetNoteNameByNoteContent(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNoteNameByNoteContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotName != tt.wantName {
				t.Errorf("GetNoteNameByNoteContent() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}

func TestParseNoteFilename(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name         string
		args         args
		wantIsZettel bool
		wantId       string
		wantName     string
	}{
		{
			name: "zk id with name",
			args: args{
				filename: "202202021234 the name",
			},
			wantIsZettel: true,
			wantId:       "202202021234",
			wantName:     "the name",
		},
		{
			name: "zk id without name",
			args: args{
				filename: "202202030405",
			},
			wantIsZettel: true,
			wantId:       "202202030405",
			wantName:     "",
		},
		{
			name: "not a zk",
			args: args{
				filename: "some filename",
			},
			wantIsZettel: false,
			wantId:       "",
			wantName:     "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsZettel, gotId, gotName := ParseNoteFilename(tt.args.filename)
			if gotIsZettel != tt.wantIsZettel {
				t.Errorf("ParseNoteFilename() gotIsZettel = %v, want %v", gotIsZettel, tt.wantIsZettel)
			}
			if gotId != tt.wantId {
				t.Errorf("ParseNoteFilename() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotName != tt.wantName {
				t.Errorf("ParseNoteFilename() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
