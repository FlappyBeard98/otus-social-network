PROGRAM_NAME := $(if $(PROGRAM_NAME),$(error))


install:
	 cd $(PROGRAM_NAME) && CGO_ENABLED=0 go build -mod=mod -v 

clean-app:
	go clean

pre-build:
	swag init --parseVendor --parseInternal --parseDependency -o . --ot json -g main.go

all: pre-build install clean-app 

.PHONY: all
