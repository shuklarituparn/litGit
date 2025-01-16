package utils

import (
	"fmt"
	"os/exec"
)

func GitPull() error {
	var cmd *exec.Cmd
	cmd = exec.Command("git", "pull")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Bruh git pull be lackin': %s")
	}
	fmt.Println(string(out))
	return nil
}
