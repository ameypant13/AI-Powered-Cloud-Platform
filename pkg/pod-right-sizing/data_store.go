package pod_right_sizing

var SamplePods = []PodMetrics{
	{
		PodName:           "api-service-1",
		Namespace:         "team1",
		CPUSamples:        []float64{110, 120, 100, 110, 105, 140, 135, 90},
		MemSamples:        []float64{200, 220, 230, 250, 205, 210, 202, 200},
		CurrentCPURequest: 250,
		CurrentMemRequest: 512,
	},
	{
		PodName:           "worker-47",
		Namespace:         "team2",
		CPUSamples:        []float64{400, 350, 420, 500, 480, 510, 390, 390},
		MemSamples:        []float64{1024, 1040, 1000, 1320, 1200, 1150, 980, 900},
		CurrentCPURequest: 800,
		CurrentMemRequest: 2048,
	},
}
