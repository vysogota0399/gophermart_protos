// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: commands/create_order/order_service.proto

package create_order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type NewOrder struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Number        string                 `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	State         string                 `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	UploadedAt    string                 `protobuf:"bytes,4,opt,name=uploaded_at,json=uploadedAt,proto3" json:"uploaded_at,omitempty"`
	AccountId     int64                  `protobuf:"varint,5,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewOrder) Reset() {
	*x = NewOrder{}
	mi := &file_commands_create_order_order_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrder) ProtoMessage() {}

func (x *NewOrder) ProtoReflect() protoreflect.Message {
	mi := &file_commands_create_order_order_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrder.ProtoReflect.Descriptor instead.
func (*NewOrder) Descriptor() ([]byte, []int) {
	return file_commands_create_order_order_service_proto_rawDescGZIP(), []int{0}
}

func (x *NewOrder) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *NewOrder) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *NewOrder) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *NewOrder) GetUploadedAt() string {
	if x != nil {
		return x.UploadedAt
	}
	return ""
}

func (x *NewOrder) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type NewOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewOrderResponse) Reset() {
	*x = NewOrderResponse{}
	mi := &file_commands_create_order_order_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrderResponse) ProtoMessage() {}

func (x *NewOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commands_create_order_order_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrderResponse.ProtoReflect.Descriptor instead.
func (*NewOrderResponse) Descriptor() ([]byte, []int) {
	return file_commands_create_order_order_service_proto_rawDescGZIP(), []int{1}
}

var File_commands_create_order_order_service_proto protoreflect.FileDescriptor

var file_commands_create_order_order_service_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x08,
	0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4e, 0x65,
	0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x3c,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x09,
	0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x4e, 0x65, 0x77, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_commands_create_order_order_service_proto_rawDescOnce sync.Once
	file_commands_create_order_order_service_proto_rawDescData []byte
)

func file_commands_create_order_order_service_proto_rawDescGZIP() []byte {
	file_commands_create_order_order_service_proto_rawDescOnce.Do(func() {
		file_commands_create_order_order_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_commands_create_order_order_service_proto_rawDesc), len(file_commands_create_order_order_service_proto_rawDesc)))
	})
	return file_commands_create_order_order_service_proto_rawDescData
}

var file_commands_create_order_order_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commands_create_order_order_service_proto_goTypes = []any{
	(*NewOrder)(nil),         // 0: NewOrder
	(*NewOrderResponse)(nil), // 1: NewOrderResponse
}
var file_commands_create_order_order_service_proto_depIdxs = []int32{
	0, // 0: CreateOrderService.Create:input_type -> NewOrder
	1, // 1: CreateOrderService.Create:output_type -> NewOrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commands_create_order_order_service_proto_init() }
func file_commands_create_order_order_service_proto_init() {
	if File_commands_create_order_order_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_commands_create_order_order_service_proto_rawDesc), len(file_commands_create_order_order_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commands_create_order_order_service_proto_goTypes,
		DependencyIndexes: file_commands_create_order_order_service_proto_depIdxs,
		MessageInfos:      file_commands_create_order_order_service_proto_msgTypes,
	}.Build()
	File_commands_create_order_order_service_proto = out.File
	file_commands_create_order_order_service_proto_goTypes = nil
	file_commands_create_order_order_service_proto_depIdxs = nil
}
