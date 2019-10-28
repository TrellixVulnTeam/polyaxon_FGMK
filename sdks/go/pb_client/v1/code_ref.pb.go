// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/code_ref.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Code Reference specification
type CodeReference struct {
	// UUID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Optional commit of the code ref
	Commit string `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	// Optional head of the code ref
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Optional branch of the code ref
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// Optional git url for this code ref
	GitUrl string `protobuf:"bytes,5,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
	// Optional whether this code ref is dirty
	IsDirty              bool     `protobuf:"varint,6,opt,name=is_dirty,json=isDirty,proto3" json:"is_dirty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeReference) Reset()         { *m = CodeReference{} }
func (m *CodeReference) String() string { return proto.CompactTextString(m) }
func (*CodeReference) ProtoMessage()    {}
func (*CodeReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_f834b60ccaa0a742, []int{0}
}

func (m *CodeReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeReference.Unmarshal(m, b)
}
func (m *CodeReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeReference.Marshal(b, m, deterministic)
}
func (m *CodeReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeReference.Merge(m, src)
}
func (m *CodeReference) XXX_Size() int {
	return xxx_messageInfo_CodeReference.Size(m)
}
func (m *CodeReference) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeReference.DiscardUnknown(m)
}

var xxx_messageInfo_CodeReference proto.InternalMessageInfo

func (m *CodeReference) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CodeReference) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *CodeReference) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *CodeReference) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CodeReference) GetGitUrl() string {
	if m != nil {
		return m.GitUrl
	}
	return ""
}

func (m *CodeReference) GetIsDirty() bool {
	if m != nil {
		return m.IsDirty
	}
	return false
}

// Request data to create/update CodeReference
type CodeRefBodyRequest struct {
	// Owner of the code ref
	Entity *ProjectEntityResourceRequest `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// Code ref object
	CodeReference        *CodeReference `protobuf:"bytes,2,opt,name=code_reference,json=codeReference,proto3" json:"code_reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CodeRefBodyRequest) Reset()         { *m = CodeRefBodyRequest{} }
func (m *CodeRefBodyRequest) String() string { return proto.CompactTextString(m) }
func (*CodeRefBodyRequest) ProtoMessage()    {}
func (*CodeRefBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f834b60ccaa0a742, []int{1}
}

func (m *CodeRefBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeRefBodyRequest.Unmarshal(m, b)
}
func (m *CodeRefBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeRefBodyRequest.Marshal(b, m, deterministic)
}
func (m *CodeRefBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeRefBodyRequest.Merge(m, src)
}
func (m *CodeRefBodyRequest) XXX_Size() int {
	return xxx_messageInfo_CodeRefBodyRequest.Size(m)
}
func (m *CodeRefBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeRefBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeRefBodyRequest proto.InternalMessageInfo

func (m *CodeRefBodyRequest) GetEntity() *ProjectEntityResourceRequest {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CodeRefBodyRequest) GetCodeReference() *CodeReference {
	if m != nil {
		return m.CodeReference
	}
	return nil
}

// Contains list runs
type ListCodeRefsResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*CodeReference `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCodeRefsResponse) Reset()         { *m = ListCodeRefsResponse{} }
func (m *ListCodeRefsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCodeRefsResponse) ProtoMessage()    {}
func (*ListCodeRefsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f834b60ccaa0a742, []int{2}
}

