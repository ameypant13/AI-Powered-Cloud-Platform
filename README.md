# AI-Powered Cloud Platform

A prototype intelligent cloud resource optimization platform that analyzes workload patterns and provides automated right-sizing recommendations.

## Overview

This project demonstrates an AI-driven approach to cloud resource optimization by:

1. Collecting metrics from monitoring systems (Prometheus)
2. Analyzing workload patterns using statistical methods
3. Clustering similar workloads using K-means
4. Generating resource optimization recommendations

## Key Features

- **Metrics Collection**: Fetches time-series metrics from Prometheus
- **Workload Analysis**: Statistical analysis of CPU and memory usage patterns
- **Pod Right-Sizing**: Recommends optimal CPU and memory settings with configurable parameters
- **Workload Clustering**: Groups similar workloads using K-means algorithm
- **REST API**: Exposes recommendations through a simple API endpoint

## Architecture

```
┌─────────────────┐     ┌───────────────┐     ┌────────────────┐
│  Data Collector │────▶│ Feature       │────▶│ Clustering     │
│  (Prometheus)   │     │ Extraction    │     │ (K-means)      │
└─────────────────┘     └───────────────┘     └────────────────┘
                                                      │
┌─────────────────┐     ┌───────────────┐            ▼
│  REST API       │◀────│ Right-Sizing  │◀────┌────────────────┐
│  (Gin)          │     │ Recommendations│     │ Analysis       │
└─────────────────┘     └───────────────┘     └────────────────┘
```

## Getting Started

1. Set up a Prometheus instance with Kubernetes metrics
2. Set the environment variable:
   ```
   export PROM_ENDPOINT=http://your-prometheus:9090
   ```
3. Build and run the API server:
   ```
   go build -o api-server ./cmd/api-server ./api-server
   ```
4. In a separate terminal, build and run the data collector
    ```
    go build -o data-collector ./cmd/data-collector ./data-collector
    ```

## API Usage

Get recommendations:
```
GET /api/v1/recommendations?namespace=team1
```

Example using curl:
```bash
curl -X GET "http://localhost:8080/api/v1/recommendations?namespace=team1"
```

Health check:
```
GET /healthz
```

## Future Enhancements

- Machine learning models for predictive recommendations
- Support for additional metric sources (CloudWatch, etc.)
- Historical analysis for detecting workload patterns
- Integration with Kubernetes for automated implementation

## Technology Stack

- Go (Golang)
- Prometheus API
- K-means clustering
- Gin web framework
- Gonum for statistical analysis