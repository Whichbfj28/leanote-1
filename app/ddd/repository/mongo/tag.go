// Code generated by ImplGen.
// Source: comment.go

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

type tag struct {
	baseRepository
	collection *mgo.Collection
}

func Newtag(_ context.Context) *tag {
	obj := &tag{
		collection: db.NoteContentHistories,
	}

	return obj
}

// New is 初始化一个新的领域对象
func (m *tag) New(ctx context.Context, data model.TagData) *model.Tag {
	return model.NewTag(newtagMutation(data, nil))
}

// Count is 获取符合 Predicater 条件的领域对象的个数
func (m *tag) Count(ctx context.Context, predicate repository.Predicater) (int, error) {
	return db.Count(m.collection, m.predicates(ctx, predicate)), nil
}

// Find is 加载一个符合 Predicater 条件的领域对象
func (m *tag) Find(ctx context.Context, predicate repository.Predicater) (*model.Tag, error) {
	tags, err := m.FindAll(ctx, repository.NewBuilder(predicate).WithLimit(1))
	if err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, nil
	}

	return tags[0], nil
}

// FindAll is 加载所有符合 Predicater 条件的领域对象
func (m *tag) FindAll(ctx context.Context, predicate repository.Predicater) ([]*model.Tag, error) {
	tags := []info.NoteTag{}
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

	q.All(&tags)

	resp := make([]*model.Tag, 0, len(tags))
	for i := range tags {
		resp = append(resp, model.NewTag(newtagMutation(model.TagData(tags[i]), &tags[i])))
	}

	return resp, nil
}

// Save is 保存领域对象到仓储
func (m *tag) Save(ctx context.Context, models ...*model.Tag) error {
	for _, model := range models {
		mut := model.TagMutation.(*tagMutation)
		saveData, err := mut.Data(ctx)
		if err != nil {
			return err
		}

		if mut.db() != nil {
			// 更新之
			if ok := db.UpdateByIdAndUserIdMap(m.collection, mut.db().TagId.Hex(), mut.db().UserId.Hex(), mut.updates); !ok {
				return errcode.NotFound(ctx, "更新对象失败", mut.db, saveData)
			}
		} else {
			// 保存之
			db.Insert(m.collection, info.NoteTag(saveData))
		}

		// 更新 model.mutation 对象状态
		*mut = *newtagMutation(saveData, (*info.NoteTag)(&saveData))

		log.G(ctx).WithField("old", mut.old).WithField("new", saveData).Info("Tag Save")
	}

	return nil
}

// Delete is 删除仓储中的领域对象
func (m *tag) Delete(ctx context.Context, models ...*model.Tag) error {
	for _, model := range models {
		data, err := model.Data(ctx)
		if err != nil {
			return err
		}

		if ok := db.DeleteByIdAndUserId(m.collection, data.TagId.Hex(), data.UserId.Hex()); !ok {
			return errcode.NotFound(ctx, "删除对象失败", data)
		}
	}

	return nil
}

// DeleteID is 通过ID删除仓储中的领域对象
func (m *tag) DeleteID(ctx context.Context, ids ...uint64) error {
	// TODO: tag.DeleteID(ctx context.Context, ids ...uint64) error Not implemented

	panic("tag.DeleteID(ctx context.Context, ids ...uint64) error Not implemented")
}

func (m *tag) predicates(ctx context.Context, predicate repository.Predicater) bson.M {
	var query bson.M

	switch predicate.Predicate() {
	case "TagNoteID":
		params := predicate.Data().(map[string]string)
		query = bson.M{"_id": bson.ObjectIdHex(params["noteID"])}
	default:
		return m.predicateToMongo(ctx, predicate)
	}

	return m.commonFields(predicate, query)
}
