build-erply:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -v -o erply ./cmd/erply