# Leanote同步机制

Leanote同步机制参考Evernote的机制, 关于Evernote的同步机制参考: http://dev.evernote.com/media/pdf/edam-sync.pdf

需要用到[Leanote API](https://github.com/leanote/leanote/wiki/leanote-api)

## 前言
Leanote主要由Notebook, Note, Tag, File(图片/附件)组成. File依附于Note存在. 当Note删除时, 其包含的File也会删除.

每个帐户(User), Notebook, Note, Tag都含有字段 Usn(Update Sequence Number), 它是整个同步系统中最重要的字段. User的Usn用于标识账户中的每一次修改, 每次修改Notebook, Note, Tag后User的Usn就会+1. 而Notebook, Note, Tag的Usn标识着一个对象最后一次被修改时的账户Usn.

举个例子, 在某一个时刻User的Usn是100. 我添加一个笔记Note1, 那么User的USN会变成101, 此时该Note1的USN也是101. 然后我再添加一个笔记Note2，这时User的Usn会变成102，Note2的Usn也是102，Note1的还是101. 这样一来我们每次同步后记录一下当时User的Usn保存为LastUSN, 下次同步的时候如果账户的Usn > LastUsn，说明账户中有东西被修改了, 此时需要先将服务器端的修改同步到本地.

当帐户第一次登录时, 此时需要进行一次全量同步, 即将服务器上所有和数据都同步到本地. 而之后用户在本地操作后, 就需要每次同步所修改的数据.

同步基本的步骤如下:

1. Pull: 判断服务端是否有新数据, 即 通过 本地LastSyncUsn 和 服务器端Usn对比, 如果本地LastSyncUsn < 服务器端Usn, 表示服务端有修改, 此时需要同步服务器上的数据到本地. 详情请见 "同步数据".
2. Push: 将本地修改的数据发送到服务器端. 详情请见 "发送改变".
3. 保存状态: 获取最新同步状态, 保存服务器端最新的Usn为本地LastSyncUsn.

## 同步数据 Pull
从服务器端同步数据到本地.

先判断服务端是否有新数据, 即 通过 本地LastSyncUsn 和 服务器端Usn对比, 如果本地LoastSyncUsn < 服务器端Usn, 表示服务端有修改, 此时需要同步服务器上的数据到本地.

同步数据步骤:

1. 同步Notebook
2. 同步Note
3. 同步Tag

同步Notebook, Note, Tag的步骤基本一致, 现拿同步Notebook作为示例, 伪代码为:

```javascript
// 获取远程要同步的数据
var lastSyncUsn = getLastSyncUsn(); // 本地保存的上次同步的Usn
function syncNotebook(lastSyncUsn) {
	var afterUsn = lastSyncUsn; // 表示取lastSyncUsn之后的notebook
	while(true) {
		// 调用api, 取afterUsn的10个笔记本
		var notebooks = api.call('/api/notebook/getSyncNotebooks?afterUsn=afterUsn&maxEntry=10');
		// 将获取到的notebooks存到本地
		updateNotebookToLocal(notebooks);

		// 如果取到的notebook == 10, 表示很可能还有要同步的notebook
		if(notebooks.length == 10) {
			afterUsn = notebooks[notebooks.length-1].Usn; // 取最大的Usn作为下一个标准
		}
		// 如果 < 则表示不够了, 没有要同步的Notebook了.
		else {
			break;
		}
	}
}

// 将远程数据保存到本地
function updateNotebookToLocal(notebooks) {
	for(var i = 0; i < notebooks.length; ++i) {
		var notebook = notebooks[i];
		// 获取本地的Notebook
		var localNotebook = getLocalNotebook(notebook.NotebookId);

		// 服务器端已删除了, 此时删除本地的
		if(notebook.IsDeleted) {
			deleteLocalNotebook(notebook.NotebookId);
		}
		else {
			// 如果本地没有修改, 那么将notebook保存到本地
			if(!localNotebook.IsDirty) {
				db.updateToLocal(notebook.NotebookId, notebook);
			}
			// 本地有更新, 此时需要处理冲突
			else {

			}
		}
	}
}
```
获取Note, Tag要同步的数据的API为

* /api/note/getSyncNotes
* /api/tag/getSyncTags

具体用法请参考: [Leanote API](https://github.com/leanote/leanote/wiki/leanote-api)

### 如何处理冲突 ?
冲突发生的原因: 本地修改了, 且服务器上也修改了. 此时同步服务器上的数据到本地, 发送本地数据的IsDirty=true. 此时需要处理冲突.

处理冲突由客户端来完成, 最极端的做法是: 客户端可以完全将服务器上的数据覆盖到本地, 或者完全舍弃服务器端的数据而使用本地修改的数据.

但这样做很可能会丢失数据, 所以当遇到冲突时, 应该将服务器上的数据下载到本地和本地冲突的数据进行关联, 最后采用哪个数据由用户来决定.

## 发送改变 Push
将本地修改的数据发送到服务器端.

发送改变步骤:

1. 发送修改的Notebook
2. 发送修改的Note
3. 发送修改的Tag

发送改变, 即得到本地修改过的Notebook, Note, Tag, 然后将修改后的信息发送到服务器端. 所以本地需要有一个标识来识别哪些数据改变了. 比如可以设置一个IsDirty的字段来标识. 如果本地的Note修改了, 更新该Note的IsDirty为true, 待发送改变成功后, 设其IsDirty=false.

下面通过发送Notebook改变作为例子:
```javascript
function sendNotebookChanges() {
	var dirtyNotebooks = getDirtyNotebooks();
	for(var i = 0; i < dirtNotebooks.length; ++i) {
		var dirtyNotebook = dirtyNotebooks[i];
		// 调用api, 发送改变, 必须要传usn, 服务器端根据传过去的usn来判断是否冲突
		var ret = api.call('/api/notebook/updateNotebook?usn=dirtyNotebook.Usn&title=dirtyNotebook.Title');
		// 修改成功, 将服务器端返回的Usn更新到本地, IsDirty设为false
		if(ret.Ok) {
			updateLocalNotebook(Notebook.Id, {Usn: ret.Usn, IsDirty: false});
		}
		// 更新失败, 有冲突, 表示服务器上的数据新于本地, 此时需要解决冲突, 
		// 解决冲突的方法可以将服务器的数据覆盖到本地
		else if(ret.Msg == "conflict") {
			var serverNote = apil.call('/api/notebook/getNotebook?notebookId=dirtyNotebook.id');
			updateNotebookToLocal(serverNote);
		}
	}
}
```

修改Note, Tag的api为:

* /api/note/updateNote, deleteNote
* /api/tag/addTag, deleteTag

具体用法请参考: [Leanote API](https://github.com/leanote/leanote/wiki/leanote-api)

## 获取最新同步状态
调用API "/api/user/getSyncState" 获取最新同步状态, 将Usn保存到本地为LastSyncUsn;
 