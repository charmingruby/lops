CLUSTER_NAME?=lops

.PHONY: setup-cluster
setup-cluster: create-cluster apply-manifests
	
.PHONY: teardown-cluster
teardown-cluster:
	@echo "deleting cluster..."
	kind delete cluster --name $(CLUSTER_NAME)
	@echo "cluster deleted successfully"

.PHONY: create-cluster
create-cluster:
	@echo "creating cluster..."
	kind create cluster --name $(CLUSTER_NAME) --config ./deploy/clusters/kind-config.yaml
	@echo "cluster created successfully"

.PHONY: apply-manifests
apply-manifests:
	@echo "applying manifests..."
	helm template base deploy/apps/base -f deploy/apps/lops/values.yaml > deploy/apps/lops/gen.yaml
	kubectl apply -k ./deploy
	@echo "manifests applied successfully"
