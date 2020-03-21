// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package avro

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum int
const (

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumAction UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 0

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageCustomersInfo UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 1

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageRoomWaitingCustomers UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 2

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageCustomersAuctionInfo UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 3

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageCarsModel UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 4

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageRoomInfo UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 5

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageContract UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 6

	 UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessage UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum = 7

)

type UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage struct {

	Action Action

	MessageCustomersInfo *MessageCustomersInfo

	MessageRoomWaitingCustomers *MessageRoomWaitingCustomers

	MessageCustomersAuctionInfo *MessageCustomersAuctionInfo

	MessageCarsModel *MessageCarsModel

	MessageRoomInfo *MessageRoomInfo

	MessageContract *MessageContract

	Message *Message

	UnionType UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum
}

func writeUnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage(r *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumAction:
		return writeAction(r.Action, w)
        
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageCustomersInfo:
		return writeMessageCustomersInfo(r.MessageCustomersInfo, w)
        
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageRoomWaitingCustomers:
		return writeMessageRoomWaitingCustomers(r.MessageRoomWaitingCustomers, w)
        
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageCustomersAuctionInfo:
		return writeMessageCustomersAuctionInfo(r.MessageCustomersAuctionInfo, w)
        
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageCarsModel:
		return writeMessageCarsModel(r.MessageCarsModel, w)
        
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageRoomInfo:
		return writeMessageRoomInfo(r.MessageRoomInfo, w)
        
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessageContract:
		return writeMessageContract(r.MessageContract, w)
        
	case UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnumMessage:
		return writeMessage(r.Message, w)
        
	}
	return fmt.Errorf("invalid value for *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage")
}

func NewUnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage() *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage {
	return &UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage{}
}

func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetString(v string) { panic("Unsupported operation") }
func (r *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetLong(v int64) { 
	r.UnionType = (UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessageTypeEnum)(v)
}
func (r *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return (*types.Int)(&r.Action)
		
	
	case 1:
		
		r.MessageCustomersInfo = NewMessageCustomersInfo()
		
		
		return r.MessageCustomersInfo
		
	
	case 2:
		
		r.MessageRoomWaitingCustomers = NewMessageRoomWaitingCustomers()
		
		
		return r.MessageRoomWaitingCustomers
		
	
	case 3:
		
		r.MessageCustomersAuctionInfo = NewMessageCustomersAuctionInfo()
		
		
		return r.MessageCustomersAuctionInfo
		
	
	case 4:
		
		r.MessageCarsModel = NewMessageCarsModel()
		
		
		return r.MessageCarsModel
		
	
	case 5:
		
		r.MessageRoomInfo = NewMessageRoomInfo()
		
		
		return r.MessageRoomInfo
		
	
	case 6:
		
		r.MessageContract = NewMessageContract()
		
		
		return r.MessageContract
		
	
	case 7:
		
		r.Message = NewMessage()
		
		
		return r.Message
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionActionMessageCustomersInfoMessageRoomWaitingCustomersMessageCustomersAuctionInfoMessageCarsModelMessageRoomInfoMessageContractMessage) Finalize()  { }