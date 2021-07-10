// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: internal/protobuf/list.proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Types that are assignable to Body:
	//	*Request_Get
	//	*Request_Del
	//	*Request_Put
	Body isRequest_Body `protobuf_oneof:"body"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_list_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (m *Request) GetBody() isRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Request) GetGet() *Id {
	if x, ok := x.GetBody().(*Request_Get); ok {
		return x.Get
	}
	return nil
}

func (x *Request) GetDel() *Id {
	if x, ok := x.GetBody().(*Request_Del); ok {
		return x.Del
	}
	return nil
}

func (x *Request) GetPut() *Ticket {
	if x, ok := x.GetBody().(*Request_Put); ok {
		return x.Put
	}
	return nil
}

type isRequest_Body interface {
	isRequest_Body()
}

type Request_Get struct {
	// if Id == 0, server will response with all todo entities.
	Get *Id `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type Request_Del struct {
	// just delete Ticket by Id
	Del *Id `protobuf:"bytes,4,opt,name=del,proto3,oneof"`
}

type Request_Put struct {
	// PUT method will change body of existing Ticket if id == 0.
	// Otherwise it will create new Ticket.
	Put *Ticket `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}

func (*Request_Get) isRequest_Body() {}

func (*Request_Del) isRequest_Body() {}

func (*Request_Put) isRequest_Body() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*Response_Ticket
	//	*Response_Text
	Body isResponse_Body `protobuf_oneof:"body"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_list_proto_msgTypes[1]
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
	return file_internal_protobuf_list_proto_rawDescGZIP(), []int{1}
}

func (m *Response) GetBody() isResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Response) GetTicket() *Ticket {
	if x, ok := x.GetBody().(*Response_Ticket); ok {
		return x.Ticket
	}
	return nil
}

func (x *Response) GetText() string {
	if x, ok := x.GetBody().(*Response_Text); ok {
		return x.Text
	}
	return ""
}

type isResponse_Body interface {
	isResponse_Body()
}

type Response_Ticket struct {
	Ticket *Ticket `protobuf:"bytes,1,opt,name=ticket,proto3,oneof"`
}

type Response_Text struct {
	Text string `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

func (*Response_Ticket) isResponse_Body() {}

func (*Response_Text) isResponse_Body() {}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_list_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body        string               `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	State       bool                 `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Ticket) isResponse_Body() {
	panic("implement me")
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_list_proto_rawDescGZIP(), []int{3}
}

func (x *Ticket) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ticket) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Ticket) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Ticket) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

func (x *Ticket) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

var File_internal_protobuf_list_proto protoreflect.FileDescriptor

var file_internal_protobuf_list_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x03, 0x67, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x49, 0x64,
	0x48, 0x00, 0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x03, 0x64, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x49, 0x64, 0x48, 0x00,
	0x52, 0x03, 0x64, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x03, 0x70, 0x75, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0x50, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x32, 0x3a, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x12, 0x2e, 0x0a,
	0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x15, 0x5a,
	0x13, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protobuf_list_proto_rawDescOnce sync.Once
	file_internal_protobuf_list_proto_rawDescData = file_internal_protobuf_list_proto_rawDesc
)

func file_internal_protobuf_list_proto_rawDescGZIP() []byte {
	file_internal_protobuf_list_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protobuf_list_proto_rawDescData)
	})
	return file_internal_protobuf_list_proto_rawDescData
}

var file_internal_protobuf_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_protobuf_list_proto_goTypes = []interface{}{
	(*Request)(nil),             // 0: todo.Request
	(*Response)(nil),            // 1: todo.Response
	(*Id)(nil),                  // 2: todo.Id
	(*Ticket)(nil),              // 3: todo.Ticket
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_internal_protobuf_list_proto_depIdxs = []int32{
	2, // 0: todo.Request.get:type_name -> todo.Id
	2, // 1: todo.Request.del:type_name -> todo.Id
	3, // 2: todo.Request.put:type_name -> todo.Ticket
	3, // 3: todo.Response.ticket:type_name -> todo.Ticket
	4, // 4: todo.Ticket.last_updated:type_name -> google.protobuf.Timestamp
	0, // 5: todo.TodoServ.TodoRequest:input_type -> todo.Request
	1, // 6: todo.TodoServ.TodoRequest:output_type -> todo.Response
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_protobuf_list_proto_init() }
func file_internal_protobuf_list_proto_init() {
	if File_internal_protobuf_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_internal_protobuf_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_protobuf_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_internal_protobuf_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
	file_internal_protobuf_list_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_Get)(nil),
		(*Request_Del)(nil),
		(*Request_Put)(nil),
	}
	file_internal_protobuf_list_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_Ticket)(nil),
		(*Response_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_protobuf_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_list_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_list_proto_depIdxs,
		MessageInfos:      file_internal_protobuf_list_proto_msgTypes,
	}.Build()
	File_internal_protobuf_list_proto = out.File
	file_internal_protobuf_list_proto_rawDesc = nil
	file_internal_protobuf_list_proto_goTypes = nil
	file_internal_protobuf_list_proto_depIdxs = nil
}
