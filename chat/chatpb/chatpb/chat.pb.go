// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: chatpb/chat.proto

package chatpb

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

type FromClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *FromClient) Reset() {
	*x = FromClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromClient) ProtoMessage() {}

func (x *FromClient) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromClient.ProtoReflect.Descriptor instead.
func (*FromClient) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *FromClient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FromClient) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type FromServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *FromServer) Reset() {
	*x = FromServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromServer) ProtoMessage() {}

func (x *FromServer) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromServer.ProtoReflect.Descriptor instead.
func (*FromServer) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *FromServer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FromServer) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_chatpb_chat_proto protoreflect.FileDescriptor

var file_chatpb_chat_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x22, 0x34, 0x0a, 0x0a, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x34, 0x0a, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x47, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x09, 0x5a, 0x07, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chatpb_chat_proto_rawDescOnce sync.Once
	file_chatpb_chat_proto_rawDescData = file_chatpb_chat_proto_rawDesc
)

func file_chatpb_chat_proto_rawDescGZIP() []byte {
	file_chatpb_chat_proto_rawDescOnce.Do(func() {
		file_chatpb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatpb_chat_proto_rawDescData)
	})
	return file_chatpb_chat_proto_rawDescData
}

var file_chatpb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chatpb_chat_proto_goTypes = []interface{}{
	(*FromClient)(nil), // 0: chatpb.FromClient
	(*FromServer)(nil), // 1: chatpb.FromServer
}
var file_chatpb_chat_proto_depIdxs = []int32{
	0, // 0: chatpb.MsgService.ChatService:input_type -> chatpb.FromClient
	1, // 1: chatpb.MsgService.ChatService:output_type -> chatpb.FromServer
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chatpb_chat_proto_init() }
func file_chatpb_chat_proto_init() {
	if File_chatpb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatpb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromClient); i {
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
		file_chatpb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromServer); i {
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
			RawDescriptor: file_chatpb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatpb_chat_proto_goTypes,
		DependencyIndexes: file_chatpb_chat_proto_depIdxs,
		MessageInfos:      file_chatpb_chat_proto_msgTypes,
	}.Build()
	File_chatpb_chat_proto = out.File
	file_chatpb_chat_proto_rawDesc = nil
	file_chatpb_chat_proto_goTypes = nil
	file_chatpb_chat_proto_depIdxs = nil
}
