package handler

import (
	"context"
	"github.com/PonyWilliam/go-borrow-logs/domain/server"
	borrowLogs "github.com/PonyWilliam/go-borrow-logs/proto"
	"github.com/PonyWilliam/go-common"
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
func(l *Logs)Confirm(ctx context.Context,req *borrowLogs.Req_Confirm,rsp *borrowLogs.Rsp_All)error{
	err := l.LogsServices.Confirm(req.ID)
	if err!=nil{
		rsp.Message = err.Error()
		rsp.Status = false
		return nil
	}
	rsp.Message = "成功"
	rsp.Status = true
	return nil
}
func(l *Logs)FindAll(ctx context.Context,req *borrowLogs.Req_Null,rsp *borrowLogs.RspLogs)error{
	res,err := l.LogsServices.FindAll()
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	temp := &borrowLogs.Rsp_Log{}
	for _,v := range res{
		_ = common.SwapTo(v, temp)
		rsp.Logs = append(rsp.Logs,temp)
	}
	return nil
}
func(l *Logs)FindByWID(ctx context.Context,req *borrowLogs.Req_Wid,rsp *borrowLogs.RspLogs)error{
	res,err := l.LogsServices.FindByWID(req.Wid)
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	temp := &borrowLogs.Rsp_Log{}
	for _,v := range res{
		_ = common.SwapTo(v, temp)
		rsp.Logs = append(rsp.Logs,temp)
	}
	return nil
}
func(l *Logs)FindByPID(ctx context.Context,req *borrowLogs.Req_Pid,rsp *borrowLogs.RspLogs)error{
	res,err := l.LogsServices.FindByPID(req.Pid)
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	temp := &borrowLogs.Rsp_Log{}
	for _,v := range res{
		_ = common.SwapTo(v, temp)
		rsp.Logs = append(rsp.Logs,temp)
	}
	return nil
}
func(l *Logs)FindByID(ctx context.Context,req *borrowLogs.Req_Id,rsp *borrowLogs.Rsp_Log)error{
	res,err := l.LogsServices.FindByID(req.Id)
	if err != nil{
		rsp = nil
		return nil
	}
	_ = common.SwapTo(res, rsp)
	return nil
}
func(l *Logs)FindByLogID(ctx context.Context,req *borrowLogs.Req_LogID,rsp *borrowLogs.RspLogs) error{
	res,err := l.LogsServices.FindByLogID(req.Logid)
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	temp := &borrowLogs.Rsp_Log{}
	for _,v := range res{
		_ = common.SwapTo(v, temp)
		rsp.Logs = append(rsp.Logs,temp)
	}
	return nil
}