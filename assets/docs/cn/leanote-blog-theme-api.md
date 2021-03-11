leanote博客模板的语法是`golang`模板语法, `golang`模板语法简洁, 很快就会上手, 大家可以参考leanote自带的主题模板.

一些最基本的用法:

* 输出: `{{$.blogInfo.UserId}}` 表示输出`blogInfo.UserId`变量, 比如 `<span>{{$.blogInfo.UserId}}</span>`
* 判断: `{{if $.blogInfo.OpendComment}} 为真的处理 {{else}} 为假为处理 {{end}}`
* range循环: `{{range $.posts}} {{.Title}} {{end}}` range循环输出所有文章标题
* 调用函数: `{{$.post.CreatedTime|datetime}}` 使用datetime函数来模式化时间, 会输出类似 `2014-11-5 12:33:22` 的数据

关于golang模板更多信息请查看 "golang模板语法帮助"

## 模板组织结构

标准的leanote主题模板组织结构如下, 其中`header.html`, `footer.html`, `paging.html`, `share_comment.html`, `highlight.html` 这些仅供其它模板引用, 可以不需要.

* theme.json 主题配置 [必须]
* header.html 头部模板, 供其它模板引用
* footer.html 底部模板, 供其它模板引用
* index.html 首页 [必须]
* cate.html 分类页 [必须]
* post.html 文章详情页 [必须]
* archive.html 归档页 [必须]
* single.html 单页 [必须]
* share_comment.html 分享与评论页, 供post.html引用
* highlight.html 代码高亮页, 供其它页面引用, index, cate, search, tag_posts, post
* paging.html 分页, 供其它模板引用
* tags.html 标签列表页 [必须]
* tag_posts.html 标签文章页 [必须]
* 404.html 错误页 [必须]
* style.css 样式
* images/ 图片文件夹
* images/screenshot.png 主题预览图

## 公用变量
公用变量表示在每个页面都可以使用的变量.
包括一些公用的url, 页面判断变量, `$.cates` 分类列表信息, `$.singles` 单页面列表 `$.themeInfo` 主题信息, `$.blogInfo` 博客信息, `$.tags` 标签, `$.recentPosts` 最近发表的5篇文章

### url 地址

因为leanote支持二级域名和自定义域名, 所以url会根据用户的定义而不同, 这里假设用户设置了`http://demo.leanote.com`二级域名, 和` http://demo.com` 自定义域名, 而leanote默认的博客域名为`http://blog.leanote.com`

| 变量        | 描述           | 
| ------------- |:-------------:|
| $.siteUrl      | 当前站点地址, 比如 `http://leanote.com`, `http://localhost:9000` |
| $.indexUrl     | 我的博客首页地址, 比如 `http://blog.leanote.com/用户名`, `http://demo.leanote.com`, `http://demo.com`, 优先级从低到高, 即如果你有自定义域名, 那么就是你的自定义域名 (下同) |
| $.cateUrl| 分类页url, 如 `http://blog.leanote.com/cate/用户名`, 或 `http://demo.leanote.com/cate` 或 `http://demo.com/cate`|
| $.postUrl | 文章详情页url, 如 `http://blog.leanote.com/post/用户名`, 或 `http://demo.leanote.com/post` 或 `http://demo.com/post` |
| $.searchUrl| 搜索页url, 如 `http://blog.leanote.com/search/用户名` 或 `http://demo.leanote.com/search` 或 `http://demo.com/search` |
|$.singleUrl | 单面面url, 如 `http://blog.leanote.com/single/用户名` 或 `http://demo.leanote.com/single` 或 `http://demo.com/single` |
|$.archivesUrl ($.archiveUrl) | 归档页url, 如 `http://blog.leanote.com/archives/用户名` 或 `http://demo.leanote.com/archives` 或 `http://demo.com/archives` |
| $.tagsUrl | 标签列表页url,  如 `http://blog.leanote.com/tags/用户名` 或 `http://demo.leanote.com/tags` 或 `http://demo.com/tags` |
| $.tagPostsUrl | 标签文章页url,  如 `http://blog.leanote.com/tag/用户名` 或 `http://demo.leanote.com/tag` 或 `http://demo.com/tag` |
| $.themeBaseUrl | 主题路径, 用于加载图片, css, js, 如 `/public/upload/123232/themes/32323` |

**注意**

上面的 `$.postUrl`, `$.searchUrl`, `$.singleUrl`, `$.cateUrl`, `$.tagPostsUrl`, `$.themeBaseUrl`是基地址, 还需要在后面加其它信息才能使用, 比如

