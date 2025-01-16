package utils

import (
	"fmt"
	"os/exec"
)

func GitInit() error {
	var cmd *exec.Cmd
	cmd = exec.Command("git", "init")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Bruh git init be lackin' : %s")
	}
	fmt.Println(string(out))
	return nil
}
