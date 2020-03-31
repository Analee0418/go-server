// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package protocol

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type MessageAuctionRecord struct {

	
	
		Goods_id int32
	

	
	
		Customer_mobile *Customer_mobileUnion
	

	
	
		Customer_mobile_region *Customer_mobile_regionUnion
	

	
	
		Customer_idcard *Customer_idcardUnion
	

	
	
		Customer_username *Customer_usernameUnion
	

	
	
		Bid_price float32
	

	
	
		Is_final bool
	

	
	
		Timestamp int64
	

}

func NewMessageAuctionRecord() (*MessageAuctionRecord) {
	return &MessageAuctionRecord{}
}

func DeserializeMessageAuctionRecord(r io.Reader) (*MessageAuctionRecord, error) {
	t := NewMessageAuctionRecord()
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

func DeserializeMessageAuctionRecordFromSchema(r io.Reader, schema string) (*MessageAuctionRecord, error) {
	t := NewMessageAuctionRecord()

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

func writeMessageAuctionRecord(r *MessageAuctionRecord, w io.Writer) error {
	var err error
	
	err = vm.WriteInt( r.Goods_id, w)
	if err != nil {
		return err			
	}
	
	err = writeCustomer_mobileUnion( r.Customer_mobile, w)
	if err != nil {
		return err			
	}
	
	err = writeCustomer_mobile_regionUnion( r.Customer_mobile_region, w)
	if err != nil {
		return err			
	}
	
	err = writeCustomer_idcardUnion( r.Customer_idcard, w)
	if err != nil {
		return err			
	}
	
	err = writeCustomer_usernameUnion( r.Customer_username, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteFloat( r.Bid_price, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteBool( r.Is_final, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteLong( r.Timestamp, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MessageAuctionRecord) Serialize(w io.Writer) error {
	return writeMessageAuctionRecord(r, w)
}

func (r *MessageAuctionRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"goods_id\",\"type\":\"int\"},{\"name\":\"customer_mobile\",\"type\":[\"null\",\"string\"]},{\"name\":\"customer_mobile_region\",\"type\":[\"null\",\"string\"]},{\"name\":\"customer_idcard\",\"type\":[\"null\",\"string\"]},{\"name\":\"customer_username\",\"type\":[\"null\",\"string\"]},{\"name\":\"bid_price\",\"type\":\"float\"},{\"name\":\"is_final\",\"type\":\"boolean\"},{\"name\":\"timestamp\",\"type\":\"long\"}],\"name\":\"proto.MessageAuctionRecord\",\"type\":\"record\"}"
}

func (r *MessageAuctionRecord) SchemaName() string {
	return "proto.MessageAuctionRecord"
}

func (_ *MessageAuctionRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageAuctionRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.Goods_id)
		
	
	case 1:
		
			r.Customer_mobile = NewCustomer_mobileUnion()
	
		
		
			return r.Customer_mobile
		
	
	case 2:
		
			r.Customer_mobile_region = NewCustomer_mobile_regionUnion()
	
		
		
			return r.Customer_mobile_region
		
	
	case 3:
		
			r.Customer_idcard = NewCustomer_idcardUnion()
	
		
		
			return r.Customer_idcard
		
	
	case 4:
		
			r.Customer_username = NewCustomer_usernameUnion()
	
		
		
			return r.Customer_username
		
	
	case 5:
		
		
			return (*types.Float)(&r.Bid_price)
		
	
	case 6:
		
		
			return (*types.Boolean)(&r.Is_final)
		
	
	case 7:
		
		
			return (*types.Long)(&r.Timestamp)
		
	
	}
	panic("Unknown field index")
}

func (r *MessageAuctionRecord) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MessageAuctionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MessageAuctionRecord) Finalize() { }
