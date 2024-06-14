// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.4
// source: payment.proto

package pb

import (
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

type CreatePaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PayTotal int64  `protobuf:"varint,2,opt,name=payTotal,proto3" json:"payTotal,omitempty"`
	OrderSn  string `protobuf:"bytes,3,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
	Platform int64  `protobuf:"varint,4,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *CreatePaymentReq) Reset() {
	*x = CreatePaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentReq) ProtoMessage() {}

func (x *CreatePaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentReq.ProtoReflect.Descriptor instead.
func (*CreatePaymentReq) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaymentReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CreatePaymentReq) GetPayTotal() int64 {
	if x != nil {
		return x.PayTotal
	}
	return 0
}

func (x *CreatePaymentReq) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *CreatePaymentReq) GetPlatform() int64 {
	if x != nil {
		return x.Platform
	}
	return 0
}

type CreatePaymentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn string `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
}

func (x *CreatePaymentResp) Reset() {
	*x = CreatePaymentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResp) ProtoMessage() {}

func (x *CreatePaymentResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResp.ProtoReflect.Descriptor instead.
func (*CreatePaymentResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePaymentResp) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

type PaymentDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sn         string  `protobuf:"bytes,2,opt,name=sn,proto3" json:"sn,omitempty"`
	Uid        int64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	TradeState int64   `protobuf:"varint,4,opt,name=tradeState,proto3" json:"tradeState,omitempty"` // 交易状态  0:未支付 1:支付成功 -1:支付失败
	PayTotal   float64 `protobuf:"fixed64,5,opt,name=payTotal,proto3" json:"payTotal,omitempty"`
	OrderSn    string  `protobuf:"bytes,6,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
	CreateTime int64   `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime int64   `protobuf:"varint,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	PayTime    int64   `protobuf:"varint,9,opt,name=payTime,proto3" json:"payTime,omitempty"`
}

func (x *PaymentDetail) Reset() {
	*x = PaymentDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentDetail) ProtoMessage() {}

func (x *PaymentDetail) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentDetail.ProtoReflect.Descriptor instead.
func (*PaymentDetail) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentDetail) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *PaymentDetail) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PaymentDetail) GetTradeState() int64 {
	if x != nil {
		return x.TradeState
	}
	return 0
}

func (x *PaymentDetail) GetPayTotal() float64 {
	if x != nil {
		return x.PayTotal
	}
	return 0
}

func (x *PaymentDetail) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *PaymentDetail) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *PaymentDetail) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *PaymentDetail) GetPayTime() int64 {
	if x != nil {
		return x.PayTime
	}
	return 0
}

type GetPaymentBySnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn string `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
}

func (x *GetPaymentBySnReq) Reset() {
	*x = GetPaymentBySnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentBySnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentBySnReq) ProtoMessage() {}

func (x *GetPaymentBySnReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentBySnReq.ProtoReflect.Descriptor instead.
func (*GetPaymentBySnReq) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentBySnReq) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

type GetPaymentBySnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentDetail *PaymentDetail `protobuf:"bytes,1,opt,name=paymentDetail,proto3" json:"paymentDetail,omitempty"`
}

func (x *GetPaymentBySnResp) Reset() {
	*x = GetPaymentBySnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentBySnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentBySnResp) ProtoMessage() {}

func (x *GetPaymentBySnResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentBySnResp.ProtoReflect.Descriptor instead.
func (*GetPaymentBySnResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{4}
}

func (x *GetPaymentBySnResp) GetPaymentDetail() *PaymentDetail {
	if x != nil {
		return x.PaymentDetail
	}
	return nil
}

type GetPaymentSuccessRefundByOrderSnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderSn string `protobuf:"bytes,1,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
}

func (x *GetPaymentSuccessRefundByOrderSnReq) Reset() {
	*x = GetPaymentSuccessRefundByOrderSnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentSuccessRefundByOrderSnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentSuccessRefundByOrderSnReq) ProtoMessage() {}

func (x *GetPaymentSuccessRefundByOrderSnReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentSuccessRefundByOrderSnReq.ProtoReflect.Descriptor instead.
func (*GetPaymentSuccessRefundByOrderSnReq) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{5}
}

func (x *GetPaymentSuccessRefundByOrderSnReq) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

type GetPaymentSuccessRefundByOrderSnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentDetail *PaymentDetail `protobuf:"bytes,1,opt,name=paymentDetail,proto3" json:"paymentDetail,omitempty"`
}

