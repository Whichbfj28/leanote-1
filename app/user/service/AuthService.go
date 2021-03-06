package service

import (
	"context"

	"gopkg.in/mgo.v2/bson"
	//	"github.com/coocn-cn/leanote/app/db"
	"github.com/coocn-cn/leanote/app/info"
	//	"github.com/revel/revel"
	"errors"
	"fmt"
	"strconv"
	"strings"

	. "github.com/coocn-cn/leanote/app/lea"
)

// 登录与权限 Login & Register

type AuthService struct {
	userSrv     *UserService
	configSrv   ConfigService
	shareSrv    ShareService
	emailSrv    EmailService
	blogSrv     BlogService
	noteSrv     NoteService
	notebookSrv NotebookService
}

func NewAuth(ctx context.Context, userSrv *UserService, configSrv ConfigService, shareSrv ShareService, emailSrv EmailService, blogSrv BlogService, noteSrv NoteService, notebookSrv NotebookService) *AuthService {
	return &AuthService{
		userSrv:     userSrv,
		configSrv:   configSrv,
		shareSrv:    shareSrv,
		emailSrv:    emailSrv,
		blogSrv:     blogSrv,
		noteSrv:     noteSrv,
		notebookSrv: notebookSrv,
	}
}

// 使用bcrypt认证或者Md5认证
// Use bcrypt (Md5 depreciated)
func (m *AuthService) Login(emailOrUsername, pwd string) (info.User, error) {
	emailOrUsername = strings.Trim(emailOrUsername, " ")
	//	pwd = strings.Trim(pwd, " ")
	userInfo := m.userSrv.GetUserInfoByName(emailOrUsername)
	if userInfo.UserId == "" || !ComparePwd(pwd, userInfo.Pwd) {
		return userInfo, errors.New("wrong username or password")
	}
	return userInfo, nil
}

// 注册
/*
注册 leanote@leanote.com userId = "5368c1aa99c37b029d000001"
添加 在博客上添加一篇欢迎note, note1 5368c1b919807a6f95000000

将nk1(只读), nk2(可写) 分享给该用户
将note1 复制到用户的生活nk上
*/
// 1. 添加用户
// 2. 将leanote共享给我
// [ok]
func (m *AuthService) Register(email, pwd, fromUserId string) (bool, string) {
	// 用户是否已存在
	if m.userSrv.IsExistsUser(email) {
		return false, "userHasBeenRegistered-" + email
	}
	passwd := GenPwd(pwd)
	if passwd == "" {
		return false, "GenerateHash error"
	}
	user := info.User{UserId: bson.NewObjectId(), Email: email, Username: email, Pwd: passwd}
	if fromUserId != "" && IsObjectId(fromUserId) {
		user.FromUserId = bson.ObjectIdHex(fromUserId)
	}
	return m.register(user)
}

func (m *AuthService) register(user info.User) (bool, string) {
	if m.userSrv.AddUser(user) {
		// 添加笔记本, 生活, 学习, 工作
		userId := user.UserId.Hex()
		notebook := info.Notebook{
			Seq:    -1,
			UserId: user.UserId}
		title2Id := map[string]bson.ObjectId{"life": bson.NewObjectId(), "study": bson.NewObjectId(), "work": bson.NewObjectId()}
		for title, objectId := range title2Id {
			notebook.Title = title
			notebook.NotebookId = objectId
			notebook.UserId = user.UserId
			m.notebookSrv.AddNotebook(notebook)
		}

		// 添加leanote -> 该用户的共享
		registerSharedUserId := m.configSrv.GetGlobalStringConfig("registerSharedUserId")
		if registerSharedUserId != "" {
			registerSharedNotebooks := m.configSrv.GetGlobalArrMapConfig("registerSharedNotebooks")
			registerSharedNotes := m.configSrv.GetGlobalArrMapConfig("registerSharedNotes")
			registerCopyNoteIds := m.configSrv.GetGlobalArrayConfig("registerCopyNoteIds")

			// 添加共享笔记本
			for _, notebook := range registerSharedNotebooks {
				perm, _ := strconv.Atoi(notebook["perm"])
				m.shareSrv.AddShareNotebookToUserId(notebook["notebookId"], perm, registerSharedUserId, userId)
			}

			// 添加共享笔记
			for _, note := range registerSharedNotes {
				perm, _ := strconv.Atoi(note["perm"])
				m.shareSrv.AddShareNoteToUserId(note["noteId"], perm, registerSharedUserId, userId)
			}

			// 复制笔记
			for _, noteId := range registerCopyNoteIds {
				note := m.noteSrv.CopySharedNote(noteId, title2Id["life"].Hex(), registerSharedUserId, user.UserId.Hex())
				//				Log(noteId)
				//				Log("Copy")
				//				LogJ(note)
				noteUpdate := bson.M{"IsBlog": false} // 不要是博客
				m.noteSrv.UpdateNote(user.UserId.Hex(), note.NoteId.Hex(), noteUpdate, -1)
			}
		}

		//---------------
		// 添加一条userBlog
		m.blogSrv.UpdateUserBlog(info.UserBlog{UserId: user.UserId,
			Title:      user.Username + " 's Blog",
			SubTitle:   "Love Leanote!",
			AboutMe:    "Hello, I am (^_^)",
			CanComment: true,
		})
		// 添加一个单页面
		m.blogSrv.AddOrUpdateSingle(user.UserId.Hex(), "", "About Me", "Hello, I am (^_^)")
	}

	return true, ""
}

//--------------
// 第三方注册

// 第三方得到用户名, 可能需要多次判断
func (m *AuthService) getUsername(thirdType, thirdUsername string) (username string) {
	username = thirdType + "-" + thirdUsername
	i := 1
	for {
		if !m.userSrv.IsExistsUserByUsername(username) {
			return
		}
		username = fmt.Sprintf("%v%v", username, i)
	}
}

func (m *AuthService) ThirdRegister(thirdType, thirdUserId, thirdUsername string) (exists bool, userInfo info.User) {
	userInfo = m.userSrv.GetUserInfoByThirdUserId(thirdUserId)
	if userInfo.UserId != "" {
		exists = true
		return
	}

	username := m.getUsername(thirdType, thirdUsername)
	userInfo = info.User{UserId: bson.NewObjectId(),
		Username:      username,
		ThirdUserId:   thirdUserId,
		ThirdUsername: thirdUsername,
	}
	_, _ = m.register(userInfo)
	return
}
