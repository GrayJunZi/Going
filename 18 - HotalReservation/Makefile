build:
	go build -o bin/api

run: build
	./bin/api

seed:
	@go run scripts/seed.go

test:
# 使用 @ 符号 隐藏输出
	@go test -v ./...