package repository

type Predicater interface {
	Data() interface{}
	Predicate() string
}

type basePredicate struct {
	data interface{}
	name string
}

func (m *basePredicate) Data() interface{} {
	return m.data
}

func (m *basePredicate) Predicate() string {
	return m.name
}
