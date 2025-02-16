// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kafka.proto

/*
	Package splunk_trace_firehose is a generated protocol buffer package.

	It is generated from these files:
		kafka.proto

	It has these top-level messages:
		Kafka
*/
package splunk_trace_firehose

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import jaeger_api_v2 "github.com/jaegertracing/jaeger/model"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Kafka struct {
	Span       *jaeger_api_v2.Span `protobuf:"bytes,1,opt,name=span" json:"span,omitempty"`
	OrgID      int64               `protobuf:"varint,2,opt,name=orgID,proto3" json:"orgID,omitempty"`
	TokenID    int64               `protobuf:"varint,3,opt,name=tokenID,proto3" json:"tokenID,omitempty"`
	IngestTime int64               `protobuf:"varint,4,opt,name=ingestTime,proto3" json:"ingestTime,omitempty"`
}

func (m *Kafka) Reset()                    { *m = Kafka{} }
func (m *Kafka) String() string            { return proto.CompactTextString(m) }
func (*Kafka) ProtoMessage()               {}
func (*Kafka) Descriptor() ([]byte, []int) { return fileDescriptorKafka, []int{0} }

func (m *Kafka) GetSpan() *jaeger_api_v2.Span {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *Kafka) GetOrgID() int64 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *Kafka) GetTokenID() int64 {
	if m != nil {
		return m.TokenID
	}
	return 0
}

func (m *Kafka) GetIngestTime() int64 {
	if m != nil {
		return m.IngestTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Kafka)(nil), "splunk.trace.firehose.Kafka")
}
func (m *Kafka) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Kafka) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Span != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKafka(dAtA, i, uint64(m.Span.Size()))
		n1, err := m.Span.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.OrgID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintKafka(dAtA, i, uint64(m.OrgID))
	}
	if m.TokenID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintKafka(dAtA, i, uint64(m.TokenID))
	}
	if m.IngestTime != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintKafka(dAtA, i, uint64(m.IngestTime))
	}
	return i, nil
}

func encodeFixed64Kafka(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Kafka(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintKafka(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Kafka) Size() (n int) {
	var l int
	_ = l
	if m.Span != nil {
		l = m.Span.Size()
		n += 1 + l + sovKafka(uint64(l))
	}
	if m.OrgID != 0 {
		n += 1 + sovKafka(uint64(m.OrgID))
	}
	if m.TokenID != 0 {
		n += 1 + sovKafka(uint64(m.TokenID))
	}
	if m.IngestTime != 0 {
		n += 1 + sovKafka(uint64(m.IngestTime))
	}
	return n
}

func sovKafka(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKafka(x uint64) (n int) {
	return sovKafka(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Kafka) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKafka
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
			return fmt.Errorf("proto: Kafka: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Kafka: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Span", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
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
				return ErrInvalidLengthKafka
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Span == nil {
				m.Span = &jaeger_api_v2.Span{}
			}
			if err := m.Span.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenID", wireType)
			}
			m.TokenID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IngestTime", wireType)
			}
			m.IngestTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKafka
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IngestTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKafka(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKafka
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
func skipKafka(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKafka
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
					return 0, ErrIntOverflowKafka
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
					return 0, ErrIntOverflowKafka
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
				return 0, ErrInvalidLengthKafka
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKafka
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
				next, err := skipKafka(dAtA[start:])
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
	ErrInvalidLengthKafka = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKafka   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("kafka.proto", fileDescriptorKafka) }

var fileDescriptorKafka = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4e, 0x4c, 0xcb,
	0x4e, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2d, 0x2e, 0xc8, 0x29, 0xcd, 0xcb, 0xd6,
	0x2b, 0x29, 0x4a, 0x4c, 0x4e, 0xd5, 0x4b, 0xcb, 0x2c, 0x4a, 0xcd, 0xc8, 0x2f, 0x4e, 0x95, 0x32,
	0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4a, 0x4c, 0x4d, 0x4f,
	0x2d, 0x02, 0x29, 0xc8, 0xcc, 0x4b, 0x87, 0xf2, 0xf4, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x20, 0x24,
	0xc4, 0x24, 0xa5, 0x06, 0x46, 0x2e, 0x56, 0x6f, 0x90, 0xc9, 0x42, 0xea, 0x5c, 0x2c, 0xc5, 0x05,
	0x89, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x10, 0x2d, 0x7a, 0x89, 0x05,
	0x99, 0xf1, 0x65, 0x46, 0x7a, 0xc1, 0x05, 0x89, 0x79, 0x41, 0x60, 0x05, 0x42, 0x22, 0x5c, 0xac,
	0xf9, 0x45, 0xe9, 0x9e, 0x2e, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x90, 0x04,
	0x17, 0x7b, 0x49, 0x7e, 0x76, 0x6a, 0x9e, 0xa7, 0x8b, 0x04, 0x33, 0x58, 0x1c, 0xc6, 0x15, 0x92,
	0xe3, 0xe2, 0xca, 0xcc, 0x4b, 0x4f, 0x2d, 0x2e, 0x09, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x01, 0x4b,
	0x22, 0x89, 0x38, 0x09, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72,
	0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0xdd, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0x7d, 0x88, 0xea, 0xf4, 0x00, 0x00, 0x00,
}
