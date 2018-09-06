// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lottery.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PurchaseRecord struct {
	Amount int64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	Number int64 `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
}

func (m *PurchaseRecord) Reset()                    { *m = PurchaseRecord{} }
func (m *PurchaseRecord) String() string            { return proto.CompactTextString(m) }
func (*PurchaseRecord) ProtoMessage()               {}
func (*PurchaseRecord) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *PurchaseRecord) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PurchaseRecord) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type PurchaseRecords struct {
	Record         []*PurchaseRecord `protobuf:"bytes,1,rep,name=record" json:"record,omitempty"`
	FundWin        int64             `protobuf:"varint,2,opt,name=fundWin" json:"fundWin,omitempty"`
	AmountOneRound int64             `protobuf:"varint,3,opt,name=amountOneRound" json:"amountOneRound,omitempty"`
}

func (m *PurchaseRecords) Reset()                    { *m = PurchaseRecords{} }
func (m *PurchaseRecords) String() string            { return proto.CompactTextString(m) }
func (*PurchaseRecords) ProtoMessage()               {}
func (*PurchaseRecords) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *PurchaseRecords) GetRecord() []*PurchaseRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *PurchaseRecords) GetFundWin() int64 {
	if m != nil {
		return m.FundWin
	}
	return 0
}

func (m *PurchaseRecords) GetAmountOneRound() int64 {
	if m != nil {
		return m.AmountOneRound
	}
	return 0
}

type Lottery struct {
	LotteryId            string                      `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
	Status               int32                       `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	CreateHeight         int64                       `protobuf:"varint,3,opt,name=createHeight" json:"createHeight,omitempty"`
	Fund                 int64                       `protobuf:"varint,4,opt,name=fund" json:"fund,omitempty"`
	PurBlockNum          int64                       `protobuf:"varint,5,opt,name=purBlockNum" json:"purBlockNum,omitempty"`
	DrawBlockNum         int64                       `protobuf:"varint,6,opt,name=drawBlockNum" json:"drawBlockNum,omitempty"`
	LastTransToPurState  int64                       `protobuf:"varint,7,opt,name=lastTransToPurState" json:"lastTransToPurState,omitempty"`
	LastTransToDrawState int64                       `protobuf:"varint,8,opt,name=lastTransToDrawState" json:"lastTransToDrawState,omitempty"`
	Records              map[string]*PurchaseRecords `protobuf:"bytes,9,rep,name=records" json:"records,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TotalPurchasedTxNum  int64                       `protobuf:"varint,10,opt,name=totalPurchasedTxNum" json:"totalPurchasedTxNum,omitempty"`
	CreateAddr           string                      `protobuf:"bytes,11,opt,name=createAddr" json:"createAddr,omitempty"`
	Round                int64                       `protobuf:"varint,12,opt,name=round" json:"round,omitempty"`
	LuckyNumber          int64                       `protobuf:"varint,13,opt,name=luckyNumber" json:"luckyNumber,omitempty"`
}

func (m *Lottery) Reset()                    { *m = Lottery{} }
func (m *Lottery) String() string            { return proto.CompactTextString(m) }
func (*Lottery) ProtoMessage()               {}
func (*Lottery) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *Lottery) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

func (m *Lottery) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Lottery) GetCreateHeight() int64 {
	if m != nil {
		return m.CreateHeight
	}
	return 0
}

func (m *Lottery) GetFund() int64 {
	if m != nil {
		return m.Fund
	}
	return 0
}

func (m *Lottery) GetPurBlockNum() int64 {
	if m != nil {
		return m.PurBlockNum
	}
	return 0
}

func (m *Lottery) GetDrawBlockNum() int64 {
	if m != nil {
		return m.DrawBlockNum
	}
	return 0
}

func (m *Lottery) GetLastTransToPurState() int64 {
	if m != nil {
		return m.LastTransToPurState
	}
	return 0
}

func (m *Lottery) GetLastTransToDrawState() int64 {
	if m != nil {
		return m.LastTransToDrawState
	}
	return 0
}

func (m *Lottery) GetRecords() map[string]*PurchaseRecords {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *Lottery) GetTotalPurchasedTxNum() int64 {
	if m != nil {
		return m.TotalPurchasedTxNum
	}
	return 0
}

func (m *Lottery) GetCreateAddr() string {
	if m != nil {
		return m.CreateAddr
	}
	return ""
}

func (m *Lottery) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Lottery) GetLuckyNumber() int64 {
	if m != nil {
		return m.LuckyNumber
	}
	return 0
}

