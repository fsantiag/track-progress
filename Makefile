.PHONY: docker

images: clean frontend-image backend-image

frontend-image: frontend-build
	cd frontend && docker build -t frontend .

backend-image: backend-build
	cd backend && docker build -t backend .

backend-build:
	cd backend && dep ensure && GOOS=linux GOARCH=amd64 go build -o backend

frontend-build:
	cd frontend && npm install && npm run build

clean:
	rm -rf backend/track-progress
	rm -rf backend/vendor
	rm -rf frontend/node_modules
	rm -rf frontend/build
