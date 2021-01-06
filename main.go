package main

import (
	"flag"
	"log"
	"net/http"

	collector "arihantdata/data_exporter/collector"
	config "arihantdata/data_exporter/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	metricsPath = flag.String("metrics-path", "/metrics", "http request path for metrics")
)
var conf *config.Config

func Init(){
	c, err := config.Load()
	if err != nil {
		panic("Could not load config")
	}
	conf = c;
}

func main(){
	flag.Parse()
	Init()
	prometheus.MustRegister(collector.NewCollector())
	http.Handle(*metricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(conf.Globals.Addr, nil))
}