// ./proto/orders/order.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.0
// source: orders/order.proto

package orders

import (
	product "github.com/microservice-veggie/grpc-service/protogen/golang/product"
	date "google.golang.org/genproto/googleapis/type/date"
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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    uint64             `protobuf:"varint,1,opt,name=order_id,proto3" json:"order_id,omitempty"`
	CustomerId uint64             `protobuf:"varint,2,opt,name=customer_id,proto3" json:"customer_id,omitempty"`
	IsActive   bool               `protobuf:"varint,3,opt,name=is_active,proto3" json:"is_active,omitempty"`
	Products   []*product.Product `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	OrderDate  *date.Date         `protobuf:"bytes,5,opt,name=order_date,proto3" json:"order_date,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Order) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Order) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Order) GetProducts() []*product.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Order) GetOrderDate() *date.Date {
	if x != nil {
		return x.OrderDate
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{1}
}

type PayloadWithSingleOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *PayloadWithSingleOrder) Reset() {
	*x = PayloadWithSingleOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orders_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadWithSingleOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadWithSingleOrder) ProtoMessage() {}

func (x *PayloadWithSingleOrder) ProtoReflect() protoreflect.Message {
	mi := &file_orders_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadWithSingleOrder.ProtoReflect.Descriptor instead.
func (*PayloadWithSingleOrder) Descriptor() ([]byte, []int) {
	return file_orders_order_proto_rawDescGZIP(), []int{2}
}

func (x *PayloadWithSingleOrder) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_orders_order_proto protoreflect.FileDescriptor

var file_orders_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x31,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x36, 0x0a, 0x16, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x32, 0x37, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x57, 0x69, 0x74, 0x68, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x76, 0x65, 0x67, 0x67, 0x69, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_order_proto_rawDescOnce sync.Once
	file_orders_order_proto_rawDescData = file_orders_order_proto_rawDesc
)

func file_orders_order_proto_rawDescGZIP() []byte {
	file_orders_order_proto_rawDescOnce.Do(func() {
		file_orders_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_order_proto_rawDescData)
	})
	return file_orders_order_proto_rawDescData
}

var file_orders_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orders_order_proto_goTypes = []interface{}{
	(*Order)(nil),                  // 0: Order
	(*Empty)(nil),                  // 1: Empty
	(*PayloadWithSingleOrder)(nil), // 2: PayloadWithSingleOrder
	(*product.Product)(nil),        // 3: Product
	(*date.Date)(nil),              // 4: google.type.Date
}
var file_orders_order_proto_depIdxs = []int32{
	3, // 0: Order.products:type_name -> Product
	4, // 1: Order.order_date:type_name -> google.type.Date
	0, // 2: PayloadWithSingleOrder.order:type_name -> Order
	2, // 3: Orders.AddOrder:input_type -> PayloadWithSingleOrder
	1, // 4: Orders.AddOrder:output_type -> Empty
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_orders_order_proto_init() }
func file_orders_order_proto_init() {
	if File_orders_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orders_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_orders_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_orders_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadWithSingleOrder); i {
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
			RawDescriptor: file_orders_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orders_order_proto_goTypes,
		DependencyIndexes: file_orders_order_proto_depIdxs,
		MessageInfos:      file_orders_order_proto_msgTypes,
	}.Build()
	File_orders_order_proto = out.File
	file_orders_order_proto_rawDesc = nil
	file_orders_order_proto_goTypes = nil
	file_orders_order_proto_depIdxs = nil
}
