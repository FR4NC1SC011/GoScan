package cmd

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func Cmd(ports []int, host string) {
	nmap, err := exec.LookPath("/usr/bin/nmap")
	if err != nil {
		log.Fatal(err)
	}

	cmd := Command(ports)
	cmd = append(cmd, host)
	//args := []string{cmd, host}
	//color.Red("Command -> %#v", args)
	env := os.Environ()

	execErr := syscall.Exec(nmap, cmd, env)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
