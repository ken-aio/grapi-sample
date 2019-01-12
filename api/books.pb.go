// Code generated by protoc-gen-go. DO NOT EDIT.
// source: books.proto

package api_pb // import "github.com/ken-aio/grapi-sample/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Book struct {
	BookId               string   `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_813ce36626b0b729, []int{0}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (dst *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(dst, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

type ListBooksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksRequest) Reset()         { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()    {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_813ce36626b0b729, []int{1}
}
func (m *ListBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksRequest.Unmarshal(m, b)
}
func (m *ListBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksRequest.Marshal(b, m, deterministic)
}
func (dst *ListBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksRequest.Merge(dst, src)
}
func (m *ListBooksRequest) XXX_Size() int {
	return xxx_messageInfo_ListBooksRequest.Size(m)
}
func (m *ListBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksRequest proto.InternalMessageInfo

type ListBooksResponse struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksResponse) Reset()         { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()    {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_813ce36626b0b729, []int{2}
}
func (m *ListBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksResponse.Unmarshal(m, b)
}
func (m *ListBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksResponse.Marshal(b, m, deterministic)
}
func (dst *ListBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksResponse.Merge(dst, src)
}
func (m *ListBooksResponse) XXX_Size() int {
	return xxx_messageInfo_ListBooksResponse.Size(m)
}
func (m *ListBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksResponse proto.InternalMessageInfo

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type GetBookRequest struct {
	BookId               string   `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_813ce36626b0b729, []int{3}
}
func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (dst *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(dst, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

type CreateBookRequest struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBookRequest) Reset()         { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()    {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_813ce36626b0b729, []int{4}
}
func (m *CreateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBookRequest.Unmarshal(m, b)
}
func (m *CreateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBookRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBookRequest.Merge(dst, src)
}
func (m *CreateBookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBookRequest.Size(m)
}
func (m *CreateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBookRequest proto.InternalMessageInfo

func (m *CreateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type UpdateBookRequest struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBookRequest) Reset()         { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()    {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_813ce36626b0b729, []int{5}
}
func (m *UpdateBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBookRequest.Unmarshal(m, b)
}
func (m *UpdateBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBookRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBookRequest.Merge(dst, src)
}
func (m *UpdateBookRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBookRequest.Size(m)
}
func (m *UpdateBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBookRequest proto.InternalMessageInfo

func (m *UpdateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type DeleteBookRequest struct {
	BookId               string   `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookRequest) Reset()         { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()    {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_books_813ce36626b0b729, []int{6}
}
func (m *DeleteBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookRequest.Unmarshal(m, b)
}
func (m *DeleteBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookRequest.Merge(dst, src)
}
func (m *DeleteBookRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBookRequest.Size(m)
}
func (m *DeleteBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookRequest proto.InternalMessageInfo

func (m *DeleteBookRequest) GetBookId() string {
	if m != nil {
		return m.BookId
	}
	return ""
}

func init() {
	proto.RegisterType((*Book)(nil), "ken_aio.grapi_sample.Book")
	proto.RegisterType((*ListBooksRequest)(nil), "ken_aio.grapi_sample.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "ken_aio.grapi_sample.ListBooksResponse")
	proto.RegisterType((*GetBookRequest)(nil), "ken_aio.grapi_sample.GetBookRequest")
	proto.RegisterType((*CreateBookRequest)(nil), "ken_aio.grapi_sample.CreateBookRequest")
	proto.RegisterType((*UpdateBookRequest)(nil), "ken_aio.grapi_sample.UpdateBookRequest")
	proto.RegisterType((*DeleteBookRequest)(nil), "ken_aio.grapi_sample.DeleteBookRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/ken_aio.grapi_sample.BookService/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/ken_aio.grapi_sample.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/ken_aio.grapi_sample.BookService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/ken_aio.grapi_sample.BookService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ken_aio.grapi_sample.BookService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*empty.Empty, error)
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ken_aio.grapi_sample.BookService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ken_aio.grapi_sample.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ken_aio.grapi_sample.BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ken_aio.grapi_sample.BookService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ken_aio.grapi_sample.BookService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ken_aio.grapi_sample.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookService_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books.proto",
}

func init() { proto.RegisterFile("books.proto", fileDescriptor_books_813ce36626b0b729) }

var fileDescriptor_books_813ce36626b0b729 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0x83, 0x22, 0xc8, 0x21, 0x21, 0x30, 0x41, 0x25, 0x55, 0x23, 0xa9, 0x46, 0x91, 0xc0,
	0xd4, 0xe0, 0x9d, 0xde, 0x81, 0xc4, 0x98, 0x78, 0x85, 0xf1, 0xc6, 0x1b, 0x32, 0x85, 0x63, 0x9d,
	0xb4, 0x74, 0xc6, 0x76, 0x20, 0xd9, 0x6c, 0xf6, 0x66, 0x5f, 0x61, 0x5f, 0x63, 0xdf, 0x66, 0x5f,
	0x61, 0x1f, 0x64, 0x33, 0xd3, 0xb2, 0xc0, 0x96, 0x76, 0x2f, 0xf6, 0x72, 0x3a, 0xff, 0xfc, 0xdf,
	0x39, 0xff, 0xe9, 0x81, 0xba, 0x2b, 0x84, 0x1f, 0x53, 0x19, 0x09, 0x25, 0x48, 0xdb, 0xc7, 0x70,
	0xce, 0xb8, 0xa0, 0x5e, 0xc4, 0x24, 0x9f, 0xc7, 0x6c, 0x25, 0x03, 0xb4, 0x5e, 0x79, 0x42, 0x78,
	0x01, 0x3a, 0x4c, 0x72, 0x87, 0x85, 0xa1, 0x50, 0x4c, 0x71, 0x11, 0xa6, 0x6f, 0xac, 0x97, 0xe9,
	0xad, 0x39, 0xb9, 0xeb, 0xbf, 0x0e, 0xae, 0xa4, 0x3a, 0x49, 0x2e, 0xed, 0x37, 0x50, 0x1e, 0x0b,
	0xe1, 0x93, 0x17, 0x50, 0xd5, 0x9c, 0x39, 0x5f, 0x76, 0x4a, 0xdd, 0x52, 0xaf, 0x36, 0xab, 0xe8,
	0xe3, 0x8f, 0xa5, 0x4d, 0xa0, 0xf9, 0x93, 0xc7, 0x4a, 0x8b, 0xe2, 0x19, 0xfe, 0x5f, 0x63, 0xac,
	0xec, 0x29, 0xb4, 0xf6, 0xbe, 0xc5, 0x52, 0x84, 0x31, 0x92, 0x4f, 0xf0, 0xc4, 0x54, 0xda, 0x29,
	0x75, 0x1f, 0xf7, 0xea, 0x23, 0x8b, 0x1e, 0x2b, 0x95, 0xea, 0x37, 0xb3, 0x44, 0x68, 0x7f, 0x84,
	0xc6, 0x77, 0x34, 0x2e, 0xa9, 0x71, 0x7e, 0x15, 0x13, 0x68, 0x4d, 0x22, 0x64, 0x0a, 0xf7, 0xd5,
	0x14, 0xca, 0xfa, 0xda, 0x48, 0x8b, 0x81, 0x46, 0xa7, 0x4d, 0x7e, 0xcb, 0xe5, 0x03, 0x4d, 0x06,
	0xd0, 0xfa, 0x86, 0x01, 0x1e, 0x9a, 0xe4, 0xd5, 0x3d, 0xba, 0x2c, 0x43, 0x5d, 0x0b, 0x7f, 0x61,
	0xb4, 0xe1, 0x0b, 0x24, 0x01, 0xd4, 0x6e, 0x93, 0x23, 0xef, 0x8f, 0xc3, 0xee, 0xc6, 0x6d, 0x7d,
	0xb8, 0x57, 0x97, 0x8c, 0xc0, 0x6e, 0x9c, 0x5f, 0x5d, 0x5f, 0x3c, 0x7a, 0x4a, 0x2a, 0x8e, 0x09,
	0x98, 0x20, 0x54, 0xd3, 0x80, 0xc9, 0xbb, 0xe3, 0x1e, 0x87, 0xf9, 0x5b, 0x05, 0xed, 0xdb, 0x1d,
	0x63, 0x4e, 0x48, 0x33, 0x31, 0x77, 0x4e, 0xd3, 0x8e, 0xcf, 0x88, 0x07, 0xb0, 0x1b, 0x0e, 0xc9,
	0xa9, 0x36, 0x33, 0xbe, 0x42, 0x58, 0xdb, 0xc0, 0x1a, 0x76, 0xda, 0xc9, 0x17, 0x93, 0x3d, 0xd9,
	0x00, 0xec, 0x06, 0x98, 0x07, 0xca, 0x8c, 0xb8, 0x10, 0xf4, 0xd6, 0x80, 0x5e, 0x8f, 0x9e, 0xed,
	0x77, 0x45, 0xb7, 0xad, 0xa5, 0x5c, 0x0f, 0x60, 0x37, 0xf3, 0x3c, 0x6e, 0xe6, 0xaf, 0xb0, 0x9e,
	0xd3, 0x64, 0xf3, 0xe8, 0x76, 0xf3, 0xe8, 0x54, 0x6f, 0xde, 0x36, 0xc9, 0x7e, 0x26, 0xc9, 0xf1,
	0xe0, 0x4f, 0xdf, 0xe3, 0xea, 0xdf, 0xda, 0xa5, 0x0b, 0xb1, 0x72, 0x7c, 0x0c, 0x87, 0x8c, 0x0b,
	0xc7, 0x60, 0x86, 0x09, 0x46, 0xef, 0xf8, 0x57, 0x4d, 0x95, 0xae, 0x5b, 0x31, 0xbe, 0x9f, 0x6f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0xb1, 0x00, 0xf6, 0x22, 0x04, 0x00, 0x00,
}