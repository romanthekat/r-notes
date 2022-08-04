package md

import "strings"

func IsFirstLevelHeader(line string) bool {
	return strings.HasPrefix(line, "# ")
}

//IsMarkdownHeader TODO more reliable parsing would be beneficial
func IsMarkdownHeader(line string) bool {
	return strings.HasPrefix(line, "# ") || strings.HasPrefix(line, "## ") || strings.HasPrefix(line, "### ")
}
