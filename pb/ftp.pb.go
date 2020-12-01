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
	Cliente  bool   `protobuf:"varint,4,opt,name=cliente,proto3" json:"cliente,omitempty"`
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

func (x *Chunk) GetCliente() bool {
	if x != nil {
		return x.Cliente
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

	Gud bool `protobuf:"varint,1,opt,name=gud,proto3" json:"gud,omitempty"`
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

func (x *Respuesta) GetGud() bool {
	if x != nil {
		return x.Gud
	}
	return false
}

type Nombre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Nombre) Reset() {
	*x = Nombre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nombre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nombre) ProtoMessage() {}

func (x *Nombre) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nombre.ProtoReflect.Descriptor instead.
func (*Nombre) Descriptor() ([]byte, []int) {
	return file_ftp_proto_rawDescGZIP(), []int{2}
}

func (x *Nombre) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Propuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lista       string `protobuf:"bytes,1,opt,name=lista,proto3" json:"lista,omitempty"`
	NombreLibro string `protobuf:"bytes,2,opt,name=NombreLibro,proto3" json:"NombreLibro,omitempty"`
}

func (x *Propuesta) Reset() {
	*x = Propuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propuesta) ProtoMessage() {}

func (x *Propuesta) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propuesta.ProtoReflect.Descriptor instead.
func (*Propuesta) Descriptor() ([]byte, []int) {
	return file_ftp_proto_rawDescGZIP(), []int{3}
}

func (x *Propuesta) GetLista() string {
	if x != nil {
		return x.Lista
	}
	return ""
}

func (x *Propuesta) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

type Aviso struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tiempo int32 `protobuf:"varint,1,opt,name=tiempo,proto3" json:"tiempo,omitempty"`
}

func (x *Aviso) Reset() {
	*x = Aviso{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aviso) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aviso) ProtoMessage() {}

func (x *Aviso) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aviso.ProtoReflect.Descriptor instead.
func (*Aviso) Descriptor() ([]byte, []int) {
	return file_ftp_proto_rawDescGZIP(), []int{4}
}

func (x *Aviso) GetTiempo() int32 {
	if x != nil {
		return x.Tiempo
	}
	return 0
}

type ListaLibros struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CantLibros  uint64   `protobuf:"varint,1,opt,name=CantLibros,proto3" json:"CantLibros,omitempty"`
	NombreLibro []string `protobuf:"bytes,2,rep,name=NombreLibro,proto3" json:"NombreLibro,omitempty"`
}

func (x *ListaLibros) Reset() {
	*x = ListaLibros{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListaLibros) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListaLibros) ProtoMessage() {}

func (x *ListaLibros) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListaLibros.ProtoReflect.Descriptor instead.
func (*ListaLibros) Descriptor() ([]byte, []int) {
	return file_ftp_proto_rawDescGZIP(), []int{5}
}

func (x *ListaLibros) GetCantLibros() uint64 {
	if x != nil {
		return x.CantLibros
	}
	return 0
}

