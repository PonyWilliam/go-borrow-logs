package server

import "github.com/PonyWilliam/go-borrow-logs/domain/repository"

/*
	service borrowLogs {
		rpc to_other(Req_to_other)returns(Rsp_All);
		rpc boss_confirm(Req_Boss_Confirm)returns(Rsp_All);
	}
*/
type IBorrowLogsService interface {
	ToOther(ReqWID int64,RspWID int64,PID int64,Reason string)error
	BossConfirm(ID int64)error
}
func NewLogsService(log repository.BorrowLogsRepository)IBorrowLogsService{
	return &LogsServices{log}
}
type LogsServices struct{
	Logs repository.BorrowLogsRepository
}
func(l *LogsServices)ToOther(ReqWID int64,RspWID int64,PID int64,Reason string)error{
	return l.Logs.ToOther(ReqWID,RspWID,PID,Reason)
}
func(l *LogsServices)BossConfirm(ID int64)error{
	return l.Logs.Confirm(ID)
}