SHELL := /bin/bash

# 项目配置
PROJECT_NAME := gradiation_artifact_demo
MYSQL_CONTAINER := proj_template_mysql
BACKEND_CONTAINER := proj_template_backend
FRONTEND_CONTAINER := proj_template_frontend

.PHONY: help push pull up down dev dev-frontend dev-backend clean rebuild ps logs

# 默认显示帮助
help:
	@echo "淘宝克隆项目 - Makefile 命令"
	@echo "======================================"
	@echo "一键启动:"
	@echo "  make up          # Podman Compose 全栈启动 (MySQL + Backend + Frontend)"
	@echo "  make down        # 停止所有容器"
	@echo "  make rebuild     # 重新构建并启动"
	@echo "  make ps          # 查看容器状态"
	@echo "  make logs         # 查看日志"
	@echo ""
	@echo "本地开发 (单独启动):"
	@echo "  make dev         # 启动 MySQL + 后端，前端需单独 npm run dev"
	@echo "  make dev-mysql   # 只启动 MySQL"
	@echo "  make dev-backend # 启动 MySQL + 后端"
	@echo "  make dev-frontend# 前端开发模式 (npm run dev)"
	@echo ""
	@echo "清理:"
	@echo "  make clean       # 停止容器 + 删除数据卷 (全新初始化)"
	@echo ""
	@echo "Git 操作:"
	@echo "  make push        # git add/commit/push"
	@echo "  make pull        # git pull"
	@echo "======================================"

# ==================== 一键启动 ====================

# 全栈容器化启动 (生产模式)
up:
	@echo "🚀 Podman Compose 全栈启动..."
	podman compose up -d --build
	@echo "✅ 启动完成!"
	@echo "   前端: http://localhost:3000"
	@echo "   后端: http://localhost:8080"
	@echo "   API:  http://localhost:8080/api/v1/health"
	@make ps

# 停止所有容器
down:
	@echo "🛑 停止所有容器..."
	podman compose down
	@echo "✅ 已停止"

# 重新构建并启动
rebuild: down clean-data
	@echo "🔄 重新构建并启动..."
	podman compose up -d --build
	@make ps

# 查看容器状态
ps:
	@echo "📊 容器状态:"
	podman compose ps

# 查看日志
logs:
	podman compose logs -f

# ==================== 本地开发 ====================

# 启动 MySQL (本地开发第一步)
dev-mysql:
	@echo "🐬 启动 MySQL..."
	podman compose up -d mysql
	@sleep 10
	@echo "✅ MySQL 已启动，等待初始化完成..."
	@podman exec -i $(MYSQL_CONTAINER) mysql -uappuser -papppassword appdb --default-character-set=utf8mb4 -e "SELECT COUNT(*) as tables FROM information_schema.tables WHERE table_schema='appdb';" 2>/dev/null || echo "⏳ MySQL 正在初始化，请稍等..."

# 启动后端 (本地开发，MySQL 必须先启动)
dev-backend:
	@echo "🔧 启动后端..."
	cd backend && go run ./cmd/server/

# 本地开发模式: MySQL + 后端同时启动
dev: dev-mysql
	@echo "🚀 启动后端 (前台运行，Ctrl+C 停止)..."
	cd backend && go run ./cmd/server/

# 前端开发模式 (需后端已启动)
dev-frontend:
	@echo "⚛️  启动前端开发服务器..."
	cd frontend && npm run dev

# ==================== 清理 ====================

# 清理数据卷 (全新初始化)
clean-data:
	@echo "🧹 删除数据卷..."
	podman volume rm $(PROJECT_NAME)_mysql_data 2>/dev/null || true

# 完全清理 (容器 + 数据卷)
clean: down clean-data
	@echo "✅ 已完全清理"

# ==================== Git 操作 ====================

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
	@echo "======================================"