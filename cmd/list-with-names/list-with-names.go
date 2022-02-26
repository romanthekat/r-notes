package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
	"sort"
)

func main() {
	folder, err := common.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	paths, err := common.GetNotesPaths(folder, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, path := range paths {
		note := common.NewNoteByPath(path)
		if note.HasId() {
			result = append(result, fmt.Sprintf("[[%s]] %s", note.Id, note.Name))
		}
	}

	sort.Strings(result)

	for _, entry := range result {
		fmt.Println(entry)
	}
}
