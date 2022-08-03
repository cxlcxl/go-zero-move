// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb/marketing.proto

package marketing

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

type PromotionCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CampaignName  string `protobuf:"bytes,1,opt,name=campaign_name,json=campaignName,proto3" json:"campaign_name,omitempty"`
	ProductType   string `protobuf:"bytes,2,opt,name=product_type,json=productType,proto3" json:"product_type,omitempty"`
	DailyBudget   int64  `protobuf:"varint,3,opt,name=daily_budget,json=dailyBudget,proto3" json:"daily_budget,omitempty"`
	AdvertiserId  string `protobuf:"bytes,4,opt,name=advertiser_id,json=advertiserId,proto3" json:"advertiser_id,omitempty"`
	MarketingGoal string `protobuf:"bytes,5,opt,name=marketing_goal,json=marketingGoal,proto3" json:"marketing_goal,omitempty"`
}

func (x *PromotionCreateReq) Reset() {
	*x = PromotionCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_marketing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionCreateReq) ProtoMessage() {}

func (x *PromotionCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_marketing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionCreateReq.ProtoReflect.Descriptor instead.
func (*PromotionCreateReq) Descriptor() ([]byte, []int) {
	return file_pb_marketing_proto_rawDescGZIP(), []int{0}
}

func (x *PromotionCreateReq) GetCampaignName() string {
	if x != nil {
		return x.CampaignName
	}
	return ""
}

func (x *PromotionCreateReq) GetProductType() string {
	if x != nil {
		return x.ProductType
	}
	return ""
}

func (x *PromotionCreateReq) GetDailyBudget() int64 {
	if x != nil {
		return x.DailyBudget
	}
	return 0
}

func (x *PromotionCreateReq) GetAdvertiserId() string {
	if x != nil {
		return x.AdvertiserId
	}
	return ""
}

func (x *PromotionCreateReq) GetMarketingGoal() string {
	if x != nil {
		return x.MarketingGoal
	}
	return ""
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_marketing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_marketing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_pb_marketing_proto_rawDescGZIP(), []int{1}
}

func (x *BaseResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_pb_marketing_proto protoreflect.FileDescriptor

var file_pb_marketing_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x1a,
	0x10, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x61, 0x6c, 0x22,
	0x1e, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0xa8, 0x04, 0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4c,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x44, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x13,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_marketing_proto_rawDescOnce sync.Once
	file_pb_marketing_proto_rawDescData = file_pb_marketing_proto_rawDesc
)

func file_pb_marketing_proto_rawDescGZIP() []byte {
	file_pb_marketing_proto_rawDescOnce.Do(func() {
		file_pb_marketing_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_marketing_proto_rawDescData)
	})
	return file_pb_marketing_proto_rawDescData
}

var file_pb_marketing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_marketing_proto_goTypes = []interface{}{
	(*PromotionCreateReq)(nil), // 0: marketing.PromotionCreateReq
	(*BaseResp)(nil),           // 1: marketing.BaseResp
	(*AccountCreateReq)(nil),   // 2: marketing.AccountCreateReq
	(*AccountUpdateReq)(nil),   // 3: marketing.AccountUpdateReq
	(*AccountInfoReq)(nil),     // 4: marketing.AccountInfoReq
	(*GetTokenReq)(nil),        // 5: marketing.GetTokenReq
	(*AccountListReq)(nil),     // 6: marketing.AccountListReq
	(*TokenInfo)(nil),          // 7: marketing.TokenInfo
	(*AccountInfo)(nil),        // 8: marketing.AccountInfo
	(*AccountListResp)(nil),    // 9: marketing.AccountListResp
}
var file_pb_marketing_proto_depIdxs = []int32{
	2, // 0: marketing.MarketingCenter.AccountCreate:input_type -> marketing.AccountCreateReq
	3, // 1: marketing.MarketingCenter.AccountUpdate:input_type -> marketing.AccountUpdateReq
	4, // 2: marketing.MarketingCenter.GetAccountInfo:input_type -> marketing.AccountInfoReq
	5, // 3: marketing.MarketingCenter.GetAccountSecretByClientId:input_type -> marketing.GetTokenReq
	6, // 4: marketing.MarketingCenter.AccountList:input_type -> marketing.AccountListReq
	5, // 5: marketing.MarketingCenter.GetToken:input_type -> marketing.GetTokenReq
	7, // 6: marketing.MarketingCenter.SetToken:input_type -> marketing.TokenInfo
	0, // 7: marketing.MarketingCenter.PromotionCreate:input_type -> marketing.PromotionCreateReq
	1, // 8: marketing.MarketingCenter.AccountCreate:output_type -> marketing.BaseResp
	1, // 9: marketing.MarketingCenter.AccountUpdate:output_type -> marketing.BaseResp
	8, // 10: marketing.MarketingCenter.GetAccountInfo:output_type -> marketing.AccountInfo
	8, // 11: marketing.MarketingCenter.GetAccountSecretByClientId:output_type -> marketing.AccountInfo
	9, // 12: marketing.MarketingCenter.AccountList:output_type -> marketing.AccountListResp
	7, // 13: marketing.MarketingCenter.GetToken:output_type -> marketing.TokenInfo
	1, // 14: marketing.MarketingCenter.SetToken:output_type -> marketing.BaseResp
	1, // 15: marketing.MarketingCenter.PromotionCreate:output_type -> marketing.BaseResp
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_marketing_proto_init() }
func file_pb_marketing_proto_init() {
	if File_pb_marketing_proto != nil {
		return
	}
	file_pb_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_marketing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionCreateReq); i {
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
		file_pb_marketing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
			RawDescriptor: file_pb_marketing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_marketing_proto_goTypes,
		DependencyIndexes: file_pb_marketing_proto_depIdxs,
		MessageInfos:      file_pb_marketing_proto_msgTypes,
	}.Build()
	File_pb_marketing_proto = out.File
	file_pb_marketing_proto_rawDesc = nil
	file_pb_marketing_proto_goTypes = nil
	file_pb_marketing_proto_depIdxs = nil
}
