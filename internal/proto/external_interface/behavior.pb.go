// Code generated by protoc-gen-go. DO NOT EDIT.
// source: behavior.proto

package Anki_Vector_external_interface

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Where in the behavior tree the SDK code should be executed.
type ControlRequest_Priority int32

const (
	// Unknown priority. Used for versions that don't understand old priority levels.
	ControlRequest_UNKNOWN ControlRequest_Priority = 0
	// Highest priority level. Suppresses most automatic physical reactions, use with caution.
	ControlRequest_OVERRIDE_BEHAVIORS ControlRequest_Priority = 10
	// Normal priority level. Directly under mandatory physical reactions.
	ControlRequest_DEFAULT ControlRequest_Priority = 20
	// Enable long-running SDK control between script execution.  Not to be used for regular behavior control.
	ControlRequest_RESERVE_CONTROL ControlRequest_Priority = 30
)

var ControlRequest_Priority_name = map[int32]string{
	0:  "UNKNOWN",
	10: "OVERRIDE_BEHAVIORS",
	20: "DEFAULT",
	30: "RESERVE_CONTROL",
}
var ControlRequest_Priority_value = map[string]int32{
	"UNKNOWN":            0,
	"OVERRIDE_BEHAVIORS": 10,
	"DEFAULT":            20,
	"RESERVE_CONTROL":    30,
}

func (x ControlRequest_Priority) String() string {
	return proto.EnumName(ControlRequest_Priority_name, int32(x))
}
func (ControlRequest_Priority) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{1, 0}
}

// Tell the behavior stream to release control from the SDK.
// The stream may stay alive, but Vector will be allowed to run
// his normal behaviors.
type ControlRelease struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlRelease) Reset()         { *m = ControlRelease{} }
func (m *ControlRelease) String() string { return proto.CompactTextString(m) }
func (*ControlRelease) ProtoMessage()    {}
func (*ControlRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{0}
}
func (m *ControlRelease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlRelease.Unmarshal(m, b)
}
func (m *ControlRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlRelease.Marshal(b, m, deterministic)
}
func (dst *ControlRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlRelease.Merge(dst, src)
}
func (m *ControlRelease) XXX_Size() int {
	return xxx_messageInfo_ControlRelease.Size(m)
}
func (m *ControlRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlRelease.DiscardUnknown(m)
}

var xxx_messageInfo_ControlRelease proto.InternalMessageInfo

// Request control of the behavior system at a given priority.
// Currently there is only one priority level.
type ControlRequest struct {
	// Where in the behavior tree the SDK code should be executed.
	Priority             ControlRequest_Priority `protobuf:"varint,1,opt,name=priority,enum=Anki.Vector.external_interface.ControlRequest_Priority" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ControlRequest) Reset()         { *m = ControlRequest{} }
func (m *ControlRequest) String() string { return proto.CompactTextString(m) }
func (*ControlRequest) ProtoMessage()    {}
func (*ControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{1}
}
func (m *ControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlRequest.Unmarshal(m, b)
}
func (m *ControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlRequest.Marshal(b, m, deterministic)
}
func (dst *ControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlRequest.Merge(dst, src)
}
func (m *ControlRequest) XXX_Size() int {
	return xxx_messageInfo_ControlRequest.Size(m)
}
func (m *ControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControlRequest proto.InternalMessageInfo

func (m *ControlRequest) GetPriority() ControlRequest_Priority {
	if m != nil {
		return m.Priority
	}
	return ControlRequest_UNKNOWN
}

// Messages that can be sent to the behavior stream. Explicitly
// requesting or releasing control.
type BehaviorControlRequest struct {
	// Types that are valid to be assigned to RequestType:
	//	*BehaviorControlRequest_ControlRelease
	//	*BehaviorControlRequest_ControlRequest
	RequestType          isBehaviorControlRequest_RequestType `protobuf_oneof:"request_type"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *BehaviorControlRequest) Reset()         { *m = BehaviorControlRequest{} }
