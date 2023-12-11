// Copyright 2015 gRPC authors.
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
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: routeguide/route_guide.proto

package routeguide

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

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi *Point `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (x *Rectangle) Reset() {
	*x = Rectangle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rectangle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rectangle) ProtoMessage() {}

func (x *Rectangle) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rectangle.ProtoReflect.Descriptor instead.
func (*Rectangle) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{1}
}

func (x *Rectangle) GetLo() *Point {
	if x != nil {
		return x.Lo
	}
	return nil
}

func (x *Rectangle) GetHi() *Point {
	if x != nil {
		return x.Hi
	}
	return nil
}

// A feature names something at a given point.
//
// If a feature could not be named, the name is empty.
type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The point where the feature is detected.
	Location *Point `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{2}
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location from which the message is sent.
	Location *Point `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The message to be sent.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RouteNote) Reset() {
	*x = RouteNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNote) ProtoMessage() {}

func (x *RouteNote) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNote.ProtoReflect.Descriptor instead.
func (*RouteNote) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{3}
}

func (x *RouteNote) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *RouteNote) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of points received.
	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount,proto3" json:"point_count,omitempty"`
	// The number of known features passed while traversing the route.
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount,proto3" json:"feature_count,omitempty"`
	// The distance covered in metres.
	Distance int32 `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	// The duration of the traversal in seconds.
	ElapsedTime int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
}

func (x *RouteSummary) Reset() {
	*x = RouteSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_guide_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteSummary) ProtoMessage() {}

func (x *RouteSummary) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_guide_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteSummary.ProtoReflect.Descriptor instead.
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return file_routeguide_route_guide_proto_rawDescGZIP(), []int{4}
}

func (x *RouteSummary) GetPointCount() int32 {
	if x != nil {
		return x.PointCount
	}
	return 0
}

func (x *RouteSummary) GetFeatureCount() int32 {
	if x != nil {
		return x.FeatureCount
	}
	return 0
}

func (x *RouteSummary) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *RouteSummary) GetElapsedTime() int32 {
	if x != nil {
		return x.ElapsedTime
	}
	return 0
}

var File_routeguide_route_guide_proto protoreflect.FileDescriptor

var file_routeguide_route_guide_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x51, 0x0a,
	0x09, 0x52, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x02, 0x6c, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x6c, 0x6f, 0x12, 0x21, 0x0a,
	0x02, 0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x68, 0x69,
	0x22, 0x4c, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54,
	0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x61, 0x70, 0x73,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65,
	0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x85, 0x02, 0x0a, 0x0a, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x12, 0x15, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x52,
	0x65, 0x63, 0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3e, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x3f, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x15,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x68, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64,
	0x65, 0x42, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routeguide_route_guide_proto_rawDescOnce sync.Once
	file_routeguide_route_guide_proto_rawDescData = file_routeguide_route_guide_proto_rawDesc
)

func file_routeguide_route_guide_proto_rawDescGZIP() []byte {
	file_routeguide_route_guide_proto_rawDescOnce.Do(func() {
		file_routeguide_route_guide_proto_rawDescData = protoimpl.X.CompressGZIP(file_routeguide_route_guide_proto_rawDescData)
	})
	return file_routeguide_route_guide_proto_rawDescData
}

var file_routeguide_route_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_routeguide_route_guide_proto_goTypes = []interface{}{
	(*Point)(nil),        // 0: routeguide.Point
	(*Rectangle)(nil),    // 1: routeguide.Rectangle
	(*Feature)(nil),      // 2: routeguide.Feature
	(*RouteNote)(nil),    // 3: routeguide.RouteNote
	(*RouteSummary)(nil), // 4: routeguide.RouteSummary
}
var file_routeguide_route_guide_proto_depIdxs = []int32{
	0, // 0: routeguide.Rectangle.lo:type_name -> routeguide.Point
	0, // 1: routeguide.Rectangle.hi:type_name -> routeguide.Point
	0, // 2: routeguide.Feature.location:type_name -> routeguide.Point
	0, // 3: routeguide.RouteNote.location:type_name -> routeguide.Point
	0, // 4: routeguide.RouteGuide.GetFeature:input_type -> routeguide.Point
	1, // 5: routeguide.RouteGuide.ListFeatures:input_type -> routeguide.Rectangle
	0, // 6: routeguide.RouteGuide.RecordRoute:input_type -> routeguide.Point
	3, // 7: routeguide.RouteGuide.RouteChat:input_type -> routeguide.RouteNote
	2, // 8: routeguide.RouteGuide.GetFeature:output_type -> routeguide.Feature
	2, // 9: routeguide.RouteGuide.ListFeatures:output_type -> routeguide.Feature
	4, // 10: routeguide.RouteGuide.RecordRoute:output_type -> routeguide.RouteSummary
	3, // 11: routeguide.RouteGuide.RouteChat:output_type -> routeguide.RouteNote
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_routeguide_route_guide_proto_init() }
func file_routeguide_route_guide_proto_init() {
	if File_routeguide_route_guide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routeguide_route_guide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_routeguide_route_guide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rectangle); i {
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
		file_routeguide_route_guide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_routeguide_route_guide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNote); i {
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
		file_routeguide_route_guide_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteSummary); i {
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
			RawDescriptor: file_routeguide_route_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routeguide_route_guide_proto_goTypes,
		DependencyIndexes: file_routeguide_route_guide_proto_depIdxs,
		MessageInfos:      file_routeguide_route_guide_proto_msgTypes,
	}.Build()
	File_routeguide_route_guide_proto = out.File
	file_routeguide_route_guide_proto_rawDesc = nil
	file_routeguide_route_guide_proto_goTypes = nil
	file_routeguide_route_guide_proto_depIdxs = nil
}
