package tests

import (
	"testing"

	"github.com/coocn-cn/leanote/app/db"
	//	. "github.com/coocn-cn/leanote/app/lea"
	//	"github.com/coocn-cn/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
