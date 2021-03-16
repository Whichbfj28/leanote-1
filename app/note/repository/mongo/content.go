// Code generated by ImplGen.
// Source: content_repository.go

package mongo

import (
	context "context"

	"github.com/coocn-cn/leanote/app/db"
	"github.com/coocn-cn/leanote/app/info"
	model "github.com/coocn-cn/leanote/app/note/model"
	repository "github.com/coocn-cn/leanote/app/note/repository"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"github.com/coocn-cn/leanote/pkg/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type content struct {
	baseRepository
	collection *mgo.Collection
}

func NewContent(_ context.Context) *content {
	obj := &content{
		collection: db.NoteContents,
	}

	return obj
}

// New is 初始化一个新的领域对象
func (m *content) New(ctx context.Context, data model.ContentData) *model.Content {
	return model.NewContent(newContentMutation(data, nil))
}

// Count is 获取符合 Predicater 条件的领域对象的个数
func (m *content) Count(ctx context.Context, predicate repository.Predicater) (int, error) {
	return db.Count(m.collection, m.predicates(ctx, predicate)), nil
}

// Find is 加载一个符合 Predicater 条件的领域对象
func (m *content) Find(ctx context.Context, predicate repository.Predicater) (*model.Content, error) {
	contents, err := m.FindAll(ctx, repository.NewBuilder(predicate).WithLimit(1))
	if err != nil {
		return nil, err
	}

	if len(contents) == 0 {
		return nil, nil
	}

	return contents[0], nil
}

// FindAll is 加载所有符合 Predicater 条件的领域对象
func (m *content) FindAll(ctx context.Context, predicate repository.Predicater) ([]*model.Content, error) {
	contents := []info.NoteContent{}
	q := m.collection.Find(m.predicates(ctx, predicate))

	if builder, ok := predicate.(repository.PredicateBuilder); ok {
		params := builder.BuildParams()

		if sort, ok := params["_common_sort"].(string); ok {
			q = q.Sort(sort)
		}
		if skip, ok := params["_common_skip"].(int); ok {
			q = q.Skip(skip)
		}
		if limit, ok := params["_common_limit"].(int); ok {
			q = q.Limit(limit)
		}
	}

	q.All(&contents)

	resp := make([]*model.Content, 0, len(contents))
	for i := range contents {
		resp = append(resp, model.NewContent(newContentMutation(model.ContentData(contents[i]), &contents[i])))
	}

	return resp, nil
}

// Save is 保存领域对象到仓储
func (m *content) Save(ctx context.Context, models ...*model.Content) error {
	for _, model := range models {
		mut := model.ContentMutation.(*contentMutation)
		saveData, err := mut.Data(ctx)
		if err != nil {
			return err
		}

		if mut.db() != nil {
			// 更新之
			if ok := db.UpdateByIdAndUserIdMap(m.collection, mut.db().NoteId.Hex(), mut.db().UserId.Hex(), mut.updates); !ok {
				return errcode.NotFound(ctx, "更新对象失败", mut.db, saveData)
			}
		} else {
			// 保存之
			db.Insert(m.collection, info.NoteContent(saveData))
		}

		// 更新 model.mutation 对象状态
		*mut = *newContentMutation(saveData, (*info.NoteContent)(&saveData))

		log.G(ctx).WithField("old", mut.old).WithField("new", saveData).Info("Content Save")
	}

	return nil
}

// Delete is 删除仓储中的领域对象
func (m *content) Delete(ctx context.Context, models ...*model.Content) error {
	for _, model := range models {
		data, err := model.Data(ctx)
		if err != nil {
			return err
		}

		if ok := db.DeleteByIdAndUserId(m.collection, data.NoteId.Hex(), data.UserId.Hex()); !ok {
			return errcode.NotFound(ctx, "删除对象失败", data)
		}
	}

	return nil
}

// DeleteID is 通过ID删除仓储中的领域对象
func (m *content) DeleteID(ctx context.Context, ids ...uint64) error {
	// TODO: content.DeleteID(ctx context.Context, ids ...uint64) error Not implemented

	panic("content.DeleteID(ctx context.Context, ids ...uint64) error Not implemented")
}

func (m *content) predicates(ctx context.Context, predicate repository.Predicater) bson.M {
	var query bson.M

	switch predicate.Predicate() {
	case "ContentNoteID":
		params := predicate.Data().(map[string]string)
		query = bson.M{"_id": bson.ObjectIdHex(params["noteID"])}
	case "ContentNoteIDs":
		params := predicate.Data().(map[string][]string)

		ids := params["noteIDs"]
		hexIDs := make([]bson.ObjectId, 0, len(ids))
		for _, v := range ids {
			hexIDs = append(hexIDs, bson.ObjectIdHex(v))
		}

		query = bson.M{"_id": bson.M{"$in": hexIDs}}
	default:
		return m.predicateToMongo(ctx, predicate)
	}

	return m.commonFields(predicate, query)
}