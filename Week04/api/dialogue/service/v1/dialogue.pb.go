// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dialogue/service/v1/dialogue.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type AddQARequest struct {
	Question             string   `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	Answers              []string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddQARequest) Reset()         { *m = AddQARequest{} }
func (m *AddQARequest) String() string { return proto.CompactTextString(m) }
func (*AddQARequest) ProtoMessage()    {}
func (*AddQARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c107f272aa950fc, []int{0}
}
func (m *AddQARequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddQARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddQARequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddQARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddQARequest.Merge(m, src)
}
func (m *AddQARequest) XXX_Size() int {
	return m.Size()
}
func (m *AddQARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddQARequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddQARequest proto.InternalMessageInfo

func (m *AddQARequest) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *AddQARequest) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

type AddQAReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddQAReply) Reset()         { *m = AddQAReply{} }
func (m *AddQAReply) String() string { return proto.CompactTextString(m) }
func (*AddQAReply) ProtoMessage()    {}
func (*AddQAReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c107f272aa950fc, []int{1}
}
func (m *AddQAReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddQAReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddQAReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddQAReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddQAReply.Merge(m, src)
}
func (m *AddQAReply) XXX_Size() int {
	return m.Size()
}
func (m *AddQAReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddQAReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddQAReply proto.InternalMessageInfo

type TalkRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TalkRequest) Reset()         { *m = TalkRequest{} }
func (m *TalkRequest) String() string { return proto.CompactTextString(m) }
func (*TalkRequest) ProtoMessage()    {}
func (*TalkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c107f272aa950fc, []int{2}
}
func (m *TalkRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TalkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TalkRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TalkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TalkRequest.Merge(m, src)
}
func (m *TalkRequest) XXX_Size() int {
	return m.Size()
}
func (m *TalkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TalkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TalkRequest proto.InternalMessageInfo

func (m *TalkRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type TalkReply struct {
	Tts                  string   `protobuf:"bytes,1,opt,name=tts,proto3" json:"tts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TalkReply) Reset()         { *m = TalkReply{} }
func (m *TalkReply) String() string { return proto.CompactTextString(m) }
func (*TalkReply) ProtoMessage()    {}
func (*TalkReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c107f272aa950fc, []int{3}
}
func (m *TalkReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TalkReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TalkReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TalkReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TalkReply.Merge(m, src)
}
func (m *TalkReply) XXX_Size() int {
	return m.Size()
}
func (m *TalkReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TalkReply.DiscardUnknown(m)
}

var xxx_messageInfo_TalkReply proto.InternalMessageInfo

func (m *TalkReply) GetTts() string {
	if m != nil {
		return m.Tts
	}
	return ""
}

func init() {
	proto.RegisterType((*AddQARequest)(nil), "dialogue.service.v1.AddQARequest")
	proto.RegisterType((*AddQAReply)(nil), "dialogue.service.v1.AddQAReply")
	proto.RegisterType((*TalkRequest)(nil), "dialogue.service.v1.TalkRequest")
	proto.RegisterType((*TalkReply)(nil), "dialogue.service.v1.TalkReply")
}

func init() { proto.RegisterFile("dialogue/service/v1/dialogue.proto", fileDescriptor_0c107f272aa950fc) }

var fileDescriptor_0c107f272aa950fc = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xc9, 0x4c, 0xcc,
	0xc9, 0x4f, 0x2f, 0x4d, 0xd5, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x33, 0xd4,
	0x87, 0x89, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0xc3, 0xf9, 0x50, 0x35, 0x7a, 0x65,
	0x86, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79,
	0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x2d, 0x4a, 0x2e, 0x5c, 0x3c, 0x8e,
	0x29, 0x29, 0x81, 0x8e, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0x60,
	0x46, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc1, 0xc5,
	0x9e, 0x98, 0x57, 0x5c, 0x9e, 0x5a, 0x54, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x19, 0x04, 0xe3,
	0x2a, 0xf1, 0x70, 0x71, 0x41, 0x4d, 0x29, 0xc8, 0xa9, 0x54, 0x52, 0xe6, 0xe2, 0x0e, 0x49, 0xcc,
	0xc9, 0x86, 0x19, 0x29, 0xc2, 0xc5, 0x5a, 0x58, 0x9a, 0x5a, 0x54, 0x09, 0x35, 0x0f, 0xc2, 0x51,
	0x92, 0xe5, 0xe2, 0x84, 0x28, 0x2a, 0xc8, 0xa9, 0x14, 0x12, 0xe0, 0x62, 0x2e, 0x29, 0x29, 0x86,
	0x2a, 0x00, 0x31, 0x8d, 0x2e, 0x33, 0x72, 0xf1, 0xbb, 0x40, 0x7d, 0x13, 0x0c, 0xf1, 0x8c, 0x50,
	0x22, 0x17, 0x2b, 0xd8, 0x16, 0x21, 0x45, 0x3d, 0x2c, 0x1e, 0xd5, 0x43, 0xf6, 0x87, 0x94, 0x3c,
	0x3e, 0x25, 0x20, 0x47, 0x8a, 0x34, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x4f, 0x89, 0x13, 0x14, 0x8e,
	0x60, 0x71, 0x2b, 0x46, 0x2d, 0xa1, 0x58, 0x2e, 0x16, 0x90, 0xab, 0x84, 0x14, 0xb0, 0x6a, 0x47,
	0xf2, 0x95, 0x94, 0x1c, 0x1e, 0x15, 0x20, 0xf3, 0x85, 0xc1, 0xe6, 0xf3, 0x0a, 0x71, 0x80, 0xcc,
	0x07, 0x09, 0x5b, 0x31, 0x6a, 0x39, 0x69, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0x89, 0x61, 0x89, 0x56, 0xeb, 0x32, 0xc3,
	0x24, 0x36, 0x70, 0xf4, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0x4d, 0x41, 0x8d, 0xf7,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DialogueServiceClient is the client API for DialogueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DialogueServiceClient interface {
	AddQA(ctx context.Context, in *AddQARequest, opts ...grpc.CallOption) (*AddQAReply, error)
	Talk(ctx context.Context, in *TalkRequest, opts ...grpc.CallOption) (*TalkReply, error)
}

