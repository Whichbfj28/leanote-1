package service

import (
	"github.com/coocn-cn/leanote/app/info"
	"gopkg.in/mgo.v2/bson"
)

type BlogService interface {
	ReCountBlogTags(userId string) bool
}

type UserService interface {
	IncrUsn(userId string) int
}

type ConfigService interface {
	GetSiteUrl() string
}
type ShareService interface {
	DeleteShareNoteAll(noteId string, userId string) bool
	HasReadPerm(userId, updatedUserId, noteId string) bool
	HasUpdatePerm(userId, updatedUserId, noteId string) bool
	HasUpdateNotebookPerm(userId, updatedUserId, notebookId string) bool
}
type NoteImageService interface {
	GetImagesByNoteIds(noteIds []bson.ObjectId) map[string][]info.File
	UpdateNoteImages(userId, noteId, imgSrc, content string) bool
	CopyNoteImages(fromNoteId, fromUserId, newNoteId, content, toUserId string) string
}
type AttachService interface {
	CopyAttachs(noteId, toNoteId, toUserId string) bool
	GetAttachsByNoteIds(noteIds []bson.ObjectId) map[string][]info.Attach
}
