## 源码结构
```
leanote/app/
    controllers 控制器
    db mongodb通用数据库访问方法, 由service调用
    info 数据表的模型和其它数据结构
    lea 通用方法
    service 服务
    view 视图
```

## Controller 控制器
```
init.go 初始化方法, 注入service

BaseController.go 基控制器, 所有控制器都继承自它

IndexController.go leanote首页
MobileController.go 移动端页面

AuthController.go 用户登录/注销/找回密码
OauthController.go 第三方登录验证, 现只有github
UserController.go 用户, 修改密码, 用户名

NotebookController.go 笔记本
NoteController.go 笔记
NoteContentHistoryController.go 笔记历史
ShareController.go 共享笔记/笔记本

BlogController.go 博客
FileController.go 文件上传, 现只有图片上传
```

## Service 服务
leanote的服务相当于php mvc的model. 服务之间可相互调用, 但服务是根据功能来划分的, 而不是根据数据表(model)
```
init.go 初始化, 注入各个service
common.go 公用方法

AuthService.go 登录与权限
PwdService.go 密码服务, 修改, 找回
UserService.go 用户
TokenService.go Token, 用于找回密码

NotebookService.go 笔记本
NoteService.go 笔记
NoteContentHistoryService.go 笔记历史
TrashService.go 废纸篓服务
TagService.go  笔记标签
ShareService.go 共享笔记/笔记本

BlogService.go 博客

SuggestionService.go 建议(已废弃)
```

## Leanote db

在db/目录下只有一个文件 Mgo.go.
* 包含表的Collection对象, 在leanote启动时会连接数据库, 并实例化所有表的Collection对象. 如
```
// 数据库连接
var err error
Session, err = mgo.Dial(url)

// 实例化各个Collection
NoteContents = Session.DB(dbname).C("note_contents")
NoteContentHistories = Session.DB(dbname).C("note_content_histories")
```

* 包含处理数据的公用方法, 如
```
func Insert(collection *mgo.Collection, i interface{}) bool {
	err := collection.Insert(i)
	return Err(err)
}

// 适合一条记录全部更新
func Update(collection *mgo.Collection, query interface{}, i interface{}) bool {
	err := collection.Update(query, i)
	return Err(err)
}
func Upsert(collection *mgo.Collection, query interface{}, i interface{}) bool {
	_, err := collection.Upsert(query, i)
	return Err(err)
}
``` 

## Leanote controller, service, db流程
leanote controller处理用户的请求, 但只做请求的分发, 会处理一些用户传的参数. 但真正实现用户的请求是通过调用service来进行处理.

controller的作用:

* 接收用户的请求
* 处理用户的参数, 构造合适的数据
* 调用service

service之间可相互调用, service其实就是数据库的操作, 是通过db来完成的.