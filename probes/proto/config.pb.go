// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/proto/config.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto2 "github.com/google/cloudprober/metrics/proto"
	proto6 "github.com/google/cloudprober/probes/dns/proto"
	proto7 "github.com/google/cloudprober/probes/external/proto"
	proto10 "github.com/google/cloudprober/probes/grpc/proto"
	proto5 "github.com/google/cloudprober/probes/http/proto"
	proto4 "github.com/google/cloudprober/probes/ping/proto"
	proto8 "github.com/google/cloudprober/probes/udp/proto"
	proto9 "github.com/google/cloudprober/probes/udplistener/proto"
	proto1 "github.com/google/cloudprober/targets/proto"
	proto3 "github.com/google/cloudprober/validators/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProbeDef_Type int32

const (
	ProbeDef_PING         ProbeDef_Type = 0
	ProbeDef_HTTP         ProbeDef_Type = 1
	ProbeDef_DNS          ProbeDef_Type = 2
	ProbeDef_EXTERNAL     ProbeDef_Type = 3
	ProbeDef_UDP          ProbeDef_Type = 4
	ProbeDef_UDP_LISTENER ProbeDef_Type = 5
	ProbeDef_GRPC         ProbeDef_Type = 6
	// One of the extension probe types. See "extensions" below for more
	// details.
	ProbeDef_EXTENSION ProbeDef_Type = 98
	// USER_DEFINED probe type is for a one off probe that you want to compile
	// into cloudprober, but you don't expect it to be reused. If you expect
	// it to be reused, you should consider adding it using the extensions
	// mechanism.
	ProbeDef_USER_DEFINED ProbeDef_Type = 99
)

var ProbeDef_Type_name = map[int32]string{
	0:  "PING",
	1:  "HTTP",
	2:  "DNS",
	3:  "EXTERNAL",
	4:  "UDP",
	5:  "UDP_LISTENER",
	6:  "GRPC",
	98: "EXTENSION",
	99: "USER_DEFINED",
}

var ProbeDef_Type_value = map[string]int32{
	"PING":         0,
	"HTTP":         1,
	"DNS":          2,
	"EXTERNAL":     3,
	"UDP":          4,
	"UDP_LISTENER": 5,
	"GRPC":         6,
	"EXTENSION":    98,
	"USER_DEFINED": 99,
}

func (x ProbeDef_Type) Enum() *ProbeDef_Type {
	p := new(ProbeDef_Type)
	*p = x
	return p
}

func (x ProbeDef_Type) String() string {
	return proto.EnumName(ProbeDef_Type_name, int32(x))
}

func (x *ProbeDef_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeDef_Type_value, data, "ProbeDef_Type")
	if err != nil {
		return err
	}
	*x = ProbeDef_Type(value)
	return nil
}

func (ProbeDef_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8261ee47d309caad, []int{0, 0}
}

// IP version to use for networking probes. If specified, this is used at the
// time of resolving a target, picking the correct IP for the source IP if
// source_interface option is provided, and to craft the packet correctly
// for PING probes.
//
// If ip_version is not configured but source_ip is provided, we get
// ip_version from it. If both are  confgiured, an error is returned if there
// is a conflict between the two.
//
// If left unspecified and both addresses are available in resolve call or on
// source interface, IPv4 is preferred.
// Future work: provide an option to prefer IPv4 and IPv6 explicitly.
type ProbeDef_IPVersion int32

const (
	ProbeDef_IP_VERSION_UNSPECIFIED ProbeDef_IPVersion = 0
	ProbeDef_IPV4                   ProbeDef_IPVersion = 1
	ProbeDef_IPV6                   ProbeDef_IPVersion = 2
)

var ProbeDef_IPVersion_name = map[int32]string{
	0: "IP_VERSION_UNSPECIFIED",
	1: "IPV4",
	2: "IPV6",
}

var ProbeDef_IPVersion_value = map[string]int32{
	"IP_VERSION_UNSPECIFIED": 0,
	"IPV4":                   1,
	"IPV6":                   2,
}

func (x ProbeDef_IPVersion) Enum() *ProbeDef_IPVersion {
	p := new(ProbeDef_IPVersion)
	*p = x
	return p
}

