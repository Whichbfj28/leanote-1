// Code generated by ImplGen.
// Source: history_repository.go

package mongo

import (
	context "context"

	"github.com/coocn-cn/leanote/app/db"
	"github.com/coocn-cn/leanote/app/info"
	model "github.com/coocn-cn/leanote/app/note/model"
	repository "github.com/coocn-cn/leanote/app/note/repository"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"github.com/coocn-cn/leanote/pkg/log"
)

type history struct {
}

// Newhistory create a new history object
func Newhistory(_ context.Context) *history {
	obj := &history{}

	return obj
}

// New is 初始化一个新的领域对象
func (m *history) New(ctx context.Context, data model.HistoryData) *model.History {
	return model.NewHistory(newHistoryMutation(data, nil))
}

// Load is 从仓储加载一个领域对象
func (m *history) Load(ctx context.Context, id uint64) (*model.History, error) {
	// TODO: history.Load(ctx context.Context, id uint64) (*model.History, error) Not implemented

	panic("history.Load(ctx context.Context, id uint64) (*model.History, error) Not implemented")
}

// Find is 加载一个符合 Predicater 条件的领域对象
func (m *history) Find(ctx context.Context, predicates repository.Predicater) (*model.History, error) {
	historys, err := m.FindAll(ctx, predicates)
	if err != nil {
		return nil, err
	}

	if len(historys) == 0 {
		return nil, nil
	}

	return historys[0], nil
}

// FindAll is 加载所有符合 Predicater 条件的领域对象
func (m *history) FindAll(ctx context.Context, predicates repository.Predicater) ([]*model.History, error) {
	switch predicates.Predicate() {
	case "HistoryUserAndID":
	default:
		return nil, errcode.Unimplemented(ctx, "加载条件未实现", predicates.Predicate())
	}
	params := predicates.Data().(map[string]string)

	// 先查是否存在历史记录, 没有则添加之
	history := info.NoteContentHistory{}
	db.GetByIdAndUserId(db.NoteContentHistories, params["id"], params["userID"], &history)

	return []*model.History{model.NewHistory(newHistoryMutation(model.HistoryData(history), &history))}, nil
}

// Save is 保存领域对象到仓储
func (m *history) Save(ctx context.Context, model *model.History) error {
	mut := model.M().(*historyMutation)
	saveData := mut.Data()

	if mut.db != nil {
		// 更新之
		db.UpdateByIdAndUserId(db.NoteContentHistories, mut.db.NoteId.String(), mut.db.UserId.String(), info.NoteContentHistory(saveData))
	} else {
		// 保存之
		db.Insert(db.NoteContentHistories, info.NoteContentHistory(saveData))
	}

	// 更新 model.mutation 对象状态
	*mut = *newHistoryMutation(saveData, (*info.NoteContentHistory)(&saveData))

	log.G(ctx).WithField("old", mut.old).WithField("new", saveData).Info("History Save")
	return nil
}

// Delete is 删除仓储中的领域对象
func (m *history) Delete(ctx context.Context, model *model.History) error {
	// TODO: history.Delete(ctx context.Context, model *model.History) error Not implemented

	panic("history.Delete(ctx context.Context, model *model.History) error Not implemented")
}

// DeleteID is 通过ID删除仓储中的领域对象
func (m *history) DeleteID(ctx context.Context, id uint64) error {
	// TODO: history.DeleteID(ctx context.Context, id uint64) error Not implemented

	panic("history.DeleteID(ctx context.Context, id uint64) error Not implemented")
}