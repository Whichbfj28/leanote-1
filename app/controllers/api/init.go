package api

import (
	"github.com/coocn-cn/leanote/app/info"
	"github.com/coocn-cn/leanote/app/service"

	note_service "github.com/coocn-cn/leanote/app/note/service"
	tag_service "github.com/coocn-cn/leanote/app/tag/service"
	user_service "github.com/coocn-cn/leanote/app/user/service"

	//		. "github.com/coocn-cn/leanote/app/lea"
	"strings"

	"github.com/revel/revel"
)

var shareService *service.ShareService
var blogService *service.BlogService
var suggestionService *service.SuggestionService
var albumService *service.AlbumService
var noteImageService *service.NoteImageService
var fileService *service.FileService
var attachService *service.AttachService
var configService *service.ConfigService
var emailService *service.EmailService
var tagService *tag_service.TagService
var noteService *note_service.NoteService
var trashService *note_service.TrashService
var notebookService *note_service.NotebookService
var userService *user_service.UserService
var sessionService *user_service.SessionService
var authService *user_service.AuthService

var pageSize = 1000
var defaultSortField = "UpdatedTime"
var leanoteUserId = "admin" // 不能更改

// 状态
const (
	S_DEFAULT                 = iota // 0
	S_NOT_LOGIN                      // 1
	S_WRONG_USERNAME_PASSWORD        // 2
	S_WRONG_CAPTCHA                  // 3
	S_NEED_CAPTCHA                   // 4
	S_NOT_OPEN_REGISTER              // 4
)

// 拦截器
// 不需要拦截的url
var commonUrl = map[string]map[string]bool{"ApiAuth": map[string]bool{"Login": true,
	"Register": true,
},
	// 文件的操作也不用登录, userId会从session中获取
	"ApiFile": map[string]bool{"GetImage": true,
		"GetAttach":     true,
		"GetAllAttachs": true,
	},
}

func needValidate(controller, method string) bool {
	// 在里面
	if v, ok := commonUrl[controller]; ok {
		// 在commonUrl里
		if _, ok2 := v[method]; ok2 {
			return false
		}
		return true
	} else {
		// controller不在这里的, 肯定要验证
		return true
	}
}

// 这里得到token, 若不是login, logout等公用操作, 必须验证是否已登录
func AuthInterceptor(c *revel.Controller) revel.Result {
	// 得到token /api/user/info?userId=xxx&token=xxxxx
	token := c.Params.Values.Get("token")
	noToken := false
	if token == "" {
		// 若无, 则取sessionId
		token = c.Session.ID()
		noToken = true
	}
	c.Session["_token"] = token

	// 全部变成首字大写
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)

	// 验证是否已登录
	// 通过sessionService判断该token下是否有userId, 并返回userId
	userId := sessionService.GetUserId(token)
	if noToken && userId == "" {
		// 从session中获取, api/file/getImage, api/file/getAttach, api/file/getAllAttach
		// 客户端
		userId, _ = c.Session["UserId"]
	}
	c.Session["_userId"] = userId

	// 是否需要验证?
	if !needValidate(controller, method) {
		return nil
	}

	if userId != "" {
		return nil // 已登录
	}

	// 没有登录, 返回错误的信息, 需要登录
	re := info.NewApiRe()
	re.Msg = "NOTLOGIN"
	return c.RenderJSON(re)
}

func init() {
	// interceptors
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &ApiAuth{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &ApiUser{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &ApiFile{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &ApiNote{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &ApiTag{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &ApiNotebook{})
}

// 最外层init.go调用
// 获取service, 单例
func InitService() {
	notebookService = service.NotebookS
	noteService = service.NoteS
	trashService = service.TrashS
	shareService = service.ShareS
	userService = service.UserS
	tagService = service.TagS
	blogService = service.BlogS
	noteImageService = service.NoteImageS
	fileService = service.FileS
	albumService = service.AlbumS
	attachService = service.AttachS
	suggestionService = service.SuggestionS
	authService = service.AuthS
	configService = service.ConfigS
	emailService = service.EmailS
	sessionService = service.SessionS
}
