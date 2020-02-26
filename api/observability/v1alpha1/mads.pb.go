// Code generated by protoc-gen-go. DO NOT EDIT.
// source: observability/v1alpha1/mads.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// MDS resource type.
//
// Describes a group of targets that need to be monitored.
type MonitoringAssignment struct {
	// MDS resource name.
	//
	// E.g., `/meshes/default/services/backend` or
	// `/meshes/default/dataplanes/backend-01`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of targets that need to be monitored.
	Targets []*MonitoringAssignment_Target `protobuf:"bytes,2,rep,name=targets,proto3" json:"targets,omitempty"`
	// Labels associated with every target in that assignment.
	//
	// E.g., `["job" : "backend"]`.
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MonitoringAssignment) Reset()         { *m = MonitoringAssignment{} }
func (m *MonitoringAssignment) String() string { return proto.CompactTextString(m) }
func (*MonitoringAssignment) ProtoMessage()    {}
func (*MonitoringAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4422a3f93d51774, []int{0}
}

func (m *MonitoringAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoringAssignment.Unmarshal(m, b)
}
func (m *MonitoringAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoringAssignment.Marshal(b, m, deterministic)
}
func (m *MonitoringAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringAssignment.Merge(m, src)
}
func (m *MonitoringAssignment) XXX_Size() int {
	return xxx_messageInfo_MonitoringAssignment.Size(m)
}
func (m *MonitoringAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringAssignment proto.InternalMessageInfo

func (m *MonitoringAssignment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MonitoringAssignment) GetTargets() []*MonitoringAssignment_Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *MonitoringAssignment) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Describes a single target that needs to be monitored.
type MonitoringAssignment_Target struct {
	// Labels associated with that particular target.
	//
	// E.g.,
	// `[
	//    "__address__" :      "192.168.0.1:8080",
	//    "__metrics_path__" : "/metrics"]`,
	//    "instance" :         "backend-01",
	//  ]`.
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MonitoringAssignment_Target) Reset()         { *m = MonitoringAssignment_Target{} }
func (m *MonitoringAssignment_Target) String() string { return proto.CompactTextString(m) }
func (*MonitoringAssignment_Target) ProtoMessage()    {}
func (*MonitoringAssignment_Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4422a3f93d51774, []int{0, 0}
}

func (m *MonitoringAssignment_Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoringAssignment_Target.Unmarshal(m, b)
}
func (m *MonitoringAssignment_Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoringAssignment_Target.Marshal(b, m, deterministic)
}
func (m *MonitoringAssignment_Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringAssignment_Target.Merge(m, src)
}
func (m *MonitoringAssignment_Target) XXX_Size() int {
	return xxx_messageInfo_MonitoringAssignment_Target.Size(m)
}
func (m *MonitoringAssignment_Target) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringAssignment_Target.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringAssignment_Target proto.InternalMessageInfo

func (m *MonitoringAssignment_Target) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*MonitoringAssignment)(nil), "kuma.observability.v1alpha1.MonitoringAssignment")
	proto.RegisterMapType((map[string]string)(nil), "kuma.observability.v1alpha1.MonitoringAssignment.LabelsEntry")
	proto.RegisterType((*MonitoringAssignment_Target)(nil), "kuma.observability.v1alpha1.MonitoringAssignment.Target")
	proto.RegisterMapType((map[string]string)(nil), "kuma.observability.v1alpha1.MonitoringAssignment.Target.LabelsEntry")
}

func init() { proto.RegisterFile("observability/v1alpha1/mads.proto", fileDescriptor_e4422a3f93d51774) }