// message for execs.game
type LotteryAction struct {
	// Types that are valid to be assigned to Value:
	//	*LotteryAction_Create
	//	*LotteryAction_Buy
	//	*LotteryAction_Draw
	//	*LotteryAction_Close
	Value isLotteryAction_Value `protobuf_oneof:"value"`
	Ty    int32                 `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
}

func (m *LotteryAction) Reset()                    { *m = LotteryAction{} }
func (m *LotteryAction) String() string            { return proto.CompactTextString(m) }
func (*LotteryAction) ProtoMessage()               {}
func (*LotteryAction) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

type isLotteryAction_Value interface {
	isLotteryAction_Value()
}

type LotteryAction_Create struct {
	Create *LotteryCreate `protobuf:"bytes,1,opt,name=create,oneof"`
}
type LotteryAction_Buy struct {
	Buy *LotteryBuy `protobuf:"bytes,2,opt,name=buy,oneof"`
}
type LotteryAction_Draw struct {
	Draw *LotteryDraw `protobuf:"bytes,3,opt,name=draw,oneof"`
}
type LotteryAction_Close struct {
	Close *LotteryClose `protobuf:"bytes,4,opt,name=close,oneof"`
}

func (*LotteryAction_Create) isLotteryAction_Value() {}
func (*LotteryAction_Buy) isLotteryAction_Value()    {}
func (*LotteryAction_Draw) isLotteryAction_Value()   {}
func (*LotteryAction_Close) isLotteryAction_Value()  {}

func (m *LotteryAction) GetValue() isLotteryAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LotteryAction) GetCreate() *LotteryCreate {
	if x, ok := m.GetValue().(*LotteryAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *LotteryAction) GetBuy() *LotteryBuy {
	if x, ok := m.GetValue().(*LotteryAction_Buy); ok {
		return x.Buy
	}
	return nil
}

func (m *LotteryAction) GetDraw() *LotteryDraw {
	if x, ok := m.GetValue().(*LotteryAction_Draw); ok {
		return x.Draw
	}
	return nil
}

func (m *LotteryAction) GetClose() *LotteryClose {
	if x, ok := m.GetValue().(*LotteryAction_Close); ok {
		return x.Close
	}
	return nil
}

func (m *LotteryAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LotteryAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LotteryAction_OneofMarshaler, _LotteryAction_OneofUnmarshaler, _LotteryAction_OneofSizer, []interface{}{
		(*LotteryAction_Create)(nil),
		(*LotteryAction_Buy)(nil),
		(*LotteryAction_Draw)(nil),
		(*LotteryAction_Close)(nil),
	}
}

func _LotteryAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LotteryAction)
	// value
	switch x := m.Value.(type) {
	case *LotteryAction_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *LotteryAction_Buy:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Buy); err != nil {
			return err
		}
	case *LotteryAction_Draw:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Draw); err != nil {
			return err
		}
	case *LotteryAction_Close:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Close); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LotteryAction.Value has unexpected type %T", x)
	}
	return nil
}

func _LotteryAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LotteryAction)
	switch tag {
	case 1: // value.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LotteryCreate)
		err := b.DecodeMessage(msg)
		m.Value = &LotteryAction_Create{msg}
		return true, err
	case 2: // value.buy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LotteryBuy)
		err := b.DecodeMessage(msg)
		m.Value = &LotteryAction_Buy{msg}
		return true, err
	case 3: // value.draw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LotteryDraw)
		err := b.DecodeMessage(msg)
		m.Value = &LotteryAction_Draw{msg}
		return true, err
	case 4: // value.close
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LotteryClose)
		err := b.DecodeMessage(msg)
		m.Value = &LotteryAction_Close{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LotteryAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LotteryAction)
	// value
	switch x := m.Value.(type) {
	case *LotteryAction_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LotteryAction_Buy:
		s := proto.Size(x.Buy)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LotteryAction_Draw:
		s := proto.Size(x.Draw)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LotteryAction_Close:
		s := proto.Size(x.Close)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LotteryCreate struct {
	PurBlockNum  int64 `protobuf:"varint,1,opt,name=purBlockNum" json:"purBlockNum,omitempty"`
	DrawBlockNum int64 `protobuf:"varint,2,opt,name=drawBlockNum" json:"drawBlockNum,omitempty"`
}

func (m *LotteryCreate) Reset()                    { *m = LotteryCreate{} }
func (m *LotteryCreate) String() string            { return proto.CompactTextString(m) }
func (*LotteryCreate) ProtoMessage()               {}
func (*LotteryCreate) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *LotteryCreate) GetPurBlockNum() int64 {
	if m != nil {
		return m.PurBlockNum
	}
	return 0
}

func (m *LotteryCreate) GetDrawBlockNum() int64 {
	if m != nil {
		return m.DrawBlockNum
	}
	return 0
}

type LotteryBuy struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
	Amount    int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Number    int64  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
}

func (m *LotteryBuy) Reset()                    { *m = LotteryBuy{} }
func (m *LotteryBuy) String() string            { return proto.CompactTextString(m) }
func (*LotteryBuy) ProtoMessage()               {}
func (*LotteryBuy) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *LotteryBuy) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

func (m *LotteryBuy) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *LotteryBuy) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type LotteryDraw struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
}

func (m *LotteryDraw) Reset()                    { *m = LotteryDraw{} }
func (m *LotteryDraw) String() string            { return proto.CompactTextString(m) }
func (*LotteryDraw) ProtoMessage()               {}
func (*LotteryDraw) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *LotteryDraw) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

type LotteryClose struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
}

func (m *LotteryClose) Reset()                    { *m = LotteryClose{} }
func (m *LotteryClose) String() string            { return proto.CompactTextString(m) }
func (*LotteryClose) ProtoMessage()               {}
func (*LotteryClose) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *LotteryClose) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

type ReceiptLottery struct {
	LotteryId   string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
	Status      int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	PrevStatus  int32  `protobuf:"varint,3,opt,name=prevStatus" json:"prevStatus,omitempty"`
	Addr        string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
	Round       int64  `protobuf:"varint,5,opt,name=round" json:"round,omitempty"`
	Number      int64  `protobuf:"varint,6,opt,name=number" json:"number,omitempty"`
	Amount      int64  `protobuf:"varint,7,opt,name=amount" json:"amount,omitempty"`
	LuckyNumber int64  `protobuf:"varint,8,opt,name=luckyNumber" json:"luckyNumber,omitempty"`
}

func (m *ReceiptLottery) Reset()                    { *m = ReceiptLottery{} }
func (m *ReceiptLottery) String() string            { return proto.CompactTextString(m) }
func (*ReceiptLottery) ProtoMessage()               {}
func (*ReceiptLottery) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *ReceiptLottery) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

func (m *ReceiptLottery) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReceiptLottery) GetPrevStatus() int32 {
	if m != nil {
		return m.PrevStatus
	}
	return 0
}

func (m *ReceiptLottery) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReceiptLottery) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *ReceiptLottery) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReceiptLottery) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReceiptLottery) GetLuckyNumber() int64 {
	if m != nil {
		return m.LuckyNumber
	}
	return 0
}

type ReqLotteryInfo struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
}

func (m *ReqLotteryInfo) Reset()                    { *m = ReqLotteryInfo{} }
func (m *ReqLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*ReqLotteryInfo) ProtoMessage()               {}
func (*ReqLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *ReqLotteryInfo) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

type ReqLotteryBuyInfo struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
	Addr      string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	Round     int64  `protobuf:"varint,3,opt,name=round" json:"round,omitempty"`
}

func (m *ReqLotteryBuyInfo) Reset()                    { *m = ReqLotteryBuyInfo{} }
func (m *ReqLotteryBuyInfo) String() string            { return proto.CompactTextString(m) }
func (*ReqLotteryBuyInfo) ProtoMessage()               {}
func (*ReqLotteryBuyInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *ReqLotteryBuyInfo) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

func (m *ReqLotteryBuyInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqLotteryBuyInfo) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type ReqLotteryBuyHistory struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
	Addr      string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	Round     int64  `protobuf:"varint,3,opt,name=round" json:"round,omitempty"`
	Count     int32  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,5,opt,name=direction" json:"direction,omitempty"`
}

