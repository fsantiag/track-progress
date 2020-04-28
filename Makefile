.PHONY: docker

build: 
	cd backend && GOOS=linux GOARCH=amd64 go build -o track-progress

image: 
	docker build -t track-progress:1.0.0 .

clean:
	rm -rf backend/track-progress

build-container: build image clean