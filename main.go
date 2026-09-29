// Repro for blackjack/webcam#74. Linux + V4L2 device required.
// go.mod replaces in the patched fork.
//
//	go run . -dev /dev/video0        # Open: video fd visible in child
//	go run . -dev /dev/video0 -safe  # OpenCloexec: not visible
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/blackjack/webcam"
)

func main() {
	dev := flag.String("dev", "/dev/video0", "video device")
	safe := flag.Bool("safe", false, "use OpenCloexec")
	flag.Parse()

	open := webcam.Open
	if *safe {
		open = webcam.OpenCloexec
	}
	cam, err := open(*dev)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer cam.Close()

	fmt.Printf("parent pid=%d dev=%s safe=%v\n", os.Getpid(), *dev, *safe)

	cmd := exec.Command("/bin/sh", "-c", "export PATH=/usr/bin:/bin; ls -l /proc/self/fd/ | grep -E 'video|v4l' || echo '(no video fd)'")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "child:", err)
	}
}