func (x ProbeDef_IPVersion) String() string {
	return proto.EnumName(ProbeDef_IPVersion_name, int32(x))
}

func (x *ProbeDef_IPVersion) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeDef_IPVersion_value, data, "ProbeDef_IPVersion")
	if err != nil {
		return err
	}
	*x = ProbeDef_IPVersion(value)
	return nil
}

func (ProbeDef_IPVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8261ee47d309caad, []int{0, 1}
}

type ProbeDef struct {
	Name *string        `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Type *ProbeDef_Type `protobuf:"varint,2,req,name=type,enum=cloudprober.probes.ProbeDef_Type" json:"type,omitempty"`
	// Which machines this probe should run on. If defined, cloudprober will run
	// this probe only if machine's hostname matches this value.
	RunOn *string `protobuf:"bytes,3,opt,name=run_on,json=runOn" json:"run_on,omitempty"`
	// Interval between two probes
	IntervalMsec *int32 `protobuf:"varint,4,opt,name=interval_msec,json=intervalMsec,def=2000" json:"interval_msec,omitempty"`
	// Timeout for each probe
	TimeoutMsec *int32 `protobuf:"varint,5,opt,name=timeout_msec,json=timeoutMsec,def=1000" json:"timeout_msec,omitempty"`
	// Targets for the probe
	Targets *proto1.TargetsDef `protobuf:"bytes,6,req,name=targets" json:"targets,omitempty"`
	// Latency distribution. If specified, latency is stored as a distribution.
	LatencyDistribution *proto2.Dist `protobuf:"bytes,7,opt,name=latency_distribution,json=latencyDistribution" json:"latency_distribution,omitempty"`
	// Latency unit. Any string that's parseable by time.ParseDuration.
	// Valid values: "ns", "us" (or "µs"), "ms", "s", "m", "h".
	LatencyUnit *string `protobuf:"bytes,8,opt,name=latency_unit,json=latencyUnit,def=us" json:"latency_unit,omitempty"`
	// Validators are in experimental phase right now and can change at any time.
	// NOTE: Only PING, HTTP and DNS probes support validators.
	Validator []*proto3.Validator `protobuf:"bytes,9,rep,name=validator" json:"validator,omitempty"`
	// Set the source IP to send packets from, either by providing an IP address
	// directly, or a network interface.
	//
	// Types that are valid to be assigned to SourceIpConfig:
	//	*ProbeDef_SourceIp
	//	*ProbeDef_SourceInterface
	SourceIpConfig isProbeDef_SourceIpConfig `protobuf_oneof:"source_ip_config"`
	IpVersion      *ProbeDef_IPVersion       `protobuf:"varint,12,opt,name=ip_version,json=ipVersion,enum=cloudprober.probes.ProbeDef_IPVersion" json:"ip_version,omitempty"`
	// How often to export stats. Probes usually run at a higher frequency (e.g.
	// every second); stats from individual probes are aggregated within
	// cloudprober until exported. In most cases, users don't need to change the
	// default.
	//
	// By default this field is set in the following way:
	// For all probes except UDP:
	//   stats_export_interval=max(interval, 10s)
	// For UDP:
	//   stats_export_interval=max(2*max(interval, timeout), 10s)
	StatsExportIntervalMsec *int32 `protobuf:"varint,13,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec" json:"stats_export_interval_msec,omitempty"`
	// Additional labels to add to the probe results. Label's value can either be
	// static or can be derived from target's labels.
	//
	// Example:
	//   additional_label {
	//     key: "src_zone"
	//     value: "{{.zone}}"
	//   }
	//   additional_label {
	//     key: "src_zone"
	//     value: "target.labels.zone"
	//   }
	AdditionalLabel []*AdditionalLabel `protobuf:"bytes,14,rep,name=additional_label,json=additionalLabel" json:"additional_label,omitempty"`
	// Types that are valid to be assigned to Probe:
	//	*ProbeDef_PingProbe
	//	*ProbeDef_HttpProbe
	//	*ProbeDef_DnsProbe
	//	*ProbeDef_ExternalProbe
	//	*ProbeDef_UdpProbe
	//	*ProbeDef_UdpListenerProbe
	//	*ProbeDef_GrpcProbe
	//	*ProbeDef_UserDefinedProbe
	Probe                        isProbeDef_Probe `protobuf_oneof:"probe"`
	DebugOptions                 *DebugOptions    `protobuf:"bytes,100,opt,name=debug_options,json=debugOptions" json:"debug_options,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}         `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *ProbeDef) Reset()         { *m = ProbeDef{} }
func (m *ProbeDef) String() string { return proto.CompactTextString(m) }
func (*ProbeDef) ProtoMessage()    {}
func (*ProbeDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_8261ee47d309caad, []int{0}
}

var extRange_ProbeDef = []proto.ExtensionRange{
	{Start: 200, End: 536870911},
}

func (*ProbeDef) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ProbeDef
}

