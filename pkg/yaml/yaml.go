package yaml

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/md"
	"github.com/romanthekat/r-notes/pkg/render"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
	"strings"
)

const Delimiter = "---"
const TitleParameter = "title:"
const (
	yamlNotFound = iota
	readingYaml
)

type Header struct {
	Content  []string
	From, To int
}

func (y Header) String() string {
	return fmt.Sprintf("%d:%d\n%s", y.From, y.To, y.Content)
}

func NewYamlHeader(content []string, from int, to int) *Header {
	return &Header{Content: content, From: from, To: to}
}

func (y Header) Exists() bool {
	return y.From != -1
}

func GetYamlHeaderContent(id, name, tags string) []string {
	return []string{
		"---",
		"title: " + strings.ToLower(name),
		"date: " + render.FormatIdAsIsoDate(id),
		"tags: " + tags,
		"---",
	}
}

func ExtractYamlHeader(content []string) *Header {
	from := -1
	state := yamlNotFound

	for i, line := range content {
		if strings.HasPrefix(line, Delimiter) {
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
	header := ExtractYamlHeader(content)
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
	header := ExtractYamlHeader(content)
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
	if !md.IsFirstLevelHeader(noteHeader) {
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

func getTagsFromYaml(header *Header) (bool, string) {
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
