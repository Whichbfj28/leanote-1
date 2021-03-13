module github.com/coocn-cn/leanote

go 1.16

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/agtorre/gocolorize v1.0.0
	github.com/golang/mock v1.4.3
	github.com/google/wire v0.5.0
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/revel/cmd v0.0.0-00010101000000-000000000000
	github.com/revel/modules v0.0.0-00010101000000-000000000000
	github.com/revel/revel v1.0.0
	github.com/robfig/config v0.0.0-20141207224736-0f78529c8c7e
	github.com/robfig/cron v1.2.0
	github.com/ssoor/gex v0.6.4
	github.com/ssoor/implgen v0.1.2
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	google.golang.org/grpc v1.36.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/mgo.v2 v2.0.0-20150124113754-c6a7dce14133
)

replace (
	github.com/revel/cmd => ./vendor_/github.com/revel/cmd
	github.com/revel/modules => ./vendor_/github.com/revel/modules
	github.com/revel/revel => ./vendor_/github.com/revel/revel
)
