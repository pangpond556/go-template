# REST API Template
Simply REST API template with Golang and Mongodb

### Library
- Golang framework https://echo.labstack.com
- Mongo driver https://github.com/mongodb/mongo-go-driver

### Development

1. Create .env from .env.example

2. Uncomment "github.com/joho/godotenv" in models/init.go

3. Change "appname" to your application name in [ app.go, routers, controllers ]

##### Init application:
```sh
$ go mod init your_app_name
```

##### To start application:
```sh
$ go run app.go
```
