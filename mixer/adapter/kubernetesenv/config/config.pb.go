// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/kubernetesenv/config/config.proto

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		mixer/adapter/kubernetesenv/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Configuration parameters for the kubernetes adapter. These params
// control the manner in which the kubernetes adapter discovers and
// generates values related to pod information.
//
// The adapter works by looking up pod information by UIDs (of the
// form: "kubernetes://pod.namespace"). It expects that the UIDs will be
// supplied in an input map for three distinct traffic classes (source,
// destination, and origin).
//
// For all valid UIDs supplied, this adapter generates output
// values containing information about the related pods.
type Params struct {
	// File path to discover kubeconfig. For in-cluster configuration,
	// this should be left unset. For local configuration, this should
	// be set to the path of a kubeconfig file that can be used to
	// reach a kubernetes API server.
	//
	// NOTE: The kubernetes adapter will use the value of the env var
	// KUBECONFIG in the case where it is set (overriding any value configured
	// through this proto).
	//
	// Default: "" (unset)
	KubeconfigPath string `protobuf:"bytes,1,opt,name=kubeconfig_path,json=kubeconfigPath,proto3" json:"kubeconfig_path,omitempty"`
	// Controls the resync period of the kubernetes cluster info cache.
	// The cache will watch for events and every so often completely resync.
	// This controls how frequently the complete resync occurs.
	//
	// Default: 5 minutes
	CacheRefreshDuration time.Duration `protobuf:"bytes,2,opt,name=cache_refresh_duration,json=cacheRefreshDuration,stdduration" json:"cache_refresh_duration"`
	// Configures the cluster domain name to use for service name normalization.
	//
	// Default: svc.cluster.local
	ClusterDomainName string `protobuf:"bytes,3,opt,name=cluster_domain_name,json=clusterDomainName,proto3" json:"cluster_domain_name,omitempty"`
	// In order to extract the service associated with a source, destination, or
	// origin, this adapter relies on pod labels. In particular, it looks for
	// the value of a specific label, as specified by this parameter.
	//
	// Default: app
	PodLabelForService string `protobuf:"bytes,4,opt,name=pod_label_for_service,json=podLabelForService,proto3" json:"pod_label_for_service,omitempty"`
	// In order to extract the service associated with a source, destination, or
	// origin, this adapter relies on pod labels. In particular, it looks for
	// the value of a specific label for istio component services, as specified
	// by this parameter.
	//
	// Default: istio
	PodLabelForIstioComponentService string `protobuf:"bytes,5,opt,name=pod_label_for_istio_component_service,json=podLabelForIstioComponentService,proto3" json:"pod_label_for_istio_component_service,omitempty"`
	//
	// Default: false
	LookupIngressSourceAndOriginValues bool `protobuf:"varint,6,opt,name=lookup_ingress_source_and_origin_values,json=lookupIngressSourceAndOriginValues,proto3" json:"lookup_ingress_source_and_origin_values,omitempty"`
	// Istio ingress service string. This is used to identify the
	// ingress service in requests.
	//
	// Default: "ingress.istio-system.svc.cluster.local"
	FullyQualifiedIstioIngressServiceName string `protobuf:"bytes,7,opt,name=fully_qualified_istio_ingress_service_name,json=fullyQualifiedIstioIngressServiceName,proto3" json:"fully_qualified_istio_ingress_service_name,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func init() {
	proto.RegisterType((*Params)(nil), "adapter.kubernetes.config.Params")
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.KubeconfigPath) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.KubeconfigPath)))
		i += copy(dAtA[i:], m.KubeconfigPath)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.CacheRefreshDuration)))
	n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.CacheRefreshDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.ClusterDomainName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ClusterDomainName)))
		i += copy(dAtA[i:], m.ClusterDomainName)
	}
	if len(m.PodLabelForService) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.PodLabelForService)))
		i += copy(dAtA[i:], m.PodLabelForService)
	}
	if len(m.PodLabelForIstioComponentService) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.PodLabelForIstioComponentService)))
		i += copy(dAtA[i:], m.PodLabelForIstioComponentService)
	}
	if m.LookupIngressSourceAndOriginValues {
		dAtA[i] = 0x30
		i++
		if m.LookupIngressSourceAndOriginValues {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.FullyQualifiedIstioIngressServiceName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.FullyQualifiedIstioIngressServiceName)))
		i += copy(dAtA[i:], m.FullyQualifiedIstioIngressServiceName)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = len(m.KubeconfigPath)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.CacheRefreshDuration)
	n += 1 + l + sovConfig(uint64(l))
	l = len(m.ClusterDomainName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.PodLabelForService)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.PodLabelForIstioComponentService)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.LookupIngressSourceAndOriginValues {
		n += 2
	}
	l = len(m.FullyQualifiedIstioIngressServiceName)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`KubeconfigPath:` + fmt.Sprintf("%v", this.KubeconfigPath) + `,`,
		`CacheRefreshDuration:` + strings.Replace(strings.Replace(this.CacheRefreshDuration.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`ClusterDomainName:` + fmt.Sprintf("%v", this.ClusterDomainName) + `,`,
		`PodLabelForService:` + fmt.Sprintf("%v", this.PodLabelForService) + `,`,
		`PodLabelForIstioComponentService:` + fmt.Sprintf("%v", this.PodLabelForIstioComponentService) + `,`,
		`LookupIngressSourceAndOriginValues:` + fmt.Sprintf("%v", this.LookupIngressSourceAndOriginValues) + `,`,
		`FullyQualifiedIstioIngressServiceName:` + fmt.Sprintf("%v", this.FullyQualifiedIstioIngressServiceName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KubeconfigPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KubeconfigPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheRefreshDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.CacheRefreshDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterDomainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterDomainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodLabelForService", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodLabelForService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodLabelForIstioComponentService", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodLabelForIstioComponentService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LookupIngressSourceAndOriginValues", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LookupIngressSourceAndOriginValues = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullyQualifiedIstioIngressServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FullyQualifiedIstioIngressServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/adapter/kubernetesenv/config/config.proto", fileDescriptorConfig)
}

var fileDescriptorConfig = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbb, 0x6e, 0xd4, 0x40,
	0x14, 0x86, 0x3d, 0x04, 0x96, 0x60, 0x24, 0x10, 0x26, 0xa0, 0x4d, 0x8a, 0xc9, 0x2a, 0x52, 0x94,
	0x15, 0x85, 0x2d, 0xa0, 0xa1, 0x25, 0x44, 0x48, 0x91, 0x10, 0x09, 0xbb, 0x12, 0x52, 0x68, 0x46,
	0xb3, 0xf6, 0xb1, 0x77, 0x14, 0x7b, 0x8e, 0x99, 0xcb, 0x0a, 0x3a, 0x1e, 0x81, 0x92, 0x47, 0xe0,
	0x51, 0xb6, 0x4c, 0x49, 0x05, 0xac, 0x69, 0x28, 0xf7, 0x11, 0x90, 0x67, 0xec, 0x84, 0x54, 0xbe,
	0xfc, 0xdf, 0xf9, 0xfe, 0xa3, 0x99, 0x30, 0xa9, 0xc4, 0x27, 0x50, 0x09, 0xcf, 0x78, 0x6d, 0x40,
	0x25, 0xe7, 0x76, 0x06, 0x4a, 0x82, 0x01, 0x0d, 0x72, 0x91, 0xa4, 0x28, 0x73, 0x51, 0x74, 0x8f,
	0xb8, 0x56, 0x68, 0x30, 0xda, 0xee, 0xd0, 0xf8, 0x0a, 0x8d, 0x3d, 0xb0, 0xb3, 0x55, 0x60, 0x81,
	0x8e, 0x4a, 0xda, 0x37, 0x3f, 0xb0, 0x43, 0x0b, 0xc4, 0xa2, 0x84, 0xc4, 0x7d, 0xcd, 0x6c, 0x9e,
	0x64, 0x56, 0x71, 0x23, 0x50, 0xfa, 0x7c, 0x6f, 0xbd, 0x11, 0x0e, 0x4e, 0xb9, 0xe2, 0x95, 0x8e,
	0x0e, 0xc2, 0xfb, 0xad, 0xd5, 0xeb, 0x58, 0xcd, 0xcd, 0x7c, 0x48, 0x46, 0x64, 0x7c, 0x67, 0x72,
	0xef, 0xea, 0xf7, 0x29, 0x37, 0xf3, 0xe8, 0x2c, 0x7c, 0x9c, 0xf2, 0x74, 0x0e, 0x4c, 0x41, 0xae,
	0x40, 0xcf, 0x59, 0xef, 0x1c, 0xde, 0x18, 0x91, 0xf1, 0xdd, 0x67, 0xdb, 0xb1, 0x2f, 0x8d, 0xfb,
	0xd2, 0xf8, 0xa8, 0x03, 0x0e, 0x37, 0x97, 0x3f, 0x77, 0x83, 0x6f, 0xbf, 0x76, 0xc9, 0x64, 0xcb,
	0x29, 0x26, 0xde, 0xd0, 0xe7, 0x51, 0x1c, 0x3e, 0x4c, 0x4b, 0xab, 0x0d, 0x28, 0x96, 0x61, 0xc5,
	0x85, 0x64, 0x92, 0x57, 0x30, 0xdc, 0x70, 0x7b, 0x3c, 0xe8, 0xa2, 0x23, 0x97, 0xbc, 0xe5, 0x15,
	0x44, 0x4f, 0xc3, 0x47, 0x35, 0x66, 0xac, 0xe4, 0x33, 0x28, 0x59, 0x8e, 0x8a, 0x69, 0x50, 0x0b,
	0x91, 0xc2, 0xf0, 0xa6, 0x9b, 0x88, 0x6a, 0xcc, 0xde, 0xb4, 0xd9, 0x6b, 0x54, 0x53, 0x9f, 0x44,
	0x27, 0xe1, 0xfe, 0xf5, 0x11, 0xa1, 0x8d, 0x40, 0x96, 0x62, 0x55, 0xa3, 0x04, 0x69, 0x2e, 0x15,
	0xb7, 0x9c, 0x62, 0xf4, 0x9f, 0xe2, 0xb8, 0x25, 0x5f, 0xf5, 0x60, 0x2f, 0x9c, 0x86, 0x07, 0x25,
	0xe2, 0xb9, 0xad, 0x99, 0x90, 0x85, 0x02, 0xad, 0x99, 0x46, 0xab, 0x52, 0x60, 0x5c, 0x66, 0x0c,
	0x95, 0x28, 0x84, 0x64, 0x0b, 0x5e, 0x5a, 0xd0, 0xc3, 0xc1, 0x88, 0x8c, 0x37, 0x27, 0x7b, 0x1e,
	0x3f, 0xf6, 0xf4, 0xd4, 0xc1, 0x2f, 0x65, 0x76, 0xe2, 0xd0, 0xf7, 0x8e, 0x8c, 0xce, 0xc2, 0x27,
	0xb9, 0x2d, 0xcb, 0xcf, 0xec, 0xa3, 0xe5, 0xa5, 0xc8, 0x05, 0x64, 0xdd, 0x9e, 0x97, 0x1d, 0xbe,
	0xdd, 0x9f, 0xcf, 0x6d, 0xb7, 0xea, 0xbe, 0x9b, 0x78, 0xd7, 0x0f, 0xb8, 0x6d, 0xfb, 0x12, 0x4f,
	0xb7, 0x67, 0x76, 0xf8, 0x62, 0xb9, 0xa2, 0xc1, 0xc5, 0x8a, 0x06, 0x3f, 0x56, 0x34, 0x58, 0xaf,
	0x68, 0xf0, 0xa5, 0xa1, 0xe4, 0x7b, 0x43, 0x83, 0x65, 0x43, 0xc9, 0x45, 0x43, 0xc9, 0xef, 0x86,
	0x92, 0xbf, 0x0d, 0x0d, 0xd6, 0x0d, 0x25, 0x5f, 0xff, 0xd0, 0xe0, 0xc3, 0xc0, 0x5f, 0xfe, 0x6c,
	0xe0, 0x2e, 0xf4, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xee, 0x79, 0xd4, 0xb7, 0x02,
	0x00, 0x00,
}