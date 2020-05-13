// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/request/orgs.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateOrg struct {
	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef organization name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef organization admin user.
	AdminUser string `protobuf:"bytes,3,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty" toml:"admin_user,omitempty" mapstructure:"admin_user,omitempty"`
	// Chef organization admin key.
	AdminKey string `protobuf:"bytes,4,opt,name=admin_key,json=adminKey,proto3" json:"admin_key,omitempty" toml:"admin_key,omitempty" mapstructure:"admin_key,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// List of projects this chef organization belongs to. May be empty.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateOrg) Reset()         { *m = CreateOrg{} }
func (m *CreateOrg) String() string { return proto.CompactTextString(m) }
func (*CreateOrg) ProtoMessage()    {}
func (*CreateOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2068df3348ec5f9, []int{0}
}

func (m *CreateOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrg.Unmarshal(m, b)
}
func (m *CreateOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrg.Marshal(b, m, deterministic)
}
func (m *CreateOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrg.Merge(m, src)
}
func (m *CreateOrg) XXX_Size() int {
	return xxx_messageInfo_CreateOrg.Size(m)
}
func (m *CreateOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrg proto.InternalMessageInfo

func (m *CreateOrg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrg) GetAdminUser() string {
	if m != nil {
		return m.AdminUser
	}
	return ""
}

func (m *CreateOrg) GetAdminKey() string {
	if m != nil {
		return m.AdminKey
	}
	return ""
}

func (m *CreateOrg) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *CreateOrg) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type UpdateOrg struct {
	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef organization name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef organization admin user.
	AdminUser string `protobuf:"bytes,3,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty" toml:"admin_user,omitempty" mapstructure:"admin_user,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,4,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// List of projects this chef organization belongs to. May be empty.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateOrg) Reset()         { *m = UpdateOrg{} }
func (m *UpdateOrg) String() string { return proto.CompactTextString(m) }
func (*UpdateOrg) ProtoMessage()    {}
func (*UpdateOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2068df3348ec5f9, []int{1}
}

func (m *UpdateOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOrg.Unmarshal(m, b)
}
func (m *UpdateOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOrg.Marshal(b, m, deterministic)
}
func (m *UpdateOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOrg.Merge(m, src)
}
func (m *UpdateOrg) XXX_Size() int {
	return xxx_messageInfo_UpdateOrg.Size(m)
}
func (m *UpdateOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOrg.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOrg proto.InternalMessageInfo

func (m *UpdateOrg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateOrg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateOrg) GetAdminUser() string {
	if m != nil {
		return m.AdminUser
	}
	return ""
}

func (m *UpdateOrg) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *UpdateOrg) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeleteOrg struct {
	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef Infra Server ID.
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteOrg) Reset()         { *m = DeleteOrg{} }
func (m *DeleteOrg) String() string { return proto.CompactTextString(m) }
func (*DeleteOrg) ProtoMessage()    {}
func (*DeleteOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2068df3348ec5f9, []int{2}
}

func (m *DeleteOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOrg.Unmarshal(m, b)
}
func (m *DeleteOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOrg.Marshal(b, m, deterministic)
}
func (m *DeleteOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOrg.Merge(m, src)
}
func (m *DeleteOrg) XXX_Size() int {
	return xxx_messageInfo_DeleteOrg.Size(m)
}
func (m *DeleteOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOrg.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOrg proto.InternalMessageInfo

func (m *DeleteOrg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteOrg) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type GetOrgs struct {
	// Chef Infra Server ID.
	ServerId             string   `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetOrgs) Reset()         { *m = GetOrgs{} }
func (m *GetOrgs) String() string { return proto.CompactTextString(m) }
func (*GetOrgs) ProtoMessage()    {}
func (*GetOrgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2068df3348ec5f9, []int{3}
}

func (m *GetOrgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrgs.Unmarshal(m, b)
}
func (m *GetOrgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrgs.Marshal(b, m, deterministic)
}
func (m *GetOrgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrgs.Merge(m, src)
}
func (m *GetOrgs) XXX_Size() int {
	return xxx_messageInfo_GetOrgs.Size(m)
}
func (m *GetOrgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrgs proto.InternalMessageInfo

func (m *GetOrgs) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type GetOrg struct {
	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef Infra Server ID.
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *GetOrg) Reset()         { *m = GetOrg{} }
func (m *GetOrg) String() string { return proto.CompactTextString(m) }
func (*GetOrg) ProtoMessage()    {}
func (*GetOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2068df3348ec5f9, []int{4}
}

func (m *GetOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrg.Unmarshal(m, b)
}
func (m *GetOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrg.Marshal(b, m, deterministic)
}
func (m *GetOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrg.Merge(m, src)
}
func (m *GetOrg) XXX_Size() int {
	return xxx_messageInfo_GetOrg.Size(m)
}
func (m *GetOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrg.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrg proto.InternalMessageInfo

func (m *GetOrg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetOrg) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type ResetOrgAdminKey struct {
	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Chef organization admin key.
	AdminKey             string   `protobuf:"bytes,3,opt,name=admin_key,json=adminKey,proto3" json:"admin_key,omitempty" toml:"admin_key,omitempty" mapstructure:"admin_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ResetOrgAdminKey) Reset()         { *m = ResetOrgAdminKey{} }
func (m *ResetOrgAdminKey) String() string { return proto.CompactTextString(m) }
func (*ResetOrgAdminKey) ProtoMessage()    {}
func (*ResetOrgAdminKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2068df3348ec5f9, []int{5}
}

func (m *ResetOrgAdminKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetOrgAdminKey.Unmarshal(m, b)
}
func (m *ResetOrgAdminKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetOrgAdminKey.Marshal(b, m, deterministic)
}
func (m *ResetOrgAdminKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetOrgAdminKey.Merge(m, src)
}
func (m *ResetOrgAdminKey) XXX_Size() int {
	return xxx_messageInfo_ResetOrgAdminKey.Size(m)
}
func (m *ResetOrgAdminKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetOrgAdminKey.DiscardUnknown(m)
}

var xxx_messageInfo_ResetOrgAdminKey proto.InternalMessageInfo

func (m *ResetOrgAdminKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResetOrgAdminKey) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *ResetOrgAdminKey) GetAdminKey() string {
	if m != nil {
		return m.AdminKey
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateOrg)(nil), "chef.automate.domain.infra_proxy.request.CreateOrg")
	proto.RegisterType((*UpdateOrg)(nil), "chef.automate.domain.infra_proxy.request.UpdateOrg")
	proto.RegisterType((*DeleteOrg)(nil), "chef.automate.domain.infra_proxy.request.DeleteOrg")
	proto.RegisterType((*GetOrgs)(nil), "chef.automate.domain.infra_proxy.request.GetOrgs")
	proto.RegisterType((*GetOrg)(nil), "chef.automate.domain.infra_proxy.request.GetOrg")
	proto.RegisterType((*ResetOrgAdminKey)(nil), "chef.automate.domain.infra_proxy.request.ResetOrgAdminKey")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/request/orgs.proto", fileDescriptor_c2068df3348ec5f9)
}

