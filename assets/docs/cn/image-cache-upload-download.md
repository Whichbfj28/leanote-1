关于图片的缓存请查看 [客户端图片缓存处理](https://github.com/leanote/leanote/wiki/%E5%AE%A2%E6%88%B7%E7%AB%AF%E5%9B%BE%E7%89%87%E7%BC%93%E5%AD%98%E5%A4%84%E7%90%86). 本文介绍客户端如何下载与上传文件(图片与附件)

文件是指图片和附件.

设计一张表 files 记录本地缓存的文件, 字段如下:

* localFileId 本地文件id
* serverFileId 服务器端文件id
* path 本地文件路径
* isAttach 是否是附件
* type 文件类型(png, jpg, doc)

## 下载文件

文件的链接可能为:

* 图片: http://leanote.com/file/outputImage?fileId=24位id [以后会废弃, 因为历史问题, 还存在]
* 图片: http://leanote.com/api/file/getImage?fileId=24位id
* 附件: http://leanote.com/api/file/getAttach?fileId=24位id
* 附件: http://leanote.com/attach/download?attachId=24位id [以后会废弃, 因为历史问题, 还存在]

根据fileId, 调用相应api(file/getImage, file/getAttach)将文件下载到本地后, 需要在files表中插入一条记录, 其中serverFileId为服务器端文件Id.

## 上传文件

图片与附件依附于笔记, 所以在add/update note时, 应当要将笔记所带的图片和附件也加上. 但是已经上传过的图片就没必要再上传了, 所以客户端必须要记录哪些文件是已经上传过了的.

对于本地新添加的文件只有localeFileId, 没有serverFileId. 所以, 没有serverFileId的文件必须要将文件的数据一并提交, 而有serverFileId的文件则只需要附带文件的元数据即可.

所以添加/修改笔记时, 需要附上: 所有文件的元数据 + 本地新添加的文件的文件数据.

提交的数据示例如下:

POST的form数据
```
{
    NoteId: "22bab16c38f41127c2000075",
    Title: "new title",
    IsBlog: false,
    IsMarkdown: true,
    Content: "content",
    NotebookId: "55bab16c38f41127c2000075",
    CreatedTime: "2015-11-11 21:28:11",
    UpdatedTime: "2015-11-11 23:28:11",
    Tags: ['a','b],
    // 以下是文件的元数据, 是一个数组
    Files: [
        {
            LocalFileId: "56435e7d19637e699b00000c",
            FileId: "", // 服务端还没有, 则需要提交文件数据
            HasBody: true, // 表示会传数据
            IsAttach: true,
            Type: "js"
        },
        {
            LocalFileId: "56435e7d19637e699b00000d",
            FileId: "", // 服务端还没有, 则需要提交文件数据
            HasBody: true, // 表示会传数据
            IsAttach: false,
            Type: "png"
        },
        {
            LocalFileId: "56435e7d19637e699b00000d",
            FileId: "32435e7d19637e699b00000d", // 服务端id已经存在
            HasBody: false, // 表示不会传数据
            IsAttach: false,
            Type: "png"
        }
    ]
}
```
POST的multipart form data 文件的数据, name是FilesData[localeFileId]
```
{
    "FileDatas['56435e7d19637e699b00000c']": {
        content_type: "application/js",
        file: "/data/1447255677791.js", // 文件路径
        filename: "a.js"
    },

    "FileDatas['56435e7d19637e699b00000d']": {
        content_type: "image/png",
        file: "/data/1447255677790.png", // 文件路径
        filename: "b.png"
    }
}
```
最终的数据传输形式为: `Title=title&Tags[0]=a&Tags[1]=b&Files[0][LocalFileId]=56435e7d19637e699b00000d&Files[0][HasBody]=true`, 对于multipart的数据上传name为FilesData[localeFileId], 其它的组织形式示语言的实现而定.

笔记添加/更新成功后, 会返回数据, 也会将Files数据返回, 其中FileId已经是服务器端FileId, 此时需要将该FileId保存至ServerFileId建议本地LocalFileId与ServerFileId的关联. 下次就不需要再传该文件的数据了.

## 笔记内容内的图片,附件链接
调用add/update note的api时, 需要将笔记内的图片, 附件的链接转成标准的leanote链接:

* 图片: http://leanote.com/api/file/getImage?fileId=24位本地LocalFileId 或 serverFileId
* 附件: http://leanote.com/api/file/getFile?fileId=24位本地LocalFileId 或 serverFileId

其中 http://leanote.com 可以为自建服务地址.

服务端会将根据传过来的Files元数据, 将24位本地LocalFileId替换成服务端的FileId 再保存到数据库中, 如果fileId为serverFileId, 则因为在Files找不到映射, 所以不会替换直接保存.