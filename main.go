//+build !test

package main

import (
	"github.com/webmalc/namecheap-dns/client"
	"github.com/webmalc/namecheap-dns/cmd"
	"github.com/webmalc/namecheap-dns/common/config"
	"github.com/webmalc/namecheap-dns/common/logger"
	"github.com/webmalc/namecheap-dns/ip"
	"github.com/webmalc/namecheap-dns/namecheap"
	"github.com/webmalc/namecheap-dns/server"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	changer := namecheap.NewChanger(log)
	cmdRouter := cmd.NewCommandRouter(
		log,
		changer,
		server.NewServer(log, changer),
		client.NewClient(log, ip.NewGetter(log)),
	)
	cmdRouter.Run()
}
