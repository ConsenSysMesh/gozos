package main

import "os/exec"

func call(name string, args ...string) error{
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	return err
}
