// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: server.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Task_Type int32

const (
	Task_TEST     Task_Type = 0
	Task_HOMEWORK Task_Type = 1
	Task_QUIZ     Task_Type = 2
)

// Enum value maps for Task_Type.
var (
	Task_Type_name = map[int32]string{
		0: "TEST",
		1: "HOMEWORK",
		2: "QUIZ",
	}
	Task_Type_value = map[string]int32{
		"TEST":     0,
		"HOMEWORK": 1,
		"QUIZ":     2,
	}
)

func (x Task_Type) Enum() *Task_Type {
	p := new(Task_Type)
	*p = x
	return p
}

func (x Task_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_enumTypes[0].Descriptor()
}

func (Task_Type) Type() protoreflect.EnumType {
	return &file_server_proto_enumTypes[0]
}

func (x Task_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_Type.Descriptor instead.
func (Task_Type) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1, 0}
}

type Subscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Lab  []string `protobuf:"bytes,2,rep,name=lab,proto3" json:"lab,omitempty"`
}

func (x *Subscribe) Reset() {
	*x = Subscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscribe) ProtoMessage() {}

func (x *Subscribe) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscribe.ProtoReflect.Descriptor instead.
func (*Subscribe) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Subscribe) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Subscribe) GetLab() []string {
	if x != nil {
		return x.Lab
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         Task_Type `protobuf:"varint,1,opt,name=type,proto3,enum=v1.Task_Type" json:"type,omitempty"`
	Deadline     int64     `protobuf:"varint,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Descriptions []string  `protobuf:"bytes,3,rep,name=descriptions,proto3" json:"descriptions,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetType() Task_Type {
	if x != nil {
		return x.Type
	}
	return Task_TEST
}

func (x *Task) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *Task) GetDescriptions() []string {
	if x != nil {
		return x.Descriptions
	}
	return nil
}

type Solution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Comment  string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Solution [][]byte `protobuf:"bytes,3,rep,name=solution,proto3" json:"solution,omitempty"`
}

func (x *Solution) Reset() {
	*x = Solution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solution) ProtoMessage() {}

func (x *Solution) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solution.ProtoReflect.Descriptor instead.
func (*Solution) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *Solution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Solution) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Solution) GetSolution() [][]byte {
	if x != nil {
		return x.Solution
	}
	return nil
}

