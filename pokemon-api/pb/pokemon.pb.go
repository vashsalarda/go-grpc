// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.27.3
// source: pokemon.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_UNKNOWN  Type = 0
	Type_NORMAL   Type = 1
	Type_FIRE     Type = 2
	Type_WATER    Type = 3
	Type_GRASS    Type = 4
	Type_ELECTRIC Type = 5
	Type_ICE      Type = 6
	Type_FIGHTING Type = 7
	Type_POISON   Type = 8
	Type_GROUND   Type = 9
	Type_FLYING   Type = 10
	Type_PSYCHIC  Type = 11
	Type_BUG      Type = 12
	Type_ROCK     Type = 13
	Type_GHOST    Type = 14
	Type_DRAGON   Type = 15
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "NORMAL",
		2:  "FIRE",
		3:  "WATER",
		4:  "GRASS",
		5:  "ELECTRIC",
		6:  "ICE",
		7:  "FIGHTING",
		8:  "POISON",
		9:  "GROUND",
		10: "FLYING",
		11: "PSYCHIC",
		12: "BUG",
		13: "ROCK",
		14: "GHOST",
		15: "DRAGON",
	}
	Type_value = map[string]int32{
		"UNKNOWN":  0,
		"NORMAL":   1,
		"FIRE":     2,
		"WATER":    3,
		"GRASS":    4,
		"ELECTRIC": 5,
		"ICE":      6,
		"FIGHTING": 7,
		"POISON":   8,
		"GROUND":   9,
		"FLYING":   10,
		"PSYCHIC":  11,
		"BUG":      12,
		"ROCK":     13,
		"GHOST":    14,
		"DRAGON":   15,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pokemon_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_pokemon_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{0}
}

type PokemonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type []Type `protobuf:"varint,3,rep,packed,name=type,proto3,enum=pokemon.Type" json:"type,omitempty"`
}

func (x *PokemonRequest) Reset() {
	*x = PokemonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonRequest) ProtoMessage() {}

func (x *PokemonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonRequest.ProtoReflect.Descriptor instead.
func (*PokemonRequest) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{0}
}

func (x *PokemonRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PokemonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PokemonRequest) GetType() []Type {
	if x != nil {
		return x.Type
	}
	return nil
}

