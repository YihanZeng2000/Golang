// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: bookstore.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 书架资源。
type Shelf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 书架的唯一 id。
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 书架上的主题(小说、诗歌等)。
	Theme string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
	// 书架的大小
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Shelf) Reset() {
	*x = Shelf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shelf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shelf) ProtoMessage() {}

func (x *Shelf) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shelf.ProtoReflect.Descriptor instead.
func (*Shelf) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *Shelf) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Shelf) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *Shelf) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// 图书资源。
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 图书的唯一 id。
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 图书的作者。
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// 书名
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// 调用 ListShelves 返回的响应消息。
type ListShelvesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 书店里的书架。
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
}

func (x *ListShelvesResponse) Reset() {
	*x = ListShelvesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShelvesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesResponse) ProtoMessage() {}

func (x *ListShelvesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShelvesResponse.ProtoReflect.Descriptor instead.
func (*ListShelvesResponse) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{2}
}

func (x *ListShelvesResponse) GetShelves() []*Shelf {
	if x != nil {
		return x.Shelves
	}
	return nil
}

// CreateShelf 方法的请求消息。
type CreateShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要创建的书架数据
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *CreateShelfRequest) Reset() {
	*x = CreateShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfRequest) ProtoMessage() {}

func (x *CreateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelfRequest.ProtoReflect.Descriptor instead.
func (*CreateShelfRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *CreateShelfRequest) GetShelf() *Shelf {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// GetShelf 方法的请求消息。
type GetShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要检索的书架的ID。
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *GetShelfRequest) Reset() {
	*x = GetShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfRequest) ProtoMessage() {}

func (x *GetShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelfRequest.ProtoReflect.Descriptor instead.
func (*GetShelfRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{4}
}

func (x *GetShelfRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// DeleteShelf 方法的请求消息。
type DeleteShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要删除的书架的ID
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *DeleteShelfRequest) Reset() {
	*x = DeleteShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfRequest) ProtoMessage() {}

func (x *DeleteShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelfRequest.ProtoReflect.Descriptor instead.
func (*DeleteShelfRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteShelfRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// Request message for ListBooks method.
type ListBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the shelf which books to list.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *ListBooksRequest) Reset() {
	*x = ListBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksRequest) ProtoMessage() {}

func (x *ListBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksRequest.ProtoReflect.Descriptor instead.
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{6}
}

func (x *ListBooksRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// ListBooks 方法的请求消息。
type ListBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 书架上的图书。
	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResponse.ProtoReflect.Descriptor instead.
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{7}
}

func (x *ListBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

// CreateBook 方法的请求消息。
type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要在其上创建图书的书架的 ID。
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// 在书架上创建的图书数据。
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{8}
}

func (x *CreateBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *CreateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

// GetBook 方法的请求消息。
type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要从中检索图书的书架的 ID。
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// 要检索的图书的 ID。
	Book int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{9}
}

func (x *GetBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *GetBookRequest) GetBook() int64 {
	if x != nil {
		return x.Book
	}
	return 0
}

// DeleteBook 方法的请求消息。
type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要从中删除图书的书架的 ID。
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// 要删除的图书的 ID。
	Book int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *DeleteBookRequest) GetBook() int64 {
	if x != nil {
		return x.Book
	}
	return 0
}

var File_bookstore_proto protoreflect.FileDescriptor

var file_bookstore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x05, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x41, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x07, 0x73, 0x68, 0x65, 0x6c,
	0x76, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65,
	0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x2a, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x22, 0x3a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x4e, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x3a, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x68, 0x65, 0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x3d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68,
	0x65, 0x6c, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x32, 0x9d, 0x06, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65,
	0x73, 0x12, 0x5a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x68, 0x65, 0x6c,
	0x66, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x3a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x55, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68,
	0x65, 0x6c, 0x66, 0x7d, 0x12, 0x61, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x2f,
	0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x7d, 0x12, 0x69, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x7d, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x64, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22,
	0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x7d, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x3a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x5f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65,
	0x6c, 0x76, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x7d, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x7d, 0x12, 0x6c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x2a, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x76,
	0x65, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x7d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x7d, 0x42, 0x4a, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x6c,
	0x69, 0x77, 0x65, 0x6e, 0x7a, 0x68, 0x6f, 0x75, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x42, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookstore_proto_rawDescOnce sync.Once
	file_bookstore_proto_rawDescData = file_bookstore_proto_rawDesc
)

func file_bookstore_proto_rawDescGZIP() []byte {
	file_bookstore_proto_rawDescOnce.Do(func() {
		file_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookstore_proto_rawDescData)
	})
	return file_bookstore_proto_rawDescData
}

