PORT?=8000
APPLICATION:=fruits-services
PACKAGE:=github.com/philip-bui/${APPLICATION}
PACKAGES=`go list ./...`
COVERAGE_FILE:=coverage-report.out
MYSQL_DOCKER:=mysql:8.0
MYSQL_NAME:=mysql
MYSQL_USER:=philip
MYSQL_PW:=password
MYSQL_DB:=fairfax
MYSQL_PORT:=3306
WORKING_DIR:=$(shell pwd)

help: ## Display this help screen
	grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

godoc: ## Open HTML API Documentation
	echo "localhost:${PORT}/pkg/${PACKAGE}"
	godoc -http=:${PORT}

mod: # Downloads dependencies
	export GOMODULES111=on
	go get ${PACKAGES}

.PHONY: test
test: ## Run Tests
	go test ./...

benchmark: ## Run Benchmark Tests
	go test -v -bench=. ./..

coverage-report.out: test ## Generate Coverage File
	go test -coverprofile=${COVERAGE_FILE} ./...

coverage: coverage-report.out ## Open HTML Test Coverage Report
	go tool cover -html=${COVERAGE_FILE}

mysql: ## Run MySQL Docker
	docker run --name ${MYSQL_NAME} -p ${MYSQL_PORT}:${MYSQL_PORT} -v ${WORKING_DIR}/resources/fairfax/:/docker-entrypoint-initdb.d -e MYSQL_ROOT_PASSWORD=${MYSQL_PW} -e MYSQL_DATABASE=${MYSQL_DB} -e MYSQL_USER=${MYSQL_USER} -e MYSQL_PASSWORD=${MYSQL_PW} ${MYSQL_DOCKER}

mysql-stop: ## Stop MySQL Docker
	docker stop ${MYSQL_NAME} && docker rm ${MYSQL_NAME}
