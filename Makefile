version = v0.0.1

build:
	go build -ldflags "-X github.com/ray846/fops/cmd.Version=$(version)"