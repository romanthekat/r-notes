package zk

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/md"
	"github.com/romanthekat/r-notes/pkg/yaml"
	"regexp"
	"strings"
)

func GetNoteNameByNoteContent(content []string) (name string, err error) {
	for _, line := range content {
		titleIdx := strings.Index(line, yaml.TitleParameter)
		hasYamlTitleParameter := titleIdx != -1

		hasFirstLevelHeader := md.IsFirstLevelHeader(line)

		//TODO it's better to rely on state machine and real parsing of yaml header
		if hasYamlTitleParameter {
			return line[titleIdx+len(yaml.TitleParameter)+1:], nil
		} else if hasFirstLevelHeader {
			return line[2:], nil
		}
	}

	return "", fmt.Errorf("not possible to detect and extract note Name from file using yaml title or # header")
}

var ValidZkId = regexp.MustCompile(`^[0-9].[a-zA-Z0-9./-]{0,31}$`)

func IsZkId(id string) bool {
	return ValidZkId.MatchString(id)
}

func ParseNoteFilename(filename string) (isZettel bool, id, name string) {
	spaceIndex := strings.Index(filename, " ")
	if spaceIndex == -1 {
		id = filename
	} else {
		id = filename[:spaceIndex]
	}

	if !IsZkId(id) {
		return false, "", ""
	}

	id = strings.TrimSpace(id)
	name = strings.TrimLeft(filename, id)
	name = strings.TrimSpace(name)

	return true, id, name
}
