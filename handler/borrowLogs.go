package handler

import (
	"context"
	"github.com/PonyWilliam/go-borrow-logs/domain/server"
	borrowLogs "github.com/PonyWilliam/go-borrow-logs/proto"
)

type Logs struct{
	LogsServices server.IBorrowLogsService
}
/*
	ToOther(ReqWID int64,RspWID int64,PID int64,Reason string)error
	BossConfirm(ID int64)error
*/
func(l *Logs)ToOther(ctx context.Context,req *borrowLogs.ReqToOther,rsp *borrowLogs.Rsp_All)error{
	err := l.LogsServices.ToOther(req.ReqWID,req.RspWID,req.PID,req.Reason)
	if err != nil{
		rsp.Message = err.Error()
		rsp.Status = false
		return nil
	}
	rsp.Message = "成功"
	rsp.Status = true
	return nil
}
func(l *Logs)BossConfirm(ctx context.Context,req *borrowLogs.Req_Boss_Confirm,rsp *borrowLogs.Rsp_All)error{
	err := l.LogsServices.BossConfirm(req.ID)
	if err!=nil{
		rsp.Message = err.Error()
		rsp.Status = false
		return nil
	}
	rsp.Message = "成功"
	rsp.Status = true
	return nil
}