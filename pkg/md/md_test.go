package md

import "testing"

func TestIsFirstLevelHeader(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "header without id",
			args: args{line: "# a usual header without id"},
			want: true,
		},
		{
			name: "not a md header",
			args: args{line: "level nil"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFirstLevelHeader(tt.args.line); got != tt.want {
				t.Errorf("IsFirstLevelHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMarkdownHeader(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "header wth id",
			args: args{line: "# 202208052214 a header"},
			want: true,
		},
		{
			name: "header without id",
			args: args{line: "# a usual header without id"},
			want: true,
		},
		{
			name: "two level header",
			args: args{line: "## level two"},
			want: true,
		},
		{
			name: "three level header",
			args: args{line: "### level three"},
			want: true,
		},
		{
			name: "not a md header",
			args: args{line: "level unknown NaN"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMarkdownHeader(tt.args.line); got != tt.want {
				t.Errorf("IsMarkdownHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
