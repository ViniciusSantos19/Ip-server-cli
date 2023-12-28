package main

import (
	"fmt"
	"lc/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partido da aplicação")
	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
