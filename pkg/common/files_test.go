package common

import "testing"

func Test_GetFilename(t *testing.T) {
	type args struct {
		path Path
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "main",
			args: args{
				path: "/somewhere/zkId/1 note.md",
			},
			want: "1 note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilename(tt.args.path); got != tt.want {
				t.Errorf("GetFullNoteName() = %v, want %v", got, tt.want)
			}
		})
	}
}
