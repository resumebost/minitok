# 命名规范

## 接口定义文件 (.proto)

- **文件名**: 使用 `snake_case`, 如: `user_service.proto`.
- **包名**: 简单小写, 避免下划线, 如: `userservice`.
- **消息名和枚举名**: 使用 `PascalCase`, 如: `UserProfile`.
- **字段名**: 使用 `snake_case`, 如: `first_name`.
- **枚举字段**: 使用 `UPPER_SNAKE_CASE`, 如: `STATUS_ACTIVE`.
- **服务名**: 使用 `PascalCase`, 如: `UserService`.
- **RPC 方法**: 使用 `PascalCase`, 如: `GetUserProfile`.
- **导入路径**: 简单小写字母和斜杠, 避免使用下划线, 如: `userservice`.

## 源文件 (.go) 

- **文件名**: 简单小写, 使用下划线分隔, 如: `user_service.go`.
- **包名**: 简单小写, 避免使用下划线, 如: `userservice`.
- **变量名**: 使用 `camelCase`, 如: `userInfo`.
- **函数名**: 使用 `PascalCase`, 并确保命名明确, 如: `GetUserByID()`.
- **常量**: 使用 `PascalCase` 或 `UPPER_SNAKE_CASE`, 如: `MaxUsers` 或 `MAX_USERS`.
- **接口名**: 使用 `PascalCase`, 并且通常以 `er` 结尾, 如: `Reader`.

