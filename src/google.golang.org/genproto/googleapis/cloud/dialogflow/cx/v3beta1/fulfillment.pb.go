// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/dialogflow/cx/v3beta1/fulfillment.proto

package cx

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// A fulfillment can do one or more of the following actions at the same time:
//
//   * Generate rich message responses.
//   * Set parameter values.
//   * Call the webhook.
//
// Fulfillments can be called at various stages in the [Page][google.cloud.dialogflow.cx.v3beta1.Page] or
// [Form][google.cloud.dialogflow.cx.v3beta1.Form] lifecycle. For example, when a [DetectIntentRequest][google.cloud.dialogflow.cx.v3beta1.DetectIntentRequest] drives a
// session to enter a new page, the page's entry fulfillment can add a static
// response to the [QueryResult][google.cloud.dialogflow.cx.v3beta1.QueryResult] in the returning [DetectIntentResponse][google.cloud.dialogflow.cx.v3beta1.DetectIntentResponse],
// call the webhook (for example, to load user data from a database), or both.
type Fulfillment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rich message responses to present to the user.
	Messages []*ResponseMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	// The webhook to call.
	// Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/webhooks/<Webhook ID>`.
	Webhook string `protobuf:"bytes,2,opt,name=webhook,proto3" json:"webhook,omitempty"`
	// The tag used by the webhook to identify which fulfillment is being called.
	// This field is required if `webhook` is specified.
	Tag string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	// Set parameter values before executing the webhook.
	SetParameterActions []*Fulfillment_SetParameterAction `protobuf:"bytes,4,rep,name=set_parameter_actions,json=setParameterActions,proto3" json:"set_parameter_actions,omitempty"`
	// Conditional cases for this fulfillment.
	ConditionalCases []*Fulfillment_ConditionalCases `protobuf:"bytes,5,rep,name=conditional_cases,json=conditionalCases,proto3" json:"conditional_cases,omitempty"`
}

func (x *Fulfillment) Reset() {
	*x = Fulfillment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fulfillment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment) ProtoMessage() {}

func (x *Fulfillment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment.ProtoReflect.Descriptor instead.
func (*Fulfillment) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescGZIP(), []int{0}
}

func (x *Fulfillment) GetMessages() []*ResponseMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Fulfillment) GetWebhook() string {
	if x != nil {
		return x.Webhook
	}
	return ""
}

func (x *Fulfillment) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Fulfillment) GetSetParameterActions() []*Fulfillment_SetParameterAction {
	if x != nil {
		return x.SetParameterActions
	}
	return nil
}

func (x *Fulfillment) GetConditionalCases() []*Fulfillment_ConditionalCases {
	if x != nil {
		return x.ConditionalCases
	}
	return nil
}

// Setting a parameter value.
type Fulfillment_SetParameterAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display name of the parameter.
	Parameter string `protobuf:"bytes,1,opt,name=parameter,proto3" json:"parameter,omitempty"`
	// The new value of the parameter. A null value clears the parameter.
	Value *structpb.Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Fulfillment_SetParameterAction) Reset() {
	*x = Fulfillment_SetParameterAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fulfillment_SetParameterAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment_SetParameterAction) ProtoMessage() {}

func (x *Fulfillment_SetParameterAction) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment_SetParameterAction.ProtoReflect.Descriptor instead.
func (*Fulfillment_SetParameterAction) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Fulfillment_SetParameterAction) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

func (x *Fulfillment_SetParameterAction) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

// A list of cascading if-else conditions. Cases are mutually exclusive.
// The first one with a matching condition is selected, all the rest ignored.
type Fulfillment_ConditionalCases struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of cascading if-else conditions.
	Cases []*Fulfillment_ConditionalCases_Case `protobuf:"bytes,1,rep,name=cases,proto3" json:"cases,omitempty"`
}

func (x *Fulfillment_ConditionalCases) Reset() {
	*x = Fulfillment_ConditionalCases{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fulfillment_ConditionalCases) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment_ConditionalCases) ProtoMessage() {}

