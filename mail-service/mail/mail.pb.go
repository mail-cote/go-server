// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: mail/mail.proto

package mail

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

type FetchQuizFromBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level string `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *FetchQuizFromBucketRequest) Reset() {
	*x = FetchQuizFromBucketRequest{}
	mi := &file_mail_mail_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchQuizFromBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchQuizFromBucketRequest) ProtoMessage() {}

func (x *FetchQuizFromBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchQuizFromBucketRequest.ProtoReflect.Descriptor instead.
func (*FetchQuizFromBucketRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{0}
}

func (x *FetchQuizFromBucketRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

type FetchQuizFromBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizId      int64  `protobuf:"varint,1,opt,name=quizId,proto3" json:"quizId,omitempty"`
	QuizContent string `protobuf:"bytes,2,opt,name=quizContent,proto3" json:"quizContent,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FetchQuizFromBucketResponse) Reset() {
	*x = FetchQuizFromBucketResponse{}
	mi := &file_mail_mail_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchQuizFromBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchQuizFromBucketResponse) ProtoMessage() {}

func (x *FetchQuizFromBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchQuizFromBucketResponse.ProtoReflect.Descriptor instead.
func (*FetchQuizFromBucketResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{1}
}

func (x *FetchQuizFromBucketResponse) GetQuizId() int64 {
	if x != nil {
		return x.QuizId
	}
	return 0
}

func (x *FetchQuizFromBucketResponse) GetQuizContent() string {
	if x != nil {
		return x.QuizContent
	}
	return ""
}

func (x *FetchQuizFromBucketResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendTo      string `protobuf:"bytes,1,opt,name=sendTo,proto3" json:"sendTo,omitempty"`
	SendFrom    string `protobuf:"bytes,2,opt,name=sendFrom,proto3" json:"sendFrom,omitempty"`
	QuizContent string `protobuf:"bytes,3,opt,name=quizContent,proto3" json:"quizContent,omitempty"`
}

func (x *SendMailRequest) Reset() {
	*x = SendMailRequest{}
	mi := &file_mail_mail_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailRequest) ProtoMessage() {}

func (x *SendMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailRequest.ProtoReflect.Descriptor instead.
func (*SendMailRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{2}
}

func (x *SendMailRequest) GetSendTo() string {
	if x != nil {
		return x.SendTo
	}
	return ""
}

func (x *SendMailRequest) GetSendFrom() string {
	if x != nil {
		return x.SendFrom
	}
	return ""
}

func (x *SendMailRequest) GetQuizContent() string {
	if x != nil {
		return x.QuizContent
	}
	return ""
}

type SendMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMailResponse) Reset() {
	*x = SendMailResponse{}
	mi := &file_mail_mail_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailResponse) ProtoMessage() {}

func (x *SendMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailResponse.ProtoReflect.Descriptor instead.
func (*SendMailResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{3}
}

func (x *SendMailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_mail_mail_proto protoreflect.FileDescriptor

var file_mail_mail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x32, 0x0a, 0x1a, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x51, 0x75, 0x69, 0x7a, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x71,
	0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x51, 0x75, 0x69, 0x7a, 0x46, 0x72, 0x6f, 0x6d, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x71, 0x75, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x71,
	0x75, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x69, 0x7a, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x69, 0x7a,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x67, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x6e, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x69, 0x7a,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71,
	0x75, 0x69, 0x7a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa9, 0x01, 0x0a, 0x04, 0x4d, 0x61, 0x69,
	0x6c, 0x12, 0x60, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x51, 0x75, 0x69, 0x7a, 0x46, 0x72,
	0x6f, 0x6d, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x51, 0x75, 0x69, 0x7a, 0x46, 0x72, 0x6f, 0x6d,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x51, 0x75, 0x69,
	0x7a, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12,
	0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x63, 0x6f, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mail_mail_proto_rawDescOnce sync.Once
	file_mail_mail_proto_rawDescData = file_mail_mail_proto_rawDesc
)

func file_mail_mail_proto_rawDescGZIP() []byte {
	file_mail_mail_proto_rawDescOnce.Do(func() {
		file_mail_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mail_mail_proto_rawDescData)
	})
	return file_mail_mail_proto_rawDescData
}

var file_mail_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mail_mail_proto_goTypes = []any{
	(*FetchQuizFromBucketRequest)(nil),  // 0: mail.v1.FetchQuizFromBucketRequest
	(*FetchQuizFromBucketResponse)(nil), // 1: mail.v1.FetchQuizFromBucketResponse
	(*SendMailRequest)(nil),             // 2: mail.v1.SendMailRequest
	(*SendMailResponse)(nil),            // 3: mail.v1.SendMailResponse
}
var file_mail_mail_proto_depIdxs = []int32{
	0, // 0: mail.v1.Mail.FetchQuizFromBucket:input_type -> mail.v1.FetchQuizFromBucketRequest
	2, // 1: mail.v1.Mail.SendMail:input_type -> mail.v1.SendMailRequest
	1, // 2: mail.v1.Mail.FetchQuizFromBucket:output_type -> mail.v1.FetchQuizFromBucketResponse
	3, // 3: mail.v1.Mail.SendMail:output_type -> mail.v1.SendMailResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mail_mail_proto_init() }
func file_mail_mail_proto_init() {
	if File_mail_mail_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mail_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mail_mail_proto_goTypes,
		DependencyIndexes: file_mail_mail_proto_depIdxs,
		MessageInfos:      file_mail_mail_proto_msgTypes,
	}.Build()
	File_mail_mail_proto = out.File
	file_mail_mail_proto_rawDesc = nil
	file_mail_mail_proto_goTypes = nil
	file_mail_mail_proto_depIdxs = nil
}
