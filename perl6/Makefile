IMAGE := rakudo-star-cgo-test

all: docker-build

docker-build:
	docker build -t $(IMAGE) .

docker-run:
	docker run --rm -it \
		-v $$(pwd)/..:/app:ro \
		-w /app \
		$(IMAGE) \
			bash

test:
	docker run --rm \
		-v $$(pwd)/..:/app:ro \
		-w /app \
		$(IMAGE) \
			prove -e perl6 -r -v perl6/t

test-valgrind:
	docker run --rm \
		-v $$(pwd)/..:/app:ro \
		-w /app \
		$(IMAGE) \
			prove -e perl6-valgrind-m -r -v perl6/t

.PHONY: all docker-build docker-run test test-valgrind
