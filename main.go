package main

import (
	"fmt"
	"log"
)

func main()  {

	var err error
	fmt.Println("starting gozos")
	zos := NewZos()

	projectName := "zos-app"

	err = zos.Init(projectName)
	if err != nil{
		log.Fatalf("initcmd error : %s", err.Error())
	}
}


