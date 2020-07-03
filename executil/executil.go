package executil

import "os/exec"

func Run(cmd string, args ...string) *exec.Cmd {
	var cmdRes *exec.Cmd = exec.Command(cmd, args...)
	cmdRes.Run()
	return cmdRes
}
