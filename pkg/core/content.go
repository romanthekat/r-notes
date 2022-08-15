package core

// IsSameContent compares to slices of strings
// TODO having ~contentHash, e.g. more generic approach for detecting changes, may be useful
// generate it on-the-fly
func IsSameContent(content1, content2 []string) bool {
	if len(content1) != len(content2) {
		return false
	}

	for i, line := range content1 {
		if line != content2[i] {
			return false
		}
	}

	return true
}