func (m *BehaviorControlRequest) String() string { return proto.CompactTextString(m) }
func (*BehaviorControlRequest) ProtoMessage()    {}
func (*BehaviorControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{2}
}
func (m *BehaviorControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BehaviorControlRequest.Unmarshal(m, b)
}
func (m *BehaviorControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BehaviorControlRequest.Marshal(b, m, deterministic)
}
func (dst *BehaviorControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BehaviorControlRequest.Merge(dst, src)
}
func (m *BehaviorControlRequest) XXX_Size() int {
	return xxx_messageInfo_BehaviorControlRequest.Size(m)
}
func (m *BehaviorControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BehaviorControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BehaviorControlRequest proto.InternalMessageInfo

type isBehaviorControlRequest_RequestType interface {
	isBehaviorControlRequest_RequestType()
}

type BehaviorControlRequest_ControlRelease struct {
	ControlRelease *ControlRelease `protobuf:"bytes,1,opt,name=control_release,json=controlRelease,oneof"`
}
type BehaviorControlRequest_ControlRequest struct {
	ControlRequest *ControlRequest `protobuf:"bytes,2,opt,name=control_request,json=controlRequest,oneof"`
}

func (*BehaviorControlRequest_ControlRelease) isBehaviorControlRequest_RequestType() {}
func (*BehaviorControlRequest_ControlRequest) isBehaviorControlRequest_RequestType() {}

func (m *BehaviorControlRequest) GetRequestType() isBehaviorControlRequest_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (m *BehaviorControlRequest) GetControlRelease() *ControlRelease {
	if x, ok := m.GetRequestType().(*BehaviorControlRequest_ControlRelease); ok {
		return x.ControlRelease
	}
	return nil
}

func (m *BehaviorControlRequest) GetControlRequest() *ControlRequest {
	if x, ok := m.GetRequestType().(*BehaviorControlRequest_ControlRequest); ok {
		return x.ControlRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BehaviorControlRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BehaviorControlRequest_OneofMarshaler, _BehaviorControlRequest_OneofUnmarshaler, _BehaviorControlRequest_OneofSizer, []interface{}{
		(*BehaviorControlRequest_ControlRelease)(nil),
		(*BehaviorControlRequest_ControlRequest)(nil),
	}
}

func _BehaviorControlRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BehaviorControlRequest)
	// request_type
	switch x := m.RequestType.(type) {
	case *BehaviorControlRequest_ControlRelease:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ControlRelease); err != nil {
			return err
		}
	case *BehaviorControlRequest_ControlRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ControlRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BehaviorControlRequest.RequestType has unexpected type %T", x)
	}
	return nil
}

