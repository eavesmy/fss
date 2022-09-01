build-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fss src/main.go

build-arm64:
	CGO_ENABLED=0 GOOS=drawin GOARCH=arm64 go build -o fss src/main.go