type GradeSolution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeGraded int64  `protobuf:"varint,1,opt,name=time_graded,json=timeGraded,proto3" json:"time_graded,omitempty"`
	Grade      int32  `protobuf:"varint,2,opt,name=grade,proto3" json:"grade,omitempty"`
	Feedback   string `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *GradeSolution) Reset() {
	*x = GradeSolution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradeSolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradeSolution) ProtoMessage() {}

func (x *GradeSolution) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradeSolution.ProtoReflect.Descriptor instead.
func (*GradeSolution) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *GradeSolution) GetTimeGraded() int64 {
	if x != nil {
		return x.TimeGraded
	}
	return 0
}

func (x *GradeSolution) GetGrade() int32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *GradeSolution) GetFeedback() string {
	if x != nil {
		return x.Feedback
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x22, 0x31, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x62, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x6c, 0x61, 0x62, 0x22, 0x93, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x28, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x53,
	0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x4f, 0x4d, 0x45, 0x57, 0x4f, 0x52, 0x4b, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x51, 0x55, 0x49, 0x5a, 0x10, 0x02, 0x22, 0x50, 0x0a, 0x08, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a,
	0x0d, 0x47, 0x72, 0x61, 0x64, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x9f, 0x01, 0x0a, 0x0e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x0a, 0x53, 0x69, 0x67, 0x6e, 0x46, 0x6f, 0x72, 0x4c, 0x61, 0x62, 0x12, 0x0d, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x1a, 0x08, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2d, 0x0a,
	0x14, 0x54, 0x45, 0x53, 0x54, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a,
	0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_server_proto_goTypes = []interface{}{
	(Task_Type)(0),        // 0: v1.Task.Type
	(*Subscribe)(nil),     // 1: v1.Subscribe
	(*Task)(nil),          // 2: v1.Task
	(*Solution)(nil),      // 3: v1.Solution
	(*GradeSolution)(nil), // 4: v1.GradeSolution
	(*Empty)(nil),         // 5: v1.Empty
}
var file_server_proto_depIdxs = []int32{
	0, // 0: v1.Task.type:type_name -> v1.Task.Type
	1, // 1: v1.StudentService.SignForLab:input_type -> v1.Subscribe
	3, // 2: v1.StudentService.SendSolution:input_type -> v1.Solution
	2, // 3: v1.StudentService.TESTCreateAssignment:input_type -> v1.Task
	2, // 4: v1.StudentService.SignForLab:output_type -> v1.Task
	4, // 5: v1.StudentService.SendSolution:output_type -> v1.GradeSolution
	5, // 6: v1.StudentService.TESTCreateAssignment:output_type -> v1.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscribe); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solution); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradeSolution); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		EnumInfos:         file_server_proto_enumTypes,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentServiceClient interface {
	SignForLab(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (StudentService_SignForLabClient, error)
	SendSolution(ctx context.Context, in *Solution, opts ...grpc.CallOption) (StudentService_SendSolutionClient, error)
	TESTCreateAssignment(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Empty, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) SignForLab(ctx context.Context, in *Subscribe, opts ...grpc.CallOption) (StudentService_SignForLabClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StudentService_serviceDesc.Streams[0], "/v1.StudentService/SignForLab", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentServiceSignForLabClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StudentService_SignForLabClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type studentServiceSignForLabClient struct {
	grpc.ClientStream
}

func (x *studentServiceSignForLabClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *studentServiceClient) SendSolution(ctx context.Context, in *Solution, opts ...grpc.CallOption) (StudentService_SendSolutionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StudentService_serviceDesc.Streams[1], "/v1.StudentService/SendSolution", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentServiceSendSolutionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StudentService_SendSolutionClient interface {
	Recv() (*GradeSolution, error)
	grpc.ClientStream
}

type studentServiceSendSolutionClient struct {
	grpc.ClientStream
}

func (x *studentServiceSendSolutionClient) Recv() (*GradeSolution, error) {
	m := new(GradeSolution)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *studentServiceClient) TESTCreateAssignment(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/v1.StudentService/TESTCreateAssignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
type StudentServiceServer interface {
	SignForLab(*Subscribe, StudentService_SignForLabServer) error
	SendSolution(*Solution, StudentService_SendSolutionServer) error
	TESTCreateAssignment(context.Context, *Task) (*Empty, error)
}

// UnimplementedStudentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (*UnimplementedStudentServiceServer) SignForLab(*Subscribe, StudentService_SignForLabServer) error {
	return status.Errorf(codes.Unimplemented, "method SignForLab not implemented")
}
func (*UnimplementedStudentServiceServer) SendSolution(*Solution, StudentService_SendSolutionServer) error {
	return status.Errorf(codes.Unimplemented, "method SendSolution not implemented")
}
func (*UnimplementedStudentServiceServer) TESTCreateAssignment(context.Context, *Task) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TESTCreateAssignment not implemented")
}

func RegisterStudentServiceServer(s *grpc.Server, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_SignForLab_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscribe)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServiceServer).SignForLab(m, &studentServiceSignForLabServer{stream})
}

type StudentService_SignForLabServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type studentServiceSignForLabServer struct {
	grpc.ServerStream
}

func (x *studentServiceSignForLabServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _StudentService_SendSolution_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Solution)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServiceServer).SendSolution(m, &studentServiceSendSolutionServer{stream})
}

type StudentService_SendSolutionServer interface {
	Send(*GradeSolution) error
	grpc.ServerStream
}

type studentServiceSendSolutionServer struct {
	grpc.ServerStream
}

func (x *studentServiceSendSolutionServer) Send(m *GradeSolution) error {
	return x.ServerStream.SendMsg(m)
}

func _StudentService_TESTCreateAssignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).TESTCreateAssignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.StudentService/TESTCreateAssignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).TESTCreateAssignment(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TESTCreateAssignment",
			Handler:    _StudentService_TESTCreateAssignment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SignForLab",
			Handler:       _StudentService_SignForLab_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendSolution",
			Handler:       _StudentService_SendSolution_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server.proto",
}
