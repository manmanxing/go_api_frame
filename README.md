# go_api_frame

go语言开发，开箱即用的API 框架的轮子，基于 Gin，xorm，beego，gomail框架
基于go 1.13版本

功能点：

+ go modules 包管理工具
+ 使用yaml配置文件
+ 默认使用mysql数据库
+ 实现指定接口使用redis缓存
+ 实现图片上传，限制图片格式，大小
+ 使用zap日志库
+ 使用 JWT 进行身份校验
+ 优雅的退出和重启服务
+ 支持Swagger文档
+ 部署到docker 
+ 支持shell脚本启动，重启，停止
+ 支持针对报错类型发送email通知
+ 采用 beego 的 validation的数据校验方式