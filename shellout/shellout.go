package shellout

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func ExecuteCommand(shellPath string, commandToRun string) (error, int, string, string) {
	if shellPath == "" {
		shellPath = "/bin/sh"
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shellPath, "-c", commandToRun)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()
	return err, exitCode, stdout.String(), stderr.String()
}

func RunExternalCmd(appToRun string, commandToRun string, defaultTimeOutSeconds time.Duration) (error, int, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var exitCode int
	var err error
	if defaultTimeOutSeconds > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOutSeconds*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, appToRun, commandToRun)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err = cmd.Run()
		exitCode = cmd.ProcessState.ExitCode()
		if ctx.Err() == context.DeadlineExceeded {
			err = fmt.Errorf("Command timed out: ", err.Error())
		}
	} else {
		cmd := exec.Command(appToRun, commandToRun)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err = cmd.Run()
		exitCode = cmd.ProcessState.ExitCode()
	}
	return err, exitCode, stdout.String(), stderr.String()
}
