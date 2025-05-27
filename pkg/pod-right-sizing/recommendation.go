package pod_right_sizing

import (
	"github.com/ameypant13/AI-Powered-Cloud-Platform/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRecommendations(c *gin.Context) {
	namespace := c.Query("namespace")

	// Normally, this would filter per tenant/namespace and fetch live metrics
	if namespace == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "namespace required"})
		return
	}

	config := RightSizerConfig{
		Quantile:    0.95,
		ScaleFactor: 1.2,
		MinCPU:      50,
		MinMemory:   64,
	}

	var recos []RightSizingRecommendation
	for _, pod := range pkg.samplePods {
		if strings.EqualFold(pod.Namespace, namespace) {
			reco := RecommendRightSize(pod, config)
			recos = append(recos, reco)
		}
	}
	c.IndentedJSON(http.StatusOK, recos)
}
