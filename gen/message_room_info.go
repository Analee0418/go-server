// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type MessageRoomInfo struct {

	
	
		Room_id int32
	

	
	
		Order_count int32
	

	
	
		Customer_info *MessageCustomersInfo
	

	
	
		Waiting_list []*MessageCustomersInfo
	

	
	
		Customer_auction_info *MessageCustomersAuctionInfo
	

	
	
		Car_model *MessageCarsModel
	

}

func NewMessageRoomInfo() (*MessageRoomInfo) {
	return &MessageRoomInfo{}
}

func DeserializeMessageRoomInfo(r io.Reader) (*MessageRoomInfo, error) {
	t := NewMessageRoomInfo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func DeserializeMessageRoomInfoFromSchema(r io.Reader, schema string) (*MessageRoomInfo, error) {
	t := NewMessageRoomInfo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func writeMessageRoomInfo(r *MessageRoomInfo, w io.Writer) error {
	var err error
	
	err = vm.WriteInt( r.Room_id, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Order_count, w)
	if err != nil {
		return err			
	}
	
	err = writeMessageCustomersInfo( r.Customer_info, w)
	if err != nil {
		return err			
	}
	
	err = writeArrayMessageCustomersInfo( r.Waiting_list, w)
	if err != nil {
		return err			
	}
	
	err = writeMessageCustomersAuctionInfo( r.Customer_auction_info, w)
	if err != nil {
		return err			
	}
	
	err = writeMessageCarsModel( r.Car_model, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MessageRoomInfo) Serialize(w io.Writer) error {
	return writeMessageRoomInfo(r, w)
}

func (r *MessageRoomInfo) Schema() string {
	return "{\"fields\":[{\"name\":\"room_id\",\"type\":\"int\"},{\"name\":\"order_count\",\"type\":\"int\"},{\"name\":\"customer_info\",\"type\":{\"fields\":[{\"name\":\"mobile\",\"type\":[\"null\",\"string\"]},{\"name\":\"idcard\",\"type\":[\"null\",\"string\"]},{\"name\":\"username\",\"type\":[\"null\",\"string\"]}],\"name\":\"MessageCustomersInfo\",\"namespace\":\"proto\",\"type\":\"record\"}},{\"name\":\"waiting_list\",\"type\":{\"items\":\"proto.MessageCustomersInfo\",\"type\":\"array\"}},{\"name\":\"customer_auction_info\",\"type\":{\"fields\":[{\"name\":\"auction_list\",\"type\":{\"type\":\"map\",\"values\":\"double\"}},{\"name\":\"discount_list\",\"type\":{\"type\":\"map\",\"values\":\"double\"}}],\"name\":\"MessageCustomersAuctionInfo\",\"namespace\":\"proto\",\"type\":\"record\"}},{\"name\":\"car_model\",\"type\":{\"fields\":[{\"name\":\"brand\",\"type\":[\"null\",\"string\"]},{\"name\":\"color\",\"type\":[\"null\",\"string\"]},{\"name\":\"series\",\"type\":[\"null\",\"string\"]}],\"name\":\"MessageCarsModel\",\"namespace\":\"proto\",\"type\":\"record\"}}],\"name\":\"proto.MessageRoomInfo\",\"type\":\"record\"}"
}

func (r *MessageRoomInfo) SchemaName() string {
	return "proto.MessageRoomInfo"
}

func (_ *MessageRoomInfo) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MessageRoomInfo) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MessageRoomInfo) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MessageRoomInfo) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MessageRoomInfo) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MessageRoomInfo) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MessageRoomInfo) SetString(v string) { panic("Unsupported operation") }
func (_ *MessageRoomInfo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageRoomInfo) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.Room_id)
		
	
	case 1:
		
		
			return (*types.Int)(&r.Order_count)
		
	
	case 2:
		
			r.Customer_info = NewMessageCustomersInfo()
	
		
		
			return r.Customer_info
		
	
	case 3:
		
			r.Waiting_list = make([]*MessageCustomersInfo, 0)
	
		
		
			return (*ArrayMessageCustomersInfoWrapper)(&r.Waiting_list)
		
	
	case 4:
		
			r.Customer_auction_info = NewMessageCustomersAuctionInfo()
	
		
		
			return r.Customer_auction_info
		
	
	case 5:
		
			r.Car_model = NewMessageCarsModel()
	
		
		
			return r.Car_model
		
	
	}
	panic("Unknown field index")
}

func (r *MessageRoomInfo) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MessageRoomInfo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MessageRoomInfo) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MessageRoomInfo) Finalize() { }