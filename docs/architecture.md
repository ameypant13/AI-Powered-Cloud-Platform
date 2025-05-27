# AI-Powered Cloud Platform Architecture

## System Overview

The AI-Powered Cloud Platform provides intelligent cloud resource optimization across multiple tenants through advanced metrics collection, analysis, and machine learning-based recommendations. The architecture follows a multi-layered approach designed for scalability and tenant isolation.

## High-Level Architecture

![High-Level Architecture](docs/images/Atlan_AI.drawio.png)

The system consists of the following major components:

### 1. Data Sources Layer
- **Multi-tenant Kubernetes Clusters**: Each tenant operates on a dedicated Amazon EKS cluster
- **Monitoring Systems**:
    - Prometheus exporters for container and pod metrics
    - Kube State Metrics for Kubernetes object information
    - Custom metric collectors for application-specific data
    - AWS CloudWatch for infrastructure metrics
    - AWS Cost & Usage Reports for cost analysis

### 2. Data Collection Pipeline

![Data Pipeline](docs/images/Atlan_data.drawio.png)

- **Data Collector Service**:
    - REST API endpoints for metrics ingestion
    - Polling mechanisms for CloudWatch metrics
    - Integration with Prometheus Query API
    - Data normalization and preprocessing

- **Data Aggregation**:
    - Combines metrics from all tenants
    - Standardizes formats for analysis
    - Securely stores time-series data in Amazon S3

### 3. AI/ML Analysis Layer
- **Feature Extraction**:
    - Transforms raw metrics into meaningful workload features:
        - Mean CPU utilization
        - CPU burstiness (standard deviation/mean)
        - Resource efficiency ratios
        - Memory utilization patterns

- **K-means Clustering Engine**:
    - Groups similar workloads based on resource usage patterns
    - Identifies common behaviors across pods/containers
    - Enables more tailored recommendations per cluster

- **Pattern Recognition**:
    - Identifies resource waste and optimization opportunities
    - Detects anomalies in resource usage
    - Powers the recommendation engine

### 4. Recommendation & Delivery Layer
- **Right-Sizing Recommendation Engine**:
    - Generates specific CPU and memory recommendations
    - Configurable parameters (quantiles, scale factors, minimums)
    - Calculates potential cost savings

- **API Gateway**:
    - Secure access to recommendations
    - Tenant-specific filtering
    - REST endpoints for integration

- **Recommendations Dashboard**:
    - Visualization of optimization opportunities
    - Savings estimates
    - Implementation suggestions

## Technical Implementation

The system is implemented using Go (Golang) with the following key components:

- **Data Collection**: Prometheus API client for metrics collection
- **Statistical Analysis**: Gonum for quantile calculations and statistical processing
- **Clustering Algorithm**: K-means implementation for workload pattern identification
- **API Layer**: Gin web framework for REST endpoints
- **Persistence**: JSON storage with capability to extend to database solutions

## Data Flow

1. Metrics are collected from Kubernetes clusters via Prometheus and CloudWatch
2. Data collector services normalize and aggregate metrics
3. Raw data is stored in Amazon S3
4. Feature extraction processes transform raw metrics into workload characteristics
5. K-means clustering groups similar workloads
6. Recommendation engine generates right-sizing suggestions
7. API Gateway exposes recommendations to clients
8. Dashboard visualizes optimization opportunities

## Security Considerations

- Tenant isolation at the Kubernetes cluster level
- Secure API access with authentication and authorization
- Data segregation throughout the pipeline
- AWS IAM roles for service access control

## Future Enhancements

- Machine learning models for predictive scaling
- Additional data sources integration
- Historical trend analysis
- Automated recommendation implementation via Kubernetes Operators