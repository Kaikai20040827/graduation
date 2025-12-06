# Introduction

# File Tree
``` directories
Secure_file_box/
├── cmd/
│   └── server/
│       └── main.go          
├── internal/
│   ├── config/              
│   │   └── config.go
│   ├── handler/             
│   │   ├── user_handler.go
│   │   ├── auth_handler.go
│   │   └── product_handler.go
│   ├── middleware/          
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── logger.go
│   ├── model/               
│   │   ├── user.go
│   │   ├── product.go
│   │   └── dto/             
│   ├── repository/          
│   │   ├── user_repo.go
│   │   └── product_repo.go
│   ├── service/             # 业务逻辑层
│   │   ├── user_service.go
│   │   └── product_service.go
│   ├── routes/              # 路由定义
│   │   ├── api.go
│   │   ├── auth_routes.go
│   │   └── user_routes.go
│   └── pkg/
│       ├── database/        # 数据库连接
│       │   └── db.go
│       ├── logger/          # 日志
│       │   └── logger.go
│       └── utils/           # 工具函数
│           └── jwt.go
├── pkg/                     # 可复用的公共包
├── api/                     # API 文档
│   └── swagger/
├── web/                     # Web 资源
│   ├── static/
│   │   ├── css/
│   │   ├── js/
│   │   └── images/
│   └── templates/
│       └── index.html
├── storage/                 # 存储目录
│   ├── uploads/
│   └── logs/
├── scripts/                 # 脚本文件
│   ├── migration/
│   └── deployment/
├── tests/                   # 测试文件
│   ├── unit/
│   └── integration/
├── docs/                    # 文档
├── .env.example             # 环境变量示例
├── .gitignore
├── go.mod
├── go.sum
├── Makefile                 # 构建脚本
└── README.md
```