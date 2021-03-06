// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imageapi-config.proto

package imageapi

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ApiServer_Status int32

const (
	ApiServer_UNKNOWN ApiServer_Status = 0
	ApiServer_OK      ApiServer_Status = 1
	ApiServer_ERROR   ApiServer_Status = 2
)

var ApiServer_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "ERROR",
}

var ApiServer_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"OK":      1,
	"ERROR":   2,
}

func (x ApiServer_Status) String() string {
	return proto.EnumName(ApiServer_Status_name, int32(x))
}

func (ApiServer_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f30bfb60cf433c3, []int{1, 0}
}

type Config struct {
	ApiServer            *ApiServer `protobuf:"bytes,1,opt,name=api_server,json=apiServer,proto3" json:"api_server,omitempty"`
	PollingInterval      string     `protobuf:"bytes,3,opt,name=polling_interval,json=pollingInterval,proto3" json:"polling_interval,omitempty"`
	MaxRetries           int32      `protobuf:"varint,4,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f30bfb60cf433c3, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetApiServer() *ApiServer {
	if m != nil {
		return m.ApiServer
	}
	return nil
}

func (m *Config) GetPollingInterval() string {
	if m != nil {
		return m.PollingInterval
	}
	return ""
}

func (m *Config) GetMaxRetries() int32 {
	if m != nil {
		return m.MaxRetries
	}
	return 0
}

func (*Config) XXX_MessageName() string {
	return "ImageAPI.Config"
}

type ApiServer struct {
	Server               string           `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Port                 int32            `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Https                bool             `protobuf:"varint,3,opt,name=https,proto3" json:"https,omitempty"`
	ApiBase              string           `protobuf:"bytes,4,opt,name=api_base,json=apiBase,proto3" json:"api_base,omitempty"`
	Status               ApiServer_Status `protobuf:"varint,5,opt,name=status,proto3,enum=ImageAPI.ApiServer_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ApiServer) Reset()         { *m = ApiServer{} }
func (m *ApiServer) String() string { return proto.CompactTextString(m) }
func (*ApiServer) ProtoMessage()    {}
func (*ApiServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f30bfb60cf433c3, []int{1}
}
func (m *ApiServer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApiServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApiServer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApiServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiServer.Merge(m, src)
}
func (m *ApiServer) XXX_Size() int {
	return m.Size()
}
func (m *ApiServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiServer.DiscardUnknown(m)
}

var xxx_messageInfo_ApiServer proto.InternalMessageInfo

func (m *ApiServer) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *ApiServer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ApiServer) GetHttps() bool {
	if m != nil {
		return m.Https
	}
	return false
}

func (m *ApiServer) GetApiBase() string {
	if m != nil {
		return m.ApiBase
	}
	return ""
}

func (m *ApiServer) GetStatus() ApiServer_Status {
	if m != nil {
		return m.Status
	}
	return ApiServer_UNKNOWN
}

func (*ApiServer) XXX_MessageName() string {
	return "ImageAPI.ApiServer"
}
func init() {
	proto.RegisterEnum("ImageAPI.ApiServer_Status", ApiServer_Status_name, ApiServer_Status_value)
	golang_proto.RegisterEnum("ImageAPI.ApiServer_Status", ApiServer_Status_name, ApiServer_Status_value)
	proto.RegisterType((*Config)(nil), "ImageAPI.Config")
	golang_proto.RegisterType((*Config)(nil), "ImageAPI.Config")
	proto.RegisterType((*ApiServer)(nil), "ImageAPI.ApiServer")
	golang_proto.RegisterType((*ApiServer)(nil), "ImageAPI.ApiServer")
}

func init() { proto.RegisterFile("imageapi-config.proto", fileDescriptor_2f30bfb60cf433c3) }
func init() { golang_proto.RegisterFile("imageapi-config.proto", fileDescriptor_2f30bfb60cf433c3) }