func (m *ProbeDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeDef.Unmarshal(m, b)
}
func (m *ProbeDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeDef.Marshal(b, m, deterministic)
}
func (m *ProbeDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeDef.Merge(m, src)
}
func (m *ProbeDef) XXX_Size() int {
	return xxx_messageInfo_ProbeDef.Size(m)
}
func (m *ProbeDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeDef.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeDef proto.InternalMessageInfo

const Default_ProbeDef_IntervalMsec int32 = 2000
const Default_ProbeDef_TimeoutMsec int32 = 1000
const Default_ProbeDef_LatencyUnit string = "us"

func (m *ProbeDef) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProbeDef) GetType() ProbeDef_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ProbeDef_PING
}

func (m *ProbeDef) GetRunOn() string {
	if m != nil && m.RunOn != nil {
		return *m.RunOn
	}
	return ""
}

func (m *ProbeDef) GetIntervalMsec() int32 {
	if m != nil && m.IntervalMsec != nil {
		return *m.IntervalMsec
	}
	return Default_ProbeDef_IntervalMsec
}

func (m *ProbeDef) GetTimeoutMsec() int32 {
	if m != nil && m.TimeoutMsec != nil {
		return *m.TimeoutMsec
	}
	return Default_ProbeDef_TimeoutMsec
}

func (m *ProbeDef) GetTargets() *proto1.TargetsDef {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ProbeDef) GetLatencyDistribution() *proto2.Dist {
	if m != nil {
		return m.LatencyDistribution
	}
	return nil
}

func (m *ProbeDef) GetLatencyUnit() string {
	if m != nil && m.LatencyUnit != nil {
		return *m.LatencyUnit
	}
	return Default_ProbeDef_LatencyUnit
}

func (m *ProbeDef) GetValidator() []*proto3.Validator {
	if m != nil {
		return m.Validator
	}
	return nil
}

type isProbeDef_SourceIpConfig interface {
	isProbeDef_SourceIpConfig()
}

type ProbeDef_SourceIp struct {
	SourceIp string `protobuf:"bytes,10,opt,name=source_ip,json=sourceIp,oneof"`
}

type ProbeDef_SourceInterface struct {
	SourceInterface string `protobuf:"bytes,11,opt,name=source_interface,json=sourceInterface,oneof"`
}

func (*ProbeDef_SourceIp) isProbeDef_SourceIpConfig() {}

func (*ProbeDef_SourceInterface) isProbeDef_SourceIpConfig() {}

func (m *ProbeDef) GetSourceIpConfig() isProbeDef_SourceIpConfig {
	if m != nil {
		return m.SourceIpConfig
	}
	return nil
}

func (m *ProbeDef) GetSourceIp() string {
	if x, ok := m.GetSourceIpConfig().(*ProbeDef_SourceIp); ok {
		return x.SourceIp
	}
	return ""
}

func (m *ProbeDef) GetSourceInterface() string {
	if x, ok := m.GetSourceIpConfig().(*ProbeDef_SourceInterface); ok {
		return x.SourceInterface
	}
	return ""
}

func (m *ProbeDef) GetIpVersion() ProbeDef_IPVersion {
	if m != nil && m.IpVersion != nil {
		return *m.IpVersion
	}
	return ProbeDef_IP_VERSION_UNSPECIFIED
}

func (m *ProbeDef) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return 0
}

