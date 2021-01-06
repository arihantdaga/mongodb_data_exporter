package collector

import "github.com/prometheus/client_golang/prometheus"


type Collector struct{
	up *prometheus.Desc
}

func NewCollector() *Collector{
	return &Collector{
		up: prometheus.NewDesc("eventstore_up", "Whether the EventStore scrape was successful", nil, nil),
	}
}

func (c* Collector) Describe(ch chan<- *prometheus.Desc){
	ch <- c.up
}

func (c* Collector) Collect(ch chan<- prometheus.Metric){
	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
}
