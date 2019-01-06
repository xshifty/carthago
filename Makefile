PROJECT_ROOT=$(shell pwd)
BINARY_DIR=${PROJECT_ROOT}/bin
BINARY=carthago

all: clean ensure
	@mkdir -p ${BINARY_DIR}
	@go build -o ${BINARY_DIR}/${BINARY} ${PROJECT_ROOT}/cmd/carthago/main.go

ensure:
	@cd ${PROJECT_ROOT}
	@dep ensure

clean:
	@rm -rf ${BINARY_DIR}/${BINARY}

update:
	@git pull origin master

.PHONY: clean update ensure all

