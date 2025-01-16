package utils

import (
	"fmt"
	"os/exec"
)

func GitStatus() error {
	var cmd *exec.Cmd
	cmd = exec.Command("git", "status")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("Bruhh, error running git status: %s", err.Error())

	}
	fmt.Println(string(out))
	return nil
}
