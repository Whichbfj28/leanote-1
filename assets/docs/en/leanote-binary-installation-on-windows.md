**This tutorial explains doing a *binary* installation on Windows systems.**

-------
# Installation overview:

1. Download the binary file of `leanote`.
2. Install the database -- `mongodb`.
3. Import initial data to `mongodb`.
4. Configure `leanote`.
5. Run `leanote`.

---------
## 1. Download the binary file of `leanote`

Choose and download the binary file corresponding to your system from [here](https://github.com/coocn-cn/leanote/releases/latest).

Suppose it is saved in `C:\user1`, extract it there to get the installation folder `C:\users1\leanote`.


-------
## 2. Install the database -- `mongodb`

### 2.1 Install `mongodb`

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image006.png)

Download the Windows installer of `mongodb` from the official site of [mongo](https://www.mongodb.com/download-center#community). Launch the installer and follow through the process. Choose default or customize your setups. (See following screenshots.)

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image008.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image009.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image010.png)

Click "Finish" to complete the installation.

Create a new folder `dbanote` in your `C` drive to store `leanote`'s data.

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image011.png)

### 2.2 Test the installation of `mongodb`

Press `Win + R`, enter `cmd` to open a command line session, type in the following (NOT including `C:\>`):
```
C:\>mongod --dbpath C:\dbanote 
```
to start the database, see the screenshot below.

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image012.png)

####**★NOTE: you can minimize this window for now, but DO NOT CLOSE IT!** 
Now open up another command line window （`Win+R`, then `cmd`）, type in `mongo` to enter the interactive session of `mongo` (NOT including `C:\>`):
```
C:\> mongo
```
A leading `>` indicates you are in the `mongo` interactive mode. How type in `show dbs` to show the databases on system:
```
> show dbs
```

If you see something like the screenshot below, the installation of `mongodb` is successful：

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image014.png)



-------
## 3. Import initial data into `mongodb`

Type in `quit()` to exit from the previous `mongo` interactive session. Then copy n paste the command below into your command line. Note the difference between version 2 and 3 of `mongodb`:

- For mongodb v2:
```
mongorestore  -h localhost -d leanote  --directoryperdb C:\user1\leanote\mongodb_backup\leanote_install_data
```

- For mongodb v3:
```
mongorestore -h localhost -d leanote --dir C:\user1\leanote\mongodb_backup\leanote_install_data
```

The screenshot below shows the import process:

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image017.png)

To test the import, type in `mongo` in the same command line session, then `show dbs` to show the databases:

```
C:\> mongo
> show dbs          # show all databases
admin    (empty)
leanote  0.078GB        # Imported Leanote data
local    0.078GB 
```

**Note that the imported database contains 2 user accounts by default**
```
user1 username: admin, password: abc123 (administrator, used for backend management and control)  
user2 username: demo@leanote.com, password: demo@leanote.com (for demonstration purposes)
```

-------
## 4. Configure `leanote`

The configurations of `leanote` is stored in the file `conf/app.conf`. 

One setting that you are **strongly suggested** to modify is `app.secret`, please change arbitrary number of digits of the string to something different, but keeping the string length unchanged. This is to avoid potential security issues.

Other settings can remain as they are for now. For some database related settings, see the [FAQ page](https://github.com/leanote/leanote/wiki/Leanote-QA-English).


-------
## 5. Run `leanote`

Open up a command line session using **administrator privilege**, then run the following (NOT including `C:\>`):
```
C:\> cd C:\users\leanote\bin
C:\> run.bat
```

Printings similar to this indicates `leanote` has started successfully:
```
...
TRACE 2013/06/06 15:01:27 watcher.go:72: Watching: /home/life/leanote/bin/src/github.com/leanote/leanote/conf/routes
Go to /@tests to run the tests.
Listening on :9000...
```
Congratulations, now fire up you browser and enter `http://localhost:9000` into the address bar. Voilà! Welcome to `leanote` and happy note-taking!


------------
If you encounter issues or want to know more about `leanote`'s configurations, refer to the [FAQ page](https://github.com/leanote/leanote/wiki/Leanote-QA-English).

