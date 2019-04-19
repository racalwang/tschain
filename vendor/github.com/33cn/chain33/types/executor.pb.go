// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Genesis struct {
	Isrun                bool     `protobuf:"varint,1,opt,name=isrun,proto3" json:"isrun,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetIsrun() bool {
	if m != nil {
		return m.Isrun
	}
	return false
}

type ExecTxList struct {
	StateHash            []byte         `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	ParentHash           []byte         `protobuf:"bytes,7,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	MainHash             []byte         `protobuf:"bytes,8,opt,name=mainHash,proto3" json:"mainHash,omitempty"`
	MainHeight           int64          `protobuf:"varint,9,opt,name=mainHeight,proto3" json:"mainHeight,omitempty"`
	BlockTime            int64          `protobuf:"varint,3,opt,name=blockTime,proto3" json:"blockTime,omitempty"`
	Height               int64          `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Difficulty           uint64         `protobuf:"varint,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	IsMempool            bool           `protobuf:"varint,6,opt,name=isMempool,proto3" json:"isMempool,omitempty"`
	Txs                  []*Transaction `protobuf:"bytes,2,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecTxList) Reset()         { *m = ExecTxList{} }
func (m *ExecTxList) String() string { return proto.CompactTextString(m) }
func (*ExecTxList) ProtoMessage()    {}
func (*ExecTxList) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{1}
}

func (m *ExecTxList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecTxList.Unmarshal(m, b)
}
func (m *ExecTxList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecTxList.Marshal(b, m, deterministic)
}
func (m *ExecTxList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecTxList.Merge(m, src)
}
func (m *ExecTxList) XXX_Size() int {
	return xxx_messageInfo_ExecTxList.Size(m)
}
func (m *ExecTxList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecTxList.DiscardUnknown(m)
}

var xxx_messageInfo_ExecTxList proto.InternalMessageInfo

func (m *ExecTxList) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ExecTxList) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *ExecTxList) GetMainHash() []byte {
	if m != nil {
		return m.MainHash
	}
	return nil
}

func (m *ExecTxList) GetMainHeight() int64 {
	if m != nil {
		return m.MainHeight
	}
	return 0
}

func (m *ExecTxList) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *ExecTxList) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ExecTxList) GetDifficulty() uint64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *ExecTxList) GetIsMempool() bool {
	if m != nil {
		return m.IsMempool
	}
	return false
}

func (m *ExecTxList) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Query struct {
	Execer               []byte   `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	FuncName             string   `protobuf:"bytes,2,opt,name=funcName,proto3" json:"funcName,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{2}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *Query) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *Query) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateTxIn struct {
	Execer               []byte   `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	ActionName           string   `protobuf:"bytes,2,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTxIn) Reset()         { *m = CreateTxIn{} }
func (m *CreateTxIn) String() string { return proto.CompactTextString(m) }
func (*CreateTxIn) ProtoMessage()    {}
func (*CreateTxIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{3}
}

func (m *CreateTxIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTxIn.Unmarshal(m, b)
}
func (m *CreateTxIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTxIn.Marshal(b, m, deterministic)
}
func (m *CreateTxIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTxIn.Merge(m, src)
}
func (m *CreateTxIn) XXX_Size() int {
	return xxx_messageInfo_CreateTxIn.Size(m)
}
func (m *CreateTxIn) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTxIn.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTxIn proto.InternalMessageInfo

func (m *CreateTxIn) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *CreateTxIn) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *CreateTxIn) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// 配置修改部分
type ArrayConfig struct {
	Value                []string `protobuf:"bytes,3,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayConfig) Reset()         { *m = ArrayConfig{} }
func (m *ArrayConfig) String() string { return proto.CompactTextString(m) }
func (*ArrayConfig) ProtoMessage()    {}
func (*ArrayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{4}
}

func (m *ArrayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayConfig.Unmarshal(m, b)
}
func (m *ArrayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayConfig.Marshal(b, m, deterministic)
}
func (m *ArrayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayConfig.Merge(m, src)
}
func (m *ArrayConfig) XXX_Size() int {
	return xxx_messageInfo_ArrayConfig.Size(m)
}
func (m *ArrayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayConfig proto.InternalMessageInfo

func (m *ArrayConfig) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringConfig struct {
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringConfig) Reset()         { *m = StringConfig{} }
func (m *StringConfig) String() string { return proto.CompactTextString(m) }
func (*StringConfig) ProtoMessage()    {}
func (*StringConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{5}
}

func (m *StringConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringConfig.Unmarshal(m, b)
}
func (m *StringConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringConfig.Marshal(b, m, deterministic)
}
func (m *StringConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringConfig.Merge(m, src)
}
func (m *StringConfig) XXX_Size() int {
	return xxx_messageInfo_StringConfig.Size(m)
}
func (m *StringConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StringConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StringConfig proto.InternalMessageInfo

func (m *StringConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Int32Config struct {
	Value                int32    `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32Config) Reset()         { *m = Int32Config{} }
