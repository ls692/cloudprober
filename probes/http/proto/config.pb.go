// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: github.com/google/cloudprober/probes/http/proto/config.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/google/cloudprober/common/oauth/proto"
	proto2 "github.com/google/cloudprober/common/tlsconfig/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProbeConf_ProtocolType int32

const (
	ProbeConf_HTTP  ProbeConf_ProtocolType = 0
	ProbeConf_HTTPS ProbeConf_ProtocolType = 1
)

// Enum value maps for ProbeConf_ProtocolType.
var (
	ProbeConf_ProtocolType_name = map[int32]string{
		0: "HTTP",
		1: "HTTPS",
	}
	ProbeConf_ProtocolType_value = map[string]int32{
		"HTTP":  0,
		"HTTPS": 1,
	}
)

func (x ProbeConf_ProtocolType) Enum() *ProbeConf_ProtocolType {
	p := new(ProbeConf_ProtocolType)
	*p = x
	return p
}

func (x ProbeConf_ProtocolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProbeConf_ProtocolType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_google_cloudprober_probes_http_proto_config_proto_enumTypes[0].Descriptor()
}

func (ProbeConf_ProtocolType) Type() protoreflect.EnumType {
	return &file_github_com_google_cloudprober_probes_http_proto_config_proto_enumTypes[0]
}

func (x ProbeConf_ProtocolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProbeConf_ProtocolType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProbeConf_ProtocolType(num)
	return nil
}

// Deprecated: Use ProbeConf_ProtocolType.Descriptor instead.
func (ProbeConf_ProtocolType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

type ProbeConf_Method int32

const (
	ProbeConf_GET     ProbeConf_Method = 0
	ProbeConf_POST    ProbeConf_Method = 1
	ProbeConf_PUT     ProbeConf_Method = 2
	ProbeConf_HEAD    ProbeConf_Method = 3
	ProbeConf_DELETE  ProbeConf_Method = 4
	ProbeConf_PATCH   ProbeConf_Method = 5
	ProbeConf_OPTIONS ProbeConf_Method = 6
)

// Enum value maps for ProbeConf_Method.
var (
	ProbeConf_Method_name = map[int32]string{
		0: "GET",
		1: "POST",
		2: "PUT",
		3: "HEAD",
		4: "DELETE",
		5: "PATCH",
		6: "OPTIONS",
	}
	ProbeConf_Method_value = map[string]int32{
		"GET":     0,
		"POST":    1,
		"PUT":     2,
		"HEAD":    3,
		"DELETE":  4,
		"PATCH":   5,
		"OPTIONS": 6,
	}
)

func (x ProbeConf_Method) Enum() *ProbeConf_Method {
	p := new(ProbeConf_Method)
	*p = x
	return p
}

func (x ProbeConf_Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProbeConf_Method) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_google_cloudprober_probes_http_proto_config_proto_enumTypes[1].Descriptor()
}

func (ProbeConf_Method) Type() protoreflect.EnumType {
	return &file_github_com_google_cloudprober_probes_http_proto_config_proto_enumTypes[1]
}

func (x ProbeConf_Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProbeConf_Method) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProbeConf_Method(num)
	return nil
}

// Deprecated: Use ProbeConf_Method.Descriptor instead.
func (ProbeConf_Method) EnumDescriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescGZIP(), []int{0, 1}
}