func _BehaviorControlRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BehaviorControlRequest)
	switch tag {
	case 1: // request_type.control_release
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ControlRelease)
		err := b.DecodeMessage(msg)
		m.RequestType = &BehaviorControlRequest_ControlRelease{msg}
		return true, err
	case 2: // request_type.control_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ControlRequest)
		err := b.DecodeMessage(msg)
		m.RequestType = &BehaviorControlRequest_ControlRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BehaviorControlRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BehaviorControlRequest)
	// request_type
	switch x := m.RequestType.(type) {
	case *BehaviorControlRequest_ControlRelease:
		s := proto.Size(x.ControlRelease)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BehaviorControlRequest_ControlRequest:
		s := proto.Size(x.ControlRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The SDK user is now free to run any actions and behaviors they like.
// Until a ControlLostResponse is received, they are directly in control
// of Vector's behavior system.
type ControlGrantedResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlGrantedResponse) Reset()         { *m = ControlGrantedResponse{} }
func (m *ControlGrantedResponse) String() string { return proto.CompactTextString(m) }
func (*ControlGrantedResponse) ProtoMessage()    {}
func (*ControlGrantedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{3}
}
func (m *ControlGrantedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlGrantedResponse.Unmarshal(m, b)
}
func (m *ControlGrantedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlGrantedResponse.Marshal(b, m, deterministic)
}
func (dst *ControlGrantedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlGrantedResponse.Merge(dst, src)
}
func (m *ControlGrantedResponse) XXX_Size() int {
	return xxx_messageInfo_ControlGrantedResponse.Size(m)
}
func (m *ControlGrantedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlGrantedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlGrantedResponse proto.InternalMessageInfo

// This informs the user that they lost control of the behavior system.
// All direct actions will be unavailable via the sdk until control is regained.
// Regaining control can be either through a call to ControlRequest, or
// can be as a result of conditions passed to the original ControlRequest.
type ControlLostResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlLostResponse) Reset()         { *m = ControlLostResponse{} }
func (m *ControlLostResponse) String() string { return proto.CompactTextString(m) }
func (*ControlLostResponse) ProtoMessage()    {}
func (*ControlLostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{4}
}
func (m *ControlLostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlLostResponse.Unmarshal(m, b)
}
func (m *ControlLostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlLostResponse.Marshal(b, m, deterministic)
}
func (dst *ControlLostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlLostResponse.Merge(dst, src)
}
func (m *ControlLostResponse) XXX_Size() int {
	return xxx_messageInfo_ControlLostResponse.Size(m)
}
func (m *ControlLostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlLostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlLostResponse proto.InternalMessageInfo

// The ability to reserve control before/after SDK scripts has been lost.
// This control can be regained through another ControlRequest.
type ReservedControlLostResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReservedControlLostResponse) Reset()         { *m = ReservedControlLostResponse{} }
func (m *ReservedControlLostResponse) String() string { return proto.CompactTextString(m) }
func (*ReservedControlLostResponse) ProtoMessage()    {}
func (*ReservedControlLostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{5}
}
func (m *ReservedControlLostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReservedControlLostResponse.Unmarshal(m, b)
}
func (m *ReservedControlLostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReservedControlLostResponse.Marshal(b, m, deterministic)
}
func (dst *ReservedControlLostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReservedControlLostResponse.Merge(dst, src)
}
func (m *ReservedControlLostResponse) XXX_Size() int {
	return xxx_messageInfo_ReservedControlLostResponse.Size(m)
}
func (m *ReservedControlLostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReservedControlLostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReservedControlLostResponse proto.InternalMessageInfo

// Responses from the behavior stream.
type BehaviorControlResponse struct {
	// Types that are valid to be assigned to ResponseType:
	//	*BehaviorControlResponse_ControlGrantedResponse
	//	*BehaviorControlResponse_ControlLostEvent
	//	*BehaviorControlResponse_KeepAlive
	//	*BehaviorControlResponse_ReservedControlLostEvent
	ResponseType         isBehaviorControlResponse_ResponseType `protobuf_oneof:"response_type"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *BehaviorControlResponse) Reset()         { *m = BehaviorControlResponse{} }
func (m *BehaviorControlResponse) String() string { return proto.CompactTextString(m) }
func (*BehaviorControlResponse) ProtoMessage()    {}
func (*BehaviorControlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_behavior_e7775c3c01601f75, []int{6}
}
func (m *BehaviorControlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BehaviorControlResponse.Unmarshal(m, b)
}
func (m *BehaviorControlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BehaviorControlResponse.Marshal(b, m, deterministic)
}
func (dst *BehaviorControlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BehaviorControlResponse.Merge(dst, src)
}
func (m *BehaviorControlResponse) XXX_Size() int {
	return xxx_messageInfo_BehaviorControlResponse.Size(m)
}
func (m *BehaviorControlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BehaviorControlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BehaviorControlResponse proto.InternalMessageInfo

type isBehaviorControlResponse_ResponseType interface {
	isBehaviorControlResponse_ResponseType()
}

type BehaviorControlResponse_ControlGrantedResponse struct {
	ControlGrantedResponse *ControlGrantedResponse `protobuf:"bytes,1,opt,name=control_granted_response,json=controlGrantedResponse,oneof"`
}
type BehaviorControlResponse_ControlLostEvent struct {
	ControlLostEvent *ControlLostResponse `protobuf:"bytes,2,opt,name=control_lost_event,json=controlLostEvent,oneof"`
}
type BehaviorControlResponse_KeepAlive struct {
	KeepAlive *KeepAlivePing `protobuf:"bytes,3,opt,name=keep_alive,json=keepAlive,oneof"`
}
type BehaviorControlResponse_ReservedControlLostEvent struct {
	ReservedControlLostEvent *ReservedControlLostResponse `protobuf:"bytes,4,opt,name=reserved_control_lost_event,json=reservedControlLostEvent,oneof"`
}

func (*BehaviorControlResponse_ControlGrantedResponse) isBehaviorControlResponse_ResponseType()   {}
func (*BehaviorControlResponse_ControlLostEvent) isBehaviorControlResponse_ResponseType()         {}
func (*BehaviorControlResponse_KeepAlive) isBehaviorControlResponse_ResponseType()                {}
func (*BehaviorControlResponse_ReservedControlLostEvent) isBehaviorControlResponse_ResponseType() {}

func (m *BehaviorControlResponse) GetResponseType() isBehaviorControlResponse_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return nil
}

func (m *BehaviorControlResponse) GetControlGrantedResponse() *ControlGrantedResponse {
	if x, ok := m.GetResponseType().(*BehaviorControlResponse_ControlGrantedResponse); ok {
		return x.ControlGrantedResponse
	}
	return nil
}

func (m *BehaviorControlResponse) GetControlLostEvent() *ControlLostResponse {
	if x, ok := m.GetResponseType().(*BehaviorControlResponse_ControlLostEvent); ok {
		return x.ControlLostEvent
	}
	return nil
}

func (m *BehaviorControlResponse) GetKeepAlive() *KeepAlivePing {
	if x, ok := m.GetResponseType().(*BehaviorControlResponse_KeepAlive); ok {
		return x.KeepAlive
	}
	return nil
}

func (m *BehaviorControlResponse) GetReservedControlLostEvent() *ReservedControlLostResponse {
	if x, ok := m.GetResponseType().(*BehaviorControlResponse_ReservedControlLostEvent); ok {
		return x.ReservedControlLostEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BehaviorControlResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BehaviorControlResponse_OneofMarshaler, _BehaviorControlResponse_OneofUnmarshaler, _BehaviorControlResponse_OneofSizer, []interface{}{
		(*BehaviorControlResponse_ControlGrantedResponse)(nil),
		(*BehaviorControlResponse_ControlLostEvent)(nil),
		(*BehaviorControlResponse_KeepAlive)(nil),
		(*BehaviorControlResponse_ReservedControlLostEvent)(nil),
	}
}

func _BehaviorControlResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BehaviorControlResponse)
	// response_type
	switch x := m.ResponseType.(type) {
	case *BehaviorControlResponse_ControlGrantedResponse:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ControlGrantedResponse); err != nil {
			return err
		}
	case *BehaviorControlResponse_ControlLostEvent:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ControlLostEvent); err != nil {
			return err
		}
	case *BehaviorControlResponse_KeepAlive:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeepAlive); err != nil {
			return err
		}
	case *BehaviorControlResponse_ReservedControlLostEvent:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReservedControlLostEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BehaviorControlResponse.ResponseType has unexpected type %T", x)
	}
	return nil
}

func _BehaviorControlResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BehaviorControlResponse)
	switch tag {
	case 1: // response_type.control_granted_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ControlGrantedResponse)
		err := b.DecodeMessage(msg)
		m.ResponseType = &BehaviorControlResponse_ControlGrantedResponse{msg}
		return true, err
	case 2: // response_type.control_lost_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ControlLostResponse)
		err := b.DecodeMessage(msg)
		m.ResponseType = &BehaviorControlResponse_ControlLostEvent{msg}
		return true, err
	case 3: // response_type.keep_alive
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KeepAlivePing)
		err := b.DecodeMessage(msg)
		m.ResponseType = &BehaviorControlResponse_KeepAlive{msg}
		return true, err
	case 4: // response_type.reserved_control_lost_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReservedControlLostResponse)
		err := b.DecodeMessage(msg)
		m.ResponseType = &BehaviorControlResponse_ReservedControlLostEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BehaviorControlResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BehaviorControlResponse)
	// response_type
	switch x := m.ResponseType.(type) {
	case *BehaviorControlResponse_ControlGrantedResponse:
		s := proto.Size(x.ControlGrantedResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BehaviorControlResponse_ControlLostEvent:
		s := proto.Size(x.ControlLostEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BehaviorControlResponse_KeepAlive:
		s := proto.Size(x.KeepAlive)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BehaviorControlResponse_ReservedControlLostEvent:
		s := proto.Size(x.ReservedControlLostEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ControlRelease)(nil), "Anki.Vector.external_interface.ControlRelease")
	proto.RegisterType((*ControlRequest)(nil), "Anki.Vector.external_interface.ControlRequest")
	proto.RegisterType((*BehaviorControlRequest)(nil), "Anki.Vector.external_interface.BehaviorControlRequest")
	proto.RegisterType((*ControlGrantedResponse)(nil), "Anki.Vector.external_interface.ControlGrantedResponse")
	proto.RegisterType((*ControlLostResponse)(nil), "Anki.Vector.external_interface.ControlLostResponse")
	proto.RegisterType((*ReservedControlLostResponse)(nil), "Anki.Vector.external_interface.ReservedControlLostResponse")
	proto.RegisterType((*BehaviorControlResponse)(nil), "Anki.Vector.external_interface.BehaviorControlResponse")
	proto.RegisterEnum("Anki.Vector.external_interface.ControlRequest_Priority", ControlRequest_Priority_name, ControlRequest_Priority_value)
}

func init() { proto.RegisterFile("behavior.proto", fileDescriptor_behavior_e7775c3c01601f75) }

var fileDescriptor_behavior_e7775c3c01601f75 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x4d, 0x19, 0x82, 0xf1, 0x0d, 0xd2, 0xc8, 0x83, 0x12, 0x31, 0x31, 0xa1, 0x9c, 0xb8, 0x90,
	0xc3, 0x26, 0xc1, 0x81, 0x53, 0xbb, 0x19, 0x32, 0xad, 0x4a, 0x86, 0xbb, 0x15, 0x71, 0xb2, 0xb2,
	0xec, 0xa3, 0x44, 0x0d, 0x76, 0xb0, 0x4d, 0xc5, 0x24, 0x7e, 0x18, 0x7f, 0x83, 0x3f, 0xc0, 0x6f,
	0x41, 0x75, 0x9c, 0xb1, 0x6e, 0xd5, 0x56, 0xae, 0xef, 0xf9, 0xbd, 0xe7, 0xf7, 0x9c, 0x80, 0x7f,
	0x8a, 0x5f, 0xf2, 0x59, 0x29, 0x55, 0x5c, 0x2b, 0x69, 0x24, 0xd9, 0xee, 0x8b, 0x69, 0x19, 0x8f,
	0xb1, 0x30, 0x52, 0xc5, 0xf8, 0xc3, 0xa0, 0x12, 0x79, 0xc5, 0x4b, 0x61, 0x50, 0x7d, 0xce, 0x0b,
	0x7c, 0xe6, 0x7f, 0x45, 0xad, 0xf3, 0x09, 0xea, 0xe6, 0x7c, 0x14, 0x80, 0xbf, 0x27, 0x85, 0x51,
	0xb2, 0x62, 0x58, 0x61, 0xae, 0x31, 0xfa, 0xd5, 0xb9, 0x04, 0x7d, 0xfb, 0x8e, 0xda, 0x90, 0x11,
	0xac, 0xd7, 0xaa, 0x94, 0xaa, 0x34, 0xe7, 0x61, 0xe7, 0x45, 0xe7, 0xa5, 0xbf, 0xf3, 0x26, 0xbe,
	0x39, 0x27, 0x5e, 0x74, 0x88, 0x8f, 0x9c, 0x9c, 0x5d, 0x18, 0x45, 0x1f, 0x60, 0xbd, 0x45, 0xc9,
	0x06, 0xdc, 0x3f, 0x49, 0x0f, 0xd3, 0xec, 0x63, 0x1a, 0x78, 0xa4, 0x07, 0x24, 0x1b, 0x53, 0xc6,
	0x0e, 0xf6, 0x29, 0x1f, 0xd0, 0xa4, 0x3f, 0x3e, 0xc8, 0xd8, 0x28, 0x80, 0xf9, 0xa1, 0x7d, 0xfa,
	0xae, 0x7f, 0x32, 0x3c, 0x0e, 0x1e, 0x93, 0x4d, 0xe8, 0x32, 0x3a, 0xa2, 0x6c, 0x4c, 0xf9, 0x5e,
	0x96, 0x1e, 0xb3, 0x6c, 0x18, 0x6c, 0x47, 0x7f, 0x3a, 0xd0, 0x1b, 0xb8, 0x3d, 0xae, 0x54, 0xf8,
	0x04, 0xdd, 0xa2, 0x41, 0xb8, 0x6a, 0x8a, 0xda, 0x26, 0x1b, 0x3b, 0xf1, 0xca, 0x4d, 0xac, 0x2a,
	0xf1, 0x98, 0x5f, 0x2c, 0x20, 0x8b, 0xd6, 0x36, 0x2d, 0xbc, 0xf3, 0x9f, 0xd6, 0x56, 0xb5, 0x60,
	0x6d, 0x91, 0x81, 0x0f, 0x0f, 0x9d, 0x25, 0x37, 0xe7, 0x35, 0x46, 0x21, 0xf4, 0x9c, 0xe6, 0xbd,
	0xca, 0x85, 0xc1, 0x33, 0x86, 0xba, 0x96, 0x42, 0x63, 0xf4, 0x04, 0x36, 0x1d, 0x33, 0x94, 0xda,
	0x5c, 0xc0, 0xcf, 0x61, 0x8b, 0xa1, 0x46, 0x35, 0xc3, 0xb3, 0x65, 0xf4, 0xef, 0x35, 0x78, 0x7a,
	0x6d, 0xb0, 0x86, 0x23, 0x0a, 0xc2, 0xb6, 0xd6, 0xa4, 0x09, 0xe3, 0xca, 0x71, 0x6e, 0xba, 0xd7,
	0x2b, 0xf6, 0xbb, 0x72, 0xd7, 0xc4, 0x63, 0xbd, 0x62, 0x29, 0x43, 0x0a, 0x20, 0x6d, 0x66, 0x25,
	0xb5, 0xe1, 0x38, 0x43, 0xd1, 0xae, 0xb9, 0xbb, 0x62, 0xda, 0xe5, 0x82, 0x89, 0xc7, 0x82, 0xe2,
	0x1f, 0x4c, 0xe7, 0x76, 0x24, 0x05, 0x98, 0x22, 0xd6, 0x3c, 0xaf, 0xca, 0x19, 0x86, 0x6b, 0xd6,
	0xfc, 0xd5, 0x6d, 0xe6, 0x87, 0x88, 0x75, 0x7f, 0x2e, 0x38, 0x2a, 0xc5, 0x24, 0xf1, 0xd8, 0x83,
	0x69, 0x0b, 0x90, 0x9f, 0xb0, 0xa5, 0xdc, 0xc6, 0x7c, 0xc9, 0xed, 0xef, 0xda, 0x80, 0xb7, 0xb7,
	0x05, 0xdc, 0xf0, 0x4c, 0x89, 0xc7, 0x42, 0x75, 0x9d, 0xb6, 0x6d, 0x06, 0x5d, 0x78, 0xd4, 0x3e,
	0x8b, 0xfd, 0x46, 0x4e, 0xef, 0xd9, 0x1f, 0x7b, 0xf7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14,
	0xdc, 0x29, 0x68, 0x1a, 0x04, 0x00, 0x00,
}
