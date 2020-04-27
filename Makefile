.PHONY: docker

build: 
	GOOS=linux GOARCH=amd64 go build -o track-progress

image: 
	docker build -t track-progress:1.0.0 .

clean:
	rm -rf track-progress

build-container: build image clean