func (m *ProbeDef) GetAdditionalLabel() []*AdditionalLabel {
	if m != nil {
		return m.AdditionalLabel
	}
	return nil
}

type isProbeDef_Probe interface {
	isProbeDef_Probe()
}

type ProbeDef_PingProbe struct {
	PingProbe *proto4.ProbeConf `protobuf:"bytes,20,opt,name=ping_probe,json=pingProbe,oneof"`
}

type ProbeDef_HttpProbe struct {
	HttpProbe *proto5.ProbeConf `protobuf:"bytes,21,opt,name=http_probe,json=httpProbe,oneof"`
}

type ProbeDef_DnsProbe struct {
	DnsProbe *proto6.ProbeConf `protobuf:"bytes,22,opt,name=dns_probe,json=dnsProbe,oneof"`
}

type ProbeDef_ExternalProbe struct {
	ExternalProbe *proto7.ProbeConf `protobuf:"bytes,23,opt,name=external_probe,json=externalProbe,oneof"`
}

type ProbeDef_UdpProbe struct {
	UdpProbe *proto8.ProbeConf `protobuf:"bytes,24,opt,name=udp_probe,json=udpProbe,oneof"`
}

type ProbeDef_UdpListenerProbe struct {
	UdpListenerProbe *proto9.ProbeConf `protobuf:"bytes,25,opt,name=udp_listener_probe,json=udpListenerProbe,oneof"`
}

type ProbeDef_GrpcProbe struct {
	GrpcProbe *proto10.ProbeConf `protobuf:"bytes,26,opt,name=grpc_probe,json=grpcProbe,oneof"`
}

type ProbeDef_UserDefinedProbe struct {
	UserDefinedProbe string `protobuf:"bytes,99,opt,name=user_defined_probe,json=userDefinedProbe,oneof"`
}

func (*ProbeDef_PingProbe) isProbeDef_Probe() {}

func (*ProbeDef_HttpProbe) isProbeDef_Probe() {}

func (*ProbeDef_DnsProbe) isProbeDef_Probe() {}

func (*ProbeDef_ExternalProbe) isProbeDef_Probe() {}

func (*ProbeDef_UdpProbe) isProbeDef_Probe() {}

func (*ProbeDef_UdpListenerProbe) isProbeDef_Probe() {}

func (*ProbeDef_GrpcProbe) isProbeDef_Probe() {}

func (*ProbeDef_UserDefinedProbe) isProbeDef_Probe() {}

func (m *ProbeDef) GetProbe() isProbeDef_Probe {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *ProbeDef) GetPingProbe() *proto4.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_PingProbe); ok {
		return x.PingProbe
	}
	return nil
}

func (m *ProbeDef) GetHttpProbe() *proto5.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_HttpProbe); ok {
		return x.HttpProbe
	}
	return nil
}

func (m *ProbeDef) GetDnsProbe() *proto6.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_DnsProbe); ok {
		return x.DnsProbe
	}
	return nil
}

func (m *ProbeDef) GetExternalProbe() *proto7.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_ExternalProbe); ok {
		return x.ExternalProbe
	}
	return nil
}

func (m *ProbeDef) GetUdpProbe() *proto8.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_UdpProbe); ok {
		return x.UdpProbe
	}
	return nil
}

func (m *ProbeDef) GetUdpListenerProbe() *proto9.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_UdpListenerProbe); ok {
		return x.UdpListenerProbe
	}
	return nil
}

func (m *ProbeDef) GetGrpcProbe() *proto10.ProbeConf {
	if x, ok := m.GetProbe().(*ProbeDef_GrpcProbe); ok {
		return x.GrpcProbe
	}
	return nil
}

func (m *ProbeDef) GetUserDefinedProbe() string {
	if x, ok := m.GetProbe().(*ProbeDef_UserDefinedProbe); ok {
		return x.UserDefinedProbe
	}
	return ""
}

