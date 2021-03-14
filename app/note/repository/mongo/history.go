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

type history struct {
	baseRepository
	collection *mgo.Collection
}

func NewHistory(_ context.Context) *history {
	obj := &history{
		collection: db.NoteContentHistories,
	}

	return obj
}

// New is 初始化一个新的领域对象
func (m *history) New(ctx context.Context, data model.HistoryData) *model.History {
	return model.NewHistory(newHistoryMutation(data, nil))
}

// Count is 获取符合 Predicater 条件的领域对象的个数
func (m *history) Count(ctx context.Context, predicate repository.Predicater) (int, error) {
	return db.Count(m.collection, m.predicates(ctx, predicate)), nil
}

// Find is 加载一个符合 Predicater 条件的领域对象
func (m *history) Find(ctx context.Context, predicate repository.Predicater) (*model.History, error) {
	historys, err := m.FindAll(ctx, repository.NewBuilder(predicate).WithLimit(1))
	if err != nil {
		return nil, err
	}

	if len(historys) == 0 {
		return nil, nil
	}

	return historys[0], nil
}

// FindAll is 加载所有符合 Predicater 条件的领域对象
func (m *history) FindAll(ctx context.Context, predicate repository.Predicater) ([]*model.History, error) {
	historys := []info.NoteContentHistory{}
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

	q.All(&historys)

	resp := make([]*model.History, 0, len(historys))
	for i := range historys {
		resp = append(resp, model.NewHistory(newHistoryMutation(model.HistoryData(historys[i]), &historys[i])))
	}

	return resp, nil
}

// Save is 保存领域对象到仓储
func (m *history) Save(ctx context.Context, models ...*model.History) error {
	for _, model := range models {
		mut := model.HistoryMutation.(*historyMutation)
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
			db.Insert(m.collection, info.NoteContentHistory(saveData))
		}

		// 更新 model.mutation 对象状态
		*mut = *newHistoryMutation(saveData, (*info.NoteContentHistory)(&saveData))

		log.G(ctx).WithField("old", mut.old).WithField("new", saveData).Info("History Save")
	}

	return nil
}

// Delete is 删除仓储中的领域对象
func (m *history) Delete(ctx context.Context, models ...*model.History) error {
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
func (m *history) DeleteID(ctx context.Context, ids ...uint64) error {
	// TODO: history.DeleteID(ctx context.Context, ids ...uint64) error Not implemented

	panic("history.DeleteID(ctx context.Context, ids ...uint64) error Not implemented")
}

func (m *history) predicates(ctx context.Context, predicate repository.Predicater) bson.M {
	var query bson.M

	switch predicate.Predicate() {
	case "HistoryNoteID":
		params := predicate.Data().(map[string]string)
		query = bson.M{"_id": bson.ObjectIdHex(params["noteID"])}
	default:
		return m.predicateToMongo(ctx, predicate)
	}

	return m.commonFields(predicate, query)
}
