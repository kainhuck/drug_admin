module drug_api

go 1.14

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20200505102242-63963976dee0
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/jinzhu/gorm v1.9.14
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.57.0 // indirect
)

replace (
	drug_api/conf => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/conf
	drug_api/middleware => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/middleware
	drug_api/models => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/models
	drug_api/pkg/e => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/pkg/e
	drug_api/pkg/setting => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/pkg/setting
	drug_api/pkg/util => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/pkg/util
	drug_api/routers => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/routers
	drug_api/routers/api => /home/kain/Documents/学校/2020_school/数据库实训/课程设计/drug_api/routers/api
)
