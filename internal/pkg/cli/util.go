package cli

import (
	"os"
	"os/exec"
)

func VerifyDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, os.ModeDir)
	}
	return nil
}

func VerifyCmd(cmd string) error {
	if _, err := exec.LookPath(cmd); err != nil {
		return err
	}
	return nil
}

func ExecCmd(cmd string) error {
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		return err
	}
	
	return nil
}