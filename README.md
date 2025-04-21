## 1.项目介绍

- 技术栈：gin框架、gorm数据库框架、jwt登录认证
- 实现模块：用户模块
- 实现接口，详情看提交记录：
	- 登录、注册
	- 获取当前用户
	- 修改用户
	- 获取用户个人资料
	- 关注/取消关注用户

## 2.项目目录结构

```go
my-realworld-go
├─ go.mod           // 使用mod管理
├─ go.sum           // 
├─ main.go          // 启动go
├─ README.md        // readme
├─ user             // 用户模块
│  ├─ user_auth.go      // 用户认证
│  ├─ user_model.go     // 用户模型
│  ├─ user_route.go     // 用户路由
│  └─ user_service.go   // 用户服务
├─ script               
│  └─ my-realworld.sql  // 数据库脚本
├─ common           // 通用模块
│  ├─ database.go       // 数据库
│  └─ utils.go          // 工具类
```

## 3.项目收获

1. 第一次接触到了gim和gorm框架，碰到了一些问题，慢慢查资料解决；
2. 感受最深的点：和 Java 最不一样的地方，重新看到了 C语言 的结构体数据结构，大一满满的回忆啊
3. 继续加油叭

## 4.待优化的点

1. 返回的用户视图、个人资料视图，代码冗余
2. 数据库表名统一设置
3. 打日志