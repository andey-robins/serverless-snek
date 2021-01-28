.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/end source/end/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/move source/move/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/ping source/ping/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/start source/start/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