// Next tag: 17
type ProbeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Which HTTP protocol to use
	Protocol *ProbeConf_ProtocolType `protobuf:"varint,1,opt,name=protocol,enum=cloudprober.probes.http.ProbeConf_ProtocolType,def=0" json:"protocol,omitempty"`
	// Relative URL (to append to all targets). Must begin with '/'
	RelativeUrl *string `protobuf:"bytes,2,opt,name=relative_url,json=relativeUrl" json:"relative_url,omitempty"`
	// Port for HTTP requests. If not specfied, port is selected in the following
	// order:
	//  - If port is provided by the targets (e.g. kubernetes endpoint or
	//    service), that port is used.
	//  - 80 for HTTP and 443 for HTTPS.
	Port *int32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	// Whether to resolve the target before making the request. If set to false,
	// we hand over the target and relative_url directly to the golang's HTTP
	// module, Otherwise, we resolve the target first to an IP address and
	// make a request using that while passing target name as Host header.
	ResolveFirst *bool `protobuf:"varint,4,opt,name=resolve_first,json=resolveFirst,def=0" json:"resolve_first,omitempty"`
	// Export response (body) count as a metric
	ExportResponseAsMetrics *bool `protobuf:"varint,5,opt,name=export_response_as_metrics,json=exportResponseAsMetrics,def=0" json:"export_response_as_metrics,omitempty"`
	// HTTP request method
	Method *ProbeConf_Method `protobuf:"varint,7,opt,name=method,enum=cloudprober.probes.http.ProbeConf_Method,def=0" json:"method,omitempty"`
	// HTTP request headers
	Headers []*ProbeConf_Header `protobuf:"bytes,8,rep,name=headers" json:"headers,omitempty"`
	// Request body.
	Body *string `protobuf:"bytes,9,opt,name=body" json:"body,omitempty"`
	// Enable HTTP keep-alive. If set to true, underlying connection is reused
	// for further probes. Default is to close the connection after every request.
	KeepAlive *bool `protobuf:"varint,10,opt,name=keep_alive,json=keepAlive" json:"keep_alive,omitempty"`
	// OAuth Config
	OauthConfig *proto1.Config `protobuf:"bytes,11,opt,name=oauth_config,json=oauthConfig" json:"oauth_config,omitempty"`
	// Disable HTTP2
	// Golang HTTP client automatically enables HTTP/2 if server supports it. This
	// option disables that behavior to enforce HTTP/1.1 for testing purpose.
	DisableHttp2 *bool `protobuf:"varint,13,opt,name=disable_http2,json=disableHttp2" json:"disable_http2,omitempty"`
	// Disable TLS certificate validation. If set to true, any certificate
	// presented by the server for any host name will be accepted
	// Deprecation: This option is now subsumed by the tls_config below. To
	// disable cert validation use:
	// tls_config {
	//   disable_cert_validation: true
	// }
	DisableCertValidation *bool `protobuf:"varint,14,opt,name=disable_cert_validation,json=disableCertValidation" json:"disable_cert_validation,omitempty"`
	// TLS config
	TlsConfig *proto2.TLSConfig `protobuf:"bytes,15,opt,name=tls_config,json=tlsConfig" json:"tls_config,omitempty"`
	// Proxy URL, e.g. http://myproxy:3128
	ProxyUrl *string `protobuf:"bytes,16,opt,name=proxy_url,json=proxyUrl" json:"proxy_url,omitempty"`
	// Requests per probe.
	// Number of HTTP requests per probe. Requests are executed concurrently and
	// each HTTP re contributes to probe results. For example, if you run two
	// requests per probe, "total" counter will be incremented by 2.
	RequestsPerProbe *int32 `protobuf:"varint,98,opt,name=requests_per_probe,json=requestsPerProbe,def=1" json:"requests_per_probe,omitempty"`
	// How long to wait between two requests to the same target
	// NOTE: This field is now deprecated and will be removed after the v0.10.3
	// releases.
	RequestsIntervalMsec *int32 `protobuf:"varint,99,opt,name=requests_interval_msec,json=requestsIntervalMsec,def=25" json:"requests_interval_msec,omitempty"`
}

// Default values for ProbeConf fields.
const (
	Default_ProbeConf_Protocol                = ProbeConf_HTTP
	Default_ProbeConf_ResolveFirst            = bool(false)
	Default_ProbeConf_ExportResponseAsMetrics = bool(false)
	Default_ProbeConf_Method                  = ProbeConf_GET
	Default_ProbeConf_RequestsPerProbe        = int32(1)
	Default_ProbeConf_RequestsIntervalMsec    = int32(25)
)

func (x *ProbeConf) Reset() {
	*x = ProbeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf) ProtoMessage() {}

func (x *ProbeConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf.ProtoReflect.Descriptor instead.
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeConf) GetProtocol() ProbeConf_ProtocolType {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return Default_ProbeConf_Protocol
}

func (x *ProbeConf) GetRelativeUrl() string {
	if x != nil && x.RelativeUrl != nil {
		return *x.RelativeUrl
	}
	return ""
}

func (x *ProbeConf) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *ProbeConf) GetResolveFirst() bool {
	if x != nil && x.ResolveFirst != nil {
		return *x.ResolveFirst
	}
	return Default_ProbeConf_ResolveFirst
}

func (x *ProbeConf) GetExportResponseAsMetrics() bool {
	if x != nil && x.ExportResponseAsMetrics != nil {
		return *x.ExportResponseAsMetrics
	}
	return Default_ProbeConf_ExportResponseAsMetrics
}

func (x *ProbeConf) GetMethod() ProbeConf_Method {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return Default_ProbeConf_Method
}

func (x *ProbeConf) GetHeaders() []*ProbeConf_Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ProbeConf) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *ProbeConf) GetKeepAlive() bool {
	if x != nil && x.KeepAlive != nil {
		return *x.KeepAlive
	}
	return false
}

func (x *ProbeConf) GetOauthConfig() *proto1.Config {
	if x != nil {
		return x.OauthConfig
	}
	return nil
}

