// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entry.proto

package sync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EntryKind int32

const (
	EntryKind_Directory EntryKind = 0
	EntryKind_File      EntryKind = 1
	EntryKind_Symlink   EntryKind = 2
)

var EntryKind_name = map[int32]string{
	0: "Directory",
	1: "File",
	2: "Symlink",
}
var EntryKind_value = map[string]int32{
	"Directory": 0,
	"File":      1,
	"Symlink":   2,
}

func (x EntryKind) String() string {
	return proto.EnumName(EntryKind_name, int32(x))
}
func (EntryKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_entry_864690241278c8e3, []int{0}
}

type Entry struct {
	Kind                 EntryKind         `protobuf:"varint,1,opt,name=kind,enum=sync.EntryKind" json:"kind,omitempty"`
	Executable           bool              `protobuf:"varint,2,opt,name=executable" json:"executable,omitempty"`
	Digest               []byte            `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Contents             map[string]*Entry `protobuf:"bytes,4,rep,name=contents" json:"contents,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Target               string            `protobuf:"bytes,5,opt,name=target" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_entry_864690241278c8e3, []int{0}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (dst *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(dst, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetKind() EntryKind {
	if m != nil {
		return m.Kind
	}
	return EntryKind_Directory
}

func (m *Entry) GetExecutable() bool {
	if m != nil {
		return m.Executable
	}
	return false
}

func (m *Entry) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Entry) GetContents() map[string]*Entry {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Entry) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func init() {
	proto.RegisterType((*Entry)(nil), "sync.Entry")
	proto.RegisterMapType((map[string]*Entry)(nil), "sync.Entry.ContentsEntry")
	proto.RegisterEnum("sync.EntryKind", EntryKind_name, EntryKind_value)
}

func init() { proto.RegisterFile("entry.proto", fileDescriptor_entry_864690241278c8e3) }

var fileDescriptor_entry_864690241278c8e3 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xcd, 0xb6, 0x5d, 0xdb, 0xa9, 0xab, 0x65, 0x0e, 0x52, 0x3d, 0x48, 0xd4, 0x4b, 0xf0,
	0x50, 0xb0, 0x22, 0x88, 0x57, 0xff, 0x20, 0x78, 0x8b, 0x9f, 0xa0, 0xdb, 0x0e, 0x4b, 0x68, 0x4d,
	0x25, 0xcd, 0x8a, 0xf9, 0xf0, 0x82, 0x34, 0x5b, 0x96, 0x7a, 0xcb, 0x7b, 0xbf, 0x24, 0xef, 0xcd,
	0x40, 0x4a, 0xda, 0x1a, 0x57, 0x7c, 0x99, 0xde, 0xf6, 0x18, 0x0e, 0x4e, 0xd7, 0x57, 0xbf, 0x0c,
	0xa2, 0x97, 0xd1, 0xc5, 0x6b, 0x08, 0x5b, 0xa5, 0x9b, 0x9c, 0x71, 0x26, 0x8e, 0xcb, 0x93, 0x62,
	0xc4, 0x85, 0x47, 0xef, 0x4a, 0x37, 0xd2, 0x43, 0xbc, 0x00, 0xa0, 0x1f, 0xaa, 0xb7, 0xb6, 0x5a,
	0x77, 0x94, 0x2f, 0x38, 0x13, 0xb1, 0x9c, 0x39, 0x78, 0x0a, 0xcb, 0x46, 0x6d, 0x68, 0xb0, 0x79,
	0xc0, 0x99, 0x38, 0x92, 0x93, 0xc2, 0x7b, 0x88, 0xeb, 0x5e, 0x5b, 0xd2, 0x76, 0xc8, 0x43, 0x1e,
	0x88, 0xb4, 0x3c, 0x9b, 0x05, 0x14, 0x4f, 0x13, 0xf3, 0x4a, 0xee, 0xaf, 0x8e, 0xdf, 0xd9, 0xca,
	0x6c, 0xc8, 0xe6, 0x11, 0x67, 0x22, 0x91, 0x93, 0x3a, 0x7f, 0x83, 0xd5, 0xbf, 0x27, 0x98, 0x41,
	0xd0, 0x92, 0xf3, 0xdd, 0x13, 0x39, 0x1e, 0xf1, 0x12, 0xa2, 0xef, 0xaa, 0xdb, 0xee, 0x4a, 0xa6,
	0x65, 0x3a, 0x8b, 0x93, 0x3b, 0xf2, 0xb8, 0x78, 0x60, 0x37, 0xb7, 0x90, 0xec, 0x67, 0xc4, 0x15,
	0x24, 0xcf, 0xca, 0x50, 0x6d, 0x7b, 0xe3, 0xb2, 0x03, 0x8c, 0x21, 0x7c, 0x55, 0x1d, 0x65, 0x0c,
	0x53, 0x38, 0xfc, 0x70, 0x9f, 0x9d, 0xd2, 0x6d, 0xb6, 0x58, 0x2f, 0xfd, 0xfe, 0xee, 0xfe, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x79, 0x70, 0x7c, 0x47, 0x4e, 0x01, 0x00, 0x00,
}