func (x *GetPaymentSuccessRefundByOrderSnResp) Reset() {
	*x = GetPaymentSuccessRefundByOrderSnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentSuccessRefundByOrderSnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentSuccessRefundByOrderSnResp) ProtoMessage() {}

func (x *GetPaymentSuccessRefundByOrderSnResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentSuccessRefundByOrderSnResp.ProtoReflect.Descriptor instead.
func (*GetPaymentSuccessRefundByOrderSnResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{6}
}

func (x *GetPaymentSuccessRefundByOrderSnResp) GetPaymentDetail() *PaymentDetail {
	if x != nil {
		return x.PaymentDetail
	}
	return nil
}

// 更新交易状态
type UpdateTradeStateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn         string `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
	TradeState int64  `protobuf:"varint,2,opt,name=tradeState,proto3" json:"tradeState,omitempty"`
}

func (x *UpdateTradeStateReq) Reset() {
	*x = UpdateTradeStateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTradeStateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTradeStateReq) ProtoMessage() {}

func (x *UpdateTradeStateReq) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTradeStateReq.ProtoReflect.Descriptor instead.
func (*UpdateTradeStateReq) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTradeStateReq) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *UpdateTradeStateReq) GetTradeState() int64 {
	if x != nil {
		return x.TradeState
	}
	return 0
}

type UpdateTradeStateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTradeStateResp) Reset() {
	*x = UpdateTradeStateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTradeStateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTradeStateResp) ProtoMessage() {}

func (x *UpdateTradeStateResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTradeStateResp.ProtoReflect.Descriptor instead.
func (*UpdateTradeStateResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{8}
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x76, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x79,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x79,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x23, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6e,
	0x22, 0xf1, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x73, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x53, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6e, 0x22, 0x4d, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x53, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x37, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x3f, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x22, 0x5f, 0x0a, 0x24, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x37, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x45, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32, 0xc6, 0x02, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x53, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x53, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x53, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x75, 0x0a, 0x20, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12,
	0x27, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_payment_proto_goTypes = []interface{}{
	(*CreatePaymentReq)(nil),                     // 0: pb.CreatePaymentReq
	(*CreatePaymentResp)(nil),                    // 1: pb.CreatePaymentResp
	(*PaymentDetail)(nil),                        // 2: pb.PaymentDetail
	(*GetPaymentBySnReq)(nil),                    // 3: pb.GetPaymentBySnReq
	(*GetPaymentBySnResp)(nil),                   // 4: pb.GetPaymentBySnResp
	(*GetPaymentSuccessRefundByOrderSnReq)(nil),  // 5: pb.GetPaymentSuccessRefundByOrderSnReq
	(*GetPaymentSuccessRefundByOrderSnResp)(nil), // 6: pb.GetPaymentSuccessRefundByOrderSnResp
	(*UpdateTradeStateReq)(nil),                  // 7: pb.UpdateTradeStateReq
	(*UpdateTradeStateResp)(nil),                 // 8: pb.UpdateTradeStateResp
}
var file_payment_proto_depIdxs = []int32{
	2, // 0: pb.GetPaymentBySnResp.paymentDetail:type_name -> pb.PaymentDetail
	2, // 1: pb.GetPaymentSuccessRefundByOrderSnResp.paymentDetail:type_name -> pb.PaymentDetail
	0, // 2: pb.payment.CreatePayment:input_type -> pb.CreatePaymentReq
	3, // 3: pb.payment.GetPaymentBySn:input_type -> pb.GetPaymentBySnReq
	7, // 4: pb.payment.UpdateTradeState:input_type -> pb.UpdateTradeStateReq
	5, // 5: pb.payment.GetPaymentSuccessRefundByOrderSn:input_type -> pb.GetPaymentSuccessRefundByOrderSnReq
	1, // 6: pb.payment.CreatePayment:output_type -> pb.CreatePaymentResp
	4, // 7: pb.payment.GetPaymentBySn:output_type -> pb.GetPaymentBySnResp
	8, // 8: pb.payment.UpdateTradeState:output_type -> pb.UpdateTradeStateResp
	6, // 9: pb.payment.GetPaymentSuccessRefundByOrderSn:output_type -> pb.GetPaymentSuccessRefundByOrderSnResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentReq); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentResp); i {
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
		file_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentDetail); i {
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
		file_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentBySnReq); i {
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
		file_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentBySnResp); i {
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
		file_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentSuccessRefundByOrderSnReq); i {
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
		file_payment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentSuccessRefundByOrderSnResp); i {
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
		file_payment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTradeStateReq); i {
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
		file_payment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTradeStateResp); i {
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
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}