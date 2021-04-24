package handler

import (
	"context"
	"fmt"
	"github.com/PonyWilliam/go-borrow-logs/domain/model"
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
	err := l.LogsServices.ToOther(req.ReqWID,req.RspWID,req.PID,req.Reason,req.Logid)
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
func(l *Logs)Reject(ctx context.Context,req *borrowLogs.Req_Reject,rsp *borrowLogs.Rsp_All)error{
	err := l.LogsServices.Reject(req.ID)
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
	result := &borrowLogs.RspLogs{}
	for _,v := range res{
		temp := &borrowLogs.Rsp_Log{}
		Swap(v,temp)
		result.Logs = append(result.Logs,temp)
	}
	rsp.Logs = result.Logs
	return nil
}
func(l *Logs)FindByReqWID(ctx context.Context,req *borrowLogs.Req_Wid,rsp *borrowLogs.RspLogs)error{
	res,err := l.LogsServices.FindByReqWID(req.Wid)
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	result := &borrowLogs.RspLogs{}
	for _,v := range res{
		temp := &borrowLogs.Rsp_Log{}
		Swap(v,temp)
		result.Logs = append(result.Logs,temp)
	}
	rsp.Logs = result.Logs
	return nil
}
func(l *Logs)FindByRspWID(ctx context.Context,req *borrowLogs.Req_Wid,rsp *borrowLogs.RspLogs)error{
	res,err := l.LogsServices.FindByRspWID(req.Wid)
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	result := &borrowLogs.RspLogs{}
	for _,v := range res{
		temp := &borrowLogs.Rsp_Log{}
		Swap(v,temp)
		result.Logs = append(result.Logs,temp)
	}
	rsp.Logs = result.Logs
	return nil
}
func(l *Logs)FindByPID(ctx context.Context,req *borrowLogs.Req_Pid,rsp *borrowLogs.RspLogs)error{
	res,err := l.LogsServices.FindByPID(req.Pid)
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	result := &borrowLogs.RspLogs{}
	for _,v := range res{
		temp := &borrowLogs.Rsp_Log{}
		Swap(v,temp)
		result.Logs = append(result.Logs,temp)
	}
	rsp.Logs = result.Logs
	return nil
}
func(l *Logs)FindByID(ctx context.Context,req *borrowLogs.Req_Id,rsp *borrowLogs.Rsp_Log)error{
	res,err := l.LogsServices.FindByID(req.Id)
	if err != nil{
		rsp = nil
		return err
	}
	Swap(res,rsp)
	fmt.Println(rsp)
	return nil
}
func(l *Logs)FindByLogID(ctx context.Context,req *borrowLogs.Req_LogID,rsp *borrowLogs.RspLogs) error{
	res,err := l.LogsServices.FindByLogID(req.Logid)
	if err != nil{
		rsp.Logs = nil
		return nil
	}
	result := &borrowLogs.RspLogs{}
	for _,v := range res{
		temp := &borrowLogs.Rsp_Log{}
		Swap(v,temp)
		result.Logs = append(result.Logs,temp)
	}
	rsp.Logs = result.Logs
	return nil
}

func Swap(req model.BorrowLogs,rsp *borrowLogs.Rsp_Log){
	rsp.RspWID = req.RspWID
	rsp.Confirm = req.Confirm
	rsp.Logid = req.Logid
	rsp.Id = req.ID
	rsp.PID = req.PID
	rsp.Reason = req.Reason
	rsp.ReqWID = req.ReqWID
	rsp.Boss = req.Boss
	rsp.Time = req.Time
	rsp.RspTime = req.RspTime
}