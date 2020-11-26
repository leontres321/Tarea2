// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: ftp.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk    []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Last     bool   `protobuf:"varint,3,opt,name=last,proto3" json:"last,omitempty"`
	First    bool   `protobuf:"varint,4,opt,name=first,proto3" json:"first,omitempty"`
	Name     string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Parts    uint64 `protobuf:"varint,6,opt,name=parts,proto3" json:"parts,omitempty"`
	ThisPart uint64 `protobuf:"varint,7,opt,name=this_part,json=thisPart,proto3" json:"this_part,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_ftp_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *Chunk) GetLast() bool {
	if x != nil {
		return x.Last
	}
	return false
}

func (x *Chunk) GetFirst() bool {
	if x != nil {
		return x.First
	}
	return false
}

func (x *Chunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chunk) GetParts() uint64 {
	if x != nil {
		return x.Parts
	}
	return 0
}

func (x *Chunk) GetThisPart() uint64 {
	if x != nil {
		return x.ThisPart
	}
	return 0
}

type Respuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Gud bool   `protobuf:"varint,2,opt,name=gud,proto3" json:"gud,omitempty"`
}

func (x *Respuesta) Reset() {
	*x = Respuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuesta) ProtoMessage() {}

func (x *Respuesta) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respuesta.ProtoReflect.Descriptor instead.
func (*Respuesta) Descriptor() ([]byte, []int) {
	return file_ftp_proto_rawDescGZIP(), []int{1}
}

func (x *Respuesta) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Respuesta) GetGud() bool {
	if x != nil {
		return x.Gud
	}
	return false
}

var File_ftp_proto protoreflect.FileDescriptor

var file_ftp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x66, 0x74, 0x70,
	0x22, 0x8e, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70,
	0x61, 0x72, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x68, 0x69, 0x73, 0x50, 0x61, 0x72,
	0x74, 0x22, 0x2d, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x67, 0x75, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x67, 0x75, 0x64,
	0x32, 0x2b, 0x0a, 0x03, 0x46, 0x54, 0x50, 0x12, 0x24, 0x0a, 0x06, 0x45, 0x6e, 0x76, 0x69, 0x61,
	0x72, 0x12, 0x0a, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0e, 0x2e,
	0x66, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x42, 0x05, 0x5a,
	0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ftp_proto_rawDescOnce sync.Once
	file_ftp_proto_rawDescData = file_ftp_proto_rawDesc
)

func file_ftp_proto_rawDescGZIP() []byte {
	file_ftp_proto_rawDescOnce.Do(func() {
		file_ftp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ftp_proto_rawDescData)
	})
	return file_ftp_proto_rawDescData
}

var file_ftp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ftp_proto_goTypes = []interface{}{
	(*Chunk)(nil),     // 0: ftp.Chunk
	(*Respuesta)(nil), // 1: ftp.Respuesta
}
var file_ftp_proto_depIdxs = []int32{
	0, // 0: ftp.FTP.Enviar:input_type -> ftp.Chunk
	1, // 1: ftp.FTP.Enviar:output_type -> ftp.Respuesta
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ftp_proto_init() }
func file_ftp_proto_init() {
	if File_ftp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ftp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_ftp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respuesta); i {
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
			RawDescriptor: file_ftp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ftp_proto_goTypes,
		DependencyIndexes: file_ftp_proto_depIdxs,
		MessageInfos:      file_ftp_proto_msgTypes,
	}.Build()
	File_ftp_proto = out.File
	file_ftp_proto_rawDesc = nil
	file_ftp_proto_goTypes = nil
	file_ftp_proto_depIdxs = nil
}
