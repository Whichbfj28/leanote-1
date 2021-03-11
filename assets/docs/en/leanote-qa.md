Here is a list of FAQ regarding the installation/configuration of `leanote`. For user manual please go to [Leanote usage manual](http://leanote.leanote.com/post/Leanote-manual-project).

If you encounter any issues during the installation or usage of `leanote`, please don't hesitate to contact me at leanote@leanote.com.

We also greatly appreciate your support by submitting issues or helping refine the documentation of `leanote`.

# Table of Content:

* [Error on runtime: "no reachable server"](https://github.com/leanote/leanote/wiki/Leanote-QA-English#1-error-on-runtime-no-reachable-server)
* [Error on runtime: "panic: auth fails"](https://github.com/leanote/leanote/wiki/Leanote-QA-English#2-error-on-runtime-panic-auth-fails)
* [Leanote is running but I cannot log on](https://github.com/leanote/leanote/wiki/Leanote-QA-English#3-leanote-is-running-but-i-cannot-log-on)
* [How to change Leanote's port](https://github.com/leanote/leanote/wiki/Leanote-QA-English#4-how-to-change-leanotes-port)
* [How to bind a domain?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#5-how-to-bind-a-domain)
* [How to add new users to mongodb database?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#6-how-to-add-new-users-to-mongodb-database)
* [How to add admin account to Leanote?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#7-how-to-add-admin-account-to-leanote)
* [How to setup https access for leanote?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#8-how-to-setup-https-access-for-leanote)
* [Import of github.com/revel/revel/modules/testrunner failed](https://github.com/leanote/leanote/wiki/Leanote-QA-English#9-import-of-githubcomrevelrevelmodulestestrunner-failed)
* [How to update the dev version of leanote?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#10-how-to-update-the-dev-version-of-leanote)
* [How to update the binary version of leanote?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#11-how-to-update-the-binary-version-of-leanote)
* [Why "site.url"?](https://github.com/leanote/leanote/wiki/Leanote-QA-English#12-why-siteurl)
* [Cannot sync images](https://github.com/leanote/leanote/wiki/Leanote-QA-English#cannot-sync-images)
* [Export as PDF configuration: wkhtmltopdf](https://github.com/leanote/leanote/wiki/Leanote-QA-English#13-export-as-pdf-configuration-wkhtmltopdf)
* [Cannot visit via ip](https://github.com/leanote/leanote/wiki/Leanote-QA-English#14-cannot-visit-via-ip)

## 1. Error on runtime: "no reachable server"

Double check that the `mongodb` database is running, if so, try changing `db.host=localhost` to `db.host=127.0.0.1` in the `conf/app.conf` file. Then re-run `leanote`.



## 2. Error on runtime: "panic: auth fails"

If you encounter the following error:
```
Go to /@tests to run the tests.
panic: auth fails

goroutine 1 [running]:
github.com/leanote/leanote/app/db.Init()
	/home/life/gopackage1/src/github.com/leanote/leanote/app/db/Mgo.go:64 +0x356
```

It is very likely that something went wrong with the database. Please check:

1. Is the database running?
2. If the database was launched in auth mode (by append the `--auth` option), check the `conf/app.conf` has the correct `db.username` and `db.password` settings.

Below is the default `conf/app.conf` settings. Note that by default `db.dbname=leanote`, `db.username` and `db.password` are both blank.
```
# mongdb
db.host=localhost
db.port=27017
db.dbname=leanote # required
db.username= # if not exists, please leave it blank
db.password= # if not exists, please leave it blank
# or you can set the mongdb url for more complex needs the format is:
# mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
# db.url=mongodb://root:root123@localhost:27017/leanote
db.urlEnv=${MONGODB_URL} # set url from env
```


## 3. Leanote is running but I cannot log on

The database is not running. Please restart database and then `leanote`.

## 4. How to change Leanote's port

Suppose you want to change the port number to 8080. Edit `conf/app.conf`:

```
http.port=8080
site.url=http://localhost:8080
```

Then restart `leanote`, use the new port for access: `http://localhost:8080`.

## 5. How to bind a domain?

Suppose you want to use the domain `a.com` to access your leanote service, you need to change leanote's port to 80: edit the `conf/app.conf` as shown below:

```
http.port=80
site.url=http://a.com
```
Then restart `leanote`. Of course you will need to bind the domain to the IP address of the leanote server.
What if there is already another service using the 80 port on the server? Then please google "use nginx to forward to different ports".

## 6. How to add new users to mongodb database?

`Mongodb` by default doesn't create any account, which is considered insecure. Therefore, after you have imported the initial data, you need to add a new user to the `leanote` database, just like the `root` account in `mysql`. Note that the "account" here is distinct from the "user" in the "user" table: it is the account used to connect to `leanote` database.

The "add user" command in `mongodb v2` and `mongodb v3` are different:

- For mongodb v2:

```
$> mongo
> use leanote;
# add a new user, root, password is abc123
> db.addUser("root", "abc123");
{
    "_id" : ObjectId("53688d1950cc1813efb9564c (Wtf is this developer!?)"),
    "user" : "root",
    "readOnly" : false,
    "pwd" : "e014bfea4a9c3c27ab34e50bd1ef0955 (Wtf is this developer!?)"
}
# test if correct
> db.auth("root", "abc123");
1 #return 1 means success
```

- For mongodb v3:
```
$> mongo
> use leanote;
# # add a new user, root, password is abc123
> db.createUser({
    user: 'root',
    pwd: 'abc123',
    roles: [{role: 'dbOwner', db: 'leanote'}]
});
# test if correct
> db.auth("root", "abc123");
1 #return 1 means success
```

You also need to modify `db.username` and `db.password` in the configuration file of `mongodb`.

Edit the file `/home/user1/gopackage/src/github.com/leanote/leanote/conf/app.conf`:
```
# mongdb
db.host=localhost
db.port=27017
db.dbname=leanote # required
db.username=root # if not exists, please leave blank
db.password=abc123 # if not exists, please leave blank
```

After you have added a root user, you can re-run `mongod` with authentication.
(Use Ctrl+c to exit `mongodb`.)

Start `mongodb` with authorization:

```
$> mongod --dbpath /home/user1/data --auth
```

Restart `leanote`:
```
revel run github.com/leanote/leanote
```

## 7. How to add admin account to Leanote?

The default "super user" account of `leanote` is `admin`, which cannot change back once you altered it. In such cases you could edit the `conf/app.conf` file to change it:
E.g. To assign user `life` as super user, change or add a line:


```
adminUsername=life
```


## 8. How to setup https access for leanote?

### 8.1. Generate SSL certificate

You can buy one certificate from online, or create one by yourself.
Below is a shell script to auto-create a certificate:

```
#!/bin/sh

# create self-signed server certificate:

read -p "Enter your domain [www.example.com]: " DOMAIN

echo "Create server key..."

openssl genrsa -des3 -out $DOMAIN.key 1024

echo "Create server certificate signing request..."

SUBJECT="/C=US/ST=Mars/L=iTranswarp/O=iTranswarp/OU=iTranswarp/CN=$DOMAIN"

openssl req -new -subj $SUBJECT -key $DOMAIN.key -out $DOMAIN.csr

echo "Remove password..."

mv $DOMAIN.key $DOMAIN.origin.key
openssl rsa -in $DOMAIN.origin.key -out $DOMAIN.key

echo "Sign SSL certificate..."

openssl x509 -req -days 3650 -in $DOMAIN.csr -signkey $DOMAIN.key -out $DOMAIN.crt
```

Suppose the script has created 2 files: `a.com.crt` and `a.com.key`.

### 8.2. Configure Nginx

Suppose `leanote` is running on port 9000, at domain `a.com`, then `nginx.conf` can be configured as follow. Note that the listing is incomplete and showing only the relevant lines.

```
# http (incomplete)
http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;
    
    upstream  a.com  {
        server   localhost:9000;
    }

    # http
    server
    {
        listen  80;
        server_name  a.com;
        
        # Force redirect to https
        # If you dont want https force redirect, comment out the rewrite line 
        rewrite ^/(.*) https://$server_name/$1 permanent;
        
        location / {
            proxy_pass        http://a.com;
            proxy_set_header   Host             $host;
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
    
    # https
    server
    {
        listen  443 ssl;
        server_name  a.com;
        ssl_certificate     /root/a.com.crt; # Change this path to point a.com.crt, same below.
        ssl_certificate_key /root/a.com.key;
        location / {
            proxy_pass        http://a.com;
            proxy_set_header   Host             $host;
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
    }
}
```

Save your changes, then reload Nginx:
```
$> sudo nginx -s reload
```
Then restart leanote.
Note that you might want to clear the caches of your browser first to see the changes take effect.


## 9. Import of `github.com/revel/revel/modules/testrunner` failed

Versions before 1.0 (beta version) may encounter this error. Version 1.0 uses `revel-0.12` so should be immune to such problems.

```
Failed to load module.  Import of github.com/revel/revel/modules/testrunner failed: cannot find package "github.com/revel/revel/modules/testrunner" in any of:
	/Users/life/app/go1.4/src/github.com/revel/revel/modules/testrunner (from $GOROOT)
	/Users/life/Documents/Go/package_base/src/github.com/revel/revel/modules/testrunner (from $GOPATH)
```

`revel 0.12` uses a different configuration, please change `conf/app.conf` as below: 
```
module.static=github.com/revel/modules/static		 	
module.testrunner=github.com/revel/modules/testrunner	 
```

## 10. How to update the dev version of leanote?

You could use `git pull` to get the latest version of `leanote`. If you have made changes to `leanote`, you may `fetch` (recommended) the latest version to local machine, then merge with the local version. E.g.
 
```
git fetch origin master:tmp # Get the latest version from remote, name it tmp
git diff tmp # Diff tmp with local version
git merge tmp # Merge to local
```
If you have difficulties accessing via git, download from https://github.com/leanote/leanote.

1. First make of copy of the previous `leanote` folder in case something screwed up.
2. Replace your `leanote` with the newly downloaded one.
3. Copy these folders/file from your old version to the new install location:
```
    * /public/upload
    * /files
    * /conf/app.conf
```

Then restart `leanote`.

In case of issues like "cannot find package "github.com/PuerkitoBio/goquery" in any of:...", that is because `leanote` has added some new dependencies, which can be obtained by the `go get` command. E.g. to get the new "github.com/PuerkitoBio/goquery":

```
go get github.com/PuerkitoBio/goquery
```

Alternatively you could download bundled dependencies with the source code from: https://github.com/leanote/leanote-all. 

## 11. How to update the binary version of leanote?

Download the new binary version of `leanote`, then copy these folders/files from the old version to the new install location:

```
* /public/upload
* /files
* /conf/app.conf
```


## 12. Why "site.url"?

`site.url` is the domain for public access. For instance you could set it to `http://a.com`, and set the port 9000 for `leanote`, then use `Nginx` to forward the traffic to 9000. 

`site.url` is also used in generating paths for images/attachments in the notes.

If deploy your `leanote` by forwarding https traffic using `nginx`, then `site.url` needs to be set as `https://a.com`. Otherwise in the exported blog pages, `css` and `js` are shown as `http` links in html, which will be blocked by new versions of browsers like `firefox`, resulting erroneous page display. For more details see https://github.com/leanote/leanote/issues/228.

## Cannot sync images

Please make sure that the `site.url` on `conf/app.conf` is same as the `Host` which set on desktop app.

## 13. Export as PDF configuration: wkhtmltopdf

`leanote` utilizes `wkhtmltopdf` to export PDFs, so it needs to be installed first.
Then log in as `admin` user, setup up the path of `wkhtmltopdf` in the control panel.

Here is now to install `wkhtmltopdf`:

### 13.1 Compile the source 

Compiling the source will set you free from dependency issues, as the compilation/installation script of `wkhtmltopdf` has packed up necessary info for needed dependencies. More details see the `install.md` file in [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf).

### 13.2 Binary install (recommended)

Download [wkhtmltopdf](http://wkhtmltopdf.org/downloads.html), choose the correct version, then install it using:
```
rpm -ivh
```

You might see the lack of some dependencies. Install them:

```
yum install -y fontconfig libX11 libXext libXrender xorg-x11-fonts-Type1 xorg-x11-fonts-75dpi libpng
```
Then re-run `rpm -ivh`.



### 13.3 Test `wkhtmltopdf`

Issue the command
```
wkhtmltopdf http://google.com google.pdf
```
to check the installation of `wkhtmltopdf`.

If still lacks dependencies, run `ldd wkhtmltopdf` to find out what is missing, then install them correspondingly.

In some systems it may give you a lack of `libpng` error. This could be the `libpng` library version is too high on your machine, while some versions of pre-compiled `wkhtmltopdf` depend on lower versions of `libpng`. In such cases, find out the right version using `yum provides \*libpng12.so.0\*`, then install it, e.g. `yum install libpng12-1.2.50-6.el7.x86_64.

### 13.4 Chinese problem

After successfully exporting the google page to PDF, try exporting a page with Chinese language, e.g. `wkhtmltopdf www.ubuntu.org.cn`. If it shows missing characters, it could be the linux system lacks proper Chinese font. You could copy some Chinese font from a Windows system (`windows/fonts`) to `/usr/share/fonts` to fix this.

## 14. Cannot visit via ip
please find and update `app.conf`
```
http.addr=0.0.0.0 # listen on all ip addresses
```
Then restart Leanote