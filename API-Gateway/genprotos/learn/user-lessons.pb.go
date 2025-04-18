// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: user-lessons.proto

package learn

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

// Foydalanuvchi darslarini qo'shish uchun xabar
type AddUserLessonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // Foydalanuvchi ID'si
	LessonId    string `protobuf:"bytes,2,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`          // Dars ID'si
	CompletedAt string `protobuf:"bytes,3,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"` // Tugatilgan vaqti
}

func (x *AddUserLessonRequest) Reset() {
	*x = AddUserLessonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_lessons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserLessonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserLessonRequest) ProtoMessage() {}

func (x *AddUserLessonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_lessons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserLessonRequest.ProtoReflect.Descriptor instead.
func (*AddUserLessonRequest) Descriptor() ([]byte, []int) {
	return file_user_lessons_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserLessonRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddUserLessonRequest) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *AddUserLessonRequest) GetCompletedAt() string {
	if x != nil {
		return x.CompletedAt
	}
	return ""
}

// Foydalanuvchi darslarini qo'shish javobi
type AddUserLessonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // Yaratilgan foydalanuvchi darsining ID'si
}

func (x *AddUserLessonResponse) Reset() {
	*x = AddUserLessonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_lessons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserLessonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserLessonResponse) ProtoMessage() {}

func (x *AddUserLessonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_lessons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserLessonResponse.ProtoReflect.Descriptor instead.
func (*AddUserLessonResponse) Descriptor() ([]byte, []int) {
	return file_user_lessons_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserLessonResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Foydalanuvchi darslarini olish uchun xabar
type GetUserLessonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // Foydalanuvchi ID'si
}

func (x *GetUserLessonsRequest) Reset() {
	*x = GetUserLessonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_lessons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLessonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLessonsRequest) ProtoMessage() {}

func (x *GetUserLessonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_lessons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLessonsRequest.ProtoReflect.Descriptor instead.
func (*GetUserLessonsRequest) Descriptor() ([]byte, []int) {
	return file_user_lessons_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserLessonsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Foydalanuvchi darslarini olish javobi
type GetUserLessonsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLessons []*UserLesson `protobuf:"bytes,1,rep,name=user_lessons,json=userLessons,proto3" json:"user_lessons,omitempty"` // Foydalanuvchi darslari ro'yxati
}

func (x *GetUserLessonsResponse) Reset() {
	*x = GetUserLessonsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_lessons_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLessonsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLessonsResponse) ProtoMessage() {}

func (x *GetUserLessonsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_lessons_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLessonsResponse.ProtoReflect.Descriptor instead.
func (*GetUserLessonsResponse) Descriptor() ([]byte, []int) {
	return file_user_lessons_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserLessonsResponse) GetUserLessons() []*UserLesson {
	if x != nil {
		return x.UserLessons
	}
	return nil
}

// Foydalanuvchi darsi haqida ma'lumotni aniqlash uchun xabar
type UserLesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                     // Foydalanuvchi darsi ID'si
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // Foydalanuvchi ID'si
	LessonId    string `protobuf:"bytes,3,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`          // Dars ID'si
	CompletedAt string `protobuf:"bytes,4,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"` // Tugatilgan vaqti
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`       // Yaratilgan vaqti
	UpdatedAt   string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`       // O'zgartirilgan vaqti
}

func (x *UserLesson) Reset() {
	*x = UserLesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_lessons_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLesson) ProtoMessage() {}

func (x *UserLesson) ProtoReflect() protoreflect.Message {
	mi := &file_user_lessons_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLesson.ProtoReflect.Descriptor instead.
func (*UserLesson) Descriptor() ([]byte, []int) {
	return file_user_lessons_proto_rawDescGZIP(), []int{4}
}

func (x *UserLesson) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserLesson) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLesson) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *UserLesson) GetCompletedAt() string {
	if x != nil {
		return x.CompletedAt
	}
	return ""
}

func (x *UserLesson) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserLesson) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_user_lessons_proto protoreflect.FileDescriptor

var file_user_lessons_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x6f, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x59, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0xb3, 0x01,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x32, 0xda, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0d, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x2e,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_lessons_proto_rawDescOnce sync.Once
	file_user_lessons_proto_rawDescData = file_user_lessons_proto_rawDesc
)

func file_user_lessons_proto_rawDescGZIP() []byte {
	file_user_lessons_proto_rawDescOnce.Do(func() {
		file_user_lessons_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_lessons_proto_rawDescData)
	})
	return file_user_lessons_proto_rawDescData
}

var file_user_lessons_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_lessons_proto_goTypes = []any{
	(*AddUserLessonRequest)(nil),   // 0: languagelearning.AddUserLessonRequest
	(*AddUserLessonResponse)(nil),  // 1: languagelearning.AddUserLessonResponse
	(*GetUserLessonsRequest)(nil),  // 2: languagelearning.GetUserLessonsRequest
	(*GetUserLessonsResponse)(nil), // 3: languagelearning.GetUserLessonsResponse
	(*UserLesson)(nil),             // 4: languagelearning.UserLesson
}
var file_user_lessons_proto_depIdxs = []int32{
	4, // 0: languagelearning.GetUserLessonsResponse.user_lessons:type_name -> languagelearning.UserLesson
	0, // 1: languagelearning.UserLessonService.AddUserLesson:input_type -> languagelearning.AddUserLessonRequest
	2, // 2: languagelearning.UserLessonService.GetUserLessons:input_type -> languagelearning.GetUserLessonsRequest
	1, // 3: languagelearning.UserLessonService.AddUserLesson:output_type -> languagelearning.AddUserLessonResponse
	3, // 4: languagelearning.UserLessonService.GetUserLessons:output_type -> languagelearning.GetUserLessonsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_lessons_proto_init() }
func file_user_lessons_proto_init() {
	if File_user_lessons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_lessons_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddUserLessonRequest); i {
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
		file_user_lessons_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddUserLessonResponse); i {
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
		file_user_lessons_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserLessonsRequest); i {
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
		file_user_lessons_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserLessonsResponse); i {
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
		file_user_lessons_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserLesson); i {
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
			RawDescriptor: file_user_lessons_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_lessons_proto_goTypes,
		DependencyIndexes: file_user_lessons_proto_depIdxs,
		MessageInfos:      file_user_lessons_proto_msgTypes,
	}.Build()
	File_user_lessons_proto = out.File
	file_user_lessons_proto_rawDesc = nil
	file_user_lessons_proto_goTypes = nil
	file_user_lessons_proto_depIdxs = nil
}
