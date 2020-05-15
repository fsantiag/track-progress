.PHONY: docker

images: clean-dependencies frontend-image backend-image bff-image clean-builds

frontend-image: frontend-build
	cd frontend && docker build -t frontend .

backend-image: backend-build
	cd backend && docker build -t backend .

bff-image: bff-build
	cd back-for-front && docker build -t bff .

backend-build:
	cd backend && dep ensure && go test -v ./... && GOOS=linux GOARCH=amd64 go build -o backend

bff-build:
	cd back-for-front && dep ensure && cd src && go test -v ./... && GOOS=linux GOARCH=amd64 go build -o bff

frontend-build:
	cd frontend && npm install && npm run build

clean-dependencies:
	rm -rf backend/vendor
	rm -rf back-for-front/vendor
	rm -rf frontend/node_modules
	
clean-builds:
	rm -rf backend/backend
	rm -rf back-for-front/src/bff
	rm -rf frontend/build