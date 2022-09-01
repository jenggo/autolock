package app

import (
	"os"
	"os/exec"

	"autolock/vars"
)

func RunCommand() error {
	c := exec.Command("sh", "-c", vars.Config.LockCommand)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
