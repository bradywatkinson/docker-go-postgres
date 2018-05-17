// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product/product.proto

package product

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ProductQuery
type ProductQuery struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Product              *ProductSchema `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProductQuery) Reset()         { *m = ProductQuery{} }
func (m *ProductQuery) String() string { return proto.CompactTextString(m) }
func (*ProductQuery) ProtoMessage()    {}
func (*ProductQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_1f429c47355f6812, []int{0}
}
func (m *ProductQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductQuery.Unmarshal(m, b)
}
func (m *ProductQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductQuery.Marshal(b, m, deterministic)
}
func (dst *ProductQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductQuery.Merge(dst, src)
}
func (m *ProductQuery) XXX_Size() int {
	return xxx_messageInfo_ProductQuery.Size(m)
}
func (m *ProductQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProductQuery proto.InternalMessageInfo

func (m *ProductQuery) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductQuery) GetProduct() *ProductSchema {
	if m != nil {
		return m.Product
	}
	return nil
}

// The product interface schema
type ProductSchema struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	MerchantId           int32    `protobuf:"varint,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSchema) Reset()         { *m = ProductSchema{} }
func (m *ProductSchema) String() string { return proto.CompactTextString(m) }
func (*ProductSchema) ProtoMessage()    {}
func (*ProductSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_1f429c47355f6812, []int{1}
}
func (m *ProductSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSchema.Unmarshal(m, b)
}
func (m *ProductSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSchema.Marshal(b, m, deterministic)
}
func (dst *ProductSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSchema.Merge(dst, src)
}
func (m *ProductSchema) XXX_Size() int {
	return xxx_messageInfo_ProductSchema.Size(m)
}
func (m *ProductSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSchema proto.InternalMessageInfo

func (m *ProductSchema) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductSchema) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductSchema) GetMerchantId() int32 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

// ProductsQuery describes a Query to get count products starting at id start
type ProductsQuery struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Start                int32    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductsQuery) Reset()         { *m = ProductsQuery{} }
func (m *ProductsQuery) String() string { return proto.CompactTextString(m) }
func (*ProductsQuery) ProtoMessage()    {}
func (*ProductsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_1f429c47355f6812, []int{2}
}
func (m *ProductsQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductsQuery.Unmarshal(m, b)
}
func (m *ProductsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductsQuery.Marshal(b, m, deterministic)
}
func (dst *ProductsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductsQuery.Merge(dst, src)
}
func (m *ProductsQuery) XXX_Size() int {
	return xxx_messageInfo_ProductsQuery.Size(m)
}
func (m *ProductsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProductsQuery proto.InternalMessageInfo

func (m *ProductsQuery) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ProductsQuery) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

