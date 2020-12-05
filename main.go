package main

import (
	"fmt"
	"github.com/headend/share-module/shellout"
	"log"
)

func main()  {
	err, exitCode, stdOut, stdErr := shellout.ExecuteCommand("./aaa", "ls")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Exit code: %d\n", exitCode)
	fmt.Printf("Stdout: %s\n", stdOut)
	fmt.Printf("Stderr: %s\n", stdErr)
}
