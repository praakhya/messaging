// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: msg/msg.proto

package msg

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

type ID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *int32                 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ID) Reset() {
	*x = ID{}
	mi := &file_msg_msg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

type Msg struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *int32                 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Topic         *int32                 `protobuf:"varint,2,req,name=topic" json:"topic,omitempty"`
	Content       *string                `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Msg) Reset() {
	*x = Msg{}
	mi := &file_msg_msg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Msg) GetTopic() int32 {
	if x != nil && x.Topic != nil {
		return *x.Topic
	}
	return 0
}

func (x *Msg) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enabled       *bool                  `protobuf:"varint,1,req,name=enabled" json:"enabled,omitempty"`
	Filter        *string                `protobuf:"bytes,2,req,name=filter" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_msg_msg_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *Filter) GetFilter() string {
	if x != nil && x.Filter != nil {
		return *x.Filter
	}
	return ""
}

type Topic struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *int32                 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Topic) Reset() {
	*x = Topic{}
	mi := &file_msg_msg_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{3}
}

func (x *Topic) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Topic) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_msg_msg_proto protoreflect.FileDescriptor

var file_msg_msg_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x3a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x2b, 0x0a,
	0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa8, 0x02, 0x0a, 0x09, 0x4d,
	0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x0b, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x73, 0x12, 0x0a, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x08, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x4d, 0x73, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x21, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x07, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x07, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x44, 0x1a, 0x08, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x41, 0x64, 0x64,
	0x4d, 0x73, 0x67, 0x12, 0x08, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x1a, 0x08, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x00, 0x12,
	0x24, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0a, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x4d, 0x73, 0x67, 0x12,
	0x08, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x1a, 0x08, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x4d, 0x73, 0x67, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6d, 0x73, 0x67,
})

var (
	file_msg_msg_proto_rawDescOnce sync.Once
	file_msg_msg_proto_rawDescData []byte
)

func file_msg_msg_proto_rawDescGZIP() []byte {
	file_msg_msg_proto_rawDescOnce.Do(func() {
		file_msg_msg_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_msg_msg_proto_rawDesc), len(file_msg_msg_proto_rawDesc)))
	})
	return file_msg_msg_proto_rawDescData
}

var file_msg_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_msg_msg_proto_goTypes = []any{
	(*ID)(nil),     // 0: msg.ID
	(*Msg)(nil),    // 1: msg.Msg
	(*Filter)(nil), // 2: msg.Filter
	(*Topic)(nil),  // 3: msg.Topic
}
var file_msg_msg_proto_depIdxs = []int32{
	2, // 0: msg.Messenger.GetTopics:input_type -> msg.Filter
	3, // 1: msg.Messenger.GetMsgs:input_type -> msg.Topic
	0, // 2: msg.Messenger.GetTopic:input_type -> msg.ID
	0, // 3: msg.Messenger.GetMsg:input_type -> msg.ID
	1, // 4: msg.Messenger.AddMsg:input_type -> msg.Msg
	3, // 5: msg.Messenger.AddTopic:input_type -> msg.Topic
	3, // 6: msg.Messenger.DelTopic:input_type -> msg.Topic
	1, // 7: msg.Messenger.DelMsg:input_type -> msg.Msg
	3, // 8: msg.Messenger.GetTopics:output_type -> msg.Topic
	1, // 9: msg.Messenger.GetMsgs:output_type -> msg.Msg
	3, // 10: msg.Messenger.GetTopic:output_type -> msg.Topic
	1, // 11: msg.Messenger.GetMsg:output_type -> msg.Msg
	1, // 12: msg.Messenger.AddMsg:output_type -> msg.Msg
	3, // 13: msg.Messenger.AddTopic:output_type -> msg.Topic
	3, // 14: msg.Messenger.DelTopic:output_type -> msg.Topic
	1, // 15: msg.Messenger.DelMsg:output_type -> msg.Msg
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_msg_proto_init() }
func file_msg_msg_proto_init() {
	if File_msg_msg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_msg_msg_proto_rawDesc), len(file_msg_msg_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_msg_proto_goTypes,
		DependencyIndexes: file_msg_msg_proto_depIdxs,
		MessageInfos:      file_msg_msg_proto_msgTypes,
	}.Build()
	File_msg_msg_proto = out.File
	file_msg_msg_proto_goTypes = nil
	file_msg_msg_proto_depIdxs = nil
}
