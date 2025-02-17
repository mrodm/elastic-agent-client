// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// This file preserves the messages used with the deprecated elastic agent Checkin() RPC for reference.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: elastic-agent-client-deprecated.proto

package proto

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

// Status codes for the current state.
type StateObserved_Status int32

const (
	// Application is starting.
	StateObserved_STARTING StateObserved_Status = 0
	// Application is currently configuring.
	StateObserved_CONFIGURING StateObserved_Status = 1
	// Application is in healthy state.
	StateObserved_HEALTHY StateObserved_Status = 2
	// Application is working but in a degraded state.
	StateObserved_DEGRADED StateObserved_Status = 3
	// Application is failing completely.
	StateObserved_FAILED StateObserved_Status = 4
	// Application is stopping.
	StateObserved_STOPPING StateObserved_Status = 5
)

// Enum value maps for StateObserved_Status.
var (
	StateObserved_Status_name = map[int32]string{
		0: "STARTING",
		1: "CONFIGURING",
		2: "HEALTHY",
		3: "DEGRADED",
		4: "FAILED",
		5: "STOPPING",
	}
	StateObserved_Status_value = map[string]int32{
		"STARTING":    0,
		"CONFIGURING": 1,
		"HEALTHY":     2,
		"DEGRADED":    3,
		"FAILED":      4,
		"STOPPING":    5,
	}
)

func (x StateObserved_Status) Enum() *StateObserved_Status {
	p := new(StateObserved_Status)
	*p = x
	return p
}

func (x StateObserved_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateObserved_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_elastic_agent_client_deprecated_proto_enumTypes[0].Descriptor()
}

func (StateObserved_Status) Type() protoreflect.EnumType {
	return &file_elastic_agent_client_deprecated_proto_enumTypes[0]
}

