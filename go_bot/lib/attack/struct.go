package attack

type HeaderFlags struct {
	synFlag bool
	ackFlag bool
	rstFlag bool
	pshFlag bool
	finFlag bool
	urgFlag bool
}

type Method struct {
	Name    string
	Payload []byte
	Flags   HeaderFlags
	Layer   int
}

type Attack struct {
	DestAddr string
	DestPort int
	Threads  int
	Running  bool
	Power    int
	Time     int
	Name     string
	Method   Method
}
