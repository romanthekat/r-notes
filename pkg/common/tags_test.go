package common

import (
	"reflect"
	"testing"
)

func Test_getTags(t *testing.T) {
	type args struct {
		content []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "no tags",
			args: args{
				content: []string{"# some header", "some body line", "", "some link [meow]"},
			},
			want: nil,
		},
		{
			name: "several tags",
			args: args{
				content: []string{"# some header", "some body line",
					" text #tag1 and then #tag2", "some link [meow]",
					"not#tag",
					//"[#notatag probably", //TODO fix case
					"ururu #tag-with_symbols hm"},
			},
			want: []string{"tag-with_symbols", "tag1", "tag2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTags(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