type Pokemon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type      []Type                 `protobuf:"varint,3,rep,packed,name=type,proto3,enum=pokemon.Type" json:"type,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Pokemon) Reset() {
	*x = Pokemon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pokemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pokemon) ProtoMessage() {}

func (x *Pokemon) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pokemon.ProtoReflect.Descriptor instead.
func (*Pokemon) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{1}
}

func (x *Pokemon) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pokemon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pokemon) GetType() []Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Pokemon) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Pokemon) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PokemonListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pokemon []*Pokemon `protobuf:"bytes,1,rep,name=pokemon,proto3" json:"pokemon,omitempty"`
}

func (x *PokemonListResponse) Reset() {
	*x = PokemonListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonListResponse) ProtoMessage() {}

func (x *PokemonListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonListResponse.ProtoReflect.Descriptor instead.
func (*PokemonListResponse) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{2}
}

func (x *PokemonListResponse) GetPokemon() []*Pokemon {
	if x != nil {
		return x.Pokemon
	}
	return nil
}

type PokemonFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PokemonFilter) Reset() {
	*x = PokemonFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonFilter) ProtoMessage() {}

func (x *PokemonFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonFilter.ProtoReflect.Descriptor instead.
func (*PokemonFilter) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{3}
}

type PokemonID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PokemonID) Reset() {
	*x = PokemonID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonID) ProtoMessage() {}

func (x *PokemonID) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonID.ProtoReflect.Descriptor instead.
func (*PokemonID) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{4}
}

func (x *PokemonID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PokemonUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type []Type `protobuf:"varint,3,rep,packed,name=type,proto3,enum=pokemon.Type" json:"type,omitempty"`
}

func (x *PokemonUpdateRequest) Reset() {
	*x = PokemonUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PokemonUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PokemonUpdateRequest) ProtoMessage() {}

func (x *PokemonUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PokemonUpdateRequest.ProtoReflect.Descriptor instead.
func (*PokemonUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{5}
}

func (x *PokemonUpdateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PokemonUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PokemonUpdateRequest) GetType() []Type {
	if x != nil {
		return x.Type
	}
	return nil
}

var File_pokemon_proto protoreflect.FileDescriptor

var file_pokemon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xc6, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x13, 0x50, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x50,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x09,
	0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x14, 0x50, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0xbf, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49,
	0x52, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x47, 0x52, 0x41, 0x53, 0x53, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4c,
	0x45, 0x43, 0x54, 0x52, 0x49, 0x43, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x43, 0x45, 0x10,
	0x06, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x47, 0x48, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x4f, 0x49, 0x53, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x47,
	0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4c, 0x59, 0x49, 0x4e,
	0x47, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x53, 0x59, 0x43, 0x48, 0x49, 0x43, 0x10, 0x0b,
	0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x47, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x43,
	0x4b, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x0e, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x52, 0x41, 0x47, 0x4f, 0x4e, 0x10, 0x0f, 0x32, 0x9e, 0x02, 0x0a, 0x07, 0x50,
	0x6f, 0x6b, 0x65, 0x64, 0x65, 0x78, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x04, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x70, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x52, 0x65, 0x61,
	0x64, 0x4f, 0x6e, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f,
	0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pokemon_proto_rawDescOnce sync.Once
	file_pokemon_proto_rawDescData = file_pokemon_proto_rawDesc
)

func file_pokemon_proto_rawDescGZIP() []byte {
	file_pokemon_proto_rawDescOnce.Do(func() {
		file_pokemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_pokemon_proto_rawDescData)
	})
	return file_pokemon_proto_rawDescData
}

var file_pokemon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pokemon_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pokemon_proto_goTypes = []interface{}{
	(Type)(0),                     // 0: pokemon.Type
	(*PokemonRequest)(nil),        // 1: pokemon.PokemonRequest
	(*Pokemon)(nil),               // 2: pokemon.Pokemon
	(*PokemonListResponse)(nil),   // 3: pokemon.PokemonListResponse
	(*PokemonFilter)(nil),         // 4: pokemon.PokemonFilter
	(*PokemonID)(nil),             // 5: pokemon.PokemonID
	(*PokemonUpdateRequest)(nil),  // 6: pokemon.PokemonUpdateRequest
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_pokemon_proto_depIdxs = []int32{
	0,  // 0: pokemon.PokemonRequest.type:type_name -> pokemon.Type
	0,  // 1: pokemon.Pokemon.type:type_name -> pokemon.Type
	7,  // 2: pokemon.Pokemon.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: pokemon.Pokemon.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 4: pokemon.PokemonListResponse.pokemon:type_name -> pokemon.Pokemon
	0,  // 5: pokemon.PokemonUpdateRequest.type:type_name -> pokemon.Type
	1,  // 6: pokemon.Pokedex.Create:input_type -> pokemon.PokemonRequest
	4,  // 7: pokemon.Pokedex.Read:input_type -> pokemon.PokemonFilter
	5,  // 8: pokemon.Pokedex.ReadOne:input_type -> pokemon.PokemonID
	6,  // 9: pokemon.Pokedex.Update:input_type -> pokemon.PokemonUpdateRequest
	5,  // 10: pokemon.Pokedex.Delete:input_type -> pokemon.PokemonID
	2,  // 11: pokemon.Pokedex.Create:output_type -> pokemon.Pokemon
	3,  // 12: pokemon.Pokedex.Read:output_type -> pokemon.PokemonListResponse
	2,  // 13: pokemon.Pokedex.ReadOne:output_type -> pokemon.Pokemon
	2,  // 14: pokemon.Pokedex.Update:output_type -> pokemon.Pokemon
	8,  // 15: pokemon.Pokedex.Delete:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pokemon_proto_init() }
func file_pokemon_proto_init() {
	if File_pokemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pokemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonRequest); i {
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
		file_pokemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pokemon); i {
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
		file_pokemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonListResponse); i {
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
		file_pokemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonFilter); i {
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
		file_pokemon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonID); i {
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
		file_pokemon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PokemonUpdateRequest); i {
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
			RawDescriptor: file_pokemon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pokemon_proto_goTypes,
		DependencyIndexes: file_pokemon_proto_depIdxs,
		EnumInfos:         file_pokemon_proto_enumTypes,
		MessageInfos:      file_pokemon_proto_msgTypes,
	}.Build()
	File_pokemon_proto = out.File
	file_pokemon_proto_rawDesc = nil
	file_pokemon_proto_goTypes = nil
	file_pokemon_proto_depIdxs = nil
}
