package main

import (
	"vline/service"
	"vline/util"
)

func main() {

	log := util.InitLogger()
	defer log.Sync()

	// parse config
	// init datasource and handle data

	go service.MssqlHandle(log)

	select {}
}