func (m *ReqLotteryBuyHistory) Reset()                    { *m = ReqLotteryBuyHistory{} }
func (m *ReqLotteryBuyHistory) String() string            { return proto.CompactTextString(m) }
func (*ReqLotteryBuyHistory) ProtoMessage()               {}
func (*ReqLotteryBuyHistory) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{11} }

func (m *ReqLotteryBuyHistory) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

func (m *ReqLotteryBuyHistory) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqLotteryBuyHistory) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *ReqLotteryBuyHistory) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqLotteryBuyHistory) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ReqLotteryLuckyInfo struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
	Round     int64  `protobuf:"varint,2,opt,name=round" json:"round,omitempty"`
}

func (m *ReqLotteryLuckyInfo) Reset()                    { *m = ReqLotteryLuckyInfo{} }
func (m *ReqLotteryLuckyInfo) String() string            { return proto.CompactTextString(m) }
func (*ReqLotteryLuckyInfo) ProtoMessage()               {}
func (*ReqLotteryLuckyInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{12} }

func (m *ReqLotteryLuckyInfo) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

func (m *ReqLotteryLuckyInfo) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type ReqLotteryLuckyHistory struct {
	LotteryId string `protobuf:"bytes,1,opt,name=lotteryId" json:"lotteryId,omitempty"`
	Round     int64  `protobuf:"varint,2,opt,name=round" json:"round,omitempty"`
	Count     int32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,4,opt,name=direction" json:"direction,omitempty"`
}

