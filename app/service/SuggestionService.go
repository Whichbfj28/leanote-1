package service

import (
	"github.com/coocn-cn/leanote/app/db"
	"github.com/coocn-cn/leanote/app/info"

	//	. "github.com/coocn-cn/leanote/app/lea"
	"gopkg.in/mgo.v2/bson"
	//	"time"
	//	"sort"
)

type SuggestionService struct {
}

// 得到某博客具体信息
func (this *SuggestionService) AddSuggestion(suggestion info.Suggestion) bool {
	if suggestion.Id == "" {
		suggestion.Id = bson.NewObjectId()
	}
	return db.Insert(db.Suggestions, suggestion)
}
