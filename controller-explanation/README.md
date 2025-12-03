This is the companion code for the Debugging K8s controllers blog post

# Quickstart

1. Apply the CRDs to cluster you're testing in:

```bash
kubectl apply \
  -f controller-explanation/crds/car.yaml \
  -f controller-explanation/crds/wheel.yaml
```
