default: release

DATE=$$(date '+%Y.%m.%d.%H%M%S')

.PHONY: release
release:
	go get -u github.com/rootlyhq/terraform-provider-rootly
	git add go.mod go.sum
	git commit -m "Automatic upgrade of terraform-provider-rootly" --allow-empty
	git tag ${DATE}
	echo "Finish release by pushing commit and tags to GitHub"