var fileDescriptor_c2068df3348ec5f9 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x69, 0xad, 0xcd, 0x1c, 0x44, 0xf6, 0x14, 0x14, 0xa1, 0xf4, 0x20, 0x3d, 0xed, 0x1e,
	0x44, 0xf0, 0x22, 0xe2, 0x07, 0x88, 0x78, 0x28, 0x14, 0x7a, 0x11, 0x21, 0x6c, 0xb3, 0xd3, 0x74,
	0xd5, 0x64, 0xe3, 0xec, 0x46, 0xec, 0x2f, 0xf0, 0x8f, 0xf8, 0x43, 0x25, 0x1b, 0x8b, 0x46, 0x3c,
	0x18, 0xf0, 0x96, 0x79, 0x6f, 0xde, 0xe3, 0x65, 0xe7, 0x81, 0x90, 0xa5, 0x16, 0xba, 0x70, 0x48,
	0x16, 0xe9, 0x45, 0xa7, 0x28, 0x74, 0xb1, 0x24, 0x99, 0x94, 0x64, 0x5e, 0xd7, 0x82, 0xf0, 0xb9,
	0x42, 0xeb, 0x84, 0xa1, 0xcc, 0xf2, 0x92, 0x8c, 0x33, 0x6c, 0x92, 0xae, 0x70, 0xc9, 0x65, 0xe5,
	0x4c, 0x2e, 0x1d, 0x72, 0x65, 0x72, 0xa9, 0x0b, 0xfe, 0x4d, 0xc4, 0x3f, 0x45, 0xe3, 0xf7, 0x00,
	0xa2, 0x4b, 0x42, 0xe9, 0x70, 0x4a, 0x19, 0xdb, 0x81, 0x50, 0xab, 0x38, 0x18, 0x05, 0x93, 0x68,
	0x16, 0x6a, 0xc5, 0x18, 0xf4, 0x0b, 0x99, 0x63, 0x1c, 0x7a, 0xc4, 0x7f, 0xb3, 0x03, 0x00, 0xa9,
	0x72, 0x5d, 0x24, 0x95, 0x45, 0x8a, 0x7b, 0x9e, 0x89, 0x3c, 0x32, 0xb7, 0x48, 0x6c, 0x1f, 0x9a,
	0x21, 0x79, 0xc4, 0x75, 0xdc, 0xf7, 0xec, 0xd0, 0x03, 0xb7, 0xb8, 0xae, 0xc9, 0xfa, 0x0f, 0x90,
	0x12, 0xad, 0xe2, 0xad, 0x86, 0x6c, 0x80, 0x1b, 0xc5, 0xf6, 0x60, 0x58, 0x92, 0x79, 0xc0, 0xd4,
	0xd9, 0x78, 0x30, 0xea, 0xd5, 0xdc, 0x66, 0x1e, 0xbf, 0x05, 0x10, 0xcd, 0x4b, 0xf5, 0xaf, 0x31,
	0xbf, 0x92, 0xf4, 0x3b, 0x24, 0x39, 0x81, 0xe8, 0x0a, 0x9f, 0xf0, 0xf7, 0x20, 0x2d, 0xd7, 0xb0,
	0xed, 0x3a, 0x3e, 0x84, 0xed, 0x6b, 0x74, 0x53, 0xca, 0x6c, 0x7b, 0x2f, 0xf8, 0xb1, 0x77, 0x0c,
	0x83, 0x66, 0xaf, 0x9b, 0xfd, 0x3d, 0xec, 0xce, 0xd0, 0x7a, 0xe1, 0xf9, 0xe6, 0xbd, 0xbb, 0x18,
	0xb4, 0x2f, 0xd7, 0x6b, 0x5f, 0xee, 0xe2, 0xec, 0xee, 0x34, 0xd3, 0x6e, 0x55, 0x2d, 0x78, 0x6a,
	0x72, 0x51, 0xd7, 0x4b, 0x6c, 0xea, 0xf5, 0xa7, 0x76, 0x2e, 0x06, 0xbe, 0x99, 0x47, 0x1f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x09, 0xeb, 0x70, 0xcc, 0x02, 0x00, 0x00,
}
