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

  
type MessageRoomWaitingCustomers struct {

	
	
		Waiting_list []*MessageCustomersInfo
	

}

func NewMessageRoomWaitingCustomers() (*MessageRoomWaitingCustomers) {
	return &MessageRoomWaitingCustomers{}
}

func DeserializeMessageRoomWaitingCustomers(r io.Reader) (*MessageRoomWaitingCustomers, error) {
	t := NewMessageRoomWaitingCustomers()
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

func DeserializeMessageRoomWaitingCustomersFromSchema(r io.Reader, schema string) (*MessageRoomWaitingCustomers, error) {
	t := NewMessageRoomWaitingCustomers()

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

func writeMessageRoomWaitingCustomers(r *MessageRoomWaitingCustomers, w io.Writer) error {
	var err error
	
	err = writeArrayMessageCustomersInfo( r.Waiting_list, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MessageRoomWaitingCustomers) Serialize(w io.Writer) error {
	return writeMessageRoomWaitingCustomers(r, w)
}

func (r *MessageRoomWaitingCustomers) Schema() string {
	return "{\"fields\":[{\"name\":\"waiting_list\",\"type\":{\"items\":{\"fields\":[{\"name\":\"mobile\",\"type\":[\"null\",\"string\"]},{\"name\":\"mobileRegion\",\"type\":[\"null\",\"string\"]},{\"name\":\"idcard\",\"type\":[\"null\",\"string\"]},{\"name\":\"username\",\"type\":[\"null\",\"string\"]},{\"name\":\"address\",\"type\":[\"null\",\"string\"]}],\"name\":\"MessageCustomersInfo\",\"namespace\":\"proto\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"proto.MessageRoomWaitingCustomers\",\"type\":\"record\"}"
}

func (r *MessageRoomWaitingCustomers) SchemaName() string {
	return "proto.MessageRoomWaitingCustomers"
}

func (_ *MessageRoomWaitingCustomers) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) SetString(v string) { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageRoomWaitingCustomers) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.Waiting_list = make([]*MessageCustomersInfo, 0)
	
		
		
			return (*ArrayMessageCustomersInfoWrapper)(&r.Waiting_list)
		
	
	}
	panic("Unknown field index")
}

func (r *MessageRoomWaitingCustomers) SetDefault(i int) {
	switch (i) {
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MessageRoomWaitingCustomers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MessageRoomWaitingCustomers) Finalize() { }
