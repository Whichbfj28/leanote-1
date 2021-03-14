package service

import (
	"context"
	"time"

	"github.com/coocn-cn/leanote/app/info"
	note_model "github.com/coocn-cn/leanote/app/note/model"
	"github.com/coocn-cn/leanote/app/note/repository"
	note_repo "github.com/coocn-cn/leanote/app/note/repository"
	tag_model "github.com/coocn-cn/leanote/app/tag/model"
	tag_repo "github.com/coocn-cn/leanote/app/tag/repository"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"github.com/coocn-cn/leanote/pkg/log"
	"gopkg.in/mgo.v2/bson"
)

/*
每添加,更新note时, 都将tag添加到tags表中
*/
type TagService struct {
	tag      tag_repo.TagRepository
	note_tag note_repo.TagRepository
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

//---------------------------
// v2
// 第二版标签, 单独一张表, 每一个tag一条记录

// 添加或更新标签, 先查下是否存在, 不存在则添加, 存在则更新
// 都要统计下tag的note数
// 什么时候调用? 笔记添加Tag, 删除Tag时
// 删除note时, 都可以调用
// 万能
func (m *TagService) AddOrUpdateTag(userId string, tag string) info.NoteTag {
	ctx := context.Background()

	modelTag, err := m.note_tag.Find(ctx, tag_repo.TagTag(tag).WithUser(userId))
	if err != nil {
		log.G(ctx).WithError(err).Error("查询笔记tag失败")
		return info.NoteTag{}
	}

	noteTag := note_model.TagData{}
	if modelTag == nil {
		// 不存在, 则创建之
		noteTag.TagId = bson.NewObjectId()
		noteTag.Count = 1
		noteTag.Tag = tag
		noteTag.UserId = bson.ObjectIdHex(userId)
		noteTag.CreatedTime = time.Now()
		noteTag.UpdatedTime = time.Now()
		noteTag.Usn = userService.IncrUsn(userId)
		noteTag.IsDeleted = false

		modelTag = m.note_tag.New(ctx, noteTag)
	} else {
		// 更新 note 数
		modelTag.SetCount(ctx, noteService.CountNoteByTag(userId, tag))
	}

	// 之前删除过的, 现在要添加回来了
	log.G(ctx).WithField("tag", tag).Info("之前删除过的, 现在要添加回来了")
	modelTag.SoftDelete(ctx, false, userService.IncrUsn(userId))

	if err := m.note_tag.Save(ctx, modelTag); err != nil {
		log.G(ctx).WithError(err).Error("保存笔记tag失败")
		return info.NoteTag{}
	}

	return info.NoteTag(modelTag.MustData(ctx))
}

// 得到标签, 按更新时间来排序
func (m *TagService) GetTags(userId string) []info.NoteTag {
	ctx := context.Background()

	tags, err := m.note_tag.FindAll(ctx, tag_repo.User(userId).WithDeleted(false).WithSort("-UpdatedTime"))
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

// 删除标签
// 也删除所有的笔记含该标签的
// 返回noteId => usn
func (m *TagService) DeleteTag(userId string, tag string) (resp map[string]int) {
	ctx := context.Background()

	err := updateNoteTag(m.note_tag, ctx, repository.User(userId), func(model *note_model.Tag) error {
		if model == nil {
			return errcode.NotFound(ctx, "notExists", userId)
		}

		if err := model.SoftDelete(ctx, true, userService.IncrUsn(userId)); err != nil {
			return err
		}

		resp = noteService.UpdateNoteToDeleteTag(userId, tag)

		return nil
	})
	if err != nil {
		return nil
	}

	return resp
}

// 删除标签, 供API调用
func (m *TagService) DeleteTagApi(userId string, tag string, usn int) (ok bool, msg string, toUsn int) {
	ctx := context.Background()

	err := updateNoteTag(m.note_tag, ctx, note_repo.TagTag(tag).WithUser(userId), func(model *note_model.Tag) error {
		if model == nil {
			return errcode.NotFound(ctx, "notExists", userId)
		}

		if model.MustData(ctx).Usn > usn {
			return errcode.DeadlineExceeded(ctx, "conflict")
		}

		return model.SoftDelete(ctx, true, userService.IncrUsn(userId))
	})
	if err != nil {
		return false, err.Error(), 0
	}

	return true, "", 0
}

// 重新统计标签的count
func (this *TagService) reCountTagCount(userId string, tags []string) {
	if tags == nil {
		return
	}
	for _, tag := range tags {
		this.AddOrUpdateTag(userId, tag)
	}
}

// 同步用
func (m *TagService) GeSyncTags(userId string, afterUsn, maxEntry int) []info.NoteTag {
	ctx := context.Background()

	tags, err := m.note_tag.FindAll(ctx, tag_repo.TagNexts(afterUsn).WithUser(userId).WithLimit(maxEntry))
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
