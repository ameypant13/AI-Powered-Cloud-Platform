package clustering

import (
	"github.com/mash/gokmeans"
	"log"
)

const numClusters = 3 // e.g., cluster workloads into 3 groups

func Clustering(featureVectors [][]float64) []int {
	// Convert [][]float64 to []gokmeans.Node
	var nodes []gokmeans.Node
	for _, fv := range featureVectors {
		nodes = append(nodes, fv)
	}

	// Use Train function and ignore the clusters variable since it's not used
	success, centroids := gokmeans.Train(nodes, numClusters, 50) // 50 iterations
	if !success {
		log.Fatalf("Clustering failed")
	}

	// Get the cluster assignments for each data point
	assignments := make([]int, len(featureVectors))
	for i, v := range featureVectors {
		assignments[i] = gokmeans.Nearest(gokmeans.Node(v), centroids)
	}

	return assignments
}
