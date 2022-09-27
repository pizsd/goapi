package file

import "os"

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Put(content []byte, filepath string) error {
	if err := os.WriteFile(filepath, content, 0644); err != nil {
		return err
	}
	return nil
}
