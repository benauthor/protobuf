// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pooltrue/pooltrue.proto

package pooltrue

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import mem "github.com/gogo/protobuf/mem"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"
import github_com_gogo_protobuf_mem "github.com/gogo/protobuf/mem"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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

	pool *mem.Pool
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
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_pooltrue_0cdd686dcfe0c146, []int{0}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
	xxx_messageInfo_Foo.Merge(dst, src)
}
func (m *Foo) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetOne() uint64 {
	if m != nil {
		return m.One
	}
	return 0
}

func (m *Foo) GetTwo() uint32 {
	if m != nil {
		return m.Two
	}
	return 0
}

func (m *Foo) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

func (m *Foo) GetRepeatedBar() []*Bar {
	if m != nil {
		return m.RepeatedBar
	}
	return nil
}

func (m *Foo) GetMapBar() map[uint32]*Bar {
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

	pool *mem.Pool
}

func (m *Bar) Reset() {
	m.OneBar = 0
	m.TwoBar = 0
}
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_pooltrue_0cdd686dcfe0c146, []int{1}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
	xxx_messageInfo_Bar.Merge(dst, src)
}
func (m *Bar) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetOneBar() uint64 {
	if m != nil {
		return m.OneBar
	}
	return 0
}

func (m *Bar) GetTwoBar() uint32 {
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
func (m *Foo) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Foo) MarshalPool() (*github_com_gogo_protobuf_mem.Bytes, error) {
	size := m.ProtoSize()
	var bytes *github_com_gogo_protobuf_mem.Bytes
	if m.pool != nil {
		bytes = m.pool.GetBytes(size)
	} else {
		bytes = github_com_gogo_protobuf_mem.NewUnmanagedBytes(size)
	}
	n, err := m.MarshalTo(bytes.Value())
	if err != nil {
		bytes.Recycle()
		return nil, err
	}
	bytes.Truncate(n)
	return bytes, nil
}

func (m *Foo) MarshalTo(dAtA []byte) (int, error) {
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
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bar) MarshalPool() (*github_com_gogo_protobuf_mem.Bytes, error) {
	size := m.ProtoSize()
	var bytes *github_com_gogo_protobuf_mem.Bytes
	if m.pool != nil {
		bytes = m.pool.GetBytes(size)
	} else {
		bytes = github_com_gogo_protobuf_mem.NewUnmanagedBytes(size)
	}
	n, err := m.MarshalTo(bytes.Value())
	if err != nil {
		bytes.Recycle()
		return nil, err
	}
	bytes.Truncate(n)
	return bytes, nil
}

