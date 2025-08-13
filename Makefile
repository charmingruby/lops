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
apply-manifests: apply-pre-manifests
	@echo "applying application..."
	helm template api deploy/apps/api -f deploy/apps/api/values.yaml > deploy/apps/api/resources.yaml
	kubectl apply -k ./deploy
	@echo "application manifests applied successfully"

apply-pre-manifests: apply-metrics-server apply-argocd

apply-metrics-server:
	@echo "applying metrics-server manifests..."
	kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml && \
	kubectl patch deploy metrics-server -n kube-system \
	  --type=json \
	  -p='[{"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--kubelet-insecure-tls"}, {"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--kubelet-preferred-address-types=InternalIP"}]'
	@echo "metrics-server manifests applied"

apply-argocd:
	@echo "applying argocd pre-requisites manifests..."
	kubectl create namespace argocd && \
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
	@echo "argocd pre-requisites manifests applied"

argo-pass:
	argocd admin initial-password -n argocd