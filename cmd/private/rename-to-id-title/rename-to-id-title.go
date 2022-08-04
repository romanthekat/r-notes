package main

import (
	"fmt"
	"github.com/romanthekat/r-notes/pkg/core"
	"github.com/romanthekat/r-notes/pkg/sys"
	"log"
	"os"
	"path/filepath"
	"strings"
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
		if !core.IsZkId(sys.GetFilename(path)) {
			fmt.Printf("filename of %s is not ZK ID, skipping\n", path)
			continue
		}

		content, err := sys.ReadFile(path)
		if err != nil {
			fmt.Printf("error during reading content of %s: %s\n", path, err)
			continue
		}

		name, err := core.GetNoteNameByNoteContent(content)
		if err != nil {
			fmt.Printf("error during getting note name of %s: %s\n", path, err)
			continue
		}

		newPath := getPathWithIdAndTitle(path, name)
		err = os.Rename(string(path), newPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s renamed to %s\n", path, newPath)
	}
}

func getPathWithIdAndTitle(path sys.Path, name string) string {
	replacer := strings.NewReplacer("/", " ", "\\", " ")
	resultName := replacer.Replace(name)
	resultName = strings.Trim(resultName, " .")

	return fmt.Sprintf("%s/%s %s%s",
		filepath.Dir(string(path)), sys.GetFilename(path), resultName, sys.MdExtension)
}
