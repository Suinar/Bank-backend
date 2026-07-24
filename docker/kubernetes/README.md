# Kubernetes

These manifests follow the conventions used by the repository and exchange-rate
services: namespace-scoped resources, standard labels, dedicated service
accounts, probes, resource limits, restricted security contexts, HPA and PDB.

This project owns the `bank` namespace and the shared three-node Kafka KRaft
cluster. Repository and exchange-rate services connect to `kafka:9092` and must
not deploy separate Kafka clusters.

## Local deployment

Build all local images:

```powershell
docker build -f docker/app/Dockerfile -t bank-backend:local .
docker build -f ..\Bank-repository-service\docker\app\Dockerfile -t bank-repository-service:local ..\Bank-repository-service
docker build -f ..\Bank-exhange-rate-service\docker\app\Dockerfile -t bank-exchange-rate-service:local ..\Bank-exhange-rate-service
```

For a cluster that cannot access Docker's local images (for example `kind`),
load the images using that cluster's image-loading command.

Apply shared infrastructure first, followed by both dependent services:

```powershell
kubectl apply -f docker/kubernetes
kubectl apply -f ..\Bank-repository-service\docker\kubernetes
kubectl apply -f ..\Bank-exhange-rate-service\docker\kubernetes
```

Verify the rollout:

```powershell
kubectl get pods,services -n bank
kubectl rollout status statefulset/kafka -n bank
kubectl rollout status deployment/bank-backend -n bank
```

Expose the backend locally:

```powershell
kubectl port-forward -n bank service/bank-backend 18080:8080
```

The API is available at `http://localhost:18080`.

## Production notes

- Replace `:local` images with immutable registry tags.
- Select a production StorageClass and size Kafka volumes for retention needs.
- Keep credentials in Kubernetes Secrets or an external secret manager.
- Add NetworkPolicies and Kafka TLS/SASL before exposing the cluster externally.
