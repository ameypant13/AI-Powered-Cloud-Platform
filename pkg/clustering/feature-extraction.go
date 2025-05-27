package clustering

import (
	pod_right_sizing "github.com/ameypant13/AI-Powered-Cloud-Platform/pkg/pod-right-sizing"
	"gonum.org/v1/gonum/stat"
)

type WorkloadFeatures struct {
	PodName       string
	Namespace     string
	MeanCPU       float64
	CPUBurstiness float64 // = stddev/mean
	CPURatio      float64 // = meanCPU / currentCPURequest
	MeanMem       float64
}

func ExtractFeatures(pod pod_right_sizing.PodMetrics) WorkloadFeatures {
	meanCPU := stat.Mean(pod.CPUSamples, nil)
	stdCPU := stat.StdDev(pod.CPUSamples, nil)
	burstiness := 0.0
	if meanCPU > 0 {
		burstiness = stdCPU / meanCPU
	}

	return WorkloadFeatures{
		PodName:       pod.PodName,
		Namespace:     pod.Namespace,
		MeanCPU:       meanCPU,
		CPUBurstiness: burstiness,
		CPURatio:      meanCPU / pod.CurrentCPURequest,
		MeanMem:       stat.Mean(pod.MemSamples, nil),
	}
}
