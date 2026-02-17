#!make
include .vars

default: version

clean:
	if [ "$$(docker images "$${CONTAINER_ORG}/$${CONTAINER_IMAGE}" --format "{{.Repository}}:{{.Tag}}")" != "" ]; then \
		docker image rm $$(docker images "$${CONTAINER_ORG}/$${CONTAINER_IMAGE}" --all --format "{{.Repository}}:{{.Tag}}"); \
	fi

docker: clean
	docker build \
	--build-arg GO_BUILDER=$${GO_BUILDER} \
	--build-arg GO_VERSION=$${GO_VERSION} \
	--build-arg GO_OS=$${GO_OS} \
	--build-arg GO_ARCH=$${GO_ARCH} \
	--build-arg GIT_HOST=$${GIT_HOST} \
	--build-arg REPO_ORG=$${REPO_ORG} \
	--build-arg REPO_NAME=$${REPO_NAME} \
	--build-arg APP_VERSION=$${APP_VERSION} \
	-t $${CONTAINER_ORG}/$${CONTAINER_IMAGE}:$${APP_VERSION} .; \
	if [ "$$(docker images --filter "dangling=true" --quiet --no-trunc)" != "" ]; then \
		docker image rm $$(docker images --filter "dangling=true" --quiet --no-trunc); \
	fi

dockerpush:
	docker push \
	$${CONTAINER_ORG}/$${CONTAINER_IMAGE}:$${APP_VERSION}

deploy:
	bash -c "./deploy.sh"

test:
	go test ./...

version:
	bash -c "./version.sh"

run:
	docker run \
	--rm \
	--tty \
	--interactive \
	--publish $${CONTAINER_IP}:$${CONTAINER_PORT}:80 \
	--workdir /go/src/$${GIT_HOST}/$${REPO_ORG}/$${REPO_NAME} \
	--volume $(shell pwd):/go/src/$${GIT_HOST}/$${REPO_ORG}/$${REPO_NAME} \
	$${GO_BUILDER}:$${GO_VERSION} \
	go run .

dockerrun:
	docker run \
	--rm \
	--tty \
	--interactive \
	--publish $${CONTAINER_IP}:$${CONTAINER_PORT}:80 \
	--volume $(shell pwd)/config.yml:/config.yml \
	$${CONTAINER_ORG}/$${CONTAINER_IMAGE}:latest