func (m *ReqLotteryLuckyHistory) Reset()                    { *m = ReqLotteryLuckyHistory{} }
func (m *ReqLotteryLuckyHistory) String() string            { return proto.CompactTextString(m) }
func (*ReqLotteryLuckyHistory) ProtoMessage()               {}
func (*ReqLotteryLuckyHistory) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{13} }

func (m *ReqLotteryLuckyHistory) GetLotteryId() string {
	if m != nil {
		return m.LotteryId
	}
	return ""
}

func (m *ReqLotteryLuckyHistory) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *ReqLotteryLuckyHistory) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqLotteryLuckyHistory) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ReplyLotteryNormalInfo struct {
	CreateHeight int64  `protobuf:"varint,1,opt,name=createHeight" json:"createHeight,omitempty"`
	PurBlockNum  int64  `protobuf:"varint,2,opt,name=purBlockNum" json:"purBlockNum,omitempty"`
	DrawBlockNum int64  `protobuf:"varint,3,opt,name=drawBlockNum" json:"drawBlockNum,omitempty"`
	CreateAddr   string `protobuf:"bytes,4,opt,name=createAddr" json:"createAddr,omitempty"`
}

func (m *ReplyLotteryNormalInfo) Reset()                    { *m = ReplyLotteryNormalInfo{} }
func (m *ReplyLotteryNormalInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyLotteryNormalInfo) ProtoMessage()               {}
func (*ReplyLotteryNormalInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{14} }

func (m *ReplyLotteryNormalInfo) GetCreateHeight() int64 {
	if m != nil {
		return m.CreateHeight
	}
	return 0
}

func (m *ReplyLotteryNormalInfo) GetPurBlockNum() int64 {
	if m != nil {
		return m.PurBlockNum
	}
	return 0
}

func (m *ReplyLotteryNormalInfo) GetDrawBlockNum() int64 {
	if m != nil {
		return m.DrawBlockNum
	}
	return 0
}

func (m *ReplyLotteryNormalInfo) GetCreateAddr() string {
	if m != nil {
		return m.CreateAddr
	}
	return ""
}

type ReplyLotteryCurrentInfo struct {
	Status               int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Fund                 int64 `protobuf:"varint,2,opt,name=fund" json:"fund,omitempty"`
	LastTransToPurState  int64 `protobuf:"varint,3,opt,name=lastTransToPurState" json:"lastTransToPurState,omitempty"`
	LastTransToDrawState int64 `protobuf:"varint,4,opt,name=lastTransToDrawState" json:"lastTransToDrawState,omitempty"`
	TotalPurchasedTxNum  int64 `protobuf:"varint,5,opt,name=totalPurchasedTxNum" json:"totalPurchasedTxNum,omitempty"`
	Round                int64 `protobuf:"varint,6,opt,name=round" json:"round,omitempty"`
	LuckyNumber          int64 `protobuf:"varint,7,opt,name=luckyNumber" json:"luckyNumber,omitempty"`
}

func (m *ReplyLotteryCurrentInfo) Reset()                    { *m = ReplyLotteryCurrentInfo{} }
func (m *ReplyLotteryCurrentInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyLotteryCurrentInfo) ProtoMessage()               {}
func (*ReplyLotteryCurrentInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{15} }

func (m *ReplyLotteryCurrentInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReplyLotteryCurrentInfo) GetFund() int64 {
	if m != nil {
		return m.Fund
	}
	return 0
}

func (m *ReplyLotteryCurrentInfo) GetLastTransToPurState() int64 {
	if m != nil {
		return m.LastTransToPurState
	}
	return 0
}

func (m *ReplyLotteryCurrentInfo) GetLastTransToDrawState() int64 {
	if m != nil {
		return m.LastTransToDrawState
	}
	return 0
}

func (m *ReplyLotteryCurrentInfo) GetTotalPurchasedTxNum() int64 {
	if m != nil {
		return m.TotalPurchasedTxNum
	}
	return 0
}

func (m *ReplyLotteryCurrentInfo) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *ReplyLotteryCurrentInfo) GetLuckyNumber() int64 {
	if m != nil {
		return m.LuckyNumber
	}
	return 0
}

type ReplyLotteryHistoryLuckyNumber struct {
	LuckyNumber []int64 `protobuf:"varint,1,rep,packed,name=luckyNumber" json:"luckyNumber,omitempty"`
}

