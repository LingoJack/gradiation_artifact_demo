# 搜索历史 API 开发需求

## 功能概述
为用户实现搜索历史管理功能，包括保存搜索记录、查询历史、删除历史等操作。

## 需求详情

### 1. API 接口设计

| 接口 | 方法 | 路径 | 认证 | 说明 |
|------|------|------|------|------|
| 获取搜索历史 | GET | `/api/v1/search-histories` | 需要 JWT | 获取当前用户的搜索历史，按时间倒序，最多返回 20 条 |
| 添加搜索历史 | POST | `/api/v1/search-histories` | 需要 JWT | 添加一条搜索记录（关键词不能为空，最大长度 200） |
| 删除单条历史 | DELETE | `/api/v1/search-histories/:id` | 需要 JWT | 删除指定的搜索历史（只能删除自己的） |
| 清空搜索历史 | DELETE | `/api/v1/search-histories` | 需要 JWT | 清空当前用户的所有搜索历史 |

### 2. 请求/响应格式

#### 2.1 获取搜索历史
```json
// Response
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "keyword": "iPhone 15",
      "created_at": "2025-01-15T10:30:00Z"
    }
  ]
}
```

#### 2.2 添加搜索历史
```json
// Request
{
  "keyword": "iPhone 15"
}

// Response
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "keyword": "iPhone 15",
    "created_at": "2025-01-15T10:30:00Z"
  }
}
```

### 3. 业务规则

1. **关键词验证**：
   - 不能为空
   - 最大长度 200 字符
   - 去除首尾空格

2. **去重逻辑**：
   - 同一用户相同关键词，更新 `created_at` 为当前时间（不重复插入）

3. **数量限制**：
   - 每个用户最多保留 50 条搜索历史
   - 超过限制时删除最旧的记录

4. **权限控制**：
   - 用户只能查看/删除自己的搜索历史
   - 删除单条时验证 `user_id` 匹配

### 4. 技术要求

1. 在 `service/` 目录创建 `search_history_service.go`
2. 在 `controller/` 目录创建 `search_history_controller.go`
3. 在 `router/router.go` 中注册路由（私有路由组）
4. 使用现有的 `SearchHistoriesDao` 进行数据库操作
5. 使用 Wire 进行依赖注入（更新 `wire.go` 和重新生成 `wire_gen.go`）

### 5. 错误码

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 1001 | 参数错误 |
| 1003 | 无权限 |

---

## 验收检查清单

### 代码实现检查
- [ ] `service/search_history_service.go` 文件存在且包含所有方法
- [ ] `controller/search_history_controller.go` 文件存在且包含所有方法
- [ ] 路由已在 `router/router.go` 中正确注册
- [ ] Wire 依赖注入已更新

### 功能测试检查
- [ ] GET `/api/v1/search-histories` 返回用户搜索历史
- [ ] POST `/api/v1/search-histories` 成功添加搜索记录
- [ ] DELETE `/api/v1/search-histories/:id` 成功删除单条记录
- [ ] DELETE `/api/v1/search-histories` 成功清空所有记录

### 业务规则检查
- [ ] 空关键词返回参数错误
- [ ] 超长关键词（>200字符）返回参数错误
- [ ] 相同关键词更新时间而非重复插入
- [ ] 删除他人记录返回无权限错误
- [ ] 未登录访问返回 401 错误

### 代码质量检查
- [ ] Go 代码编译通过 (`go build ./...`)
- [ ] 无明显的代码风格问题
