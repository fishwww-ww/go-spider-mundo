# go爬虫小demo

爬取mundo中的组队信息,并保存在数据库中

## 快速启动

1. 本地运行docker的mysql容器，配置信息见db.go
2. 安装依赖
```bash
  go mod tidy
```
3. 运行

## 项目亮点

- 使用gorm框架操作数据,而不是原生sql语句
- 使用docker容器,简化mysql的安装和配置

## 后续完善
- 学习docker目录挂载,将数据持久化在宿主机中