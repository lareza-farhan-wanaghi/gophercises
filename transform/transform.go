package transform

import (
	"log"
	"os/exec"
	"strconv"
)

func Primitive(in, out string, mode, n int) error {
	path, err := exec.LookPath("primitive")
	if err != nil {
		return err
	}
	cmd := exec.Command(path, "-i", in, "-o", out, "-m", strconv.Itoa(mode), "-n", strconv.Itoa(n))
	if cmd.Err != nil {
		return cmd.Err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}
	if cmd.Err != nil {
		return cmd.Err
	}

	log.Printf("created primitive transformed image on %s", out)
	return nil
}
