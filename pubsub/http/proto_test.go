// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package pubsub is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	TestProto
*/
package http

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type TestProto struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *TestProto) Reset()         { *m = TestProto{} }
func (m *TestProto) String() string { return proto.CompactTextString(m) }
func (*TestProto) ProtoMessage()    {}

func init() {
}