func (m *Int32Config) String() string { return proto.CompactTextString(m) }
func (*Int32Config) ProtoMessage()    {}
func (*Int32Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{6}
}

func (m *Int32Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32Config.Unmarshal(m, b)
}
func (m *Int32Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32Config.Marshal(b, m, deterministic)
}
func (m *Int32Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32Config.Merge(m, src)
}
func (m *Int32Config) XXX_Size() int {
	return xxx_messageInfo_Int32Config.Size(m)
}
func (m *Int32Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32Config.DiscardUnknown(m)
}

var xxx_messageInfo_Int32Config proto.InternalMessageInfo

func (m *Int32Config) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ConfigItem struct {
	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*ConfigItem_Arr
	//	*ConfigItem_Str
	//	*ConfigItem_Int
	Value                isConfigItem_Value `protobuf_oneof:"value"`
	Ty                   int32              `protobuf:"varint,11,opt,name=Ty,proto3" json:"Ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ConfigItem) Reset()         { *m = ConfigItem{} }
func (m *ConfigItem) String() string { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()    {}
func (*ConfigItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{7}
}

func (m *ConfigItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigItem.Unmarshal(m, b)
}
func (m *ConfigItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigItem.Marshal(b, m, deterministic)
}
func (m *ConfigItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigItem.Merge(m, src)
}
func (m *ConfigItem) XXX_Size() int {
	return xxx_messageInfo_ConfigItem.Size(m)
}
func (m *ConfigItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigItem.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigItem proto.InternalMessageInfo

func (m *ConfigItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigItem) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type isConfigItem_Value interface {
	isConfigItem_Value()
}

type ConfigItem_Arr struct {
	Arr *ArrayConfig `protobuf:"bytes,3,opt,name=arr,proto3,oneof"`
}

type ConfigItem_Str struct {
	Str *StringConfig `protobuf:"bytes,4,opt,name=str,proto3,oneof"`
}

type ConfigItem_Int struct {
	Int *Int32Config `protobuf:"bytes,5,opt,name=int,proto3,oneof"`
}

func (*ConfigItem_Arr) isConfigItem_Value() {}

func (*ConfigItem_Str) isConfigItem_Value() {}

func (*ConfigItem_Int) isConfigItem_Value() {}

func (m *ConfigItem) GetValue() isConfigItem_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ConfigItem) GetArr() *ArrayConfig {
	if x, ok := m.GetValue().(*ConfigItem_Arr); ok {
		return x.Arr
	}
	return nil
}

func (m *ConfigItem) GetStr() *StringConfig {
	if x, ok := m.GetValue().(*ConfigItem_Str); ok {
		return x.Str
	}
	return nil
}

func (m *ConfigItem) GetInt() *Int32Config {
	if x, ok := m.GetValue().(*ConfigItem_Int); ok {
		return x.Int
	}
	return nil
}

func (m *ConfigItem) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigItem_OneofMarshaler, _ConfigItem_OneofUnmarshaler, _ConfigItem_OneofSizer, []interface{}{
		(*ConfigItem_Arr)(nil),
		(*ConfigItem_Str)(nil),
		(*ConfigItem_Int)(nil),
	}
}

func _ConfigItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigItem)
	// value
	switch x := m.Value.(type) {
	case *ConfigItem_Arr:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Arr); err != nil {
			return err
		}
	case *ConfigItem_Str:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Str); err != nil {
			return err
		}
	case *ConfigItem_Int:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Int); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConfigItem.Value has unexpected type %T", x)
	}
	return nil
}

func _ConfigItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigItem)
	switch tag {
	case 3: // value.arr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ArrayConfig)
		err := b.DecodeMessage(msg)
		m.Value = &ConfigItem_Arr{msg}
		return true, err
	case 4: // value.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringConfig)
		err := b.DecodeMessage(msg)
		m.Value = &ConfigItem_Str{msg}
		return true, err
	case 5: // value.int
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Int32Config)
		err := b.DecodeMessage(msg)
		m.Value = &ConfigItem_Int{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigItem)
	// value
	switch x := m.Value.(type) {
	case *ConfigItem_Arr:
		s := proto.Size(x.Arr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigItem_Str:
		s := proto.Size(x.Str)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigItem_Int:
		s := proto.Size(x.Int)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ModifyConfig struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Op                   string   `protobuf:"bytes,3,opt,name=op,proto3" json:"op,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyConfig) Reset()         { *m = ModifyConfig{} }
func (m *ModifyConfig) String() string { return proto.CompactTextString(m) }
func (*ModifyConfig) ProtoMessage()    {}
func (*ModifyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{8}
}

func (m *ModifyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyConfig.Unmarshal(m, b)
}
func (m *ModifyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyConfig.Marshal(b, m, deterministic)
}
func (m *ModifyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyConfig.Merge(m, src)
}
func (m *ModifyConfig) XXX_Size() int {
	return xxx_messageInfo_ModifyConfig.Size(m)
}
func (m *ModifyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyConfig proto.InternalMessageInfo

func (m *ModifyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ModifyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ModifyConfig) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *ModifyConfig) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ReceiptConfig struct {
	Prev                 *ConfigItem `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Current              *ConfigItem `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReceiptConfig) Reset()         { *m = ReceiptConfig{} }
func (m *ReceiptConfig) String() string { return proto.CompactTextString(m) }
func (*ReceiptConfig) ProtoMessage()    {}
func (*ReceiptConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{9}
}

func (m *ReceiptConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptConfig.Unmarshal(m, b)
}
func (m *ReceiptConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptConfig.Marshal(b, m, deterministic)
}
func (m *ReceiptConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptConfig.Merge(m, src)
}
func (m *ReceiptConfig) XXX_Size() int {
	return xxx_messageInfo_ReceiptConfig.Size(m)
}
func (m *ReceiptConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptConfig proto.InternalMessageInfo

func (m *ReceiptConfig) GetPrev() *ConfigItem {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptConfig) GetCurrent() *ConfigItem {
	if m != nil {
		return m.Current
	}
	return nil
}

type ReplyConfig struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyConfig) Reset()         { *m = ReplyConfig{} }
func (m *ReplyConfig) String() string { return proto.CompactTextString(m) }
func (*ReplyConfig) ProtoMessage()    {}
func (*ReplyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{10}
}

func (m *ReplyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyConfig.Unmarshal(m, b)
}
func (m *ReplyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyConfig.Marshal(b, m, deterministic)
}
func (m *ReplyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyConfig.Merge(m, src)
}
func (m *ReplyConfig) XXX_Size() int {
	return xxx_messageInfo_ReplyConfig.Size(m)
}
func (m *ReplyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyConfig proto.InternalMessageInfo

func (m *ReplyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HistoryCertStore struct {
	Rootcerts            [][]byte `protobuf:"bytes,1,rep,name=rootcerts,proto3" json:"rootcerts,omitempty"`
	IntermediateCerts    [][]byte `protobuf:"bytes,2,rep,name=intermediateCerts,proto3" json:"intermediateCerts,omitempty"`
	RevocationList       [][]byte `protobuf:"bytes,3,rep,name=revocationList,proto3" json:"revocationList,omitempty"`
	CurHeigth            int64    `protobuf:"varint,4,opt,name=curHeigth,proto3" json:"curHeigth,omitempty"`
	NxtHeight            int64    `protobuf:"varint,5,opt,name=nxtHeight,proto3" json:"nxtHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryCertStore) Reset()         { *m = HistoryCertStore{} }
func (m *HistoryCertStore) String() string { return proto.CompactTextString(m) }
func (*HistoryCertStore) ProtoMessage()    {}
func (*HistoryCertStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{11}
}

func (m *HistoryCertStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryCertStore.Unmarshal(m, b)
}
func (m *HistoryCertStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryCertStore.Marshal(b, m, deterministic)
}
func (m *HistoryCertStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryCertStore.Merge(m, src)
}
func (m *HistoryCertStore) XXX_Size() int {
	return xxx_messageInfo_HistoryCertStore.Size(m)
}
func (m *HistoryCertStore) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryCertStore.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryCertStore proto.InternalMessageInfo

func (m *HistoryCertStore) GetRootcerts() [][]byte {
	if m != nil {
		return m.Rootcerts
	}
	return nil
}

func (m *HistoryCertStore) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *HistoryCertStore) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *HistoryCertStore) GetCurHeigth() int64 {
	if m != nil {
		return m.CurHeigth
	}
	return 0
}

func (m *HistoryCertStore) GetNxtHeight() int64 {
	if m != nil {
		return m.NxtHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Genesis)(nil), "types.Genesis")
	proto.RegisterType((*ExecTxList)(nil), "types.ExecTxList")
	proto.RegisterType((*Query)(nil), "types.Query")
	proto.RegisterType((*CreateTxIn)(nil), "types.CreateTxIn")
	proto.RegisterType((*ArrayConfig)(nil), "types.ArrayConfig")
	proto.RegisterType((*StringConfig)(nil), "types.StringConfig")
	proto.RegisterType((*Int32Config)(nil), "types.Int32Config")
	proto.RegisterType((*ConfigItem)(nil), "types.ConfigItem")
	proto.RegisterType((*ModifyConfig)(nil), "types.ModifyConfig")
	proto.RegisterType((*ReceiptConfig)(nil), "types.ReceiptConfig")
	proto.RegisterType((*ReplyConfig)(nil), "types.ReplyConfig")
	proto.RegisterType((*HistoryCertStore)(nil), "types.HistoryCertStore")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f) }

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xed, 0x6a, 0xdb, 0x4a,
	0x10, 0xbd, 0x92, 0xec, 0x38, 0x1e, 0xfb, 0x86, 0x64, 0xef, 0xa5, 0x88, 0xd0, 0x26, 0x46, 0x49,
	0x53, 0x43, 0x8b, 0x03, 0x36, 0x7d, 0x80, 0xc6, 0x94, 0x26, 0xd0, 0x14, 0xaa, 0xb8, 0x7f, 0xf2,
	0xa3, 0xb0, 0x59, 0x8f, 0xed, 0x25, 0xf6, 0xae, 0x58, 0x8d, 0x82, 0xf5, 0x36, 0x7d, 0x96, 0xd2,
	0x07, 0x2b, 0xbb, 0x92, 0x2d, 0xd1, 0xa4, 0x85, 0xfe, 0xd3, 0x9c, 0x73, 0x74, 0x76, 0xce, 0xec,
	0x07, 0xec, 0xe1, 0x1a, 0x45, 0x46, 0xda, 0x0c, 0x12, 0xa3, 0x49, 0xb3, 0x26, 0xe5, 0x09, 0xa6,
	0x87, 0x07, 0x64, 0xb8, 0x4a, 0xb9, 0x20, 0xa9, 0x55, 0xc1, 0x44, 0xc7, 0xd0, 0xfa, 0x80, 0x0a,
	0x53, 0x99, 0xb2, 0xff, 0xa1, 0x29, 0x53, 0x93, 0xa9, 0xd0, 0xeb, 0x79, 0xfd, 0xdd, 0xb8, 0x28,
	0xa2, 0x6f, 0x3e, 0xc0, 0xfb, 0x35, 0x8a, 0xc9, 0xfa, 0xa3, 0x4c, 0x89, 0x3d, 0x87, 0x76, 0x4a,
	0x9c, 0xf0, 0x92, 0xa7, 0x0b, 0x27, 0xec, 0xc6, 0x15, 0xc0, 0x8e, 0x00, 0x12, 0x6e, 0x50, 0x91,
	0xa3, 0x5b, 0x8e, 0xae, 0x21, 0xec, 0x10, 0x76, 0x57, 0x5c, 0x2a, 0xc7, 0xee, 0x3a, 0x76, 0x5b,
	0xdb, 0x7f, 0xdd, 0x37, 0xca, 0xf9, 0x82, 0xc2, 0x76, 0xcf, 0xeb, 0x07, 0x71, 0x0d, 0xb1, 0x2b,
	0xdf, 0x2d, 0xb5, 0xb8, 0x9f, 0xc8, 0x15, 0x86, 0x81, 0xa3, 0x2b, 0x80, 0x3d, 0x83, 0x9d, 0x45,
	0xf1, 0x67, 0xc3, 0x51, 0x65, 0x65, 0x5d, 0xa7, 0x72, 0x36, 0x93, 0x22, 0x5b, 0x52, 0x1e, 0x36,
	0x7b, 0x5e, 0xbf, 0x11, 0xd7, 0x10, 0xeb, 0x2a, 0xd3, 0x6b, 0x5c, 0x25, 0x5a, 0x2f, 0xc3, 0x1d,
	0x17, 0xbc, 0x02, 0xd8, 0x29, 0x04, 0xb4, 0x4e, 0x43, 0xbf, 0x17, 0xf4, 0x3b, 0x43, 0x36, 0x70,
	0x53, 0x1c, 0x4c, 0xaa, 0x21, 0xc6, 0x96, 0x8e, 0xbe, 0x40, 0xf3, 0x73, 0x86, 0x26, 0xb7, 0x4d,
	0xd8, 0xc1, 0xa3, 0x29, 0x27, 0x53, 0x56, 0x36, 0xf6, 0x2c, 0x53, 0xe2, 0x13, 0x5f, 0x61, 0xe8,
	0xf7, 0xbc, 0x7e, 0x3b, 0xde, 0xd6, 0x2c, 0x84, 0x56, 0xc2, 0xf3, 0xa5, 0xe6, 0x53, 0x17, 0xaa,
	0x1b, 0x6f, 0xca, 0xe8, 0x2b, 0xc0, 0xd8, 0x20, 0x27, 0x9c, 0xac, 0xaf, 0xd4, 0x6f, 0xbd, 0x8f,
	0x00, 0x8a, 0x5e, 0x6a, 0xee, 0x35, 0xe4, 0x0f, 0xfe, 0x27, 0xd0, 0x79, 0x67, 0x0c, 0xcf, 0xc7,
	0x5a, 0xcd, 0xe4, 0xdc, 0x6e, 0xff, 0x03, 0x5f, 0x66, 0x76, 0xb6, 0x41, 0xbf, 0x1d, 0x17, 0x45,
	0x74, 0x0a, 0xdd, 0x1b, 0x32, 0x52, 0xcd, 0x1f, 0xab, 0xbc, 0x4a, 0x75, 0x02, 0x9d, 0x2b, 0x45,
	0xa3, 0xe1, 0x53, 0xa2, 0xe6, 0x46, 0xf4, 0xc3, 0x03, 0x28, 0x04, 0x57, 0x84, 0x2b, 0xb6, 0x0f,
	0xc1, 0x3d, 0xe6, 0x2e, 0x4d, 0x3b, 0xb6, 0x9f, 0x8c, 0x41, 0x83, 0x4f, 0xa7, 0xa6, 0x0c, 0xe1,
	0xbe, 0xd9, 0x19, 0x04, 0xdc, 0x18, 0x67, 0x54, 0xed, 0x40, 0xad, 0xed, 0xcb, 0x7f, 0x62, 0x2b,
	0x60, 0xaf, 0x20, 0x48, 0xc9, 0xb8, 0xcd, 0xef, 0x0c, 0xff, 0x2b, 0x75, 0xf5, 0xce, 0xad, 0x30,
	0x25, 0x67, 0x28, 0x15, 0xb9, 0x93, 0x50, 0x19, 0xd6, 0x9a, 0xb7, 0x3a, 0xa9, 0x88, 0xed, 0x81,
	0x3f, 0xc9, 0xc3, 0x8e, 0x0b, 0xe0, 0x4f, 0xf2, 0x8b, 0x56, 0x99, 0x29, 0xba, 0x85, 0xee, 0xb5,
	0x9e, 0xca, 0xd9, 0x66, 0x6e, 0x8f, 0x73, 0x6c, 0xe3, 0xfb, 0xb5, 0x19, 0x59, 0x43, 0x9d, 0x94,
	0x63, 0xf3, 0x75, 0xb2, 0x4d, 0xdb, 0xa8, 0xd2, 0x46, 0x02, 0xfe, 0x8d, 0x51, 0xa0, 0x4c, 0xa8,
	0x34, 0x7f, 0x09, 0x8d, 0xc4, 0xe0, 0x83, 0x73, 0xef, 0x0c, 0x0f, 0xca, 0x76, 0xab, 0x29, 0xc6,
	0x8e, 0x66, 0xaf, 0xa1, 0x25, 0x32, 0x63, 0xaf, 0x99, 0x5b, 0xf3, 0x49, 0xe5, 0x46, 0x11, 0xbd,
	0x85, 0x4e, 0x8c, 0xc9, 0xf2, 0x2f, 0xfb, 0x8f, 0xbe, 0x7b, 0xb0, 0x7f, 0x29, 0x53, 0xd2, 0x26,
	0x1f, 0xa3, 0xa1, 0x1b, 0xd2, 0x06, 0xed, 0xf5, 0x31, 0x5a, 0x93, 0x40, 0x43, 0x69, 0xe8, 0xf5,
	0x02, 0xfb, 0x1c, 0x6c, 0x01, 0xf6, 0x06, 0x0e, 0xa4, 0x22, 0x34, 0x2b, 0x9c, 0x4a, 0x4e, 0x38,
	0x76, 0x2a, 0xdf, 0xa9, 0x1e, 0x13, 0xec, 0x0c, 0xf6, 0x0c, 0x3e, 0x68, 0xc1, 0xed, 0xd9, 0xb5,
	0x8f, 0x8d, 0x3b, 0x89, 0xdd, 0xf8, 0x17, 0xd4, 0xae, 0x29, 0x32, 0x63, 0x5f, 0x05, 0x5a, 0x94,
	0xb7, 0xbd, 0x02, 0x2c, 0xab, 0xd6, 0x54, 0xbe, 0x22, 0xcd, 0x82, 0xdd, 0x02, 0x17, 0xc7, 0xb7,
	0x2f, 0xe6, 0x92, 0x16, 0xd9, 0xdd, 0x40, 0xe8, 0xd5, 0xf9, 0x68, 0x24, 0xd4, 0xb9, 0x58, 0x70,
	0xa9, 0x46, 0xa3, 0x73, 0x37, 0xb0, 0xbb, 0x1d, 0xf7, 0x2c, 0x8e, 0x7e, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x22, 0x46, 0xb1, 0x62, 0x42, 0x05, 0x00, 0x00,
}