func (x *Fulfillment_ConditionalCases) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment_ConditionalCases.ProtoReflect.Descriptor instead.
func (*Fulfillment_ConditionalCases) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Fulfillment_ConditionalCases) GetCases() []*Fulfillment_ConditionalCases_Case {
	if x != nil {
		return x.Cases
	}
	return nil
}

// Each case has a Boolean condition. When it is evaluated to be True, the
// corresponding messages will be selected and evaluated recursively.
type Fulfillment_ConditionalCases_Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The condition to activate and select this case. Empty means the
	// condition is always true. The condition is evaluated against [form
	// parameters][Form.parameters] or [session
	// parameters][SessionInfo.parameters].
	//
	// See the [conditions
	// reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition).
	Condition string `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	// A list of case content.
	CaseContent []*Fulfillment_ConditionalCases_Case_CaseContent `protobuf:"bytes,2,rep,name=case_content,json=caseContent,proto3" json:"case_content,omitempty"`
}

func (x *Fulfillment_ConditionalCases_Case) Reset() {
	*x = Fulfillment_ConditionalCases_Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fulfillment_ConditionalCases_Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment_ConditionalCases_Case) ProtoMessage() {}

func (x *Fulfillment_ConditionalCases_Case) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment_ConditionalCases_Case.ProtoReflect.Descriptor instead.
func (*Fulfillment_ConditionalCases_Case) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Fulfillment_ConditionalCases_Case) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *Fulfillment_ConditionalCases_Case) GetCaseContent() []*Fulfillment_ConditionalCases_Case_CaseContent {
	if x != nil {
		return x.CaseContent
	}
	return nil
}

// The list of messages or conditional cases to activate for this case.
type Fulfillment_ConditionalCases_Case_CaseContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Either a message is returned or additional cases to be evaluated.
	//
	// Types that are assignable to CasesOrMessage:
	//	*Fulfillment_ConditionalCases_Case_CaseContent_Message
	//	*Fulfillment_ConditionalCases_Case_CaseContent_AdditionalCases
	CasesOrMessage isFulfillment_ConditionalCases_Case_CaseContent_CasesOrMessage `protobuf_oneof:"cases_or_message"`
}

func (x *Fulfillment_ConditionalCases_Case_CaseContent) Reset() {
	*x = Fulfillment_ConditionalCases_Case_CaseContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fulfillment_ConditionalCases_Case_CaseContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fulfillment_ConditionalCases_Case_CaseContent) ProtoMessage() {}

func (x *Fulfillment_ConditionalCases_Case_CaseContent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fulfillment_ConditionalCases_Case_CaseContent.ProtoReflect.Descriptor instead.
func (*Fulfillment_ConditionalCases_Case_CaseContent) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescGZIP(), []int{0, 1, 0, 0}
}

func (m *Fulfillment_ConditionalCases_Case_CaseContent) GetCasesOrMessage() isFulfillment_ConditionalCases_Case_CaseContent_CasesOrMessage {
	if m != nil {
		return m.CasesOrMessage
	}
	return nil
}

func (x *Fulfillment_ConditionalCases_Case_CaseContent) GetMessage() *ResponseMessage {
	if x, ok := x.GetCasesOrMessage().(*Fulfillment_ConditionalCases_Case_CaseContent_Message); ok {
		return x.Message
	}
	return nil
}

func (x *Fulfillment_ConditionalCases_Case_CaseContent) GetAdditionalCases() *Fulfillment_ConditionalCases {
	if x, ok := x.GetCasesOrMessage().(*Fulfillment_ConditionalCases_Case_CaseContent_AdditionalCases); ok {
		return x.AdditionalCases
	}
	return nil
}

type isFulfillment_ConditionalCases_Case_CaseContent_CasesOrMessage interface {
	isFulfillment_ConditionalCases_Case_CaseContent_CasesOrMessage()
}

type Fulfillment_ConditionalCases_Case_CaseContent_Message struct {
	// Returned message.
	Message *ResponseMessage `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

