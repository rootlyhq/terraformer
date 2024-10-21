default: build

.PHONY: build
build:
	go get -u github.com/rootlyhq/terraform-provider-rootly/v2
	yarn install
	node generators/rootly/rootly.js
