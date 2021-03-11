[中文版](https://github.com/leanote/leanote/wiki/leanote-blog-theme-api)

The grammar of leanote blog template is based on the grammar of golang template.
As the grammar of golang template is brief, it is easy for you to master it. You can refer to leanote's theme template.

Here is some basic usage:

* print: `{{$.blogInfo.UserId}}` means print blogInfo.UserId variable, e.g.:`{{$.blogInfo.UserId}}`

* judge：`{{if $.blogInfo.OpenComment}}` some statement `{{else}}` other statement `{{end}}`

* range loop: `{{range $.posts}} {{.Title}} {{end}}` means loop through the posts and print the title of each post.

* call function: `{{$.post.CreatedTime|datetime}}` means formatting the time by calling `datetime` function. 

More information about golang template, please refer to golang template help page.


## Template organization structure:

Standard theme template of leanote is as follows:

(note: header.html, footer.html, paging.html, share_comment.html, highlight.html are not be required, for they are just referred by other template.)

* theme.json --theme configuration [must]
* header.html --header template, referred by other template
* footer.html --footer template, referred by other template
* index.html --index page [must]
* cate.html --category page [must]
* post.html --post detail page [must]
* archive.html -- archive page [must]
* single.html -- single page [must]
* share_comment.html --share and comment page , referred by post.html
* highlight.html --code highlight page, referred by others
* paging.html -- page html, referred by other template
* tags.html --tag list page [must]
* tag_posts.html --tag post page [must]
* 404.html --error page [must]
* style.css -- style
* images/ --image folder
* mages/screenshot.png --theme preview image

## Public variable

Public variable could be used by all pages, including some public urls, page judge variable, `$.cates` --category list, `$.singles` --single page list, `$.themeInfo` theme info; `$.blogInfo` blog info, `$.tags` tag, `$.recentPosts` recent 5 posts.

### Url address:
As leanote support second-level domains and custom domain, your url will be different based on you settings.
e.g. you could set a second-level domain like `http://demo.leanote.com` and a custom domain like `http://demo.com`.
Default leanote's blog domain is `http://blog.leanote.com`

| variable	|  description       |
| ------------- |:-------------:|
|$.siteUrl	|current site address, e.g. `http://leanote.com`, `http://localhost:9000` |
|$.indexUrl	|my blog index address, e.g. `http://blog.leanote.com/username` or `http://demo.leanote.com`, `http://demo.com`, The priority is From low to high, so if you have custom domain like `http://domain.com`, `$.indexUrl` is `http://demo.com`|
|$.cateUrl	|category page url, e.g. `http://blog.leanote.com/cate/username`, or `http://demo.leanote.com/cate`, or `http://demo.com/cate` |
|$.postUrl	|post detail url, e.g. 如 `http://blog.leanote.com/post/username`, or `http://demo.leanote.com/post`, or `http://demo.com/post`  |
|$.searchUrl	|search page url, e.g. `http://blog.leanote.com/search/username`, or `http://demo.leanote.com/search`, or `http://demo.com/search` |
|$.singleUrl	|single page url, e.g. `http://blog.leanote.com/single/username`, or `http://demo.leanote.com/single` or `http://demo.com/single` |
|$.archiveUrl  |archives page url, e.g. `http://blog.leanote.com/archives/username`, or `http://demo.leanote.com/archives`, or `http://demo.com/archives`| 
|$.tagsUrl	|tag list url, e.g. `http://blog.leanote.com/tags/username`, `http://demo.leanote.com/tags`, `http://demo.com/tags` |
|$.tagPostsUrl	|tag's posts url, e.g. `http://blog.leanote.com/tag/usename`, or `http://demo.leanote.com/tag`, or `http://demo.com/tag` | 
|$.themeBaseUrl	|the current theme's path, e.g. `/public/upload/123232/themes/32323`|


** note** 
The `$.postUrl`, `$.searchUrl`, `$.singleUrl`, `$.cateUrl`, `$.tagPostsUrl`, `$.themeBaseUrl` are all base address, you need to add other related information to use, e.g.:

* the url of a post is `{{$.postUrl}}/UrlTitle or NoteId`
* the url of a search page is `{{$.searchUrl}}?keywords=your_search_keywords`
* the url of posts list of a category is `{{$.cateUrl}}/UrlTitle or your_categoryId`
* the url of a single page is  `{{$.singleUrl}}/UrlTitle or single_page_id`
* the url of post list of a tag is   `{{$.tagPostsUrl}}/your_tag`
* the url of referred static file is   `{{$.themeBaseUrl}}/style.css`

Leanote has defined some public static file url:

|variable      |description       |
| ------------- |:-------------:|
|$.blogCommonJsUrl   |	blog public js address |
|$.jQueryUrl    | jQuery url(version:1.9.0) |
|$.fontAwesomeUrl   | font awesome address  |
|$.shareCommentCssUrl	|  share and comment css address |
|$.shareCommentJsUrl	|  share and comment js address  |
|$.bootstrapCssUrl	|   bootstrap css address  |
|$.bootstrapJsUrl	|  bootstrap js address  |
|$.prettifyJsUrl	|  google code prettify js address |
|$.prettifyCssUrl	|  google code prettify css address |


There are some js files under `public/blog/js` directory as follows, you can load them by referring to address like `/public/blog/js/xx.js` :
* bootstrap-dialog.min.js
* bootstrap-hover-dropdown.js
* jquery-cookie-min.js
* jquery.qrcode.min.js
* jsrender.js


If you still need other js and css file, you could new a file and load them by `{{$.themeBaseUrl}}/your_static_file`

### Page judge:
Page judge is used for judging what is the current page, like index page or category page.

|variable	|  description |
| ------------- |:-------------:|
|$.curlsIndex	| whether current page is the index page |
|$.curlsCate	| whether current page is the category page |
|$.curlsPost	| whether current page is the post detail page |
|$.curlsSearch	| whether current page is the search page |
|$.curlsArchive	| whether current page is the archive page |
|$.curlsSingle	| whether current page is the single page |
|$.curlsTags	| whether current page is the tag list page |
|$.curlsTagPosts| whether current page is the tag post page |

e.g. you could show different title by judging current page:
```
<title>
{{if $.curIsIndex}}
    {{$.blogInfo.Title}}
{{else if $.curIsCate}}
    category-{{$.curCateTitle}}
{{else if $.curIsSearch}}
   	search-{{$.keywords}}
{{else if $.curIsTags}}
   	my tag
{{else if $.curIsTagPosts}}
    tag-{{$.curTag}}
{{else if $.curIsPost}}
    {{$.post.Title}}
{{else if $.curIsSingle}}
    {{$.single.Title}}
{{else if $.curIsArchive}}
    archive
{{end}}
</title>
```

### $.cates --category list

data type is `[]Cate`, `Cate` data structure is in the following.

### $.singles  --single page list

Data type is `[]Single`, `Single` data structure is in the following , *note* this single has no content

### $.blogInfo --blog info

Data type is `BlogInfo`, `BlogInfo` data structure is in the following 

### $.themeInfo  --theme info

Data type is `ThemeInfo`, `ThemeInfo` data structure is in the following

### $.tags -- tag list

Data type is  `[]TagCount`, `TagCount` data structure is in the following

### $.recentPosts --recent 5 posts

data type is `[]Post`, `Post` data structure is in the following

## Data structure:

Please note that the following structure should begin with a capital letter, like {{$.blogInfo.Title}}

### BlogInfo
```
{
    "UserId": "user id",
    "Username":    "user name"
    "UserLogo":    "user logo, including http://",
    "Title":       "blog title"
    "SubTitle":    "blog description",
    "Logo":        "blog Logo, include http://",
    "OpenComment": true, // whether enabling comment function
    "CommentType": "leanote" // comment systme type, including leanote and disqus
    "DisqusId": "leanote", // id of disqus comment system
    "ThemeId":    "xxxxxxxx", // theme Id
    "SubDomain":   "second level domain, like demo",
    "Domain":     "custom domain, like http://demo.com"
}
```

### ThemeInfo

The following just contains the basic information, it depends on the configuration of theme.json.
```
{
    "Name": "leanote-elegant-theme",
    "Desc": "",
    "Version": "1.0",
    "Author": "leanote.com",
    "AuthorUrl": "http://leanote.com"
    // other configuration
}
```

### Cate (category)
```
{
    "CateId": "1232232",
    "Title": "title",
    "UrlTitle": "friendly url"
}
```

### Post
```
{
    "NoteId":      "1232323" // a post is also a note, so using noteId as the primary key
    "Title":      "title",
    "UrlTitle":   "friendly Url",
    "ImgSrc":      "post image url, including http",
    "CreatedTime": time.Now(), // golang time type, creating time
    "UpdatedTime": time.Now(), // updated time
    "PublicTime":  time.Now(), // publish time
    "Desc":        "description, pure Text",
    "Abstract":    "abstract, may contains HTML",
    "Content":     "content",
    "Tags":        []string{"tag1", "tag2"}, // array, element type is tag name
    "CommentNum":  1232, // comment number
    "ReadNum":     32, // read number
    "LikeNum":     33, // liked number
    "IsMarkdown":  false, // boolean type, whether a markdown post, if true, the content is markdown,or is html.
}
```

### Archive
```
{
    "Year": 2012 // year
    "Posts": []Post, // post list, array type, element type is Post
    "MonthArchives": []MonthArchive // archived by month, MonthArchive type is in the following.
}
```

### MonthArchive

Archived by month, which is used in Archive type.
```
{
    "Month": 12 // month
    "Posts": []Post, // post list, array type, element type is Post
}
```

### Single
```
{
    "SingleId":    "xxxxxxx", // id
    "Title":       "title",
    "UrlTitle":    "friendly url",
    "Content":     "content",
    "CreatedTime": time.Now(), // golang time type, creating time
    "UpdatedTime": time.Now(), // update time
}
```

### TagCount
```
{
    "Tag": "tag name", 
    "Count": 32 // post counts
}
```

## theme.json, theme configuration file

The configuration use json data type, note that the Name, Desc, Version, Author, AuthorUrl are required and these fields begin with capital character. Of course you could define other configuration as you like, such as FriendLinks and you could get value by using `$.themeInfo.FriendLinks`.

**Note**:
1. As the json grammar is strict, the key must be within double quotations and must not end with ',' 
2. The following configurations could not contain any comment, otherwise error would occur.
```
{
    "Name": "leanote-elegant-theme",
    "Desc": "some descriptions or some work should be done after installation",
    "Version": "1.0",
    "Author": "leanote.com",
    "AuthorUrl": "http://leanote.com",
    "FriendLinks": [ 
        {"Title": "leanote", "Url": "http://leanote.com"} 
    ]
}
```

Next we will analyze other special variables in each page in detail.

### index.html
|variables    |           desc|
| ------------- |:-------------:|
|$.posts       |          data type is []Post, post list|
|$.pagingBaseUrl     |    paging base url, like `http://demo.leanote.com` , `$.pagingBaseUrl` exists in cate.html, search.html, tag_posts.html, but the values are different.|


### cate.html
|variables      |         desc |
| ------------- |:-------------:|
|$.curCateTitle  |       current category title |
|$.curCateId     |        current category id |
|$.pagingBaseUrl  |       paging base url |
|$.posts          |       data type is []Post, post list of the current category |


### post.html 
|variables         |      desc |
| ------------- |:-------------:|
|$.post         | data type is Post, current post information |
|$.prePost      | data type is Post, the last post, no content |
|$.nextPost     | data type is Post, the next post, no content |


### search.html
|variables     |          desc |
| ------------- |:-------------:|
|$.keywords     |         search keyword |
|$.pagingBaseUrl |        paging base url |
|$.posts         |       data type is []Post, post list |

### archive.html
|variables         |      desc  |
| ------------- |:-------------:|
|$.archives     |        data type is []Archive |
|$.curCateTitle      |    current category title, if passing cateId, the archives could be shown by category |
|$.curCateId     |        current category id, if passing cateId, the archives could be shown by category|

The posts could be shown by year or by month.

By year:
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

By month:
```
<ul>
    {{range $.archives}}
    <li><span class="archive-year">{{.Year}}</span>
        <ul>
            {{range .MonthArchives}}
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

The archive page could also be queried by category, year or month. Such as url: `{{$.archiveUrl}}?year=2014&month=11&cateId=xxxxxx`, it query the archives that the category is xxxxxx and the date is November 2014.

### single.html
| variables     |          desc |
| ------------- |:-------------:|
|$.curSingleId   |        current single page id |
|$.single        |       data type is Single, single page information |


### tags.html
|variables      |         desc |
| ------------- |:-------------:|
|$.tags         |        data type is []TagCount |


### tag_posts.html
| variables       |        desc |
| ------------- |:-------------:|
|$.curTag       |        current tag |
|$.pagingBaseUrl  |       current paging base url |
|$.posts          |       data type is []Post, post list of the current category |


### 404.html   error page
no special variable

## Golang template grammar help:

### about $ and .

`$` means the root context, a variable begins with `$` is recommended. By using `$`, you could still fetch the value when the context changes, like {{$.post.Title}}

`.` means current context, context would change when using {{range}} or {{for}} , for example, in index.html:
```
{{range $.posts}}
    {{.Title}} <!-- print each post's title, the . here means the current post -->
{{end}}
```

### Template reference
A template could contain other templates, for example, {{template "header.html" $}} in index.html

**Note**:
the context $ is necessary here, otherwise you could not refer to variables in header.html

## Function:
There are two methods to call function:

1) pipeline  
2) call function

call function is recommended which is easy to understand.

the golang official doc has described it in detail:
```
{{printf "%q" "output"}}
    A function call. 
{{"output" | printf "%q"}}
    A function call whose final argument comes from the previous command
{{printf "%q" (print "out" "put")}}
    A parenthesized argument.
```

More complicated example:
```
{{"put" | printf "%s%s" "out" | printf "%q"}}
    A more elaborate call. 
{{"output" | printf "%s" | printf "%q"}}
    A longer chain.
```

1) pipeline:

`{{param1| function name}}`

`{{$.post.CreatedTime | datetime}}`, the datetime is a function and $.post.CreatedTime act as the function's parameter.

2) calling function

`{{function_name params1 params2 params3}}`

The previous example could also be implemented by `{{datetime $.post.CreatedTime}}`


## if else condition judgement

The value after if could be a variable like bool, string(not nil), object(not nil) or expression

If the value is a variable , there is a classic if else condition judgement in header.html:
```
<title>
{{if $.curIsIndex}}
    {{$.blogInfo.Title}}
{{else if $.curIsCate}}
    category-{{$.curCateTitle}}
{{else if $.curIsPost}}
    {{$.post.Title}}
{{else}}
    my blog
{{end}}
```

If the value is expression, there is a complex example in share_comment.html:
```
{{if and $.blogInfo.OpenComment (eq $.blogInfo.CommentType "disqus")}}
{{end}}
```

Golang template could use `eq`, `lt`, `gt`, `and` , `or`, please note that `eq`, `lt`, `gt`, `and`, `or` are all function.

The previous example is equals to pseudocode as follows:
```
if ($.blogInfo.OpenComment && $.blogInfo.CommentType == "disqus")
```

## leanote build-in custom function
* blogTags, print tags,  e.g. `{{blogTags $ .Tags}}`
* datetime, date, time format, e.g. `{{$.post.UpdatedTime | datetime}}`, `{{$.post.UpdatedTime | date}}`
* dateFormat, custom datetime format, e.g. `{{dateFormat .PublicTime "2006-01-02"}}`
* raw, shows raw html, e.g. `{{$.post.Content | raw}}`