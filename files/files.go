package files

import (
	"fmt"
	"os"
)

func ReadFile(title string) ([]byte, error) {
	bytes, readErr := os.ReadFile(title)
	if readErr != nil {
		return nil, fmt.Errorf("reading error file %w", readErr)
	}
	
	return bytes, nil
}

func WriteFile(title string, content []byte) error {
	file, createErr := os.Create(title)
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
