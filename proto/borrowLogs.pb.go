// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.0
// source: proto/borrowLogs.proto

//
//ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
//ReqWID int64 `json:"req_wid"`//转借者ID
//RspWID int64 `json:"rsp_wid"`//收到者ID
//Confirm bool `json:"confirm"`//对方是否已确认
//Reason string `json:"reason"`//原因说明
//Boss int8 `json:"boss"`//是否需要boss确认，1代表不需要，2代表需要且未确认，3代表需要且确认

package borrowLogs

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Req_Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Req_Null) Reset() {
	*x = Req_Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_Null) ProtoMessage() {}

func (x *Req_Null) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_Null.ProtoReflect.Descriptor instead.
func (*Req_Null) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{0}
}

type Req_LogID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logid int64 `protobuf:"varint,1,opt,name=logid,proto3" json:"logid"`
}

func (x *Req_LogID) Reset() {
	*x = Req_LogID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_LogID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_LogID) ProtoMessage() {}

func (x *Req_LogID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_LogID.ProtoReflect.Descriptor instead.
func (*Req_LogID) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{1}
}

func (x *Req_LogID) GetLogid() int64 {
	if x != nil {
		return x.Logid
	}
	return 0
}

type Req_Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *Req_Id) Reset() {
	*x = Req_Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_Id) ProtoMessage() {}

func (x *Req_Id) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_Id.ProtoReflect.Descriptor instead.
func (*Req_Id) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{2}
}

func (x *Req_Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Req_Wid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wid int64 `protobuf:"varint,1,opt,name=wid,proto3" json:"wid"`
}

func (x *Req_Wid) Reset() {
	*x = Req_Wid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_Wid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_Wid) ProtoMessage() {}

func (x *Req_Wid) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_Wid.ProtoReflect.Descriptor instead.
func (*Req_Wid) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{3}
}

func (x *Req_Wid) GetWid() int64 {
	if x != nil {
		return x.Wid
	}
	return 0
}

type Req_Pid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int64 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid"`
}

func (x *Req_Pid) Reset() {
	*x = Req_Pid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_Pid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_Pid) ProtoMessage() {}

func (x *Req_Pid) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_Pid.ProtoReflect.Descriptor instead.
func (*Req_Pid) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{4}
}

func (x *Req_Pid) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type Rsp_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//type BorrowLogs struct{
	//ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	//ReqWID int64 `json:"req_wid"`//转借者ID
	//RspWID int64 `json:"rsp_wid"`//收到者ID
	//PID int64 `json:"pid"`//转借商品ID
	//Confirm bool `json:"confirm"`//对方是否已确认
	//Reason string `json:"reason"`//原因说明
	//Boss int8 `json:"boss"`//是否需要boss确认，1代表不需要，2代表需要且未确认，3代表需要且确认
	//Hashcode string `json:"hashcode"`//区块链确认，暂时不用
	//}
	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ReqWID  int64  `protobuf:"varint,2,opt,name=ReqWID,proto3" json:"ReqWID"`
	RspWID  int64  `protobuf:"varint,3,opt,name=RspWID,proto3" json:"RspWID"`
	PID     int64  `protobuf:"varint,4,opt,name=PID,proto3" json:"PID"`
	Confirm int64  `protobuf:"varint,5,opt,name=Confirm,proto3" json:"Confirm"`
	Reason  string `protobuf:"bytes,6,opt,name=Reason,proto3" json:"Reason"`
	Boss    int64  `protobuf:"varint,7,opt,name=Boss,proto3" json:"Boss"`
	Logid   int64  `protobuf:"varint,8,opt,name=Logid,proto3" json:"Logid"`
}

func (x *Rsp_Log) Reset() {
	*x = Rsp_Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rsp_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rsp_Log) ProtoMessage() {}

func (x *Rsp_Log) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rsp_Log.ProtoReflect.Descriptor instead.
func (*Rsp_Log) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{5}
}

func (x *Rsp_Log) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rsp_Log) GetReqWID() int64 {
	if x != nil {
		return x.ReqWID
	}
	return 0
}

func (x *Rsp_Log) GetRspWID() int64 {
	if x != nil {
		return x.RspWID
	}
	return 0
}

func (x *Rsp_Log) GetPID() int64 {
	if x != nil {
		return x.PID
	}
	return 0
}

func (x *Rsp_Log) GetConfirm() int64 {
	if x != nil {
		return x.Confirm
	}
	return 0
}

func (x *Rsp_Log) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Rsp_Log) GetBoss() int64 {
	if x != nil {
		return x.Boss
	}
	return 0
}

func (x *Rsp_Log) GetLogid() int64 {
	if x != nil {
		return x.Logid
	}
	return 0
}

type RspLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Rsp_Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs"`
}

func (x *RspLogs) Reset() {
	*x = RspLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspLogs) ProtoMessage() {}

