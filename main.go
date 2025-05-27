package main

import (
	"fmt"
	"github.com/ameypant13/AI-Powered-Cloud-Platform/pkg/pod-right-sizing"
	"github.com/gin-gonic/gin"
)

func main() {
	metrics := pod_right_sizing.PodMetrics{
		PodName:           "api-service-234",
		Namespace:         "customer1",
		CPUSamples:        []float64{80, 90, 120, 300, 120, 130, 80, 100}, // Example millicores
		MemSamples:        []float64{180, 200, 220, 300, 250, 190, 200},   // Example MiB
		CurrentCPURequest: 400,                                            // 400 millicores
		CurrentMemRequest: 512,                                            // 512 MiB
	}

	config := pod_right_sizing.RightSizerConfig{
		Quantile:    0.95, // Use p95 for request calculation
		ScaleFactor: 1.2,  // 20% headroom
		MinCPU:      50,   // 50 millicores
		MinMemory:   64,   // 64MiB
	}

	reco := pod_right_sizing.RecommendRightSize(metrics, config)
	fmt.Printf("Pod: %s/%s\n", reco.Namespace, reco.PodName)
	fmt.Printf("Recommended CPU: %.0fm, Memory: %.0fMiB\n", reco.RecommendedCPU, reco.RecommendedMemory)

	r := gin.Default()
	r.GET("/api/v1/recommendations", pod_right_sizing.GetRecommendations)

	// (Optional) healthcheck
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	r.Run(":8080") // Listen on port 8080
}
