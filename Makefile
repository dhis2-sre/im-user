tag ?= latest
version ?= $(shell yq e '.version' helm/Chart.yaml)
clean-cmd = docker compose down --remove-orphans --volumes

keys:
	openssl genpkey -algorithm RSA -out ./rsa_private.pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in ./rsa_private.pem -pubout -out ./rsa_public.pem

init:
	direnv allow
	pip install pre-commit
	pre-commit install --install-hooks --overwrite

check:
	pre-commit run --all-files --show-diff-on-failure

smoke-test:
	docker compose up -d database redis
	sleep 3
	IMAGE_TAG=$(tag) docker compose up -d prod

docker-image:
	IMAGE_TAG=$(tag) docker compose build prod

push-docker-image:
	IMAGE_TAG=$(tag) docker compose push prod

dev:
	docker compose up --build dev database redis

test: clean
	docker compose up -d database redis
	docker compose run --no-deps test
	$(clean-cmd)

clean:
	$(clean-cmd)
	go clean

swagger-check-install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger-clean:
	rm -rf swagger/sdk/*
	rm swagger/swagger.yaml

swagger-docs: swagger-check-install
	swagger generate spec -o swagger/swagger.yaml -x swagger/sdk --scan-models
	swagger validate swagger/swagger.yaml

swagger-client: swagger-check-install
	swagger generate client -f swagger/swagger.yaml -t swagger/sdk

swagger: swagger-clean swagger-docs swagger-client

.PHONY: check docker-image init push-docker-image dev test