func (x *RspLogs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspLogs.ProtoReflect.Descriptor instead.
func (*RspLogs) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{6}
}

func (x *RspLogs) GetLogs() []*Rsp_Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type ReqToOther struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqWID int64  `protobuf:"varint,1,opt,name=ReqWID,proto3" json:"ReqWID"`
	RspWID int64  `protobuf:"varint,2,opt,name=RspWID,proto3" json:"RspWID"`
	PID    int64  `protobuf:"varint,3,opt,name=PID,proto3" json:"PID"` //产品ID
	Reason string `protobuf:"bytes,4,opt,name=Reason,proto3" json:"Reason"`
	Logid  int64  `protobuf:"varint,5,opt,name=Logid,proto3" json:"Logid"`
}

func (x *ReqToOther) Reset() {
	*x = ReqToOther{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqToOther) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqToOther) ProtoMessage() {}

func (x *ReqToOther) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqToOther.ProtoReflect.Descriptor instead.
func (*ReqToOther) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{7}
}

func (x *ReqToOther) GetReqWID() int64 {
	if x != nil {
		return x.ReqWID
	}
	return 0
}

func (x *ReqToOther) GetRspWID() int64 {
	if x != nil {
		return x.RspWID
	}
	return 0
}

func (x *ReqToOther) GetPID() int64 {
	if x != nil {
		return x.PID
	}
	return 0
}

func (x *ReqToOther) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ReqToOther) GetLogid() int64 {
	if x != nil {
		return x.Logid
	}
	return 0
}

type Req_Confirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID"` //根据ID确认同意
}

func (x *Req_Confirm) Reset() {
	*x = Req_Confirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_Confirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_Confirm) ProtoMessage() {}

func (x *Req_Confirm) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_Confirm.ProtoReflect.Descriptor instead.
func (*Req_Confirm) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{8}
}

func (x *Req_Confirm) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type Req_Reject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID"`
}

func (x *Req_Reject) Reset() {
	*x = Req_Reject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req_Reject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req_Reject) ProtoMessage() {}

func (x *Req_Reject) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req_Reject.ProtoReflect.Descriptor instead.
func (*Req_Reject) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{9}
}

func (x *Req_Reject) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type Rsp_All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status"` //是否成功
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
}

func (x *Rsp_All) Reset() {
	*x = Rsp_All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_borrowLogs_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rsp_All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rsp_All) ProtoMessage() {}

func (x *Rsp_All) ProtoReflect() protoreflect.Message {
	mi := &file_proto_borrowLogs_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rsp_All.ProtoReflect.Descriptor instead.
func (*Rsp_All) Descriptor() ([]byte, []int) {
	return file_proto_borrowLogs_proto_rawDescGZIP(), []int{10}
}

func (x *Rsp_All) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Rsp_All) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_borrowLogs_proto protoreflect.FileDescriptor

var file_proto_borrowLogs_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x4c, 0x6f, 0x67, 0x73, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x5f, 0x4e, 0x75, 0x6c, 0x6c,
	0x22, 0x21, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x5f, 0x4c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x5f, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x5f, 0x57, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x77, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x5f, 0x50, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x07, 0x52, 0x73, 0x70, 0x5f,
	0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x57, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x65, 0x71, 0x57, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x73, 0x70, 0x57, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x73, 0x70,
	0x57, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x50, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x73, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x42, 0x6f, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x64, 0x22, 0x33, 0x0a, 0x08, 0x52, 0x73, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x27, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x4c, 0x6f, 0x67,
	0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x7e, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x5f, 0x74, 0x6f,
	0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x57, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x65, 0x71, 0x57, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x73, 0x70, 0x57, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x52, 0x73, 0x70, 0x57, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x50, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x5f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1c, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x5f, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x07, 0x52, 0x73, 0x70, 0x5f, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x8c, 0x04, 0x0a, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x38, 0x0a, 0x07, 0x54, 0x6f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x74, 0x6f, 0x5f, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x41, 0x6c, 0x6c, 0x12, 0x37, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x1a, 0x13, 0x2e,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x41,
	0x6c, 0x6c, 0x12, 0x35, 0x0a, 0x06, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x41, 0x6c, 0x6c, 0x12, 0x35, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x4e, 0x75, 0x6c, 0x6c, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x39, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x52, 0x65, 0x71, 0x57, 0x49, 0x44,
	0x12, 0x13, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x5f, 0x57, 0x69, 0x64, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f,
	0x67, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0c, 0x46,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x52, 0x73, 0x70, 0x57, 0x49, 0x44, 0x12, 0x13, 0x2e, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x57, 0x69, 0x64,
	0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x73,
	0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x50, 0x49, 0x44, 0x12, 0x13, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73,
	0x2e, 0x52, 0x65, 0x71, 0x5f, 0x50, 0x69, 0x64, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x33,
	0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x5f, 0x49, 0x64, 0x1a, 0x13,
	0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f,
	0x4c, 0x6f, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x4c, 0x6f, 0x67,
	0x49, 0x44, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e,
	0x52, 0x65, 0x71, 0x5f, 0x4c, 0x6f, 0x67, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x52, 0x73, 0x70, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x42,
	0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x4c,
	0x6f, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_borrowLogs_proto_rawDescOnce sync.Once
	file_proto_borrowLogs_proto_rawDescData = file_proto_borrowLogs_proto_rawDesc
)

