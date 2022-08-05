package core

import (
	"reflect"
	"testing"
)

func TestSortByRank(t *testing.T) {
	type args struct {
		notes []*Note
	}
	tests := []struct {
		name string
		args args
		want []*Note
	}{
		{
			name: "tags rank",
			args: args{notes: []*Note{
				NewNoteWithTags("3", "3", "", nil, nil),
				NewNoteWithTags("1", "1", "", nil, map[string]any{"tag1": 0, "tag2": 0}),
				NewNoteWithTags("2", "1", "", nil, map[string]any{"tag": 0}),
			}},
			want: []*Note{
				NewNoteWithTags("1", "1", "", nil, map[string]any{"tag1": 0, "tag2": 0}),
				NewNoteWithTags("2", "1", "", nil, map[string]any{"tag": 0}),
				NewNoteWithTags("3", "3", "", nil, nil),
			},
		},
		{
			name: "special tags rank",
			args: args{notes: []*Note{
				NewNoteWithTags("2", "1", "", nil, map[string]any{"tag1": 0, "tag2": 0}),
				NewNoteWithTags("3", "3", "", nil, nil),
				NewNoteWithTags("1", "1", "", nil, map[string]any{"index": 0}),
			}},
			want: []*Note{
				NewNoteWithTags("1", "1", "", nil, map[string]any{"index": 0}),
				NewNoteWithTags("2", "1", "", nil, map[string]any{"tag1": 0, "tag2": 0}),
				NewNoteWithTags("3", "3", "", nil, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortByRank(tt.args.notes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