func (x *ProbeConf) GetDisableHttp2() bool {
	if x != nil && x.DisableHttp2 != nil {
		return *x.DisableHttp2
	}
	return false
}

func (x *ProbeConf) GetDisableCertValidation() bool {
	if x != nil && x.DisableCertValidation != nil {
		return *x.DisableCertValidation
	}
	return false
}

func (x *ProbeConf) GetTlsConfig() *proto2.TLSConfig {
	if x != nil {
		return x.TlsConfig
	}
	return nil
}

func (x *ProbeConf) GetProxyUrl() string {
	if x != nil && x.ProxyUrl != nil {
		return *x.ProxyUrl
	}
	return ""
}

func (x *ProbeConf) GetRequestsPerProbe() int32 {
	if x != nil && x.RequestsPerProbe != nil {
		return *x.RequestsPerProbe
	}
	return Default_ProbeConf_RequestsPerProbe
}

func (x *ProbeConf) GetRequestsIntervalMsec() int32 {
	if x != nil && x.RequestsIntervalMsec != nil {
		return *x.RequestsIntervalMsec
	}
	return Default_ProbeConf_RequestsIntervalMsec
}

type ProbeConf_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *ProbeConf_Header) Reset() {
	*x = ProbeConf_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf_Header) ProtoMessage() {}

func (x *ProbeConf_Header) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf_Header.ProtoReflect.Descriptor instead.
func (*ProbeConf_Header) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProbeConf_Header) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ProbeConf_Header) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_github_com_google_cloudprober_probes_http_proto_config_proto protoreflect.FileDescriptor

var file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x6c, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x07, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x51, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x04, 0x48, 0x54, 0x54, 0x50,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x2a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x42, 0x0a,
	0x1a, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x61, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x17, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x46, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x03, 0x47, 0x45,
	0x54, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x43, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x32,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x48,
	0x74, 0x74, 0x70, 0x32, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0a,
	0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74,
	0x6c, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x09, 0x74, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x12, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x18, 0x62, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x31, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x38, 0x0a, 0x16, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x63, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x32, 0x35, 0x52,
	0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x4d, 0x73, 0x65, 0x63, 0x1a, 0x32, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54,
	0x50, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x53, 0x10, 0x01, 0x22, 0x52,
	0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50,
	0x55, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53,
	0x10, 0x06, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescOnce sync.Once
	file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescData = file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDesc
)

func file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescData)
	})
	return file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDescData
}

var file_github_com_google_cloudprober_probes_http_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_google_cloudprober_probes_http_proto_config_proto_goTypes = []interface{}{
	(ProbeConf_ProtocolType)(0), // 0: cloudprober.probes.http.ProbeConf.ProtocolType
	(ProbeConf_Method)(0),       // 1: cloudprober.probes.http.ProbeConf.Method
	(*ProbeConf)(nil),           // 2: cloudprober.probes.http.ProbeConf
	(*ProbeConf_Header)(nil),    // 3: cloudprober.probes.http.ProbeConf.Header
	(*proto1.Config)(nil),       // 4: cloudprober.oauth.Config
	(*proto2.TLSConfig)(nil),    // 5: cloudprober.tlsconfig.TLSConfig
}
var file_github_com_google_cloudprober_probes_http_proto_config_proto_depIdxs = []int32{
	0, // 0: cloudprober.probes.http.ProbeConf.protocol:type_name -> cloudprober.probes.http.ProbeConf.ProtocolType
	1, // 1: cloudprober.probes.http.ProbeConf.method:type_name -> cloudprober.probes.http.ProbeConf.Method
	3, // 2: cloudprober.probes.http.ProbeConf.headers:type_name -> cloudprober.probes.http.ProbeConf.Header
	4, // 3: cloudprober.probes.http.ProbeConf.oauth_config:type_name -> cloudprober.oauth.Config
	5, // 4: cloudprober.probes.http.ProbeConf.tls_config:type_name -> cloudprober.tlsconfig.TLSConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_google_cloudprober_probes_http_proto_config_proto_init() }
func file_github_com_google_cloudprober_probes_http_proto_config_proto_init() {
	if File_github_com_google_cloudprober_probes_http_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf_Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_google_cloudprober_probes_http_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_google_cloudprober_probes_http_proto_config_proto_depIdxs,
		EnumInfos:         file_github_com_google_cloudprober_probes_http_proto_config_proto_enumTypes,
		MessageInfos:      file_github_com_google_cloudprober_probes_http_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_google_cloudprober_probes_http_proto_config_proto = out.File
	file_github_com_google_cloudprober_probes_http_proto_config_proto_rawDesc = nil
	file_github_com_google_cloudprober_probes_http_proto_config_proto_goTypes = nil
	file_github_com_google_cloudprober_probes_http_proto_config_proto_depIdxs = nil
}