func (x *ListaLibros) GetNombreLibro() []string {
	if x != nil {
		return x.NombreLibro
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_ftp_proto_rawDescGZIP(), []int{6}
}

var File_ftp_proto protoreflect.FileDescriptor

var file_ftp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x66, 0x74, 0x70,
	0x22, 0x92, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x69, 0x73,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x68, 0x69,
	0x73, 0x50, 0x61, 0x72, 0x74, 0x22, 0x1d, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x67, 0x75, 0x64, 0x22, 0x1c, 0x0a, 0x06, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x43, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x69, 0x73, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c,
	0x69, 0x62, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x22, 0x1f, 0x0a, 0x05, 0x41, 0x76, 0x69, 0x73, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6d, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x74, 0x69, 0x65, 0x6d, 0x70, 0x6f, 0x22, 0x4f, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x61, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x74, 0x4c,
	0x69, 0x62, 0x72, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x43, 0x61, 0x6e,
	0x74, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x51, 0x0a, 0x03, 0x46, 0x54, 0x50, 0x12, 0x24, 0x0a, 0x06, 0x45, 0x6e, 0x76,
	0x69, 0x61, 0x72, 0x12, 0x0a, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a,
	0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12,
	0x24, 0x0a, 0x09, 0x44, 0x65, 0x73, 0x63, 0x61, 0x72, 0x67, 0x61, 0x72, 0x12, 0x0b, 0x2e, 0x66,
	0x74, 0x70, 0x2e, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x1a, 0x0a, 0x2e, 0x66, 0x74, 0x70, 0x2e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x8d, 0x01, 0x0a, 0x0e, 0x46, 0x54, 0x50, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x12, 0x25, 0x0a, 0x07, 0x45, 0x6e, 0x76, 0x69,
	0x61, 0x72, 0x44, 0x12, 0x0a, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a,
	0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12,
	0x25, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x63, 0x61, 0x72, 0x67, 0x61, 0x72, 0x44, 0x12, 0x0b, 0x2e,
	0x66, 0x74, 0x70, 0x2e, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x1a, 0x0a, 0x2e, 0x66, 0x74, 0x70,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x2d, 0x0a, 0x0f, 0x41, 0x76, 0x69, 0x73, 0x6f, 0x45,
	0x73, 0x63, 0x72, 0x69, 0x74, 0x75, 0x72, 0x61, 0x44, 0x12, 0x0a, 0x2e, 0x66, 0x74, 0x70, 0x2e,
	0x41, 0x76, 0x69, 0x73, 0x6f, 0x1a, 0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x32, 0x98, 0x01, 0x0a, 0x03, 0x4c, 0x4f, 0x47, 0x12, 0x31, 0x0a,
	0x0f, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x12, 0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x1a, 0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x12, 0x31, 0x0a, 0x12, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x55, 0x62, 0x69,
	0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x4e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x1a, 0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65,
	0x73, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x0b, 0x50, 0x65, 0x64, 0x69, 0x72, 0x4c, 0x69, 0x62, 0x72,
	0x6f, 0x73, 0x12, 0x0a, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x66, 0x74, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73,
	0x32, 0xa6, 0x01, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x69, 0x64, 0x6f, 0x12, 0x32, 0x0a, 0x10, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f,
	0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x44, 0x12, 0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x13, 0x53, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x61, 0x72, 0x55, 0x62, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x44, 0x12, 0x0b,
	0x2e, 0x66, 0x74, 0x70, 0x2e, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x1a, 0x0e, 0x2e, 0x66, 0x74,
	0x70, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x0c, 0x50,
	0x65, 0x64, 0x69, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x44, 0x12, 0x0a, 0x2e, 0x66, 0x74,
	0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x66, 0x74, 0x70, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x61, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_ftp_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ftp_proto_goTypes = []interface{}{
	(*Chunk)(nil),       // 0: ftp.Chunk
	(*Respuesta)(nil),   // 1: ftp.Respuesta
	(*Nombre)(nil),      // 2: ftp.Nombre
	(*Propuesta)(nil),   // 3: ftp.Propuesta
	(*Aviso)(nil),       // 4: ftp.Aviso
	(*ListaLibros)(nil), // 5: ftp.ListaLibros
	(*Empty)(nil),       // 6: ftp.Empty
}
var file_ftp_proto_depIdxs = []int32{
	0,  // 0: ftp.FTP.Enviar:input_type -> ftp.Chunk
	2,  // 1: ftp.FTP.Descargar:input_type -> ftp.Nombre
	0,  // 2: ftp.FTPDistribuido.EnviarD:input_type -> ftp.Chunk
	2,  // 3: ftp.FTPDistribuido.DescargarD:input_type -> ftp.Nombre
	4,  // 4: ftp.FTPDistribuido.AvisoEscrituraD:input_type -> ftp.Aviso
	3,  // 5: ftp.LOG.EnviarPropuesta:input_type -> ftp.Propuesta
	2,  // 6: ftp.LOG.SolicitarUbicacion:input_type -> ftp.Nombre
	6,  // 7: ftp.LOG.PedirLibros:input_type -> ftp.Empty
	3,  // 8: ftp.LOGDistribuido.EnviarPropuestaD:input_type -> ftp.Propuesta
	2,  // 9: ftp.LOGDistribuido.SolicitarUbicacionD:input_type -> ftp.Nombre
	6,  // 10: ftp.LOGDistribuido.PedirLibrosD:input_type -> ftp.Empty
	1,  // 11: ftp.FTP.Enviar:output_type -> ftp.Respuesta
	0,  // 12: ftp.FTP.Descargar:output_type -> ftp.Chunk
	1,  // 13: ftp.FTPDistribuido.EnviarD:output_type -> ftp.Respuesta
	0,  // 14: ftp.FTPDistribuido.DescargarD:output_type -> ftp.Chunk
	1,  // 15: ftp.FTPDistribuido.AvisoEscrituraD:output_type -> ftp.Respuesta
	3,  // 16: ftp.LOG.EnviarPropuesta:output_type -> ftp.Propuesta
	3,  // 17: ftp.LOG.SolicitarUbicacion:output_type -> ftp.Propuesta
	5,  // 18: ftp.LOG.PedirLibros:output_type -> ftp.ListaLibros
	3,  // 19: ftp.LOGDistribuido.EnviarPropuestaD:output_type -> ftp.Propuesta
	3,  // 20: ftp.LOGDistribuido.SolicitarUbicacionD:output_type -> ftp.Propuesta
	5,  // 21: ftp.LOGDistribuido.PedirLibrosD:output_type -> ftp.ListaLibros
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
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
		file_ftp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nombre); i {
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
		file_ftp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propuesta); i {
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
		file_ftp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aviso); i {
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
		file_ftp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListaLibros); i {
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
		file_ftp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   4,
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
