package main

import (
	"fmt"
	"github.com/ameypant13/AI-Powered-Cloud-Platform/pkg/data-collector-pipeline/prometheus"
	"github.com/ameypant13/AI-Powered-Cloud-Platform/pkg/data-collector-pipeline/storage"
	"os"
	"time"
)

func main() {
	promEndpoint := os.Getenv("PROM_ENDPOINT")
	fmt.Println("Using Prometheus endpoint:", promEndpoint)
	prom, err := prometheus.NewPrometheusFetcher(promEndpoint)
	if err != nil {
		panic(err)
	}
	// Example query: K8s pod CPU usage
	query := `sum(rate(container_cpu_usage_seconds_total{container!=""}[5m])) by (pod,namespace)`
	start := time.Now().Add(-time.Hour)
	end := time.Now()
	step := time.Minute * 5

	metrics, err := prom.FetchRange(query, start, end, step)
	if err != nil {
		fmt.Println("Prometheus fetch error:", err)
		return
	}
	err = storage.SaveToFile("prom_metrics.json", metrics)
	if err != nil {
		fmt.Println("Failed to write metrics:", err)
	}
	fmt.Println("Fetched and saved Prometheus data.")

	// Repeat the above with CloudWatch fetcher, and then save results to file.

	// Schedule as a cron Job or run as a long-lived service.
}
