// Code generated by protoc-gen-go.
// source: task_execution.proto
// DO NOT EDIT!

/*
Package ga4gh_task_exec is a generated protocol buffer package.

It is generated from these files:
	task_execution.proto

It has these top-level messages:
	FileParameter
	DockerExecutor
	Disk
	Resources
	Task
	TaskArgs
	TaskRunRequest
	TaskId
	TaskOpId
	TaskOp
*/
package ga4gh_task_exec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// / Mapping long term file storage to path used in docker container
type FileParameter struct {
	Path        string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	StoragePath string `protobuf:"bytes,2,opt,name=storagePath" json:"storagePath,omitempty"`
}

func (m *FileParameter) Reset()                    { *m = FileParameter{} }
func (m *FileParameter) String() string            { return proto.CompactTextString(m) }
func (*FileParameter) ProtoMessage()               {}
func (*FileParameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// / A command line to be executed and the docker container to run it
type DockerExecutor struct {
	ImageName string `protobuf:"bytes,1,opt,name=imageName" json:"imageName,omitempty"`
	Cmd       string `protobuf:"bytes,2,opt,name=cmd" json:"cmd,omitempty"`
}

func (m *DockerExecutor) Reset()                    { *m = DockerExecutor{} }
func (m *DockerExecutor) String() string            { return proto.CompactTextString(m) }
func (*DockerExecutor) ProtoMessage()               {}
func (*DockerExecutor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// / Attached disk request.
type Disk struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SizeGb uint32 `protobuf:"varint,2,opt,name=sizeGb" json:"sizeGb,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	// / could identify. Leave blank if is to be a newly created disk
	AutoDelete bool   `protobuf:"varint,4,opt,name=autoDelete" json:"autoDelete,omitempty"`
	ReadOnly   bool   `protobuf:"varint,5,opt,name=readOnly" json:"readOnly,omitempty"`
	MountPoint string `protobuf:"bytes,6,opt,name=mountPoint" json:"mountPoint,omitempty"`
}

func (m *Disk) Reset()                    { *m = Disk{} }
func (m *Disk) String() string            { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()               {}
func (*Disk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Resources struct {
	MinimumCpuCores uint32   `protobuf:"varint,1,opt,name=minimumCpuCores" json:"minimumCpuCores,omitempty"`
	Preemptible     bool     `protobuf:"varint,2,opt,name=preemptible" json:"preemptible,omitempty"`
	MinimumRamGb    uint32   `protobuf:"varint,3,opt,name=minimumRamGb" json:"minimumRamGb,omitempty"`
	Disks           []*Disk  `protobuf:"bytes,4,rep,name=disks" json:"disks,omitempty"`
	Zones           []string `protobuf:"bytes,5,rep,name=zones" json:"zones,omitempty"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Resources) GetDisks() []*Disk {
	if m != nil {
		return m.Disks
	}
	return nil
}

// / The description of a task to be run
type Task struct {
	Name             string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ProjectId        string            `protobuf:"bytes,2,opt,name=projectId" json:"projectId,omitempty"`
	Description      string            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	InputParameters  []*FileParameter  `protobuf:"bytes,4,rep,name=inputParameters" json:"inputParameters,omitempty"`
	OutputParameters []*FileParameter  `protobuf:"bytes,5,rep,name=outputParameters" json:"outputParameters,omitempty"`
	Resources        *Resources        `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
	TaskId           string            `protobuf:"bytes,7,opt,name=taskId" json:"taskId,omitempty"`
	Docker           []*DockerExecutor `protobuf:"bytes,8,rep,name=docker" json:"docker,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Task) GetInputParameters() []*FileParameter {
	if m != nil {
		return m.InputParameters
	}
	return nil
}

func (m *Task) GetOutputParameters() []*FileParameter {
	if m != nil {
		return m.OutputParameters
	}
	return nil
}

func (m *Task) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Task) GetDocker() []*DockerExecutor {
	if m != nil {
		return m.Docker
	}
	return nil
}

// / Arguments for task to be instanced
type TaskArgs struct {
	ProjectId string                    `protobuf:"bytes,1,opt,name=projectId" json:"projectId,omitempty"`
	Inputs    map[string]*FileParameter `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Outputs   map[string]*FileParameter `protobuf:"bytes,3,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Resources *Resources                `protobuf:"bytes,4,opt,name=resources" json:"resources,omitempty"`
}

