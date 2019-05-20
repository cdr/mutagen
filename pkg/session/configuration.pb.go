// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session/configuration.proto

package session

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	behavior "github.com/havoc-io/mutagen/pkg/filesystem/behavior"
	staging "github.com/havoc-io/mutagen/pkg/staging"
	sync "github.com/havoc-io/mutagen/pkg/sync"
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

// Configuration encodes session configuration parameters. It is used for create
// commands to specify configuration options, for loading global configuration
// options, and for storing a merged configuration inside sessions.
type Configuration struct {
	// SynchronizationMode specifies the synchronization mode that should be
	// used in synchronization.
	SynchronizationMode sync.SynchronizationMode `protobuf:"varint,11,opt,name=synchronizationMode,proto3,enum=sync.SynchronizationMode" json:"synchronizationMode,omitempty"`
	// MaximumEntryCount specifies the maximum number of filesystem entries that
	// endpoints will tolerate managing. A zero value indicates no limit.
	MaximumEntryCount uint64 `protobuf:"varint,12,opt,name=maximumEntryCount,proto3" json:"maximumEntryCount,omitempty"`
	// MaximumStagingFileSize is the maximum (individual) file size that
	// endpoints will stage. A zero value indicates no limit.
	MaximumStagingFileSize uint64 `protobuf:"varint,13,opt,name=maximumStagingFileSize,proto3" json:"maximumStagingFileSize,omitempty"`
	// ProbeMode specifies the filesystem probing mode.
	ProbeMode behavior.ProbeMode `protobuf:"varint,14,opt,name=probeMode,proto3,enum=behavior.ProbeMode" json:"probeMode,omitempty"`
	// StagingMode specifies the file staging mode.
	StagingMode staging.StagingMode `protobuf:"varint,16,opt,name=stagingMode,proto3,enum=staging.StagingMode" json:"stagingMode,omitempty"`
	// SymlinkMode specifies the symlink mode that should be used in
	// synchronization.
	SymlinkMode sync.SymlinkMode `protobuf:"varint,1,opt,name=symlinkMode,proto3,enum=sync.SymlinkMode" json:"symlinkMode,omitempty"`
	// WatchMode specifies the filesystem watching mode.
	WatchMode WatchMode `protobuf:"varint,21,opt,name=watchMode,proto3,enum=session.WatchMode" json:"watchMode,omitempty"`
	// WatchPollingInterval specifies the interval (in seconds) for poll-based
	// file monitoring. A value of 0 specifies that the default interval should
	// be used.
	WatchPollingInterval uint32 `protobuf:"varint,22,opt,name=watchPollingInterval,proto3" json:"watchPollingInterval,omitempty"`
	// DefaultIgnores specifies the ignore patterns brought in from the global
	// configuration.
	// DEPRECATED: This field is no longer used when loading from global
	// configuration. Instead, ignores provided by global configuration are
	// simply merged into the ignore list of the main configuration. However,
	// older sessions still use this field.
	DefaultIgnores []string `protobuf:"bytes,31,rep,name=defaultIgnores,proto3" json:"defaultIgnores,omitempty"`
	// Ignores specifies the ignore patterns brought in from the create request.
	Ignores []string `protobuf:"bytes,32,rep,name=ignores,proto3" json:"ignores,omitempty"`
	// IgnoreVCSMode specifies the VCS ignore mode that should be used in
	// synchronization.
	IgnoreVCSMode sync.IgnoreVCSMode `protobuf:"varint,33,opt,name=ignoreVCSMode,proto3,enum=sync.IgnoreVCSMode" json:"ignoreVCSMode,omitempty"`
	// DefaultFileMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultFileMode uint32 `protobuf:"varint,63,opt,name=defaultFileMode,proto3" json:"defaultFileMode,omitempty"`
	// DefaultDirectoryMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultDirectoryMode uint32 `protobuf:"varint,64,opt,name=defaultDirectoryMode,proto3" json:"defaultDirectoryMode,omitempty"`
	// DefaultOwner specifies the default owner identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultOwner string `protobuf:"bytes,65,opt,name=defaultOwner,proto3" json:"defaultOwner,omitempty"`
	// DefaultGroup specifies the default group identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultGroup         string   `protobuf:"bytes,66,opt,name=defaultGroup,proto3" json:"defaultGroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_34b104152623cac3, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetSynchronizationMode() sync.SynchronizationMode {
	if m != nil {
		return m.SynchronizationMode
	}
	return sync.SynchronizationMode_SynchronizationModeDefault
}

