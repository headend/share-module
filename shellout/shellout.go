package shellout

import (
	"bytes"
	"os/exec"
)

func ExecuteCommand(shell string, command string) (error, int, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()
	return err, exitCode, stdout.String(), stderr.String()
}
