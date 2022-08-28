// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: notice.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 通知读者的消息
type NoticeReaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// Types that are assignable to NoticeWay:
	//	*NoticeReaderRequest_Email
	//	*NoticeReaderRequest_Phone
	NoticeWay isNoticeReaderRequest_NoticeWay `protobuf_oneof:"notice_way"`
}

func (x *NoticeReaderRequest) Reset() {
	*x = NoticeReaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeReaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeReaderRequest) ProtoMessage() {}

func (x *NoticeReaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeReaderRequest.ProtoReflect.Descriptor instead.
func (*NoticeReaderRequest) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{0}
}

func (x *NoticeReaderRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (m *NoticeReaderRequest) GetNoticeWay() isNoticeReaderRequest_NoticeWay {
	if m != nil {
		return m.NoticeWay
	}
	return nil
}

func (x *NoticeReaderRequest) GetEmail() string {
	if x, ok := x.GetNoticeWay().(*NoticeReaderRequest_Email); ok {
		return x.Email
	}
	return ""
}

func (x *NoticeReaderRequest) GetPhone() string {
	if x, ok := x.GetNoticeWay().(*NoticeReaderRequest_Phone); ok {
		return x.Phone
	}
	return ""
}

type isNoticeReaderRequest_NoticeWay interface {
	isNoticeReaderRequest_NoticeWay()
}

type NoticeReaderRequest_Email struct {
	Email string `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

type NoticeReaderRequest_Phone struct {
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3,oneof"`
}

func (*NoticeReaderRequest_Email) isNoticeReaderRequest_NoticeWay() {}

func (*NoticeReaderRequest_Phone) isNoticeReaderRequest_NoticeWay() {}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// int64 price = 3;
	// google.protobuf.Int64Value price = 3;
	Price     *int64                  `protobuf:"varint,3,opt,name=price,proto3,oneof" json:"price,omitempty"`
	SalePrice *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"` // float64
	Memo      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`                            // string
	Info      *Book_Info              `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`                            // 嵌套
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[1]
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
	return file_notice_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPrice() int64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *Book) GetSalePrice() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SalePrice
	}
	return nil
}

func (x *Book) GetMemo() *wrapperspb.StringValue {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *Book) GetInfo() *Book_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

// UpdateBookRequest 更新书籍的消息
type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 操作人
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// 要更新的书籍信息
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	// 要更新的字段
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBookRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *UpdateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UpdateBookRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type Book_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"` // book.info.b
}

func (x *Book_Info) Reset() {
	*x = Book_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Info) ProtoMessage() {}

func (x *Book_Info) ProtoReflect() protoreflect.Message {
	mi := &file_notice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Info.ProtoReflect.Descriptor instead.
func (*Book_Info) Descriptor() ([]byte, []int) {
	return file_notice_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Book_Info) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Book_Info) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

var File_notice_proto protoreflect.FileDescriptor

var file_notice_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x77, 0x61, 0x79, 0x22, 0x90, 0x02, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3b,
	0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x22, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x1a, 0x22, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x62, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x7f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x70, 0x12, 0x1d, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x64, 0x65, 0x6d,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notice_proto_rawDescOnce sync.Once
	file_notice_proto_rawDescData = file_notice_proto_rawDesc
)

func file_notice_proto_rawDescGZIP() []byte {
	file_notice_proto_rawDescOnce.Do(func() {
		file_notice_proto_rawDescData = protoimpl.X.CompressGZIP(file_notice_proto_rawDescData)
	})
	return file_notice_proto_rawDescData
}

var file_notice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notice_proto_goTypes = []interface{}{
	(*NoticeReaderRequest)(nil),    // 0: api.NoticeReaderRequest
	(*Book)(nil),                   // 1: api.Book
	(*UpdateBookRequest)(nil),      // 2: api.UpdateBookRequest
	(*Book_Info)(nil),              // 3: api.Book.Info
	(*wrapperspb.DoubleValue)(nil), // 4: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*fieldmaskpb.FieldMask)(nil),  // 6: google.protobuf.FieldMask
}
var file_notice_proto_depIdxs = []int32{
	4, // 0: api.Book.sale_price:type_name -> google.protobuf.DoubleValue
	5, // 1: api.Book.memo:type_name -> google.protobuf.StringValue
	3, // 2: api.Book.info:type_name -> api.Book.Info
	1, // 3: api.UpdateBookRequest.book:type_name -> api.Book
	6, // 4: api.UpdateBookRequest.update_mask:type_name -> google.protobuf.FieldMask
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_notice_proto_init() }
func file_notice_proto_init() {
	if File_notice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeReaderRequest); i {
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
		file_notice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_notice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
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
		file_notice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book_Info); i {
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
	file_notice_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NoticeReaderRequest_Email)(nil),
		(*NoticeReaderRequest_Phone)(nil),
	}
	file_notice_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notice_proto_goTypes,
		DependencyIndexes: file_notice_proto_depIdxs,
		MessageInfos:      file_notice_proto_msgTypes,
	}.Build()
	File_notice_proto = out.File
	file_notice_proto_rawDesc = nil
	file_notice_proto_goTypes = nil
	file_notice_proto_depIdxs = nil
}