* 查看一篇文章的链接为 `{{$.postUrl}}/相应的文章Id`
* 以某关键字进行搜索时 `{{$.searchUrl}}?keywords=相应的关键字`
* 查看某分类下的文章列表时 `{{$.cateUrl}}/相应的cateId`
* 查看某单页面时 `{{$.singleUrl}}/某单页面id`
* 查看某标签的文章列表时 `{{$.tagPostsUrl}}/某tag`
* 引用主题下style.css文件时 `{{$.themeBaseUrl}}/style.css`

leanote还定义了一些公用的静态文件url:

| 变量        | 描述           | 
| ------------- |:-------------:|
 $.blogCommonJsUrl | blog 公用 js 地址 |
| $.jQueryUrl | jQuery地址(1.9.0) |
| $.fontAwesomeUrl | font awsome 地址 |
| $.shareCommentCssUrl | 博客的分享与评论css地址 |
| $.shareCommentJsUrl | 博客的分享与评论js地址 |
| $.bootstrapCssUrl | bootstrap css 地址 |
| $.bootstrapJsUrl | bootstrap js 地址 |
| $.prettifyJsUrl | google code prettify js 地址 |
| $.prettifyCssUrl | google code prettify css 地址 |

在leanote的public/blog/js目录下含有以下js, 你可以使用`/public/blog/js/xx.js` 来加载, 如` <script src="/public/blog/js/jquery-cookie-min.js"></script>`

* bootstrap-dialog.min.js
* bootstrap-hover-dropdown.js
* jquery-cookie-min.js
* jquery.qrcode.min.js
* jsrender.js

若你还需要其它的js, css, 你可以在主题下新建js, css, 到时只需使用 `{{$.themeBaseUrl}}/你的.js或css` 来加载就行

### 页面判断
页面判断用于判断当前页是哪个页面, 比如是否在index,或cate页?

| 变量        | 描述           | 
| ------------- |:-------------:|
| $.curIsIndex| 是否在博客主页 |
| $.curIsCate | 是否在分类页|
| $.curIsPost | 否在文章详情页|
| $.curIsSearch | 是否在搜索页|
| $.curIsArchive | 是否在归档页|
| $.curIsSingle | 是否在单页|
| $.curIsTags | 是否在标签列表|
| $.curIsTagPosts | 是否在标签文章页|

比如, 你可以在`header.html`中通过判断页面来显示不同的标题:
```
<title>
{{if $.curIsIndex}}
    {{$.blogInfo.Title}}
{{else if $.curIsCate}}
    分类-{{$.curCateTitle}}
{{else if $.curIsSearch}}
    搜索-{{$.keywords}}
{{else if $.curIsTags}}
    我的标签
{{else if $.curIsTagPosts}}
    标签-{{$.curTag}}
{{else if $.curIsPost}}
    {{$.post.Title}}
{{else if $.curIsSingle}}
    {{$.single.Title}}
{{else if $.curIsArchive}}
    归档
{{end}}
</title>
```

### $.cates 分类列表

类型为`[]Cate`, `Cate`数据结构请查看下章内容

### $.singles 单页列表

类型为`[]Single`, `Single`数据结构请查看下章内容 , *注意* 这里的Single没有Content

### $.blogInfo 博客信息

类型为`BlogInfo`, `BlogInfo`数据结构请查看下章内容

### $.themeInfo 主题信息

类型为`ThemeInfo`, `ThemeInfo`数据结构请查看下章内容

### $.tags 标签集

类型为 `[]TagCount`, `TagCount` 数据结构请查看下章内容

### $.recentPosts 最近发表的5篇文章

类型为 `[]Post`, `Post` 数据结构请查看下章内容


## 数据结构

以下结构请注意首字大写!! 如你要在页面显示博客标题, 你需要这样:
`{{$.blogInfo.Title}}`

## BlogInfo 博客信息
```
{
    "UserId": "用户Id,即博客的作者的用户Id",
    "Username":    "用户名, "
    "UserLogo":    "用户的Logo, 包含http://",
    "Title":       "博客标题"
    "SubTitle":    "博客描述",
    "Logo":        "博客Logo, 包含http://",
    "OpenComment": true, // 是否开启评论
    "CommentType": "leanote" // 评论系统类型, 分为 leanote, or disqus
    "DisqusId": "leanote", // Disqus评论系统id
    "ThemeId":    "xxxxxxxx", // 主题Id
    "SubDomain":   "leanote的二级域名, 如demo",
    "Domain":     "http://demo.com 自定义域名"
}
```

### ThemeInfo 主题信息

以下包含最基本的, 是否还有其它字段, 这需要依据theme.json里的配置.
```
{
    "Name": "leanote-elegant-ing",
    "Desc": "一些描述, 或说明安装后还需要做的工作",
    "Version": "1.0",
    "Author": "leanote.com",
    "AuthorUrl": "http://leanote.com"
    // 可能还有其它配置...
|
```

### Cate 分类
```
{
    "CateId": "1232232",
    "Title": "分类标题",
    "UrlTitle": "友好的url"
}
```