func file_proto_borrowLogs_proto_rawDescGZIP() []byte {
	file_proto_borrowLogs_proto_rawDescOnce.Do(func() {
		file_proto_borrowLogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_borrowLogs_proto_rawDescData)
	})
	return file_proto_borrowLogs_proto_rawDescData
}

var file_proto_borrowLogs_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_borrowLogs_proto_goTypes = []interface{}{
	(*Req_Null)(nil),    // 0: borrowLogs.Req_Null
	(*Req_LogID)(nil),   // 1: borrowLogs.Req_LogID
	(*Req_Id)(nil),      // 2: borrowLogs.Req_Id
	(*Req_Wid)(nil),     // 3: borrowLogs.Req_Wid
	(*Req_Pid)(nil),     // 4: borrowLogs.Req_Pid
	(*Rsp_Log)(nil),     // 5: borrowLogs.Rsp_Log
	(*RspLogs)(nil),     // 6: borrowLogs.Rsp_logs
	(*ReqToOther)(nil),  // 7: borrowLogs.Req_to_other
	(*Req_Confirm)(nil), // 8: borrowLogs.Req_Confirm
	(*Req_Reject)(nil),  // 9: borrowLogs.Req_Reject
	(*Rsp_All)(nil),     // 10: borrowLogs.Rsp_All
}
var file_proto_borrowLogs_proto_depIdxs = []int32{
	5,  // 0: borrowLogs.Rsp_logs.logs:type_name -> borrowLogs.Rsp_Log
	7,  // 1: borrowLogs.borrowLogs.ToOther:input_type -> borrowLogs.Req_to_other
	8,  // 2: borrowLogs.borrowLogs.Confirm:input_type -> borrowLogs.Req_Confirm
	9,  // 3: borrowLogs.borrowLogs.Reject:input_type -> borrowLogs.Req_Reject
	0,  // 4: borrowLogs.borrowLogs.FindAll:input_type -> borrowLogs.Req_Null
	3,  // 5: borrowLogs.borrowLogs.FindByReqWID:input_type -> borrowLogs.Req_Wid
	3,  // 6: borrowLogs.borrowLogs.FindByRspWID:input_type -> borrowLogs.Req_Wid
	4,  // 7: borrowLogs.borrowLogs.FindByPID:input_type -> borrowLogs.Req_Pid
	2,  // 8: borrowLogs.borrowLogs.FindByID:input_type -> borrowLogs.Req_Id
	1,  // 9: borrowLogs.borrowLogs.FindByLogID:input_type -> borrowLogs.Req_LogID
	10, // 10: borrowLogs.borrowLogs.ToOther:output_type -> borrowLogs.Rsp_All
	10, // 11: borrowLogs.borrowLogs.Confirm:output_type -> borrowLogs.Rsp_All
	10, // 12: borrowLogs.borrowLogs.Reject:output_type -> borrowLogs.Rsp_All
	6,  // 13: borrowLogs.borrowLogs.FindAll:output_type -> borrowLogs.Rsp_logs
	6,  // 14: borrowLogs.borrowLogs.FindByReqWID:output_type -> borrowLogs.Rsp_logs
	6,  // 15: borrowLogs.borrowLogs.FindByRspWID:output_type -> borrowLogs.Rsp_logs
	6,  // 16: borrowLogs.borrowLogs.FindByPID:output_type -> borrowLogs.Rsp_logs
	5,  // 17: borrowLogs.borrowLogs.FindByID:output_type -> borrowLogs.Rsp_Log
	6,  // 18: borrowLogs.borrowLogs.FindByLogID:output_type -> borrowLogs.Rsp_logs
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_borrowLogs_proto_init() }
func file_proto_borrowLogs_proto_init() {
	if File_proto_borrowLogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_borrowLogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_Null); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_LogID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_Id); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_Wid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_Pid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rsp_Log); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspLogs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqToOther); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_Confirm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req_Reject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_borrowLogs_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rsp_All); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_borrowLogs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_borrowLogs_proto_goTypes,
		DependencyIndexes: file_proto_borrowLogs_proto_depIdxs,
		MessageInfos:      file_proto_borrowLogs_proto_msgTypes,
	}.Build()
	File_proto_borrowLogs_proto = out.File
	file_proto_borrowLogs_proto_rawDesc = nil
	file_proto_borrowLogs_proto_goTypes = nil
	file_proto_borrowLogs_proto_depIdxs = nil
}
