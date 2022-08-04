package main

import (
	"github.com/romanthekat/r-notes/pkg/core"
	"log"
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

	for _, path := range paths {
		content, err := core.ReadFile(path)
		if err != nil {
			log.Printf("error while reading %s: %s\n", path, err)
			continue
		}

		updatedContent, canBeMoved := core.MoveHeaderFromTopToBottom(path, content)
		if canBeMoved {
			core.WriteToFile(path, updatedContent)
			log.Printf("yaml header moved to bottom for file %s\n", path)
		}
	}
}
