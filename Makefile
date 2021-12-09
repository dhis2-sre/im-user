tag ?= latest
version ?= $(shell yq e '.version' helm/Chart.yaml)
clean-cmd = docker compose down --remove-orphans --volumes

binary:
	go build -o im-users -ldflags "-s -w" ./cmd/serve

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

dev-test: clean
	docker compose up -d database redis
	docker compose run --no-deps dev-test
	$(clean-cmd)

clean:
	$(clean-cmd)
	go clean

helm-chart:
	@helm package helm/chart

publish-helm:
	@curl --user "$(CHART_AUTH_USER):$(CHART_AUTH_PASS)" \
        -F "chart=@im-users-$(version).tgz" \
        https://helm-charts.fitfit.dk/api/charts

swagger-check-install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger-docs: swagger-check-install
	swagger generate spec -o ./swagger/swagger.yaml --scan-models

swagger-client: swagger-check-install
	swagger generate client -f ./swagger/swagger.yaml -t swagger/sdk

swagger: swagger-docs swagger-client

di:
	wire gen ./internal/di

.PHONY: binary docker-image push-docker-image dev test dev-test helm-chart publish-helm
