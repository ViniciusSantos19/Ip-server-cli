package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Busca ips por linha de comando"
	app.Usage = "Busca Ips e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Buscar IPS de enderecos na internet",
			Flags:  flags,
			Action: buscarIps,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "servidores",
			Usage:  "Buscar servidor de enderecos na internet",
			Flags:  flags,
			Action: buscarServidor,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")
	servidores, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}

}