### Post 文章
```
{
    "NoteId":      "1232323" // 一篇文件也是笔记, 所以使用NoteId作为主键
    "Title":      "标题",
    "UrlTitle":   "友好的Url",
    "ImgSrc":      "文章主图url地址, 包含http",
    "CreatedTime": time.Now(), // golang时间类型 创建时间
    "UpdatedTime": time.Now(), // golang时间类型 更新时间
    "PublicTime":  time.Now(), // golang时间类型 发布时间
    "Desc":        "文章的描述, 纯Text",
    "Abstract":    "文章的摘要, 有HTML",
    "Content":     "文章内容",
    "Tags":        []string{"标签1", "标签2"}, // 数组, 元素是标签名
    "CommentNum":  1232, // 评论数量
    "ReadNum":     32, // 阅读量
    "LikeNum":     33, // 赞的量
    "IsMarkdown":  false, // bool类型, 是否是markdown文章, 如果是, 那么Conent就是markdown内容, 不是html内容
}
```

### Archive 归档, 按年分类的所有文章
```
{
    "Year": 2012 // 年
    "Posts": []Post, // 文章列表, 数组, 每个元素是Post类型
    "MonthArchives": []MonthArchive // 按月的归档, MonthArchive类型见下
}
```

### MonthArchive 按月归档, 用于在类型Archive内使用
```
{
    "Month": 12 // 月
    "Posts": []Post, // 文章列表, 数组, 每个元素是Post类型
}
```

### Single 单页信息
```
{
    "SingleId":    "xxxxxxx", // id
    "Title":       "标签",
    "UrlTitle":    "友好的url",
    "Content":     "内容",
    "CreatedTime": time.Now(), // golang时间类型 创建时间
    "UpdatedTime": time.Now(), // golang时间类型 更新时间
}
```

### TagCount 标签统计
```
{
    "Tag": "标签名", 
    "Count": 32 // 文章数
}
```

## theme.json 主题配置

主题配置采用JSON格式

其中Name, Desc, Version, Author, AuthorUrl是必填项(注意首字大写)

你也可以定义其它的配置, 如FriendLinks, 在模板文件使用 $.themeInfo.FriendLinks来获取值

注意: 

1. JSON语法严格, 键必须用双引号, 最后不得有空','来结尾
2. 以下配置不能包含任何注释, 不然解析会出错!

```
{
    "Name": "leanote-elegant-ing",
    "Desc": "一些描述, 或说明安装后还需要做的的工作",
    "Version": "1.0",
    "Author": "leanote.com",
    "AuthorUrl": "http://leanote.com",
    "FriendLinks": [ 
        {"Title": "leanote", "Url": "http://leanote.com"} 
    ]
}
```

下面分析每个页面的其它变量.

## index.html 首页

| 变量        | 描述           | 
| ------------- |:-------------:|
| $.posts| 类型是 []Post, 文章列表|
| $.pagingBaseUrl | 分页基本url, 比如 `http://demo.leanote.com`, $.pagingBaseUrl 在cate.html, search.html, tag_posts.html都有, 只是值不一样|

## cate.html 分类页
| 变量        | 描述           | 
| ------------- |:-------------:|
| $.curCateTitle | 当前分类标题 |
| $.curCateId| 当前分类id |
| $.pagingBaseUrl | 分页基本url|
| $.posts| 类型是 []Post, 当前分类下的文章列表|

## post.html 文章详情页
| 变量        | 描述           | 
| ------------- |:-------------:|
| $.post | 类型是Post, 当前文章信息 |
| $.prePost| 类型是Post, 上一篇文章, 无Content信息|
| $.nextPost| 类型是Post, 下一篇文章, 无Content信息|

## search.html 搜索页
| 变量        | 描述           | 
| ------------- |:-------------:|
| $.keywords | 搜索的关键字 |
| $.pagingBaseUrl | 分页基本url|
| $.posts| 类型是 []Post, 文章列表|

## archive.html 归档页
| 变量        | 描述           | 
| ------------- |:-------------:|
| $.archives | 类型是 []Archive |
| $.curCateTitle | 当前分类标题, 如果有传cateId, 那么可按分类来显示归档 |
| $.curCateId| 当前分类id, 如果有传cateId, 那么可按分类来显示归档|

可以按年来显示文章, 也可以按月

按年:
```
<ul>
    {{range $.archives}}
    <li><span class="archive-year">{{.Year}}</span>
        <ul>
            {{range .Posts}}
            <li>
                {{dateFormat .PublicTime "2006-01-02"}} <a href="{{$.postUrl}}/{{.NoteId}}">{{.Title}}</a>
            </li>
            {{end}}
        </ul>
    </li>
    {{end}}
</ul>
```

