module github.com/Coder-stars/gin-example

go 1.16

replace (
	github.com/Coder-stars/gin-example/conf => ./pkg/conf
	github.com/Coder-stars/gin-example/middleware => ./middleware
	github.com/Coder-stars/gin-example/models => ./models
	github.com/Coder-stars/gin-example/pkg/e => ./pkg/e
	github.com/Coder-stars/gin-example/pkg/setting => ./pkg/setting
	github.com/Coder-stars/gin-example/pkg/util => ./pkg/util
	github.com/Coder-stars/gin-example/routers => ./routers
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/swaggo/gin-swagger v1.3.1
	github.com/swaggo/swag v1.5.1
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
