version = v0.0.2

build:
	go build -ldflags "-X github.com/ray846/fops/cmd.Version=$(version)"

tests:
	go test ./...