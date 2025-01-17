# uniperm
The unified permission package

# Install permctl
```
go install github.com/go-the-way/uniperm/cmd/permctl@latest
```

# Models
- User 用户
- Role 角色
- Permission 权限
- RolePermission 角色权限

# Services

## DB 数据库
- SetDB 设置DB
- AutoMigrate 自动合并  
- SetPagination 设置分配器

## User 用户模块
- UserGetPage 分页查询
- UserGet 查询详情
- UserGetPerm 查询用户权限
- UserGetPermButton 查询用户按钮权限
- UserAdd 新增
- UserUpdate 修改
- UserUpdatePassword 修改密码
- UserUpdateRole 修改用户角色
- UserDelete 删除
- UserEnable 启用
- UserDisable 禁用
- UserLogin 用户登录
- UserLogout 用户注销

## Role 角色模块
- role.GetPage 分页查询
- role.Get 查询详情
- role.GetPerm 查询角色权限
- role.UpdatePerm 修改角色权限
- role.Add 新增
- role.Update 修改
- role.Delete 删除
- role.Enable 启用
- role.Disable 禁用

## Permission 权限模块
- permission.Tree 查询权限树
- permission.Get 查询
- permission.Add 新增
- permission.Update 编辑
- permission.Delete 删除
