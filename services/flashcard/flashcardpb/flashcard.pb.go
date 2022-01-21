// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: flashcard.proto

package flashcardpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Creator string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Users   []string `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{0}
}

func (x *Board) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Board) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Board) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetFlashcardByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlashcardId string `protobuf:"bytes,1,opt,name=flashcard_id,json=flashcardId,proto3" json:"flashcard_id,omitempty"`
}

func (x *GetFlashcardByIdRequest) Reset() {
	*x = GetFlashcardByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlashcardByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlashcardByIdRequest) ProtoMessage() {}

func (x *GetFlashcardByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlashcardByIdRequest.ProtoReflect.Descriptor instead.
func (*GetFlashcardByIdRequest) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{1}
}

func (x *GetFlashcardByIdRequest) GetFlashcardId() string {
	if x != nil {
		return x.FlashcardId
	}
	return ""
}

type GetFlashcardByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlashcardId string `protobuf:"bytes,1,opt,name=flashcard_id,json=flashcardId,proto3" json:"flashcard_id,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Answer      string `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
	Board       *Board `protobuf:"bytes,5,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *GetFlashcardByIdResponse) Reset() {
	*x = GetFlashcardByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlashcardByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlashcardByIdResponse) ProtoMessage() {}

func (x *GetFlashcardByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlashcardByIdResponse.ProtoReflect.Descriptor instead.
func (*GetFlashcardByIdResponse) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{2}
}

func (x *GetFlashcardByIdResponse) GetFlashcardId() string {
	if x != nil {
		return x.FlashcardId
	}
	return ""
}

func (x *GetFlashcardByIdResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetFlashcardByIdResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *GetFlashcardByIdResponse) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type GetFlashcardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Board   string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *GetFlashcardsRequest) Reset() {
	*x = GetFlashcardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlashcardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlashcardsRequest) ProtoMessage() {}

func (x *GetFlashcardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlashcardsRequest.ProtoReflect.Descriptor instead.
func (*GetFlashcardsRequest) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{3}
}

func (x *GetFlashcardsRequest) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *GetFlashcardsRequest) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

type CreateFlashcardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Board string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	Tag   string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CreateFlashcardRequest) Reset() {
	*x = CreateFlashcardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlashcardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlashcardRequest) ProtoMessage() {}

func (x *CreateFlashcardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlashcardRequest.ProtoReflect.Descriptor instead.
func (*CreateFlashcardRequest) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFlashcardRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateFlashcardRequest) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *CreateFlashcardRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type UpdateFlashcardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlashcardId string `protobuf:"bytes,1,opt,name=flashcard_id,json=flashcardId,proto3" json:"flashcard_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Answer      string `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
	Board       string `protobuf:"bytes,4,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *UpdateFlashcardRequest) Reset() {
	*x = UpdateFlashcardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlashcardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlashcardRequest) ProtoMessage() {}

func (x *UpdateFlashcardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlashcardRequest.ProtoReflect.Descriptor instead.
func (*UpdateFlashcardRequest) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFlashcardRequest) GetFlashcardId() string {
	if x != nil {
		return x.FlashcardId
	}
	return ""
}

func (x *UpdateFlashcardRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateFlashcardRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *UpdateFlashcardRequest) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

type DeleteFlashcardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlashcardId string `protobuf:"bytes,1,opt,name=flashcard_id,json=flashcardId,proto3" json:"flashcard_id,omitempty"`
}

func (x *DeleteFlashcardRequest) Reset() {
	*x = DeleteFlashcardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlashcardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlashcardRequest) ProtoMessage() {}

func (x *DeleteFlashcardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlashcardRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlashcardRequest) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFlashcardRequest) GetFlashcardId() string {
	if x != nil {
		return x.FlashcardId
	}
	return ""
}

type Default struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Default) Reset() {
	*x = Default{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flashcard_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Default) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Default) ProtoMessage() {}

func (x *Default) ProtoReflect() protoreflect.Message {
	mi := &file_flashcard_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Default.ProtoReflect.Descriptor instead.
func (*Default) Descriptor() ([]byte, []int) {
	return file_flashcard_proto_rawDescGZIP(), []int{7}
}

func (x *Default) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Default) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_flashcard_proto protoreflect.FileDescriptor

var file_flashcard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x70, 0x62, 0x22, 0x4b,
	0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3c, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6c,
	0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6c,
	0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61,
	0x72, 0x64, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x56, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x22, 0x7f, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x22, 0x3b, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73,
	0x68, 0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22,
	0x39, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc4, 0x03, 0x0a, 0x10, 0x46,
	0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x61, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x43, 0x61, 0x72, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x6c, 0x61, 0x73,
	0x68, 0x63, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x63, 0x61, 0x72, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x21, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61,
	0x72, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72,
	0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x4e, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x63, 0x61, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x6c, 0x61, 0x73,
	0x68, 0x63, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x63, 0x61, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x6c, 0x61, 0x73,
	0x68, 0x63, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x63, 0x61, 0x72, 0x64, 0x12, 0x23, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72, 0x64,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x6c, 0x61, 0x73,
	0x68, 0x63, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x63, 0x61, 0x72,
	0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flashcard_proto_rawDescOnce sync.Once
	file_flashcard_proto_rawDescData = file_flashcard_proto_rawDesc
)

func file_flashcard_proto_rawDescGZIP() []byte {
	file_flashcard_proto_rawDescOnce.Do(func() {
		file_flashcard_proto_rawDescData = protoimpl.X.CompressGZIP(file_flashcard_proto_rawDescData)
	})
	return file_flashcard_proto_rawDescData
}

var file_flashcard_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_flashcard_proto_goTypes = []interface{}{
	(*Board)(nil),                    // 0: flashcardpb.Board
	(*GetFlashcardByIdRequest)(nil),  // 1: flashcardpb.GetFlashcardByIdRequest
	(*GetFlashcardByIdResponse)(nil), // 2: flashcardpb.GetFlashcardByIdResponse
	(*GetFlashcardsRequest)(nil),     // 3: flashcardpb.GetFlashcardsRequest
	(*CreateFlashcardRequest)(nil),   // 4: flashcardpb.CreateFlashcardRequest
	(*UpdateFlashcardRequest)(nil),   // 5: flashcardpb.UpdateFlashcardRequest
	(*DeleteFlashcardRequest)(nil),   // 6: flashcardpb.DeleteFlashcardRequest
	(*Default)(nil),                  // 7: flashcardpb.Default
}
var file_flashcard_proto_depIdxs = []int32{
	0, // 0: flashcardpb.GetFlashcardByIdResponse.board:type_name -> flashcardpb.Board
	1, // 1: flashcardpb.FlashcardService.GetFlashCardById:input_type -> flashcardpb.GetFlashcardByIdRequest
	3, // 2: flashcardpb.FlashcardService.GetFlashcards:input_type -> flashcardpb.GetFlashcardsRequest
	4, // 3: flashcardpb.FlashcardService.CreateFlashcard:input_type -> flashcardpb.CreateFlashcardRequest
	5, // 4: flashcardpb.FlashcardService.UpdateFlashcard:input_type -> flashcardpb.UpdateFlashcardRequest
	6, // 5: flashcardpb.FlashcardService.DeleteFlashcard:input_type -> flashcardpb.DeleteFlashcardRequest
	2, // 6: flashcardpb.FlashcardService.GetFlashCardById:output_type -> flashcardpb.GetFlashcardByIdResponse
	2, // 7: flashcardpb.FlashcardService.GetFlashcards:output_type -> flashcardpb.GetFlashcardByIdResponse
	7, // 8: flashcardpb.FlashcardService.CreateFlashcard:output_type -> flashcardpb.Default
	7, // 9: flashcardpb.FlashcardService.UpdateFlashcard:output_type -> flashcardpb.Default
	7, // 10: flashcardpb.FlashcardService.DeleteFlashcard:output_type -> flashcardpb.Default
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flashcard_proto_init() }
func file_flashcard_proto_init() {
	if File_flashcard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flashcard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_flashcard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlashcardByIdRequest); i {
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
		file_flashcard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlashcardByIdResponse); i {
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
		file_flashcard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlashcardsRequest); i {
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
		file_flashcard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFlashcardRequest); i {
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
		file_flashcard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlashcardRequest); i {
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
		file_flashcard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlashcardRequest); i {
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
		file_flashcard_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Default); i {
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
			RawDescriptor: file_flashcard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flashcard_proto_goTypes,
		DependencyIndexes: file_flashcard_proto_depIdxs,
		MessageInfos:      file_flashcard_proto_msgTypes,
	}.Build()
	File_flashcard_proto = out.File
	file_flashcard_proto_rawDesc = nil
	file_flashcard_proto_goTypes = nil
	file_flashcard_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FlashcardServiceClient is the client API for FlashcardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlashcardServiceClient interface {
	GetFlashCardById(ctx context.Context, in *GetFlashcardByIdRequest, opts ...grpc.CallOption) (*GetFlashcardByIdResponse, error)
	GetFlashcards(ctx context.Context, in *GetFlashcardsRequest, opts ...grpc.CallOption) (FlashcardService_GetFlashcardsClient, error)
	CreateFlashcard(ctx context.Context, in *CreateFlashcardRequest, opts ...grpc.CallOption) (*Default, error)
	UpdateFlashcard(ctx context.Context, in *UpdateFlashcardRequest, opts ...grpc.CallOption) (*Default, error)
	DeleteFlashcard(ctx context.Context, in *DeleteFlashcardRequest, opts ...grpc.CallOption) (*Default, error)
}

type flashcardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlashcardServiceClient(cc grpc.ClientConnInterface) FlashcardServiceClient {
	return &flashcardServiceClient{cc}
}

func (c *flashcardServiceClient) GetFlashCardById(ctx context.Context, in *GetFlashcardByIdRequest, opts ...grpc.CallOption) (*GetFlashcardByIdResponse, error) {
	out := new(GetFlashcardByIdResponse)
	err := c.cc.Invoke(ctx, "/flashcardpb.FlashcardService/GetFlashCardById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashcardServiceClient) GetFlashcards(ctx context.Context, in *GetFlashcardsRequest, opts ...grpc.CallOption) (FlashcardService_GetFlashcardsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FlashcardService_serviceDesc.Streams[0], "/flashcardpb.FlashcardService/GetFlashcards", opts...)
	if err != nil {
		return nil, err
	}
	x := &flashcardServiceGetFlashcardsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlashcardService_GetFlashcardsClient interface {
	Recv() (*GetFlashcardByIdResponse, error)
	grpc.ClientStream
}

type flashcardServiceGetFlashcardsClient struct {
	grpc.ClientStream
}

func (x *flashcardServiceGetFlashcardsClient) Recv() (*GetFlashcardByIdResponse, error) {
	m := new(GetFlashcardByIdResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flashcardServiceClient) CreateFlashcard(ctx context.Context, in *CreateFlashcardRequest, opts ...grpc.CallOption) (*Default, error) {
	out := new(Default)
	err := c.cc.Invoke(ctx, "/flashcardpb.FlashcardService/CreateFlashcard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashcardServiceClient) UpdateFlashcard(ctx context.Context, in *UpdateFlashcardRequest, opts ...grpc.CallOption) (*Default, error) {
	out := new(Default)
	err := c.cc.Invoke(ctx, "/flashcardpb.FlashcardService/UpdateFlashcard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flashcardServiceClient) DeleteFlashcard(ctx context.Context, in *DeleteFlashcardRequest, opts ...grpc.CallOption) (*Default, error) {
	out := new(Default)
	err := c.cc.Invoke(ctx, "/flashcardpb.FlashcardService/DeleteFlashcard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlashcardServiceServer is the server API for FlashcardService service.
type FlashcardServiceServer interface {
	GetFlashCardById(context.Context, *GetFlashcardByIdRequest) (*GetFlashcardByIdResponse, error)
	GetFlashcards(*GetFlashcardsRequest, FlashcardService_GetFlashcardsServer) error
	CreateFlashcard(context.Context, *CreateFlashcardRequest) (*Default, error)
	UpdateFlashcard(context.Context, *UpdateFlashcardRequest) (*Default, error)
	DeleteFlashcard(context.Context, *DeleteFlashcardRequest) (*Default, error)
}

// UnimplementedFlashcardServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFlashcardServiceServer struct {
}

func (*UnimplementedFlashcardServiceServer) GetFlashCardById(context.Context, *GetFlashcardByIdRequest) (*GetFlashcardByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlashCardById not implemented")
}
func (*UnimplementedFlashcardServiceServer) GetFlashcards(*GetFlashcardsRequest, FlashcardService_GetFlashcardsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFlashcards not implemented")
}
func (*UnimplementedFlashcardServiceServer) CreateFlashcard(context.Context, *CreateFlashcardRequest) (*Default, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlashcard not implemented")
}
func (*UnimplementedFlashcardServiceServer) UpdateFlashcard(context.Context, *UpdateFlashcardRequest) (*Default, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlashcard not implemented")
}
func (*UnimplementedFlashcardServiceServer) DeleteFlashcard(context.Context, *DeleteFlashcardRequest) (*Default, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFlashcard not implemented")
}

func RegisterFlashcardServiceServer(s *grpc.Server, srv FlashcardServiceServer) {
	s.RegisterService(&_FlashcardService_serviceDesc, srv)
}

func _FlashcardService_GetFlashCardById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlashcardByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashcardServiceServer).GetFlashCardById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flashcardpb.FlashcardService/GetFlashCardById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashcardServiceServer).GetFlashCardById(ctx, req.(*GetFlashcardByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashcardService_GetFlashcards_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFlashcardsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlashcardServiceServer).GetFlashcards(m, &flashcardServiceGetFlashcardsServer{stream})
}

type FlashcardService_GetFlashcardsServer interface {
	Send(*GetFlashcardByIdResponse) error
	grpc.ServerStream
}

type flashcardServiceGetFlashcardsServer struct {
	grpc.ServerStream
}

func (x *flashcardServiceGetFlashcardsServer) Send(m *GetFlashcardByIdResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FlashcardService_CreateFlashcard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlashcardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashcardServiceServer).CreateFlashcard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flashcardpb.FlashcardService/CreateFlashcard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashcardServiceServer).CreateFlashcard(ctx, req.(*CreateFlashcardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashcardService_UpdateFlashcard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlashcardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashcardServiceServer).UpdateFlashcard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flashcardpb.FlashcardService/UpdateFlashcard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashcardServiceServer).UpdateFlashcard(ctx, req.(*UpdateFlashcardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlashcardService_DeleteFlashcard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlashcardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlashcardServiceServer).DeleteFlashcard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flashcardpb.FlashcardService/DeleteFlashcard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlashcardServiceServer).DeleteFlashcard(ctx, req.(*DeleteFlashcardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlashcardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flashcardpb.FlashcardService",
	HandlerType: (*FlashcardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlashCardById",
			Handler:    _FlashcardService_GetFlashCardById_Handler,
		},
		{
			MethodName: "CreateFlashcard",
			Handler:    _FlashcardService_CreateFlashcard_Handler,
		},
		{
			MethodName: "UpdateFlashcard",
			Handler:    _FlashcardService_UpdateFlashcard_Handler,
		},
		{
			MethodName: "DeleteFlashcard",
			Handler:    _FlashcardService_DeleteFlashcard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFlashcards",
			Handler:       _FlashcardService_GetFlashcards_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "flashcard.proto",
}
