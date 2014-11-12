package main

import (
	// this import is needed for the HN plugin.  the ssl cert for HN uses
	// an algorithm that isn't available by default (sha384)
	// https://groups.google.com/forum/#!topic/Golang-nuts/hqm_ssUNUtI for details
	_ "crypto/sha512"
	"github.com/Sirupsen/logrus"
	"github.com/ehazlett/phoenix/plugins"

	"os"
	"strings"
)

var (
	pluginList []string
	listenAddr string
	logger     = logrus.New()
	version    = "0.1"
)

func main() {
	logger.Infof("phoenix v%s", version)
	pluginList = strings.Split(os.Getenv("PLUGINS"), ",")
	logger.WithFields(logrus.Fields{
		"names": pluginList,
	}).Info("enabling plugins")
	manager := plugins.New(pluginList)
	server := NewServer(manager, ":"+os.Getenv("PORT"))
	server.Run()
}
