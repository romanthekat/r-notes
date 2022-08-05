package core

import (
	"reflect"
	"testing"
)

func TestJoinContent(t *testing.T) {
	type args struct {
		notes []*Note
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				notes: []*Note{
					NewNote("1", "1", "", []string{"# 1 1", "1 2"}),
					NewNote("2", "2", "", []string{"# 2 1", "2 2"}),
				},
			},
			want: []string{"# 1 1", "1 2", JoinedNotesSeparator, "# 2 1", "2 2", JoinedNotesSeparator},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinContent(tt.args.notes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
