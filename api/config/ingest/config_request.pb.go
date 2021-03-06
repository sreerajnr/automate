// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/ingest/config_request.proto

package ingest

import (
	fmt "fmt"
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0}
}

func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(m, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Log                  *ConfigRequest_V1_System_Log     `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 0}
}

func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(m, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	// NOTE: purge actions are no longer configurable via config
	// and are now managed with gRPC endpoints. They are not
	// reserved so we can migrate the initial values from config to
	// to the purge cereal workflows.
	// Setting these values is prevented in the Validate() callback.
	PurgeConvergeHistoryAfterDays         *wrappers.Int32Value `protobuf:"bytes,3,opt,name=purge_converge_history_after_days,json=purgeConvergeHistoryAfterDays,proto3" json:"purge_converge_history_after_days,omitempty" toml:"purge_converge_history_after_days,omitempty" mapstructure:"purge_converge_history_after_days,omitempty"` // Deprecated: Do not use.
	PurgeActionsAfterDays                 *wrappers.Int32Value `protobuf:"bytes,4,opt,name=purge_actions_after_days,json=purgeActionsAfterDays,proto3" json:"purge_actions_after_days,omitempty" toml:"purge_actions_after_days,omitempty" mapstructure:"purge_actions_after_days,omitempty"`                           // Deprecated: Do not use.
	MaxNumberOfBundledRunMsgs             *wrappers.Int32Value `protobuf:"bytes,6,opt,name=max_number_of_bundled_run_msgs,json=maxNumberOfBundledRunMsgs,proto3" json:"max_number_of_bundled_run_msgs,omitempty" toml:"max_number_of_bundled_run_msgs,omitempty" mapstructure:"max_number_of_bundled_run_msgs,omitempty"`
	MaxNumberOfBundledActionMsgs          *wrappers.Int32Value `protobuf:"bytes,7,opt,name=max_number_of_bundled_action_msgs,json=maxNumberOfBundledActionMsgs,proto3" json:"max_number_of_bundled_action_msgs,omitempty" toml:"max_number_of_bundled_action_msgs,omitempty" mapstructure:"max_number_of_bundled_action_msgs,omitempty"`
	NumberOfRunMsgsTransformers           *wrappers.Int32Value `protobuf:"bytes,8,opt,name=number_of_run_msgs_transformers,json=numberOfRunMsgsTransformers,proto3" json:"number_of_run_msgs_transformers,omitempty" toml:"number_of_run_msgs_transformers,omitempty" mapstructure:"number_of_run_msgs_transformers,omitempty"`
	NumberOfRunMsgPublishers              *wrappers.Int32Value `protobuf:"bytes,9,opt,name=number_of_run_msg_publishers,json=numberOfRunMsgPublishers,proto3" json:"number_of_run_msg_publishers,omitempty" toml:"number_of_run_msg_publishers,omitempty" mapstructure:"number_of_run_msg_publishers,omitempty"`
	MessageBufferSize                     *wrappers.Int32Value `protobuf:"bytes,10,opt,name=message_buffer_size,json=messageBufferSize,proto3" json:"message_buffer_size,omitempty" toml:"message_buffer_size,omitempty" mapstructure:"message_buffer_size,omitempty"`
	NodesMissingRunningDefault            *wrappers.BoolValue  `protobuf:"bytes,11,opt,name=nodes_missing_running_default,json=nodesMissingRunningDefault,proto3" json:"nodes_missing_running_default,omitempty" toml:"nodes_missing_running_default,omitempty" mapstructure:"nodes_missing_running_default,omitempty"`
	MissingNodesForDeletionRunningDefault *wrappers.BoolValue  `protobuf:"bytes,12,opt,name=missing_nodes_for_deletion_running_default,json=missingNodesForDeletionRunningDefault,proto3" json:"missing_nodes_for_deletion_running_default,omitempty" toml:"missing_nodes_for_deletion_running_default,omitempty" mapstructure:"missing_nodes_for_deletion_running_default,omitempty"`
	XXX_NoUnkeyedLiteral                  struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized                      []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache                         int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

// Deprecated: Do not use.
func (m *ConfigRequest_V1_System_Service) GetPurgeConvergeHistoryAfterDays() *wrappers.Int32Value {
	if m != nil {
		return m.PurgeConvergeHistoryAfterDays
	}
	return nil
}

// Deprecated: Do not use.
func (m *ConfigRequest_V1_System_Service) GetPurgeActionsAfterDays() *wrappers.Int32Value {
	if m != nil {
		return m.PurgeActionsAfterDays
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetMaxNumberOfBundledRunMsgs() *wrappers.Int32Value {
	if m != nil {
		return m.MaxNumberOfBundledRunMsgs
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetMaxNumberOfBundledActionMsgs() *wrappers.Int32Value {
	if m != nil {
		return m.MaxNumberOfBundledActionMsgs
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetNumberOfRunMsgsTransformers() *wrappers.Int32Value {
	if m != nil {
		return m.NumberOfRunMsgsTransformers
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetNumberOfRunMsgPublishers() *wrappers.Int32Value {
	if m != nil {
		return m.NumberOfRunMsgPublishers
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetMessageBufferSize() *wrappers.Int32Value {
	if m != nil {
		return m.MessageBufferSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetNodesMissingRunningDefault() *wrappers.BoolValue {
	if m != nil {
		return m.NodesMissingRunningDefault
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetMissingNodesForDeletionRunningDefault() *wrappers.BoolValue {
	if m != nil {
		return m.MissingNodesForDeletionRunningDefault
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Format               *wrappers.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	Level                *wrappers.StringValue `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8c399c32386e790, []int{0, 0, 1}
}

