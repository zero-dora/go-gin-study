module github.com/zero-dora/go-gin-example

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/astaxie/beego v1.12.3
	github.com/boombuler/barcode v1.0.1
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/robfig/cron v1.2.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.3.1
	github.com/swaggo/swag v1.7.1
	github.com/tealeg/xlsx v1.0.5
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210819135213-f52c844e1c1c // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/zero-dora/go-gin-example/conf => ./conf
	github.com/zero-dora/go-gin-example/docs => ./docs
	github.com/zero-dora/go-gin-example/middleware => ./middleware
	github.com/zero-dora/go-gin-example/models => ./models
	github.com/zero-dora/go-gin-example/pkg/app => ./pkg/app
	github.com/zero-dora/go-gin-example/pkg/e => ./pkg/e
	github.com/zero-dora/go-gin-example/pkg/file => ./pkg/file
	github.com/zero-dora/go-gin-example/pkg/gredis => ./pkg/gredis
	github.com/zero-dora/go-gin-example/pkg/logging => ./pkg/logging
	github.com/zero-dora/go-gin-example/pkg/qrcode => ./pkg/qrcode
	github.com/zero-dora/go-gin-example/pkg/setting => ./pkg/setting
	github.com/zero-dora/go-gin-example/pkg/upload => ./pkg/upload
	github.com/zero-dora/go-gin-example/pkg/util => ./pkg/util
	github.com/zero-dora/go-gin-example/routers => ./routers
	github.com/zero-dora/go-gin-example/service => ./service
)
