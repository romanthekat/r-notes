package yaml

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
	"strings"
)

const YamlDelimiter = "---"
const (
	yamlNotFound = iota
	readingYaml
)

type YamlHeader struct {
	Content  []string
	From, To int
}

func (y YamlHeader) String() string {
	return fmt.Sprintf("%d:%d\n%s", y.From, y.To, y.Content)
}

func NewYamlHeader(content []string, from int, to int) *YamlHeader {
	return &YamlHeader{Content: content, From: from, To: to}
}

func (y YamlHeader) Exists() bool {
	return y.From != -1
}

func ParseForYamlHeader(content []string) *YamlHeader {
	from := -1
	state := yamlNotFound

	for i, line := range content {
		if strings.HasPrefix(line, YamlDelimiter) {
			switch state {
			case yamlNotFound:
				state = readingYaml
				from = i
			case readingYaml:
				to := i
				return NewYamlHeader(content[from:to+1], from, to)
			}
		}
	}

	return NewYamlHeader(nil, -1, -1)
}

func MoveHeaderFromTopToBottom(path sys.Path, content []string) ([]string, bool) {
	header := ParseForYamlHeader(content)
	if !header.Exists() {
		log.Printf("file %s doesn't have yaml header, skipping\n", path)
		return content, false
	}

	if header.From > 0 {
		log.Printf("file %s has yaml header not on top, skipping\n", path)
		return content, false
	}

	body := content[header.To+1:]

	result := append(body, " ")
	return append(result, header.Content...), true
}

func RemoveHeader(path sys.Path, content []string) ([]string, bool) {
	header := ParseForYamlHeader(content)
	if !header.Exists() {
		log.Printf("file %s doesn't have yaml header, skipping\n", path)
		return content, false
	}

	if header.From > 0 {
		log.Printf("file %s has yaml header not on top, skipping\n", path)
		return content, false
	}

	body := content[header.To+1:]

	noteHeader := body[0]
	if !core.IsFirstLevelHeader(noteHeader) {
		log.Printf("file %s first line after yaml header is not markdown header, skipping\n", path)
		return content, false
	}

	result := []string{noteHeader}

	hasTags, tags := getTagsFromYaml(header)
	if hasTags {
		result = append(result, tags)
	}

	return append(result, body[1:]...), true
}

func getTagsFromYaml(header *YamlHeader) (bool, string) {
	for _, line := range header.Content {
		line = strings.TrimSpace(line)

		tagsPrefix := "tags:"
		prefixLength := len(tagsPrefix)

		hasTagsPrefix := strings.HasPrefix(line, tagsPrefix)
		if !hasTagsPrefix {
			continue
		}

		if len(line) == prefixLength {
			continue
		}

		return true, strings.TrimSpace(line[prefixLength:])
	}

	return false, ""
}