func (m *ReplyLotteryHistoryLuckyNumber) Reset()                    { *m = ReplyLotteryHistoryLuckyNumber{} }
func (m *ReplyLotteryHistoryLuckyNumber) String() string            { return proto.CompactTextString(m) }
func (*ReplyLotteryHistoryLuckyNumber) ProtoMessage()               {}
func (*ReplyLotteryHistoryLuckyNumber) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{16} }

func (m *ReplyLotteryHistoryLuckyNumber) GetLuckyNumber() []int64 {
	if m != nil {
		return m.LuckyNumber
	}
	return nil
}

type ReplyLotteryShowInfo struct {
	Records []*LotteryBuyRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *ReplyLotteryShowInfo) Reset()                    { *m = ReplyLotteryShowInfo{} }
func (m *ReplyLotteryShowInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyLotteryShowInfo) ProtoMessage()               {}
func (*ReplyLotteryShowInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{17} }

func (m *ReplyLotteryShowInfo) GetRecords() []*LotteryBuyRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type LotteryNumberRecord struct {
	Number int64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Amount int64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *LotteryNumberRecord) Reset()                    { *m = LotteryNumberRecord{} }
func (m *LotteryNumberRecord) String() string            { return proto.CompactTextString(m) }
func (*LotteryNumberRecord) ProtoMessage()               {}
func (*LotteryNumberRecord) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{18} }

func (m *LotteryNumberRecord) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *LotteryNumberRecord) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// used for execlocal
type LotteryBuyRecord struct {
	Number int64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Amount int64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Round  int64 `protobuf:"varint,3,opt,name=round" json:"round,omitempty"`
}

func (m *LotteryBuyRecord) Reset()                    { *m = LotteryBuyRecord{} }
func (m *LotteryBuyRecord) String() string            { return proto.CompactTextString(m) }
func (*LotteryBuyRecord) ProtoMessage()               {}
func (*LotteryBuyRecord) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{19} }

func (m *LotteryBuyRecord) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *LotteryBuyRecord) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *LotteryBuyRecord) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type LotteryBuyRecords struct {
	Records []*LotteryBuyRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *LotteryBuyRecords) Reset()                    { *m = LotteryBuyRecords{} }
func (m *LotteryBuyRecords) String() string            { return proto.CompactTextString(m) }
func (*LotteryBuyRecords) ProtoMessage()               {}
func (*LotteryBuyRecords) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{20} }

func (m *LotteryBuyRecords) GetRecords() []*LotteryBuyRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type LotteryDrawRecord struct {
	Number int64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Round  int64 `protobuf:"varint,2,opt,name=round" json:"round,omitempty"`
}

func (m *LotteryDrawRecord) Reset()                    { *m = LotteryDrawRecord{} }
func (m *LotteryDrawRecord) String() string            { return proto.CompactTextString(m) }
func (*LotteryDrawRecord) ProtoMessage()               {}
func (*LotteryDrawRecord) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{21} }

func (m *LotteryDrawRecord) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *LotteryDrawRecord) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type LotteryDrawRecords struct {
	Records []*LotteryDrawRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *LotteryDrawRecords) Reset()                    { *m = LotteryDrawRecords{} }
func (m *LotteryDrawRecords) String() string            { return proto.CompactTextString(m) }
func (*LotteryDrawRecords) ProtoMessage()               {}
func (*LotteryDrawRecords) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{22} }

func (m *LotteryDrawRecords) GetRecords() []*LotteryDrawRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*PurchaseRecord)(nil), "types.PurchaseRecord")
	proto.RegisterType((*PurchaseRecords)(nil), "types.PurchaseRecords")
	proto.RegisterType((*Lottery)(nil), "types.Lottery")
	proto.RegisterType((*LotteryAction)(nil), "types.LotteryAction")
	proto.RegisterType((*LotteryCreate)(nil), "types.LotteryCreate")
	proto.RegisterType((*LotteryBuy)(nil), "types.LotteryBuy")
	proto.RegisterType((*LotteryDraw)(nil), "types.LotteryDraw")
	proto.RegisterType((*LotteryClose)(nil), "types.LotteryClose")
	proto.RegisterType((*ReceiptLottery)(nil), "types.ReceiptLottery")
	proto.RegisterType((*ReqLotteryInfo)(nil), "types.ReqLotteryInfo")
	proto.RegisterType((*ReqLotteryBuyInfo)(nil), "types.ReqLotteryBuyInfo")
	proto.RegisterType((*ReqLotteryBuyHistory)(nil), "types.ReqLotteryBuyHistory")
	proto.RegisterType((*ReqLotteryLuckyInfo)(nil), "types.ReqLotteryLuckyInfo")
	proto.RegisterType((*ReqLotteryLuckyHistory)(nil), "types.ReqLotteryLuckyHistory")
	proto.RegisterType((*ReplyLotteryNormalInfo)(nil), "types.ReplyLotteryNormalInfo")
	proto.RegisterType((*ReplyLotteryCurrentInfo)(nil), "types.ReplyLotteryCurrentInfo")
	proto.RegisterType((*ReplyLotteryHistoryLuckyNumber)(nil), "types.ReplyLotteryHistoryLuckyNumber")
	proto.RegisterType((*ReplyLotteryShowInfo)(nil), "types.ReplyLotteryShowInfo")
	proto.RegisterType((*LotteryNumberRecord)(nil), "types.LotteryNumberRecord")
	proto.RegisterType((*LotteryBuyRecord)(nil), "types.LotteryBuyRecord")
	proto.RegisterType((*LotteryBuyRecords)(nil), "types.LotteryBuyRecords")
	proto.RegisterType((*LotteryDrawRecord)(nil), "types.LotteryDrawRecord")
	proto.RegisterType((*LotteryDrawRecords)(nil), "types.LotteryDrawRecords")
}

