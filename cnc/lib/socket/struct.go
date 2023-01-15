package socket

import "net"

var (
	MasterList []*Master
	BotList    []*Bot
)

type Socket struct {
	Connected bool
	Socket    net.Conn
}

type Master struct {
	Network   *Socket
	Logged    bool
	LogAttemp int
}

type Bot struct {
	Cpu     int
	Mem     int
	Disk    int
	Auth    bool
	Arch    string
	Version string
	Vector  string
	Network *Socket
}
