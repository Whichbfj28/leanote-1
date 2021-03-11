package tests

import (
	"testing"

	"github.com/coocn-cn/leanote/app/db"
	"github.com/revel/revel"

	//	. "github.com/coocn-cn/leanote/app/lea"
	"github.com/coocn-cn/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func init() {
	revel.Init("dev", "github.com/coocn-cn/leanote", "/Users/life/Documents/Go/package_base/src")
	db.Init("mongodb://localhost:27017/leanote", "leanote")
	service.InitService()
	service.ConfigS.InitGlobalConfigs()
}

// 测试登录
func TestSendMail(t *testing.T) {
	ok, err := service.EmailS.SendEmail("life@leanote.com", "你好", "你好吗")
	t.Log(ok)
	t.Log(err)
}
