package main

import "github.com/prometheus/client_golang/prometheus"
var (

	buildInfo = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "skeleton_app_build_info",
		Help:        "Skeleton App build information.",
		ConstLabels: prometheus.Labels{"revision": GitCommit, "buildDate": BuildDate},
	})
)

func initialiseMetrics() {
	prometheus.MustRegister(buildInfo)

	buildInfo.Set(1)
}