func (m *TaskArgs) Reset()                    { *m = TaskArgs{} }
func (m *TaskArgs) String() string            { return proto.CompactTextString(m) }
func (*TaskArgs) ProtoMessage()               {}
func (*TaskArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TaskArgs) GetInputs() map[string]*FileParameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TaskArgs) GetOutputs() map[string]*FileParameter {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TaskArgs) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

// / Task run request
type TaskRunRequest struct {
	TaskArgs *TaskArgs `protobuf:"bytes,1,opt,name=taskArgs" json:"taskArgs,omitempty"`
	// Define either a taskId or an ephemeralTask
	TaskId        string `protobuf:"bytes,2,opt,name=taskId" json:"taskId,omitempty"`
	EphemeralTask *Task  `protobuf:"bytes,3,opt,name=ephemeralTask" json:"ephemeralTask,omitempty"`
}

func (m *TaskRunRequest) Reset()                    { *m = TaskRunRequest{} }
func (m *TaskRunRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskRunRequest) ProtoMessage()               {}
func (*TaskRunRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TaskRunRequest) GetTaskArgs() *TaskArgs {
	if m != nil {
		return m.TaskArgs
	}
	return nil
}

func (m *TaskRunRequest) GetEphemeralTask() *Task {
	if m != nil {
		return m.EphemeralTask
	}
	return nil
}

