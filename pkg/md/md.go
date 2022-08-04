package md

import "strings"

func IsFirstLevelHeader(line string) bool {
	return strings.HasPrefix(line, "# ")
}
