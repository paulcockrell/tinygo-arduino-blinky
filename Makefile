GOPATH:=$(shell go env GOPATH)

.PHONY: deploy
deploy:
	tinygo flash -target arduino ./main.go
