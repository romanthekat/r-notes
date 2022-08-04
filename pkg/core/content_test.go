package core

import (
	"testing"
)

func Test_IsSameContent(t *testing.T) {
	type args struct {
		content1 []string
		content2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same values",
			args: struct {
				content1 []string
				content2 []string
			}{
				[]string{"first line", "second line"},
				[]string{"first line", "second line"},
			},
			want: true,
		},
		{
			name: "diff values diff lines count",
			args: struct {
				content1 []string
				content2 []string
			}{
				[]string{"first line", "second line"},
				[]string{"first line", "second line", "third line"},
			},
			want: false,
		},
		{
			name: "diff values line differs",
			args: struct {
				content1 []string
				content2 []string
			}{
				[]string{"first line", "second line", "third line"},
				[]string{"first line", "second line", "third line differs"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameContent(tt.args.content1, tt.args.content2); got != tt.want {
				t.Errorf("isSameContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
