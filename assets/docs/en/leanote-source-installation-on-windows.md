**This tutorial explains doing a *source code* installation on Windows systems.**

`leanote` source distribution is suitable for developers or those who need to use the new features sooner.

- For **binary** installation on Windows, see [here](https://github.com/leanote/leanote/wiki/leanote-binary-installation-on-Windows-(En)).
- For **source** installation on Mac and Linux, see [here](https://github.com/leanote/leanote/wiki/Leanote-source-installation-on-Mac-and-Linux-(En)).
- For **binary** installation on Mac and Linux, see [here](https://github.com/leanote/leanote/wiki/leanote-binary-installation-on-Mac-and-Linux-(En)). 


-------------
# Installation overview:

1. Download necessary files.
2. Install `golang`.
3. Install the database -- `mongodb`.
4. Install `leanote`.
5. Import initial data of `leanote`.
6. Configure `leanote`.
7. Run `leanote`.

**NOTE: For compatibility issues, please try following the instructions stringently, most importantly: use 32 bit installers and put the source codes to `C` drive. Other configurations can be customized according to your needs.**

----------
## 1. Download necessary files (32 bit)

* Download `golang` (1.7+): http://www.golangtc.com/static/go/1.8/go1.8.windows-386.msi
* Download `mongodb`: https://fastdl.mongodb.org/win32/mongodb-win32-i386-2.6.8-signed.msi?_ga=1.163324924.1783433278.1426342651  
* Download `leanote-all` (choose the version suitable for your system): https://github.com/leanote/leanote-all/archive/master.zip  

--------------
## 2. Install `golang`

Follow through the installation procedure using the `golang` installer (see screenshot below).

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image001.png)

Use default or customize your installation. After installation, press `Win + R`, then type in `cmd` to launch a command line session. Type in `go version` in the command line. If you see the following output, the `golang` is installed successfully.

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image004.png)

Now add the environmental variables `GOPATH` and `GOROOT` for `go`. **NOTE the difference between `GOPATH` and `GOROOT` !**:


Right click on "My computer" -- "Properties" -- "Advanced" -- "Environment variables" (see screenshot below)

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image005.png)


-------------------
## 3. Install the database -- `mongodb`

### 3.1 Install `mongodb`

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image006.png)

Download the Windows installer of `mongodb` from the official site of [mongo](https://www.mongodb.com/download-center#community). Launch the installer and follow through the process. Choose default or customize your setups. (See following screenshots.)

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image008.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image009.png)
![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image010.png)

Click "Finish" to complete the installation.

Create a new folder `dbanote` in your `C` drive to store `leanote`'s data.

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image011.png)


### 3.2 Test the installation of `mongodb`

Press `Win + R`, enter `cmd` to open a command line session, type in the following (NOT including `C:\>`) to start the database, see the screenshot below.:
```
C:\>mongod --dbpath C:\dbanote 
```


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

----------------------
## 4. Install `leanote`

Extract the downloaded `leanote-all` package to any folder, navigate into the folder and copy the `src` subfolder to `C:\GO\` (see screenshot below). If prompted whether to overwrite or not, click Yes.

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image016.png)

-----------------------
## 5. Import the initial data of `leanote`

Open up a new command line session, then copy n paste the command below into your command line. Note the difference between version 2 and 3 of `mongodb`:

- For mongodb v2:
```
mongorestore  -h localhost -d leanote  --directoryperdb C:\Go\src\github.com\leanote\leanote\mongodb_backup\leanote_install_data
```

- For mongodb v3:
```
mongorestore -h localhost -d leanote --dir C:\Go\src\github.com\leanote\leanote\mongodb_backup\leanote_install_data
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
## 6. Configure `leanote`

The configurations of `leanote` is stored in the file `leanote/conf/app.conf`. 

One setting that you are **strongly suggested** to modify is `app.secret`, please change arbitrary number of digits of the string to something different, but keeping the string length unchanged. This is to avoid potential security issues. (See screenshot below)

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image019.png)

Other settings can remain as they are for now. For some database related settings, see the [FAQ page](https://github.com/leanote/leanote/wiki/Leanote-QA-English).

----------------
## 7. Run `leanote`

Open up a new command line session and type in to generate the `revel` command:
```
go install github.com\\revel\\cmd\\revel
```  

Then launch `leanote` by:
```
revel run github.com\\leanote\\leanote  
```

A successful startup will show the following outputs:

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image021.png)

**★NOTE: You can minimize the command line window for now, but DO NOT CLOSE IT! (same as the mongodb window)**

With both command line windows (`mongodb` and `revel`) still open, fire up your browser and enter in the address bar `http://localhost`. You can use either of the default accounts

```
user1 username: admin, password: abc123 (administrator account)
user2 username: demo@leanote.com, password: demo@leanote.com (for demonstration)
```

or create you new own account.

![enter image description here](http://7xi5m5.com1.z0.glb.clouddn.com/leanote/image/image023.png)

------------

# Attention!!!!!

Please note that you run `Mongodb` with no `auth` option which mentioned in this paper, if your server is exposed to the internet, anyone can access and modify and delete it!!!!!! So it's very dangerous to run `Mongodb` in this way. You must add user and password to `Mongodb` and run it with `auth` option. Please see [How to add new users to mongodb database?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#6-how-to-add-new-users-to-mongodb-database)

# Trouble shooting

If you encounter issues or want to know more about `leanote`'s configurations, refer to the [FAQ page](https://github.com/leanote/leanote/wiki/Leanote-QA-English).
