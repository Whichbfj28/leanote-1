package repository

type Predicater interface {
	Data() interface{}
	Predicate() string
}

// All is 获取所有
func All() *PredicateBuild {
	data := true

	return NewPredicate("All", data)
}

// ID is 查询条件 - 从仓储加载一个领域对象
func ID(id string) *PredicateBuild {
	data := map[string]interface{}{
		"id": id,
	}

	return NewPredicate("ID", data)
}

// IDs is 查询条件 - 从仓储加载多个领域对象
func IDs(ids []string) *PredicateBuild {
	data := map[string]interface{}{
		"ids": ids,
	}

	return NewPredicate("IDs", data)
}

// User is 查询条件 - 从仓储加载多个属于指定用户的领域对象
func User(userID string) *PredicateBuild {
	data := map[string]interface{}{
		"userID": userID,
	}

	return NewPredicate("User", data)
}
