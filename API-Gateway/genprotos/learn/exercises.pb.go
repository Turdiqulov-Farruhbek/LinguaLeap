// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: exercises.proto

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

// Savol haqida ma'lumotni aniqlash uchun xabar
type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                            // Savolning UUID si
	Type          string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`                                        // Savol turi
	LanguageCode  string   `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`    // Til kodi, masalan, 'en', 'uz'
	Difficulty    string   `protobuf:"bytes,4,opt,name=difficulty,proto3" json:"difficulty,omitempty"`                            // Qiyinlik darajasi
	Question      string   `protobuf:"bytes,5,opt,name=question,proto3" json:"question,omitempty"`                                // Savol matni
	Options       []string `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`                                  // Javob variantlari
	CorrectAnswer string   `protobuf:"bytes,7,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"` // To'g'ri javob
	Explanation   string   `protobuf:"bytes,8,opt,name=explanation,proto3" json:"explanation,omitempty"`                          // Javob izohi
	CreatedAt     string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`             // Yaratilgan vaqti
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{0}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Question) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *Question) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Question) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Question) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

func (x *Question) GetExplanation() string {
	if x != nil {
		return x.Explanation
	}
	return ""
}

func (x *Question) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// Savol yaratish uchun xabar
type CreateQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`                                        // Savol turi
	LanguageCode  string   `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`    // Til kodi, masalan, 'en', 'uz'
	Difficulty    string   `protobuf:"bytes,3,opt,name=difficulty,proto3" json:"difficulty,omitempty"`                            // Qiyinlik darajasi
	Question      string   `protobuf:"bytes,4,opt,name=question,proto3" json:"question,omitempty"`                                // Savol matni
	Options       []string `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty"`                                  // Javob variantlari
	CorrectAnswer string   `protobuf:"bytes,6,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"` // To'g'ri javob
	Explanation   string   `protobuf:"bytes,7,opt,name=explanation,proto3" json:"explanation,omitempty"`                          // Javob izohi
}

func (x *CreateQuestionRequest) Reset() {
	*x = CreateQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionRequest) ProtoMessage() {}

func (x *CreateQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionRequest.ProtoReflect.Descriptor instead.
func (*CreateQuestionRequest) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuestionRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateQuestionRequest) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *CreateQuestionRequest) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *CreateQuestionRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CreateQuestionRequest) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CreateQuestionRequest) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

func (x *CreateQuestionRequest) GetExplanation() string {
	if x != nil {
		return x.Explanation
	}
	return ""
}

// Savol yaratish javobi
type CreateQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Yaratilgan savolning UUID si
}

func (x *CreateQuestionResponse) Reset() {
	*x = CreateQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionResponse) ProtoMessage() {}

func (x *CreateQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionResponse.ProtoReflect.Descriptor instead.
func (*CreateQuestionResponse) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQuestionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Savol haqida ma'lumot so'rash uchun xabar
type GetQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Savolning UUID si
}

func (x *GetQuestionRequest) Reset() {
	*x = GetQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionRequest) ProtoMessage() {}

func (x *GetQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionRequest) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{3}
}

func (x *GetQuestionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Savol haqida ma'lumot javobi
type GetQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"` // Savol haqida ma'lumot
}

func (x *GetQuestionResponse) Reset() {
	*x = GetQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionResponse) ProtoMessage() {}

func (x *GetQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionResponse.ProtoReflect.Descriptor instead.
func (*GetQuestionResponse) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{4}
}

func (x *GetQuestionResponse) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

// Barcha savollar ro'yxatini so'rash uchun xabar
type ListQuestionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQuestionsRequest) Reset() {
	*x = ListQuestionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestionsRequest) ProtoMessage() {}

func (x *ListQuestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestionsRequest.ProtoReflect.Descriptor instead.
func (*ListQuestionsRequest) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{5}
}

// Barcha savollar ro'yxatini javobi
type ListQuestionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"` // Savollarning ro'yxati
}

func (x *ListQuestionsResponse) Reset() {
	*x = ListQuestionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestionsResponse) ProtoMessage() {}

func (x *ListQuestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestionsResponse.ProtoReflect.Descriptor instead.
func (*ListQuestionsResponse) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{6}
}

func (x *ListQuestionsResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

// Savolni yangilash uchun xabar
type UpdateQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                            // Savolning UUID si
	Type          string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`                                        // Savol turi
	LanguageCode  string   `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`    // Til kodi, masalan, 'en', 'uz'
	Difficulty    string   `protobuf:"bytes,4,opt,name=difficulty,proto3" json:"difficulty,omitempty"`                            // Qiyinlik darajasi
	Question      string   `protobuf:"bytes,5,opt,name=question,proto3" json:"question,omitempty"`                                // Savol matni
	Options       []string `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`                                  // Javob variantlari
	CorrectAnswer string   `protobuf:"bytes,7,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"` // To'g'ri javob
	Explanation   string   `protobuf:"bytes,8,opt,name=explanation,proto3" json:"explanation,omitempty"`                          // Javob izohi
}

func (x *UpdateQuestionRequest) Reset() {
	*x = UpdateQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionRequest) ProtoMessage() {}

func (x *UpdateQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuestionRequest) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateQuestionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateQuestionRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateQuestionRequest) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *UpdateQuestionRequest) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *UpdateQuestionRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *UpdateQuestionRequest) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *UpdateQuestionRequest) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

func (x *UpdateQuestionRequest) GetExplanation() string {
	if x != nil {
		return x.Explanation
	}
	return ""
}

// Savolni yangilash javobi
type UpdateQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"` // Yangilangan savol haqida ma'lumot
}

