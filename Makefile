.PHONY: all dev clean build env-up env-down run

all: clean build env-up run

dev: build run

##### BUILD
build:
	@echo "Build ..."
	@go build
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@cd fixtures && docker-compose -f docker-compose.yaml -f docker-compose-couch.yaml up --force-recreate -d	
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fixtures && docker-compose -f docker-compose.yaml -f docker-compose-couch.yaml down	
	@echo "Environment down"

##### RUN
run:
	@echo "Start app ..."
	@./gosdk-example

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/gosdk-example-* gosdk-example
	@docker rm -f -v `docker ps -a --no-trunc | grep "gosdk-example" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "gosdk-example" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"
