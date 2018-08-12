// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pooltrue/pooltrue.proto

package pooltrue

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import mem "github.com/gogo/protobuf/mem"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import encoding_binary "encoding/binary"
import github_com_gogo_protobuf_mem "github.com/gogo/protobuf/mem"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = mem.PoolMarkerNone

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Foo struct {
	One                  uint64          `protobuf:"varint,1,opt,name=one,proto3" json:"one,omitempty"`
	Two                  uint32          `protobuf:"fixed32,2,opt,name=two,proto3" json:"two,omitempty"`
	Bar                  *Bar            `protobuf:"bytes,3,opt,name=bar" json:"bar,omitempty"`
	RepeatedBar          []*Bar          `protobuf:"bytes,4,rep,name=repeated_bar,json=repeatedBar" json:"repeated_bar,omitempty"`
	MapBar               map[uint32]*Bar `protobuf:"bytes,5,rep,name=map_bar,json=mapBar" json:"map_bar,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`

	// poolMarker is used to provide safeguards for pooling functions.
	// This variable should never be modified by users directly.
	poolMarker mem.PoolMarker
}

func (m *Foo) Reset() {
	m.One = 0
	m.Two = 0
	m.Bar = nil
	if m.RepeatedBar != nil {
		m.RepeatedBar = m.RepeatedBar[:0]
	}
	m.MapBar = nil
}
func (m *Foo) String() string {
	m.checkNotRecycled()
	return proto.CompactTextString(m)
}
func (m *Foo) ProtoMessage() {
	m.checkNotRecycled()
}
func (m *Foo) Descriptor() ([]byte, []int) {
	m.checkNotRecycled()
	return fileDescriptor_pooltrue_fbd4e022f4ee01e5, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	m.checkNotRecycled()
	return m.Unmarshal(b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	m.checkNotRecycled()
	if deterministic {
		return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Foo) XXX_Merge(src proto.Message) {
	dst.checkNotRecycled()
	xxx_messageInfo_Foo.Merge(dst, src)
}
func (m *Foo) XXX_Size() int {
	m.checkNotRecycled()
	return m.ProtoSize()
}
func (m *Foo) XXX_DiscardUnknown() {
	m.checkNotRecycled()
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetOne() uint64 {
	m.checkNotRecycled()
	if m != nil {
		return m.One
	}
	return 0
}

func (m *Foo) GetTwo() uint32 {
	m.checkNotRecycled()
	if m != nil {
		return m.Two
	}
	return 0
}

func (m *Foo) GetBar() *Bar {
	m.checkNotRecycled()
	if m != nil {
		return m.Bar
	}
	return nil
}

func (m *Foo) GetRepeatedBar() []*Bar {
	m.checkNotRecycled()
	if m != nil {
		return m.RepeatedBar
	}
	return nil
}

func (m *Foo) GetMapBar() map[uint32]*Bar {
	m.checkNotRecycled()
	if m != nil {
		return m.MapBar
	}
	return nil
}

type Bar struct {
	OneBar               uint64   `protobuf:"varint,1,opt,name=one_bar,json=oneBar,proto3" json:"one_bar,omitempty"`
	TwoBar               uint32   `protobuf:"fixed32,2,opt,name=two_bar,json=twoBar,proto3" json:"two_bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`

	// poolMarker is used to provide safeguards for pooling functions.
	// This variable should never be modified by users directly.
	poolMarker mem.PoolMarker
}

func (m *Bar) Reset() {
	m.OneBar = 0
	m.TwoBar = 0
}
func (m *Bar) String() string {
	m.checkNotRecycled()
	return proto.CompactTextString(m)
}
func (m *Bar) ProtoMessage() {
	m.checkNotRecycled()
}
func (m *Bar) Descriptor() ([]byte, []int) {
	m.checkNotRecycled()
	return fileDescriptor_pooltrue_fbd4e022f4ee01e5, []int{1}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	m.checkNotRecycled()
	return m.Unmarshal(b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	m.checkNotRecycled()
	if deterministic {
		return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Bar) XXX_Merge(src proto.Message) {
	dst.checkNotRecycled()
	xxx_messageInfo_Bar.Merge(dst, src)
}
func (m *Bar) XXX_Size() int {
	m.checkNotRecycled()
	return m.ProtoSize()
}
func (m *Bar) XXX_DiscardUnknown() {
	m.checkNotRecycled()
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetOneBar() uint64 {
	m.checkNotRecycled()
	if m != nil {
		return m.OneBar
	}
	return 0
}

func (m *Bar) GetTwoBar() uint32 {
	m.checkNotRecycled()
	if m != nil {
		return m.TwoBar
	}
	return 0
}

func init() {
	proto.RegisterType((*Foo)(nil), "pooltrue.Foo")
	proto.RegisterMapType((map[uint32]*Bar)(nil), "pooltrue.Foo.MapBarEntry")
	proto.RegisterType((*Bar)(nil), "pooltrue.Bar")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BazClient is the client API for Baz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BazClient interface {
	Hello(ctx context.Context, in *Foo, opts ...grpc.CallOption) (*Bar, error)
}

type bazClient struct {
	cc *grpc.ClientConn
}

func NewBazClient(cc *grpc.ClientConn) BazClient {
	return &bazClient{cc}
}

func (c *bazClient) Hello(ctx context.Context, in *Foo, opts ...grpc.CallOption) (*Bar, error) {
	out := new(Bar)
	err := c.cc.Invoke(ctx, "/pooltrue.Baz/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Baz service

type BazServer interface {
	Hello(context.Context, *Foo) (*Bar, error)
}

func RegisterBazServer(s *grpc.Server, srv BazServer) {
	s.RegisterService(&_Baz_serviceDesc, srv)
}

func _Baz_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := GetFoo()
	defer in.Recycle()
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BazServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pooltrue.Baz/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BazServer).Hello(ctx, req.(*Foo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Baz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pooltrue.Baz",
	HandlerType: (*BazServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Baz_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pooltrue/pooltrue.proto",
}

func (m *Foo) Marshal() (dAtA []byte, err error) {
	m.checkNotRecycled()
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Foo) MarshalPool() (*github_com_gogo_protobuf_mem.Bytes, error) {
	m.checkNotRecycled()
	size := m.ProtoSize()
	bytes := github_com_gogo_protobuf_mem.GetBytes(size)
	bytes.Truncate(size)
	n, err := m.MarshalTo(bytes.Value())
	if err != nil {
		bytes.Recycle()
		return nil, err
	}
	bytes.Truncate(n)
	return bytes, nil
}

func (m *Foo) MarshalTo(dAtA []byte) (int, error) {
	m.checkNotRecycled()
	var i int
	_ = i
	var l int
	_ = l
	if m.One != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPooltrue(dAtA, i, uint64(m.One))
	}
	if m.Two != 0 {
		dAtA[i] = 0x15
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Two))
		i += 4
	}
	if m.Bar != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPooltrue(dAtA, i, uint64(m.Bar.ProtoSize()))
		n1, err := m.Bar.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.RepeatedBar) > 0 {
		for _, msg := range m.RepeatedBar {
			dAtA[i] = 0x22
			i++
			i = encodeVarintPooltrue(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.MapBar) > 0 {
		for k, _ := range m.MapBar {
			dAtA[i] = 0x2a
			i++
			v := m.MapBar[k]
			msgSize := 0
			if v != nil {
				msgSize = v.ProtoSize()
				msgSize += 1 + sovPooltrue(uint64(msgSize))
			}
			mapSize := 1 + sovPooltrue(uint64(k)) + msgSize
			i = encodeVarintPooltrue(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintPooltrue(dAtA, i, uint64(k))
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintPooltrue(dAtA, i, uint64(v.ProtoSize()))
				n2, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n2
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Bar) Marshal() (dAtA []byte, err error) {
	m.checkNotRecycled()
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bar) MarshalPool() (*github_com_gogo_protobuf_mem.Bytes, error) {
	m.checkNotRecycled()
	size := m.ProtoSize()
	bytes := github_com_gogo_protobuf_mem.GetBytes(size)
	bytes.Truncate(size)
	n, err := m.MarshalTo(bytes.Value())
	if err != nil {
		bytes.Recycle()
		return nil, err
	}
	bytes.Truncate(n)
	return bytes, nil
}

func (m *Bar) MarshalTo(dAtA []byte) (int, error) {
	m.checkNotRecycled()
	var i int
	_ = i
	var l int
	_ = l
	if m.OneBar != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPooltrue(dAtA, i, uint64(m.OneBar))
	}
	if m.TwoBar != 0 {
		dAtA[i] = 0x15
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.TwoBar))
		i += 4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPooltrue(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

// GetFoo gets a reset *Foo.
func GetFoo() *Foo {
	if !mem.PoolingEnabled() {
		return &Foo{}
	}
	value := globalFooObjectPool.Get().(*Foo)
	value.poolMarker = mem.PoolMarkerAllocatedByPool
	return value
}

// GetBar gets a reset *Bar.
func GetBar() *Bar {
	if !mem.PoolingEnabled() {
		return &Bar{}
	}
	value := globalBarObjectPool.Get().(*Bar)
	value.poolMarker = mem.PoolMarkerAllocatedByPool
	return value
}

// Recycle puts the message back in the Pool that created it.
//
// Any non-nil fields that are messages will be recycled as well, including elements of repeated fields and values of map fields.
// Once Recycle is called, the message should no longer be used.
//
// If the message is nil, the message was not allocated from a Pool, or pooling is disabled, this is a no-op.
func (m *Foo) Recycle() {
	if !mem.PoolingEnabled() {
		return
	}
	if m == nil {
		return
	}
	if m.poolMarker&mem.PoolMarkerAllocatedByPool != mem.PoolMarkerAllocatedByPool {
		return
	}
	if m.poolMarker&mem.PoolMarkerRecycled == mem.PoolMarkerRecycled {
		panic(mem.PanicDoubleRecycle)
	}
	m.Bar.Recycle()
	if m.RepeatedBar != nil {
		for _, value := range m.RepeatedBar {
			value.Recycle()
		}
		m.RepeatedBar = m.RepeatedBar[:0]
	}
	if m.MapBar != nil {
		for _, elem := range m.MapBar {
			elem.Recycle()
		}
		m.MapBar = nil
	}
	m.poolMarker = mem.PoolMarkerRecycled
	globalFooObjectPool.Put(m)
}

// Recycle puts the message back in the Pool that created it.
//
// Any non-nil fields that are messages will be recycled as well, including elements of repeated fields and values of map fields.
// Once Recycle is called, the message should no longer be used.
//
// If the message is nil, the message was not allocated from a Pool, or pooling is disabled, this is a no-op.
func (m *Bar) Recycle() {
	if !mem.PoolingEnabled() {
		return
	}
	if m == nil {
		return
	}
	if m.poolMarker&mem.PoolMarkerAllocatedByPool != mem.PoolMarkerAllocatedByPool {
		return
	}
	if m.poolMarker&mem.PoolMarkerRecycled == mem.PoolMarkerRecycled {
		panic(mem.PanicDoubleRecycle)
	}
	m.poolMarker = mem.PoolMarkerRecycled
	globalBarObjectPool.Put(m)
}

// checkNotRecycled checks that the message has not been recycled, and if it has, panics.
//
// This is used by generated functions.
func (m *Foo) checkNotRecycled() {
	if m == nil {
		return
	}
	if m.poolMarker&mem.PoolMarkerRecycled == mem.PoolMarkerRecycled {
		panic(mem.PanicUseAfterRecycle)
	}
}

// checkNotRecycled checks that the message has not been recycled, and if it has, panics.
//
// This is used by generated functions.
func (m *Bar) checkNotRecycled() {
	if m == nil {
		return
	}
	if m.poolMarker&mem.PoolMarkerRecycled == mem.PoolMarkerRecycled {
		panic(mem.PanicUseAfterRecycle)
	}
}

func newFoo() mem.Object {
	return &Foo{}
}

func newBar() mem.Object {
	return &Bar{}
}

var (
	globalFooObjectPool = mem.NewObjectPool(newFoo)
	globalBarObjectPool = mem.NewObjectPool(newBar)
)

func init() {
	mem.RegisterGlobalObjectPool(
		func() *mem.ObjectPool { return globalFooObjectPool },
		func(options ...mem.ObjectPoolOption) { globalFooObjectPool = mem.NewObjectPool(newFoo, options...) },
	)
	mem.RegisterGlobalObjectPool(
		func() *mem.ObjectPool { return globalBarObjectPool },
		func(options ...mem.ObjectPoolOption) { globalBarObjectPool = mem.NewObjectPool(newBar, options...) },
	)
}

func (m *Foo) ProtoSize() (n int) {
	m.checkNotRecycled()
	var l int
	_ = l
	if m.One != 0 {
		n += 1 + sovPooltrue(uint64(m.One))
	}
	if m.Two != 0 {
		n += 5
	}
	if m.Bar != nil {
		l = m.Bar.ProtoSize()
		n += 1 + l + sovPooltrue(uint64(l))
	}
	if len(m.RepeatedBar) > 0 {
		for _, e := range m.RepeatedBar {
			l = e.ProtoSize()
			n += 1 + l + sovPooltrue(uint64(l))
		}
	}
	if len(m.MapBar) > 0 {
		for k, v := range m.MapBar {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.ProtoSize()
				l += 1 + sovPooltrue(uint64(l))
			}
			mapEntrySize := 1 + sovPooltrue(uint64(k)) + l
			n += mapEntrySize + 1 + sovPooltrue(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Bar) ProtoSize() (n int) {
	m.checkNotRecycled()
	var l int
	_ = l
	if m.OneBar != 0 {
		n += 1 + sovPooltrue(uint64(m.OneBar))
	}
	if m.TwoBar != 0 {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPooltrue(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPooltrue(x uint64) (n int) {
	return sovPooltrue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Foo) Unmarshal(dAtA []byte) error {
	m.checkNotRecycled()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPooltrue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Foo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Foo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field One", wireType)
			}
			m.One = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPooltrue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.One |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Two", wireType)
			}
			m.Two = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Two = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bar", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPooltrue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPooltrue
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bar == nil {
				m.Bar = GetBar()
			}
			if err := m.Bar.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepeatedBar", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPooltrue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPooltrue
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RepeatedBar = append(m.RepeatedBar, GetBar())
			if err := m.RepeatedBar[len(m.RepeatedBar)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MapBar", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPooltrue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPooltrue
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MapBar == nil {
				m.MapBar = make(map[uint32]*Bar)
			}
			var mapkey uint32
			var mapvalue *Bar
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPooltrue
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPooltrue
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPooltrue
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthPooltrue
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthPooltrue
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue := GetBar()
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPooltrue(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthPooltrue
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MapBar[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPooltrue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPooltrue
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
func (m *Bar) Unmarshal(dAtA []byte) error {
	m.checkNotRecycled()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPooltrue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Bar: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bar: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneBar", wireType)
			}
			m.OneBar = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPooltrue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OneBar |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field TwoBar", wireType)
			}
			m.TwoBar = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.TwoBar = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		default:
			iNdEx = preIndex
			skippy, err := skipPooltrue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPooltrue
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
func skipPooltrue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPooltrue
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
					return 0, ErrIntOverflowPooltrue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPooltrue
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPooltrue
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPooltrue
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPooltrue(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPooltrue = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPooltrue   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pooltrue/pooltrue.proto", fileDescriptor_pooltrue_fbd4e022f4ee01e5) }

var fileDescriptor_pooltrue_fbd4e022f4ee01e5 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4e, 0x83, 0x40,
	0x14, 0x86, 0x33, 0x9d, 0x96, 0x9a, 0xa9, 0x24, 0xcd, 0x6c, 0x8a, 0x2c, 0x90, 0xd4, 0x98, 0xb0,
	0x50, 0x30, 0xb8, 0xd0, 0xb8, 0x24, 0xb1, 0xe9, 0xc6, 0x0d, 0x17, 0x30, 0x83, 0x8e, 0x68, 0x04,
	0x1e, 0x99, 0x0c, 0x92, 0x7a, 0x1a, 0x8f, 0xe3, 0xb2, 0x67, 0x68, 0x6f, 0xe1, 0xca, 0xcc, 0x43,
	0xd4, 0xa6, 0xbb, 0xff, 0x9f, 0xff, 0x7f, 0x79, 0xdf, 0x3c, 0x36, 0xab, 0x01, 0x0a, 0xad, 0x1a,
	0x19, 0xf5, 0x22, 0xac, 0x15, 0x68, 0xe0, 0x07, 0xbd, 0x77, 0xcf, 0xf3, 0x17, 0xfd, 0xdc, 0x64,
	0xe1, 0x03, 0x94, 0x51, 0x0e, 0x39, 0x44, 0x58, 0xc8, 0x9a, 0x27, 0x74, 0x68, 0x50, 0x75, 0x83,
	0xf3, 0x2f, 0xc2, 0xe8, 0x02, 0x80, 0x4f, 0x19, 0x85, 0x4a, 0x3a, 0xc4, 0x27, 0xc1, 0x30, 0x35,
	0xd2, 0xbc, 0xe8, 0x16, 0x9c, 0x81, 0x4f, 0x82, 0x71, 0x6a, 0x24, 0x3f, 0x66, 0x34, 0x13, 0xca,
	0xa1, 0x3e, 0x09, 0x26, 0xb1, 0x1d, 0xfe, 0x22, 0x24, 0x42, 0xa5, 0x26, 0xe1, 0x17, 0xec, 0x50,
	0xc9, 0x5a, 0x0a, 0x2d, 0x1f, 0xef, 0x4d, 0x73, 0xe8, 0xd3, 0xfd, 0xe6, 0xa4, 0xaf, 0x24, 0x42,
	0xf1, 0x98, 0x8d, 0x4b, 0x51, 0x63, 0x79, 0x84, 0xe5, 0xa3, 0xbf, 0xf2, 0x02, 0x20, 0xbc, 0x13,
	0x75, 0x22, 0xd4, 0x6d, 0xa5, 0xd5, 0x2a, 0xb5, 0x4a, 0x34, 0xee, 0x92, 0x4d, 0xfe, 0x3d, 0x1b,
	0xce, 0x57, 0xb9, 0x42, 0x72, 0x3b, 0x35, 0x92, 0x9f, 0xb0, 0xd1, 0x9b, 0x28, 0x1a, 0x89, 0xec,
	0x7b, 0xfb, 0xbb, 0xec, 0x66, 0x70, 0x4d, 0xe6, 0x57, 0x8c, 0x1a, 0x88, 0x19, 0x1b, 0x43, 0x25,
	0x11, 0xa2, 0xfb, 0xbf, 0x05, 0x95, 0xfc, 0x09, 0x74, 0x0b, 0x18, 0x74, 0x67, 0xb0, 0x74, 0x0b,
	0x89, 0x50, 0xf1, 0x99, 0x19, 0x7c, 0xe7, 0xa7, 0x6c, 0xb4, 0x94, 0x45, 0x01, 0xdc, 0xde, 0xa1,
	0x76, 0x77, 0x37, 0x26, 0xd3, 0xcf, 0x8d, 0x47, 0xd6, 0x1b, 0x8f, 0x7c, 0x6c, 0x3d, 0xb2, 0xde,
	0x7a, 0x24, 0xb3, 0xf0, 0xf8, 0x97, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x26, 0xaf, 0xc9,
	0xd0, 0x01, 0x00, 0x00,
}
