package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"log"
	"sort"
)

func main() {
	folder, err := core.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	paths, err := core.GetNotesPaths(folder, core.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, path := range paths {
		note := core.NewNoteByPath(path)
		if note.HasId() {
			result = append(result, fmt.Sprintf("[[%s]] %s", note.Id, note.Name))
		}
	}

	sort.Strings(result)

	for _, entry := range result {
		fmt.Println(entry)
	}
}
