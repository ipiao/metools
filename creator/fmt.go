package creator

import (
	"os/exec"
)

// FmtFile is format
func FmtFile(fname string) {
	cmd := exec.Command("gofmt", "-w", fname)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
