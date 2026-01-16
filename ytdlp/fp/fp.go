package fp

import "path/filepath"

// FilePathOperations interface abstracts file path operations.
type FilePathOperations interface {
	// Join joins any number of path elements into a single path,
	// adding a separator if necessary.
	Join(elem ...string) string
}

type FilePathOperationsProvider struct{}

// Join joins any number of path elements into a single path,
// adding a separator if necessary.
func (FilePathOperationsProvider) Join(elem ...string) string {
	return filepath.Join(elem...)
}
