// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/events/accrual_started_event.proto

package events

import (
	common "github.com/vysogota0399/gophermart_protos/gen/common"
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

type AccrualStartedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventUuid     *common.Uuid           `protobuf:"bytes,1,opt,name=event_uuid,json=eventUuid,proto3" json:"event_uuid,omitempty"`
	OrderUuid     *common.Uuid           `protobuf:"bytes,2,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	OrderNumber   string                 `protobuf:"bytes,3,opt,name=order_number,json=orderNumber,proto3" json:"order_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccrualStartedEvent) Reset() {
	*x = AccrualStartedEvent{}
	mi := &file_proto_events_accrual_started_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccrualStartedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccrualStartedEvent) ProtoMessage() {}

func (x *AccrualStartedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_events_accrual_started_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccrualStartedEvent.ProtoReflect.Descriptor instead.
func (*AccrualStartedEvent) Descriptor() ([]byte, []int) {
	return file_proto_events_accrual_started_event_proto_rawDescGZIP(), []int{0}
}

func (x *AccrualStartedEvent) GetEventUuid() *common.Uuid {
	if x != nil {
		return x.EventUuid
	}
	return nil
}

func (x *AccrualStartedEvent) GetOrderUuid() *common.Uuid {
	if x != nil {
		return x.OrderUuid
	}
	return nil
}

func (x *AccrualStartedEvent) GetOrderNumber() string {
	if x != nil {
		return x.OrderNumber
	}
	return ""
}

var File_proto_events_accrual_started_event_proto protoreflect.FileDescriptor

var file_proto_events_accrual_started_event_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x72, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x13,
	0x41, 0x63, 0x63, 0x72, 0x75, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x2b, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x75,
	0x69, 0x64, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x79, 0x73, 0x6f, 0x67, 0x6f, 0x74, 0x61, 0x30, 0x33, 0x39, 0x39, 0x2f, 0x67, 0x6f, 0x70, 0x68,
	0x65, 0x72, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_events_accrual_started_event_proto_rawDescOnce sync.Once
	file_proto_events_accrual_started_event_proto_rawDescData []byte
)

func file_proto_events_accrual_started_event_proto_rawDescGZIP() []byte {
	file_proto_events_accrual_started_event_proto_rawDescOnce.Do(func() {
		file_proto_events_accrual_started_event_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_events_accrual_started_event_proto_rawDesc), len(file_proto_events_accrual_started_event_proto_rawDesc)))
	})
	return file_proto_events_accrual_started_event_proto_rawDescData
}

var file_proto_events_accrual_started_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_events_accrual_started_event_proto_goTypes = []any{
	(*AccrualStartedEvent)(nil), // 0: events.AccrualStartedEvent
	(*common.Uuid)(nil),         // 1: common.Uuid
}
var file_proto_events_accrual_started_event_proto_depIdxs = []int32{
	1, // 0: events.AccrualStartedEvent.event_uuid:type_name -> common.Uuid
	1, // 1: events.AccrualStartedEvent.order_uuid:type_name -> common.Uuid
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_events_accrual_started_event_proto_init() }
func file_proto_events_accrual_started_event_proto_init() {
	if File_proto_events_accrual_started_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_events_accrual_started_event_proto_rawDesc), len(file_proto_events_accrual_started_event_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_events_accrual_started_event_proto_goTypes,
		DependencyIndexes: file_proto_events_accrual_started_event_proto_depIdxs,
		MessageInfos:      file_proto_events_accrual_started_event_proto_msgTypes,
	}.Build()
	File_proto_events_accrual_started_event_proto = out.File
	file_proto_events_accrual_started_event_proto_goTypes = nil
	file_proto_events_accrual_started_event_proto_depIdxs = nil
}
