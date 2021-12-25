SERVICE := am-users-api
NAMESPACE := aftermath
REGISTRY := docker.io/vkouzin
# 
VERSION = $(shell git rev-parse --short HEAD)
TAG := ${REGISTRY}/${SERVICE}

echo:
	@echo "Tag:" ${TAG}

pull:
	git pull

build:
	 docker buildx build --push -t ${TAG}:${VERSION} -t ${TAG}:latest .

apply:
	kubectl apply -f ./_kube-yml

restart:
	kubectl rollout restart deployments/${SERVICE} -n ${NAMESPACE}

ctx:
	kubectl config set-context --current --namespace=${NAMESPACE}