
.PHONY: build
build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 go build -a -installsuffix cgo -o weatherSvc
	redis-server &
	./weatherSvc

#TODO: ADD REDIS start svc here

.PHONY: clean
clean: ## Remove all the temporary and build files
	go clean




