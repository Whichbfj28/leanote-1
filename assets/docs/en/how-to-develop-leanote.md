## Technologies and Architecture Overview

The main Leanote projects are:

* Leanote Server (https://github.com/leanote/leanote). It’s the Leanote server and web client. Which is written in Google's [Go](https://golang.org/) language using the [revel](https://github.com/revel/revel) framework. The data is stored in a MongoDB database. You can run Mongo in the cloud (recommended) or on your own computer (for development purposes). The server must be running all the time to allow clients (Desktop, iOS, Android) to connect and sync data to it. These apps communicate with the Leanote server via REST requests in JSON format via an [API](https://github.com/leanote/leanote/wiki/leanote-api-en). The technologies which the Leanote Server use are:
   * Golang
   * Mongodb
   * jQuery (Just jQuery, no other frameworks like Angularjs or React)
   * Bootstrap
   * Less

* Leanote Desktop App (https://github.com/leanote/desktop-app) which is the Leanote Desktop App. The data is synced from Leanote Server. The desktop apps allow offline note editing and sync seamlessly with the server once an internet connection is reestablished. Uses Leanote Server API. The technologies which the Leanote Server use are:
   * Electron (github.com/electron/electron)
   * jQuery (Just jQuery, no other frameworks like Angularjs or React)
   * Bootstrap
   * Less

* Leanote iOS (https://github.com/leanote/leanote-ios). Uses Leanote Server [API](https://github.com/leanote/leanote/wiki/leanote-api-en). The technologies which the Leanote Server use are:
   * Objective-c
   * Core Data

* Leanote Android (https://github.com/leanote/leanote-android). Uses Leanote Server [API](https://github.com/leanote/leanote/wiki/leanote-api-en). The technologies which the Leanote Server use are:
   * Java

## Install Leanote Server & Web

If you want to develop Leanote, you must install the Leanote server (https://github.com/leanote/leanote) first! Please see [Install Leanote](https://github.com/leanote/leanote/wiki/leanote-develop-distribution-installation-tutorial)

### Run Leanote Server & Web app
In order to get Leanote ready for development, [follow this guide](https://github.com/leanote/leanote/wiki/leanote-develop-distribution-installation-tutorial) to help you set up the server/web app.
Go to the applications `/app` folder and run `go get` to install dependencies. Then you should be able to run Leanote with dev mode or production mode.

Dev mode:
```
$> revel run github.com/leanote/leanote dev [port]
eg:
$> revel run github.com/leanote/leanote dev 9000
```

Production mode:
```
$> revel run github.com/leanote/leanote prod [port]
eg:
$> revel run github.com/leanote/leanote prod 8080
```

Please note that, for dev mode, Leanote use "note-dev.html" as the note view, and for product mode, Leanote use "note.html" as the note view.

### Build Leanote Server & web static file
Leanote use "Gulp" to build Leanote static file (js and note-dev.html -> note.html). You should install nodejs, gulp firstly.

On Leanote home, please run:
```
$> npm install
$> gulp
```

## How to develop Leanote Clients

* Leanote Desktop, Please see https://github.com/leanote/desktop-app/wiki
* Leanote iOS, Please see https://github.com/leanote/leanote-ios
* Leanote android, Please see https://github.com/leanote/leanote-android