func (x StateObserved_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateObserved_Status.Descriptor instead.
func (StateObserved_Status) EnumDescriptor() ([]byte, []int) {
	return file_elastic_agent_client_deprecated_proto_rawDescGZIP(), []int{0, 0}
}

type StateExpected_State int32

const (
	// Expects that the application is running.
	StateExpected_RUNNING StateExpected_State = 0
	// Expects that the application is stopping.
	StateExpected_STOPPING StateExpected_State = 1
)

// Enum value maps for StateExpected_State.
var (
	StateExpected_State_name = map[int32]string{
		0: "RUNNING",
		1: "STOPPING",
	}
	StateExpected_State_value = map[string]int32{
		"RUNNING":  0,
		"STOPPING": 1,
	}
)

func (x StateExpected_State) Enum() *StateExpected_State {
	p := new(StateExpected_State)
	*p = x
	return p
}

func (x StateExpected_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateExpected_State) Descriptor() protoreflect.EnumDescriptor {
	return file_elastic_agent_client_deprecated_proto_enumTypes[1].Descriptor()
}

func (StateExpected_State) Type() protoreflect.EnumType {
	return &file_elastic_agent_client_deprecated_proto_enumTypes[1]
}

func (x StateExpected_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateExpected_State.Descriptor instead.
func (StateExpected_State) EnumDescriptor() ([]byte, []int) {
	return file_elastic_agent_client_deprecated_proto_rawDescGZIP(), []int{1, 0}
}

// A status observed message is streamed from the application to Elastic Agent.
//
// This message contains the currently applied `config_state_idx` (0 in the case of initial start, 1 is the first
// applied config index) along with the status of the application. In the case that the sent `config_state_idx`
// doesn't match the expected `config_state_idx` that Elastic Agent expects, the application is always marked as
// `CONFIGURING`.
type StateObserved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token that is used to uniquely identify the application to agent. When agent started this
	// application it would have provided it this token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Current index of the applied configuration.
	ConfigStateIdx uint64 `protobuf:"varint,2,opt,name=config_state_idx,json=configStateIdx,proto3" json:"config_state_idx,omitempty"`
	// Status code.
	Status StateObserved_Status `protobuf:"varint,3,opt,name=status,proto3,enum=proto.StateObserved_Status" json:"status,omitempty"`
	// Message for the health status.
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// JSON encoded payload for the status.
	Payload string `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StateObserved) Reset() {
	*x = StateObserved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elastic_agent_client_deprecated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateObserved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateObserved) ProtoMessage() {}

func (x *StateObserved) ProtoReflect() protoreflect.Message {
	mi := &file_elastic_agent_client_deprecated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateObserved.ProtoReflect.Descriptor instead.
func (*StateObserved) Descriptor() ([]byte, []int) {
	return file_elastic_agent_client_deprecated_proto_rawDescGZIP(), []int{0}
}

func (x *StateObserved) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StateObserved) GetConfigStateIdx() uint64 {
	if x != nil {
		return x.ConfigStateIdx
	}
	return 0
}

func (x *StateObserved) GetStatus() StateObserved_Status {
	if x != nil {
		return x.Status
	}
	return StateObserved_STARTING
}

func (x *StateObserved) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StateObserved) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

// A state expected message is streamed from the Elastic Agent to the application informing the application
// what Elastic Agent expects the applications state to be.
type StateExpected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expected state of the application.
	State StateExpected_State `protobuf:"varint,1,opt,name=state,proto3,enum=proto.StateExpected_State" json:"state,omitempty"`
	// Index of the either current configuration or new configuration provided.
	ConfigStateIdx uint64 `protobuf:"varint,2,opt,name=config_state_idx,json=configStateIdx,proto3" json:"config_state_idx,omitempty"`
	// Resulting configuration. (If the application already has the current `config_state_idx` this
	// will be empty.)
	Config string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *StateExpected) Reset() {
	*x = StateExpected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elastic_agent_client_deprecated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateExpected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateExpected) ProtoMessage() {}

func (x *StateExpected) ProtoReflect() protoreflect.Message {
	mi := &file_elastic_agent_client_deprecated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateExpected.ProtoReflect.Descriptor instead.
func (*StateExpected) Descriptor() ([]byte, []int) {
	return file_elastic_agent_client_deprecated_proto_rawDescGZIP(), []int{1}
}

func (x *StateExpected) GetState() StateExpected_State {
	if x != nil {
		return x.State
	}
	return StateExpected_RUNNING
}

func (x *StateExpected) GetConfigStateIdx() uint64 {
	if x != nil {
		return x.ConfigStateIdx
	}
	return 0
}

func (x *StateExpected) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

var File_elastic_agent_client_deprecated_proto protoreflect.FileDescriptor

var file_elastic_agent_client_deprecated_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96,
	0x02, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x64, 0x78,
	0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x45, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f,
	0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x22, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x42, 0x14, 0x5a, 0x0f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_elastic_agent_client_deprecated_proto_rawDescOnce sync.Once
	file_elastic_agent_client_deprecated_proto_rawDescData = file_elastic_agent_client_deprecated_proto_rawDesc
)

func file_elastic_agent_client_deprecated_proto_rawDescGZIP() []byte {
	file_elastic_agent_client_deprecated_proto_rawDescOnce.Do(func() {
		file_elastic_agent_client_deprecated_proto_rawDescData = protoimpl.X.CompressGZIP(file_elastic_agent_client_deprecated_proto_rawDescData)
	})
	return file_elastic_agent_client_deprecated_proto_rawDescData
}

var file_elastic_agent_client_deprecated_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_elastic_agent_client_deprecated_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_elastic_agent_client_deprecated_proto_goTypes = []interface{}{
	(StateObserved_Status)(0), // 0: proto.StateObserved.Status
	(StateExpected_State)(0),  // 1: proto.StateExpected.State
	(*StateObserved)(nil),     // 2: proto.StateObserved
	(*StateExpected)(nil),     // 3: proto.StateExpected
}
var file_elastic_agent_client_deprecated_proto_depIdxs = []int32{
	0, // 0: proto.StateObserved.status:type_name -> proto.StateObserved.Status
	1, // 1: proto.StateExpected.state:type_name -> proto.StateExpected.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_elastic_agent_client_deprecated_proto_init() }
func file_elastic_agent_client_deprecated_proto_init() {
	if File_elastic_agent_client_deprecated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_elastic_agent_client_deprecated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateObserved); i {
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
		file_elastic_agent_client_deprecated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateExpected); i {
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
			RawDescriptor: file_elastic_agent_client_deprecated_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_elastic_agent_client_deprecated_proto_goTypes,
		DependencyIndexes: file_elastic_agent_client_deprecated_proto_depIdxs,
		EnumInfos:         file_elastic_agent_client_deprecated_proto_enumTypes,
		MessageInfos:      file_elastic_agent_client_deprecated_proto_msgTypes,
	}.Build()
	File_elastic_agent_client_deprecated_proto = out.File
	file_elastic_agent_client_deprecated_proto_rawDesc = nil
	file_elastic_agent_client_deprecated_proto_goTypes = nil
	file_elastic_agent_client_deprecated_proto_depIdxs = nil
}
