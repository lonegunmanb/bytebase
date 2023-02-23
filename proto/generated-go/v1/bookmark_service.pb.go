// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: v1/bookmark_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource of the bookmark.
	// Format: users/{user}, user is a server-generated unique IDs.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The bookmark to create.
	Bookmark *Bookmark `protobuf:"bytes,2,opt,name=bookmark,proto3" json:"bookmark,omitempty"`
}

func (x *CreateBookmarkRequest) Reset() {
	*x = CreateBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookmarkRequest) ProtoMessage() {}

func (x *CreateBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookmarkRequest.ProtoReflect.Descriptor instead.
func (*CreateBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookmarkRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateBookmarkRequest) GetBookmark() *Bookmark {
	if x != nil {
		return x.Bookmark
	}
	return nil
}

type DeleteBookmarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the bookmark to delete.
	// Format: users/{user}/bookmarks/{bookmark}, user and bookmark are server-generated unique IDs.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteBookmarkRequest) Reset() {
	*x = DeleteBookmarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookmarkRequest) ProtoMessage() {}

func (x *DeleteBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookmarkRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteBookmarkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListBookmarksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource of the bookmark.
	// Format: users/{user}, user is a server-generated unique ID.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Not used. The maximum number of bookmarks to return. The service may return fewer than
	// this value.
	// If unspecified, at most 50 bookmarks will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Not used. A page token, received from a previous `ListBookmarks` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListBookmarks` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBookmarksRequest) Reset() {
	*x = ListBookmarksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookmarksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookmarksRequest) ProtoMessage() {}

func (x *ListBookmarksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookmarksRequest.ProtoReflect.Descriptor instead.
func (*ListBookmarksRequest) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBookmarksRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListBookmarksRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBookmarksRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBookmarksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of bookmarks.
	Bookmarks []*Bookmark `protobuf:"bytes,1,rep,name=bookmarks,proto3" json:"bookmarks,omitempty"`
	// Not used. A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBookmarksResponse) Reset() {
	*x = ListBookmarksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookmarksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookmarksResponse) ProtoMessage() {}

func (x *ListBookmarksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookmarksResponse.ProtoReflect.Descriptor instead.
func (*ListBookmarksResponse) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListBookmarksResponse) GetBookmarks() []*Bookmark {
	if x != nil {
		return x.Bookmarks
	}
	return nil
}

func (x *ListBookmarksResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type Bookmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the bookmark.
	// Format: users/{user}/bookmarks/{bookmark}, user and bookmark are server-generated unique IDs.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource link of the bookmark. Only support issue link for now.
	// Format:
	// Issue: /issue/slug(issue_name)-{issue_uid}
	// Example: /issue/start-here-add-email-column-to-employee-table-101
	Link string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Bookmark) Reset() {
	*x = Bookmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookmark_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookmark) ProtoMessage() {}

func (x *Bookmark) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookmark_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookmark.ProtoReflect.Descriptor instead.
func (*Bookmark) Descriptor() ([]byte, []int) {
	return file_v1_bookmark_service_proto_rawDescGZIP(), []int{4}
}

func (x *Bookmark) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bookmark) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_v1_bookmark_service_proto protoreflect.FileDescriptor

var file_v1_bookmark_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x68, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x37, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02,
	0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x74, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x3e, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32,
	0xaa, 0x03, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x22, 0x42, 0xda, 0x41, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x7b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x2d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2f,
	0x2a, 0x7d, 0x12, 0x87, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x73, 0x12, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0xda, 0x41, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x2a, 0x7d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x42, 0x11, 0x5a, 0x0f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bookmark_service_proto_rawDescOnce sync.Once
	file_v1_bookmark_service_proto_rawDescData = file_v1_bookmark_service_proto_rawDesc
)

func file_v1_bookmark_service_proto_rawDescGZIP() []byte {
	file_v1_bookmark_service_proto_rawDescOnce.Do(func() {
		file_v1_bookmark_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bookmark_service_proto_rawDescData)
	})
	return file_v1_bookmark_service_proto_rawDescData
}

var file_v1_bookmark_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_bookmark_service_proto_goTypes = []interface{}{
	(*CreateBookmarkRequest)(nil), // 0: bytebase.v1.CreateBookmarkRequest
	(*DeleteBookmarkRequest)(nil), // 1: bytebase.v1.DeleteBookmarkRequest
	(*ListBookmarksRequest)(nil),  // 2: bytebase.v1.ListBookmarksRequest
	(*ListBookmarksResponse)(nil), // 3: bytebase.v1.ListBookmarksResponse
	(*Bookmark)(nil),              // 4: bytebase.v1.Bookmark
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_v1_bookmark_service_proto_depIdxs = []int32{
	4, // 0: bytebase.v1.CreateBookmarkRequest.bookmark:type_name -> bytebase.v1.Bookmark
	4, // 1: bytebase.v1.ListBookmarksResponse.bookmarks:type_name -> bytebase.v1.Bookmark
	0, // 2: bytebase.v1.BookmarkService.CreateBookmark:input_type -> bytebase.v1.CreateBookmarkRequest
	1, // 3: bytebase.v1.BookmarkService.DeleteBookmark:input_type -> bytebase.v1.DeleteBookmarkRequest
	2, // 4: bytebase.v1.BookmarkService.ListBookmarks:input_type -> bytebase.v1.ListBookmarksRequest
	4, // 5: bytebase.v1.BookmarkService.CreateBookmark:output_type -> bytebase.v1.Bookmark
	5, // 6: bytebase.v1.BookmarkService.DeleteBookmark:output_type -> google.protobuf.Empty
	3, // 7: bytebase.v1.BookmarkService.ListBookmarks:output_type -> bytebase.v1.ListBookmarksResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_bookmark_service_proto_init() }
func file_v1_bookmark_service_proto_init() {
	if File_v1_bookmark_service_proto != nil {
		return
	}
	file_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_bookmark_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookmarkRequest); i {
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
		file_v1_bookmark_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookmarkRequest); i {
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
		file_v1_bookmark_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookmarksRequest); i {
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
		file_v1_bookmark_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookmarksResponse); i {
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
		file_v1_bookmark_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookmark); i {
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
			RawDescriptor: file_v1_bookmark_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bookmark_service_proto_goTypes,
		DependencyIndexes: file_v1_bookmark_service_proto_depIdxs,
		MessageInfos:      file_v1_bookmark_service_proto_msgTypes,
	}.Build()
	File_v1_bookmark_service_proto = out.File
	file_v1_bookmark_service_proto_rawDesc = nil
	file_v1_bookmark_service_proto_goTypes = nil
	file_v1_bookmark_service_proto_depIdxs = nil
}