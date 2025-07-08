/blog-backend/
├── config/
│   └── config.go
├── controllers/
│   ├── auth_controller.go
│   ├── comment_controller.go
│   └── post_controller.go
├── middlewares/
│   └── auth.go
├── models/
│   ├── comment.go
│   ├── post.go
│   └── user.go
├── routes/
│   └── routes.go
├── utils/
│   ├── jwt.go
│   └── response.go
└── main.go

## 8. 环境变量

创建 `.env` 文件：

```
JWT_SECRET=your_jwt_secret_key
PORT=8080
```

## 9. 依赖管理

`go.mod` 文件：

```go
module github.com/yourusername/blog-backend

go 1.21

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.9.1
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.12.0
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.2
)

require (
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

## 10. API 文档

### 用户认证

- **注册用户**
  - POST `/auth/register`
  - Body: `{ "username": "user1", "email": "user1@example.com", "password": "password123" }`

- **用户登录**
  - POST `/auth/login`
  - Body: `{ "username": "user1", "password": "password123" }`
  - Response: `{ "token": "jwt_token" }`

### 文章管理

- **获取所有文章**
  - GET `/posts`

- **获取单个文章**
  - GET `/posts/:id`

- **创建文章** (需要认证)
  - POST `/posts`
  - Headers: `Authorization: Bearer jwt_token`
  - Body: `{ "title": "First Post", "content": "This is my first post" }`

- **更新文章** (需要认证，只能更新自己的文章)
  - PUT `/posts/:id`
  - Headers: `Authorization: Bearer jwt_token`
  - Body: `{ "title": "Updated Post", "content": "Updated content" }`

- **删除文章** (需要认证，只能删除自己的文章)
  - DELETE `/posts/:id`
  - Headers: `Authorization: Bearer jwt_token`

### 评论管理

- **获取文章评论**
  - GET `/posts/:postId/comments`

- **创建评论** (需要认证)
  - POST `/posts/:postId/comments`
  - Headers: `Authorization: Bearer jwt_token`
  - Body: `{ "content": "Great post!" }`

## 11. 运行项目

1. 初始化项目：
```bash
go mod init github.com/yourusername/blog-backend
go mod tidy
```

2. 运行项目：
```bash
go run main.go
```

3. 测试API：
   - 使用Postman或curl测试各个API端点

这个实现包含了所有要求的功能：
- 用户认证（注册/登录，JWT）
- 文章CRUD（带权限控制）
- 评论功能
- 错误处理
- 日志记录
- 数据库模型定义

你可以根据需要进一步扩展功能，如添加分页、文章分类、标签等功能。