package main

import (
	"github.com/PonyWilliam/go-borrow-logs/domain/repository"
	"github.com/PonyWilliam/go-borrow-logs/domain/server"
	"github.com/PonyWilliam/go-borrow-logs/handler"
	Proto "github.com/PonyWilliam/go-borrow-logs/proto"
	common "github.com/PonyWilliam/go-common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul "github.com/micro/go-plugins/registry/consul/v2"
	"strconv"
	"time"
)

func main() {
	// New Service
	consulRegistry := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{"0.0.0.0"}
			options.Timeout = time.Second * 10

		})
	service := micro.NewService(
		micro.Name("go.micro.service.borrowLogs"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8089"),
		micro.Registry(consulRegistry),
	)
	consulConfig,err := common.GetConsualConfig("127.0.0.1",8500,"/micro/config")
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	if err!=nil{
		log.Fatal(err)
	}
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil{
		log.Error(err)

	}
	defer db.Close()
	db.SingularTable(true)
	rp := repository.NewBorrowLogsRepository(db)
	err = rp.InitTable()
	if err != nil{
		log.Fatal(err)
	}

	// Initialise service
	service.Init()
	err = Proto.RegisterBorrowLogsHandler(service.Server(),&handler.Logs{LogsServices: server.NewLogsService(rp)})
	if err!=nil{
		log.Fatal(err)
	}
	// Register Handler
	//_ = borrowLogs.RegisterBorrowLogsHandler(service.Server(), new(handler))

	// Register Struct as Subscriber
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
