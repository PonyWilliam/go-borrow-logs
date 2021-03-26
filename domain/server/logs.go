package server

import (
	"github.com/PonyWilliam/go-borrow-logs/domain/model"
	"github.com/PonyWilliam/go-borrow-logs/domain/repository"
)

/*
	service borrowLogs {
		rpc to_other(Req_to_other)returns(Rsp_All);
		rpc boss_confirm(Req_Boss_Confirm)returns(Rsp_All);
	}
*/
type IBorrowLogsService interface {
	ToOther(ReqWID int64,RspWID int64,PID int64,Reason string,Logid int64)error
	Confirm(ID int64)error
	FindAll() ([]model.BorrowLogs,error)
	FindByID(id int64) (model.BorrowLogs,error)
	FindByWID(wid int64)([]model.BorrowLogs,error)
	FindByPID(pid int64)([]model.BorrowLogs,error)
	FindByLogID(logid int64)([]model.BorrowLogs,error)
}
func NewLogsService(log repository.ILogs)IBorrowLogsService{
	return &LogsServices{log}
}
type LogsServices struct{
	Logs repository.ILogs
}
func(l *LogsServices)ToOther(ReqWID int64,RspWID int64,PID int64,Reason string,Logid int64)error{
	return l.Logs.ToOther(ReqWID,RspWID,PID,Reason,Logid)
}
func(l *LogsServices)Confirm(ID int64)error{
	return l.Logs.Confirm(ID)
}
func(l *LogsServices)FindAll() ([]model.BorrowLogs,error){
	return l.Logs.FindAll()
}
func(l *LogsServices)FindByID(id int64) (model.BorrowLogs,error){
	return l.Logs.FindByID(id)
}
func(l *LogsServices)FindByWID(wid int64) ([]model.BorrowLogs,error){
	return l.Logs.FindByWID(wid)
}
func(l *LogsServices)FindByPID(pid int64) ([]model.BorrowLogs,error){
	return l.Logs.FindByPID(pid)
}
func(l *LogsServices)FindByLogID(logid int64)([]model.BorrowLogs,error){
	return l.Logs.FindByLogID(logid)
}