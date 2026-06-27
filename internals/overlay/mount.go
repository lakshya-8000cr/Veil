package overlay

import (
	"fmt"
	"os/exec"
)


// in the current architectire iam not using the mount.SysCall() which is mostly used by the orgnizations 
// . we using the command for easy debugging 

func Mount(lower string, upper string, work string, merged string) error {
	opts := fmt.Sprintf(
		"lowerdir=%s,upperdir=%s,workdir=%s",
		lower,
		upper,
		work,
	)

	cmd := exec.Command(
		"sudo",
		"mount",
		"-t",
		"overlay",
		"overlay",
		"-o",
		opts,
		merged,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("mount failed: %s", string(out))
	}

	return nil
}