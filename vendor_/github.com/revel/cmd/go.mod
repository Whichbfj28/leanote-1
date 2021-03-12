module github.com/revel/cmd

go 1.16

require (
	github.com/agtorre/gocolorize v1.0.0
	github.com/revel/modules v0.0.0-00010101000000-000000000000
	github.com/revel/revel v0.18.1-0.20171106060741-b8b463b10112
)

replace (
	github.com/revel/modules => ../modules
	github.com/revel/revel => ../revel
)
