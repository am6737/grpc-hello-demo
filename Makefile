.PHONY: build-server build-client run-server run-client build-push push

REGISTRY=hub.hitosea.com
REPOSITORY=csiaddons
TAG=grpc-demo

CSI_IMAGE_NAME=$(REGISTRY)/$(REPOSITORY)
SERVER_IMAGE=$(CSI_IMAGE_NAME)/$(TAG)-server
CLIENT_IMAGE=$(CSI_IMAGE_NAME)/$(TAG)-client

# Build server Docker image


build-server:
	docker build -t $(SERVER_IMAGE) --build-arg TARGET=server .

# Build client Docker image
build-client:
	docker build -t $(CLIENT_IMAGE) --build-arg TARGET=client .

# Run server container
run-server:
	docker run --rm $(SERVER_IMAGE)

# Run client container
run-client:
	docker run --rm $(CLIENT_IMAGE)

# Build and push Docker images
build-push: build-server build-client
	docker push $(SERVER_IMAGE)
	docker push $(CLIENT_IMAGE)

push:
	docker push $(SERVER_IMAGE)
	docker push $(CLIENT_IMAGE)