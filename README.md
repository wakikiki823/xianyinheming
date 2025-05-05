# xianyinheming-backend
弦音和鸣，基于ai的陪练平台。 这是一个关于弦音和鸣app的后端仓库。

## 功能模块

### 用户系统
- 手机号/邮箱注册登录
- JWT鉴权机制
- 个人信息管理

### 资源管理
- 七牛云对象存储

### AI接口
- 大模型API集成

## 快速启动

1. 复制配置文件模板：
```bash
cp config/config.yaml.example config/config.yaml
```

2. 根据实际情况修改配置文件

3. 启动服务：
```bash
go run cmd/main.go
```

## 项目结构
```
├── config        # 配置管理
├── controllers   # 控制器层
├── dao           # 数据访问层
├── database      # 数据库连接
├── models        # 数据模型
├── proto         # gRPC协议文件
├── routes        # 路由配置
├── service       # 业务服务层
└── utils         # 通用工具
```

