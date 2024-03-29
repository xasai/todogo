// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: internal/protobuf/note.proto

package protobuf

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type RequestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId int32 `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (x *RequestId) Reset() {
	*x = RequestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestId) ProtoMessage() {}

func (x *RequestId) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestId.ProtoReflect.Descriptor instead.
func (*RequestId) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_note_proto_rawDescGZIP(), []int{0}
}

func (x *RequestId) GetNoteId() int32 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Note *Note  `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_note_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_note_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_note_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Response) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      bool                 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"` // may be TODO or DONE
	WasCreated  *timestamp.Timestamp `protobuf:"bytes,5,opt,name=was_created,json=wasCreated,proto3" json:"was_created,omitempty"`
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_note_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_note_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_note_proto_rawDescGZIP(), []int{2}
}

func (x *Note) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Note) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Note) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Note) GetWasCreated() *timestamp.Timestamp {
	if x != nil {
		return x.WasCreated
	}
	return nil
}

func (x *Note) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

var File_internal_protobuf_note_proto protoreflect.FileDescriptor

var file_internal_protobuf_note_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6e,
	0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x04,
	0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x77, 0x61, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x77, 0x61, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x32, 0xaa, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x0e, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0a, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c,
	0x4d, 0x61, 0x72, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a,
	0x13, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protobuf_note_proto_rawDescOnce sync.Once
	file_internal_protobuf_note_proto_rawDescData = file_internal_protobuf_note_proto_rawDesc
)

func file_internal_protobuf_note_proto_rawDescGZIP() []byte {
	file_internal_protobuf_note_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protobuf_note_proto_rawDescData)
	})
	return file_internal_protobuf_note_proto_rawDescData
}

var file_internal_protobuf_note_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_protobuf_note_proto_goTypes = []interface{}{
	(*RequestId)(nil),           // 0: todo.RequestId
	(*Response)(nil),            // 1: todo.Response
	(*Note)(nil),                // 2: todo.Note
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_internal_protobuf_note_proto_depIdxs = []int32{
	2, // 0: todo.Response.note:type_name -> todo.Note
	3, // 1: todo.Note.was_created:type_name -> google.protobuf.Timestamp
	3, // 2: todo.Note.last_updated:type_name -> google.protobuf.Timestamp
	2, // 3: todo.TodoService.CreateNewNote:input_type -> todo.Note
	0, // 4: todo.TodoService.GetNoteById:input_type -> todo.RequestId
	0, // 5: todo.TodoService.GetAllNotes:input_type -> todo.RequestId
	2, // 6: todo.TodoService.ChangeNoteContent:input_type -> todo.Note
	0, // 7: todo.TodoService.MarkDoneById:input_type -> todo.RequestId
	0, // 8: todo.TodoService.DelNoteById:input_type -> todo.RequestId
	1, // 9: todo.TodoService.CreateNewNote:output_type -> todo.Response
	1, // 10: todo.TodoService.GetNoteById:output_type -> todo.Response
	2, // 11: todo.TodoService.GetAllNotes:output_type -> todo.Note
	1, // 12: todo.TodoService.ChangeNoteContent:output_type -> todo.Response
	1, // 13: todo.TodoService.MarkDoneById:output_type -> todo.Response
	1, // 14: todo.TodoService.DelNoteById:output_type -> todo.Response
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_protobuf_note_proto_init() }
func file_internal_protobuf_note_proto_init() {
	if File_internal_protobuf_note_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestId); i {
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
		file_internal_protobuf_note_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_internal_protobuf_note_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
			RawDescriptor: file_internal_protobuf_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_note_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_note_proto_depIdxs,
		MessageInfos:      file_internal_protobuf_note_proto_msgTypes,
	}.Build()
	File_internal_protobuf_note_proto = out.File
	file_internal_protobuf_note_proto_rawDesc = nil
	file_internal_protobuf_note_proto_goTypes = nil
	file_internal_protobuf_note_proto_depIdxs = nil
}
