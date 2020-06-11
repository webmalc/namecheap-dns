//+build !test

package main

import (
	"github.com/webmalc/namecheap-dns/cmd"
	"github.com/webmalc/namecheap-dns/common/config"
	"github.com/webmalc/namecheap-dns/common/logger"
	"github.com/webmalc/namecheap-dns/namecheap"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	cmdRouter := cmd.NewCommandRouter(log, namecheap.NewChanger(log))
	cmdRouter.Run()
}
