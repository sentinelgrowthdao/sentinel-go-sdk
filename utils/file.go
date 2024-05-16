package utils

import (
	"os"
)

// RemoveFile deletes the file at the specified path.
// It returns nil if the file does not exist or is successfully deleted.
// If the file removal fails, it returns an error.
func RemoveFile(path string) error {
	// Check if the file exists at the given path.
	if _, err := os.Stat(path); err != nil {
		return nil
	}

	// Remove the file and return the resulting error, if any.
	return os.Remove(path)
}
