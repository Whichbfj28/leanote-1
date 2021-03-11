package controllers

import (
	"github.com/revel/revel"
	//	"encoding/json"
	//	"gopkg.in/mgo.v2/bson"
	. "github.com/coocn-cn/leanote/app/lea"
	"github.com/coocn-cn/leanote/app/lea/captcha"

	//	"github.com/coocn-cn/leanote/app/types"
	//	"io/ioutil"
	//	"fmt"
	//	"math"
	//	"os"
	//	"path"
	//	"strconv"
	"io"
	"net/http"
)

// 验证码服务
type Captcha struct {
	BaseController
}

type Ca string

func (r Ca) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "image/png")
}

func (c Captcha) Get() revel.Result {
	c.Response.ContentType = "image/png"
	image, str := captcha.Fetch()
	out := io.Writer(c.Response.GetWriter())
	image.WriteTo(out)

	sessionId := c.Session["_ID"]
	//	LogJ(c.Session)
	//	Log("------")
	//	Log(str)
	//	Log(sessionId)
	Log("..")
	sessionService.SetCaptcha(sessionId, str)

	return c.Render()
}
