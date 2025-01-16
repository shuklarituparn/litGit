package utils

import (
	"fmt"
	"os/exec"
)

func GitPush() error {
	var cmd *exec.Cmd
	cmd = exec.Command("git", "push")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Bruh, git push be lackin: %s", err)
	}
	fmt.Println(string(out))
	return nil
}