func (m *ListCodeRefsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCodeRefsResponse.Unmarshal(m, b)
}
func (m *ListCodeRefsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCodeRefsResponse.Marshal(b, m, deterministic)
}
func (m *ListCodeRefsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCodeRefsResponse.Merge(m, src)
}
func (m *ListCodeRefsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCodeRefsResponse.Size(m)
}
func (m *ListCodeRefsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCodeRefsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCodeRefsResponse proto.InternalMessageInfo

func (m *ListCodeRefsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListCodeRefsResponse) GetResults() []*CodeReference {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListCodeRefsResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListCodeRefsResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func init() {
	proto.RegisterType((*CodeReference)(nil), "v1.CodeReference")
	proto.RegisterType((*CodeRefBodyRequest)(nil), "v1.CodeRefBodyRequest")
	proto.RegisterType((*ListCodeRefsResponse)(nil), "v1.ListCodeRefsResponse")
}

func init() { proto.RegisterFile("v1/code_ref.proto", fileDescriptor_f834b60ccaa0a742) }

var fileDescriptor_f834b60ccaa0a742 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0xa3, 0x30,
	0x10, 0xc6, 0x45, 0xfe, 0x90, 0xc4, 0x51, 0x56, 0x8a, 0x15, 0xed, 0xb2, 0x5c, 0x16, 0xe5, 0x14,
	0x69, 0x25, 0x10, 0xd9, 0x4b, 0xf6, 0xd8, 0x7f, 0xb7, 0x1e, 0x2a, 0xab, 0x3d, 0x23, 0x02, 0x13,
	0xe4, 0x0a, 0x30, 0xb5, 0xc7, 0xa8, 0x79, 0x83, 0xf6, 0xb5, 0xfa, 0x64, 0x15, 0x36, 0x54, 0xad,
	0xd4, 0x9b, 0x7f, 0x33, 0x1f, 0xc3, 0x37, 0xdf, 0x90, 0x75, 0x1b, 0x47, 0x99, 0xc8, 0x21, 0x91,
	0x70, 0x0a, 0x1b, 0x29, 0x50, 0xd0, 0x51, 0x1b, 0xfb, 0x7f, 0x0a, 0x21, 0x8a, 0x12, 0x22, 0x53,
	0x39, 0xea, 0x53, 0x84, 0xbc, 0x02, 0x85, 0x69, 0xd5, 0x58, 0x91, 0xbf, 0x6a, 0xe3, 0xe8, 0x98,
	0x2a, 0xb0, 0xb8, 0x7d, 0x73, 0xc8, 0xea, 0x4a, 0xe4, 0xc0, 0xe0, 0x04, 0x12, 0xea, 0x0c, 0x28,
	0x25, 0x13, 0xad, 0x79, 0xee, 0x39, 0x81, 0xb3, 0x5b, 0x30, 0xf3, 0xa6, 0x3f, 0x89, 0x9b, 0x89,
	0xaa, 0xe2, 0xe8, 0x8d, 0x4c, 0xb5, 0x27, 0xfa, 0x9f, 0x10, 0xdd, 0xe4, 0x29, 0x42, 0x9e, 0xa4,
	0xe8, 0x8d, 0x03, 0x67, 0xb7, 0xdc, 0xfb, 0xa1, 0xb5, 0x10, 0x0e, 0x16, 0xc2, 0xfb, 0xc1, 0x02,
	0x5b, 0xf4, 0xea, 0x0b, 0xec, 0x46, 0x2a, 0x4c, 0x51, 0x2b, 0x6f, 0x62, 0x47, 0x5a, 0xa2, 0xbf,
	0xc8, 0xac, 0xe0, 0x98, 0x68, 0x59, 0x7a, 0x53, 0xdb, 0x28, 0x38, 0x3e, 0xc8, 0x92, 0xfe, 0x26,
	0x73, 0xae, 0x92, 0x9c, 0x4b, 0x3c, 0x7b, 0x6e, 0xe0, 0xec, 0xe6, 0x6c, 0xc6, 0xd5, 0x75, 0x87,
	0xdb, 0x17, 0x87, 0xd0, 0x7e, 0x89, 0x4b, 0x91, 0x9f, 0x19, 0x3c, 0x69, 0x50, 0x48, 0x0f, 0xc4,
	0x85, 0x1a, 0x39, 0x9e, 0xcd, 0x2e, 0xcb, 0x7d, 0x10, 0xb6, 0x71, 0x78, 0x27, 0xc5, 0x23, 0x64,
	0x78, 0x63, 0x1a, 0x0c, 0x94, 0xd0, 0x32, 0x83, 0xfe, 0x0b, 0xd6, 0xeb, 0xe9, 0x81, 0xfc, 0x18,
	0xb2, 0xb5, 0xa9, 0x98, 0xbd, 0x97, 0xfb, 0x75, 0x37, 0xe1, 0x4b, 0x5c, 0x6c, 0x95, 0x7d, 0xc6,
	0xed, 0xab, 0x43, 0x36, 0xb7, 0x5c, 0x61, 0x2f, 0x52, 0x0c, 0x54, 0x23, 0x6a, 0x05, 0x74, 0x43,
	0xa6, 0x99, 0xd0, 0x35, 0x1a, 0x2f, 0x53, 0x66, 0x81, 0xfe, 0x25, 0x33, 0x09, 0x4a, 0x97, 0xa8,
	0xbc, 0x51, 0x30, 0xfe, 0xfe, 0x0f, 0x83, 0x82, 0xfa, 0x64, 0xde, 0x48, 0x68, 0xb9, 0xd0, 0xca,
	0x64, 0xbd, 0x60, 0x1f, 0xdc, 0x5d, 0xad, 0x86, 0x67, 0xec, 0xc3, 0x34, 0xef, 0xa3, 0x6b, 0x2e,
	0xf0, 0xef, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xe5, 0x48, 0x32, 0x2b, 0x02, 0x00, 0x00,
}