func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.ingest.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.ingest.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.ingest.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/ingest/config_request.proto", fileDescriptor_a8c399c32386e790)
}

var fileDescriptor_a8c399c32386e790 = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x6f, 0x6f, 0xdc, 0x34,
	0x1c, 0xc7, 0xd5, 0xbb, 0xec, 0xda, 0xb9, 0x1b, 0x2a, 0xe6, 0x8f, 0x42, 0xda, 0x95, 0x0d, 0x09,
	0x84, 0x2a, 0x2e, 0xa1, 0xb7, 0x22, 0x21, 0x34, 0x1e, 0xec, 0x5a, 0x6d, 0x50, 0xda, 0x0d, 0xa5,
	0x53, 0x85, 0x40, 0x10, 0x39, 0xc9, 0x2f, 0x3e, 0x4b, 0x8e, 0x1d, 0x6c, 0xe7, 0xd8, 0xed, 0x25,
	0x20, 0x5e, 0x02, 0xef, 0x81, 0xf7, 0xb0, 0xb7, 0xb4, 0x27, 0x3c, 0x44, 0xb1, 0x7d, 0xb7, 0x76,
	0x87, 0x7a, 0xdb, 0x1e, 0xdd, 0x45, 0xf6, 0xe7, 0xf3, 0xfd, 0xf9, 0x17, 0xdb, 0x41, 0x9f, 0x91,
	0x86, 0x25, 0x85, 0x14, 0x15, 0xa3, 0x09, 0x13, 0x14, 0xb4, 0xf1, 0x4f, 0x99, 0x82, 0xdf, 0x5b,
	0xd0, 0x26, 0x6e, 0x94, 0x34, 0x12, 0x6f, 0x17, 0x13, 0xa8, 0x62, 0xd2, 0x1a, 0x59, 0x13, 0x03,
	0x71, 0x29, 0x6b, 0xc2, 0x44, 0xec, 0x88, 0x68, 0xf7, 0x82, 0x44, 0x4f, 0x88, 0x82, 0x32, 0xa1,
	0x5c, 0xe6, 0x84, 0x3b, 0x38, 0xda, 0x5e, 0x1e, 0x37, 0x5c, 0xfb, 0xc1, 0xe3, 0x42, 0xd6, 0x8d,
	0x14, 0x20, 0x8c, 0x4e, 0xe6, 0xfe, 0x21, 0x55, 0x4d, 0x91, 0xd8, 0xf1, 0x62, 0x48, 0x41, 0x0c,
	0xc9, 0x68, 0xe8, 0xf9, 0x4e, 0x45, 0x46, 0xdd, 0x43, 0x42, 0x84, 0x90, 0x86, 0x18, 0x26, 0xc5,
	0xdc, 0xb5, 0x4b, 0xa5, 0xa4, 0x1c, 0x1c, 0x99, 0xb7, 0x55, 0xf2, 0x87, 0x22, 0x4d, 0x03, 0xca,
	0x8f, 0x7f, 0xf2, 0xf7, 0x4d, 0x74, 0xf3, 0xd0, 0x7a, 0x52, 0xb7, 0x3a, 0xfc, 0x2d, 0xea, 0x4d,
	0xf7, 0xc3, 0xfe, 0xed, 0xb5, 0xcf, 0x37, 0x47, 0xc3, 0xf8, 0x8a, 0x45, 0xc6, 0x97, 0xb8, 0xf8,
	0x7c, 0x3f, 0xed, 0x4d, 0xf7, 0xa3, 0x7f, 0x6e, 0xa0, 0xde, 0xf9, 0x3e, 0x7e, 0x80, 0xfa, 0x7a,
	0xa6, 0xc3, 0x35, 0xab, 0x39, 0x78, 0x23, 0x4d, 0x7c, 0x36, 0xd3, 0x06, 0xea, 0xb4, 0x13, 0xe0,
	0x87, 0xa8, 0xaf, 0xa7, 0x45, 0xd8, 0xb3, 0x9e, 0xaf, 0xde, 0xd0, 0x03, 0x6a, 0xca, 0x0a, 0x48,
	0x3b, 0x43, 0xf4, 0xd7, 0x26, 0x1a, 0x38, 0x31, 0x3e, 0x40, 0x41, 0xcd, 0x35, 0xf1, 0xc5, 0xdd,
	0x7e, 0x45, 0xca, 0x44, 0xa5, 0x48, 0xec, 0x7a, 0x1b, 0x9f, 0x72, 0x4d, 0x52, 0x3b, 0x1b, 0x9f,
	0xa3, 0x75, 0xed, 0x84, 0xbe, 0x9a, 0x7b, 0x6f, 0xb3, 0xaa, 0x45, 0x51, 0x73, 0x19, 0xbe, 0x87,
	0xfa, 0x86, 0x6b, 0xdf, 0xf0, 0xbd, 0xab, 0x8a, 0x79, 0x72, 0x72, 0x76, 0xa8, 0xa0, 0x04, 0x61,
	0x18, 0xe1, 0x3a, 0xed, 0x30, 0x7c, 0x8c, 0xfa, 0x5c, 0xd2, 0x30, 0xb0, 0xf4, 0xd7, 0x6f, 0x55,
	0xd1, 0x89, 0xa4, 0x69, 0x27, 0x89, 0xfe, 0x5d, 0x47, 0xeb, 0xbe, 0x3c, 0xfc, 0x25, 0x0a, 0x26,
	0x52, 0x1b, 0xdf, 0xa3, 0x9d, 0xd8, 0x6d, 0xa3, 0x78, 0xbe, 0x8d, 0xe2, 0x33, 0xa3, 0x98, 0xa0,
	0xe7, 0x84, 0xb7, 0x90, 0xda, 0x99, 0xf8, 0x21, 0x0a, 0x1a, 0xa9, 0x8c, 0x6f, 0xce, 0xf6, 0x12,
	0xf1, 0xbd, 0x30, 0x77, 0x47, 0x16, 0x18, 0x7f, 0xf8, 0xfc, 0x45, 0x88, 0x17, 0xed, 0xdc, 0xfa,
	0xf3, 0x71, 0x14, 0x74, 0xdb, 0x3b, 0xb5, 0x02, 0xcc, 0xd0, 0x9d, 0xa6, 0x55, 0x14, 0xb2, 0x42,
	0x8a, 0x29, 0x74, 0x7f, 0x26, 0x4c, 0x1b, 0xa9, 0x66, 0x19, 0xa9, 0x0c, 0xa8, 0xac, 0x24, 0xb3,
	0x79, 0xbb, 0xae, 0x4c, 0xe9, 0x85, 0x6b, 0xe9, 0x2d, 0x6b, 0x3a, 0xf4, 0xa2, 0xef, 0x9c, 0xe7,
	0x7e, 0xa7, 0x39, 0x22, 0x33, 0x8d, 0x7f, 0x42, 0xa1, 0x8b, 0x22, 0x85, 0x3d, 0x34, 0x17, 0x13,
	0x82, 0xd7, 0x4b, 0xf8, 0xc0, 0x0a, 0xee, 0x3b, 0xfe, 0xa5, 0xf9, 0x37, 0xb4, 0x5b, 0x93, 0xa7,
	0x99, 0x68, 0xeb, 0x1c, 0x54, 0x26, 0xab, 0x2c, 0x6f, 0x45, 0xc9, 0xa1, 0xcc, 0x54, 0x2b, 0xb2,
	0x5a, 0x53, 0x1d, 0x0e, 0x56, 0xfa, 0xd3, 0x8f, 0x6a, 0xf2, 0xf4, 0x91, 0x35, 0x3c, 0xae, 0xc6,
	0x8e, 0x4f, 0x5b, 0x71, 0xaa, 0xa9, 0xc6, 0x25, 0xba, 0xf3, 0xff, 0x7e, 0xb7, 0x12, 0x17, 0xb1,
	0xbe, 0x3a, 0x62, 0x67, 0x39, 0xc2, 0xad, 0xc5, 0xa6, 0x10, 0xf4, 0xf1, 0xcb, 0x84, 0x79, 0xe5,
	0x99, 0x51, 0x44, 0xe8, 0x4a, 0xaa, 0x1a, 0x94, 0x0e, 0x37, 0x56, 0x67, 0x6c, 0x0b, 0x1f, 0xe0,
	0x8b, 0x7f, 0x72, 0x81, 0xc7, 0xbf, 0xa0, 0x9d, 0xa5, 0x88, 0xac, 0x69, 0x73, 0xce, 0xf4, 0xa4,
	0xf3, 0x5f, 0x5f, 0xed, 0x0f, 0x2f, 0xfb, 0x7f, 0x5c, 0xc0, 0xf8, 0x07, 0xf4, 0x5e, 0x0d, 0x5a,
	0x13, 0x0a, 0x59, 0xde, 0x56, 0x15, 0xa8, 0x4c, 0xb3, 0x67, 0x10, 0xa2, 0xd5, 0xce, 0x77, 0x3d,
	0x37, 0xb6, 0xd8, 0x19, 0x7b, 0x06, 0xf8, 0x57, 0x74, 0x4b, 0xc8, 0x12, 0x74, 0x56, 0x33, 0xad,
	0x99, 0xa0, 0x5d, 0xb5, 0xa2, 0xfb, 0x2d, 0xa1, 0x22, 0x2d, 0x37, 0xe1, 0xa6, 0xd5, 0x46, 0x4b,
	0xda, 0xb1, 0x94, 0xdc, 0x59, 0x23, 0x2b, 0x38, 0x75, 0x7c, 0xea, 0xf0, 0x23, 0x47, 0xe3, 0x16,
	0xed, 0xcd, 0xc5, 0x2e, 0xa6, 0x92, 0x2a, 0x2b, 0x81, 0x83, 0x7d, 0x9f, 0xaf, 0x66, 0xdd, 0x58,
	0x99, 0xf5, 0xa9, 0xb7, 0x3d, 0xea, 0x64, 0x0f, 0xa4, 0x3a, 0xf2, 0xaa, 0xcb, 0xb1, 0xc7, 0xc1,
	0xc6, 0xb5, 0xad, 0x41, 0x24, 0x51, 0xff, 0x44, 0x52, 0x7c, 0x80, 0x06, 0xdd, 0x7b, 0x21, 0xaf,
	0x77, 0xee, 0xfd, 0x5c, 0x3c, 0x42, 0xd7, 0x38, 0x4c, 0x81, 0xfb, 0xa3, 0x7f, 0x35, 0xe4, 0xa6,
	0x46, 0xd7, 0x17, 0x57, 0xcd, 0x37, 0xef, 0x3f, 0x7f, 0x11, 0x6e, 0xa1, 0x77, 0xdc, 0x2d, 0x35,
	0xf4, 0x97, 0xc2, 0x71, 0xb0, 0xb1, 0xb6, 0xd5, 0x1f, 0x7f, 0xf1, 0xf3, 0x1e, 0x65, 0x66, 0xd2,
	0xe6, 0x71, 0x21, 0xeb, 0xa4, 0xbb, 0xdd, 0x16, 0x5f, 0xc4, 0x64, 0xe9, 0x3b, 0x9d, 0x0f, 0x6c,
	0xe2, 0xdd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x5d, 0xe8, 0x1a, 0xc3, 0x07, 0x00, 0x00,
}
