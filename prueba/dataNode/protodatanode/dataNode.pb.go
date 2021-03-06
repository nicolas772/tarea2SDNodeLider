// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: dataNode/protodatanode/dataNode.proto

package dataNode

import (
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

type InfoJugada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumeroJugador int32 `protobuf:"varint,1,opt,name=numeroJugador,proto3" json:"numeroJugador,omitempty"`
	Ronda         int32 `protobuf:"varint,2,opt,name=ronda,proto3" json:"ronda,omitempty"`
	Jugada        int32 `protobuf:"varint,3,opt,name=jugada,proto3" json:"jugada,omitempty"`
}

func (x *InfoJugada) Reset() {
	*x = InfoJugada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataNode_protodatanode_dataNode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoJugada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoJugada) ProtoMessage() {}

func (x *InfoJugada) ProtoReflect() protoreflect.Message {
	mi := &file_dataNode_protodatanode_dataNode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoJugada.ProtoReflect.Descriptor instead.
func (*InfoJugada) Descriptor() ([]byte, []int) {
	return file_dataNode_protodatanode_dataNode_proto_rawDescGZIP(), []int{0}
}

func (x *InfoJugada) GetNumeroJugador() int32 {
	if x != nil {
		return x.NumeroJugador
	}
	return 0
}

func (x *InfoJugada) GetRonda() int32 {
	if x != nil {
		return x.Ronda
	}
	return 0
}

func (x *InfoJugada) GetJugada() int32 {
	if x != nil {
		return x.Jugada
	}
	return 0
}

type Registro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jugadas []int32 `protobuf:"varint,1,rep,packed,name=jugadas,proto3" json:"jugadas,omitempty"`
}

func (x *Registro) Reset() {
	*x = Registro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataNode_protodatanode_dataNode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registro) ProtoMessage() {}

func (x *Registro) ProtoReflect() protoreflect.Message {
	mi := &file_dataNode_protodatanode_dataNode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registro.ProtoReflect.Descriptor instead.
func (*Registro) Descriptor() ([]byte, []int) {
	return file_dataNode_protodatanode_dataNode_proto_rawDescGZIP(), []int{1}
}

func (x *Registro) GetJugadas() []int32 {
	if x != nil {
		return x.Jugadas
	}
	return nil
}

var File_dataNode_protodatanode_dataNode_proto protoreflect.FileDescriptor

var file_dataNode_protodatanode_dataNode_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x61,
	0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x60, 0x0a, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x4a, 0x75,
	0x67, 0x61, 0x64, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x4a, 0x75,
	0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x6e, 0x64, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x6e, 0x64, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x22, 0x24, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73, 0x32, 0x54,
	0x0a, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x72, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61,
	0x73, 0x12, 0x41, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x1a, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x6f, 0x42, 0x28, 0x5a, 0x26, 0x73, 0x69, 0x73, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dataNode_protodatanode_dataNode_proto_rawDescOnce sync.Once
	file_dataNode_protodatanode_dataNode_proto_rawDescData = file_dataNode_protodatanode_dataNode_proto_rawDesc
)

func file_dataNode_protodatanode_dataNode_proto_rawDescGZIP() []byte {
	file_dataNode_protodatanode_dataNode_proto_rawDescOnce.Do(func() {
		file_dataNode_protodatanode_dataNode_proto_rawDescData = protoimpl.X.CompressGZIP(file_dataNode_protodatanode_dataNode_proto_rawDescData)
	})
	return file_dataNode_protodatanode_dataNode_proto_rawDescData
}

var file_dataNode_protodatanode_dataNode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dataNode_protodatanode_dataNode_proto_goTypes = []interface{}{
	(*InfoJugada)(nil), // 0: protodatanode.infoJugada
	(*Registro)(nil),   // 1: protodatanode.registro
}
var file_dataNode_protodatanode_dataNode_proto_depIdxs = []int32{
	0, // 0: protodatanode.informarJugadas.infoJugadas:input_type -> protodatanode.infoJugada
	1, // 1: protodatanode.informarJugadas.infoJugadas:output_type -> protodatanode.registro
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dataNode_protodatanode_dataNode_proto_init() }
func file_dataNode_protodatanode_dataNode_proto_init() {
	if File_dataNode_protodatanode_dataNode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dataNode_protodatanode_dataNode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoJugada); i {
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
		file_dataNode_protodatanode_dataNode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registro); i {
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
			RawDescriptor: file_dataNode_protodatanode_dataNode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dataNode_protodatanode_dataNode_proto_goTypes,
		DependencyIndexes: file_dataNode_protodatanode_dataNode_proto_depIdxs,
		MessageInfos:      file_dataNode_protodatanode_dataNode_proto_msgTypes,
	}.Build()
	File_dataNode_protodatanode_dataNode_proto = out.File
	file_dataNode_protodatanode_dataNode_proto_rawDesc = nil
	file_dataNode_protodatanode_dataNode_proto_goTypes = nil
	file_dataNode_protodatanode_dataNode_proto_depIdxs = nil
}
