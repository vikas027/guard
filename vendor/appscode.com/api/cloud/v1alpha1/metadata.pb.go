// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegionListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
}

func (m *RegionListRequest) Reset()                    { *m = RegionListRequest{} }
func (m *RegionListRequest) String() string            { return proto.CompactTextString(m) }
func (*RegionListRequest) ProtoMessage()               {}
func (*RegionListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *RegionListRequest) GetCloudCredential() string {
	if m != nil {
		return m.CloudCredential
	}
	return ""
}

type RegionListResponse struct {
	Regions []string `protobuf:"bytes,1,rep,name=regions" json:"regions,omitempty"`
}

func (m *RegionListResponse) Reset()                    { *m = RegionListResponse{} }
func (m *RegionListResponse) String() string            { return proto.CompactTextString(m) }
func (*RegionListResponse) ProtoMessage()               {}
func (*RegionListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RegionListResponse) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

type ZoneListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	Region          string `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
}

func (m *ZoneListRequest) Reset()                    { *m = ZoneListRequest{} }
func (m *ZoneListRequest) String() string            { return proto.CompactTextString(m) }
func (*ZoneListRequest) ProtoMessage()               {}
func (*ZoneListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ZoneListRequest) GetCloudCredential() string {
	if m != nil {
		return m.CloudCredential
	}
	return ""
}

func (m *ZoneListRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type ZoneListResponse struct {
	Zones []string `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
}

func (m *ZoneListResponse) Reset()                    { *m = ZoneListResponse{} }
func (m *ZoneListResponse) String() string            { return proto.CompactTextString(m) }
func (*ZoneListResponse) ProtoMessage()               {}
func (*ZoneListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ZoneListResponse) GetZones() []string {
	if m != nil {
		return m.Zones
	}
	return nil
}

type BucketListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	GceProject      string `protobuf:"bytes,2,opt,name=gce_project,json=gceProject" json:"gce_project,omitempty"`
	ClusterUid      string `protobuf:"bytes,3,opt,name=cluster_uid,json=clusterUid" json:"cluster_uid,omitempty"`
	SecretNamespace string `protobuf:"bytes,4,opt,name=secret_namespace,json=secretNamespace" json:"secret_namespace,omitempty"`
	SecretName      string `protobuf:"bytes,5,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	Provider        string `protobuf:"bytes,6,opt,name=provider" json:"provider,omitempty"`
}

func (m *BucketListRequest) Reset()                    { *m = BucketListRequest{} }
func (m *BucketListRequest) String() string            { return proto.CompactTextString(m) }
func (*BucketListRequest) ProtoMessage()               {}
func (*BucketListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *BucketListRequest) GetCloudCredential() string {
	if m != nil {
		return m.CloudCredential
	}
	return ""
}

func (m *BucketListRequest) GetGceProject() string {
	if m != nil {
		return m.GceProject
	}
	return ""
}

func (m *BucketListRequest) GetClusterUid() string {
	if m != nil {
		return m.ClusterUid
	}
	return ""
}

func (m *BucketListRequest) GetSecretNamespace() string {
	if m != nil {
		return m.SecretNamespace
	}
	return ""
}

func (m *BucketListRequest) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *BucketListRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

type BucketListResponse struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *BucketListResponse) Reset()                    { *m = BucketListResponse{} }
func (m *BucketListResponse) String() string            { return proto.CompactTextString(m) }
func (*BucketListResponse) ProtoMessage()               {}
func (*BucketListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *BucketListResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*RegionListRequest)(nil), "appscode.cloud.v1alpha1.RegionListRequest")
	proto.RegisterType((*RegionListResponse)(nil), "appscode.cloud.v1alpha1.RegionListResponse")
	proto.RegisterType((*ZoneListRequest)(nil), "appscode.cloud.v1alpha1.ZoneListRequest")
	proto.RegisterType((*ZoneListResponse)(nil), "appscode.cloud.v1alpha1.ZoneListResponse")
	proto.RegisterType((*BucketListRequest)(nil), "appscode.cloud.v1alpha1.BucketListRequest")
	proto.RegisterType((*BucketListResponse)(nil), "appscode.cloud.v1alpha1.BucketListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metadata service

type MetadataClient interface {
	ListRegions(ctx context.Context, in *RegionListRequest, opts ...grpc.CallOption) (*RegionListResponse, error)
	ListZones(ctx context.Context, in *ZoneListRequest, opts ...grpc.CallOption) (*ZoneListResponse, error)
	ListBuckets(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) ListRegions(ctx context.Context, in *RegionListRequest, opts ...grpc.CallOption) (*RegionListResponse, error) {
	out := new(RegionListResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha1.Metadata/ListRegions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) ListZones(ctx context.Context, in *ZoneListRequest, opts ...grpc.CallOption) (*ZoneListResponse, error) {
	out := new(ZoneListResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha1.Metadata/ListZones", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) ListBuckets(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error) {
	out := new(BucketListResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha1.Metadata/ListBuckets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	ListRegions(context.Context, *RegionListRequest) (*RegionListResponse, error)
	ListZones(context.Context, *ZoneListRequest) (*ZoneListResponse, error)
	ListBuckets(context.Context, *BucketListRequest) (*BucketListResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha1.Metadata/ListRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListRegions(ctx, req.(*RegionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_ListZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListZones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha1.Metadata/ListZones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListZones(ctx, req.(*ZoneListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha1.Metadata/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListBuckets(ctx, req.(*BucketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.cloud.v1alpha1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRegions",
			Handler:    _Metadata_ListRegions_Handler,
		},
		{
			MethodName: "ListZones",
			Handler:    _Metadata_ListZones_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _Metadata_ListBuckets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x13, 0x1a, 0x92, 0xa9, 0xa0, 0xed, 0x0a, 0x81, 0x65, 0x21, 0x5a, 0xf9, 0x94, 0x06,
	0x64, 0xab, 0x70, 0x01, 0x81, 0x38, 0x84, 0x03, 0x17, 0x40, 0x51, 0x04, 0x97, 0x5e, 0xac, 0xed,
	0x7a, 0x64, 0xb6, 0x38, 0xbb, 0x8b, 0x77, 0xdd, 0x03, 0x88, 0x4b, 0x7f, 0x81, 0x1f, 0xe0, 0xca,
	0x07, 0xf0, 0x25, 0xfc, 0x02, 0x07, 0x3e, 0x82, 0x03, 0xf2, 0xae, 0xb7, 0x09, 0x0d, 0xad, 0xd2,
	0xdc, 0x32, 0x6f, 0xdf, 0xcc, 0x7b, 0x99, 0x79, 0x32, 0xdc, 0x9c, 0xa1, 0xa1, 0x39, 0x35, 0x34,
	0x51, 0x95, 0x34, 0x92, 0xdc, 0xa1, 0x4a, 0x69, 0x26, 0x73, 0x4c, 0x58, 0x29, 0xeb, 0x3c, 0x39,
	0x39, 0xa0, 0xa5, 0x7a, 0x4f, 0x0f, 0xa2, 0xbb, 0x85, 0x94, 0x45, 0x89, 0x29, 0x55, 0x3c, 0xa5,
	0x42, 0x48, 0x43, 0x0d, 0x97, 0x42, 0xbb, 0xb6, 0xe8, 0x9e, 0x6f, 0xfb, 0xff, 0x7b, 0xfc, 0x1c,
	0x76, 0xa6, 0x58, 0x70, 0x29, 0x5e, 0x71, 0x6d, 0xa6, 0xf8, 0xb1, 0x46, 0x6d, 0xc8, 0x3e, 0x6c,
	0x5b, 0x91, 0x8c, 0x55, 0x98, 0xa3, 0x30, 0x9c, 0x96, 0x61, 0xb0, 0x17, 0x0c, 0x07, 0xd3, 0x2d,
	0x8b, 0xbf, 0x38, 0x83, 0xe3, 0x04, 0xc8, 0x62, 0xbf, 0x56, 0x52, 0x68, 0x24, 0x21, 0x5c, 0xaf,
	0x2c, 0xaa, 0xc3, 0x60, 0xaf, 0x3b, 0x1c, 0x4c, 0x7d, 0x19, 0xbf, 0x85, 0xad, 0x43, 0x29, 0x70,
	0x3d, 0x35, 0x72, 0x1b, 0x7a, 0x6e, 0x50, 0xd8, 0xb1, 0x84, 0xb6, 0x8a, 0x87, 0xb0, 0x3d, 0x9f,
	0xda, 0x7a, 0xb8, 0x05, 0x1b, 0x9f, 0xa4, 0x40, 0xef, 0xc0, 0x15, 0xf1, 0xef, 0x00, 0x76, 0xc6,
	0x35, 0xfb, 0x80, 0x66, 0x4d, 0x0b, 0xbb, 0xb0, 0x59, 0x30, 0xcc, 0x54, 0x25, 0x8f, 0x91, 0x99,
	0xd6, 0x07, 0x14, 0x0c, 0x27, 0x0e, 0x69, 0x08, 0xac, 0xac, 0xb5, 0xc1, 0x2a, 0xab, 0x79, 0x1e,
	0x76, 0x1d, 0xa1, 0x85, 0xde, 0xf1, 0xbc, 0x11, 0xd3, 0xc8, 0x2a, 0x34, 0x99, 0xa0, 0x33, 0xd4,
	0x8a, 0x32, 0x0c, 0xaf, 0x39, 0x31, 0x87, 0xbf, 0xf1, 0x70, 0x33, 0x6b, 0x81, 0x1a, 0x6e, 0xb8,
	0x59, 0x73, 0x16, 0x89, 0xa0, 0xaf, 0x2a, 0x79, 0xc2, 0x73, 0xac, 0xc2, 0x9e, 0x7d, 0x3d, 0xab,
	0xe3, 0x11, 0x90, 0xc5, 0x7f, 0x3a, 0x5f, 0x8b, 0x95, 0xf5, 0x6b, 0xb1, 0xc5, 0xc3, 0x3f, 0x5d,
	0xe8, 0xbf, 0x6e, 0x03, 0x47, 0xbe, 0x05, 0xb0, 0xe9, 0x7a, 0xec, 0xcd, 0xc8, 0x28, 0xb9, 0x20,
	0x7b, 0xc9, 0x52, 0x74, 0xa2, 0xfb, 0x2b, 0x71, 0x9d, 0x97, 0xf8, 0xf1, 0xe9, 0x8f, 0xb0, 0xd3,
	0x0f, 0x4e, 0x7f, 0xfe, 0xfa, 0xda, 0x79, 0x40, 0x46, 0x69, 0xf6, 0x4f, 0x58, 0x6d, 0x7f, 0xea,
	0xfb, 0xd3, 0x36, 0x41, 0xe9, 0xb1, 0x96, 0x82, 0x7c, 0x0f, 0x60, 0xd0, 0x8c, 0x6a, 0xae, 0xae,
	0xc9, 0xf0, 0x42, 0xd1, 0x73, 0x59, 0x8b, 0xf6, 0x57, 0x60, 0xb6, 0xe6, 0x5e, 0x2e, 0x98, 0x7b,
	0x4a, 0x9e, 0xac, 0x66, 0xee, 0xb3, 0xfb, 0xf1, 0x25, 0xb5, 0x69, 0x73, 0x5e, 0xfd, 0x3a, 0xdd,
	0x31, 0x2e, 0x5b, 0xe7, 0x52, 0x30, 0x2f, 0x59, 0xe7, 0xf2, 0x69, 0xaf, 0xb6, 0xce, 0x23, 0xe7,
	0xc6, 0x5a, 0x1c, 0x3f, 0x83, 0x5d, 0x26, 0x67, 0x73, 0x2d, 0xaa, 0xf8, 0x39, 0xbd, 0xf1, 0x0d,
	0x1f, 0x8f, 0x49, 0xf3, 0xdd, 0x98, 0x04, 0x87, 0x7d, 0xff, 0x74, 0xd4, 0xb3, 0x9f, 0x92, 0x47,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x87, 0xab, 0x2e, 0x15, 0xb3, 0x04, 0x00, 0x00,
}
