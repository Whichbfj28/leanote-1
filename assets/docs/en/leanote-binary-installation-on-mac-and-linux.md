**This tutorial explains doing a *binary* installation on Mac and Linux systems.**

If you wish to experience the new feature sooner or help develop `Leanote`, please try using the source distribution. 

----------------
# Installation Overview:

1. Download the binary file of `Leanote`.
2. Install the database -- `Mongodb`.
3. Import initial data to `Mongodb`.
4. Configure `Leanote`.
5. Run `Leanote`.

----------------------
## 1. Download the binary file of `Leanote`

Choose and download the binary file corresponding to your system from [here](https://github.com/coocn-cn/leanote/releases/latest).

Suppose it is saved in the `/home/user1` folder, extract the `.zip` file there using:
```
$> cd /home/user1
$> unzip master.zip
```

This will create a `Leanote` directory under `/home/user1`. 

---------------------
## 2. Install the database -- `Mongodb`

### 2.1 Download `Mongodb` and configure

You could download a more up-to-date version from the official site of [Mongodb](http://www.mongodb.org/downloads).
Or, you could use the following links to get the versions that are validated to be working by the developers.

Fast download:
* 64-bit linux Mongodb 3.0.1: https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-3.0.1.tgz

Save the file to `/home/user1`, then extract it:
```
$> cd /home/user1
$> tar -xzvf mongodb-linux-x86_64-3.0.1.tgz/
```

To make sure that you can reference the `Mongodb` command from anywhere, configure its environment variable by
editing your `~/bash_profile` or `/etc/profile`:
```
sudo vim /etc/profile
```

Here I'm using the `vim` editor, feel free to use whatever text editor you prefer (e.g. `nano`). Add the following line to the file, and remember to replace with you own username and version strings:
```
export PATH=$PATH:/home/user1/mongodb-linux-x86_64-3.0.1/bin
```

To make your modification take effect:
```
$> source /etc/profile
```

### 2.2 Test `Mongodb` installation

To verify the installation of `mongodb`, make a new folder (e.g. `data`) under `/home/user1` to store data:
```
$> mkdir /home/user1/data
```

Then start the `Mongodb` database server. You might want it to run in the background, so append `&` to the end: 
```
$> mongod --dbpath /home/user1/data &
```

Now `Mongodb` is up and running, you can open a new terminal (or in the same terminal session if you have `mongod` run in the background) and launch it:
```
$> mongo
> show dbs
```

Should no error pops up, your `Mongodb` installation is complete, let's import initial data to `Mongodb`.

-------------
## 3. Import initial `Leanote` data

`Leanote`'s initial data is stored in `PATH_TO_LEANOTE/mongodb_backup/leanote_install_data`

Open a terminal and paste in the following command to import initial data. Note the difference between the version 2 and 3 of `Mongodb`:

```
$> mongorestore -h localhost -d leanote --dir PATH_TO_LEANOTE/mongodb_backup/leanote_install_data/
```

Now `Mongodb` has created a `Leanote` database, you can have a peek into it, for instance query how many tables `leanote` database has:
```
$> mongo
> show dbs 
leanote	0.203125GB
local	0.078125GB
```

Tell `Mongodb` to use our newly created `leanote` database:
```
> use leanote 
switched to db leanote
```

Bit more playing around:
```
> show collections # a collection in Mongodb is a table in mysql
files
has_share_notes
note_content_histories
note_contents
notebooks
...
```

The initial `users` table has two accounts:
```
user1 username: admin, password: abc123 (administrator who can manage Leanote)
user2 username: demo@leanote.com, password: demo@leanote.com (just for demonstration)
```

-------------
## 4. Configure `Leanote`

The configuration of `Leanote` is controlled by this file: `PATH_TO_LEANOTE/conf/app.conf`.

One setting that you are **strongly suggested** to modify is `app.secret`, please change arbitrary number of digits of the string to something different, but keeping the string length unchanged. This is to avoid potential security issues.

Other optional changes you can make includes `db.username`, `db.password` (more on these in the Trouble Shooting section) and etc..

-------------
## 5.  Run `Leanote`

If you have successfully come to this stage, there is just one more step to go:

```
$> cd /home/user1/leanote/bin
$> bash run.sh
```

If you see a message similar to this, `Leanote` has started successfully:
```
...
TRACE 2013/06/06 15:01:27 watcher.go:72: Watching: /home/user1/leanote/bin/src/github.com/leanote/leanote/conf/routes
Go to /@tests to run the tests.
Listening on :9000...
```

Congratulations, now fire up you browser and enter `http://localhost:9000` into the address bar. Voilà! Welcome to `Leanote` and happy note-taking!


------------

# Attention!!!!!

Please note that you run `Mongodb` with no `auth` option which mentioned in this paper, if your server is exposed to the internet, anyone can access and modify and delete it!!!!!! So it's very dangerous to run `Mongodb` in this way. You must add user and password to `Mongodb` and run it with `auth` option. Please see [How to add new users to mongodb database?](https://github.com/coocn-cn/leanote/blob/master/assets/docs/en/leanote-qa.md#6-how-to-add-new-users-to-mongodb-database)

# Trouble shooting

If you encounter issues or want to know more about `Leanote`'s configurations, refer to the [FAQ page](https://github.com/coocn-cn/leanote/blob/master/assets/docs/en/leanote-qa.md).

