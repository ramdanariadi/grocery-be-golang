// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: proto/transaction.proto

package transaction

import (
	"github.com/ramdanariadi/grocery-product-service/main/response"
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

type TransactionUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransactionUserId) Reset() {
	*x = TransactionUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionUserId) ProtoMessage() {}

func (x *TransactionUserId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionUserId.ProtoReflect.Descriptor instead.
func (*TransactionUserId) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionUserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TransactionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransactionId) Reset() {
	*x = TransactionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionId) ProtoMessage() {}

func (x *TransactionId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionId.ProtoReflect.Descriptor instead.
func (*TransactionId) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TransactionProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Total     uint32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *TransactionProduct) Reset() {
	*x = TransactionProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionProduct) ProtoMessage() {}

func (x *TransactionProduct) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionProduct.ProtoReflect.Descriptor instead.
func (*TransactionProduct) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *TransactionProduct) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type TransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string                `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Products []*TransactionProduct `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *TransactionBody) Reset() {
	*x = TransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBody) ProtoMessage() {}

func (x *TransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBody.ProtoReflect.Descriptor instead.
func (*TransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionBody) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TransactionBody) GetProducts() []*TransactionProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type TransactionProductDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Category    string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	ImageUrl    string `protobuf:"bytes,4,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Price       uint64 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Weight      uint32 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	CategoryId  string `protobuf:"bytes,7,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	PerUnit     uint64 `protobuf:"varint,8,opt,name=per_unit,json=perUnit,proto3" json:"per_unit,omitempty"`
	Description string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Total       uint32 `protobuf:"varint,10,opt,name=total,proto3" json:"total,omitempty"`
	ProductId   string `protobuf:"bytes,11,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *TransactionProductDetail) Reset() {
	*x = TransactionProductDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionProductDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionProductDetail) ProtoMessage() {}

func (x *TransactionProductDetail) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionProductDetail.ProtoReflect.Descriptor instead.
func (*TransactionProductDetail) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionProductDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TransactionProductDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransactionProductDetail) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *TransactionProductDetail) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *TransactionProductDetail) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TransactionProductDetail) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *TransactionProductDetail) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *TransactionProductDetail) GetPerUnit() uint64 {
	if x != nil {
		return x.PerUnit
	}
	return 0
}

func (x *TransactionProductDetail) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TransactionProductDetail) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TransactionProductDetail) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string                      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalPrice      uint64                      `protobuf:"varint,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	TransactionDate int64                       `protobuf:"varint,4,opt,name=transaction_date,json=transactionDate,proto3" json:"transaction_date,omitempty"`
	Products        []*TransactionProductDetail `protobuf:"bytes,5,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Transaction) GetTotalPrice() uint64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Transaction) GetTransactionDate() int64 {
	if x != nil {
		return x.TransactionDate
	}
	return 0
}

func (x *Transaction) GetProducts() []*TransactionProductDetail {
	if x != nil {
		return x.Products
	}
	return nil
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Transaction `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *TransactionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransactionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransactionResponse) GetData() *Transaction {
	if x != nil {
		return x.Data
	}
	return nil
}

type MultipleTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Transaction `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MultipleTransactionResponse) Reset() {
	*x = MultipleTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleTransactionResponse) ProtoMessage() {}

func (x *MultipleTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleTransactionResponse.ProtoReflect.Descriptor instead.
func (*MultipleTransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{7}
}

func (x *MultipleTransactionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MultipleTransactionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MultipleTransactionResponse) GetData() []*Transaction {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_transaction_proto protoreflect.FileDescriptor

var file_proto_transaction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x12,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x61, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xb7, 0x02, 0x0a, 0x18, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x22, 0xbf, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x6f, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x1b, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x8d, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2c, 0x0a, 0x15, 0x69, 0x64, 0x2e, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x2e, 0x74,
	0x75, 0x6e, 0x61, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x11, 0x2f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transaction_proto_rawDescOnce sync.Once
	file_proto_transaction_proto_rawDescData = file_proto_transaction_proto_rawDesc
)

func file_proto_transaction_proto_rawDescGZIP() []byte {
	file_proto_transaction_proto_rawDescOnce.Do(func() {
		file_proto_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transaction_proto_rawDescData)
	})
	return file_proto_transaction_proto_rawDescData
}

var file_proto_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_transaction_proto_goTypes = []interface{}{
	(*TransactionUserId)(nil),           // 0: proto.TransactionUserId
	(*TransactionId)(nil),               // 1: proto.TransactionId
	(*TransactionProduct)(nil),          // 2: proto.TransactionProduct
	(*TransactionBody)(nil),             // 3: proto.TransactionBody
	(*TransactionProductDetail)(nil),    // 4: proto.TransactionProductDetail
	(*Transaction)(nil),                 // 5: proto.Transaction
	(*TransactionResponse)(nil),         // 6: proto.TransactionResponse
	(*MultipleTransactionResponse)(nil), // 7: proto.MultipleTransactionResponse
	(*response.Response)(nil),           // 8: proto.Response
}
var file_proto_transaction_proto_depIdxs = []int32{
	2, // 0: proto.TransactionBody.products:type_name -> proto.TransactionProduct
	4, // 1: proto.Transaction.products:type_name -> proto.TransactionProductDetail
	5, // 2: proto.TransactionResponse.data:type_name -> proto.Transaction
	5, // 3: proto.MultipleTransactionResponse.data:type_name -> proto.Transaction
	1, // 4: proto.TransactionService.FindByTransactionId:input_type -> proto.TransactionId
	0, // 5: proto.TransactionService.FindByUserId:input_type -> proto.TransactionUserId
	3, // 6: proto.TransactionService.Save:input_type -> proto.TransactionBody
	1, // 7: proto.TransactionService.Delete:input_type -> proto.TransactionId
	6, // 8: proto.TransactionService.FindByTransactionId:output_type -> proto.TransactionResponse
	7, // 9: proto.TransactionService.FindByUserId:output_type -> proto.MultipleTransactionResponse
	8, // 10: proto.TransactionService.Save:output_type -> proto.Response
	8, // 11: proto.TransactionService.Delete:output_type -> proto.Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_transaction_proto_init() }
func file_proto_transaction_proto_init() {
	if File_proto_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionUserId); i {
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
		file_proto_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionId); i {
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
		file_proto_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionProduct); i {
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
		file_proto_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBody); i {
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
		file_proto_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionProductDetail); i {
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
		file_proto_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_proto_transaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
		file_proto_transaction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleTransactionResponse); i {
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
			RawDescriptor: file_proto_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_transaction_proto_goTypes,
		DependencyIndexes: file_proto_transaction_proto_depIdxs,
		MessageInfos:      file_proto_transaction_proto_msgTypes,
	}.Build()
	File_proto_transaction_proto = out.File
	file_proto_transaction_proto_rawDesc = nil
	file_proto_transaction_proto_goTypes = nil
	file_proto_transaction_proto_depIdxs = nil
}
