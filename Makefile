.PHONY: create-cluster
create-cluster:
	@echo "creating cluster..."
	kind create cluster --name lops-cluster --config ./k8s/clusters/dev/kind-config.yaml
	@echo "cluster created successfully"