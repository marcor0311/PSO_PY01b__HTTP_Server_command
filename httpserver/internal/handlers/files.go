package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var fileLock sync.Mutex

// /createfile?name=filename&content=text&repeat=x: Generates a file by writing the given text x times.
func CreateFile(name, content string, repeat int) error {
	// Locking to ensure safe concurrent access to the file system.
	fileLock.Lock()
	defer fileLock.Unlock()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user's home directory: %w", err)
	}

	downloadsPath := filepath.Join(homeDir, "Downloads")
	fullPath := filepath.Join(downloadsPath, name)

	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file in Downloads: %w", err)
	}
	defer f.Close()

	for i := 0; i < repeat; i++ {
		_, err := f.WriteString(content + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}

// /deletefile?name=filename: Elimina el archivo especificado si existe.
func DeleteFile(name string) error {
	// Locking to ensure safe concurrent access to the file system.
	fileLock.Lock()
	defer fileLock.Unlock()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user's home directory: %w", err)
	}

	fullPath := filepath.Join(homeDir, "Downloads", name)

	err = os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("error deleting file in Downloads: %w", err)
	}
	return nil
}