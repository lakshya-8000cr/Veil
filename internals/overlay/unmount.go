package overlay

import (
	"fmt"
	"os/exec"
)

func Unmount(target string) error {  // unmount
	cmd := exec.Command("sudo", "umount", target)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unmount failed: %s", string(out))
	}

	return nil
}