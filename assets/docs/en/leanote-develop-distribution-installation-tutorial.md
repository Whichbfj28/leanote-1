Leanote develop distribution is suitable for developer or those who need to use the new feature quickly.

Installation steps:

1. Install golang
2. Fetch revel and leanote source code
3. Install mongodb
4. Import initial data
5. Use revel to run leanote


## Install golang

Go to http://golang.org to download the newest golang(1.3.1+) corresponding to your os.

If you download the file to /home/user1, extract the file
```
$> cd /home/user1
$> tar -xzvf go1.3.1.linux-amd64.tar.gz
```
Make a new directory gopackage under /home/user1 directory (gopackage will store go package and the compiled file)

```
$> mkdir /home/user1/gopackage
```

Configure environment variable and edit /etc/profile:
```
export GOROOT=/home/user1/go
export GOPATH=/home/user1/gopackage
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
To make the environment variable take effect
```
$> source /etc/profile
```

Check if go has installed successfully.
```
$> go version
# if the terminal print the following message, it means success.
go version go1.3.1 linux/amd64
```

## Fetch revel and leanote


#### **Recommend** Method 1: Download all the leanote's source and dependencies from [leanote-all](https://github.com/leanote/leanote-all)

Download [leante-all-master.zip](https://github.com/leanote/leanote-all/archive/master.zip). Extract it and move the src directory to `/home/user1/gopackage/`

Use the following cmd to generate `revel` which will be used to run leanote.
```
$> go install github.com/revel/cmd/revel
```
### Method 2

This method use 「go get」 to download package. As 「go get」 will call git and mercurial, so you need to install them first.

#### Install git

```
$> sudo apt-get install git-core mercurial openssh-server openssh-client
```

#### Fetch revel and leanote
Open the terminal, type the following command to download revel , leanote and related dependent packages.
It may take long time to download these files, please wait patiently.
```
$> go get github.com/revel/cmd/revel
$> go get github.com/leanote/leanote/app
```
Source code of leanote is under /home/user1/gopackage/src/github.com/leanote/leanote directory.

## Install mongodb

Download address:  http://www.mongodb.org/downloads 

Fast download:
* 64-bit linux mongodb 2.6.4: http://www.mongodb.org/dr//fastdl.mongodb.org/linux/mongodb-linux-x86_64-2.6.4.tgz/download
* 64-bit linux mongodb 3.0.1: https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-3.0.1.tgz

Download the file to /home/user1, extract the file.
```
$> cd /home/user1
$> tar -xzvf mongodb-linux-x86_64-2.6.4.tgz/
```

To guarantee you can reference mongodb command from anywhere, you can configure environment variable.
Edit ~/bash_profile or /etc/profile and add mongodb path to PATH. 

```
$> sudo vim /etc/profile
```
add:
```
export PATH=$PATH:/home/user1/mongodb-linux-x86_64-2.6.4/bin
```
To make your modification take effect
```
$> source /etc/profile
```

### Test mongodb

Make a new folder data under /home/user1 to store data.
```
mkdir /home/user1/data
```

```
# start mongod (mongodb server part)
mongod --dbpath /home/user1/data
```

Now mongod has started, you can open terminal and test using mongodb.
```
$> mongo
> show dbs
```

Now your mongodb installation is complete, let's import initial data for mongodb.

## Import initial leanote data

Leanote's initial data is in `/home/user1/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data`

Open terminal and type the following the following command to import initial data.

The import data cmd in mognodb v2 and mongodb v3 is different.

For mongodb v2:
```
$> mongorestore -h localhost -d leanote --directoryperdb /home/user1/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data/
```
For mongodb v3:
```
mongorestore -h localhost -d leanote --dir /home/user1/leanote/mongodb_backup/leanote_install_data/
```

Now mongodb has created leanote database, you can query how many tables leanote database has.
```
$> mongo
> show dbs 
leanote	0.203125GB
local	0.078125GB
> use leanote 
switched to db leanote
> show collections # a collection in mongodb is a table in mysql
files
has_share_notes
note_content_histories
note_contents
notebooks
...
```

The initial users table has two accounts:
```
user1 username: admin, password: abc123 (administrator who can manage background system)
user2 username: demo@leanote.com, password: demo@leanote.com (just for experiencing)
```

## Configure leanote

Edit `/home/user1/gopackage/src/github.com/leanote/leanote/conf/app.conf`, You need to modify `app.secret`, please change it to a different value, if not, there will be secure problem.
Change `db.username` `db.password` and other db options if needed.

## Run leanote
```
$> revel run github.com/leanote/leanote
```
If you are using revel 0.12 or above, please refer to this [post](https://github.com/leanote/leanote/pull/98)

Congratulations, open you browser and enter `http://localhost:9000` as the address, you just need to experience leanote.

## Trouble shooting

### Issue: Cannot running
```
Go to /@tests to run the tests.
panic: auth fails

goroutine 1 [running]:
github.com/leanote/leanote/app/db.Init()
	/home/life/gopackage1/src/github.com/leanote/leanote/app/db/Mgo.go:64 +0x356
gi
```
answer: the database configuration is not correct, please check if the username and password in conf/app.conf is correct.

### Issue: How to modify the default port to 80 ?

Modify the file: conf/app.conf. Update port to 80:
```
http.port=80

site.url=http://localhost
```

### Issue: How to add database administrator ?

OK, the data has been imported. Now you need to add a new user to leanote database, like root account in mysql. Mongodb initially doesn't have any account, which is not secure, so you need to add a new user to connect to leanoate database(note: the account here is not the user in user table but the account which is used to connect to leanote database).

The add user cmd in mongodb v2 and mongodb v3 is different:

For mongodb v2:

```
> use leanote;
# add a new user, root, pasword is abc123
> db.addUser("root", "abc123");
{
    "_id" : ObjectId("53688d1950cc1813efb9564c"),
    "user" : "root",
    "readOnly" : false,
    "pwd" : "e014bfea4a9c3c27ab34e50bd1ef0955"
}
# test if correct
> db.auth("root", "abc123");
1 #return 1 means success 
```

For mongodb v3:
```
> use leanote;
# # add a new user, root, pasword is abc123
> db.createUser({
    user: 'root',
    pwd: 'abc123',
    roles: [{role: 'dbOwner', db: 'leanote'}]
});
# test if correct
> db.auth("root", "abc123");
1 #return 1 means success 
```

You must modify db.username and db.password on the configuration of mongodb. 

Modify file: `conf/app.conf`
```
# mongdb
db.host=localhost
db.port=27017
db.dbname=leanote # required
db.username=root # if not exists, please leave blank
db.password=abc123 # if not exists, please leave blank
```

After you have added the root user, you can re-run mongod, and open access authentication. You can enter ctrl+c to exit mongodb.

Start mongodb with authorization:

```
$> mongod --dbpath /home/user1/data --auth
```

Restart leanote via `revel run github.com/leanote/leanote`

### Issue: How to update leanote?

You can use git pull to fetch the newest version of leanote. If you have modified leanote, you can fetch(
fetch is recommended) the newest leanote to local and merge with your local version.
e.g.
```
git fetch origin master:tmp # fetch the newest version leanote ,alias tmp 
git diff tmp # compare and diff
git merge tmp # merge the newest leanote with your local version
```