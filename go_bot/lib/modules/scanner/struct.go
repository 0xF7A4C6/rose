package scanner

import "io"

var BlacklistIPs = map[string]struct{}{

	/*
		Loopback
	*/
	"127.": {},

	/*
		General Electric Company
	*/
	"3.": {},

	/*
		Hewlett-Packard Company
	*/
	"15.": {}, "16.": {},

	/*
		US Postal Service
	*/
	"56.": {},

	/*
		Internal network
	*/
	"10.": {}, "192.": {}, "172.": {},

	/*
		NAT
	*/
	"100.": {}, "169.": {},

	/*
		Special use
	*/
	"198.": {},

	/*
		Multicast
	*/
	"224.": {},

	/*
		CIA
	*/
	"162.": {},

	/*
		Cloudflare
	*/
	"104.": {},

	/*
		NASA Kennedy Space Center
	*/
	"163.": {}, "164.": {},

	/*
		Naval Air Systems Command, VA
	*/
	"199.": {},

	/*
		Department of the Navy, Space and Naval Warfare System Command, Washington DC - SPAWAR
	*/
	"205.": {},

	/*
		FBI controlled Linux servers & IPs/IP-Ranges
	*/
	"207.": {},

	/*
		Amazon + Microsoft
	*/
	"13.": {}, "52.": {}, "54.": {},

	/*
		Ministry of Education Computer Science
	*/
	"120.": {}, "188.": {}, "78.": {},

	/*
		Total Server Solutions
	*/
	"107.": {}, "184.": {}, "206.": {}, "98.": {},

	/*
		Blazingfast & Nforce
	*/
	"63.": {}, "70.": {}, "74.": {},

	/*
		Choopa & Vultr
	*/
	"64.": {}, "185.": {}, "208.": {}, "209.": {}, "45.": {}, "66.": {}, "108.": {}, "216.": {},

	/*
		OVH
	*/
	"149.": {}, "151.": {}, "167.": {}, "176.": {}, "178.": {}, "37.": {}, "46.": {}, "51.": {},
	"5.": {}, "91.": {},

	/*
		Department of Defense
	*/
	"6.": {}, "7.": {}, "11.": {}, "21.": {}, "22.": {}, "26.": {}, "28.": {}, "29.": {},
	"30.": {}, "33.": {}, "55.": {}, "214.": {}, "215.": {},
}

var (
	EvilProcess = []string{
		"i",
		".i",
		"mozi.m",
		"Mozi.m",
		"mozi.a",
		"Mozi.a",
		"kaiten",
		"Nbrute",
		"minerd",
		"/bin/busybox",
	}
)

type ScannerThread struct {
	Found    int
	ThreadID int
	Worker   int
}

type Exploit struct {
	ExploitBody       io.Reader
	ExploitName       string
	ExploitMethod     string
	ExploitHeader     string
	ExploitAgent      string
	ExploitAccept     string
	ExploitContType   string
	ExploitConnection string
}
