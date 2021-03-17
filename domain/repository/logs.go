package repository

import (
	"github.com/PonyWilliam/go-borrow-logs/domain/model"
	"github.com/jinzhu/gorm"
)
type ILogs interface {
	InitTable() error
	ToOther(int64,int64,int64,string) error
	Confirm(int64) error
	//FIndAll() ([]*model.BorrowLogs,error)
	//FindByID(id int64) (*model.BorrowLogs,error)
	//FindByWID(wid int64)([]*model.BorrowLogs,error)
	//FindByPID(pid int64)([]*model.BorrowLogs,error)
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
func(b *BorrowLogsRepository) ToOther(ReqWID int64,RspWID int64,PID int64,Reason string) error {
	return b.mysqlDB.Model(&model.BorrowLogs{}).Create(&model.BorrowLogs{
		ReqWID: ReqWID,
		RspWID: RspWID,
		PID: PID,
		Reason: Reason,
		Confirm: false,
	}).Error
}
func(b *BorrowLogsRepository) Confirm(id int64)error{
	logs := &model.BorrowLogs{}
	b.mysqlDB.Where("id = ?",id).First(&logs)
	return b.mysqlDB.Model(logs).Update("Confirm",true).Error
}