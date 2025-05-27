package clustering

import pod_right_sizing "github.com/ameypant13/AI-Powered-Cloud-Platform/pkg/pod-right-sizing"

func vectorizeFeatures(samplePods []pod_right_sizing.PodMetrics) [][]float64 {
	// Prepare data matrix for clustering: [][]float64
	var featureVectors [][]float64
	var metaList []WorkloadFeatures
	for _, pod := range samplePods {
		feat := ExtractFeatures(pod)
		metaList = append(metaList, feat)
		featureVectors = append(featureVectors, []float64{
			feat.MeanCPU,
			feat.CPUBurstiness,
			feat.CPURatio,
			feat.MeanMem,
		})
	}
	return featureVectors
}
