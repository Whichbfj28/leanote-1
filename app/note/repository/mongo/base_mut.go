package mongo

import (
	context "context"
	"reflect"

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