var file_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_bookstore_proto_goTypes = []interface{}{
	(*Shelf)(nil),               // 0: bookstore.Shelf
	(*Book)(nil),                // 1: bookstore.Book
	(*ListShelvesResponse)(nil), // 2: bookstore.ListShelvesResponse
	(*CreateShelfRequest)(nil),  // 3: bookstore.CreateShelfRequest
	(*GetShelfRequest)(nil),     // 4: bookstore.GetShelfRequest
	(*DeleteShelfRequest)(nil),  // 5: bookstore.DeleteShelfRequest
	(*ListBooksRequest)(nil),    // 6: bookstore.ListBooksRequest
	(*ListBooksResponse)(nil),   // 7: bookstore.ListBooksResponse
	(*CreateBookRequest)(nil),   // 8: bookstore.CreateBookRequest
	(*GetBookRequest)(nil),      // 9: bookstore.GetBookRequest
	(*DeleteBookRequest)(nil),   // 10: bookstore.DeleteBookRequest
	(*empty.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_bookstore_proto_depIdxs = []int32{
	0,  // 0: bookstore.ListShelvesResponse.shelves:type_name -> bookstore.Shelf
	0,  // 1: bookstore.CreateShelfRequest.shelf:type_name -> bookstore.Shelf
	1,  // 2: bookstore.ListBooksResponse.books:type_name -> bookstore.Book
	1,  // 3: bookstore.CreateBookRequest.book:type_name -> bookstore.Book
	11, // 4: bookstore.Bookstore.ListShelves:input_type -> google.protobuf.Empty
	3,  // 5: bookstore.Bookstore.CreateShelf:input_type -> bookstore.CreateShelfRequest
	4,  // 6: bookstore.Bookstore.GetShelf:input_type -> bookstore.GetShelfRequest
	5,  // 7: bookstore.Bookstore.DeleteShelf:input_type -> bookstore.DeleteShelfRequest
	6,  // 8: bookstore.Bookstore.ListBooks:input_type -> bookstore.ListBooksRequest
	8,  // 9: bookstore.Bookstore.CreateBook:input_type -> bookstore.CreateBookRequest
	9,  // 10: bookstore.Bookstore.GetBook:input_type -> bookstore.GetBookRequest
	10, // 11: bookstore.Bookstore.DeleteBook:input_type -> bookstore.DeleteBookRequest
	2,  // 12: bookstore.Bookstore.ListShelves:output_type -> bookstore.ListShelvesResponse
	0,  // 13: bookstore.Bookstore.CreateShelf:output_type -> bookstore.Shelf
	0,  // 14: bookstore.Bookstore.GetShelf:output_type -> bookstore.Shelf
	11, // 15: bookstore.Bookstore.DeleteShelf:output_type -> google.protobuf.Empty
	7,  // 16: bookstore.Bookstore.ListBooks:output_type -> bookstore.ListBooksResponse
	1,  // 17: bookstore.Bookstore.CreateBook:output_type -> bookstore.Book
	1,  // 18: bookstore.Bookstore.GetBook:output_type -> bookstore.Book
	11, // 19: bookstore.Bookstore.DeleteBook:output_type -> google.protobuf.Empty
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_bookstore_proto_init() }
func file_bookstore_proto_init() {
	if File_bookstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shelf); i {
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
		file_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShelvesResponse); i {
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
		file_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelfRequest); i {
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
		file_bookstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelfRequest); i {
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
		file_bookstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelfRequest); i {
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
		file_bookstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksRequest); i {
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
		file_bookstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksResponse); i {
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
		file_bookstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_bookstore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequest); i {
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
		file_bookstore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
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
			RawDescriptor: file_bookstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookstore_proto_goTypes,
		DependencyIndexes: file_bookstore_proto_depIdxs,
		MessageInfos:      file_bookstore_proto_msgTypes,
	}.Build()
	File_bookstore_proto = out.File
	file_bookstore_proto_rawDesc = nil
	file_bookstore_proto_goTypes = nil
	file_bookstore_proto_depIdxs = nil
}
