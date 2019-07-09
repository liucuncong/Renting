// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostHousesImage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	SessionId string `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	//图片
	Image []byte `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	//房屋id
	Id string `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
	//图片大小
	FileSize int64 `protobuf:"varint,4,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	//图片名
	FileName             string   `protobuf:"bytes,5,opt,name=FileName,proto3" json:"FileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Request) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *Request) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type Response struct {
	Error  string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	//返回url
	Url                  string   `protobuf:"bytes,3,opt,name=Url,proto3" json:"Url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.PostHousesImage.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostHousesImage.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostHousesImage.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.PostHousesImage.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.PostHousesImage.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.PostHousesImage.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.PostHousesImage.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0xfd, 0xd2, 0xf4, 0x77, 0xf8, 0xd0, 0xba, 0x14, 0x09, 0x6d, 0x91, 0x12, 0x50, 0xe2, 0x4d,
	0x04, 0x7d, 0x86, 0x8a, 0x11, 0x94, 0xb2, 0xc5, 0x07, 0x48, 0xdb, 0x21, 0x04, 0xbb, 0x99, 0xba,
	0xb3, 0x15, 0xf5, 0xde, 0xf7, 0x96, 0x6c, 0x36, 0xad, 0x88, 0x3f, 0x57, 0x99, 0x33, 0xe7, 0x64,
	0xe6, 0x1c, 0x66, 0x61, 0xb4, 0xd1, 0x64, 0xe8, 0x02, 0x5f, 0x52, 0xb5, 0x59, 0x63, 0xfd, 0x8d,
	0x6d, 0x57, 0x8c, 0x33, 0x8a, 0x55, 0xbe, 0xd4, 0x14, 0xb3, 0x7e, 0x8e, 0x67, 0xc4, 0xe6, 0x86,
	0xb6, 0x8c, 0x9c, 0xa8, 0x34, 0xc3, 0x70, 0x04, 0x9d, 0x3b, 0x64, 0x4e, 0x33, 0x14, 0x7d, 0xf0,
	0x39, 0x7d, 0x0d, 0xbc, 0x89, 0x17, 0xf5, 0x64, 0x59, 0x86, 0xef, 0x1e, 0x74, 0x24, 0x3e, 0x6d,
	0x91, 0x8d, 0x18, 0x43, 0x6f, 0x8e, 0xcc, 0x39, 0x15, 0xc9, 0xca, 0x69, 0xf6, 0x0d, 0x31, 0x80,
	0x96, 0x9d, 0x17, 0x34, 0x26, 0x5e, 0xf4, 0x5f, 0x56, 0x40, 0x1c, 0x40, 0x23, 0x59, 0x05, 0xbe,
	0x15, 0x37, 0x92, 0x95, 0x18, 0x42, 0xf7, 0x3a, 0x5f, 0xe3, 0x3c, 0x7f, 0xc3, 0xa0, 0x39, 0xf1,
	0x22, 0x5f, 0xee, 0x70, 0xcd, 0xdd, 0xa7, 0x0a, 0x83, 0x96, 0xfd, 0x63, 0x87, 0xc3, 0x5b, 0xe8,
	0x4a, 0xe4, 0x0d, 0x15, 0x8c, 0xe5, 0xa6, 0xa9, 0xd6, 0xa4, 0x9d, 0x87, 0x0a, 0x88, 0x63, 0x68,
	0x4f, 0xb5, 0x56, 0x9c, 0x59, 0x03, 0x3d, 0xe9, 0x50, 0x99, 0xe9, 0x41, 0xaf, 0x9d, 0x85, 0xb2,
	0x0c, 0x23, 0xe8, 0xcf, 0x8d, 0xc6, 0x54, 0xe5, 0x45, 0x56, 0x67, 0x1b, 0x40, 0x6b, 0x49, 0xdb,
	0xc2, 0xd8, 0x99, 0xbe, 0xac, 0x40, 0x78, 0x0e, 0x47, 0x9f, 0x94, 0xfb, 0xf5, 0xdf, 0x48, 0x4f,
	0xa0, 0x39, 0xcb, 0x8b, 0xac, 0xb4, 0xc1, 0x46, 0xd3, 0x23, 0x3a, 0xda, 0x21, 0xcb, 0xd3, 0xcf,
	0xfc, 0xa5, 0x82, 0xce, 0xb4, 0x3a, 0x9a, 0x58, 0xc0, 0xe1, 0x97, 0x1b, 0x89, 0xd3, 0xf8, 0xb7,
	0x13, 0xc6, 0x2e, 0xc5, 0xf0, 0xec, 0x2f, 0x59, 0x15, 0x21, 0xfc, 0xb7, 0x68, 0xdb, 0x97, 0x71,
	0xf5, 0x11, 0x00, 0x00, 0xff, 0xff, 0x86, 0x07, 0xb6, 0x68, 0x38, 0x02, 0x00, 0x00,
}
