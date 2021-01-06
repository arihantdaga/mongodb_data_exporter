package config

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/bson"
)

type QueryConfig struct{
	AggStages interface{}
}

type GlobalConfig struct {
	Addr string `json:"addr"`
	MetricsPath string `json:"metricsPath"`

}

type MetricsConfig struct {
	Name string
	TypeString string //'Gauge/Counter' etc
	Help string
	KeyLabels []string // expose these columns as labels from Query
	Values	[]string
	ValueType	prometheus.ValueType // TypeString Converted to Prometheus Value Type
	QueryString string
	query *QueryConfig
}

type Config struct{
	Globals  *GlobalConfig
	Metrics []*MetricsConfig
}


func Load() (*Config, error){
	c := Config{
		Globals: &GlobalConfig{
			Addr: ":8080",
			MetricsPath: "/metrics",
		},
	}
	m := MetricsConfig{
		Name: "error_count",
		TypeString: "counter",
		Help: "Total Error Count",
		KeyLabels: []string{"instance", "status_code"},
		Values: []string{"count"},
		ValueType: prometheus.CounterValue,
		QueryString: `[
			{
				"$match": {
					"status": 500
				}
			},
			{
				"$project": {
					"instance": "$instance",
					"count": "$count",
					"status_code": "$status_code"
				}
			}
		]`,
	}
	m.query = &QueryConfig{}
	bson.UnmarshalExtJSON([]byte(m.QueryString), true, &m.query.AggStages)
	c.Metrics = []*MetricsConfig{&m}

	return &c, nil
}