func init() { proto.RegisterFile("lottery.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0xaf, 0xff, 0x25, 0xd7, 0x71, 0x5b, 0xae, 0xdb, 0xd2, 0xb3, 0x00, 0x55, 0x95, 0x25, 0x50,
	0xa5, 0x3b, 0x22, 0x08, 0x42, 0x42, 0x7c, 0xa2, 0x39, 0x0e, 0xa5, 0x52, 0x55, 0x4e, 0xdb, 0x22,
	0x10, 0x7c, 0x72, 0xed, 0x3d, 0x1a, 0xd5, 0xb1, 0xc3, 0x7a, 0x7d, 0xc5, 0xdf, 0x10, 0x8f, 0x80,
	0xc4, 0x0b, 0xf0, 0x14, 0x3c, 0x09, 0x0f, 0xc0, 0x93, 0xa0, 0x9d, 0xdd, 0xc4, 0x6b, 0xc7, 0x69,
	0x72, 0xdc, 0x37, 0xef, 0xec, 0xf8, 0xb7, 0xf3, 0xfb, 0xcd, 0xec, 0xcc, 0xc2, 0x6e, 0x9a, 0x0b,
	0xc1, 0x78, 0x35, 0x98, 0xf1, 0x5c, 0xe4, 0xc4, 0x13, 0xd5, 0x8c, 0x15, 0xe1, 0x57, 0xb0, 0xf7,
	0xb2, 0xe4, 0xf1, 0x6d, 0x54, 0x30, 0xca, 0xe2, 0x9c, 0x27, 0xe4, 0x08, 0x7a, 0xd1, 0x34, 0x2f,
	0x33, 0x11, 0x58, 0x27, 0xd6, 0xa9, 0x43, 0xf5, 0x4a, 0xda, 0xb3, 0x72, 0x7a, 0xc3, 0x78, 0x60,
	0x2b, 0xbb, 0x5a, 0x85, 0xbf, 0x5b, 0xf0, 0x4e, 0x13, 0xa2, 0x20, 0x1f, 0x43, 0x8f, 0xe3, 0x67,
	0x60, 0x9d, 0x38, 0xa7, 0xfe, 0xf0, 0xdd, 0x01, 0x9e, 0x36, 0x68, 0xfa, 0x51, 0xed, 0x44, 0x02,
	0xe8, 0xbf, 0x2a, 0xb3, 0xe4, 0xfb, 0x49, 0xa6, 0xb1, 0xe7, 0x4b, 0xf2, 0x11, 0xec, 0xa9, 0xe3,
	0xbf, 0xcd, 0x18, 0xcd, 0xcb, 0x2c, 0x09, 0x1c, 0x74, 0x68, 0x59, 0xc3, 0xbf, 0x5d, 0xe8, 0x5f,
	0x28, 0x7e, 0xe4, 0x03, 0xd8, 0xd6, 0x54, 0xcf, 0x13, 0xe4, 0xb0, 0x4d, 0x6b, 0x83, 0xa4, 0x51,
	0x88, 0x48, 0x94, 0x05, 0x1e, 0xe5, 0x51, 0xbd, 0x22, 0x21, 0xec, 0xc4, 0x9c, 0x45, 0x82, 0x8d,
	0xd9, 0xe4, 0xe7, 0x5b, 0xa1, 0xcf, 0x69, 0xd8, 0x08, 0x01, 0x57, 0x06, 0x16, 0xb8, 0xb8, 0x87,
	0xdf, 0xe4, 0x04, 0xfc, 0x59, 0xc9, 0x47, 0x69, 0x1e, 0xdf, 0x5d, 0x96, 0xd3, 0xc0, 0xc3, 0x2d,
	0xd3, 0x24, 0x91, 0x13, 0x1e, 0xdd, 0x2f, 0x5c, 0x7a, 0x0a, 0xd9, 0xb4, 0x91, 0x4f, 0xe0, 0x20,
	0x8d, 0x0a, 0x71, 0xcd, 0xa3, 0xac, 0xb8, 0xce, 0x5f, 0x96, 0xfc, 0x4a, 0x44, 0x82, 0x05, 0x7d,
	0x74, 0xed, 0xda, 0x22, 0x43, 0x38, 0x34, 0xcc, 0x5f, 0xf3, 0xe8, 0x5e, 0xfd, 0xf2, 0x08, 0x7f,
	0xe9, 0xdc, 0x23, 0x9f, 0x43, 0x5f, 0x29, 0x5e, 0x04, 0xdb, 0x98, 0x97, 0xf7, 0x75, 0x5e, 0xb4,
	0x74, 0x03, 0x9d, 0xbf, 0x17, 0x99, 0xe0, 0x15, 0x9d, 0xfb, 0xca, 0xe0, 0x44, 0x2e, 0xa2, 0x74,
	0x9e, 0xbd, 0xe4, 0xfa, 0x57, 0xc9, 0x03, 0x54, 0x70, 0x1d, 0x5b, 0xe4, 0x18, 0x40, 0x09, 0x77,
	0x96, 0x24, 0x3c, 0xf0, 0x31, 0x07, 0x86, 0x85, 0x1c, 0x82, 0xc7, 0x31, 0x9b, 0x3b, 0x88, 0xa1,
	0x16, 0x52, 0xca, 0xb4, 0x8c, 0xef, 0xaa, 0x4b, 0x55, 0x66, 0xbb, 0x4a, 0x4a, 0xc3, 0xf4, 0x1e,
	0x85, 0x1d, 0x33, 0x44, 0xf2, 0x18, 0x9c, 0x3b, 0x56, 0xe9, 0x24, 0xcb, 0x4f, 0xf2, 0x0c, 0xbc,
	0xd7, 0x51, 0x5a, 0x32, 0xcc, 0xae, 0x3f, 0x3c, 0xea, 0x2c, 0xbc, 0x82, 0x2a, 0xa7, 0x2f, 0xed,
	0x2f, 0xac, 0xf0, 0x1f, 0x0b, 0x76, 0x35, 0xff, 0xb3, 0x58, 0x4c, 0xf2, 0x8c, 0x0c, 0xa0, 0xa7,
	0x62, 0x45, 0x60, 0x7f, 0x78, 0xd8, 0x54, 0xe9, 0xb9, 0x2a, 0x89, 0x2d, 0xaa, 0xbd, 0xc8, 0x87,
	0xe0, 0xdc, 0x94, 0x95, 0x3e, 0x71, 0xbf, 0xe9, 0x3c, 0x2a, 0xab, 0xf1, 0x16, 0x95, 0xfb, 0xe4,
	0x14, 0x5c, 0x99, 0x73, 0xac, 0x2c, 0x7f, 0x48, 0x9a, 0x7e, 0x32, 0x49, 0xe3, 0x2d, 0x8a, 0x1e,
	0xe4, 0x29, 0x78, 0x71, 0x9a, 0x17, 0x0c, 0x0b, 0xcd, 0x1f, 0x1e, 0xb4, 0xce, 0x97, 0x5b, 0xe3,
	0x2d, 0xaa, 0x7c, 0xc8, 0x1e, 0xd8, 0xa2, 0xc2, 0x64, 0x78, 0xd4, 0x16, 0xd5, 0xa8, 0xaf, 0x15,
	0x08, 0xbf, 0x5b, 0xf0, 0x52, 0x11, 0xb7, 0x4b, 0xd5, 0x5a, 0x5f, 0xaa, 0xf6, 0x72, 0xa9, 0x86,
	0x3f, 0x02, 0xd4, 0xdc, 0xd6, 0x5f, 0x36, 0xdd, 0x4b, 0xec, 0x15, 0xbd, 0xc4, 0x69, 0xf4, 0x92,
	0xa7, 0xe0, 0x1b, 0x7a, 0x3c, 0x0c, 0x1e, 0x3e, 0x83, 0x1d, 0x53, 0x91, 0x35, 0xde, 0xff, 0x5a,
	0xb0, 0x47, 0x59, 0xcc, 0x26, 0x33, 0xf1, 0x76, 0x8d, 0xe2, 0x18, 0x60, 0xc6, 0xd9, 0xeb, 0x2b,
	0xb5, 0xe7, 0xe0, 0x9e, 0x61, 0x91, 0x4d, 0x22, 0x92, 0x55, 0xef, 0x22, 0x20, 0x7e, 0xd7, 0xf5,
	0xee, 0x99, 0xf5, 0x5e, 0xab, 0xd0, 0x33, 0x55, 0x30, 0x54, 0xeb, 0x37, 0x54, 0x6b, 0xdd, 0x8f,
	0x47, 0x4b, 0xf7, 0x23, 0x1c, 0x48, 0x8e, 0xbf, 0x68, 0x7e, 0xe7, 0xd9, 0xab, 0x7c, 0x8d, 0x28,
	0x3f, 0xc1, 0x7e, 0xed, 0x3f, 0x2a, 0x37, 0xf8, 0x65, 0x41, 0xcf, 0xee, 0xa2, 0xe7, 0x18, 0xf4,
	0xc2, 0x3f, 0x2c, 0x38, 0x6c, 0xa0, 0x8f, 0x27, 0x85, 0xc8, 0xd7, 0xea, 0xbe, 0xf1, 0x01, 0xd2,
	0x1a, 0xa3, 0x4c, 0x2e, 0x26, 0x41, 0x2d, 0x24, 0x7a, 0x32, 0xe1, 0x0c, 0xaf, 0x32, 0xea, 0xed,
	0xd1, 0xda, 0x10, 0x9e, 0xc3, 0x41, 0x1d, 0xd3, 0x85, 0x94, 0x6e, 0x03, 0xce, 0x8b, 0xe3, 0x6d,
	0x93, 0xdf, 0x6f, 0x16, 0x1c, 0xb5, 0xb0, 0x36, 0x63, 0xd8, 0x09, 0x57, 0xb3, 0x71, 0x56, 0xb2,
	0x71, 0xdb, 0x6c, 0xfe, 0xc2, 0x10, 0x66, 0x69, 0xa5, 0x83, 0xb8, 0xcc, 0xf9, 0x34, 0x4a, 0x91,
	0x51, 0x7b, 0x9e, 0x59, 0x1d, 0xf3, 0xac, 0xd5, 0x10, 0xec, 0xf5, 0x0d, 0xc1, 0xe9, 0x98, 0x5d,
	0xcd, 0x66, 0xef, 0xb6, 0x9b, 0x7d, 0xf8, 0xa7, 0x0d, 0x4f, 0xcc, 0x20, 0x9f, 0x97, 0x9c, 0xb3,
	0x4c, 0x60, 0x94, 0xf5, 0x25, 0xb3, 0x1a, 0x97, 0x6c, 0x3e, 0x69, 0x6d, 0x63, 0xd2, 0xae, 0x98,
	0x91, 0xce, 0x9b, 0xcf, 0x48, 0xf7, 0x81, 0x19, 0xb9, 0x62, 0xd8, 0x79, 0xab, 0x87, 0xdd, 0x22,
	0x9d, 0xbd, 0x07, 0x86, 0x59, 0x7f, 0xf9, 0xb2, 0x8e, 0xe0, 0xd8, 0x94, 0x45, 0xd7, 0xce, 0x45,
	0xed, 0xd1, 0xc6, 0x90, 0x6f, 0xa9, 0x16, 0xc6, 0xb9, 0xbc, 0x62, 0x35, 0xc6, 0xd5, 0x6d, 0x7e,
	0x8f, 0xba, 0x7e, 0x5a, 0x4f, 0x7a, 0xf5, 0x02, 0x7b, 0xb2, 0x34, 0x96, 0xf4, 0x1b, 0x6c, 0xee,
	0x17, 0xbe, 0x80, 0x83, 0x79, 0x15, 0x21, 0x76, 0xfd, 0x1c, 0xcc, 0xe6, 0xc7, 0x77, 0x37, 0xa9,
	0x46, 0x6b, 0x0f, 0x7f, 0x80, 0xc7, 0xed, 0x33, 0xde, 0x14, 0x63, 0x45, 0x3f, 0xf9, 0x06, 0xf6,
	0xdb, 0xc8, 0xc5, 0xff, 0x21, 0x7a, 0xb6, 0xc0, 0x91, 0x59, 0x5f, 0x13, 0x62, 0xf7, 0xd5, 0x1f,
	0x03, 0x59, 0x82, 0x28, 0xc8, 0xb0, 0x1d, 0x4b, 0xb0, 0x3c, 0xe3, 0x5b, 0xc1, 0xdc, 0xf4, 0xf0,
	0x35, 0xfe, 0xd9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xcb, 0x3f, 0xe2, 0x9e, 0x0b, 0x00,
	0x00,
}