package main

import (
	"github.com/romanthekat/r-notes/pkg/sys"
	"github.com/romanthekat/r-notes/pkg/yaml"
	"log"
)

func main() {
	folder, err := sys.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}

	paths, err := sys.GetNotesPaths(folder, sys.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range paths {
		content, err := sys.ReadFile(path)
		if err != nil {
			log.Printf("error while reading %s: %s\n", path, err)
			continue
		}

		updatedContent, canBeMoved := yaml.MoveHeaderFromTopToBottom(path, content)
		if canBeMoved {
			sys.WriteToFile(path, updatedContent)
			log.Printf("yaml header moved to bottom for file %s\n", path)
		}
	}
}
