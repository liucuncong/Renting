// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostRet

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
	//手机号
	Mobile string `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	//密码
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	//短信验证码
	SmsCode              string   `protobuf:"bytes,3,opt,name=Sms_code,json=SmsCode,proto3" json:"Sms_code,omitempty"`
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

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Request) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

type Response struct {
	Error  string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	//将SessionId返回
	SessionId            string   `protobuf:"bytes,3,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
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

func (m *Response) GetSessionId() string {
	if m != nil {
		return m.SessionId
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
	proto.RegisterType((*Message)(nil), "go.micro.srv.PostRet.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostRet.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostRet.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.PostRet.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.PostRet.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.PostRet.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.PostRet.Pong")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x57, 0xe7, 0xd6, 0xee, 0x3e, 0xcd, 0x30, 0x64, 0x76, 0x3a, 0x24, 0x0f, 0x3a, 0x5f,
	0xea, 0xd0, 0x8f, 0x20, 0x03, 0x15, 0x06, 0xa3, 0x05, 0xf1, 0x45, 0xa4, 0x5b, 0x2f, 0xa5, 0xb8,
	0xf6, 0xce, 0xdc, 0xcc, 0x3f, 0x5f, 0x5e, 0xa4, 0x69, 0xe6, 0x44, 0x3a, 0x7c, 0x6a, 0x4f, 0x7e,
	0xa7, 0xa7, 0xe7, 0x26, 0x81, 0xc1, 0x4a, 0x91, 0xa6, 0x4b, 0xfc, 0x88, 0xf3, 0xd5, 0x12, 0x37,
	0xcf, 0xc0, 0xac, 0x8a, 0x5e, 0x4a, 0x41, 0x9e, 0x2d, 0x14, 0x05, 0xac, 0xde, 0x82, 0x19, 0xb1,
	0x0e, 0x51, 0xcb, 0x01, 0xb8, 0x53, 0x64, 0x8e, 0x53, 0x14, 0x5d, 0x68, 0x72, 0xfc, 0xd9, 0x77,
	0x4e, 0x9d, 0x51, 0x27, 0x2c, 0x5f, 0xe5, 0x23, 0xb8, 0x21, 0xbe, 0xae, 0x91, 0xb5, 0x38, 0x84,
	0xf6, 0x94, 0xe6, 0xd9, 0x12, 0x2d, 0xb7, 0x4a, 0xf8, 0xe0, 0xcd, 0x62, 0xe6, 0x77, 0x52, 0x49,
	0x7f, 0xcf, 0x90, 0x1f, 0x2d, 0x8e, 0xc0, 0x8b, 0x72, 0x7e, 0x5e, 0x50, 0x82, 0xfd, 0xa6, 0x61,
	0x6e, 0x94, 0xf3, 0x0d, 0x25, 0x28, 0x1f, 0xc0, 0x0b, 0x91, 0x57, 0x54, 0x30, 0x8a, 0x1e, 0xb4,
	0x26, 0x4a, 0x91, 0xb2, 0xc9, 0x95, 0x28, 0x7f, 0x38, 0x51, 0x2a, 0xe7, 0xd4, 0xc6, 0x5a, 0x25,
	0x8e, 0xa1, 0x13, 0x21, 0x73, 0x46, 0xc5, 0x5d, 0x62, 0x53, 0xb7, 0x0b, 0x72, 0x04, 0xdd, 0x48,
	0x2b, 0x8c, 0xf3, 0xac, 0x48, 0x37, 0xd5, 0x7b, 0xd0, 0x5a, 0xd0, 0xba, 0xd0, 0x26, 0xbf, 0x19,
	0x56, 0x42, 0x5e, 0xc0, 0xc1, 0x2f, 0xe7, 0xb6, 0x4a, 0x8d, 0x75, 0x08, 0xfb, 0xb3, 0xac, 0x48,
	0xcb, 0x4a, 0xac, 0x15, 0xbd, 0xa0, 0xc5, 0x56, 0x19, 0x4e, 0xbb, 0xf9, 0xd5, 0x97, 0x03, 0xee,
	0xa4, 0x3a, 0x0b, 0x71, 0x0f, 0xae, 0xdd, 0x7a, 0x71, 0x12, 0xd4, 0x9d, 0x48, 0x60, 0x6b, 0xfb,
	0xc3, 0x5d, 0xb8, 0xea, 0x2a, 0x1b, 0xe2, 0x09, 0xda, 0xd5, 0x08, 0xe2, 0xac, 0xde, 0xfb, 0x77,
	0x2b, 0xfc, 0xf3, 0x7f, 0x7d, 0x9b, 0xf0, 0xb1, 0x23, 0x6e, 0xc1, 0x2b, 0xc7, 0x36, 0xa3, 0xf9,
	0xf5, 0x1f, 0x96, 0xdc, 0xdf, 0xc5, 0xa8, 0x48, 0x65, 0x63, 0xe4, 0x8c, 0x9d, 0x79, 0xdb, 0xdc,
	0xc0, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x11, 0xdc, 0x9b, 0xa0, 0x02, 0x00, 0x00,
}
