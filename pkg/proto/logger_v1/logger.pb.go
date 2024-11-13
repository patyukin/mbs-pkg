// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: logger.proto

package logger_v1

import (
	error_v1 "github.com/patyukin/mbs-pkg/pkg/proto/error_v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime   string `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime     string `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	ServiceName string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *LogReportRequest) Reset() {
	*x = LogReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogReportRequest) ProtoMessage() {}

func (x *LogReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogReportRequest.ProtoReflect.Descriptor instead.
func (*LogReportRequest) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{0}
}

func (x *LogReportRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *LogReportRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *LogReportRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type LogReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *emptypb.Empty          `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  *error_v1.ErrorResponse `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LogReportResponse) Reset() {
	*x = LogReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogReportResponse) ProtoMessage() {}

func (x *LogReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogReportResponse.ProtoReflect.Descriptor instead.
func (*LogReportResponse) Descriptor() ([]byte, []int) {
	return file_logger_proto_rawDescGZIP(), []int{1}
}

func (x *LogReportResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *LogReportResponse) GetError() *error_v1.ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_logger_proto protoreflect.FileDescriptor

var file_logger_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x1a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x5a, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x74, 0x79, 0x75, 0x6b, 0x69, 0x6e, 0x2f, 0x6d, 0x62, 0x73, 0x2d,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logger_proto_rawDescOnce sync.Once
	file_logger_proto_rawDescData = file_logger_proto_rawDesc
)

func file_logger_proto_rawDescGZIP() []byte {
	file_logger_proto_rawDescOnce.Do(func() {
		file_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_logger_proto_rawDescData)
	})
	return file_logger_proto_rawDescData
}

var file_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logger_proto_goTypes = []any{
	(*LogReportRequest)(nil),       // 0: logger_v1.LogReportRequest
	(*LogReportResponse)(nil),      // 1: logger_v1.LogReportResponse
	(*emptypb.Empty)(nil),          // 2: google.protobuf.Empty
	(*error_v1.ErrorResponse)(nil), // 3: error_v1.ErrorResponse
}
var file_logger_proto_depIdxs = []int32{
	2, // 0: logger_v1.LogReportResponse.result:type_name -> google.protobuf.Empty
	3, // 1: logger_v1.LogReportResponse.error:type_name -> error_v1.ErrorResponse
	0, // 2: logger_v1.LoggerService.GetLogReport:input_type -> logger_v1.LogReportRequest
	1, // 3: logger_v1.LoggerService.GetLogReport:output_type -> logger_v1.LogReportResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_logger_proto_init() }
func file_logger_proto_init() {
	if File_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logger_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LogReportRequest); i {
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
		file_logger_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LogReportResponse); i {
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
			RawDescriptor: file_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logger_proto_goTypes,
		DependencyIndexes: file_logger_proto_depIdxs,
		MessageInfos:      file_logger_proto_msgTypes,
	}.Build()
	File_logger_proto = out.File
	file_logger_proto_rawDesc = nil
	file_logger_proto_goTypes = nil
	file_logger_proto_depIdxs = nil
}
