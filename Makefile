
tidy:
	go mod tidy

build:
	CGO_ENABLED=0 GOOS=linux go build .

docker-build:
	docker build . 