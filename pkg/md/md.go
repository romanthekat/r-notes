package md

import "strings"

const MarkdownLineBreak = "  "

func IsFirstLevelHeader(line string) bool {
	return strings.HasPrefix(line, "# ")
}

//IsMarkdownHeader TODO more reliable parsing would be beneficial
func IsMarkdownHeader(line string) bool {
	return IsFirstLevelHeader(line) || strings.HasPrefix(line, "## ") || strings.HasPrefix(line, "### ")
}