func (m *Configuration) GetMaximumEntryCount() uint64 {
	if m != nil {
		return m.MaximumEntryCount
	}
	return 0
}

func (m *Configuration) GetMaximumStagingFileSize() uint64 {
	if m != nil {
		return m.MaximumStagingFileSize
	}
	return 0
}

func (m *Configuration) GetProbeMode() behavior.ProbeMode {
	if m != nil {
		return m.ProbeMode
	}
	return behavior.ProbeMode_ProbeModeDefault
}

func (m *Configuration) GetStagingMode() staging.StagingMode {
	if m != nil {
		return m.StagingMode
	}
	return staging.StagingMode_StagingModeDefault
}

func (m *Configuration) GetSymlinkMode() sync.SymlinkMode {
	if m != nil {
		return m.SymlinkMode
	}
	return sync.SymlinkMode_SymlinkModeDefault
}

func (m *Configuration) GetWatchMode() WatchMode {
	if m != nil {
		return m.WatchMode
	}
	return WatchMode_WatchModeDefault
}

func (m *Configuration) GetWatchPollingInterval() uint32 {
	if m != nil {
		return m.WatchPollingInterval
	}
	return 0
}

func (m *Configuration) GetDefaultIgnores() []string {
	if m != nil {
		return m.DefaultIgnores
	}
	return nil
}

func (m *Configuration) GetIgnores() []string {
	if m != nil {
		return m.Ignores
	}
	return nil
}

func (m *Configuration) GetIgnoreVCSMode() sync.IgnoreVCSMode {
	if m != nil {
		return m.IgnoreVCSMode
	}
	return sync.IgnoreVCSMode_IgnoreVCSModeDefault
}

func (m *Configuration) GetDefaultFileMode() uint32 {
	if m != nil {
		return m.DefaultFileMode
	}
	return 0
}

func (m *Configuration) GetDefaultDirectoryMode() uint32 {
	if m != nil {
		return m.DefaultDirectoryMode
	}
	return 0
}

func (m *Configuration) GetDefaultOwner() string {
	if m != nil {
		return m.DefaultOwner
	}
	return ""
}

func (m *Configuration) GetDefaultGroup() string {
	if m != nil {
		return m.DefaultGroup
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "session.Configuration")
}

func init() { proto.RegisterFile("session/configuration.proto", fileDescriptor_34b104152623cac3) }

