module github.com/coocn-cn/leanote

go 1.16

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/agtorre/gocolorize v1.0.0
	github.com/revel/cmd v0.0.0-00010101000000-000000000000
	github.com/revel/modules v0.0.0-00010101000000-000000000000
	github.com/revel/revel v1.0.0
	github.com/robfig/config v0.0.0-20141207224736-0f78529c8c7e
	github.com/robfig/cron v1.2.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/mgo.v2 v2.0.0-20150124113754-c6a7dce14133
)

replace (
	github.com/revel/cmd => ./vendor_/github.com/revel/cmd
	github.com/revel/modules => ./vendor_/github.com/revel/modules
	github.com/revel/revel => ./vendor_/github.com/revel/revel
)
