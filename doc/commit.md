# Commit 规范

## Commit 消息结构

每次提交的消息应该只包含一个**type**和一个**subject**. 结构如下: 

```
<type>: <subject>
```

## Commit 类型

- **feat**: 新功能
- **fix**: 修复bug
- **docs**: 文档变更
- **style**: 代码格式 (不影响代码运行的变动)
- **refactor**: 代码重构
- **perf**: 改善性能
- **test**: 测试相关
- **chore**: 构建或辅助工具的变动

## Subject

主题应简短, 清晰并描述 commit 的目的. 不要超过 50 个字符. 

## 示例

1. 新增功能: 
```bash
git add .
git commit -m "feat: 添加用户登录功能"
```

2. 修复bug: 
```bash
git add .
git commit -m "fix: 修复登录 bug"
```

3. 更新文档: 
```bash
git add  README.md
git commit -m "docs: 更新安装说明"
```