按月:
```
<ul>
    {{range $.archives}}
    <li><span class="archive-year">{{.Year}}</span>
	    <ul>
            {{range .MonthAchives}}
	        <li>
	            <span class="archive-month">{{.Month}}</span>
	            <ul>
	                {{range .Posts}}
	                <li>
	                    {{dateFormat .PublicTime "2006-01-02"}} <a href="{{$.postUrl}}/{{.NoteId}}">{{.Title}}</a>
	                </li>
	                {{end}}
	            </ul>
	            
	        </li>
	        {{end}}
	    </ul>
    </li>
    {{end}}
</ul>	        
```

归档页还可以按分类, 年, 月来查询, 比如url: ```{{$.archiveUrl}}?year=2014&month=11&cateId=xxxxxx``` 表示查询分类为xxxxxx, 2014年11月的归档.

## single.html 单页面
| 变量        | 描述           | 
| ------------- |:-------------:|
| $.curSingleId | 当前单页id|
| $.single| 类型是 Single, 单页信息 |

## tags.html 标签列表页
| 变量        | 描述           | 
| ------------- |:-------------:|
| $.tags | 类型是[]TagCount|

## tag_posts.html 标签文章页
| 变量        | 描述           | 
| ------------- |:-------------:|
| $.curTag | 当前标签 |
| $.pagingBaseUrl | 分页基本url|
| $.posts| 类型是 []Post, 当前分类下的文章列表|

## 404.html 错误页
无特殊变量

## golang模板语法帮助

### 关于 `$` 和 `.`
`$` 表示根上下文, 推荐在每个变量前加`$`, 这样当上下文改变时还可以正确获取想要的值, 如`{{$.post.Title}}`.

而`.`表示当前的上下文, 上下文会因 {{range}}, {{for}}而改变, 比如在`index.html`页:
```
{{range $.posts}}
    {{.Title}} <!-- 输出每个post的标题, 这里的.就表示当前的post -->
{{end}}
```

### 模板引用

一个模板可以包含其它模板, 比如在 `index.html`中
`{{template "header.html" $}}`

**注意**

这里还需要传递上下文 `$`, 不然在`header.html`中不能引用变量.

### 函数函数
有两种方法可以调用函数, 1) 管道 2) 函数调用

推荐使用函数调用的方式, 比管道更容易理解.

golang官方文档上很详细:
```
{{printf "%q" "output"}}
	A function call. 函数调用
{{"output" | printf "%q"}}
	A function call whose final argument comes from the previous command
        printf有两个参数, 参数1为"%q", 参数2为"output". 相当于  {{printf "%q" "output"}}
{{printf "%q" (print "out" "put")}}
	A parenthesized argument. 使用()来嵌套函数调用

下面是更复杂的例子
{{"put" | printf "%s%s" "out" | printf "%q"}}
	A more elaborate call. 
{{"output" | printf "%s" | printf "%q"}}
	A longer chain.
```

1) 管道

`{{参数1| 函数名}}`

`{{$.post.CreatedTime | datetime}}`, 这里datetime是一个函数, `$.post.CreatedTime`作为其参数

2) 函数调用

`{{函数名 参数1 参数2 参数3}}`

上面的例子也可以像函数一样这样调用 
`{{datetime $.post.CreatedTime}}`

### if else 条件判断
if 是一个值, 这个值可以是变量(bool, string, object都行, 字符串不空为值, object不为nil为值, 值也可以是表达式.

如果值是变量, 在`header.html`中有一段经典的 if else 条件判断:
```
<title>
{{if $.curIsIndex}}
    {{$.blogInfo.Title}}
{{else if $.curIsCate}}
    分类-{{$.curCateTitle}}
{{else if $.curIsPost}}
    {{$.post.Title}}
{{else}}
    我的博客
{{end}}
```

如果值是表达式, 在`share_comment.html`中有一段复杂的例子:
```
{{if and $.blogInfo.OpenComment (eq $.blogInfo.CommentType "disqus")}}
{{end}}
```
golang模板可以使用`eq`, `lt`, `gt`之类的来作比较, 也有`and`, `or`, 需要注意的是 `eq`, `lt`, `gt`, `and`, `or`全是函数. 上面这个例子相当于伪代码:

`if ($.blogInfo.OpenComment && $.blogInfo.CommentType == "disqus")`

### leanote自带的自定义函数
* blogTags 输出标签, 如  `{{blogTags $ .Tags}}`
* datetime, date 日期时间格式化, 如 `{{$.post.UpdatedTime | datetime}} `  `{{$.post.UpdatedTime | date}} `
* dateFormat 自定义日期时间格式化, 如 ` {{dateFormat .PublicTime "2006-01-02"}}`
* raw 显示原生 html, 如`{{$.post.Content | raw}}`

### for, range 循环