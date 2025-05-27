package pod_right_sizing

// go.mod will need "gonum.org/v1/gonum/stat"
import (
	"gonum.org/v1/gonum/stat"
	"sort"
)

func GetQuantile(data []float64, quantile float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	clone := append([]float64{}, data...)
	sort.Float64s(clone)
	return stat.Quantile(quantile, stat.Empirical, clone, nil)
}
