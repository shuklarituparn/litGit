package utils

import (
	"fmt"
	"os/exec"
)

func GitClone(url string) error {
	var cmd *exec.Cmd
	cmd = exec.Command("git", "clone", url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Bruh git clone be lackin': %v", err)
	}
	fmt.Println(string(output))
	return nil
}
