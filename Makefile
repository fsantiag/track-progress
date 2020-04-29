.PHONY: docker

build-backend:
	cd backend && dep ensure && GOOS=linux GOARCH=amd64 go build -o track-progress

build-frontend:
	cd frontend && npm install && npm run build

move-frontend:
	mkdir -p backend/static && mv frontend/build/* backend/static/ && rmdir frontend/build

build: clean build-frontend move-frontend build-backend

image:
	docker build -t track-progress:1.0.0 .

clean:
	rm -rf backend/static
	rm -rf backend/vendor
	rm -rf backend/track-progress
	rm -rf frontend/build
	rm -rf frontend/node_modules

build-container: build image