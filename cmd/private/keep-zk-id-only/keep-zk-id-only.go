package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
	"os"
	"path/filepath"
)

func main() {
	folder, err := sys.GetNotesFolderArg()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("reading notes at " + folder)

	paths, err := sys.GetNotesPaths(folder, sys.MdExtension)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("found notes:", len(paths))

	for _, path := range paths {
		isZettel, id, name := parseNoteNameByPath(path)
		if !isZettel {
			fmt.Printf("%s is not a zettel\n", path)
			continue
		}

		if name == "" {
			fmt.Printf("%s is already in 'zk id only' format\n", path)
			continue
		}

		content, err := sys.ReadFile(path)
		if err != nil {
			panic(err)
		}

		header := core.GetYamlHeader(id, name, "")

		sys.WriteToFile(path, append(header, content...))

		newPath := getFilepathOnlyId(path, id)
		err = os.Rename(string(path), newPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s renamed to %s\n", path, newPath)
	}
}

func getFilepathOnlyId(oldPath sys.Path, id string) string {
	return filepath.Dir(string(oldPath)) + "/" + id + ".md"
}

func parseNoteNameByPath(path sys.Path) (isZettel bool, id, name string) {
	if filepath.Ext(string(path)) != ".md" {
		return false, "", ""
	}

	return core.ParseNoteFilename(sys.GetFilename(path))
}
