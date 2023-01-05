// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: plugins/strategy/proto/v1/strategy.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/hashicorp/nomad-autoscaler/plugins/shared/proto/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action            *v1.ScalingAction       `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Check             *v1.ScalingPolicyCheck  `protobuf:"bytes,2,opt,name=check,proto3" json:"check,omitempty"`
	TimestampedMetric []*v1.TimestampedMetric `protobuf:"bytes,3,rep,name=timestamped_metric,json=timestampedMetric,proto3" json:"timestamped_metric,omitempty"`
	Count             int64                   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_strategy_proto_v1_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_strategy_proto_v1_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_plugins_strategy_proto_v1_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *RunRequest) GetAction() *v1.ScalingAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *RunRequest) GetCheck() *v1.ScalingPolicyCheck {
	if x != nil {
		return x.Check
	}
	return nil
}

func (x *RunRequest) GetTimestampedMetric() []*v1.TimestampedMetric {
	if x != nil {
		return x.TimestampedMetric
	}
	return nil
}

func (x *RunRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action            *v1.ScalingAction       `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Check             *v1.ScalingPolicyCheck  `protobuf:"bytes,2,opt,name=check,proto3" json:"check,omitempty"`
	TimestampedMetric []*v1.TimestampedMetric `protobuf:"bytes,3,rep,name=timestamped_metric,json=timestampedMetric,proto3" json:"timestamped_metric,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_strategy_proto_v1_strategy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_strategy_proto_v1_strategy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_plugins_strategy_proto_v1_strategy_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetAction() *v1.ScalingAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *RunResponse) GetCheck() *v1.ScalingPolicyCheck {
	if x != nil {
		return x.Check
	}
	return nil
}

func (x *RunResponse) GetTimestampedMetric() []*v1.TimestampedMetric {
	if x != nil {
		return x.TimestampedMetric
	}
	return nil
}

var File_plugins_strategy_proto_v1_strategy_proto protoreflect.FileDescriptor

var file_plugins_strategy_proto_v1_strategy_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x1a, 0x24, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5c, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x46, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61,
	0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x74,
	0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74,
	0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xbc, 0x02, 0x0a, 0x0b, 0x52,
	0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74,
	0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x05, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x74, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x45, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61,
	0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x32, 0xa6, 0x01, 0x0a, 0x15, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x40, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f, 0x61, 0x75,
	0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_plugins_strategy_proto_v1_strategy_proto_rawDescOnce sync.Once
	file_plugins_strategy_proto_v1_strategy_proto_rawDescData = file_plugins_strategy_proto_v1_strategy_proto_rawDesc
)

func file_plugins_strategy_proto_v1_strategy_proto_rawDescGZIP() []byte {
	file_plugins_strategy_proto_v1_strategy_proto_rawDescOnce.Do(func() {
		file_plugins_strategy_proto_v1_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_strategy_proto_v1_strategy_proto_rawDescData)
	})
	return file_plugins_strategy_proto_v1_strategy_proto_rawDescData
}

var file_plugins_strategy_proto_v1_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_plugins_strategy_proto_v1_strategy_proto_goTypes = []interface{}{
	(*RunRequest)(nil),            // 0: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunRequest
	(*RunResponse)(nil),           // 1: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunResponse
	(*v1.ScalingAction)(nil),      // 2: hashicorp.nomad_autoscaler.plugins.shared.proto.v1.ScalingAction
	(*v1.ScalingPolicyCheck)(nil), // 3: hashicorp.nomad_autoscaler.plugins.shared.proto.v1.ScalingPolicyCheck
	(*v1.TimestampedMetric)(nil),  // 4: hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimestampedMetric
}
var file_plugins_strategy_proto_v1_strategy_proto_depIdxs = []int32{
	2, // 0: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunRequest.action:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.ScalingAction
	3, // 1: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunRequest.check:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.ScalingPolicyCheck
	4, // 2: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunRequest.timestamped_metric:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimestampedMetric
	2, // 3: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunResponse.action:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.ScalingAction
	3, // 4: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunResponse.check:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.ScalingPolicyCheck
	4, // 5: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunResponse.timestamped_metric:type_name -> hashicorp.nomad_autoscaler.plugins.shared.proto.v1.TimestampedMetric
	0, // 6: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.StrategyPluginService.Run:input_type -> hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunRequest
	1, // 7: hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.StrategyPluginService.Run:output_type -> hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.RunResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_plugins_strategy_proto_v1_strategy_proto_init() }
func file_plugins_strategy_proto_v1_strategy_proto_init() {
	if File_plugins_strategy_proto_v1_strategy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugins_strategy_proto_v1_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_plugins_strategy_proto_v1_strategy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
			RawDescriptor: file_plugins_strategy_proto_v1_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugins_strategy_proto_v1_strategy_proto_goTypes,
		DependencyIndexes: file_plugins_strategy_proto_v1_strategy_proto_depIdxs,
		MessageInfos:      file_plugins_strategy_proto_v1_strategy_proto_msgTypes,
	}.Build()
	File_plugins_strategy_proto_v1_strategy_proto = out.File
	file_plugins_strategy_proto_v1_strategy_proto_rawDesc = nil
	file_plugins_strategy_proto_v1_strategy_proto_goTypes = nil
	file_plugins_strategy_proto_v1_strategy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StrategyPluginServiceClient is the client API for StrategyPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StrategyPluginServiceClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type strategyPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyPluginServiceClient(cc grpc.ClientConnInterface) StrategyPluginServiceClient {
	return &strategyPluginServiceClient{cc}
}

func (c *strategyPluginServiceClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.StrategyPluginService/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrategyPluginServiceServer is the server API for StrategyPluginService service.
type StrategyPluginServiceServer interface {
	Run(context.Context, *RunRequest) (*RunResponse, error)
}

// UnimplementedStrategyPluginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStrategyPluginServiceServer struct {
}

func (*UnimplementedStrategyPluginServiceServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterStrategyPluginServiceServer(s *grpc.Server, srv StrategyPluginServiceServer) {
	s.RegisterService(&_StrategyPluginService_serviceDesc, srv)
}

func _StrategyPluginService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyPluginServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.StrategyPluginService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyPluginServiceServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StrategyPluginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad_autoscaler.plugins.strategy.proto.v1.StrategyPluginService",
	HandlerType: (*StrategyPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _StrategyPluginService_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins/strategy/proto/v1/strategy.proto",
}