// ProductsResponse is a list of products
type ProductsResponse struct {
	Products             []*ProductSchema `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProductsResponse) Reset()         { *m = ProductsResponse{} }
func (m *ProductsResponse) String() string { return proto.CompactTextString(m) }
func (*ProductsResponse) ProtoMessage()    {}
func (*ProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_1f429c47355f6812, []int{3}
}
func (m *ProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductsResponse.Unmarshal(m, b)
}
func (m *ProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductsResponse.Marshal(b, m, deterministic)
}
func (dst *ProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductsResponse.Merge(dst, src)
}
func (m *ProductsResponse) XXX_Size() int {
	return xxx_messageInfo_ProductsResponse.Size(m)
}
func (m *ProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductsResponse proto.InternalMessageInfo

func (m *ProductsResponse) GetProducts() []*ProductSchema {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductQuery)(nil), "product.ProductQuery")
	proto.RegisterType((*ProductSchema)(nil), "product.ProductSchema")
	proto.RegisterType((*ProductsQuery)(nil), "product.ProductsQuery")
	proto.RegisterType((*ProductsResponse)(nil), "product.ProductsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Create a product
	CreateProduct(ctx context.Context, in *ProductSchema, opts ...grpc.CallOption) (*ProductSchema, error)
	// Get a product by id
	GetProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*ProductSchema, error)
	// Update a product by id
	UpdateProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*ProductSchema, error)
	// Delete a product by id
	DeleteProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*wrappers.StringValue, error)
	// GetProducts gets a list of products by start and count
	GetProducts(ctx context.Context, in *ProductsQuery, opts ...grpc.CallOption) (*ProductsResponse, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *ProductSchema, opts ...grpc.CallOption) (*ProductSchema, error) {
	out := new(ProductSchema)
	err := c.cc.Invoke(ctx, "/product.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*ProductSchema, error) {
	out := new(ProductSchema)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*ProductSchema, error) {
	out := new(ProductSchema)
	err := c.cc.Invoke(ctx, "/product.ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *ProductQuery, opts ...grpc.CallOption) (*wrappers.StringValue, error) {
	out := new(wrappers.StringValue)
	err := c.cc.Invoke(ctx, "/product.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *ProductsQuery, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// Create a product
	CreateProduct(context.Context, *ProductSchema) (*ProductSchema, error)
	// Get a product by id
	GetProduct(context.Context, *ProductQuery) (*ProductSchema, error)
	// Update a product by id
	UpdateProduct(context.Context, *ProductQuery) (*ProductSchema, error)
	// Delete a product by id
	DeleteProduct(context.Context, *ProductQuery) (*wrappers.StringValue, error)
	// GetProducts gets a list of products by start and count
	GetProducts(context.Context, *ProductsQuery) (*ProductsResponse, error)
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*ProductSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*ProductQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*ProductQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*ProductQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*ProductsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/product.proto",
}

func init() { proto.RegisterFile("product/product.proto", fileDescriptor_product_1f429c47355f6812) }

var fileDescriptor_product_1f429c47355f6812 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4b, 0x4f, 0xf2, 0x40,
	0x14, 0xa5, 0x85, 0x7e, 0x8f, 0xcb, 0x57, 0xf2, 0x65, 0x02, 0xa6, 0x12, 0xa3, 0xcd, 0xac, 0x58,
	0x15, 0x83, 0x4b, 0x63, 0xe2, 0x2b, 0x18, 0x77, 0x58, 0xa2, 0x5b, 0x33, 0xb4, 0x57, 0xa8, 0x81,
	0xce, 0x64, 0x66, 0xaa, 0xf1, 0xdf, 0xfa, 0x53, 0x0c, 0x9d, 0x0e, 0x3e, 0x08, 0x18, 0x57, 0xed,
	0x3d, 0x73, 0xce, 0xb9, 0xf7, 0xcc, 0x5c, 0xe8, 0x08, 0xc9, 0xd3, 0x22, 0xd1, 0xfd, 0xea, 0x1b,
	0x09, 0xc9, 0x35, 0x27, 0xbf, 0xab, 0xb2, 0xbb, 0x3f, 0xe5, 0x7c, 0x3a, 0xc7, 0x7e, 0x09, 0x4f,
	0x8a, 0x87, 0xfe, 0xb3, 0x64, 0x42, 0xa0, 0x54, 0x86, 0x48, 0x47, 0xf0, 0x6f, 0x64, 0xa8, 0x37,
	0x05, 0xca, 0x17, 0xd2, 0x02, 0x37, 0x4b, 0x03, 0x27, 0x74, 0x7a, 0x5e, 0xec, 0x66, 0x29, 0x39,
	0x04, 0x6b, 0x15, 0xb8, 0xa1, 0xd3, 0x6b, 0x0e, 0x76, 0x22, 0xdb, 0xa9, 0xd2, 0x8d, 0x93, 0x19,
	0x2e, 0x58, 0x6c, 0x69, 0xf4, 0x11, 0xfc, 0x4f, 0x27, 0x6b, 0x96, 0x04, 0x1a, 0x39, 0x5b, 0x60,
	0xe9, 0xf7, 0x37, 0x2e, 0xff, 0x49, 0x1b, 0x3c, 0x21, 0xb3, 0x04, 0x83, 0x7a, 0xe8, 0xf4, 0x9c,
	0xd8, 0x14, 0xe4, 0x00, 0x9a, 0x0b, 0x94, 0xc9, 0x8c, 0xe5, 0xfa, 0x3e, 0x4b, 0x83, 0x46, 0x69,
	0x01, 0x16, 0xba, 0x4e, 0xe9, 0xf1, 0xaa, 0x97, 0x32, 0xe3, 0xb7, 0xc1, 0x4b, 0x78, 0x91, 0xeb,
	0xaa, 0x9d, 0x29, 0x96, 0xa8, 0xd2, 0x4c, 0x9a, 0x08, 0x5e, 0x6c, 0x0a, 0x3a, 0x84, 0xff, 0x56,
	0x1c, 0xa3, 0x12, 0x3c, 0x57, 0x48, 0x06, 0xf0, 0xa7, 0xca, 0xa1, 0x02, 0x27, 0xac, 0x6f, 0xc9,
	0xbb, 0xe2, 0x0d, 0x5e, 0x5d, 0x68, 0xd9, 0x33, 0x94, 0x4f, 0xcb, 0xc1, 0xcf, 0xc0, 0xbf, 0x90,
	0xc8, 0x34, 0x56, 0x38, 0xd9, 0xe0, 0xd2, 0xdd, 0x80, 0xd3, 0x1a, 0x39, 0x01, 0xb8, 0x42, 0x6d,
	0xf5, 0x9d, 0xaf, 0xbc, 0x32, 0xee, 0x16, 0xf9, 0x29, 0xf8, 0xb7, 0x22, 0xfd, 0x30, 0xc1, 0x8f,
	0x1d, 0x86, 0xe0, 0x5f, 0xe2, 0x1c, 0xbf, 0x75, 0xd8, 0x8b, 0xcc, 0x8a, 0x45, 0x76, 0xc5, 0xa2,
	0xb1, 0x96, 0x59, 0x3e, 0xbd, 0x63, 0xf3, 0x02, 0x69, 0x8d, 0x9c, 0x43, 0xf3, 0x3d, 0x88, 0x5a,
	0xbf, 0x09, 0xf3, 0x72, 0xdd, 0xdd, 0x35, 0xdc, 0x3e, 0x0a, 0xad, 0x4d, 0x7e, 0x95, 0xde, 0x47,
	0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x1a, 0x38, 0x7e, 0xee, 0x02, 0x00, 0x00,
}
