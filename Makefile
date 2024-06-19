default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done


arch?=$(shell go env GOARCH)
os?=$(shell go env GOOS)
ARCH=$(arch)
OS=$(os)
CGO=1
BIN_FOLDER=bin
PLUGINS_FOLDER=plugins
PLUGIN_GO_FILE=plugin.go
PLUGIN_NAME=plugin.so
ERRORG_DIRECTORY= ~/.errorg



.PHONY: go-plugin
init:  go-plugin http-plugin # Build whole plugins


.PHONY: go-plugin
go-plugin: # To build a  golang standard errors plugin.
	@echo "  >  Building go-plugin..."
	chmod  -R   700  ${CURDIR}/${PLUGINS_FOLDER}
	CGO_ENABLED=${CGO} GOOS=${OS} GOARCH=${ARCH} go build  -buildmode=plugin   -o  ${CURDIR}/${PLUGINS_FOLDER}/go/${PLUGIN_NAME} ${PLUGINS_FOLDER}/go/${PLUGIN_GO_FILE}

.PHONY: http-plugin
http-plugin: # To build a  http errors plugin.
	@echo "  >  Building http-plugin..."
	chmod -R   700  ${CURDIR}/${PLUGINS_FOLDER}
	CGO_ENABLED=${CGO} GOOS=${OS} GOARCH=${ARCH} go build  -buildmode=plugin   -o  ${CURDIR}/${PLUGINS_FOLDER}/http/${PLUGIN_NAME} ${PLUGINS_FOLDER}/http/${PLUGIN_GO_FILE}


.PHONY: example
example: # To example
	@echo "  >  Building example go app..."
	CGO_ENABLED=${CGO} GOOS=${OS} GOARCH=${ARCH} go build     -o  ${BIN_FOLDER}/errorg example/main.go
	chmod  +x  ${BIN_FOLDER}/errorg
	./${BIN_FOLDER}/errorg
