## Install mongodb

see http://www.mongodb.org to install mongodb first.

## Import initial data to mongodb

```
$> mongod

Open another terminal and 

$> mongorestore -h localhost -d leanote --directoryperdb PATH_TO_LEANOTE/mongodb_backup/leanote_install_data
```

## Set user and password

```
$> mongo
> show dbs;
> use leanote;
> db.addUser("root", "root123");
> db.auth("root", "root123");
```

## Config app/app.conf
Put the mongodb configuration into app/app.conf