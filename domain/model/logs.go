package model
type BorrowLogs struct{
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	ReqWID int64 `json:"req_wid"`//转借者ID
	RspWID int64 `json:"rsp_wid"`//收到者ID
	PID int64 `json:"pid"`//转借商品ID
	Confirm int64 `json:"confirm"`//1代表未确认，2代表已确认，3代表已拒绝
	Reason string `json:"reason"`//原因说明
	Boss int64 `json:"boss"`//是否需要boss确认，1代表不需要，2代表需要且未确认，3代表需要且确认
	Hashcode string `json:"hashcode"`//区块链确认，暂时不用
	Logid int64 `json:"logid"`//borrow表中的id，用于第二次confirm
	Time int64 `json:"time"`
	RspTime int64 `json:"rsp_time"`
}