var fileDescriptor_34b104152623cac3 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdb, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x55, 0x71, 0x99, 0xea, 0xae, 0x1d, 0x75, 0xb7, 0x29, 0x8c, 0x87, 0x85, 0x3d, 0x40,
	0x90, 0x20, 0x81, 0x4d, 0x9a, 0xc4, 0x13, 0xb0, 0x72, 0x51, 0x85, 0x10, 0x93, 0x23, 0x81, 0xc4,
	0x5b, 0x9a, 0xba, 0x89, 0xb5, 0xc4, 0xa7, 0xb2, 0x9d, 0x8e, 0xec, 0x5f, 0xe6, 0x9f, 0x40, 0x3d,
	0xb1, 0x69, 0xba, 0x95, 0xa7, 0xc4, 0xbf, 0xef, 0xfb, 0x92, 0x73, 0x49, 0xc8, 0x13, 0xcd, 0xb5,
	0x16, 0x20, 0xa3, 0x14, 0xe4, 0x5c, 0x64, 0x95, 0x4a, 0x8c, 0x00, 0x19, 0x2e, 0x14, 0x18, 0xa0,
	0x3b, 0x56, 0x3c, 0x3a, 0x9e, 0x8b, 0x82, 0xeb, 0x5a, 0x1b, 0x5e, 0x46, 0x53, 0x9e, 0x27, 0x4b,
	0x01, 0x2a, 0x5a, 0x28, 0x98, 0xf2, 0xc6, 0x79, 0x34, 0x72, 0x8f, 0xb9, 0x4e, 0x4c, 0x9a, 0x5b,
	0x78, 0xa0, 0x4d, 0x92, 0x09, 0x99, 0x45, 0xf6, 0x6a, 0xf1, 0x50, 0xd7, 0x32, 0x8d, 0x44, 0x26,
	0x41, 0xb9, 0xf8, 0x1e, 0xa2, 0x12, 0x66, 0x0e, 0x50, 0x04, 0xba, 0x2e, 0x0b, 0x21, 0xaf, 0x1a,
	0x76, 0xf2, 0xe7, 0x01, 0xe9, 0x8f, 0xdb, 0x55, 0xd2, 0xaf, 0x64, 0xb4, 0xf2, 0xe5, 0x0a, 0xa4,
	0xb8, 0x41, 0xf4, 0x0d, 0x66, 0xdc, 0xeb, 0xf9, 0x9d, 0x60, 0x70, 0xfa, 0x38, 0x5c, 0x69, 0x61,
	0x7c, 0xd7, 0xc0, 0xb6, 0xa5, 0xe8, 0x4b, 0x32, 0x2c, 0x93, 0xdf, 0xa2, 0xac, 0xca, 0x4f, 0xd2,
	0xa8, 0x7a, 0x0c, 0x95, 0x34, 0xde, 0xae, 0xdf, 0x09, 0xee, 0xb3, 0xbb, 0x02, 0x3d, 0x27, 0x87,
	0x16, 0xc6, 0x4d, 0x73, 0x9f, 0x45, 0xc1, 0x63, 0x71, 0xc3, 0xbd, 0x3e, 0x46, 0xfe, 0xa3, 0xd2,
	0x37, 0xa4, 0x8b, 0x73, 0xc3, 0x42, 0x07, 0x58, 0xe8, 0x28, 0x74, 0x23, 0x0d, 0x2f, 0x9d, 0xc4,
	0xd6, 0x2e, 0x7a, 0x4e, 0x7a, 0x76, 0x80, 0x18, 0x7a, 0x84, 0xa1, 0xfd, 0xd0, 0x0d, 0x35, 0x5e,
	0x6b, 0xac, 0x6d, 0xa4, 0x67, 0xa4, 0x67, 0x07, 0x88, 0xb9, 0x0e, 0xe6, 0x86, 0x6e, 0x2a, 0xff,
	0x04, 0xd6, 0x76, 0xd1, 0xd7, 0xa4, 0x8b, 0x2b, 0xc4, 0xc8, 0x01, 0x46, 0x68, 0x68, 0x97, 0x1b,
	0xfe, 0x74, 0x0a, 0x5b, 0x9b, 0xe8, 0x29, 0xd9, 0xc7, 0xc3, 0x25, 0x14, 0x85, 0x90, 0xd9, 0x44,
	0x1a, 0xae, 0x96, 0x49, 0xe1, 0x1d, 0xfa, 0x9d, 0xa0, 0xcf, 0xb6, 0x6a, 0xf4, 0x19, 0x19, 0xcc,
	0xf8, 0x3c, 0xa9, 0x0a, 0x33, 0xc1, 0xcf, 0x40, 0x7b, 0xc7, 0xfe, 0xbd, 0xa0, 0xcb, 0x6e, 0x51,
	0xea, 0x91, 0x1d, 0x61, 0x0d, 0x3e, 0x1a, 0xdc, 0x91, 0xbe, 0x25, 0xfd, 0xe6, 0xf6, 0xc7, 0x38,
	0xc6, 0x5a, 0x9f, 0xda, 0x59, 0x62, 0x7b, 0x93, 0xb6, 0xc4, 0x36, 0x9d, 0x34, 0x20, 0x7b, 0xf6,
	0x35, 0xab, 0xad, 0x60, 0xf8, 0x1d, 0xd6, 0x7a, 0x1b, 0xaf, 0x5a, 0xb3, 0xe8, 0xa3, 0x50, 0x3c,
	0x35, 0xa0, 0x6a, 0xb4, 0xbf, 0x6f, 0x5a, 0xdb, 0xa6, 0xd1, 0x13, 0xb2, 0x6b, 0xf9, 0xf7, 0x6b,
	0xc9, 0x95, 0xf7, 0xc1, 0xef, 0x04, 0x5d, 0xb6, 0xc1, 0x5a, 0x9e, 0x2f, 0x0a, 0xaa, 0x85, 0x77,
	0xb1, 0xe1, 0x41, 0x76, 0xf1, 0xe2, 0xd7, 0xf3, 0x4c, 0x98, 0xbc, 0x9a, 0x86, 0x29, 0x94, 0x51,
	0x9e, 0x2c, 0x21, 0x7d, 0x25, 0x20, 0x2a, 0x2b, 0x93, 0x64, 0x5c, 0x46, 0x8b, 0xab, 0x2c, 0xb2,
	0x6b, 0x99, 0x3e, 0xc4, 0xff, 0xe3, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x40, 0x29,
	0x47, 0xcc, 0x03, 0x00, 0x00,
}
