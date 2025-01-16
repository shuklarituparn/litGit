package utils

import (
	"fmt"
	"os/exec"
)

func GitAdd(args ...string) error {
	var cmd *exec.Cmd
	gitAddArgs := append([]string{"add"}, args...)
	cmd = exec.Command("git", gitAddArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Bruh, following errror happened: %s", string(output))
	}
	return nil
}
