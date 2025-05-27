package pod_right_sizing

import (
	"fmt"
)

type RightSizingRecommendation struct {
	PodName           string
	Namespace         string
	RecommendedCPU    float64
	RecommendedMemory float64
	CurrentCPU        float64
	CurrentMemory     float64
	Confidence        float64
	Rationale         string
	PotentialSavings  string
}

type RightSizerConfig struct {
	Quantile    float64 // e.g., 0.95
	ScaleFactor float64 // e.g., 1.2 (buffer for p95 usage → recommended request)
	MinCPU      float64 // minimum CPU request (to avoid undersizing)
	MinMemory   float64 // minimum memory request
}

func RecommendRightSize(metrics PodMetrics, config RightSizerConfig) RightSizingRecommendation {
	// Calculate p95, p99, mean, etc.
	cpuQuantile := GetQuantile(metrics.CPUSamples, config.Quantile)
	memQuantile := GetQuantile(metrics.MemSamples, config.Quantile)

	recoCPU := cpuQuantile * config.ScaleFactor
	recoMem := memQuantile * config.ScaleFactor

	// Enforce resource floor
	if recoCPU < config.MinCPU {
		recoCPU = config.MinCPU
	}
	if recoMem < config.MinMemory {
		recoMem = config.MinMemory
	}

	confidence := 0.95 // could be improved with historical error measurement

	// Estimate potential savings compared to current settings
	cpuSavings := ""
	memSavings := ""
	if recoCPU < metrics.CurrentCPURequest {
		pct := 100 * (metrics.CurrentCPURequest - recoCPU) / metrics.CurrentCPURequest
		cpuSavings = fmt.Sprintf("Potential %.1f%% CPU reduction", pct)
	}
	if recoMem < metrics.CurrentMemRequest {
		pct := 100 * (metrics.CurrentMemRequest - recoMem) / metrics.CurrentMemRequest
		memSavings = fmt.Sprintf("Potential %.1f%% Mem reduction", pct)
	}

	return RightSizingRecommendation{
		PodName:           metrics.PodName,
		Namespace:         metrics.Namespace,
		RecommendedCPU:    recoCPU,
		RecommendedMemory: recoMem,
		CurrentCPU:        metrics.CurrentCPURequest,
		CurrentMemory:     metrics.CurrentMemRequest,
		Confidence:        confidence,
		Rationale:         fmt.Sprintf("Recommended %.0fth percentile x scale factor. (%s, %s)", config.Quantile*100, cpuSavings, memSavings),
		PotentialSavings:  fmt.Sprintf("%s, %s", cpuSavings, memSavings),
	}
}
