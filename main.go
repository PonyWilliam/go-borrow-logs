package main

import (
	"github.com/PonyWilliam/go-borrow-logs/domain/repository"
	services2 "github.com/PonyWilliam/go-borrow-logs/domain/server"
	"github.com/PonyWilliam/go-borrow-logs/handler"
	works "github.com/PonyWilliam/go-borrow-logs/proto"
	"github.com/opentracing/opentracing-go"
	"strconv"
	"time"

	common "github.com/PonyWilliam/go-common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
)
var QPS = 1000
func main() {
	consulConfig,err := common.GetConsualConfig("1.116.62.214",8500,"/micro/config")
	//配置中心
	if err != nil{
		log.Fatal(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(
		func(options *registry.Options){
			options.Addrs = []string{"1.116.62.214"}
			options.Timeout = time.Second * 10
		},
	)
	t,io,err := common.NewTracer("go.micro.service.borrowlog",":6832")
	if err != nil{
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	srv := micro.NewService(
		micro.Name("go.micro.service.borrowlog"),
		micro.Version("latest"),
		micro.Address(":15340"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil{
		log.Error(err)

	}
	defer db.Close()
	db.SingularTable(true)
	srv.Init()
	go common.PrometheusBoot("5001")//　协程
	rp := repository.NewBorrowLogsRepository(db)
	err =rp.InitTable()
	if err!=nil{
		err := rp.InitTable()
		if err!=nil{
			log.Error(err)
		}
	}

	LogServices := services2.NewLogsService(repository.NewBorrowLogsRepository(db))
	err = works.RegisterBorrowLogsHandler(srv.Server(),&handler.Logs{LogsServices:LogServices})
	if err:=srv.Run();err!=nil{
		log.Fatal(err)
	}
}