**This tutorial explains doing a *source code* installation on Mac and Linux systems.**

`Leanote` source distribution is suitable for developers or those who need to use the new features sooner.

- For **binary** installation on Mac and Linux, see [here](https://github.com/leanote/leanote/wiki/leanote-binary-installation-on-Mac-and-Linux-(En)). 
- For **source** installation on Windows, see [here](https://github.com/leanote/leanote/wiki/leanote-source-installation-on-Windows-(En)).
- For **binary** installation on Windows, see [here](https://github.com/leanote/leanote/wiki/leanote-binary-installation-on-Windows-(En)).

-------------
**Installation overview:**

1. Install the execution environment -- `Golang`
2. Fetch the `Leanote` source code
3. Install the database -- `Mongodb`
4. Import initial data of `Leanote`
5. Use `revel` to run `Leanote`


-------------
## 1. Install `Golang`

Go to http://golang.org and download the latest Golang(1.7+) corresponding to your OS.

Suppose you downloaded the .tar.gz file to your HOME directory (e.g. /home/user1), extract the file there:
```
$> cd /home/user1
$> tar -xzvf go1.6.linux-amd64.tar.gz
```

Make a new directory `gopackage` under `/home/user1`, to store the `go` packages and the compiled files:
```
$> mkdir /home/user1/gopackage
```

Edit `/etc/profile` to configure some environment variables:
```
$> sudo vim /etc/profile
```

Here I'm using the `vim` editor. Feel free to use whatever editor you prefer (e.g. `nano`).
Add the following lines to your `/etc/profile` file, and remember to replace "user1" with your own username:
```
export GOROOT=/home/user1/go
export GOPATH=/home/user1/gopackage
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

To make the changes take effect:
```
$> source /etc/profile
```

Now check whether `go` is installed successfully:
```
$> go version
```

If the terminal prints a message similar to the following, the installation is successful:
```
go version go1.6 linux/amd64
```

-------------
## 2. Fetch `Leanote`


### 2.1 Method 1 (**Recommend**): 

Download [leante-all-master.zip](https://github.com/leanote/leanote-all/archive/master.zip). Extract it to any folder and move the `src` directory to `/home/user1/gopackage/`:

```
$> wget https://github.com/leanote/leanote-all/archive/master.zip
$> unzip master.zip
$> cp -r ./master/src /home/user1/gopackage
```

Then use the following command to generate `revel` which will be used to run `Leanote`:
```
$> go install github.com/revel/cmd/revel
```

### 2.2 Method 2

Alternatively, you could also use `go get`to download the `Leanote` package. As `go get` will call the `git` and `mercurial` commands, you need to install them first.

To Install `git`:

```
$> sudo apt-get install git-core mercurial openssh-server openssh-client
```

Then Fetch `Revel`, `Leanote` and related dependencies:
```
$> go get github.com/revel/cmd/revel
$> go get github.com/leanote/leanote/app
```

It may take a while to download these files, please be patient.
The source code of `Leanote` is stored in `/home/user1/gopackage/src/github.com/leanote/leanote`.

-------------
## 3. Install `Mongodb`

### 3.1 Download `Mongodb` and configure

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
adding the following line to your `~/bash_profile` or `/etc/profile` (make sure you type in the correct username and version strings): 
```
export PATH=$PATH:/home/user1/mongodb-linux-x86_64-3.0.1/bin
```

Again to make your modification take effect:
```
$> source /etc/profile
```

### 3.2 Test `Mongodb` installation

To verify the installation of `Mongodb`, make a new folder (e.g. `data`) under `/home/user1` to store data:
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
## 4. Import initial `c` data

`Leanote`'s initial data is stored in `/home/user1/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data`

Open a terminal and paste in the following command to import initial data.

```
$> mongorestore -h localhost -d leanote --dir /home/$USER/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data/
```

Now `Mongodb` has created a `leanote` database, you can have a peek into it, for instance query how many tables `leanote` database has:
```
$> mongo
> show dbs 
leanote	0.203125GB
local	0.078125GB
```

Tell `mongodb` to use our newly created `leanote` database:
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
## 5. Configure `Leanote`

The configuration of `Leanote` is controlled by this file: `/home/user1/gopackage/src/github.com/leanote/leanote/conf/app.conf`.

One setting that you are **strongly suggested** to modify is `app.secret`, please change arbitrary number of digits of the string to something different, but keeping the string length unchanged. This is to avoid potential security issues.

Other optional changes you can make includes `db.username`, `db.password` (more on these in the Trouble Shooting section) and etc..

-------------
## 6.  Run `Leanote`

If you have successfully come to this stage, there is just one more step to go. 

**Make sure the `Mongodb` is still up and running**, and your `9000` port (the default port, which can be changed later) is open. Then run:
```
$> revel run github.com/leanote/leanote
```

Note that if you are using revel 0.12 or above, please refer to [this post](https://github.com/leanote/leanote/pull/98)

Congratulations, now fire up you browser and enter `http://localhost:9000` (or `http://IP_ADDRESS_OF_SERVER:9000`) into the address bar. Voilà! Welcome to `Leanote` and happy note-taking!


------------

# Attention!!!!!

Please note that you run `Mongodb` with no `auth` option which mentioned in this paper, if your server is exposed to the internet, anyone can access and modify and delete it!!!!!! So it's very dangerous to run `Mongodb` in this way. You must add user and password to `Mongodb` and run it with `auth` option. Please see [How to add new users to mongodb database?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#6-how-to-add-new-users-to-mongodb-database)

# Trouble shooting

If you encounter issues or want to know more about `Leanote`'s configurations, refer to the [FAQ page](https://github.com/leanote/leanote/wiki/Leanote-QA-English).