// / ID of a Task description
type TaskId struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *TaskId) Reset()                    { *m = TaskId{} }
func (m *TaskId) String() string            { return proto.CompactTextString(m) }
func (*TaskId) ProtoMessage()               {}
func (*TaskId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// / ID of an instance of a Task
type TaskOpId struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *TaskOpId) Reset()                    { *m = TaskOpId{} }
func (m *TaskOpId) String() string            { return proto.CompactTextString(m) }
func (*TaskOpId) ProtoMessage()               {}
func (*TaskOpId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// / The description of the running instance of a task
type TaskOp struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Done     bool              `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
}

func (m *TaskOp) Reset()                    { *m = TaskOp{} }
func (m *TaskOp) String() string            { return proto.CompactTextString(m) }
func (*TaskOp) ProtoMessage()               {}
func (*TaskOp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TaskOp) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*FileParameter)(nil), "ga4gh_task_exec.FileParameter")
	proto.RegisterType((*DockerExecutor)(nil), "ga4gh_task_exec.DockerExecutor")
	proto.RegisterType((*Disk)(nil), "ga4gh_task_exec.Disk")
	proto.RegisterType((*Resources)(nil), "ga4gh_task_exec.Resources")
	proto.RegisterType((*Task)(nil), "ga4gh_task_exec.Task")
	proto.RegisterType((*TaskArgs)(nil), "ga4gh_task_exec.TaskArgs")
	proto.RegisterType((*TaskRunRequest)(nil), "ga4gh_task_exec.TaskRunRequest")
	proto.RegisterType((*TaskId)(nil), "ga4gh_task_exec.TaskId")
	proto.RegisterType((*TaskOpId)(nil), "ga4gh_task_exec.TaskOpId")
	proto.RegisterType((*TaskOp)(nil), "ga4gh_task_exec.TaskOp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for TaskService service

type TaskServiceClient interface {
	// / Create a task
	CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	// / Delete a task
	DeleteTask(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*TaskId, error)
	// / Get a task by its ID
	GetTask(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*Task, error)
	// / Run a task
	RunTask(ctx context.Context, in *TaskRunRequest, opts ...grpc.CallOption) (*TaskOpId, error)
	// / Get info about a running task
	GetTaskOp(ctx context.Context, in *TaskOpId, opts ...grpc.CallOption) (*TaskOp, error)
	// / Cancel a running task
	CancelTaskOp(ctx context.Context, in *TaskOpId, opts ...grpc.CallOption) (*TaskOpId, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/ga4gh_task_exec.TaskService/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteTask(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*TaskId, error) {
	out := new(TaskId)
	err := grpc.Invoke(ctx, "/ga4gh_task_exec.TaskService/DeleteTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *TaskId, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/ga4gh_task_exec.TaskService/GetTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) RunTask(ctx context.Context, in *TaskRunRequest, opts ...grpc.CallOption) (*TaskOpId, error) {
	out := new(TaskOpId)
	err := grpc.Invoke(ctx, "/ga4gh_task_exec.TaskService/RunTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTaskOp(ctx context.Context, in *TaskOpId, opts ...grpc.CallOption) (*TaskOp, error) {
	out := new(TaskOp)
	err := grpc.Invoke(ctx, "/ga4gh_task_exec.TaskService/GetTaskOp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) CancelTaskOp(ctx context.Context, in *TaskOpId, opts ...grpc.CallOption) (*TaskOpId, error) {
	out := new(TaskOpId)
	err := grpc.Invoke(ctx, "/ga4gh_task_exec.TaskService/CancelTaskOp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceServer interface {
	// / Create a task
	CreateTask(context.Context, *Task) (*Task, error)
	// / Delete a task
	DeleteTask(context.Context, *TaskId) (*TaskId, error)
	// / Get a task by its ID
	GetTask(context.Context, *TaskId) (*Task, error)
	// / Run a task
	RunTask(context.Context, *TaskRunRequest) (*TaskOpId, error)
	// / Get info about a running task
	GetTaskOp(context.Context, *TaskOpId) (*TaskOp, error)
	// / Cancel a running task
	CancelTaskOp(context.Context, *TaskOpId) (*TaskOpId, error)
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_exec.TaskService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_exec.TaskService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteTask(ctx, req.(*TaskId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_exec.TaskService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*TaskId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_RunTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).RunTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_exec.TaskService/RunTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).RunTask(ctx, req.(*TaskRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTaskOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskOpId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTaskOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_exec.TaskService/GetTaskOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTaskOp(ctx, req.(*TaskOpId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_CancelTaskOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskOpId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CancelTaskOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_exec.TaskService/CancelTaskOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CancelTaskOp(ctx, req.(*TaskOpId))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ga4gh_task_exec.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TaskService_DeleteTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "RunTask",
			Handler:    _TaskService_RunTask_Handler,
		},
		{
			MethodName: "GetTaskOp",
			Handler:    _TaskService_GetTaskOp_Handler,
		},
		{
			MethodName: "CancelTaskOp",
			Handler:    _TaskService_CancelTaskOp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x4f, 0xe3, 0x46,
	0x14, 0x56, 0x7e, 0xdb, 0x2f, 0x71, 0x80, 0x21, 0x15, 0x89, 0x55, 0x95, 0xca, 0x6a, 0x2b, 0x44,
	0x4b, 0xd2, 0xa6, 0xad, 0x40, 0xdc, 0xaa, 0x40, 0x11, 0x07, 0x08, 0x4a, 0xb9, 0x54, 0x3d, 0xd0,
	0xc1, 0x1e, 0x05, 0x6f, 0x6c, 0x8f, 0x77, 0x3c, 0x46, 0x0b, 0x68, 0x2f, 0xab, 0x3d, 0xef, 0x65,
	0xff, 0x80, 0xfd, 0x17, 0xf6, 0x7f, 0xd9, 0xe3, 0x5e, 0xf7, 0x0f, 0xd9, 0x99, 0xb1, 0x93, 0x60,
	0x12, 0xc3, 0x4a, 0x7b, 0xcc, 0xcc, 0xf7, 0xbe, 0xf7, 0x7d, 0xdf, 0x7b, 0xe3, 0x40, 0x8b, 0xe3,
	0x68, 0x72, 0x41, 0x5e, 0x10, 0x3b, 0xe6, 0x2e, 0x0d, 0xba, 0x21, 0xa3, 0x9c, 0xa2, 0x95, 0x31,
	0xfe, 0x63, 0x7c, 0x75, 0x31, 0xbb, 0x33, 0xbf, 0x1d, 0x53, 0x3a, 0xf6, 0x48, 0x0f, 0x87, 0x6e,
	0x0f, 0x07, 0x01, 0xe5, 0x58, 0xa2, 0xa3, 0x04, 0x6e, 0xf5, 0xc1, 0xf8, 0xdb, 0xf5, 0xc8, 0x19,
	0x66, 0xd8, 0x27, 0x9c, 0x30, 0xd4, 0x80, 0x72, 0x88, 0xf9, 0x55, 0xbb, 0xf0, 0x7d, 0x61, 0x4b,
	0x47, 0xeb, 0x50, 0x8f, 0x38, 0x65, 0x78, 0x2c, 0x10, 0xe2, 0xb0, 0x28, 0x0f, 0xad, 0x5f, 0xa1,
	0x79, 0x40, 0xed, 0x09, 0x61, 0x87, 0xaa, 0x37, 0x65, 0x68, 0x0d, 0x74, 0xd7, 0x17, 0xa0, 0x53,
	0x41, 0x92, 0x56, 0xd6, 0xa1, 0x64, 0xfb, 0x4e, 0x5a, 0x11, 0x40, 0xf9, 0xc0, 0x8d, 0x26, 0x92,
	0x3c, 0x98, 0x43, 0x9a, 0x50, 0x8d, 0xdc, 0x5b, 0x72, 0x74, 0xa9, 0x50, 0x86, 0xfa, 0x4d, 0x63,
	0x66, 0x93, 0x76, 0x49, 0xdd, 0x23, 0x00, 0x2c, 0xe8, 0x0f, 0x88, 0x27, 0x94, 0xb5, 0xcb, 0xe2,
	0x4c, 0x43, 0xab, 0xa0, 0x31, 0x82, 0x9d, 0x61, 0xe0, 0xdd, 0xb4, 0x2b, 0xea, 0x44, 0xa0, 0x7c,
	0x1a, 0x07, 0xfc, 0x8c, 0xba, 0x01, 0x6f, 0x57, 0x55, 0xbf, 0xd7, 0x05, 0xd0, 0x47, 0x24, 0x21,
	0x8b, 0xd0, 0x06, 0xac, 0xf8, 0x6e, 0xe0, 0xfa, 0xb1, 0x3f, 0x08, 0xe3, 0x01, 0x65, 0x24, 0x52,
	0x02, 0x0c, 0xe9, 0x2e, 0x64, 0x84, 0xf8, 0x21, 0x77, 0x2f, 0x3d, 0xa2, 0x54, 0x68, 0xa8, 0x05,
	0x8d, 0x14, 0x3d, 0xc2, 0xbe, 0xd0, 0x56, 0x52, 0xd0, 0x1f, 0xa0, 0xe2, 0x08, 0x07, 0x91, 0x90,
	0x51, 0xda, 0xaa, 0xf7, 0xbf, 0xe9, 0x3e, 0x88, 0xb9, 0xab, 0xfc, 0x19, 0x50, 0xb9, 0xa5, 0x81,
	0xe0, 0xaf, 0x08, 0x94, 0x6e, 0xbd, 0x2b, 0x42, 0xf9, 0x1c, 0x2f, 0xf8, 0x16, 0x69, 0x89, 0xf0,
	0x9f, 0x11, 0x9b, 0x1f, 0xa7, 0x01, 0x49, 0x25, 0x0e, 0x89, 0x6c, 0xe6, 0x86, 0x72, 0x38, 0xa9,
	0xff, 0x5d, 0x58, 0x71, 0x83, 0x30, 0xe6, 0xb3, 0xe1, 0x4c, 0xbb, 0x7f, 0xb7, 0xd0, 0x3d, 0x3b,
	0xc3, 0x3d, 0x58, 0xa5, 0x31, 0xcf, 0x56, 0x56, 0xbe, 0xa8, 0x72, 0x07, 0x74, 0x36, 0xcd, 0x4d,
	0x65, 0x59, 0xef, 0x9b, 0x0b, 0x25, 0xf3, 0x64, 0xc5, 0xc4, 0xe4, 0xb1, 0xb0, 0x51, 0x53, 0x8a,
	0x7b, 0x50, 0x75, 0xd4, 0x66, 0xb4, 0x35, 0xd5, 0x6e, 0x73, 0x31, 0xa6, 0xcc, 0xe2, 0x58, 0x1f,
	0x8b, 0xa0, 0xc9, 0x84, 0xfe, 0x62, 0xe3, 0x28, 0x9b, 0x4b, 0x12, 0xd5, 0x9f, 0x50, 0x55, 0x11,
	0x44, 0x22, 0x27, 0x49, 0xf8, 0xe3, 0x02, 0xe1, 0xb4, 0xba, 0x7b, 0xac, 0x70, 0x87, 0x01, 0x67,
	0x37, 0x22, 0xb9, 0x5a, 0x12, 0x40, 0x24, 0xa2, 0x94, 0x75, 0x3f, 0xe5, 0xd7, 0x0d, 0x13, 0x60,
	0x52, 0x98, 0xf1, 0x5f, 0x7e, 0xca, 0xbf, 0x79, 0x02, 0xf5, 0xfb, 0x6d, 0xc5, 0xce, 0x4f, 0xc8,
	0x4d, 0x2a, 0x7d, 0x07, 0x2a, 0xd7, 0xd8, 0x8b, 0x93, 0xb5, 0x7a, 0x32, 0xf9, 0xfd, 0xe2, 0x5e,
	0xc1, 0x3c, 0x85, 0x46, 0x46, 0xcd, 0x57, 0xf2, 0x59, 0x77, 0xd0, 0x94, 0x36, 0x47, 0x71, 0x30,
	0x22, 0xcf, 0x63, 0x12, 0x71, 0xf4, 0x33, 0x68, 0x3c, 0x35, 0xae, 0x68, 0xeb, 0xfd, 0x4e, 0x6e,
	0x32, 0xf7, 0xa6, 0x9b, 0x2c, 0xe9, 0x2f, 0x60, 0x90, 0xf0, 0x8a, 0xf8, 0x84, 0x61, 0x4f, 0x82,
	0xd4, 0x9a, 0x2e, 0x7b, 0x0b, 0xf2, 0xd2, 0xda, 0x80, 0xea, 0xb9, 0xaa, 0x96, 0xaf, 0x22, 0x51,
	0xae, 0x8c, 0x58, 0x9d, 0x64, 0xe4, 0xc3, 0x70, 0xf1, 0xea, 0x4d, 0x21, 0x29, 0x1a, 0x86, 0x0f,
	0x9e, 0xcc, 0x2e, 0x68, 0xc2, 0x16, 0x76, 0x30, 0xc7, 0x8f, 0x6e, 0xc2, 0x30, 0xec, 0x9e, 0xa4,
	0xb8, 0x24, 0x42, 0x41, 0xe3, 0x88, 0x17, 0xa9, 0xa4, 0x6a, 0x66, 0x0f, 0x8c, 0xec, 0x75, 0x26,
	0x61, 0xe3, 0x7e, 0xc2, 0xba, 0x4c, 0xb0, 0xff, 0xbe, 0x0c, 0x75, 0xc9, 0xfb, 0x0f, 0x61, 0xd7,
	0xae, 0x4d, 0xd0, 0x19, 0xc0, 0x40, 0x7c, 0x7f, 0x38, 0x51, 0xcf, 0x7a, 0xb9, 0x73, 0x33, 0x27,
	0x90, 0xd6, 0xab, 0x0f, 0x9f, 0xde, 0x16, 0x9b, 0x96, 0xde, 0xbb, 0xfe, 0xad, 0x27, 0xef, 0xa2,
	0xfd, 0xc2, 0x36, 0xfa, 0x17, 0x20, 0xf9, 0xc0, 0x29, 0xc6, 0x8d, 0xa5, 0xa5, 0xc7, 0x8e, 0x99,
	0x77, 0x61, 0x75, 0x14, 0xeb, 0xfa, 0xf6, 0xda, 0x8c, 0xb5, 0x77, 0xa7, 0x5c, 0xbc, 0x44, 0xe7,
	0x50, 0x3b, 0x22, 0xfc, 0x71, 0xde, 0x1c, 0xad, 0x29, 0x2b, 0x5a, 0xc2, 0x7a, 0x01, 0x35, 0xb1,
	0x50, 0x8a, 0x75, 0x73, 0x69, 0xf1, 0x7c, 0xdd, 0xcc, 0x4e, 0xce, 0x90, 0x84, 0xee, 0xb6, 0xea,
	0x80, 0x2c, 0x63, 0x9e, 0x06, 0x8b, 0x03, 0x99, 0xc8, 0x7f, 0xa0, 0xa7, 0xb2, 0xc5, 0x1a, 0xe4,
	0x33, 0xe4, 0x44, 0x32, 0x0c, 0x2d, 0x53, 0x51, 0xb7, 0x10, 0x9a, 0x52, 0xd3, 0x70, 0xa6, 0xfe,
	0x7f, 0x68, 0x0c, 0x70, 0x60, 0x13, 0xef, 0x69, 0xfe, 0x47, 0xc4, 0xa7, 0x1d, 0xb6, 0x97, 0x74,
	0xb8, 0xac, 0xaa, 0x3f, 0xd6, 0xdf, 0x3f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1d, 0x03, 0xe7,
	0x9f, 0x07, 0x00, 0x00,
}
