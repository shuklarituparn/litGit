package utils

import (
	"fmt"
	"os/exec"
)

func GitCommit(msg string, all bool) error {
	var cmd *exec.Cmd
	if all {
		cmd = exec.Command("git", "commit", "-am", msg)
	} else {
		cmd = exec.Command("git", "commit", "-m", msg)
	}
	outPut, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git commit failed: %s", string(outPut))
	}
	return nil
}
