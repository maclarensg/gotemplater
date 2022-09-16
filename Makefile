compile: 
	go build .

build:
	docker build . -t maclarensg/gotemplater:latest

push-local:
	docker push maclarensg/gotemplater:latest

test:
	go test . -coverprofile cover.out
	go tool cover -html=cover.out

.PHONY: compile build push-local test