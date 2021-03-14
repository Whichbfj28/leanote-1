package service

import (
	"github.com/coocn-cn/leanote/app/info"
	"gopkg.in/mgo.v2/bson"
)

type EmailService interface {
	FindPwdSendEmail(token, email string) (ok bool, msg string)
	RegisterSendActiveEmail(userInfo info.User, email string) bool
}

type ShareService interface {
	DeleteAllShareNotebookGroup(groupId string) bool
	DeleteAllShareNoteGroup(groupId string) bool
	DeleteShareNotebookGroupWhenDeleteGroupUser(userId, groupId string) bool
	DeleteShareNoteGroupWhenDeleteGroupUser(userId, groupId string) bool

	AddShareNotebookToUserId(notebookId string, perm int, userId, toUserId string) (bool, string, string)
	AddShareNoteToUserId(noteId string, perm int, userId, toUserId string) (bool, string, string)
}

type BlogService interface {
	GetUserBlog(userId string) info.UserBlog
	GetBlogUrls(userBlog *info.UserBlog, userInfo *info.User) info.BlogUrls
	GetUserBlogUrl(userBlog *info.UserBlog, username string) string

	UpdateUserBlog(userBlog info.UserBlog) bool
	AddOrUpdateSingle(userId, singleId, title, content string) (ok bool)
}

type ConfigService interface {
	GetAdminUserId() string

	GetGlobalStringConfig(key string) string
	GetGlobalArrMapConfig(key string) []map[string]string
	GetGlobalArrayConfig(key string) []string
}

type NoteService interface {
	CopySharedNote(noteId, notebookId, fromUserId, myUserId string) info.Note
	UpdateNote(updatedUserId, noteId string, needUpdate bson.M, usn int) (bool, string, int)
}

type NotebookService interface {
	AddNotebook(notebook info.Notebook) (bool, info.Notebook)
}

// 分页, 排序处理
func parsePageAndSort(pageNumber, pageSize int, sortField string, isAsc bool) (skipNum int, sortFieldR string) {
	skipNum = (pageNumber - 1) * pageSize
	if sortField == "" {
		sortField = "UpdatedTime"
	}
	if !isAsc {
		sortFieldR = "-" + sortField
	} else {
		sortFieldR = sortField
	}
	return
}
