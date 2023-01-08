clean:
	go clean

test:
	go test -v ./...

coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

build:
	go build -o ./build