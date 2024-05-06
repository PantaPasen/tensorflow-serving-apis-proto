clean: ## Clean dist folder expcet third party which we no longer generate
	rm  -rf tmp go


protos: 
	buf generate

pull:
	sh pull_protos.sh

prune:
	python3 prune_protos.py

build: clean pull prune protos

.PHONY: build

.PHONY: help
.DEFAULT_GOAL := help

help:
	@for i in $(MAKEFILE_LIST); do grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $$i | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}';done