type dialogueServiceClient struct {
	cc *grpc.ClientConn
}

func NewDialogueServiceClient(cc *grpc.ClientConn) DialogueServiceClient {
	return &dialogueServiceClient{cc}
}

func (c *dialogueServiceClient) AddQA(ctx context.Context, in *AddQARequest, opts ...grpc.CallOption) (*AddQAReply, error) {
	out := new(AddQAReply)
	err := c.cc.Invoke(ctx, "/dialogue.service.v1.DialogueService/AddQA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dialogueServiceClient) Talk(ctx context.Context, in *TalkRequest, opts ...grpc.CallOption) (*TalkReply, error) {
	out := new(TalkReply)
	err := c.cc.Invoke(ctx, "/dialogue.service.v1.DialogueService/Talk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DialogueServiceServer is the server API for DialogueService service.
type DialogueServiceServer interface {
	AddQA(context.Context, *AddQARequest) (*AddQAReply, error)
	Talk(context.Context, *TalkRequest) (*TalkReply, error)
}

// UnimplementedDialogueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDialogueServiceServer struct {
}

func (*UnimplementedDialogueServiceServer) AddQA(ctx context.Context, req *AddQARequest) (*AddQAReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQA not implemented")
}
func (*UnimplementedDialogueServiceServer) Talk(ctx context.Context, req *TalkRequest) (*TalkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Talk not implemented")
}

func RegisterDialogueServiceServer(s *grpc.Server, srv DialogueServiceServer) {
	s.RegisterService(&_DialogueService_serviceDesc, srv)
}

func _DialogueService_AddQA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddQARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DialogueServiceServer).AddQA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialogue.service.v1.DialogueService/AddQA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DialogueServiceServer).AddQA(ctx, req.(*AddQARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DialogueService_Talk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TalkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DialogueServiceServer).Talk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialogue.service.v1.DialogueService/Talk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DialogueServiceServer).Talk(ctx, req.(*TalkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DialogueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialogue.service.v1.DialogueService",
	HandlerType: (*DialogueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddQA",
			Handler:    _DialogueService_AddQA_Handler,
		},
		{
			MethodName: "Talk",
			Handler:    _DialogueService_Talk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dialogue/service/v1/dialogue.proto",
}

func (m *AddQARequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddQARequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddQARequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Answers) > 0 {
		for iNdEx := len(m.Answers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Answers[iNdEx])
			copy(dAtA[i:], m.Answers[iNdEx])
			i = encodeVarintDialogue(dAtA, i, uint64(len(m.Answers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintDialogue(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddQAReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddQAReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddQAReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *TalkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TalkRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TalkRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintDialogue(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TalkReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TalkReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TalkReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tts) > 0 {
		i -= len(m.Tts)
		copy(dAtA[i:], m.Tts)
		i = encodeVarintDialogue(dAtA, i, uint64(len(m.Tts)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDialogue(dAtA []byte, offset int, v uint64) int {
	offset -= sovDialogue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddQARequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovDialogue(uint64(l))
	}
	if len(m.Answers) > 0 {
		for _, s := range m.Answers {
			l = len(s)
			n += 1 + l + sovDialogue(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AddQAReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TalkRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovDialogue(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TalkReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tts)
	if l > 0 {
		n += 1 + l + sovDialogue(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDialogue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDialogue(x uint64) (n int) {
	return sovDialogue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddQARequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDialogue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddQARequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddQARequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDialogue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDialogue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDialogue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDialogue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDialogue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDialogue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answers = append(m.Answers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDialogue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddQAReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDialogue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddQAReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddQAReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDialogue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TalkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDialogue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TalkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TalkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDialogue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDialogue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDialogue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDialogue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TalkReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDialogue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TalkReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TalkReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDialogue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDialogue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDialogue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tts = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDialogue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDialogue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDialogue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDialogue
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDialogue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDialogue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDialogue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDialogue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDialogue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDialogue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDialogue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDialogue = fmt.Errorf("proto: unexpected end of group")
)