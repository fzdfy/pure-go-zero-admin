// Code generated by goctl. DO NOT EDIT.
package types

type UserLoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Status           string `json:"status"`
	CurrentAuthority string `json:"currentAuthority"`
	Id               int64  `json:"id"`
	UserName         string `json:"userName"`
	AccessToken      string `json:"token"`
	AccessExpire     int64  `json:"accessExpire"`
	RefreshAfter     int64  `json:"refreshAfter"`
}

type UserInfoResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
}

type ListUserReq struct {
	Current  int64  `form:"current,default=1"`
	PageSize int64  `form:"pageSize,default=20"`
	Name     string `form:"name,optional"`
	NickName string `form:"nickName,optional"`
	Mobile   string `form:"mobile,optional"`
	Email    string `form:"email,optional"`
	Status   int64  `form:"status,optional,default=-1"`
}

type ListUserData struct {
	Id             int64  `json:"id"`             // 编号
	Name           string `json:"name"`           // 用户名
	NickName       string `json:"nickName"`       // 昵称
	Avatar         string `json:"avatar"`         // 头像
	Password       string `json:"password"`       // 密码
	Salt           string `json:"salt"`           // 加密盐
	Email          string `json:"email"`          // 邮箱
	Mobile         string `json:"mobile"`         // 手机号
	Status         int64  `json:"status"`         // 状态  0：禁用   1：正常
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
	DelFlag        int64  `json:"delFlag"`        // 是否删除  -1：已删除  0：正常
}

type ListUserResp struct {
	Code     string          `json:"code"`
	Message  string          `json:"message"`
	Current  int64           `json:"current,default=1"`
	Data     []*ListUserData `json:"data"`
	PageSize int64           `json:"pageSize,default=20"`
	Success  bool            `json:"success"`
	Total    int64           `json:"total"`
}

type AddUserReq struct {
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
}

type AddUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type UpdateUserReq struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Status   int64  `json:"status"`
}

type UpdateUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeleteUserReq struct {
	Id int64 `json:"id"`
}

type DeleteUserResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AddMenuReq struct {
	Name     string `json:"name"`              // 菜单名称
	ParentId int64  `json:"parentId,optional"` // 父菜单ID，一级菜单为0
	Path     string `json:"path,optional"`     // 菜单URL,类型：1.普通页面（如用户管理， /sysmodel/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	Perms    string `json:"perms,optional"`    // 授权(多个用逗号分隔，如：sysmodel:user:add,sysmodel:user:edit)
	Type     int64  `json:"type,optional"`     // 类型   0：目录   1：菜单   2：按钮
	Icon     string `json:"icon,optional"`     // 菜单图标
	OrderNum int64  `json:"orderNum,optional"` // 排序
}

type AddMenuResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ListMenuReq struct {
	Current  int64  `form:"current,default=1"`
	PageSize int64  `form:"pageSize,default=20"`
	Name     string `form:"name,optional"`
	Path     string `form:"path,optional"`
	Status   int64  `form:"status,optional,default=-1"`
}

type ListMenuData struct {
	Id             int64  `json:"id"`             // 编号
	Name           string `json:"name"`           // 菜单名称
	ParentId       int64  `json:"parentId"`       // 父菜单ID，一级菜单为0
	Path           string `json:"path"`           // 菜单Path,类型：1.普通页面（如用户管理， /sysmodel/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	Perms          string `json:"perms"`          // 授权(多个用逗号分隔，如：sysmodel:user:add,sysmodel:user:edit)
	Type           int64  `json:"type"`           // 类型   0：目录   1：菜单   2：按钮
	Icon           string `json:"icon"`           // 菜单图标
	OrderNum       int64  `json:"orderNum"`       // 排序
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
	DelFlag        int64  `json:"delFlag"`        // 是否删除  -1：已删除  0：正常
}

type ListMenuResp struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    []*ListMenuData `json:"data"`
	Success bool            `json:"success"`
	Total   int64           `json:"total"`
}

type UpdateMenuReq struct {
	Id       int64  `json:"id"`                // 编号
	Name     string `json:"name"`              // 菜单名称
	ParentId int64  `json:"parentId"`          // 父菜单ID，一级菜单为0
	Path     string `json:"path,optional"`     // 菜单URL,类型：1.普通页面（如用户管理， /sysmodel/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	Perms    string `json:"perms,optional"`    // 授权(多个用逗号分隔，如：sysmodel:user:add,sysmodel:user:edit)
	Type     int64  `json:"type,optional"`     // 类型   0：目录   1：菜单   2：按钮
	Icon     string `json:"icon,optional"`     // 菜单图标
	OrderNum int64  `json:"orderNum,optional"` // 排序
	Status   int64  `json:"status,optional"`
}

type UpdateMenuResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeleteMenuReq struct {
	Id int64 `json:"id"`
}

type DeleteMenuResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AddRoleReq struct {
	Name   string `json:"name"`   // 角色名称
	Remark string `json:"remark"` // 备注
}

type AddRoleResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ListRoleReq struct {
	Current  int64  `form:"current,default=1"`
	PageSize int64  `form:"pageSize,default=20"`
	Name     string `form:"name,optional "`
	Status   int64  `form:"status,optional,default=-1"`
}

type ListRoleData struct {
	Id             int64  `json:"id"`             // 编号
	Name           string `json:"name"`           // 角色名称
	Remark         string `json:"remark"`         // 备注
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
	DelFlag        int64  `json:"delFlag"`        // 是否删除  -1：已删除  0：正常
	Status         int64  `json:"status"`         // 状态
}

type ListRoleResp struct {
	Code     string          `json:"code"`
	Message  string          `json:"message"`
	Current  int64           `json:"current,default=1"`
	Data     []*ListRoleData `json:"data"`
	PageSize int64           `json:"pageSize,default=20"`
	Success  bool            `json:"success"`
	Total    int64           `json:"total"`
}

type UpdateRoleReq struct {
	Id     int64  `json:"id"`     // 编号
	Name   string `json:"name"`   // 角色名称
	Remark string `json:"remark"` // 备注
	Status int64  `json:"status"` // 状态
}

type UpdateRoleResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeleteRoleReq struct {
	Id int64 `json:"id"`
}

type DeleteRoleResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
