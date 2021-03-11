# Leanote API v1

## Preface

### api url

All api url are start with /api/, e.g:

`/api/user/info?userId=xxxx&token=xxxx`

Except for  `/auth/login`, `/auth/register`, all other url must include param `token=xxxx`

### Flow
1. Call `/auth/login` and get token
2. Call other api with token 

## Return result structure
* All result is JSON besides the binary file (image, attachment), if the result is not JSON, it must be wrong.
* All error result are formatted as {Ok: false, Msg: "error msg"}
* The correct result has two different formats:
1. For operation api, such as update notebook, update note, the result is {Ok: true, Msg:""}
2. For get data api, such as get note, get notebook, the result is the data, such as get note result data:
```
{
  "NoteId": "54bdc7e305fcd13ea3000000",
  "NotebookId": "54bdc65599c37b0da9000003",
  "UserId": "54bdc65599c37b0da9000002",
  "Title": "Note title",
  "Desc": "",
  "Tags": null,
  "Abstract": "",
  "Content": "",
  "IsMarkdown": false,
  "IsBlog": false,
  "IsTrash": true,
  "IsDeleted": false,
  "Usn": 15,
  "Files": [],
  "CreatedTime": "2015-01-20T11:13:41.34+08:00",
  "UpdatedTime": "2015-01-20T11:13:41.34+08:00",
  "PublicTime": "0001-01-01T00:00:00Z"
}
```

* All the date result are formatted as "2015-01-20T11:13:41.34+08:00"
* All Key's first letter are in capitals

-----------------

## Auth Login & Register

### Login: /auth/login 

```
Params: email, pwd
Method: GET
Result: 
Error: {"Ok":false, "Msg":"pwdError"}
Success:
{
    "Ok":true,
    "Token":"5500830738f41138e90003232",
    "UserId":"52d26b4e99c37b609a000001",
    "Email":"leanote@leanote.com",
    "Username":"leanote"
}
```
### Logout: /auth/logout 
```
Params: token
Method: GET
Return:
Error: {Ok: false, Msg: ""}
Success: {Ok: true, Msg: ""}
```

### Register: /auth/register
```
Params: email, pwd
Method: POST
Return:
Error: {Ok: false, Msg: ""}
Success: {Ok: true, Msg: ""}
```

## User
### Get user's info: /user/info
```
Params: userId
Method: GET
Return:
Error: {Ok: false, Msg: ""}
Success: type.User
```

### Update username: /user/updateUsername
```
Params: username
Method: POST
Return: 
Error: {Ok: false, Msg: ""}
Success: {Ok: true, Msg: ""}
```

### Update password /user/updatePwd
```
Params: oldPwd(old password), pwd(new password)
Method: POST
Return:
Error: {Ok: false, Msg: ""}
Success: {Ok: true, Msg: ""}
```

### Update logo: /user/updateLogo
```
Params: file(avatar file)
Method: POST
Return:
Error: {Ok: false, Msg: ""}
Success: {Ok: true, Msg: ""}
```

### Get user's latest sync state: /user/getSyncState
```
Params: no params
Method: GET
Return:
Error: {Ok: false, Msg: ""}
Success: {LastSyncUsn: 3232, LastSyncTime: "2015-01-20T11:13:41.34+08:00"}
```

-----

## Notebook

### Get notebooks which need be synced: /notebook/getSyncNotebooks 
```
Params: afterUsn(int, the usn bigger than it is need be synced), maxEntry(int)
Method: GET
Return: 
Error: {Ok: false, Msg: ""}
Success: [type.Notebook] Array
```

### Get all notebooks: /notebook/getNotebooks 
```
No params
Method: GET
Return: 
Error: {Ok: false, Msg: ""}
Success: [type.Notebook] Array
```

### Create notebook: /notebook/addNotebook
```
Params: title(string), parentNotebookId(string, parent notebookId, optional), seq(int sequence)
Method: POST
Return: 
Error: {Ok: false, Msg:""}
Success: type.Notebook
```

### Update notebook: /notebook/updateNotebook
```
Params: notebookId, title, parentNotebookId, seq(int), usn(int)
Method: POST
Return: 
Error: {Ok: false, msg: ""} msg == "conflict" means the notebook posted is conflicted with server
Success: type.Notebook
```

### Delete notebook: /notebook/deleteNotebook
```
Params: notebookId, usn(int)
Method: POST
Return: 
Error: {Ok: false, msg: ""} msg == "conflict"
Success: {Ok: true}
```

----

## Note

### Get notes which need be synced: /note/getSyncNotes
```
Params: afterUsn(int,  the usn bigger than it is need be synced), maxEntry(int)
Method: GET
Return: 
Error: {Ok: false, Msg: ""}
Success: [type.Note] Array, Exclude Abstract and Content
```


