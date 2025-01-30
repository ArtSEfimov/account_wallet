package files

import (
	"fmt"
	"os"
)

type JSONdb struct {
	filename string
}

func NewJSONdb(filename string) *JSONdb {
	return &JSONdb{
		filename: filename,
	}
}
func (db *JSONdb) Read() ([]byte, error) {
	bytes, readErr := os.ReadFile(db.filename)
	if readErr != nil {
		return nil, fmt.Errorf("reading error file %w", readErr)
	}

	return bytes, nil
}

func (db *JSONdb) Write(content []byte) error {
	file, createErr := os.Create(db.filename)
	if createErr != nil {
		return fmt.Errorf("failed to create file: %w", createErr)
	}

	_, writeErr := file.Write(content)
	if writeErr != nil {
		return fmt.Errorf("failed to write to file: %w", writeErr)
	}
	defer file.Close()
	fmt.Println("Файл записан успешно")
	return nil
}
