module shortlink

go 1.13

require (
	github.com/chasex/redis-go-cluster v1.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/gorm v1.9.14
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1 // indirect
	google.golang.org/grpc v1.30.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/chasex/redis-go-cluster => github.com/chasex/redis-go-cluster v1.0.1-0.20161207023922-222d81891f1d
