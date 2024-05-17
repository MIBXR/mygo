package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CopyFile(oldPath, newPath string, test bool) error {
	if (test) {
		fmt.Println("Copying file from", oldPath, "to", newPath)
		return nil
	}
	
	files, err := os.ReadDir(oldPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		src := filepath.Join(oldPath, file.Name())
		dst := filepath.Join(newPath, file.Name())

		input, err := os.ReadFile(src)
		if err != nil {
			return err
		}

		err = os.WriteFile(dst, input, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
