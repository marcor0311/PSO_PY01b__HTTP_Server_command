package handlers

import (
	"fmt"
	"os"
	"path/filepath"
)

// /createfile?name=filename&content=text&repeat=x: Genera un archivo escribiendo el texto dado x veces.
func CreateFile(name, content string, repeat int) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("no se pudo obtener el directorio del usuario: %w", err)
	}

	downloadsPath := filepath.Join(homeDir, "Downloads")
	fullPath := filepath.Join(downloadsPath, name)

	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error al crear archivo en Downloads: %w", err)
	}
	defer f.Close()

	for i := 0; i < repeat; i++ {
		_, err := f.WriteString(content + "\n")
		if err != nil {
			return fmt.Errorf("error al escribir en archivo: %w", err)
		}
	}

	return nil
}
// /deletefile?name=filename: Elimina el archivo especificado si existe.
func DeleteFile(name string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("no se pudo obtener el directorio del usuario: %w", err)
	}

	fullPath := filepath.Join(homeDir, "Downloads", name)

	err = os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("error al eliminar archivo en Downloads: %w", err)
	}
	return nil
}