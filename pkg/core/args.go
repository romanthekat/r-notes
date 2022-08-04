package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetNoteFileArgument(extension string) (Path, Path, error) {
	if len(os.Args) != 2 {
		return "", "", fmt.Errorf("specify notes path")
	}

	filename := os.Args[1]
	if filepath.Ext(filename) != extension {
		return "", "", fmt.Errorf("specify %s file path", extension)
	}

	return Path(filename), Path(filepath.Dir(filename)), nil
}

func GetNotesFolderArg() (Path, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf("specify notes folder")
	}

	return Path(os.Args[1]), nil
}
