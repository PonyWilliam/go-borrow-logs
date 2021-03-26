package repository

import (
	"github.com/PonyWilliam/go-borrow-logs/domain/model"
	"github.com/jinzhu/gorm"
)
type ILogs interface {
	InitTable() error
	ToOther(int64,int64,int64,string,int64) error
	Confirm(int64) error
	FindAll() ([]model.BorrowLogs,error)
	FindByID(id int64) (model.BorrowLogs,error)
	FindByWID(wid int64)([]model.BorrowLogs,error)
	FindByPID(pid int64)([]model.BorrowLogs,error)
	FindByLogID(logid int64)([]model.BorrowLogs,error)
}
func NewBorrowLogsRepository(db *gorm.DB)ILogs{
	return &BorrowLogsRepository{mysqlDB: db}
}
type BorrowLogsRepository struct{
	mysqlDB *gorm.DB
}
func(b *BorrowLogsRepository) InitTable() error{
	if b.mysqlDB.HasTable(&model.BorrowLogs{}){
		return nil
	}
	return b.mysqlDB.CreateTable(&model.BorrowLogs{}).Error
}
func(b *BorrowLogsRepository) ToOther(ReqWID int64,RspWID int64,PID int64,Reason string,Logid int64) error {
	return b.mysqlDB.Model(&model.BorrowLogs{}).Create(&model.BorrowLogs{
		ReqWID: ReqWID,
		RspWID: RspWID,
		PID: PID,
		Reason: Reason,
		Confirm: false,
		Logid:Logid,
	}).Error
}
func(b *BorrowLogsRepository) Confirm(id int64)error{
	logs := &model.BorrowLogs{}
	b.mysqlDB.Where("id = ?",id).First(&logs)
	return b.mysqlDB.Model(logs).Update("Confirm",true).Error
}
func(b *BorrowLogsRepository) FindAll() (logs []model.BorrowLogs,err error){
	return logs,b.mysqlDB.Model(&logs).Find(&logs).Error
}
func(b *BorrowLogsRepository) FindByID(id int64) (log model.BorrowLogs,err error){
	return log,b.mysqlDB.Where("id = ?",id).Find(&log).Error
}
func(b *BorrowLogsRepository) FindByWID(wid int64) (logs []model.BorrowLogs,err error){
	return  logs,b.mysqlDB.Where("wid = ?",wid).Find(&logs).Error
}
func(b *BorrowLogsRepository) FindByPID(pid int64) (logs []model.BorrowLogs,err error){
	return  logs,b.mysqlDB.Where("pid = ?",pid).Find(&logs).Error
}
func(b *BorrowLogsRepository) FindByLogID(logid int64)(logs []model.BorrowLogs,err error){
	return logs,b.mysqlDB.Where("logid = ?",logid).Find(&logs).Error
}