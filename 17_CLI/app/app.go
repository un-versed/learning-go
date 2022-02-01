package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns the CLI to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App Learning"
	app.Usage = "Search IP's and name servers on the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search Ip's",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "ns",
			Usage:  "Search name servers",
			Flags:  flags,
			Action: searchNameServers,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchNameServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