type Fulfillment_ConditionalCases_Case_CaseContent_AdditionalCases struct {
	// Additional cases to be evaluated.
	AdditionalCases *Fulfillment_ConditionalCases `protobuf:"bytes,2,opt,name=additional_cases,json=additionalCases,proto3,oneof"`
}

func (*Fulfillment_ConditionalCases_Case_CaseContent_Message) isFulfillment_ConditionalCases_Case_CaseContent_CasesOrMessage() {
}

func (*Fulfillment_ConditionalCases_Case_CaseContent_AdditionalCases) isFulfillment_ConditionalCases_Case_CaseContent_CasesOrMessage() {
}

var File_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x78, 0x2f, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x07, 0x0a,
	0x0b, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x40, 0x0a,
	0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26,
	0xfa, 0x41, 0x23, 0x0a, 0x21, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x76, 0x0a, 0x15, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x73, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6d, 0x0a, 0x11, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63,
	0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x60, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xf0, 0x03, 0x0a, 0x10, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x5b, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0xfe, 0x02, 0x0a,
	0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x74, 0x0a, 0x0c, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x61,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0xe1, 0x01, 0x0a, 0x0b, 0x43, 0x61,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63,
	0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x5f, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xaf, 0x01,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78,
	0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x10, 0x46, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x63, 0x78, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33, 0x42, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_goTypes = []interface{}{
	(*Fulfillment)(nil),                                   // 0: google.cloud.dialogflow.cx.v3beta1.Fulfillment
	(*Fulfillment_SetParameterAction)(nil),                // 1: google.cloud.dialogflow.cx.v3beta1.Fulfillment.SetParameterAction
	(*Fulfillment_ConditionalCases)(nil),                  // 2: google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases
	(*Fulfillment_ConditionalCases_Case)(nil),             // 3: google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case
	(*Fulfillment_ConditionalCases_Case_CaseContent)(nil), // 4: google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.CaseContent
	(*ResponseMessage)(nil),                               // 5: google.cloud.dialogflow.cx.v3beta1.ResponseMessage
	(*structpb.Value)(nil),                                // 6: google.protobuf.Value
}
var file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_depIdxs = []int32{
	5, // 0: google.cloud.dialogflow.cx.v3beta1.Fulfillment.messages:type_name -> google.cloud.dialogflow.cx.v3beta1.ResponseMessage
	1, // 1: google.cloud.dialogflow.cx.v3beta1.Fulfillment.set_parameter_actions:type_name -> google.cloud.dialogflow.cx.v3beta1.Fulfillment.SetParameterAction
	2, // 2: google.cloud.dialogflow.cx.v3beta1.Fulfillment.conditional_cases:type_name -> google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases
	6, // 3: google.cloud.dialogflow.cx.v3beta1.Fulfillment.SetParameterAction.value:type_name -> google.protobuf.Value
	3, // 4: google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.cases:type_name -> google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case
	4, // 5: google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.case_content:type_name -> google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.CaseContent
	5, // 6: google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.CaseContent.message:type_name -> google.cloud.dialogflow.cx.v3beta1.ResponseMessage
	2, // 7: google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases.Case.CaseContent.additional_cases:type_name -> google.cloud.dialogflow.cx.v3beta1.Fulfillment.ConditionalCases
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto != nil {
		return
	}
	file_google_cloud_dialogflow_cx_v3beta1_response_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fulfillment); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fulfillment_SetParameterAction); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fulfillment_ConditionalCases); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fulfillment_ConditionalCases_Case); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fulfillment_ConditionalCases_Case_CaseContent); i {
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
	file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Fulfillment_ConditionalCases_Case_CaseContent_Message)(nil),
		(*Fulfillment_ConditionalCases_Case_CaseContent_AdditionalCases)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_depIdxs,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_fulfillment_proto_depIdxs = nil
}
