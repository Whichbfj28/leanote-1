package mongo

import (
	context "context"
	"reflect"

	repository "github.com/coocn-cn/leanote/app/note/repository"
	"github.com/coocn-cn/leanote/pkg/errcode"
	"gopkg.in/mgo.v2/bson"
)

type baseMutation struct {
	db      interface{}
	old     interface{}
	updates bson.M
}

func newbaseMutation(data interface{}, db interface{}) *baseMutation {
	m := &baseMutation{
		db:      db,
		old:     data,
		updates: bson.M{},
	}

	return m
}

func (m *baseMutation) Data(_ context.Context, out interface{}) error {
	new := reflect.New(reflect.TypeOf(m.old))

	dataBin, err := bson.Marshal(m.updates)
	if err != nil {
		return err
	}

	if err := bson.Unmarshal(dataBin, new.Interface()); err != nil {
		return err
	}

	if err := merge(out, m.old); err != nil {
		return err
	}

	if err := merge(out, new.Elem().Interface()); err != nil {
		return err
	}

	return nil
}

func (m *baseMutation) SetField(name string, value interface{}) {
	m.updates[name] = value
}

func (m *baseMutation) SetFields(fields map[string]interface{}) {
	for name, value := range fields {
		m.SetField(name, value)
	}
}

type baseRepository struct {
}

func (m *baseRepository) commonFields(predicate repository.Predicater, query bson.M) bson.M {
	builder, ok := predicate.(repository.PredicateBuilder)
	if !ok {
		return query
	}
	params := builder.BuildParams()

	if userID, ok := params["_common_fields_user"].(string); ok {
		query["UserId"] = bson.ObjectIdHex(userID)
	}

	if trash, ok := params["_common_fields_trash"].(bool); ok && !trash {
		query["IsTrash"] = false
	}

	if deleted, ok := params["_common_fields_deleted"].(bool); ok && !deleted {
		query["IsDeleted"] = false
	}

	return query
}

func (m *baseRepository) predicateToMongo(ctx context.Context, predicate repository.Predicater) bson.M {
	var query bson.M

	switch predicate.Predicate() {
	case "All":
		query = bson.M{}
	case "ID":
		params := predicate.Data().(map[string]interface{})
		query = bson.M{
			"_id": bson.ObjectIdHex(params["id"].(string)),
		}
	case "IDs":
		params := predicate.Data().(map[string]interface{})

		ids := params["ids"].([]string)
		hexIDs := make([]bson.ObjectId, 0, len(ids))
		for _, v := range ids {
			hexIDs = append(hexIDs, bson.ObjectIdHex(v))
		}

		query = bson.M{
			"_id": bson.M{"$in": hexIDs},
		}
	case "User":
		params := predicate.Data().(map[string]interface{})
		query = bson.M{
			"UserId": bson.ObjectIdHex(params["userID"].(string)),
		}
	default:
		panic(errcode.Unimplemented(ctx, "加载条件未实现", predicate.Predicate()).Error())
	}

	return m.commonFields(predicate, query)

}
