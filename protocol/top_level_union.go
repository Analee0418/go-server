// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package protocol

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type TopLevelUnionTypeEnum int
const (

	 TopLevelUnionTypeEnumGlobalState TopLevelUnionTypeEnum = 0

	 TopLevelUnionTypeEnumAction TopLevelUnionTypeEnum = 1

	 TopLevelUnionTypeEnumMessageSession TopLevelUnionTypeEnum = 2

	 TopLevelUnionTypeEnumMessageCustomersInfo TopLevelUnionTypeEnum = 3

	 TopLevelUnionTypeEnumMessageRoomWaitingCustomers TopLevelUnionTypeEnum = 4

	 TopLevelUnionTypeEnumMessageCustomersAuctionInfo TopLevelUnionTypeEnum = 5

	 TopLevelUnionTypeEnumMessageCarsModel TopLevelUnionTypeEnum = 6

	 TopLevelUnionTypeEnumMessageRoomInfo TopLevelUnionTypeEnum = 7

	 TopLevelUnionTypeEnumMessageContract TopLevelUnionTypeEnum = 8

	 TopLevelUnionTypeEnumMessageAuctionRecord TopLevelUnionTypeEnum = 9

	 TopLevelUnionTypeEnumMessageAuctionGoods TopLevelUnionTypeEnum = 10

	 TopLevelUnionTypeEnumMessageAuctionInfo TopLevelUnionTypeEnum = 11

	 TopLevelUnionTypeEnumMessageGlobalInfo TopLevelUnionTypeEnum = 12

	 TopLevelUnionTypeEnumMessageForward TopLevelUnionTypeEnum = 13

	 TopLevelUnionTypeEnumMessage TopLevelUnionTypeEnum = 14

)

type TopLevelUnion struct {

	GlobalState GlobalState

	Action Action

	MessageSession *MessageSession

	MessageCustomersInfo *MessageCustomersInfo

	MessageRoomWaitingCustomers *MessageRoomWaitingCustomers

	MessageCustomersAuctionInfo *MessageCustomersAuctionInfo

	MessageCarsModel *MessageCarsModel

	MessageRoomInfo *MessageRoomInfo

	MessageContract *MessageContract

	MessageAuctionRecord *MessageAuctionRecord

	MessageAuctionGoods *MessageAuctionGoods

	MessageAuctionInfo *MessageAuctionInfo

	MessageGlobalInfo *MessageGlobalInfo

	MessageForward *MessageForward

	Message *Message

	UnionType TopLevelUnionTypeEnum
}

func writeTopLevelUnion(r *TopLevelUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case TopLevelUnionTypeEnumGlobalState:
		return writeGlobalState(r.GlobalState, w)
        
	case TopLevelUnionTypeEnumAction:
		return writeAction(r.Action, w)
        
	case TopLevelUnionTypeEnumMessageSession:
		return writeMessageSession(r.MessageSession, w)
        
	case TopLevelUnionTypeEnumMessageCustomersInfo:
		return writeMessageCustomersInfo(r.MessageCustomersInfo, w)
        
	case TopLevelUnionTypeEnumMessageRoomWaitingCustomers:
		return writeMessageRoomWaitingCustomers(r.MessageRoomWaitingCustomers, w)
        
	case TopLevelUnionTypeEnumMessageCustomersAuctionInfo:
		return writeMessageCustomersAuctionInfo(r.MessageCustomersAuctionInfo, w)
        
	case TopLevelUnionTypeEnumMessageCarsModel:
		return writeMessageCarsModel(r.MessageCarsModel, w)
        
	case TopLevelUnionTypeEnumMessageRoomInfo:
		return writeMessageRoomInfo(r.MessageRoomInfo, w)
        
	case TopLevelUnionTypeEnumMessageContract:
		return writeMessageContract(r.MessageContract, w)
        
	case TopLevelUnionTypeEnumMessageAuctionRecord:
		return writeMessageAuctionRecord(r.MessageAuctionRecord, w)
        
	case TopLevelUnionTypeEnumMessageAuctionGoods:
		return writeMessageAuctionGoods(r.MessageAuctionGoods, w)
        
	case TopLevelUnionTypeEnumMessageAuctionInfo:
		return writeMessageAuctionInfo(r.MessageAuctionInfo, w)
        
	case TopLevelUnionTypeEnumMessageGlobalInfo:
		return writeMessageGlobalInfo(r.MessageGlobalInfo, w)
        
	case TopLevelUnionTypeEnumMessageForward:
		return writeMessageForward(r.MessageForward, w)
        
	case TopLevelUnionTypeEnumMessage:
		return writeMessage(r.Message, w)
        
	}
	return fmt.Errorf("invalid value for *TopLevelUnion")
}

func NewTopLevelUnion() *TopLevelUnion {
	return &TopLevelUnion{}
}

func (_ *TopLevelUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *TopLevelUnion) SetString(v string) { panic("Unsupported operation") }
func (r *TopLevelUnion) SetLong(v int64) { 
	r.UnionType = (TopLevelUnionTypeEnum)(v)
}
func (r *TopLevelUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return (*types.Int)(&r.GlobalState)
		
	
	case 1:
		
		
		return (*types.Int)(&r.Action)
		
	
	case 2:
		
		r.MessageSession = NewMessageSession()
		
		
		return r.MessageSession
		
	
	case 3:
		
		r.MessageCustomersInfo = NewMessageCustomersInfo()
		
		
		return r.MessageCustomersInfo
		
	
	case 4:
		
		r.MessageRoomWaitingCustomers = NewMessageRoomWaitingCustomers()
		
		
		return r.MessageRoomWaitingCustomers
		
	
	case 5:
		
		r.MessageCustomersAuctionInfo = NewMessageCustomersAuctionInfo()
		
		
		return r.MessageCustomersAuctionInfo
		
	
	case 6:
		
		r.MessageCarsModel = NewMessageCarsModel()
		
		
		return r.MessageCarsModel
		
	
	case 7:
		
		r.MessageRoomInfo = NewMessageRoomInfo()
		
		
		return r.MessageRoomInfo
		
	
	case 8:
		
		r.MessageContract = NewMessageContract()
		
		
		return r.MessageContract
		
	
	case 9:
		
		r.MessageAuctionRecord = NewMessageAuctionRecord()
		
		
		return r.MessageAuctionRecord
		
	
	case 10:
		
		r.MessageAuctionGoods = NewMessageAuctionGoods()
		
		
		return r.MessageAuctionGoods
		
	
	case 11:
		
		r.MessageAuctionInfo = NewMessageAuctionInfo()
		
		
		return r.MessageAuctionInfo
		
	
	case 12:
		
		r.MessageGlobalInfo = NewMessageGlobalInfo()
		
		
		return r.MessageGlobalInfo
		
	
	case 13:
		
		r.MessageForward = NewMessageForward()
		
		
		return r.MessageForward
		
	
	case 14:
		
		r.Message = NewMessage()
		
		
		return r.Message
		
	
	}
	panic("Unknown field index")
}
func (_ *TopLevelUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *TopLevelUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *TopLevelUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *TopLevelUnion) Finalize()  { }
