package util

import (
	"log"
	"os/exec"
	"runtime"
	"syscall"
)

func OpenWinBrowser(uri string) error {
	if runtime.GOOS != "windows" {
		return nil
	}

	cmd := exec.Command(`cmd`, `/c`, `start`, uri)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	return nil
}
