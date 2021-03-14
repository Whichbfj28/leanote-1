package service

import (
	"context"

	"github.com/coocn-cn/leanote/app/info"
	"github.com/coocn-cn/leanote/app/note/repository"
	note_repo "github.com/coocn-cn/leanote/app/note/repository"
	note_mongo "github.com/coocn-cn/leanote/app/note/repository/mongo"
	tag_model "github.com/coocn-cn/leanote/app/tag/model"
	tag_repo "github.com/coocn-cn/leanote/app/tag/repository"
	tag_mongo "github.com/coocn-cn/leanote/app/tag/repository/mongo"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"github.com/coocn-cn/leanote/pkg/log"
)

/*
每添加,更新note时, 都将tag添加到tags表中
*/
type TagService struct {
	tag     tag_repo.TagRepository
	noteTag note_repo.TagRepository
}

func NewTag(ctx context.Context) *TagService {
	return &TagService{
		tag:     tag_mongo.NewTag(ctx),
		noteTag: note_mongo.NewTag(ctx),
	}
}

func (this *TagService) AddTagsI(userId string, tags interface{}) bool {
	if ts, ok := tags.([]string); ok {
		return this.AddTags(userId, ts)
	}
	return false
}

func (m *TagService) AddTags(userId string, tags []string) bool {
	ctx := context.Background()

	err := updateTag(m.tag, ctx, repository.User(userId), func(tag *tag_model.Tag) error {
		if tag == nil {
			return errcode.NotFound(ctx, "notExists", userId)
		}

		return tag.AddTags(ctx, tags)
	})
	if err != nil {
		return false
	}
	return true
}

// 同步用
func (m *TagService) GeSyncTags(userId string, afterUsn, maxEntry int) []info.NoteTag {
	ctx := context.Background()

	tags, err := m.noteTag.FindAll(ctx, tag_repo.TagNexts(afterUsn).WithUser(userId).WithLimit(maxEntry))
	if err != nil {
		log.G(ctx).WithError(err).Error("查询笔记tag失败")
		return nil
	}

	resp := make([]info.NoteTag, 0, len(tags))
	for _, v := range tags {
		resp = append(resp, info.NoteTag(v.MustData(ctx)))
	}

	return resp
}
