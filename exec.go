package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func call(dir string, commandName string, args ...string) error {
	fmt.Printf("$ %s %v\n", commandName, args)
	cmd := exec.Command(commandName, args...)
	cmd.Dir = dir
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return err
	}
	if outb.Len() > 0 {
		fmt.Printf("%s\n", outb.String())
	}
	if errb.Len() > 0 {
		fmt.Printf("%s\n", errb.String())
	}
	return nil

}
