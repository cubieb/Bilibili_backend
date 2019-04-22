// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Comparator.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BitComparator_BitwiseOp int32

const (
	BitComparator_AND BitComparator_BitwiseOp = 1
	BitComparator_OR  BitComparator_BitwiseOp = 2
	BitComparator_XOR BitComparator_BitwiseOp = 3
)

var BitComparator_BitwiseOp_name = map[int32]string{
	1: "AND",
	2: "OR",
	3: "XOR",
}
var BitComparator_BitwiseOp_value = map[string]int32{
	"AND": 1,
	"OR":  2,
	"XOR": 3,
}

func (x BitComparator_BitwiseOp) Enum() *BitComparator_BitwiseOp {
	p := new(BitComparator_BitwiseOp)
	*p = x
	return p
}
func (x BitComparator_BitwiseOp) String() string {
	return proto.EnumName(BitComparator_BitwiseOp_name, int32(x))
}
func (x *BitComparator_BitwiseOp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BitComparator_BitwiseOp_value, data, "BitComparator_BitwiseOp")
	if err != nil {
		return err
	}
	*x = BitComparator_BitwiseOp(value)
	return nil
}
func (BitComparator_BitwiseOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{5, 0} }

type Comparator struct {
	Name                 *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	SerializedComparator []byte  `protobuf:"bytes,2,opt,name=serialized_comparator,json=serializedComparator" json:"serialized_comparator,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *Comparator) Reset()                    { *m = Comparator{} }
func (m *Comparator) String() string            { return proto.CompactTextString(m) }
func (*Comparator) ProtoMessage()               {}
func (*Comparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Comparator) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Comparator) GetSerializedComparator() []byte {
	if m != nil {
		return m.SerializedComparator
	}
	return nil
}

type ByteArrayComparable struct {
	Value            []byte `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ByteArrayComparable) Reset()                    { *m = ByteArrayComparable{} }
func (m *ByteArrayComparable) String() string            { return proto.CompactTextString(m) }
func (*ByteArrayComparable) ProtoMessage()               {}
func (*ByteArrayComparable) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *ByteArrayComparable) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BinaryComparator struct {
	Comparable       *ByteArrayComparable `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *BinaryComparator) Reset()                    { *m = BinaryComparator{} }
func (m *BinaryComparator) String() string            { return proto.CompactTextString(m) }
func (*BinaryComparator) ProtoMessage()               {}
func (*BinaryComparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *BinaryComparator) GetComparable() *ByteArrayComparable {
	if m != nil {
		return m.Comparable
	}
	return nil
}

type LongComparator struct {
	Comparable       *ByteArrayComparable `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *LongComparator) Reset()                    { *m = LongComparator{} }
func (m *LongComparator) String() string            { return proto.CompactTextString(m) }
func (*LongComparator) ProtoMessage()               {}
func (*LongComparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *LongComparator) GetComparable() *ByteArrayComparable {
	if m != nil {
		return m.Comparable
	}
	return nil
}