func (x *UpdateQuestionResponse) Reset() {
	*x = UpdateQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionResponse) ProtoMessage() {}

func (x *UpdateQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionResponse.ProtoReflect.Descriptor instead.
func (*UpdateQuestionResponse) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateQuestionResponse) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

// Savolni o'chirish uchun xabar
type DeleteQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Savolning UUID si
}

func (x *DeleteQuestionRequest) Reset() {
	*x = DeleteQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestionRequest) ProtoMessage() {}

func (x *DeleteQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestionRequest.ProtoReflect.Descriptor instead.
func (*DeleteQuestionRequest) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteQuestionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Savolni o'chirish javobi
type DeleteQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // O'chirilgan savolning UUID si
}

func (x *DeleteQuestionResponse) Reset() {
	*x = DeleteQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercises_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestionResponse) ProtoMessage() {}

func (x *DeleteQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercises_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestionResponse.ProtoReflect.Descriptor instead.
func (*DeleteQuestionResponse) Descriptor() ([]byte, []int) {
	return file_exercises_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteQuestionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_exercises_proto protoreflect.FileDescriptor

var file_exercises_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x22, 0x91, 0x02, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66,
	0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70,
	0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xef, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78,
	0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x51, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0xff, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xfe, 0x03, 0x0a, 0x0f, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x26, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x63, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exercises_proto_rawDescOnce sync.Once
	file_exercises_proto_rawDescData = file_exercises_proto_rawDesc
)

func file_exercises_proto_rawDescGZIP() []byte {
	file_exercises_proto_rawDescOnce.Do(func() {
		file_exercises_proto_rawDescData = protoimpl.X.CompressGZIP(file_exercises_proto_rawDescData)
	})
	return file_exercises_proto_rawDescData
}

var file_exercises_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_exercises_proto_goTypes = []any{
	(*Question)(nil),               // 0: languagelearning.Question
	(*CreateQuestionRequest)(nil),  // 1: languagelearning.CreateQuestionRequest
	(*CreateQuestionResponse)(nil), // 2: languagelearning.CreateQuestionResponse
	(*GetQuestionRequest)(nil),     // 3: languagelearning.GetQuestionRequest
	(*GetQuestionResponse)(nil),    // 4: languagelearning.GetQuestionResponse
	(*ListQuestionsRequest)(nil),   // 5: languagelearning.ListQuestionsRequest
	(*ListQuestionsResponse)(nil),  // 6: languagelearning.ListQuestionsResponse
	(*UpdateQuestionRequest)(nil),  // 7: languagelearning.UpdateQuestionRequest
	(*UpdateQuestionResponse)(nil), // 8: languagelearning.UpdateQuestionResponse
	(*DeleteQuestionRequest)(nil),  // 9: languagelearning.DeleteQuestionRequest
	(*DeleteQuestionResponse)(nil), // 10: languagelearning.DeleteQuestionResponse
}
var file_exercises_proto_depIdxs = []int32{
	0,  // 0: languagelearning.GetQuestionResponse.question:type_name -> languagelearning.Question
	0,  // 1: languagelearning.ListQuestionsResponse.questions:type_name -> languagelearning.Question
	0,  // 2: languagelearning.UpdateQuestionResponse.question:type_name -> languagelearning.Question
	1,  // 3: languagelearning.QuestionService.CreateQuestion:input_type -> languagelearning.CreateQuestionRequest
	3,  // 4: languagelearning.QuestionService.GetQuestion:input_type -> languagelearning.GetQuestionRequest
	5,  // 5: languagelearning.QuestionService.ListQuestions:input_type -> languagelearning.ListQuestionsRequest
	7,  // 6: languagelearning.QuestionService.UpdateQuestion:input_type -> languagelearning.UpdateQuestionRequest
	9,  // 7: languagelearning.QuestionService.DeleteQuestion:input_type -> languagelearning.DeleteQuestionRequest
	2,  // 8: languagelearning.QuestionService.CreateQuestion:output_type -> languagelearning.CreateQuestionResponse
	4,  // 9: languagelearning.QuestionService.GetQuestion:output_type -> languagelearning.GetQuestionResponse
	6,  // 10: languagelearning.QuestionService.ListQuestions:output_type -> languagelearning.ListQuestionsResponse
	8,  // 11: languagelearning.QuestionService.UpdateQuestion:output_type -> languagelearning.UpdateQuestionResponse
	10, // 12: languagelearning.QuestionService.DeleteQuestion:output_type -> languagelearning.DeleteQuestionResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_exercises_proto_init() }
func file_exercises_proto_init() {
	if File_exercises_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exercises_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Question); i {
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
		file_exercises_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateQuestionRequest); i {
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
		file_exercises_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateQuestionResponse); i {
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
		file_exercises_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetQuestionRequest); i {
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
		file_exercises_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetQuestionResponse); i {
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
		file_exercises_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListQuestionsRequest); i {
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
		file_exercises_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListQuestionsResponse); i {
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
		file_exercises_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateQuestionRequest); i {
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
		file_exercises_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateQuestionResponse); i {
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
		file_exercises_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteQuestionRequest); i {
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
		file_exercises_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteQuestionResponse); i {
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
			RawDescriptor: file_exercises_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exercises_proto_goTypes,
		DependencyIndexes: file_exercises_proto_depIdxs,
		MessageInfos:      file_exercises_proto_msgTypes,
	}.Build()
	File_exercises_proto = out.File
	file_exercises_proto_rawDesc = nil
	file_exercises_proto_goTypes = nil
	file_exercises_proto_depIdxs = nil
}
