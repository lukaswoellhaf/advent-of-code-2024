package tools

import "os"

func ReadFile(filepath string) (string, error) {
	content, readErr := os.ReadFile(filepath)
	if readErr != nil {
		return "", readErr
	}

	return string(content), nil
}
