.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: build
build:
	go build -o bot cmd/bot/main.go

tests:
	go test -timeout 30s github.com/ozonmp/omp-bot/internal/service/logistic/group

