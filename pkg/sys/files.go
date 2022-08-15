package sys

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const MdExtension = ".md"

type Path string

func GetNotesPaths(folderPath Path, extension string) ([]Path, error) {
	var paths []Path

	err := filepath.Walk(string(folderPath),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Ext(path) == extension {
				paths = append(paths, Path(path))
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func GetFilename(path Path) string {
	filename := filepath.Base(string(path))
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func ReadFile(path Path) ([]string, error) {
	file, err := os.Open(string(path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// WriteToFile overwrites file content
func WriteToFile(path Path, content []string) {
	f, err := os.Create(string(path))
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range content {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
