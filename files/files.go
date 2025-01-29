package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(name string, content string) error {
	file, createErr := os.Create(name)
	if createErr != nil {
		return fmt.Errorf("failed to create file: %w", createErr)
	}

	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		return fmt.Errorf("failed to write to file: %w", writeErr)
	}
	defer file.Close()
	fmt.Println("Файл записан успешно")
	return nil
}
