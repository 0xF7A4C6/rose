package grafana

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var BotCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bot_count",
		Help: "Number of bots",
	},
	[]string{"type", "arch"},
)

var CpuCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "cpu_count",
		Help: "Number of cores",
	},
	[]string{"cpu_core"},
)

var MethodCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "method_count",
		Help: "Number of method",
	},
	[]string{"method_name"},
)

var DownloadCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "download_count",
		Help: "Number of download",
	},
	[]string{"arch"},
)

var ReqCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "req_count",
		Help: "Number of download",
	},
	[]string{"endpoint"},
)

var HoneypotCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "honeypot_count",
		Help: "Number of download",
	},
	[]string{"match"},
)

func StartPrometheus() {
	prometheus.MustRegister(BotCount)
	prometheus.MustRegister(CpuCount)
	prometheus.MustRegister(ReqCount)
	prometheus.MustRegister(MethodCount)
	prometheus.MustRegister(DownloadCount)
	prometheus.MustRegister(HoneypotCount)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}