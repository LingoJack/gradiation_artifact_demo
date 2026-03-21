SHELL := /bin/bash

.PHONY: push

push:
	@git add . \
	&& (git commit -m "优化" || exit 0) \
	&& git push origin main;

pull: 
	@git pull origin main;

current_dir: ## 显示当前目录信息
	@echo "🔍 当前目录信息:"
	@echo "======================================"
	@echo "目录: $$(pwd)"
	@echo "版本: $(VERSION)"
	@echo "分支: $(GIT_BRANCH)"
	@echo "======================================"