func (m *ProbeDef) GetDebugOptions() *DebugOptions {
	if m != nil {
		return m.DebugOptions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProbeDef) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProbeDef_SourceIp)(nil),
		(*ProbeDef_SourceInterface)(nil),
		(*ProbeDef_PingProbe)(nil),
		(*ProbeDef_HttpProbe)(nil),
		(*ProbeDef_DnsProbe)(nil),
		(*ProbeDef_ExternalProbe)(nil),
		(*ProbeDef_UdpProbe)(nil),
		(*ProbeDef_UdpListenerProbe)(nil),
		(*ProbeDef_GrpcProbe)(nil),
		(*ProbeDef_UserDefinedProbe)(nil),
	}
}

type AdditionalLabel struct {
	Key *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	// Value can either be a static value or can be derived from target's labels.
	// To get value from target's labels, use target.labels.<target's label key>
	// as value.
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdditionalLabel) Reset()         { *m = AdditionalLabel{} }
func (m *AdditionalLabel) String() string { return proto.CompactTextString(m) }
func (*AdditionalLabel) ProtoMessage()    {}
func (*AdditionalLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_8261ee47d309caad, []int{1}
}

func (m *AdditionalLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdditionalLabel.Unmarshal(m, b)
}
func (m *AdditionalLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdditionalLabel.Marshal(b, m, deterministic)
}
func (m *AdditionalLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdditionalLabel.Merge(m, src)
}
func (m *AdditionalLabel) XXX_Size() int {
	return xxx_messageInfo_AdditionalLabel.Size(m)
}
func (m *AdditionalLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_AdditionalLabel.DiscardUnknown(m)
}

var xxx_messageInfo_AdditionalLabel proto.InternalMessageInfo

