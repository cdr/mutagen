// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/forwarding/forwarding.proto

package forwarding

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	forwarding "github.com/mutagen-io/mutagen/pkg/forwarding"
	selection "github.com/mutagen-io/mutagen/pkg/selection"
	url "github.com/mutagen-io/mutagen/pkg/url"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreationSpecification struct {
	// Source is the source endpoint URL for the session.
	Source *url.URL `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Destination is the destination endpoint URL for the session.
	Destination *url.URL `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Configuration is the base session configuration. It is the result of
	// merging the global configuration (unless disabled), any manually
	// specified configuration file, and any command line configuration
	// parameters.
	Configuration *forwarding.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// ConfigurationSource is the source-specific session configuration. It is
	// determined based on command line configuration parameters.
	ConfigurationSource *forwarding.Configuration `protobuf:"bytes,4,opt,name=configurationSource,proto3" json:"configurationSource,omitempty"`
	// ConfigurationDestination is the destination-specific session
	// configuration. It is determined based on command line configuration
	// parameters.
	ConfigurationDestination *forwarding.Configuration `protobuf:"bytes,5,opt,name=configurationDestination,proto3" json:"configurationDestination,omitempty"`
	// Name is the name for the session object.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Labels are the labels for the session object.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not to create the session pre-paused.
	Paused               bool     `protobuf:"varint,8,opt,name=paused,proto3" json:"paused,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreationSpecification) Reset()         { *m = CreationSpecification{} }
func (m *CreationSpecification) String() string { return proto.CompactTextString(m) }
func (*CreationSpecification) ProtoMessage()    {}
func (*CreationSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{0}
}

func (m *CreationSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreationSpecification.Unmarshal(m, b)
}
func (m *CreationSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreationSpecification.Marshal(b, m, deterministic)
}
func (m *CreationSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreationSpecification.Merge(m, src)
}
func (m *CreationSpecification) XXX_Size() int {
	return xxx_messageInfo_CreationSpecification.Size(m)
}
func (m *CreationSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_CreationSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_CreationSpecification proto.InternalMessageInfo

func (m *CreationSpecification) GetSource() *url.URL {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *CreationSpecification) GetDestination() *url.URL {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *CreationSpecification) GetConfiguration() *forwarding.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *CreationSpecification) GetConfigurationSource() *forwarding.Configuration {
	if m != nil {
		return m.ConfigurationSource
	}
	return nil
}

func (m *CreationSpecification) GetConfigurationDestination() *forwarding.Configuration {
	if m != nil {
		return m.ConfigurationDestination
	}
	return nil
}

func (m *CreationSpecification) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreationSpecification) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreationSpecification) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

type CreateRequest struct {
	Specification        *CreationSpecification `protobuf:"bytes,1,opt,name=specification,proto3" json:"specification,omitempty"`
	Response             string                 `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetSpecification() *CreationSpecification {
	if m != nil {
		return m.Specification
	}
	return nil
}

func (m *CreateRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type CreateResponse struct {
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string   `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type ListRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	PreviousStateIndex   uint64               `protobuf:"varint,2,opt,name=previousStateIndex,proto3" json:"previousStateIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{3}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

func (m *ListRequest) GetPreviousStateIndex() uint64 {
	if m != nil {
		return m.PreviousStateIndex
	}
	return 0
}

type ListResponse struct {
	StateIndex           uint64              `protobuf:"varint,1,opt,name=stateIndex,proto3" json:"stateIndex,omitempty"`
	SessionStates        []*forwarding.State `protobuf:"bytes,2,rep,name=sessionStates,proto3" json:"sessionStates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{4}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetStateIndex() uint64 {
	if m != nil {
		return m.StateIndex
	}
	return 0
}

func (m *ListResponse) GetSessionStates() []*forwarding.State {
	if m != nil {
		return m.SessionStates
	}
	return nil
}

type PauseRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PauseRequest) Reset()         { *m = PauseRequest{} }
func (m *PauseRequest) String() string { return proto.CompactTextString(m) }
func (*PauseRequest) ProtoMessage()    {}
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{5}
}

func (m *PauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseRequest.Unmarshal(m, b)
}
func (m *PauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseRequest.Marshal(b, m, deterministic)
}
func (m *PauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseRequest.Merge(m, src)
}
func (m *PauseRequest) XXX_Size() int {
	return xxx_messageInfo_PauseRequest.Size(m)
}
func (m *PauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PauseRequest proto.InternalMessageInfo

func (m *PauseRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

type PauseResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseResponse) Reset()         { *m = PauseResponse{} }
func (m *PauseResponse) String() string { return proto.CompactTextString(m) }
func (*PauseResponse) ProtoMessage()    {}
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{6}
}

func (m *PauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseResponse.Unmarshal(m, b)
}
func (m *PauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseResponse.Marshal(b, m, deterministic)
}
func (m *PauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseResponse.Merge(m, src)
}
func (m *PauseResponse) XXX_Size() int {
	return xxx_messageInfo_PauseResponse.Size(m)
}
func (m *PauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PauseResponse proto.InternalMessageInfo

func (m *PauseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResumeRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	Response             string               `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResumeRequest) Reset()         { *m = ResumeRequest{} }
func (m *ResumeRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeRequest) ProtoMessage()    {}
func (*ResumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{7}
}

func (m *ResumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRequest.Unmarshal(m, b)
}
func (m *ResumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRequest.Marshal(b, m, deterministic)
}
func (m *ResumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRequest.Merge(m, src)
}
func (m *ResumeRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeRequest.Size(m)
}
func (m *ResumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRequest proto.InternalMessageInfo

func (m *ResumeRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

func (m *ResumeRequest) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type ResumeResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Prompt               string   `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeResponse) Reset()         { *m = ResumeResponse{} }
func (m *ResumeResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeResponse) ProtoMessage()    {}
func (*ResumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{8}
}

func (m *ResumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeResponse.Unmarshal(m, b)
}
func (m *ResumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeResponse.Marshal(b, m, deterministic)
}
func (m *ResumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeResponse.Merge(m, src)
}
func (m *ResumeResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeResponse.Size(m)
}
func (m *ResumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeResponse proto.InternalMessageInfo

func (m *ResumeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResumeResponse) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type TerminateRequest struct {
	Selection            *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TerminateRequest) Reset()         { *m = TerminateRequest{} }
func (m *TerminateRequest) String() string { return proto.CompactTextString(m) }
func (*TerminateRequest) ProtoMessage()    {}
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{9}
}

func (m *TerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateRequest.Unmarshal(m, b)
}
func (m *TerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateRequest.Marshal(b, m, deterministic)
}
func (m *TerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateRequest.Merge(m, src)
}
func (m *TerminateRequest) XXX_Size() int {
	return xxx_messageInfo_TerminateRequest.Size(m)
}
func (m *TerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateRequest proto.InternalMessageInfo

func (m *TerminateRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

type TerminateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateResponse) Reset()         { *m = TerminateResponse{} }
func (m *TerminateResponse) String() string { return proto.CompactTextString(m) }
func (*TerminateResponse) ProtoMessage()    {}
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{10}
}

func (m *TerminateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateResponse.Unmarshal(m, b)
}
func (m *TerminateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateResponse.Marshal(b, m, deterministic)
}
func (m *TerminateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateResponse.Merge(m, src)
}
func (m *TerminateResponse) XXX_Size() int {
	return xxx_messageInfo_TerminateResponse.Size(m)
}
func (m *TerminateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateResponse proto.InternalMessageInfo

func (m *TerminateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreationSpecification)(nil), "forwarding.CreationSpecification")
	proto.RegisterMapType((map[string]string)(nil), "forwarding.CreationSpecification.LabelsEntry")
	proto.RegisterType((*CreateRequest)(nil), "forwarding.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "forwarding.CreateResponse")
	proto.RegisterType((*ListRequest)(nil), "forwarding.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "forwarding.ListResponse")
	proto.RegisterType((*PauseRequest)(nil), "forwarding.PauseRequest")
	proto.RegisterType((*PauseResponse)(nil), "forwarding.PauseResponse")
	proto.RegisterType((*ResumeRequest)(nil), "forwarding.ResumeRequest")
	proto.RegisterType((*ResumeResponse)(nil), "forwarding.ResumeResponse")
	proto.RegisterType((*TerminateRequest)(nil), "forwarding.TerminateRequest")
	proto.RegisterType((*TerminateResponse)(nil), "forwarding.TerminateResponse")
}

func init() {
	proto.RegisterFile("service/forwarding/forwarding.proto", fileDescriptor_3507425a8852e9f1)
}

var fileDescriptor_3507425a8852e9f1 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5b, 0x6f, 0xd3, 0x4a,
	0x10, 0x3e, 0x4e, 0xd2, 0x34, 0x9e, 0xd4, 0x55, 0xbb, 0xa7, 0xa7, 0xc7, 0xb1, 0xce, 0xa9, 0x82,
	0x79, 0x09, 0x48, 0x75, 0x50, 0x40, 0x5c, 0x1f, 0x90, 0x52, 0x5a, 0x84, 0xa8, 0x10, 0xda, 0xd0,
	0x17, 0x84, 0x84, 0x5c, 0x67, 0x6a, 0xac, 0xc6, 0x97, 0xee, 0xda, 0x85, 0xfe, 0x66, 0xfe, 0x00,
	0x8f, 0xc8, 0x6b, 0x6f, 0xb2, 0xdb, 0xa6, 0xe4, 0xa1, 0x6f, 0x3b, 0x33, 0xdf, 0x5c, 0xbf, 0x19,
	0x2d, 0xdc, 0xe7, 0xc8, 0x2e, 0xa3, 0x00, 0x87, 0x67, 0x29, 0xfb, 0xee, 0xb3, 0x69, 0x94, 0x84,
	0xca, 0xd3, 0xcb, 0x58, 0x9a, 0xa7, 0x04, 0x16, 0x1a, 0xa7, 0xc7, 0x71, 0x86, 0x41, 0x1e, 0xa5,
	0xc9, 0x70, 0xfe, 0xaa, 0x60, 0xce, 0x9e, 0x12, 0x23, 0x48, 0x93, 0xb3, 0x28, 0x2c, 0x98, 0xaf,
	0xd8, 0x77, 0x15, 0x3b, 0xcf, 0xfd, 0x1c, 0x6b, 0xbd, 0x55, 0xb0, 0xd9, 0xb0, 0x60, 0xb3, 0x4a,
	0x74, 0x7f, 0x35, 0xe1, 0x9f, 0x03, 0x86, 0xc2, 0x73, 0x92, 0x61, 0x10, 0x9d, 0x45, 0x81, 0x10,
	0x48, 0x1f, 0xda, 0x3c, 0x2d, 0x58, 0x80, 0xb6, 0xd1, 0x37, 0x06, 0xdd, 0x51, 0xc7, 0x2b, 0xbd,
	0x4e, 0xe8, 0x31, 0xad, 0xf5, 0xe4, 0x21, 0x74, 0xa7, 0xc8, 0xf3, 0x28, 0x11, 0x0e, 0x76, 0xe3,
	0x1a, 0x4c, 0x35, 0x92, 0xd7, 0x60, 0x69, 0x55, 0xda, 0x4d, 0x81, 0xee, 0x79, 0x4a, 0xff, 0x07,
	0x2a, 0x80, 0xea, 0x78, 0xf2, 0x1e, 0xfe, 0xd6, 0x14, 0x93, 0xaa, 0xb6, 0xd6, 0xaa, 0x30, 0xcb,
	0xbc, 0xc8, 0x09, 0xd8, 0x9a, 0xfa, 0x8d, 0xd2, 0xc6, 0xda, 0xaa, 0x88, 0xb7, 0xba, 0x12, 0x02,
	0xad, 0xc4, 0x8f, 0xd1, 0x6e, 0xf7, 0x8d, 0x81, 0x49, 0xc5, 0x9b, 0x1c, 0x42, 0x7b, 0xe6, 0x9f,
	0xe2, 0x8c, 0xdb, 0xeb, 0xfd, 0xe6, 0xa0, 0x3b, 0xda, 0xd7, 0x02, 0x2f, 0x9b, 0xbc, 0x77, 0x2c,
	0xf0, 0x87, 0x49, 0xce, 0xae, 0x68, 0xed, 0x4c, 0x76, 0xa1, 0x9d, 0xf9, 0x05, 0xc7, 0xa9, 0xdd,
	0xe9, 0x1b, 0x83, 0x0e, 0xad, 0x25, 0xe7, 0x05, 0x74, 0x15, 0x38, 0xd9, 0x82, 0xe6, 0x39, 0x5e,
	0x09, 0xc6, 0x4c, 0x5a, 0x3e, 0xc9, 0x0e, 0xac, 0x5d, 0xfa, 0xb3, 0x02, 0x05, 0x3d, 0x26, 0xad,
	0x84, 0x97, 0x8d, 0xe7, 0x86, 0x9b, 0x83, 0x25, 0xf2, 0x23, 0xc5, 0x8b, 0x02, 0x79, 0x4e, 0xde,
	0x82, 0xc5, 0xd5, 0x42, 0x6a, 0xe2, 0xef, 0xad, 0xac, 0x98, 0xea, 0x7e, 0xc4, 0x81, 0x0e, 0x43,
	0x9e, 0xa5, 0x09, 0x97, 0x69, 0xe7, 0xb2, 0xfb, 0x05, 0x36, 0x65, 0xd6, 0x4a, 0x43, 0x6c, 0x58,
	0xe7, 0xc8, 0xb9, 0x4c, 0x68, 0x52, 0x29, 0x96, 0x96, 0x18, 0x39, 0xf7, 0x43, 0x19, 0x46, 0x8a,
	0x62, 0x1c, 0x2c, 0x8d, 0xb3, 0x5c, 0xec, 0x91, 0x49, 0x6b, 0xc9, 0xbd, 0x80, 0xee, 0x71, 0xc4,
	0x73, 0xd9, 0xd1, 0x08, 0xcc, 0xf9, 0xdd, 0xd4, 0xdd, 0xec, 0x78, 0x8b, 0x4b, 0x9a, 0xc8, 0x17,
	0x5d, 0xc0, 0x88, 0x07, 0x24, 0x63, 0x78, 0x19, 0xa5, 0x05, 0x9f, 0x94, 0x77, 0xf3, 0x2e, 0x99,
	0xe2, 0x0f, 0x91, 0xbf, 0x45, 0x97, 0x58, 0xdc, 0x10, 0x36, 0xaa, 0x94, 0x75, 0x3b, 0x7b, 0x00,
	0x7c, 0xe1, 0x67, 0x08, 0x3f, 0x45, 0x43, 0x9e, 0x81, 0x55, 0xf7, 0x27, 0x82, 0x70, 0xbb, 0x21,
	0xf6, 0x62, 0x5b, 0x9d, 0xb2, 0xb0, 0x50, 0x1d, 0xe7, 0x8e, 0x61, 0xe3, 0x63, 0x49, 0xfa, 0x1d,
	0x9a, 0x73, 0x1f, 0x80, 0x55, 0xc7, 0x58, 0x0c, 0x5f, 0x8e, 0xd8, 0xd0, 0x46, 0xec, 0x7e, 0x05,
	0x8b, 0x22, 0x2f, 0xe2, 0xbb, 0xe4, 0xfb, 0xe3, 0x26, 0x8c, 0x61, 0x53, 0x26, 0x58, 0x55, 0x8c,
	0xc2, 0x77, 0x43, 0xe3, 0xfb, 0x08, 0xb6, 0x3e, 0x21, 0x8b, 0xcb, 0x03, 0xbc, 0xd3, 0x5c, 0xf6,
	0x61, 0x5b, 0x89, 0xb3, 0xaa, 0x9c, 0xd1, 0xcf, 0x06, 0xc0, 0xd1, 0x9c, 0xae, 0xf2, 0xc6, 0xab,
	0x9d, 0x26, 0xbd, 0x1b, 0xb7, 0x22, 0xcb, 0x72, 0x9c, 0x65, 0xa6, 0x7a, 0x14, 0x7f, 0x0d, 0x8c,
	0x47, 0x06, 0x79, 0x05, 0xad, 0x72, 0x93, 0xc8, 0xbf, 0x2a, 0x52, 0x59, 0x67, 0xc7, 0xbe, 0x69,
	0x90, 0x01, 0xc8, 0x18, 0xd6, 0x04, 0xb3, 0x44, 0x03, 0xa9, 0x0b, 0xe3, 0xf4, 0x96, 0x58, 0xb4,
	0x02, 0x0e, 0xa1, 0x5d, 0x31, 0xa2, 0xf7, 0xa1, 0xad, 0x81, 0xde, 0x87, 0x4e, 0x60, 0x1d, 0xe6,
	0x03, 0x98, 0xf3, 0x61, 0x92, 0xff, 0x54, 0xf8, 0x75, 0xae, 0x9c, 0xff, 0x6f, 0xb1, 0xaa, 0xf1,
	0xc6, 0x4f, 0x3f, 0x3f, 0x09, 0xa3, 0xfc, 0x5b, 0x71, 0xea, 0x05, 0x69, 0x3c, 0x8c, 0x8b, 0xdc,
	0x0f, 0x31, 0xd9, 0x8f, 0x52, 0xf9, 0x1c, 0x66, 0xe7, 0xe1, 0xf0, 0xe6, 0xd7, 0x7a, 0xda, 0x16,
	0x5f, 0xdc, 0xe3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xdf, 0x16, 0xc3, 0x77, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ForwardingClient is the client API for Forwarding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForwardingClient interface {
	Create(ctx context.Context, opts ...grpc.CallOption) (Forwarding_CreateClient, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Pause(ctx context.Context, opts ...grpc.CallOption) (Forwarding_PauseClient, error)
	Resume(ctx context.Context, opts ...grpc.CallOption) (Forwarding_ResumeClient, error)
	Terminate(ctx context.Context, opts ...grpc.CallOption) (Forwarding_TerminateClient, error)
}

type forwardingClient struct {
	cc *grpc.ClientConn
}

func NewForwardingClient(cc *grpc.ClientConn) ForwardingClient {
	return &forwardingClient{cc}
}

func (c *forwardingClient) Create(ctx context.Context, opts ...grpc.CallOption) (Forwarding_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Forwarding_serviceDesc.Streams[0], "/forwarding.Forwarding/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &forwardingCreateClient{stream}
	return x, nil
}

type Forwarding_CreateClient interface {
	Send(*CreateRequest) error
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type forwardingCreateClient struct {
	grpc.ClientStream
}

func (x *forwardingCreateClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *forwardingCreateClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *forwardingClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/forwarding.Forwarding/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) Pause(ctx context.Context, opts ...grpc.CallOption) (Forwarding_PauseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Forwarding_serviceDesc.Streams[1], "/forwarding.Forwarding/Pause", opts...)
	if err != nil {
		return nil, err
	}
	x := &forwardingPauseClient{stream}
	return x, nil
}

type Forwarding_PauseClient interface {
	Send(*PauseRequest) error
	Recv() (*PauseResponse, error)
	grpc.ClientStream
}

type forwardingPauseClient struct {
	grpc.ClientStream
}

func (x *forwardingPauseClient) Send(m *PauseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *forwardingPauseClient) Recv() (*PauseResponse, error) {
	m := new(PauseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *forwardingClient) Resume(ctx context.Context, opts ...grpc.CallOption) (Forwarding_ResumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Forwarding_serviceDesc.Streams[2], "/forwarding.Forwarding/Resume", opts...)
	if err != nil {
		return nil, err
	}
	x := &forwardingResumeClient{stream}
	return x, nil
}

type Forwarding_ResumeClient interface {
	Send(*ResumeRequest) error
	Recv() (*ResumeResponse, error)
	grpc.ClientStream
}

type forwardingResumeClient struct {
	grpc.ClientStream
}

func (x *forwardingResumeClient) Send(m *ResumeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *forwardingResumeClient) Recv() (*ResumeResponse, error) {
	m := new(ResumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *forwardingClient) Terminate(ctx context.Context, opts ...grpc.CallOption) (Forwarding_TerminateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Forwarding_serviceDesc.Streams[3], "/forwarding.Forwarding/Terminate", opts...)
	if err != nil {
		return nil, err
	}
	x := &forwardingTerminateClient{stream}
	return x, nil
}

type Forwarding_TerminateClient interface {
	Send(*TerminateRequest) error
	Recv() (*TerminateResponse, error)
	grpc.ClientStream
}

type forwardingTerminateClient struct {
	grpc.ClientStream
}

func (x *forwardingTerminateClient) Send(m *TerminateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *forwardingTerminateClient) Recv() (*TerminateResponse, error) {
	m := new(TerminateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ForwardingServer is the server API for Forwarding service.
type ForwardingServer interface {
	Create(Forwarding_CreateServer) error
	List(context.Context, *ListRequest) (*ListResponse, error)
	Pause(Forwarding_PauseServer) error
	Resume(Forwarding_ResumeServer) error
	Terminate(Forwarding_TerminateServer) error
}

// UnimplementedForwardingServer can be embedded to have forward compatible implementations.
type UnimplementedForwardingServer struct {
}

func (*UnimplementedForwardingServer) Create(srv Forwarding_CreateServer) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedForwardingServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedForwardingServer) Pause(srv Forwarding_PauseServer) error {
	return status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (*UnimplementedForwardingServer) Resume(srv Forwarding_ResumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (*UnimplementedForwardingServer) Terminate(srv Forwarding_TerminateServer) error {
	return status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}

func RegisterForwardingServer(s *grpc.Server, srv ForwardingServer) {
	s.RegisterService(&_Forwarding_serviceDesc, srv)
}

func _Forwarding_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ForwardingServer).Create(&forwardingCreateServer{stream})
}

type Forwarding_CreateServer interface {
	Send(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type forwardingCreateServer struct {
	grpc.ServerStream
}

func (x *forwardingCreateServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *forwardingCreateServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Forwarding_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarding.Forwarding/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_Pause_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ForwardingServer).Pause(&forwardingPauseServer{stream})
}

type Forwarding_PauseServer interface {
	Send(*PauseResponse) error
	Recv() (*PauseRequest, error)
	grpc.ServerStream
}

type forwardingPauseServer struct {
	grpc.ServerStream
}

func (x *forwardingPauseServer) Send(m *PauseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *forwardingPauseServer) Recv() (*PauseRequest, error) {
	m := new(PauseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Forwarding_Resume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ForwardingServer).Resume(&forwardingResumeServer{stream})
}

type Forwarding_ResumeServer interface {
	Send(*ResumeResponse) error
	Recv() (*ResumeRequest, error)
	grpc.ServerStream
}

type forwardingResumeServer struct {
	grpc.ServerStream
}

func (x *forwardingResumeServer) Send(m *ResumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *forwardingResumeServer) Recv() (*ResumeRequest, error) {
	m := new(ResumeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Forwarding_Terminate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ForwardingServer).Terminate(&forwardingTerminateServer{stream})
}

type Forwarding_TerminateServer interface {
	Send(*TerminateResponse) error
	Recv() (*TerminateRequest, error)
	grpc.ServerStream
}

type forwardingTerminateServer struct {
	grpc.ServerStream
}

func (x *forwardingTerminateServer) Send(m *TerminateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *forwardingTerminateServer) Recv() (*TerminateRequest, error) {
	m := new(TerminateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Forwarding_serviceDesc = grpc.ServiceDesc{
	ServiceName: "forwarding.Forwarding",
	HandlerType: (*ForwardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Forwarding_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _Forwarding_Create_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Pause",
			Handler:       _Forwarding_Pause_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Resume",
			Handler:       _Forwarding_Resume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Terminate",
			Handler:       _Forwarding_Terminate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service/forwarding/forwarding.proto",
}
