package cmd

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

func Run(name string, arg ...string) error {

	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	process := bufio.NewReader(stdout)

	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	for {
		line, _, err := process.ReadLine()

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
				return err
			}
			break
		}

		log.Printf("%s: %s\n", name, line)
	}

	return cmd.Wait()
}
