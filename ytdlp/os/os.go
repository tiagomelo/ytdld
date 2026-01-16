package os

import "os"

// OSOperations defines an interface for OS operations.
type OSOperations interface {
	// MkdirTemp creates a new temporary directory in the directory dir
	// with a name beginning with pattern and returns the pathname of the new directory.
	MkdirTemp(dir string, pattern string) (string, error)

	// WriteFie writes data to a file named by filename.
	WriteFile(name string, data []byte, perm os.FileMode) error

	// RemoveAll removes path and any children it contains.
	RemoveAll(path string) error
}

// OSOperationsProvider is the default implementation of OSOperations
type OSOperationsProvider struct{}

// MkdirTemp creates a new temporary directory in the directory dir
// with a name beginning with pattern and returns the pathname of the new directory.
func (OSOperationsProvider) MkdirTemp(dir string, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

// WriteFie writes data to a file named by filename.
func (OSOperationsProvider) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

// RemoveAll removes path and any children it contains.
func (OSOperationsProvider) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
