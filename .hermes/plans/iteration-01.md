# AI Product Research 迭代计划

## 功能 1: Web UI 搜索和筛选
- 添加搜索框（产品名称/关键词）
- 添加分类筛选下拉框
- 添加价格范围筛选
- 表格排序功能

## 功能 2: 定时爬虫 Cron Job
- 集成 cron 库实现定时任务
- Toolify 爬虫定时运行
- Product Hunt 爬虫定时运行
- 可配置的间隔时间

## 实现步骤

### 前端 (Next.js)
1. 更新 Dashboard 页面添加搜索和筛选组件
2. 添加状态管理（搜索关键词、分类、价格范围）
3. 更新 API 调用传递筛选参数

### 后端 (Go)
1. 添加 cron 包依赖
2. 创建 scheduler 包
3. 在 main.go 中初始化定时任务
4. 添加 API 用于配置定时任务
