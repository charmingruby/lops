CLUSTER_NAME:=lops

.PHONY: setup-cluster
setup-cluster:
	@echo "creating cluster..."
	kind create cluster --name $(CLUSTER_NAME) --config ./deploy/clusters/dev/kind-config.yaml
	@echo "cluster created successfully"

.PHONY: teardown-cluster
teardown-cluster:
	@echo "deleting cluster..."
	kind delete cluster --name $(CLUSTER_NAME)
	@echo "cluster deleted successfully"
