本文介绍客户端如何展示/缓存笔记内的图片.

## 图片URL地址

图片在笔记内容的呈现方式有以下几种:

非Markdown笔记:

```
<img src="http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1">
<img src="http://leanote.com/api/file/getImage?fileId=5503537b38f4111dcb0000d1">
```
Markdown笔记:
```
![](http://leanote.com/file/outputImage?fileId=5503537b38f4111dcb0000d1)
![](http://leanote.com/api/file/getImage?fileId=5503537b38f4111dcb0000d1)
```

因为历史的原因, 图片的URL可能有以上两种, 又因为https的原因, 很可能图片的URL是https的. 所以 `http://leanote.com`  有可能是`https://leanote.com`, 又因为Leanote客户端支持自建服务, 所以域名可能会用户自定义的域名.

## 图片处理流程
笔记的内容最终要展示在webview里, 那么客户端需要如何处理笔记的图片呢, 怎么展示, 怎么缓存? 基本思路如下:

1. 自定义协议, 如leanote://, 用来处理图片请求.
2. 将笔记内的图片url替换成固定的url形式, 如替换成 `leanote://file/getImage?fileId=5503537b38f4111dcb0000d1`.
3. 在自定义协议里, 得到访问的图片请求, 得到fileId, 通过fileId去本地数据库查是否有该fileId的文件, 
	1.  如果没有, 则根据fileId去调用Api: /file/getImage, 将图片下载到本地, 并将图片路径和fileId对应起来存到本地数据库中.
	2. 如果存在, 则根据图片路径直接返回图片.

## 设计实现

所以,  本地数据库需要有一个表来存储fileId与本地图片路径的表. 表设计files:

* fileId (24位)
* path

伪代码:

### 自定义protocol
```javascript
initProtocol: function () {
	// 自定义leanote协议
	protocol.registerFileProtocol('leanote', function(request, callback) {
		// 解析url得到fileId
		var url = request.url;
		var ret = /fileId=([a-zA-Z0-9]{24})/.exec(url);
		if (ret && ret[1]) {
			// 得到fileId
			var fileId = ret[1];
			// 调用File服务, 获取图片
			FileService.getImage(fileId, function(fileLocalPath) {
				callback({path: fileLocalPath});
			});
		}
	});
}
```

### 图片服务, 下载图片, 获取本地图片

```javascript
getImage: function(fileId, callback) {
    var Api = require('api');

    // 访问api, 得到图片
    function getImageFromApi() {
        Api.getImage(fileId, function(fileLocalPath, filename) { 
            if(fileLocalPath) {
                // 保存到本地数据库中
                me.addImageForce(fileId, fileLocalPath, function(doc) {
                    callback(fileLocalPath);
                });
            }
        }); 
    }

    // 先查看本地是否有该文件, 如果有, 则直接返回, 没有, 则调用API
    me.getImageLocalPath(fileId, function(has, fileLocalPath) {
        // 本地数据库有记录
        if(fileLocalPath) {
            // 看本地是否存在该文件, 如果不存在, 还是要重新调用API获取
            fs.exists(fileLocalPath, function(exists) {
                if(exists) {
                    callback(fileLocalPath);
                } else {
                    getImageFromApi();
                }
            });
        // 不存在, 调用API获取
        } else {
            getImageFromApi();
        }
    });
}
```