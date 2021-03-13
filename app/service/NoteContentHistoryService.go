package service

import (
	"context"

	"github.com/coocn-cn/leanote/app/info"
	"github.com/coocn-cn/leanote/app/note/model"
	"github.com/coocn-cn/leanote/app/note/repository"
	"github.com/coocn-cn/leanote/pkg/log"

	//	. "github.com/coocn-cn/leanote/app/lea"
	"gopkg.in/mgo.v2/bson"
	//	"time"
)

// 历史记录
type NoteContentHistoryService struct {
	repo repository.HistoryRepository
}

// 新建一个note, 不需要添加历史记录
// 添加历史
func (m *NoteContentHistoryService) AddHistory(noteId, userId string, eachHistory info.EachHistory) {
	ctx := context.Background()
	// 检查是否是空
	if eachHistory.Content == "" {
		return
	}

	// 先查是否存在历史记录, 没有则添加之
	history, err := m.repo.Find(ctx, repository.HistoryUserAndID(userId, noteId))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记历史列表失败")
		return
	}

	if history == nil {
		history = m.repo.New(ctx, model.HistoryData{
			UserId: bson.ObjectIdHex(userId),
			NoteId: bson.ObjectIdHex(noteId),
		})
	}

	history.AddHistory(eachHistory)

	err = m.repo.Save(ctx, history)
	if nil != err {
		log.G(ctx).WithError(err).Error("保存笔记历史失败")
	}
}

// 列表展示
func (m *NoteContentHistoryService) ListHistories(noteId, userId string) []info.EachHistory {
	ctx := context.Background()

	history, err := m.repo.Find(ctx, repository.HistoryUserAndID(userId, noteId))
	if err != nil {
		log.G(ctx).WithError(err).Error("获取笔记历史列表失败")
		return []info.EachHistory{}
	}

	if history == nil {
		return []info.EachHistory{}
	}

	return history.Data().Histories
}
