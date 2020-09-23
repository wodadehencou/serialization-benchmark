// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

package gogoprotobuf

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Simple struct {
	StrField    string `protobuf:"bytes,1,opt,name=StrField,proto3" json:"StrField,omitempty"`
	Bytes1Field []byte `protobuf:"bytes,2,opt,name=Bytes1Field,proto3" json:"Bytes1Field,omitempty"`
	Bytes2Field []byte `protobuf:"bytes,3,opt,name=Bytes2Field,proto3" json:"Bytes2Field,omitempty"`
	Timestamp   int64  `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (m *Simple) Reset()      { *m = Simple{} }
func (*Simple) ProtoMessage() {}
func (*Simple) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{0}
}
func (m *Simple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Simple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Simple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Simple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple.Merge(m, src)
}
func (m *Simple) XXX_Size() int {
	return m.Size()
}
func (m *Simple) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple.DiscardUnknown(m)
}

var xxx_messageInfo_Simple proto.InternalMessageInfo

func (m *Simple) GetStrField() string {
	if m != nil {
		return m.StrField
	}
	return ""
}

func (m *Simple) GetBytes1Field() []byte {
	if m != nil {
		return m.Bytes1Field
	}
	return nil
}

func (m *Simple) GetBytes2Field() []byte {
	if m != nil {
		return m.Bytes2Field
	}
	return nil
}

func (m *Simple) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type SimpleArr struct {
	Simples []*Simple `protobuf:"bytes,1,rep,name=Simples,proto3" json:"Simples,omitempty"`
}

func (m *SimpleArr) Reset()      { *m = SimpleArr{} }
func (*SimpleArr) ProtoMessage() {}
func (*SimpleArr) Descriptor() ([]byte, []int) {
	return fileDescriptor_f80abaa17e25ccc8, []int{1}
}
func (m *SimpleArr) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleArr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleArr.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleArr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleArr.Merge(m, src)
}
func (m *SimpleArr) XXX_Size() int {
	return m.Size()
}
func (m *SimpleArr) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleArr.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleArr proto.InternalMessageInfo

func (m *SimpleArr) GetSimples() []*Simple {
	if m != nil {
		return m.Simples
	}
	return nil
}

func init() {
	proto.RegisterType((*Simple)(nil), "pb.Simple")
	proto.RegisterType((*SimpleArr)(nil), "pb.SimpleArr")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_f80abaa17e25ccc8) }

var fileDescriptor_f80abaa17e25ccc8 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0xfd, 0x5a, 0x54, 0x5a, 0x97, 0x29, 0x53, 0x85, 0xd0, 0x53, 0x54, 0x31, 0x64, 0x21,
	0x55, 0xcb, 0xca, 0x42, 0x07, 0x3e, 0x20, 0x65, 0x62, 0x8b, 0x13, 0x93, 0x58, 0x24, 0xb1, 0xe5,
	0x38, 0x42, 0x30, 0x31, 0x31, 0xf3, 0x19, 0x7c, 0x0a, 0x63, 0xc6, 0x8e, 0xc4, 0x59, 0x18, 0xfb,
	0x09, 0x48, 0xb5, 0x28, 0xdd, 0xde, 0x3d, 0xf7, 0x0c, 0x4f, 0x97, 0x8e, 0x15, 0x0b, 0x95, 0x96,
	0x46, 0x7a, 0x03, 0xc5, 0xe6, 0xef, 0x40, 0x47, 0x1b, 0x51, 0xaa, 0x82, 0x7b, 0xe7, 0x74, 0xbc,
	0x31, 0xfa, 0x4e, 0xf0, 0x22, 0x9d, 0x81, 0x0f, 0xc1, 0x24, 0x3a, 0x64, 0xcf, 0xa7, 0xd3, 0xf5,
	0x8b, 0xe1, 0xf5, 0xd2, 0xd5, 0x03, 0x1f, 0x82, 0xb3, 0xe8, 0x18, 0x1d, 0x8c, 0x95, 0x33, 0x86,
	0x47, 0x86, 0x43, 0xde, 0x05, 0x9d, 0xdc, 0x8b, 0x92, 0xd7, 0x26, 0x2e, 0xd5, 0xec, 0xc4, 0x87,
	0x60, 0x18, 0xfd, 0x83, 0xf9, 0x92, 0x4e, 0xdc, 0x1f, 0xb7, 0x5a, 0x7b, 0x97, 0xf4, 0xd4, 0x85,
	0x7a, 0x06, 0xfe, 0x30, 0x98, 0xae, 0x68, 0xa8, 0x58, 0xe8, 0x50, 0xf4, 0x57, 0xad, 0x75, 0xdb,
	0x21, 0xd9, 0x76, 0x48, 0x76, 0x1d, 0xc2, 0x9b, 0x45, 0xf8, 0xb4, 0x08, 0x5f, 0x16, 0xa1, 0xb5,
	0x08, 0xdf, 0x16, 0xe1, 0xc7, 0x22, 0xd9, 0x59, 0x84, 0x8f, 0x1e, 0x49, 0xdb, 0x23, 0xd9, 0xf6,
	0x48, 0x1e, 0x6e, 0x32, 0x61, 0xf2, 0x86, 0x85, 0x89, 0x2c, 0x17, 0xcf, 0x32, 0x8d, 0x53, 0x9e,
	0xf3, 0x2a, 0x91, 0xcd, 0xa2, 0xe6, 0x5a, 0xc4, 0x85, 0x78, 0x8d, 0x8d, 0x90, 0xd5, 0x15, 0xe3,
	0x55, 0x92, 0x97, 0xb1, 0x7e, 0x5a, 0x64, 0x32, 0x93, 0xfb, 0xc1, 0x58, 0xf3, 0xc8, 0x46, 0xfb,
	0xeb, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x8a, 0x90, 0xbd, 0x46, 0x01, 0x00, 0x00,
}

func (this *Simple) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Simple)
	if !ok {
		that2, ok := that.(Simple)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.StrField != that1.StrField {
		return false
	}
	if !bytes.Equal(this.Bytes1Field, that1.Bytes1Field) {
		return false
	}
	if !bytes.Equal(this.Bytes2Field, that1.Bytes2Field) {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	return true
}
func (this *SimpleArr) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SimpleArr)
	if !ok {
		that2, ok := that.(SimpleArr)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Simples) != len(that1.Simples) {
		return false
	}
	for i := range this.Simples {
		if !this.Simples[i].Equal(that1.Simples[i]) {
			return false
		}
	}
	return true
}
func (this *Simple) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&gogoprotobuf.Simple{")
	s = append(s, "StrField: "+fmt.Sprintf("%#v", this.StrField)+",\n")
	s = append(s, "Bytes1Field: "+fmt.Sprintf("%#v", this.Bytes1Field)+",\n")
	s = append(s, "Bytes2Field: "+fmt.Sprintf("%#v", this.Bytes2Field)+",\n")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SimpleArr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&gogoprotobuf.SimpleArr{")
	if this.Simples != nil {
		s = append(s, "Simples: "+fmt.Sprintf("%#v", this.Simples)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPb(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Simple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Simple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Simple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintPb(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Bytes2Field) > 0 {
		i -= len(m.Bytes2Field)
		copy(dAtA[i:], m.Bytes2Field)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Bytes2Field)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bytes1Field) > 0 {
		i -= len(m.Bytes1Field)
		copy(dAtA[i:], m.Bytes1Field)
		i = encodeVarintPb(dAtA, i, uint64(len(m.Bytes1Field)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StrField) > 0 {
		i -= len(m.StrField)
		copy(dAtA[i:], m.StrField)
		i = encodeVarintPb(dAtA, i, uint64(len(m.StrField)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimpleArr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleArr) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleArr) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Simples) > 0 {
		for iNdEx := len(m.Simples) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Simples[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPb(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPb(dAtA []byte, offset int, v uint64) int {
	offset -= sovPb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Simple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StrField)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	l = len(m.Bytes1Field)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	l = len(m.Bytes2Field)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovPb(uint64(m.Timestamp))
	}
	return n
}

func (m *SimpleArr) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Simples) > 0 {
		for _, e := range m.Simples {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	return n
}

func sovPb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPb(x uint64) (n int) {
	return sovPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Simple) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Simple{`,
		`StrField:` + fmt.Sprintf("%v", this.StrField) + `,`,
		`Bytes1Field:` + fmt.Sprintf("%v", this.Bytes1Field) + `,`,
		`Bytes2Field:` + fmt.Sprintf("%v", this.Bytes2Field) + `,`,
		`Timestamp:` + fmt.Sprintf("%v", this.Timestamp) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SimpleArr) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForSimples := "[]*Simple{"
	for _, f := range this.Simples {
		repeatedStringForSimples += strings.Replace(f.String(), "Simple", "Simple", 1) + ","
	}
	repeatedStringForSimples += "}"
	s := strings.Join([]string{`&SimpleArr{`,
		`Simples:` + repeatedStringForSimples + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPb(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Simple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: Simple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Simple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
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
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StrField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes1Field", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes1Field = append(m.Bytes1Field[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes1Field == nil {
				m.Bytes1Field = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes2Field", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes2Field = append(m.Bytes2Field[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes2Field == nil {
				m.Bytes2Field = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SimpleArr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
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
			return fmt.Errorf("proto: SimpleArr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleArr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Simples", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Simples = append(m.Simples, &Simple{})
			if err := m.Simples[len(m.Simples)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPb
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
					return 0, ErrIntOverflowPb
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
					return 0, ErrIntOverflowPb
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
				return 0, ErrInvalidLengthPb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPb = fmt.Errorf("proto: unexpected end of group")
)
