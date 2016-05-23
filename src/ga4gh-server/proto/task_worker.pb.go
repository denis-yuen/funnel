// Code generated by protoc-gen-go.
// source: task_worker.proto
// DO NOT EDIT!

/*
Package ga4gh_task_ref is a generated protocol buffer package.

It is generated from these files:
	task_worker.proto

It has these top-level messages:
	WorkerInfo
	JobRequest
	JobResult
	UpdateStatusRequest
*/
package ga4gh_task_ref

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ga4gh_task_exec "ga4gh-tasks"

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

// *
// Worker Info
type WorkerInfo struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	LastPing int64  `protobuf:"varint,2,opt,name=last_ping" json:"last_ping,omitempty"`
	Hostname string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *WorkerInfo) Reset()                    { *m = WorkerInfo{} }
func (m *WorkerInfo) String() string            { return proto.CompactTextString(m) }
func (*WorkerInfo) ProtoMessage()               {}
func (*WorkerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type JobRequest struct {
	Worker    *WorkerInfo                `protobuf:"bytes,1,opt,name=worker" json:"worker,omitempty"`
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,2,opt,name=resources" json:"resources,omitempty"`
}

func (m *JobRequest) Reset()                    { *m = JobRequest{} }
func (m *JobRequest) String() string            { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()               {}
func (*JobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobRequest) GetWorker() *WorkerInfo {
	if m != nil {
		return m.Worker
	}
	return nil
}

func (m *JobRequest) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

type JobResult struct {
	Task *ga4gh_task_exec.TaskOp `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
}

func (m *JobResult) Reset()                    { *m = JobResult{} }
func (m *JobResult) String() string            { return proto.CompactTextString(m) }
func (*JobResult) ProtoMessage()               {}
func (*JobResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JobResult) GetTask() *ga4gh_task_exec.TaskOp {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateStatusRequest struct {
	Id       string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Step     int64                      `protobuf:"varint,2,opt,name=step" json:"step,omitempty"`
	State    ga4gh_task_exec.State      `protobuf:"varint,3,opt,name=state,enum=ga4gh_task_exec.State" json:"state,omitempty"`
	Log      *ga4gh_task_exec.TaskOpLog `protobuf:"bytes,4,opt,name=log" json:"log,omitempty"`
	WorkerId string                     `protobuf:"bytes,5,opt,name=worker_id" json:"worker_id,omitempty"`
}

func (m *UpdateStatusRequest) Reset()                    { *m = UpdateStatusRequest{} }
func (m *UpdateStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateStatusRequest) ProtoMessage()               {}
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateStatusRequest) GetLog() *ga4gh_task_exec.TaskOpLog {
	if m != nil {
		return m.Log
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkerInfo)(nil), "ga4gh_task_ref.WorkerInfo")
	proto.RegisterType((*JobRequest)(nil), "ga4gh_task_ref.JobRequest")
	proto.RegisterType((*JobResult)(nil), "ga4gh_task_ref.JobResult")
	proto.RegisterType((*UpdateStatusRequest)(nil), "ga4gh_task_ref.UpdateStatusRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Scheduler service

type SchedulerClient interface {
	GetJobToRun(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResult, error)
	UpdateTaskOpStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*ga4gh_task_exec.TaskOpId, error)
	WorkerPing(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*WorkerInfo, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetJobToRun(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResult, error) {
	out := new(JobResult)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/GetJobToRun", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UpdateTaskOpStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*ga4gh_task_exec.TaskOpId, error) {
	out := new(ga4gh_task_exec.TaskOpId)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/UpdateTaskOpStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) WorkerPing(ctx context.Context, in *WorkerInfo, opts ...grpc.CallOption) (*WorkerInfo, error) {
	out := new(WorkerInfo)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/WorkerPing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	GetJobToRun(context.Context, *JobRequest) (*JobResult, error)
	UpdateTaskOpStatus(context.Context, *UpdateStatusRequest) (*ga4gh_task_exec.TaskOpId, error)
	WorkerPing(context.Context, *WorkerInfo) (*WorkerInfo, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_GetJobToRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetJobToRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/GetJobToRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetJobToRun(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UpdateTaskOpStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UpdateTaskOpStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/UpdateTaskOpStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UpdateTaskOpStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_WorkerPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).WorkerPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/WorkerPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).WorkerPing(ctx, req.(*WorkerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ga4gh_task_ref.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobToRun",
			Handler:    _Scheduler_GetJobToRun_Handler,
		},
		{
			MethodName: "UpdateTaskOpStatus",
			Handler:    _Scheduler_UpdateTaskOpStatus_Handler,
		},
		{
			MethodName: "WorkerPing",
			Handler:    _Scheduler_WorkerPing_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0x4b, 0x4f, 0xc2, 0x40,
	0x10, 0xa6, 0xbc, 0x62, 0x07, 0x43, 0x64, 0x35, 0x5a, 0x7b, 0x32, 0x35, 0x44, 0x63, 0x62, 0x0f,
	0xd5, 0xab, 0x57, 0x0c, 0xc6, 0x44, 0x03, 0xa8, 0xc7, 0xa6, 0xd0, 0xa1, 0x10, 0x6a, 0xb7, 0x76,
	0x77, 0xa3, 0xbf, 0xc3, 0xbf, 0xea, 0x1f, 0x70, 0x1f, 0x10, 0x79, 0xc9, 0x6d, 0xd3, 0xf9, 0xe6,
	0x7b, 0x75, 0xa0, 0xc5, 0x23, 0x36, 0x0b, 0x3f, 0x69, 0x31, 0xc3, 0xc2, 0xcf, 0x0b, 0xca, 0x29,
	0x69, 0x26, 0xd1, 0x6d, 0x32, 0x09, 0xf5, 0xa0, 0xc0, 0xb1, 0x7b, 0xa4, 0x5f, 0xf8, 0x85, 0x23,
	0xc1, 0xa7, 0x34, 0x33, 0x28, 0xef, 0x0e, 0xe0, 0x4d, 0x6f, 0x75, 0xb3, 0x31, 0x25, 0x00, 0xe5,
	0x69, 0xec, 0x58, 0x67, 0xd6, 0xa5, 0x4d, 0x5a, 0x60, 0xa7, 0x11, 0xe3, 0x61, 0x3e, 0xcd, 0x12,
	0xa7, 0x2c, 0x3f, 0x55, 0xc8, 0x01, 0xec, 0x4d, 0x28, 0xe3, 0x59, 0xf4, 0x8e, 0x4e, 0x45, 0x81,
	0xbc, 0x04, 0xe0, 0x81, 0x0e, 0x7b, 0xf8, 0x21, 0x90, 0x71, 0x72, 0x05, 0x75, 0x63, 0x41, 0x53,
	0x34, 0x02, 0xd7, 0x5f, 0xf5, 0xe0, 0x2f, 0x49, 0x5d, 0x83, 0x5d, 0x20, 0xa3, 0xa2, 0x18, 0x21,
	0xd3, 0xf4, 0x6b, 0x70, 0x65, 0xd4, 0xef, 0x2d, 0x10, 0x5e, 0x00, 0xb6, 0x16, 0x62, 0x22, 0xe5,
	0xa4, 0x0d, 0x55, 0x85, 0x99, 0xab, 0x9c, 0x6c, 0xac, 0x0d, 0xe4, 0xeb, 0x29, 0xf7, 0xbe, 0x2d,
	0x38, 0x7c, 0xc9, 0xe3, 0x88, 0x63, 0x9f, 0x47, 0x5c, 0xb0, 0x85, 0xcd, 0xe5, 0x94, 0xfb, 0x50,
	0x65, 0x1c, 0xf3, 0x79, 0xc0, 0x36, 0xd4, 0x98, 0x84, 0x9a, 0x74, 0xcd, 0xe0, 0x78, 0x83, 0x59,
	0x11, 0x21, 0xb9, 0x80, 0x4a, 0x4a, 0x13, 0xa7, 0xfa, 0x8f, 0x6b, 0x23, 0xff, 0x48, 0x13, 0xd5,
	0xa1, 0x29, 0x24, 0x94, 0x82, 0x35, 0x25, 0x18, 0xfc, 0x58, 0x60, 0xf7, 0x47, 0x13, 0x8c, 0x45,
	0x8a, 0x05, 0xe9, 0x40, 0xe3, 0x1e, 0xb9, 0x4c, 0x36, 0xa0, 0x3d, 0x91, 0x91, 0x8d, 0xc2, 0xfe,
	0xca, 0x75, 0x4f, 0xb7, 0xce, 0x54, 0x1f, 0x5e, 0x89, 0xbc, 0x02, 0x31, 0x49, 0x8d, 0xb6, 0xc9,
	0x4b, 0xce, 0xd7, 0x57, 0xb6, 0xb4, 0xb1, 0xca, 0xbb, 0xe4, 0xbf, 0x1b, 0x4b, 0xde, 0xce, 0xe2,
	0x3c, 0x9e, 0xe5, 0x15, 0x90, 0x1d, 0xff, 0xd3, 0xdd, 0x31, 0xf3, 0x4a, 0xc3, 0xba, 0xbe, 0xb6,
	0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xc8, 0xba, 0x55, 0xa8, 0x02, 0x00, 0x00,
}
