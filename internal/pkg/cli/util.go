package cli

import (
	"os"
)

func VerifyDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, os.ModeDir)
	}
	return nil
}