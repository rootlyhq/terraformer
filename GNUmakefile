default: build

.PHONY: build
build:
	go get -u github.com/rootlyhq/terraform-provider-rootly
	yarn install
	node generators/rootly/rootly.js
