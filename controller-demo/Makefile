IMAGE=ghcr.io/nycdavid/contrived-controller:latest

PHONY: image
image:
	@docker build -t "${IMAGE}" .

PHONY: push
push: image
	@docker push "${IMAGE}"