var fileDescriptor_e4422a3f93d51774 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0x86, 0x59, 0x3b, 0x5c, 0x60, 0xd3, 0xa0, 0x55, 0x24, 0x2c, 0x5f, 0x04, 0xc7, 0x91, 0xe2,
	0xa0, 0x58, 0x93, 0xa3, 0x09, 0x27, 0x51, 0x70, 0x3a, 0xa8, 0xa0, 0x71, 0xa0, 0x41, 0x34, 0x73,
	0xe7, 0x91, 0xb3, 0xca, 0x7a, 0xd7, 0xec, 0xee, 0xad, 0xe4, 0x96, 0x02, 0xd1, 0xf3, 0x04, 0x3c,
	0x08, 0x4f, 0xc1, 0x2b, 0xf0, 0x14, 0x54, 0x88, 0x75, 0x7c, 0xe2, 0x22, 0x13, 0x09, 0x50, 0xba,
	0xf5, 0xfc, 0xff, 0xfc, 0xdf, 0xc8, 0x9a, 0xa1, 0xf7, 0xf4, 0xd2, 0xa2, 0xf1, 0xb0, 0x14, 0x52,
	0xb8, 0x26, 0xf3, 0x47, 0x20, 0xeb, 0x53, 0x38, 0xca, 0x2a, 0x28, 0x2c, 0xaf, 0x8d, 0x76, 0x9a,
	0x0d, 0xcf, 0xd6, 0x15, 0xf0, 0x2d, 0x1f, 0xef, 0x7c, 0xe9, 0x01, 0x2a, 0xaf, 0x9b, 0x0c, 0x6a,
	0x91, 0xf9, 0x69, 0x56, 0x08, 0xbb, 0xd2, 0x1e, 0x4d, 0xd3, 0xb6, 0xa6, 0x07, 0xa5, 0xd6, 0xa5,
	0xc4, 0x20, 0x83, 0x52, 0xda, 0x81, 0x13, 0x5a, 0x9d, 0x07, 0xa7, 0xb7, 0x3d, 0x48, 0x51, 0x80,
	0xc3, 0xac, 0x7b, 0xb4, 0xc2, 0xf8, 0x6b, 0x4c, 0xf7, 0x5f, 0x69, 0x25, 0x9c, 0x36, 0x42, 0x95,
	0xcf, 0xac, 0x15, 0xa5, 0xaa, 0x50, 0x39, 0x36, 0xa4, 0x3b, 0x0a, 0x2a, 0x4c, 0xc8, 0x88, 0x4c,
	0x6e, 0xce, 0x77, 0x7f, 0xcc, 0x77, 0x4c, 0x34, 0x22, 0x79, 0x28, 0xb2, 0x9c, 0xee, 0x3a, 0x30,
	0x25, 0x3a, 0x9b, 0x44, 0xa3, 0x78, 0xb2, 0x37, 0x3d, 0xe6, 0x97, 0x4c, 0xce, 0xfb, 0x00, 0xfc,
	0x75, 0x08, 0xc8, 0xbb, 0x20, 0xf6, 0x86, 0x0e, 0x24, 0x2c, 0x51, 0xda, 0x24, 0x0e, 0x91, 0x4f,
	0xff, 0x3e, 0xf2, 0x65, 0xe8, 0x7f, 0xae, 0x9c, 0x69, 0xf2, 0xf3, 0xb0, 0xf4, 0x0b, 0xa1, 0x83,
	0x16, 0xc5, 0xde, 0x6d, 0x08, 0x24, 0x10, 0x16, 0xff, 0x3a, 0x74, 0x2f, 0xe8, 0x09, 0xdd, 0xfb,
	0xad, 0xcc, 0x6e, 0xd1, 0xf8, 0x0c, 0x9b, 0xf6, 0xf7, 0xe5, 0xbf, 0x9e, 0x6c, 0x9f, 0x5e, 0xf7,
	0x20, 0xd7, 0x98, 0x44, 0xa1, 0xd6, 0x7e, 0xcc, 0xa2, 0x63, 0xf2, 0x1f, 0xad, 0xd3, 0x8f, 0x31,
	0x3d, 0xec, 0x9b, 0x74, 0xd1, 0xad, 0xc7, 0x09, 0x1a, 0x2f, 0x56, 0xc8, 0x2a, 0x9a, 0x2e, 0x50,
	0x3a, 0xe8, 0x33, 0x5b, 0x76, 0x9f, 0x87, 0xe5, 0xe2, 0x50, 0x0b, 0xee, 0xa7, 0x3c, 0x38, 0x37,
	0x11, 0x39, 0xbe, 0x5f, 0xa3, 0x75, 0xe9, 0xe1, 0xe5, 0x26, 0x5b, 0x6b, 0x65, 0x71, 0x7c, 0x6d,
	0x42, 0x1e, 0x11, 0x56, 0xd0, 0xe1, 0x89, 0x33, 0x08, 0x55, 0x3f, 0xef, 0xce, 0x85, 0xa8, 0x8b,
	0xa8, 0xbb, 0x7f, 0xd4, 0xb7, 0x28, 0x9f, 0x08, 0x4d, 0x5f, 0xa0, 0x5b, 0x9d, 0x5e, 0x11, 0xe5,
	0xc1, 0x87, 0x6f, 0xdf, 0x3f, 0x47, 0xe9, 0x38, 0xd9, 0xba, 0xb8, 0x59, 0xb5, 0xa1, 0x05, 0x3d,
	0x9e, 0x91, 0x87, 0x73, 0xfa, 0xf6, 0x46, 0xb7, 0x3a, 0xcb, 0x41, 0xb8, 0xad, 0xc7, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x6d, 0x58, 0x50, 0xf2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitoringAssignmentDiscoveryServiceClient is the client API for MonitoringAssignmentDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitoringAssignmentDiscoveryServiceClient interface {
	DeltaMonitoringAssignments(ctx context.Context, opts ...grpc.CallOption) (MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignmentsClient, error)
	StreamMonitoringAssignments(ctx context.Context, opts ...grpc.CallOption) (MonitoringAssignmentDiscoveryService_StreamMonitoringAssignmentsClient, error)
	FetchMonitoringAssignments(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type monitoringAssignmentDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitoringAssignmentDiscoveryServiceClient(cc *grpc.ClientConn) MonitoringAssignmentDiscoveryServiceClient {
	return &monitoringAssignmentDiscoveryServiceClient{cc}
}

func (c *monitoringAssignmentDiscoveryServiceClient) DeltaMonitoringAssignments(ctx context.Context, opts ...grpc.CallOption) (MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitoringAssignmentDiscoveryService_serviceDesc.Streams[0], "/kuma.observability.v1alpha1.MonitoringAssignmentDiscoveryService/DeltaMonitoringAssignments", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsClient{stream}
	return x, nil
}

type MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignmentsClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsClient struct {
	grpc.ClientStream
}

func (x *monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *monitoringAssignmentDiscoveryServiceClient) StreamMonitoringAssignments(ctx context.Context, opts ...grpc.CallOption) (MonitoringAssignmentDiscoveryService_StreamMonitoringAssignmentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitoringAssignmentDiscoveryService_serviceDesc.Streams[1], "/kuma.observability.v1alpha1.MonitoringAssignmentDiscoveryService/StreamMonitoringAssignments", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsClient{stream}
	return x, nil
}

type MonitoringAssignmentDiscoveryService_StreamMonitoringAssignmentsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsClient struct {
	grpc.ClientStream
}

func (x *monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *monitoringAssignmentDiscoveryServiceClient) FetchMonitoringAssignments(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/kuma.observability.v1alpha1.MonitoringAssignmentDiscoveryService/FetchMonitoringAssignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringAssignmentDiscoveryServiceServer is the server API for MonitoringAssignmentDiscoveryService service.
type MonitoringAssignmentDiscoveryServiceServer interface {
	DeltaMonitoringAssignments(MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignmentsServer) error
	StreamMonitoringAssignments(MonitoringAssignmentDiscoveryService_StreamMonitoringAssignmentsServer) error
	FetchMonitoringAssignments(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedMonitoringAssignmentDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMonitoringAssignmentDiscoveryServiceServer struct {
}

func (*UnimplementedMonitoringAssignmentDiscoveryServiceServer) DeltaMonitoringAssignments(srv MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaMonitoringAssignments not implemented")
}
func (*UnimplementedMonitoringAssignmentDiscoveryServiceServer) StreamMonitoringAssignments(srv MonitoringAssignmentDiscoveryService_StreamMonitoringAssignmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMonitoringAssignments not implemented")
}
func (*UnimplementedMonitoringAssignmentDiscoveryServiceServer) FetchMonitoringAssignments(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMonitoringAssignments not implemented")
}

func RegisterMonitoringAssignmentDiscoveryServiceServer(s *grpc.Server, srv MonitoringAssignmentDiscoveryServiceServer) {
	s.RegisterService(&_MonitoringAssignmentDiscoveryService_serviceDesc, srv)
}

func _MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignments_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MonitoringAssignmentDiscoveryServiceServer).DeltaMonitoringAssignments(&monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsServer{stream})
}

type MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignmentsServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsServer struct {
	grpc.ServerStream
}

func (x *monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *monitoringAssignmentDiscoveryServiceDeltaMonitoringAssignmentsServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MonitoringAssignmentDiscoveryService_StreamMonitoringAssignments_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MonitoringAssignmentDiscoveryServiceServer).StreamMonitoringAssignments(&monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsServer{stream})
}

type MonitoringAssignmentDiscoveryService_StreamMonitoringAssignmentsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsServer struct {
	grpc.ServerStream
}

func (x *monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *monitoringAssignmentDiscoveryServiceStreamMonitoringAssignmentsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MonitoringAssignmentDiscoveryService_FetchMonitoringAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringAssignmentDiscoveryServiceServer).FetchMonitoringAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kuma.observability.v1alpha1.MonitoringAssignmentDiscoveryService/FetchMonitoringAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringAssignmentDiscoveryServiceServer).FetchMonitoringAssignments(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitoringAssignmentDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kuma.observability.v1alpha1.MonitoringAssignmentDiscoveryService",
	HandlerType: (*MonitoringAssignmentDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchMonitoringAssignments",
			Handler:    _MonitoringAssignmentDiscoveryService_FetchMonitoringAssignments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaMonitoringAssignments",
			Handler:       _MonitoringAssignmentDiscoveryService_DeltaMonitoringAssignments_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamMonitoringAssignments",
			Handler:       _MonitoringAssignmentDiscoveryService_StreamMonitoringAssignments_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "observability/v1alpha1/mads.proto",
}
