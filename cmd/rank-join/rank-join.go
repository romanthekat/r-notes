package main

import (
	"flag"
	"fmt"
	"github.com/romanthekat/r-notes/pkg/common"
	"log"
)

func main() {
	folder, outputPath, err := parseArguments()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("obtaining notes")
	notes := getNotes(folder)

	notes = common.SortByRank(notes)
	//for _, note := range notes {
	//	fmt.Println(note.String())
	//}

	//result := common.JoinContent(notes)
	//common.WriteToFile(outputPath, result)

	log.Println("file saved to", outputPath)
}

func getNotes(folder common.Path) []*common.Note {
	paths, err := common.GetNotesPaths(folder, common.MdExtension)
	if err != nil {
		log.Fatal(err)
	}

	notes := common.NewNotesByPaths(paths)
	common.FillLinks(notes)
	common.FillTags(notes)

	return notes
}

func parseArguments() (common.Path, common.Path, error) {
	notesPath := flag.String("notesPath", "", "a path to notes folder")
	outputPath := flag.String("outputPath", "./", "a path to result join file")
	flag.Parse()

	if *notesPath == "" || *outputPath == "" {
		return "", "", fmt.Errorf("provide both 'notesPath' and 'outputPath'")
	}

	return common.Path(*notesPath), common.Path(*outputPath), nil
}