### Get notebook's notes (Exclude content) /note/getNotes
```
Params: notebookId
Method: GET
Return: 
Error: {Ok: false, Msg: ""}
Success: [type.Note] Array, Exclude Abstract and Content
```

### Get note and content: /note/getNoteAndContent
```
Params: noteId
Method: GET
Return: 
Error: {Ok: false, Msg: ""}
Success: type.Note
```

### Get note's content: /note/getNoteContent
```
Params: noteId
Method: GET
Return: 
Error: {Ok: false, Msg: ""}
Success: type.NoteContent
```

### Create note: /note/addNote
```
Params: (Note the first letter is capitalized)
    NotebookId string required
    Title string required
    Tags []string optional
    Content string required
    Abstract string optional
    IsMarkdown bool optional
    CreatedTime string optional, e.g: 2012-12-01 12:32:11
    UpdatedTime string optional, e.g: 2012-12-01 12:32:11
    Files []type.NoteFiles Array optional
Method: POST
Return: 
Error: {Ok: false, Msg: ""}
Success: type.Note, Exclude Abstract and Content
```

**About note's image & attachment**

Please see: [客户端图片, 附件下载与上传](https://github.com/leanote/leanote/wiki/%E5%AE%A2%E6%88%B7%E7%AB%AF%E5%9B%BE%E7%89%87,-%E9%99%84%E4%BB%B6%E4%B8%8B%E8%BD%BD%E4%B8%8E%E4%B8%8A%E4%BC%A0)

### Update note: /note/updateNote

You just pass the param that has been modified, so if the titled is modified and the content is not modified, you don't need pass the content.

```
Params: (Note the first letter is capitalized)
    NoteId string required
    Usn int required
    NotebookId string optional
    Title string optional
    Tags []string optional
    Content string optional
    Abstract string optional
    IsMarkdown bool optional
    IsTrash bool optional
    UpdatedTime string optional e.g: 2012-12-01 12:32:11
    Files []type.NoteFiles optional
Method: POST
Return: 
Error: {Ok: false, msg: ''} msg == 'conflict' means conflicted!!
Success: type.Note, Exclude Abstract and Content
```

### Delete note permanently: /note/deleteTrash
```
Params: noteId, usn
Method: POST
Return: 
Error: {Ok: false, msg: ''} msg == 'conflict' means conflicted!!
Success: type.UpdateRet
```

-------

## Tag

### /tag/getSyncTags
```
Params: afterUsn, maxEntry
Method: GET
Return: 
Error: {Ok: false, Msg: ""}
Success: [type.Tag] Array
```

### /tag/addTag
```
Params: tag(string)
Method: POST
Return: 
Error: {Ok: false, Msg: ""}
Success: type.Tag
```

### /tag/deleteTag
```
Params: tag(string)
Method: POST
Return: 
Error: {Ok: false, Msg: ""}
Success: type.UpdateRet
```

### File (image, attachment)

### /file/getImage
```
Params: fileId
Method: GET
Return: 
Error: Not binary data
Success: binary data
```

### /file/getAttach
```
Params: fileId
Method: GET
Return: 
Error: Not binary data
Success: binary data
```

### /file/getAllAttachs
```
Params: noteId
Method: GET
Return: 
Error: Not binary data
Success: binary data
```

--------

## Data structure

### type.User

```
User {
    UserId  string
    Username string
    Email string
    Verified bool
    Logo string
}
```

### type.Notebook

```
Notebook {
    NotebookId        
    UserId           
    ParentNotebookId 
    Seq              int
    Title            string 
    IsBlog           bool  
    IsDeleted    bool
    CreatedTime      time.Time   
    UpdatedTime      time.Time 
    
    Usn int  // UpdateSequenceNum 
}
```

### type.Note
```
Note {
    NoteId     string
    NotebookId string
    UserId     string
    Title      string
    Tags       []string
    Content    string
    IsMarkdown bool
    IsBlog     bool 
    IsTrash bool
    Files []NoteFile // image, attachment
    CreatedTime time.Time
    UpdatedTime time.Time
    PublicTime time.Time
    
    Usn int
}
```

### type.NoteContent
```
NoteContent {
    NoteId string
    UserId string
    Content string
}
```

### type.NoteFile
```
NoteFile {
    FileId string // server file id
    LocalFileId string // client local file id
    Type string // images/png, doc, xls
    Title string
    HasBody bool // If it's true, the file must be posted
    IsAttach bool //
}
```

### type.Tag
```
Tag {
    TagId string
    UserId string
    Tag string 
    CreatedTime
    UpdatedTime
    IsDeleted bool
    Usn 
}
```

### type.UpdateRe
```
ReUpdate {
    Ok bool
    Msg string
    
    Usn int
}
```