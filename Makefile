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
	@echo "applying Kustomize manifests..."
	kubectl apply -k ./deploy
	@echo "application manifests applied successfully"

.PHONY: apply-pre-manifests
apply-pre-manifests: apply-metrics-server apply-argocd

.PHONY: apply-metrics-server
apply-metrics-server:
	@kubectl get deploy metrics-server -n kube-system >/dev/null 2>&1 && \
		(echo "metrics-server already installed, skipping"; exit 0) || \
		(echo "applying metrics-server manifests..." && \
		kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml && \
		kubectl patch deploy metrics-server -n kube-system \
			--type=json \
			-p='[{"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--kubelet-insecure-tls"}, {"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--kubelet-preferred-address-types=InternalIP"}]' && \
		echo "metrics-server manifests applied")

.PHONY: apply-argocd
apply-argocd:
	@kubectl get ns argocd >/dev/null 2>&1 && \
		(echo "argocd already installed, skipping"; exit 0) || \
		(echo "applying argocd manifests..." && \
		kubectl create namespace argocd && \
		kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml && \
		echo "argocd manifests applied")

.PHONY: argo-pass
argo-pass:
	@kubectl get ns argocd >/dev/null 2>&1 && \
		argocd admin initial-password -n argocd || \
		echo "argocd not installed in cluster"