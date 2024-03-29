// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: timeservice.proto

package protobuff

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

type Time struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Time) Reset() {
	*x = Time{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Time) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Time) ProtoMessage() {}

func (x *Time) ProtoReflect() protoreflect.Message {
	mi := &file_timeservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Time.ProtoReflect.Descriptor instead.
func (*Time) Descriptor() ([]byte, []int) {
	return file_timeservice_proto_rawDescGZIP(), []int{0}
}

func (x *Time) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TimeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *Time `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *TimeUpdate) Reset() {
	*x = TimeUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeUpdate) ProtoMessage() {}

func (x *TimeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_timeservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeUpdate.ProtoReflect.Descriptor instead.
func (*TimeUpdate) Descriptor() ([]byte, []int) {
	return file_timeservice_proto_rawDescGZIP(), []int{1}
}

func (x *TimeUpdate) GetTime() *Time {
	if x != nil {
		return x.Time
	}
	return nil
}

type NowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NowRequest) Reset() {
	*x = NowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowRequest) ProtoMessage() {}

func (x *NowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_timeservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowRequest.ProtoReflect.Descriptor instead.
func (*NowRequest) Descriptor() ([]byte, []int) {
	return file_timeservice_proto_rawDescGZIP(), []int{2}
}

type TimeStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *TimeStreamRequest) Reset() {
	*x = TimeStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timeservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeStreamRequest) ProtoMessage() {}

func (x *TimeStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_timeservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeStreamRequest.ProtoReflect.Descriptor instead.
func (*TimeStreamRequest) Descriptor() ([]byte, []int) {
	return file_timeservice_proto_rawDescGZIP(), []int{3}
}

func (x *TimeStreamRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

var File_timeservice_proto protoreflect.FileDescriptor

var file_timeservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x27, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x4e, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x32, 0x5b, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x4e, 0x6f, 0x77, 0x12, 0x0b, 0x2e, 0x4e, 0x6f,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x12, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timeservice_proto_rawDescOnce sync.Once
	file_timeservice_proto_rawDescData = file_timeservice_proto_rawDesc
)

func file_timeservice_proto_rawDescGZIP() []byte {
	file_timeservice_proto_rawDescOnce.Do(func() {
		file_timeservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_timeservice_proto_rawDescData)
	})
	return file_timeservice_proto_rawDescData
}

var file_timeservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_timeservice_proto_goTypes = []interface{}{
	(*Time)(nil),              // 0: Time
	(*TimeUpdate)(nil),        // 1: TimeUpdate
	(*NowRequest)(nil),        // 2: NowRequest
	(*TimeStreamRequest)(nil), // 3: TimeStreamRequest
}
var file_timeservice_proto_depIdxs = []int32{
	0, // 0: TimeUpdate.time:type_name -> Time
	2, // 1: TimeService.Now:input_type -> NowRequest
	3, // 2: TimeService.Stream:input_type -> TimeStreamRequest
	1, // 3: TimeService.Now:output_type -> TimeUpdate
	1, // 4: TimeService.Stream:output_type -> TimeUpdate
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_timeservice_proto_init() }
func file_timeservice_proto_init() {
	if File_timeservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timeservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Time); i {
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
		file_timeservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeUpdate); i {
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
		file_timeservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NowRequest); i {
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
		file_timeservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeStreamRequest); i {
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
			RawDescriptor: file_timeservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_timeservice_proto_goTypes,
		DependencyIndexes: file_timeservice_proto_depIdxs,
		MessageInfos:      file_timeservice_proto_msgTypes,
	}.Build()
	File_timeservice_proto = out.File
	file_timeservice_proto_rawDesc = nil
	file_timeservice_proto_goTypes = nil
	file_timeservice_proto_depIdxs = nil
}