func (m *Bar) MarshalTo(dAtA []byte) (int, error) {
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

// RegisterToPoolPooltrue registers constructors for all messages in this file to the given Pool.
//
// Users must call this function at initialization for the Pool to pool messages in this file.
func RegisterToPoolPooltrue(pool *github_com_gogo_protobuf_mem.Pool) {
	pool.RegisterConstructor("pooltrue.Foo", registerToPoolFoo)
	pool.RegisterConstructor("pooltrue.Bar", registerToPoolBar)
}

// GetFoo gets a reset *Foo from the global Pool.
//
// It is generally preferable to create a Pool instance for your application and
// use GetFooFromPool instead, however this will prove difficult
// for existing applications wishing to migrate to the Pool API. Additionally,
// RegisterToPoolPooltrueis implicitly called on the global Pool so there
// is no need for additional setup to use this function.
func GetFoo() *Foo {
	return GetFooFromPool(github_com_gogo_protobuf_mem.Global())
}

// GetFooFromPool gets a reset *Foo from the given Pool.
//
// If the Pool is nil, or PooltruePoolRegister was not called on this Pool, this will return a new message not managed by any Pool.
func GetFooFromPool(pool *github_com_gogo_protobuf_mem.Pool) *Foo {
	if pool == nil {
		return &Foo{}
	}
	pooledMessage := pool.Get("pooltrue.Foo")
	if pooledMessage == nil {
		return &Foo{}
	}
	return pooledMessage.(*Foo)
}

// GetBar gets a reset *Bar from the global Pool.
//
// It is generally preferable to create a Pool instance for your application and
// use GetBarFromPool instead, however this will prove difficult
// for existing applications wishing to migrate to the Pool API. Additionally,
// RegisterToPoolPooltrueis implicitly called on the global Pool so there
// is no need for additional setup to use this function.
func GetBar() *Bar {
	return GetBarFromPool(github_com_gogo_protobuf_mem.Global())
}

// GetBarFromPool gets a reset *Bar from the given Pool.
//
// If the Pool is nil, or PooltruePoolRegister was not called on this Pool, this will return a new message not managed by any Pool.
func GetBarFromPool(pool *github_com_gogo_protobuf_mem.Pool) *Bar {
	if pool == nil {
		return &Bar{}
	}
	pooledMessage := pool.Get("pooltrue.Bar")
	if pooledMessage == nil {
		return &Bar{}
	}
	return pooledMessage.(*Bar)
}

// Recycle puts the message back in the Pool that created it.
//
// Any non-nil fields that are messages will be recycled as well, including elements of repeated fields and values of map fields.
//
// If the message is nil, or the message was not allocated from a Pool, this is a no-op.
func (this *Foo) Recycle() {
	if this == nil {
		return
	}
	var iface interface{} = this.Bar
	if pooledMessage, ok := iface.(github_com_gogo_protobuf_mem.PooledMessage); ok {
		pooledMessage.Recycle()
	}
	for _, elem := range this.RepeatedBar {
		var iface interface{} = elem
		if pooledMessage, ok := iface.(github_com_gogo_protobuf_mem.PooledMessage); ok {
			pooledMessage.Recycle()
		}
	}
	for _, elem := range this.MapBar {
		var iface interface{} = elem
		if pooledMessage, ok := iface.(github_com_gogo_protobuf_mem.PooledMessage); ok {
			pooledMessage.Recycle()
		}
	}
	if this.pool == nil {
		return
	}
	this.pool.Put("pooltrue.Foo", this)
}

// Recycle puts the message back in the Pool that created it.
//
// Any non-nil fields that are messages will be recycled as well, including elements of repeated fields and values of map fields.
//
// If the message is nil, or the message was not allocated from a Pool, this is a no-op.
func (this *Bar) Recycle() {
	if this == nil {
		return
	}
	if this.pool == nil {
		return
	}
	this.pool.Put("pooltrue.Bar", this)
}

func registerToPoolFoo(pool *github_com_gogo_protobuf_mem.Pool) github_com_gogo_protobuf_mem.PooledMessage {
	return &Foo{pool: pool}
}

func registerToPoolBar(pool *github_com_gogo_protobuf_mem.Pool) github_com_gogo_protobuf_mem.PooledMessage {
	return &Bar{pool: pool}
}

func init() {
	RegisterToPoolPooltrue(github_com_gogo_protobuf_mem.Global())
}

func (m *Foo) ProtoSize() (n int) {
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
				if m.pool == nil {
					m.Bar = &Bar{}
				} else if pooledMessage := m.pool.Get("pooltrue.Bar"); pooledMessage != nil {
					m.Bar = pooledMessage.(*Bar)
				} else {
					m.Bar = &Bar{}
				}
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
			if m.pool == nil {
				m.RepeatedBar = append(m.RepeatedBar, &Bar{})
			} else if pooledMessage := m.pool.Get("pooltrue.Bar"); pooledMessage != nil {
				m.RepeatedBar = append(m.RepeatedBar, pooledMessage.(*Bar))
			} else {
				m.RepeatedBar = append(m.RepeatedBar, &Bar{})
			}
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
					if m.pool == nil {
						mapvalue = &Bar{}
					} else if pooledMessage := m.pool.Get("pooltrue.Bar"); pooledMessage != nil {
						mapvalue = pooledMessage.(*Bar)
					} else {
						mapvalue = &Bar{}
					}
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

func init() { proto.RegisterFile("pooltrue/pooltrue.proto", fileDescriptor_pooltrue_0cdd686dcfe0c146) }

var fileDescriptor_pooltrue_0cdd686dcfe0c146 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0xba, 0x4d, 0xd1, 0x85, 0x4a, 0x95, 0x97, 0x86, 0x0e, 0x21, 0x2a, 0x4b, 0x16,
	0x12, 0x54, 0x06, 0x10, 0x63, 0x24, 0x2a, 0x16, 0x96, 0xbc, 0x00, 0x72, 0xc0, 0x04, 0x44, 0x93,
	0xb3, 0x2c, 0x87, 0xa8, 0x6f, 0xc3, 0xe3, 0x30, 0xf6, 0x19, 0xda, 0xb7, 0x60, 0x42, 0xbe, 0x10,
	0x40, 0xea, 0xf6, 0xff, 0xf7, 0x7f, 0x96, 0xff, 0x3b, 0x98, 0x69, 0xc4, 0xb5, 0x35, 0x8d, 0x4a,
	0x7b, 0x91, 0x68, 0x83, 0x16, 0xc5, 0x51, 0xef, 0xe7, 0xe7, 0xe5, 0xab, 0x7d, 0x69, 0x8a, 0xe4,
	0x11, 0xab, 0xb4, 0xc4, 0x12, 0x53, 0x02, 0x8a, 0xe6, 0x99, 0x1c, 0x19, 0x52, 0xdd, 0xc3, 0xc5,
	0x17, 0x03, 0xbe, 0x42, 0x14, 0x53, 0xe0, 0x58, 0xab, 0x80, 0x45, 0x2c, 0x1e, 0xe6, 0x4e, 0xba,
	0x89, 0x6d, 0x31, 0x18, 0x44, 0x2c, 0x1e, 0xe7, 0x4e, 0x8a, 0x53, 0xe0, 0x85, 0x34, 0x01, 0x8f,
	0x58, 0xec, 0x2f, 0x27, 0xc9, 0x6f, 0x85, 0x4c, 0x9a, 0xdc, 0x25, 0xe2, 0x02, 0x8e, 0x8d, 0xd2,
	0x4a, 0x5a, 0xf5, 0xf4, 0xe0, 0xc8, 0x61, 0xc4, 0x0f, 0x49, 0xbf, 0x47, 0x32, 0x69, 0xc4, 0x12,
	0xc6, 0x95, 0xd4, 0x04, 0x8f, 0x08, 0x3e, 0xf9, 0x83, 0x57, 0x88, 0xc9, 0xbd, 0xd4, 0x99, 0x34,
	0xb7, 0xb5, 0x35, 0x9b, 0xdc, 0xab, 0xc8, 0xcc, 0xef, 0xc0, 0xff, 0x37, 0x76, 0x3d, 0xdf, 0xd4,
	0x86, 0x9a, 0x4f, 0x72, 0x27, 0xc5, 0x19, 0x8c, 0xde, 0xe5, 0xba, 0x51, 0xd4, 0xfd, 0xe0, 0xff,
	0x2e, 0xbb, 0x19, 0x5c, 0xb3, 0xc5, 0x15, 0x70, 0x57, 0x62, 0x06, 0x63, 0xac, 0x15, 0x95, 0xe8,
	0xf6, 0xf7, 0xb0, 0x56, 0x3f, 0x81, 0x6d, 0x91, 0x82, 0xee, 0x0c, 0x9e, 0x6d, 0x31, 0x93, 0x26,
	0x9b, 0x7e, 0xee, 0x42, 0xb6, 0xdd, 0x85, 0xec, 0x63, 0x1f, 0xb2, 0xed, 0x3e, 0x64, 0x85, 0x47,
	0xe7, 0xbc, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xea, 0xa6, 0x12, 0xa2, 0x01, 0x00, 0x00,
}