func (m *AdditionalLabel) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *AdditionalLabel) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type DebugOptions struct {
	// Whether to log metrics or not.
	LogMetrics           *bool    `protobuf:"varint,1,opt,name=log_metrics,json=logMetrics" json:"log_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugOptions) Reset()         { *m = DebugOptions{} }
func (m *DebugOptions) String() string { return proto.CompactTextString(m) }
func (*DebugOptions) ProtoMessage()    {}
func (*DebugOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8261ee47d309caad, []int{2}
}

func (m *DebugOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugOptions.Unmarshal(m, b)
}
func (m *DebugOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugOptions.Marshal(b, m, deterministic)
}
func (m *DebugOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugOptions.Merge(m, src)
}
func (m *DebugOptions) XXX_Size() int {
	return xxx_messageInfo_DebugOptions.Size(m)
}
func (m *DebugOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DebugOptions proto.InternalMessageInfo

func (m *DebugOptions) GetLogMetrics() bool {
	if m != nil && m.LogMetrics != nil {
		return *m.LogMetrics
	}
	return false
}

func init() {
	proto.RegisterEnum("cloudprober.probes.ProbeDef_Type", ProbeDef_Type_name, ProbeDef_Type_value)
	proto.RegisterEnum("cloudprober.probes.ProbeDef_IPVersion", ProbeDef_IPVersion_name, ProbeDef_IPVersion_value)
	proto.RegisterType((*ProbeDef)(nil), "cloudprober.probes.ProbeDef")
	proto.RegisterType((*AdditionalLabel)(nil), "cloudprober.probes.AdditionalLabel")
	proto.RegisterType((*DebugOptions)(nil), "cloudprober.probes.DebugOptions")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/proto/config.proto", fileDescriptor_8261ee47d309caad)
}

var fileDescriptor_8261ee47d309caad = []byte{
	// 965 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xe1, 0x6e, 0xdb, 0x36,
	0x10, 0xc7, 0x2b, 0xc7, 0x4e, 0xac, 0xb3, 0x93, 0x08, 0x5c, 0xda, 0xaa, 0x06, 0x86, 0x6a, 0x1e,
	0xb6, 0xb9, 0x1b, 0xe0, 0x64, 0xc1, 0xd6, 0x21, 0xcd, 0x80, 0xb5, 0x8d, 0xd4, 0x46, 0x40, 0xea,
	0x08, 0x8c, 0x13, 0x6c, 0x9f, 0x04, 0x59, 0xa2, 0x55, 0x61, 0x0a, 0x25, 0x48, 0x54, 0xd6, 0x7c,
	0xcb, 0xcb, 0xec, 0x5d, 0xf6, 0x58, 0x03, 0x45, 0xca, 0xb1, 0x52, 0x25, 0x30, 0xfa, 0x49, 0xe4,
	0xf1, 0x7f, 0x3f, 0x1e, 0xef, 0x74, 0x24, 0xfc, 0x16, 0x46, 0xec, 0x63, 0x31, 0x1b, 0xfb, 0xc9,
	0xe5, 0x6e, 0x98, 0x24, 0x61, 0x4c, 0x76, 0xfd, 0x38, 0x29, 0x82, 0x34, 0x4b, 0x66, 0x24, 0xdb,
	0x2d, 0x3f, 0x39, 0xff, 0xb0, 0x64, 0xd7, 0x4f, 0xe8, 0x3c, 0x0a, 0xc7, 0xe5, 0x04, 0xa1, 0x25,
	0xd9, 0x58, 0xc8, 0x06, 0x2f, 0x1f, 0x86, 0x5d, 0x12, 0x96, 0x45, 0x7e, 0x45, 0x0b, 0xa2, 0x9c,
	0x09, 0xd6, 0xe0, 0x70, 0xa5, 0x20, 0x02, 0xda, 0x14, 0xc8, 0xe0, 0xf5, 0x4a, 0xce, 0xe4, 0x13,
	0x23, 0x19, 0xf5, 0xe2, 0x26, 0xc2, 0xef, 0x2b, 0x11, 0xc2, 0x2c, 0xf5, 0xbf, 0xdc, 0xfb, 0x23,
	0x63, 0xe9, 0x97, 0x7b, 0xa7, 0x11, 0x0d, 0x9b, 0xbc, 0x57, 0x4b, 0x5c, 0x11, 0x34, 0x6e, 0x7d,
	0xb4, 0xaa, 0x73, 0x1c, 0xe5, 0x8c, 0x50, 0x61, 0xba, 0x0b, 0x39, 0x78, 0x18, 0xc2, 0xbc, 0x2c,
	0x24, 0xac, 0xaa, 0x9b, 0x9c, 0xad, 0x16, 0xfc, 0x95, 0x17, 0x47, 0x81, 0xc7, 0x92, 0xac, 0xa9,
	0xea, 0xc3, 0x7f, 0xfb, 0xd0, 0x75, 0xb8, 0xd0, 0x24, 0x73, 0x84, 0xa0, 0x4d, 0xbd, 0x4b, 0xa2,
	0x2b, 0x46, 0x6b, 0xa4, 0xe2, 0x72, 0x8c, 0x7e, 0x85, 0x36, 0xbb, 0x4e, 0x89, 0xde, 0x32, 0x5a,
	0xa3, 0xad, 0xfd, 0x6f, 0xc6, 0x9f, 0xff, 0xae, 0xe3, 0xca, 0x7f, 0x3c, 0xbd, 0x4e, 0x09, 0x2e,
	0xe5, 0xe8, 0x31, 0xac, 0x67, 0x05, 0x75, 0x13, 0xaa, 0xaf, 0x19, 0xca, 0x48, 0xc5, 0x9d, 0xac,
	0xa0, 0xa7, 0x14, 0xbd, 0x80, 0xcd, 0x88, 0x32, 0x92, 0x5d, 0x79, 0xb1, 0x7b, 0x99, 0x13, 0x5f,
	0x6f, 0x1b, 0xca, 0xa8, 0xf3, 0xaa, 0xbd, 0xbf, 0xb7, 0xb7, 0x87, 0xfb, 0xd5, 0xd2, 0x87, 0x9c,
	0xf8, 0xe8, 0x07, 0xe8, 0xb3, 0xe8, 0x92, 0x24, 0x05, 0x13, 0xca, 0x8e, 0x50, 0xfe, 0xcc, 0x95,
	0x3d, 0xb9, 0x52, 0x0a, 0x0f, 0x60, 0x43, 0x26, 0x44, 0x5f, 0x37, 0x5a, 0xa3, 0xde, 0xfe, 0xf3,
	0x5a, 0x90, 0x55, 0xb2, 0xa6, 0xe2, 0x6b, 0x92, 0x39, 0xae, 0xf4, 0xe8, 0x04, 0x76, 0x62, 0x8f,
	0x11, 0xea, 0x5f, 0xbb, 0xbc, 0x8d, 0xb2, 0x68, 0x56, 0xb0, 0x28, 0xa1, 0xfa, 0x86, 0xa1, 0x8c,
	0x7a, 0xfb, 0xcf, 0x6a, 0x1c, 0xd9, 0x75, 0x63, 0x33, 0xca, 0x19, 0xfe, 0x4a, 0xba, 0x99, 0x4b,
	0x5e, 0xe8, 0x3b, 0xe8, 0x57, 0xb4, 0x82, 0x46, 0x4c, 0xef, 0xf2, 0x93, 0xbf, 0x6a, 0x15, 0x39,
	0xee, 0x49, 0xfb, 0x39, 0x8d, 0x18, 0xfa, 0x03, 0xd4, 0x45, 0x4d, 0x74, 0xd5, 0x58, 0x1b, 0xf5,
	0xee, 0xa4, 0xf5, 0xb6, 0x62, 0xe3, 0x8b, 0x6a, 0x88, 0x6f, 0x7d, 0xd0, 0xd7, 0xa0, 0xe6, 0x49,
	0x91, 0xf9, 0xc4, 0x8d, 0x52, 0x1d, 0xf8, 0x26, 0xc7, 0x8f, 0x70, 0x57, 0x98, 0xec, 0x14, 0xfd,
	0x04, 0x5a, 0xb5, 0xcc, 0xf3, 0x39, 0xf7, 0x7c, 0xa2, 0xf7, 0xa4, 0x6a, 0x5b, 0xaa, 0xaa, 0x05,
	0x64, 0x01, 0x44, 0xa9, 0x7b, 0x45, 0xb2, 0x9c, 0x9f, 0xbb, 0x6f, 0x28, 0xa3, 0xad, 0xfd, 0xef,
	0x1f, 0x2c, 0xb2, 0xed, 0x5c, 0x08, 0x35, 0x56, 0xa3, 0x54, 0x0e, 0xd1, 0x21, 0x0c, 0x72, 0xe6,
	0xb1, 0xdc, 0x25, 0x9f, 0xd2, 0x24, 0x63, 0x6e, 0xbd, 0xc8, 0x9b, 0xbc, 0x74, 0xf8, 0x69, 0xa9,
	0xb0, 0x4a, 0x81, 0xbd, 0x5c, 0xe9, 0x09, 0x68, 0x5e, 0x10, 0x44, 0x3c, 0x87, 0x5e, 0xec, 0xc6,
	0xde, 0x8c, 0xc4, 0xfa, 0x56, 0x99, 0x97, 0x6f, 0x9b, 0x22, 0x79, 0xb3, 0xd0, 0x9e, 0x70, 0x29,
	0xde, 0xf6, 0xea, 0x06, 0x74, 0x04, 0xc0, 0x1b, 0xdd, 0x2d, 0xf5, 0xfa, 0x4e, 0x59, 0xcb, 0x61,
	0x13, 0x89, 0xab, 0xc4, 0xc1, 0x8e, 0x12, 0x3a, 0x3f, 0x56, 0xb0, 0xca, 0x2d, 0xa5, 0x81, 0x43,
	0xf8, 0x5d, 0x23, 0x21, 0x8f, 0xef, 0x87, 0x70, 0x55, 0x1d, 0xc2, 0x2d, 0x02, 0xf2, 0x1a, 0xd4,
	0x80, 0xe6, 0x92, 0xf1, 0xa4, 0x64, 0x34, 0x76, 0x50, 0x40, 0xf3, 0x1a, 0xa2, 0x1b, 0xd0, 0x5c,
	0x10, 0x4e, 0x61, 0xab, 0xba, 0x72, 0x25, 0xe6, 0x69, 0x89, 0x69, 0xac, 0x51, 0xa5, 0xac, 0xb1,
	0x36, 0x2b, 0xeb, 0x22, 0xa4, 0x22, 0xa8, 0x8e, 0xa5, 0xdf, 0x1f, 0x52, 0x11, 0xd4, 0x4f, 0xd5,
	0x2d, 0x02, 0x79, 0xa8, 0xbf, 0x00, 0x71, 0x42, 0x75, 0x9b, 0x49, 0xd4, 0xb3, 0x12, 0xf5, 0xe2,
	0x1e, 0x54, 0x25, 0xae, 0x21, 0xb5, 0x22, 0x48, 0x4f, 0xe4, 0xc2, 0x22, 0xe9, 0xfc, 0x79, 0x90,
	0xc8, 0xc1, 0xfd, 0x49, 0xe7, 0xaa, 0x7a, 0xd2, 0xb9, 0x45, 0x40, 0xc6, 0x80, 0x8a, 0x9c, 0x64,
	0x6e, 0x40, 0xe6, 0x11, 0x25, 0x81, 0x84, 0xf9, 0x65, 0x07, 0xf0, 0x4d, 0x73, 0x92, 0x99, 0x62,
	0x49, 0xe8, 0x2d, 0xd8, 0x0c, 0xc8, 0xac, 0x08, 0xdd, 0x24, 0xe5, 0x7f, 0x51, 0xae, 0x07, 0xe5,
	0xbe, 0x46, 0xd3, 0xbe, 0x26, 0x17, 0x9e, 0x0a, 0x1d, 0xee, 0x07, 0x4b, 0xb3, 0xe1, 0x3f, 0xd0,
	0xe6, 0xf7, 0x1f, 0xea, 0x42, 0xdb, 0xb1, 0x27, 0xef, 0xb5, 0x47, 0x7c, 0x74, 0x3c, 0x9d, 0x3a,
	0x9a, 0x82, 0x36, 0x60, 0xcd, 0x9c, 0x9c, 0x69, 0x2d, 0xd4, 0x87, 0xae, 0xf5, 0xe7, 0xd4, 0xc2,
	0x93, 0x37, 0x27, 0xda, 0x1a, 0x37, 0x9f, 0x9b, 0x8e, 0xd6, 0x46, 0x1a, 0xf4, 0xcf, 0x4d, 0xc7,
	0x3d, 0xb1, 0xcf, 0xa6, 0xd6, 0xc4, 0xc2, 0x5a, 0x87, 0xfb, 0xbe, 0xc7, 0xce, 0x91, 0xb6, 0x8e,
	0x36, 0x41, 0xe5, 0x2e, 0x93, 0x33, 0xfb, 0x74, 0xa2, 0xcd, 0x4a, 0xe9, 0x99, 0x85, 0x5d, 0xd3,
	0x7a, 0x67, 0x4f, 0x2c, 0x53, 0xf3, 0x87, 0x87, 0xa0, 0x2e, 0x7a, 0x12, 0x0d, 0xe0, 0x89, 0xed,
	0xb8, 0x17, 0x16, 0xe6, 0x72, 0xf7, 0x7c, 0x72, 0xe6, 0x58, 0x47, 0xf6, 0x3b, 0xdb, 0x32, 0x45,
	0x3c, 0xb6, 0x73, 0xf1, 0x8b, 0xa6, 0xc8, 0xd1, 0x4b, 0xad, 0xf5, 0xa3, 0xda, 0xfd, 0x4f, 0xd1,
	0x6e, 0x6e, 0x6e, 0x6e, 0x5a, 0x6f, 0xd1, 0xed, 0xbd, 0x91, 0xba, 0xe2, 0x91, 0x78, 0xbb, 0x01,
	0x9d, 0xf2, 0xe4, 0xc3, 0x03, 0xd8, 0xbe, 0xd3, 0x77, 0x48, 0x83, 0xb5, 0xbf, 0xc9, 0xb5, 0x7c,
	0x2c, 0xf8, 0x10, 0xed, 0x40, 0xe7, 0xca, 0x8b, 0x0b, 0xf1, 0x58, 0xa8, 0x58, 0x4c, 0x86, 0xbb,
	0xd0, 0x5f, 0x4e, 0x1b, 0x7a, 0x0e, 0xbd, 0x38, 0x09, 0x5d, 0x79, 0x9f, 0xea, 0x8a, 0xa1, 0x8c,
	0xba, 0x18, 0xe2, 0x24, 0xfc, 0x20, 0x2c, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xc2, 0x42,
	0xbe, 0x4c, 0x09, 0x00, 0x00,
}
