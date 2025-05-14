package handlers_test

import (
	"os"
	"path/filepath"
	"testing"

	"httpserver/internal/handlers"
)

// Helper para obtener ruta completa en ~/Downloads
func getTestFilePath(name string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Downloads", name)
}

func TestCreateFile_Basic(t *testing.T) {
	filename := "test_createfile.txt"
	fullPath := getTestFilePath(filename)

	// Asegúrate de eliminar si ya existe
	os.Remove(fullPath)

	err := handlers.CreateFile(filename, "contenido de prueba", 3)
	if err != nil {
		t.Fatalf("Error while creating file: %v", err)
	}

	// Verifica que el archivo exista
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Errorf("File %s was not created", fullPath)
	}
}

func TestDeleteFile_FileExists(t *testing.T) {
	filename := "test_deletefile.txt"
	fullPath := getTestFilePath(filename)

	err := handlers.CreateFile(filename, "delete", 1)
	if err != nil {
		t.Fatalf("Error while creating file: %v", err)
	}

	err = handlers.DeleteFile(filename)
	if err != nil {
		t.Fatalf("Error while deleting file: %v", err)
	}

	if _, err := os.Stat(fullPath); err == nil {
		t.Error("File was not deleted, it still exists")
	}
}

func TestDeleteFile_NotExist(t *testing.T) {
	err := handlers.DeleteFile("archivo_inexistente.txt")
	if err == nil {
		t.Error("Expected error when deleting non-existent file, but got none")
	}
}
