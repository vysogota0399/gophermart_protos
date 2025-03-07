// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/queries/accounting/withdrawal.proto

package accounting

import (
	money "google.golang.org/genproto/googleapis/type/money"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Withdrawal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderNumber   string                 `protobuf:"bytes,1,opt,name=order_number,json=orderNumber,proto3" json:"order_number,omitempty"`
	Sum           *money.Money           `protobuf:"bytes,2,opt,name=sum,proto3" json:"sum,omitempty"`
	ProcessedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=processed_at,json=processedAt,proto3" json:"processed_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Withdrawal) Reset() {
	*x = Withdrawal{}
	mi := &file_proto_queries_accounting_withdrawal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Withdrawal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdrawal) ProtoMessage() {}

func (x *Withdrawal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_queries_accounting_withdrawal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdrawal.ProtoReflect.Descriptor instead.
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return file_proto_queries_accounting_withdrawal_proto_rawDescGZIP(), []int{0}
}

func (x *Withdrawal) GetOrderNumber() string {
	if x != nil {
		return x.OrderNumber
	}
	return ""
}

func (x *Withdrawal) GetSum() *money.Money {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *Withdrawal) GetProcessedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ProcessedAt
	}
	return nil
}

var File_proto_queries_accounting_withdrawal_proto protoreflect.FileDescriptor

var file_proto_queries_accounting_withdrawal_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x71, 0x75, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x1a,
	0x34, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x3d,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x42, 0x42, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x79, 0x73, 0x6f,
	0x67, 0x6f, 0x74, 0x61, 0x30, 0x33, 0x39, 0x39, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x6d,
	0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_queries_accounting_withdrawal_proto_rawDescOnce sync.Once
	file_proto_queries_accounting_withdrawal_proto_rawDescData []byte
)

func file_proto_queries_accounting_withdrawal_proto_rawDescGZIP() []byte {
	file_proto_queries_accounting_withdrawal_proto_rawDescOnce.Do(func() {
		file_proto_queries_accounting_withdrawal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_queries_accounting_withdrawal_proto_rawDesc), len(file_proto_queries_accounting_withdrawal_proto_rawDesc)))
	})
	return file_proto_queries_accounting_withdrawal_proto_rawDescData
}

var file_proto_queries_accounting_withdrawal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_queries_accounting_withdrawal_proto_goTypes = []any{
	(*Withdrawal)(nil),            // 0: queries.accounting.Withdrawal
	(*money.Money)(nil),           // 1: google.type.Money
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_queries_accounting_withdrawal_proto_depIdxs = []int32{
	1, // 0: queries.accounting.Withdrawal.sum:type_name -> google.type.Money
	2, // 1: queries.accounting.Withdrawal.processed_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_queries_accounting_withdrawal_proto_init() }
func file_proto_queries_accounting_withdrawal_proto_init() {
	if File_proto_queries_accounting_withdrawal_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_queries_accounting_withdrawal_proto_rawDesc), len(file_proto_queries_accounting_withdrawal_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_queries_accounting_withdrawal_proto_goTypes,
		DependencyIndexes: file_proto_queries_accounting_withdrawal_proto_depIdxs,
		MessageInfos:      file_proto_queries_accounting_withdrawal_proto_msgTypes,
	}.Build()
	File_proto_queries_accounting_withdrawal_proto = out.File
	file_proto_queries_accounting_withdrawal_proto_goTypes = nil
	file_proto_queries_accounting_withdrawal_proto_depIdxs = nil
}