type BinaryPrefixComparator struct {
	Comparable       *ByteArrayComparable `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *BinaryPrefixComparator) Reset()                    { *m = BinaryPrefixComparator{} }
func (m *BinaryPrefixComparator) String() string            { return proto.CompactTextString(m) }
func (*BinaryPrefixComparator) ProtoMessage()               {}
func (*BinaryPrefixComparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *BinaryPrefixComparator) GetComparable() *ByteArrayComparable {
	if m != nil {
		return m.Comparable
	}
	return nil
}

type BitComparator struct {
	Comparable       *ByteArrayComparable     `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
	BitwiseOp        *BitComparator_BitwiseOp `protobuf:"varint,2,req,name=bitwise_op,json=bitwiseOp,enum=pb.BitComparator_BitwiseOp" json:"bitwise_op,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *BitComparator) Reset()                    { *m = BitComparator{} }
func (m *BitComparator) String() string            { return proto.CompactTextString(m) }
func (*BitComparator) ProtoMessage()               {}
func (*BitComparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *BitComparator) GetComparable() *ByteArrayComparable {
	if m != nil {
		return m.Comparable
	}
	return nil
}

func (m *BitComparator) GetBitwiseOp() BitComparator_BitwiseOp {
	if m != nil && m.BitwiseOp != nil {
		return *m.BitwiseOp
	}
	return BitComparator_AND
}

type NullComparator struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *NullComparator) Reset()                    { *m = NullComparator{} }
func (m *NullComparator) String() string            { return proto.CompactTextString(m) }
func (*NullComparator) ProtoMessage()               {}
func (*NullComparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

type RegexStringComparator struct {
	Pattern          *string `protobuf:"bytes,1,req,name=pattern" json:"pattern,omitempty"`
	PatternFlags     *int32  `protobuf:"varint,2,req,name=pattern_flags,json=patternFlags" json:"pattern_flags,omitempty"`
	Charset          *string `protobuf:"bytes,3,req,name=charset" json:"charset,omitempty"`
	Engine           *string `protobuf:"bytes,4,opt,name=engine" json:"engine,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RegexStringComparator) Reset()                    { *m = RegexStringComparator{} }
func (m *RegexStringComparator) String() string            { return proto.CompactTextString(m) }
func (*RegexStringComparator) ProtoMessage()               {}
func (*RegexStringComparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *RegexStringComparator) GetPattern() string {
	if m != nil && m.Pattern != nil {
		return *m.Pattern
	}
	return ""
}

func (m *RegexStringComparator) GetPatternFlags() int32 {
	if m != nil && m.PatternFlags != nil {
		return *m.PatternFlags
	}
	return 0
}

func (m *RegexStringComparator) GetCharset() string {
	if m != nil && m.Charset != nil {
		return *m.Charset
	}
	return ""
}

func (m *RegexStringComparator) GetEngine() string {
	if m != nil && m.Engine != nil {
		return *m.Engine
	}
	return ""
}

type SubstringComparator struct {
	Substr           *string `protobuf:"bytes,1,req,name=substr" json:"substr,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SubstringComparator) Reset()                    { *m = SubstringComparator{} }
func (m *SubstringComparator) String() string            { return proto.CompactTextString(m) }
func (*SubstringComparator) ProtoMessage()               {}
func (*SubstringComparator) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *SubstringComparator) GetSubstr() string {
	if m != nil && m.Substr != nil {
		return *m.Substr
	}
	return ""
}

func init() {
	proto.RegisterType((*Comparator)(nil), "pb.Comparator")
	proto.RegisterType((*ByteArrayComparable)(nil), "pb.ByteArrayComparable")
	proto.RegisterType((*BinaryComparator)(nil), "pb.BinaryComparator")
	proto.RegisterType((*LongComparator)(nil), "pb.LongComparator")
	proto.RegisterType((*BinaryPrefixComparator)(nil), "pb.BinaryPrefixComparator")
	proto.RegisterType((*BitComparator)(nil), "pb.BitComparator")
	proto.RegisterType((*NullComparator)(nil), "pb.NullComparator")
	proto.RegisterType((*RegexStringComparator)(nil), "pb.RegexStringComparator")
	proto.RegisterType((*SubstringComparator)(nil), "pb.SubstringComparator")
	proto.RegisterEnum("pb.BitComparator_BitwiseOp", BitComparator_BitwiseOp_name, BitComparator_BitwiseOp_value)
}

func init() { proto.RegisterFile("Comparator.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x95, 0x37, 0x6d, 0xaa, 0x0c, 0x6d, 0xb4, 0x72, 0xdb, 0xb0, 0x12, 0x97, 0x68, 0x11, 0x52,
	0x04, 0x62, 0x0f, 0xe5, 0x80, 0xc4, 0xad, 0x0b, 0xaa, 0x40, 0xa0, 0xa6, 0xb8, 0x42, 0xe2, 0x16,
	0xd9, 0xc9, 0x64, 0x63, 0x69, 0xbb, 0xb6, 0x6c, 0x2f, 0x34, 0x7c, 0x41, 0x3f, 0x83, 0x3b, 0x3f,
	0x89, 0xbc, 0xeb, 0x24, 0x0b, 0xea, 0x31, 0xb7, 0x79, 0x6f, 0xe6, 0x3d, 0x3f, 0xd9, 0x1e, 0x88,
	0xdf, 0xab, 0x3b, 0xcd, 0x0d, 0x77, 0xca, 0x64, 0xda, 0x28, 0xa7, 0x68, 0xa4, 0x45, 0xfa, 0x0d,
	0x60, 0xc7, 0x53, 0x0a, 0x07, 0x15, 0xbf, 0xc3, 0x84, 0x8c, 0xa3, 0xc9, 0x80, 0x35, 0x35, 0x7d,
	0x03, 0xe7, 0x16, 0x8d, 0xe4, 0xa5, 0xfc, 0x85, 0x8b, 0xd9, 0x7c, 0x3b, 0x9c, 0x44, 0x63, 0x32,
	0x39, 0x66, 0x67, 0xbb, 0xe6, 0xce, 0x28, 0x7d, 0x05, 0xa7, 0xf9, 0xda, 0xe1, 0xa5, 0x31, 0x7c,
	0x1d, 0x68, 0x51, 0x22, 0x3d, 0x83, 0xc3, 0x1f, 0xbc, 0xac, 0xfd, 0x01, 0x5e, 0xdb, 0x82, 0xf4,
	0x33, 0xc4, 0xb9, 0xac, 0xb8, 0x59, 0x77, 0x92, 0xbc, 0x05, 0x98, 0x6f, 0x75, 0x4d, 0x9e, 0x27,
	0x17, 0x4f, 0x33, 0x2d, 0xb2, 0x47, 0x6c, 0x59, 0x67, 0x34, 0xfd, 0x04, 0xc3, 0x2f, 0xaa, 0x2a,
	0xf6, 0x61, 0xf5, 0x15, 0x46, 0x6d, 0xae, 0x1b, 0x83, 0x4b, 0x79, 0xbf, 0x0f, 0xcb, 0x3f, 0x04,
	0x4e, 0x72, 0xe9, 0xf6, 0x60, 0x45, 0xdf, 0x01, 0x08, 0xe9, 0x7e, 0x4a, 0x8b, 0x33, 0xa5, 0x93,
	0x68, 0x1c, 0x4d, 0x86, 0x17, 0xcf, 0x1a, 0x61, 0xd7, 0xdf, 0x23, 0x3f, 0x33, 0xd5, 0x6c, 0x20,
	0x36, 0x65, 0xfa, 0x02, 0x06, 0x5b, 0x9e, 0x1e, 0x41, 0xef, 0xf2, 0xfa, 0x43, 0x4c, 0x68, 0x1f,
	0xa2, 0x29, 0x8b, 0x23, 0x4f, 0x7c, 0x9f, 0xb2, 0xb8, 0x97, 0xc6, 0x30, 0xbc, 0xae, 0xcb, 0xb2,
	0xf3, 0xae, 0x0f, 0x04, 0xce, 0x19, 0x16, 0x78, 0x7f, 0xeb, 0x8c, 0xfc, 0xe7, 0x96, 0x13, 0x38,
	0xd2, 0xdc, 0x39, 0x34, 0x55, 0xf8, 0x3d, 0x1b, 0x48, 0x9f, 0xc3, 0x49, 0x28, 0x67, 0xcb, 0x92,
	0x17, 0xb6, 0xc9, 0x7a, 0xc8, 0x8e, 0x03, 0x79, 0xe5, 0x39, 0x2f, 0x9f, 0xaf, 0xb8, 0xb1, 0xe8,
	0x92, 0x5e, 0x2b, 0x0f, 0x90, 0x8e, 0xa0, 0x8f, 0x55, 0x21, 0x2b, 0x4c, 0x0e, 0xc6, 0x64, 0x32,
	0x60, 0x01, 0xa5, 0xaf, 0xe1, 0xf4, 0xb6, 0x16, 0xf6, 0xff, 0x1c, 0x23, 0xe8, 0xdb, 0x86, 0x0e,
	0x31, 0x02, 0xca, 0xaf, 0xe0, 0xa5, 0x32, 0x45, 0xc6, 0x35, 0x9f, 0xaf, 0x30, 0x5b, 0xf1, 0x85,
	0x52, 0x3a, 0x5b, 0x09, 0x6e, 0xb1, 0xdd, 0x06, 0x51, 0x2f, 0xb3, 0x02, 0x2b, 0x34, 0xdc, 0xe1,
	0x22, 0xef, 0x2c, 0xcb, 0x8d, 0xef, 0xda, 0x8f, 0xe4, 0x81, 0x90, 0xdf, 0x84, 0xfc, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xff, 0x76, 0x5f, 0x6f, 0x47, 0x03, 0x00, 0x00,
}