var fileDescriptor_2f30bfb60cf433c3 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x6b, 0xc2, 0x30,
	0x18, 0xc6, 0x17, 0x67, 0xab, 0x7d, 0x85, 0xad, 0x64, 0x7f, 0xe8, 0x3c, 0x74, 0xc5, 0x53, 0x77,
	0xb0, 0x42, 0x77, 0xdc, 0x49, 0xc7, 0x0e, 0x22, 0xe8, 0x88, 0x8c, 0xc1, 0x2e, 0x92, 0x4a, 0xac,
	0x01, 0x6b, 0x42, 0x13, 0xc5, 0x8f, 0xb0, 0x8f, 0xb5, 0xd3, 0xf0, 0xb8, 0x8f, 0x30, 0xf4, 0x8b,
	0x0c, 0xd3, 0x3a, 0x76, 0xd8, 0xed, 0x7d, 0x7e, 0x79, 0xde, 0xf7, 0xc9, 0x9b, 0xc0, 0x15, 0xcf,
	0x68, 0xca, 0xa8, 0xe4, 0xed, 0xa9, 0x58, 0xce, 0x78, 0x1a, 0xc9, 0x5c, 0x68, 0x81, 0xeb, 0xfd,
	0x03, 0xee, 0x3e, 0xf7, 0x9b, 0xed, 0x94, 0xeb, 0xf9, 0x2a, 0x89, 0xa6, 0x22, 0xeb, 0xa4, 0x22,
	0x15, 0x1d, 0x63, 0x48, 0x56, 0x33, 0xa3, 0x8c, 0x30, 0x55, 0xd1, 0xd8, 0x7a, 0x47, 0x60, 0x3f,
	0x9a, 0x49, 0x38, 0x06, 0xa0, 0x92, 0x4f, 0x14, 0xcb, 0xd7, 0x2c, 0xf7, 0x50, 0x80, 0xc2, 0x46,
	0x7c, 0x11, 0x1d, 0x07, 0x47, 0x5d, 0xc9, 0xc7, 0xe6, 0x88, 0x38, 0xf4, 0x58, 0xe2, 0x3b, 0x70,
	0xa5, 0x58, 0x2c, 0xf8, 0x32, 0x9d, 0xf0, 0xa5, 0x66, 0xf9, 0x9a, 0x2e, 0xbc, 0xd3, 0x00, 0x85,
	0x0e, 0x39, 0x2f, 0x79, 0xbf, 0xc4, 0xf8, 0x16, 0x1a, 0x19, 0xdd, 0x4c, 0x72, 0xa6, 0x73, 0xce,
	0x94, 0x57, 0x0d, 0x50, 0x68, 0x11, 0xc8, 0xe8, 0x86, 0x14, 0xa4, 0xf5, 0x89, 0xc0, 0xf9, 0x0d,
	0xc1, 0xd7, 0x60, 0xff, 0xb9, 0x89, 0x43, 0x4a, 0x85, 0x31, 0x54, 0xa5, 0xc8, 0xb5, 0x57, 0x31,
	0xfd, 0xa6, 0xc6, 0x97, 0x60, 0xcd, 0xb5, 0x96, 0xca, 0x44, 0xd7, 0x49, 0x21, 0xf0, 0x0d, 0xd4,
	0x0f, 0xfb, 0x24, 0x54, 0x31, 0x93, 0xe6, 0x90, 0x1a, 0x95, 0xbc, 0x47, 0x15, 0xc3, 0x31, 0xd8,
	0x4a, 0x53, 0xbd, 0x52, 0x9e, 0x15, 0xa0, 0xf0, 0x2c, 0x6e, 0xfe, 0xb3, 0x66, 0x34, 0x36, 0x0e,
	0x52, 0x3a, 0x5b, 0x21, 0xd8, 0x05, 0xc1, 0x0d, 0xa8, 0xbd, 0x0c, 0x07, 0xc3, 0xd1, 0xeb, 0xd0,
	0x3d, 0xc1, 0x36, 0x54, 0x46, 0x03, 0x17, 0x61, 0x07, 0xac, 0x27, 0x42, 0x46, 0xc4, 0xad, 0xf4,
	0x82, 0xed, 0xce, 0x47, 0x5f, 0x3b, 0x1f, 0x7d, 0xef, 0x7c, 0xf4, 0xb1, 0xf7, 0xd1, 0x76, 0xef,
	0xa3, 0x37, 0x88, 0x1e, 0x8e, 0x7f, 0x97, 0xd8, 0xe6, 0xf1, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xb7, 0xbc, 0xf1, 0x7f, 0xce, 0x01, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxRetries != 0 {
		i = encodeVarintImageapiConfig(dAtA, i, uint64(m.MaxRetries))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PollingInterval) > 0 {
		i -= len(m.PollingInterval)
		copy(dAtA[i:], m.PollingInterval)
		i = encodeVarintImageapiConfig(dAtA, i, uint64(len(m.PollingInterval)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ApiServer != nil {
		{
			size, err := m.ApiServer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintImageapiConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApiServer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApiServer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApiServer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintImageapiConfig(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ApiBase) > 0 {
		i -= len(m.ApiBase)
		copy(dAtA[i:], m.ApiBase)
		i = encodeVarintImageapiConfig(dAtA, i, uint64(len(m.ApiBase)))
		i--
		dAtA[i] = 0x22
	}
	if m.Https {
		i--
		if m.Https {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Port != 0 {
		i = encodeVarintImageapiConfig(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Server) > 0 {
		i -= len(m.Server)
		copy(dAtA[i:], m.Server)
		i = encodeVarintImageapiConfig(dAtA, i, uint64(len(m.Server)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintImageapiConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovImageapiConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApiServer != nil {
		l = m.ApiServer.Size()
		n += 1 + l + sovImageapiConfig(uint64(l))
	}
	l = len(m.PollingInterval)
	if l > 0 {
		n += 1 + l + sovImageapiConfig(uint64(l))
	}
	if m.MaxRetries != 0 {
		n += 1 + sovImageapiConfig(uint64(m.MaxRetries))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ApiServer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Server)
	if l > 0 {
		n += 1 + l + sovImageapiConfig(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovImageapiConfig(uint64(m.Port))
	}
	if m.Https {
		n += 2
	}
	l = len(m.ApiBase)
	if l > 0 {
		n += 1 + l + sovImageapiConfig(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovImageapiConfig(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovImageapiConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozImageapiConfig(x uint64) (n int) {
	return sovImageapiConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageapiConfig
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiServer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
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
				return ErrInvalidLengthImageapiConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthImageapiConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ApiServer == nil {
				m.ApiServer = &ApiServer{}
			}
			if err := m.ApiServer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollingInterval", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
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
				return ErrInvalidLengthImageapiConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageapiConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PollingInterval = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetries", wireType)
			}
			m.MaxRetries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRetries |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImageapiConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImageapiConfig
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
func (m *ApiServer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImageapiConfig
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
			return fmt.Errorf("proto: ApiServer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiServer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
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
				return ErrInvalidLengthImageapiConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageapiConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Server = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Https", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Https = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
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
				return ErrInvalidLengthImageapiConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthImageapiConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImageapiConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ApiServer_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImageapiConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthImageapiConfig
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
func skipImageapiConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImageapiConfig
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
					return 0, ErrIntOverflowImageapiConfig
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
					return 0, ErrIntOverflowImageapiConfig
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
				return 0, ErrInvalidLengthImageapiConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupImageapiConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthImageapiConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthImageapiConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImageapiConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupImageapiConfig = fmt.Errorf("proto: unexpected end of group")
)
