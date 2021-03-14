package repository

type PredicateBuilder interface {
	Predicater
	BuildParams() map[string]interface{}
}

type PredicateBuild struct {
	name   string
	data   interface{}
	params map[string]interface{}
}

func NewPredicate(name string, data interface{}) *PredicateBuild {
	return &PredicateBuild{
		data:   data,
		name:   name,
		params: make(map[string]interface{}),
	}
}

func NewBuilder(p Predicater) *PredicateBuild {
	if builder, ok := p.(*PredicateBuild); ok {
		return builder.Clone()
	}

	return NewPredicate(p.Predicate(), p.Data())
}

func (m *PredicateBuild) Data() interface{} {
	return m.data
}

func (m *PredicateBuild) Predicate() string {
	return m.name
}

func (m *PredicateBuild) BuildParams() map[string]interface{} {
	return m.params
}

func (m *PredicateBuild) Clone() *PredicateBuild {
	return &PredicateBuild{
		data:   m.data,
		name:   m.name,
		params: make(map[string]interface{}),
	}
}

func (m *PredicateBuild) WithSelect(fields []string) *PredicateBuild {
	params := m.params

	params["_common_select"] = fields

	return m
}

func (m *PredicateBuild) WithLimit(limit int) *PredicateBuild {
	params := m.params

	params["_common_limit"] = limit

	return m
}

func (m *PredicateBuild) WithPage(skip, limit int) *PredicateBuild {
	params := m.params

	params["_common_skip"] = skip
	params["_common_limit"] = limit

	return m
}

func (m *PredicateBuild) WithSort(sort string) *PredicateBuild {
	params := m.params

	params["_common_sort"] = sort

	return m
}

func (m *PredicateBuild) WithUser(userID string) *PredicateBuild {
	params := m.params

	params["_common_fields_user"] = userID

	return m
}

func (m *PredicateBuild) WithTrash(trash bool) *PredicateBuild {
	params := m.params

	params["_common_fields_trash"] = trash

	return m
}

func (m *PredicateBuild) WithDeleted(deleted bool) *PredicateBuild {
	params := m.params

	params["_common_fields_deleted"] = deleted

	return m
}
