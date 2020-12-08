package shellout

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func ExecuteCommand(shell string, commandToRun string) (error, int, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", commandToRun)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()
	return err, exitCode, stdout.String(), stderr.String()
}

func RunExternalCmd(shell string, commandToRun string, defaultTimeOutSeconds time.Duration) (error, int, string, string) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOutSeconds*time.Second)
	defer cancel()
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, shell, commandToRun)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()
	if ctx.Err() == context.DeadlineExceeded {
		err = fmt.Errorf("Command timed out: ", err)
	}
	return err, exitCode, stdout.String(), stderr.String()
}
