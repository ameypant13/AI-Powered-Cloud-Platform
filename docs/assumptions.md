#  Assumptions Made and Their Justifications

## Overview

Certain architectural and technical assumptions underpin the platform’s design and prototype. These assumptions are made explicit to clarify boundaries, avoid over-engineering, and ensure a pragmatic, feasible solution in the MVP and initial rollout.

---

## Assumptions

### **A. Data Availability and Quality**

- **Assumption:**  
  Accurate, granular metrics (CPU, memory, etc.) for all workloads are consistently collected from Prometheus or CloudWatch.
- **Justification:**  
  Right-sizing and pattern recognition fundamentally rely on actual usage metrics. Kubernetes makes Prometheus the de facto standard for metrics, and CloudWatch is available across all AWS services. Outage or data gaps are captured as errors.

---

### **B. Homogeneity of Workloads (in MVP Phase)**

- **Assumption:**  
  Most workloads are homogeneous (i.e., similar in their scheduling, scaling, and SLO model).
- **Justification:**  
  The first implementation focuses optimization on generic resource types (CPU, memory, storage), not on specialized configurations or exotic runtime environments. Heterogeneous/Specialty workloads (e.g., GPUs, FPGA, burstable instances) can be handled in later iterations.

---

### **C. Multi-Tenancy Boundaries**

- **Assumption:**  
  Tenants are logically partitioned at the namespace/account level, and each pod/service maps clearly to a single tenant.
- **Justification:**  
  Simplifies tagging, permissions, and tenant-specific recommendations. Complex scenarios (shared workloads, federated clusters) are out of scope for the MVP.

---

### **D. Cost and Usage Attribution**

- **Assumption:**  
  AWS Cost & Usage Report (CUR) maps cleanly to tenant resources via resource tags or namespaces.
- **Justification:**  
  Key for cost savings to align with infra usage. In practice, this may require robust tagging policies or annotation enforcement.

---

### **E. Model Selection and Performance**

- **Assumption:**  
  Quantile-based and clustering methods are sufficient for actionable recommendations in the MVP.
- **Justification:**  
  These algorithms are industry-accepted, low-latency, explainable, and can be implemented natively in Go. More advanced models are justified only when clear incremental value is demonstrated.

---

### **F. Integration Patterns**

- **Assumption:**  
  Platform integrates with existing GitOps (ArgoCD) and CI/CD workflows.
- **Justification:**  
  Ensures minimal change management friction and speeds up the feedback loop from recommendation to adoption.

---

### **G. Security and Compliance**

- **Assumption:**  
  Strong tenant isolation is required for all storage and API surfaces.
- **Justification:**  
  Meets enterprise/cross-industry data protection standards and customer trust requirements.

---

### **H. MVP Focus and Incremental Delivery**

- **Assumption:**  
  Deep functionality in one component (e.g., right-sizing or data pipeline) delivers more value than shallow implementation across all features.
- **Justification:**  
  “Implement one key thing extremely well” aligns with agile/lean delivery and increases odds of early, demonstrable ROI.

---

## Summary

All assumptions are based on both industry best practices in multi-tenant SaaS platform engineering and the specific realities of operating at cloud-native scale on AWS/EKS.  
Each assumption is validated or made configurable for minimal refactoring or expansion as the platform matures.

---
