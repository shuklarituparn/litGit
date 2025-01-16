package utils

import (
	"fmt"
	"os/exec"
)

func GitReset(commitHash string) error {
	var cmd *exec.Cmd
	cmd = exec.Command("git", "reset", "--hard", commitHash)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git reset --hard %s failed: %s", commitHash, err)
	}
	fmt.Println(string(output))
	return nil

}
