// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: book/price.proto

package book

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

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketPrice int64 `protobuf:"varint,1,opt,name=market_price,json=marketPrice,proto3" json:"market_price,omitempty"`
	SalePrice   int64 `protobuf:"varint,2,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_book_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_book_price_proto_rawDescGZIP(), []int{0}
}

func (x *Price) GetMarketPrice() int64 {
	if x != nil {
		return x.MarketPrice
	}
	return 0
}

func (x *Price) GetSalePrice() int64 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

var File_book_price_proto protoreflect.FileDescriptor

var file_book_price_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x49, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_price_proto_rawDescOnce sync.Once
	file_book_price_proto_rawDescData = file_book_price_proto_rawDesc
)

func file_book_price_proto_rawDescGZIP() []byte {
	file_book_price_proto_rawDescOnce.Do(func() {
		file_book_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_price_proto_rawDescData)
	})
	return file_book_price_proto_rawDescData
}

var file_book_price_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_book_price_proto_goTypes = []interface{}{
	(*Price)(nil), // 0: book.Price
}
var file_book_price_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_book_price_proto_init() }
func file_book_price_proto_init() {
	if File_book_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
			RawDescriptor: file_book_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_book_price_proto_goTypes,
		DependencyIndexes: file_book_price_proto_depIdxs,
		MessageInfos:      file_book_price_proto_msgTypes,
	}.Build()
	File_book_price_proto = out.File
	file_book_price_proto_rawDesc = nil
	file_book_price_proto_goTypes = nil
	file_book_price_proto_depIdxs = nil
}
