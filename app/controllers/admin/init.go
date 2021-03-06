package admin

import (
	"github.com/coocn-cn/leanote/app/info"
	note_service "github.com/coocn-cn/leanote/app/note/service"
	"github.com/coocn-cn/leanote/app/service"
	user_service "github.com/coocn-cn/leanote/app/user/service"

	//	. "github.com/coocn-cn/leanote/app/lea"
	"github.com/revel/revel"
	//	"strings"
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
var upgradeService *service.UpgradeService
var noteService *note_service.NoteService
var notebookService *note_service.NotebookService
var userService *user_service.UserService
var authService *user_service.AuthService

// 拦截器
// 不需要拦截的url
// Index 除了Note之外都不需要
var commonUrl = map[string]map[string]bool{"Index": map[string]bool{"Index": true,
	"Login":              true,
	"DoLogin":            true,
	"Logout":             true,
	"Register":           true,
	"DoRegister":         true,
	"FindPasswword":      true,
	"DoFindPassword":     true,
	"FindPassword2":      true,
	"FindPasswordUpdate": true,
	"Suggestion":         true,
},
	"Blog": map[string]bool{"Index": true,
		"View":       true,
		"AboutMe":    true,
		"SearchBlog": true,
	},
	// 用户的激活与修改邮箱都不需要登录, 通过链接地址
	"User": map[string]bool{"UpdateEmail": true,
		"ActiveEmail": true,
	},
	"Oauth":  map[string]bool{"GithubCallback": true},
	"File":   map[string]bool{"OutputImage": true, "OutputFile": true},
	"Attach": map[string]bool{"Download": true, "DownloadAll": true},
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
func AuthInterceptor(c *revel.Controller) revel.Result {
	// 全部变成首字大写
	/*
		var controller = strings.Title(c.Name)
		var method = strings.Title(c.MethodName)
		// 是否需要验证?
		if !needValidate(controller, method) {
			return nil
		}
	*/

	// 验证是否已登录
	// 必须是管理员
	if username, ok := c.Session["Username"]; ok && username == configService.GetAdminUsername() {
		return nil // 已登录
	}

	// 没有登录, 判断是否是ajax操作
	if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		re := info.NewRe()
		re.Msg = "NOTLOGIN"
		return c.RenderJSON(re)
	}

	return c.Redirect("/login")
}

// 最外层init.go调用
// 获取service, 单例
func InitService() {
	notebookService = service.NotebookS
	noteService = service.NoteS
	shareService = service.ShareS
	userService = service.UserS
	blogService = service.BlogS
	noteImageService = service.NoteImageS
	fileService = service.FileS
	albumService = service.AlbumS
	attachService = service.AttachS
	suggestionService = service.SuggestionS
	authService = service.AuthS
	configService = service.ConfigS
	emailService = service.EmailS
	upgradeService = service.UpgradeS
}

func init() {
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &Admin{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &AdminSetting{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &AdminUser{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &AdminBlog{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &AdminEmail{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &AdminUpgrade{})
	revel.InterceptFunc(AuthInterceptor, revel.BEFORE, &AdminData{})
}
