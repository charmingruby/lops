CLUSTER_NAME?=lops
ENV?=local

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
	kind create cluster --name $(CLUSTER_NAME) --config ./deploy/clusters/local/kind-config.yaml
	@echo "cluster deleted successfully"

.PHONY: apply-manifests
apply-manifests:
	@echo "applying manifests..."
	helm template base deploy/apps/base -f deploy/apps/lops/overlays/$(ENV)/values.yaml > deploy/apps/lops/overlays/$(ENV)/gen.yaml
	kubectl apply -k ./deploy
	@echo "manifests applied successfully"
