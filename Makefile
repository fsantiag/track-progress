.PHONY: docker

build: 
	cd api && GOOS=linux GOARCH=amd64 go build -o track-progress

image: 
	docker build -t track-progress:1.0.0 .

clean:
	rm -rf api/track-progress

build-container: build image clean