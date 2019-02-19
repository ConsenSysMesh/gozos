package main

import (
	"fmt"
	"log"
)

func main() {

	var err error
	fmt.Println("starting gozos")

	zos := NewZos()

	projectName := "zos-app"

	err = zos.Init(
		projectName,
		InitVersion("0.8.7"),
	)
	if err != nil {
		log.Fatalf("initcmd error : %s", err.Error())
	}
}
