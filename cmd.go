package main

const(
	zos = "zos"
	initCmd = "init"
)


type Cmd interface {
	Exec() error
}


