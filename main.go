package main

import (
	"fmt"
	"github.com/headend/share-module/shellout"
	"log"
)

func main()  {
	cmdtorun := []string{"udp://239.0.0.1:1234", "-v", "quiet", "-show_format", "-show_streams", "-print_format", "json"}
	err, exitCode, stdOut, stdErr := shellout.RunExternalCmd("/opt/iptv/sbin/ffprobe", cmdtorun, 20)
	fmt.Println("/opt/iptv/sbin/ffmpeg", cmdtorun)
	if err != nil {
		fmt.Printf("Exit code loi loi: %d\n", exitCode)
		log.Fatal(err)
	}
	fmt.Printf("Exit code: %d\n", exitCode)
	fmt.Printf("Stdout: %s\n", stdOut)
	fmt.Printf("Stderr: %s\n", stdErr)
}
