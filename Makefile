.PHONY: service lint test

all: web

web:
	go build -o bin/web cmd/web/main.go && chmod +x bin/web

fmt-fetch:
	go get github.com/daixiang0/gci
	go get mvdan.cc/gofumpt@v0.1.0

fmt: fmt-fetch
	go fmt ./pkg/...
	find ./pkg -type f -name "*.go" -not -path "./pkg/dao/models/*" -not -name "*_test.go"  -not -name "mock_*.go" | xargs gofumpt -l -w
	find ./pkg -type f -name "*.go" -not -path "./pkg/dao/models/*" -not -name "*_test.go"  -not -name "mock_*.go" | xargs gci -w

lint:
	$(GOPATH)/bin/golangci-lint run ./pkg/...

models:
	bash sqlboiler.env
	bash sqlboiler.sh
	make fmt

# 查看当前运行的 go 进程, kill -9 杀死后再执行 make run
ps_go:
	ps aux | grep go

run:
	nohup go run cmd/web/main.go &
	nohup go run cmd/cron/main.go &
