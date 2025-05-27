package pod_right_sizing

type PodMetrics struct {
	PodName           string
	Namespace         string
	CPUSamples        []float64 // in millicores, e.g. 100m = 0.1 CPU
	MemSamples        []float64 // in MiB
	CurrentCPURequest float64   // in millicores
	CurrentMemRequest float64   // in MiB
}
