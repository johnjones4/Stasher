PROJECT=$(shell basename $(shell pwd))
TAG=ghcr.io/johnjones4/${PROJECT}
VERSION=$(shell date +%s)

info:
	echo ${PROJECT} ${VERSION}

container:
	docker build -t ${TAG} .
	docker push ${TAG}:latest
	docker image rm ${TAG}:latest

bookmarklet:
	@cat bookmarklet.js | sed 's/STASH_KEY/$(STASH_KEY)/g' | sed 's/URL_ROOT/$(URL_ROOT)/g' | tr '\n' ' '

ci: container
