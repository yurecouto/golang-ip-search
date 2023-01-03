package main

import (
	"fmt"
	"golang-ip-search/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start Point")

	application := app.Generate()
	erro := application.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}

}
