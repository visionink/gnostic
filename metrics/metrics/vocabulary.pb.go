// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: vocabulary.proto

package gnostic_metrics_v1

import (
	proto "google.golang.org/protobuf/proto"
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
var _ = proto.Error

type WordCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word  string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *WordCount) Reset() {
	*x = WordCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordCount) ProtoMessage() {}

func (x *WordCount) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordCount.ProtoReflect.Descriptor instead.
func (*WordCount) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{0}
}

func (x *WordCount) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *WordCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Vocabulary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schemas    []*WordCount `protobuf:"bytes,1,rep,name=schemas,proto3" json:"schemas,omitempty"`
	Properties []*WordCount `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty"`
	Operations []*WordCount `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	Parameters []*WordCount `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *Vocabulary) Reset() {
	*x = Vocabulary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vocabulary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vocabulary) ProtoMessage() {}

func (x *Vocabulary) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vocabulary.ProtoReflect.Descriptor instead.
func (*Vocabulary) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{1}
}

func (x *Vocabulary) GetSchemas() []*WordCount {
	if x != nil {
		return x.Schemas
	}
	return nil
}

func (x *Vocabulary) GetProperties() []*WordCount {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Vocabulary) GetOperations() []*WordCount {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *Vocabulary) GetParameters() []*WordCount {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	NewTerms         *Vocabulary `protobuf:"bytes,1,opt,name=newTerms,proto3" json:"newTerms,omitempty"`
	DeletedTerms     *Vocabulary `protobuf:"bytes,2,opt,name=deletedTerms,proto3" json:"deletedTerms,omitempty"`
	NewTermCount     int32       `protobuf:"varint,4,opt,name=newTermCount,proto3" json:"newTermCount,omitempty"`
	DeletedTermCount int32       `protobuf:"varint,5,opt,name=deletedTermCount,proto3" json:"deletedTermCount,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{2}
}

func (x *Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Version) GetNewTerms() *Vocabulary {
	if x != nil {
		return x.NewTerms
	}
	return nil
}

func (x *Version) GetDeletedTerms() *Vocabulary {
	if x != nil {
		return x.DeletedTerms
	}
	return nil
}

func (x *Version) GetNewTermCount() int32 {
	if x != nil {
		return x.NewTermCount
	}
	return 0
}

func (x *Version) GetDeletedTermCount() int32 {
	if x != nil {
		return x.DeletedTermCount
	}
	return 0
}

type VersionHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []*Version `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *VersionHistory) Reset() {
	*x = VersionHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vocabulary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionHistory) ProtoMessage() {}

func (x *VersionHistory) ProtoReflect() protoreflect.Message {
	mi := &file_vocabulary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionHistory.ProtoReflect.Descriptor instead.
func (*VersionHistory) Descriptor() ([]byte, []int) {
	return file_vocabulary_proto_rawDescGZIP(), []int{3}
}

func (x *VersionHistory) GetVersions() []*Version {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *VersionHistory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_vocabulary_proto protoreflect.FileDescriptor

var file_vocabulary_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x35, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x82, 0x02,
	0x0a, 0x0a, 0x56, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x07,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6e, 0x6f, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x6f, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x22, 0xed, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x12, 0x42,
	0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x65, 0x72,
	0x6d, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x54, 0x65, 0x72, 0x6d, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x54, 0x65, 0x72,
	0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x54, 0x65, 0x72, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x5d, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x3b, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vocabulary_proto_rawDescOnce sync.Once
	file_vocabulary_proto_rawDescData = file_vocabulary_proto_rawDesc
)

func file_vocabulary_proto_rawDescGZIP() []byte {
	file_vocabulary_proto_rawDescOnce.Do(func() {
		file_vocabulary_proto_rawDescData = protoimpl.X.CompressGZIP(file_vocabulary_proto_rawDescData)
	})
	return file_vocabulary_proto_rawDescData
}

var file_vocabulary_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vocabulary_proto_goTypes = []interface{}{
	(*WordCount)(nil),      // 0: gnostic.metrics.v1.WordCount
	(*Vocabulary)(nil),     // 1: gnostic.metrics.v1.Vocabulary
	(*Version)(nil),        // 2: gnostic.metrics.v1.Version
	(*VersionHistory)(nil), // 3: gnostic.metrics.v1.VersionHistory
}
var file_vocabulary_proto_depIdxs = []int32{
	0, // 0: gnostic.metrics.v1.Vocabulary.schemas:type_name -> gnostic.metrics.v1.WordCount
	0, // 1: gnostic.metrics.v1.Vocabulary.properties:type_name -> gnostic.metrics.v1.WordCount
	0, // 2: gnostic.metrics.v1.Vocabulary.operations:type_name -> gnostic.metrics.v1.WordCount
	0, // 3: gnostic.metrics.v1.Vocabulary.parameters:type_name -> gnostic.metrics.v1.WordCount
	1, // 4: gnostic.metrics.v1.Version.newTerms:type_name -> gnostic.metrics.v1.Vocabulary
	1, // 5: gnostic.metrics.v1.Version.deletedTerms:type_name -> gnostic.metrics.v1.Vocabulary
	2, // 6: gnostic.metrics.v1.VersionHistory.versions:type_name -> gnostic.metrics.v1.Version
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_vocabulary_proto_init() }
func file_vocabulary_proto_init() {
	if File_vocabulary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vocabulary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordCount); i {
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
		file_vocabulary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vocabulary); i {
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
		file_vocabulary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_vocabulary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionHistory); i {
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
			RawDescriptor: file_vocabulary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vocabulary_proto_goTypes,
		DependencyIndexes: file_vocabulary_proto_depIdxs,
		MessageInfos:      file_vocabulary_proto_msgTypes,
	}.Build()
	File_vocabulary_proto = out.File
	file_vocabulary_proto_rawDesc = nil
	file_vocabulary_proto_goTypes = nil
	file_vocabulary_proto_depIdxs = nil
}
