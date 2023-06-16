build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o NavIndex

docker-build-and-push:
	docker build --platform linux/amd64 -t zze326/navindex:v1.0 .
	docker push zze326/navindex:v1.0