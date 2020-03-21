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


type UnionNullMessageRoomWaitingCustomersTypeEnum int
const (

	 UnionNullMessageRoomWaitingCustomersTypeEnumNull UnionNullMessageRoomWaitingCustomersTypeEnum = 0

	 UnionNullMessageRoomWaitingCustomersTypeEnumMessageRoomWaitingCustomers UnionNullMessageRoomWaitingCustomersTypeEnum = 1

)

type UnionNullMessageRoomWaitingCustomers struct {

	Null *types.NullVal

	MessageRoomWaitingCustomers *MessageRoomWaitingCustomers

	UnionType UnionNullMessageRoomWaitingCustomersTypeEnum
}

func writeUnionNullMessageRoomWaitingCustomers(r *UnionNullMessageRoomWaitingCustomers, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullMessageRoomWaitingCustomersTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullMessageRoomWaitingCustomersTypeEnumMessageRoomWaitingCustomers:
		return writeMessageRoomWaitingCustomers(r.MessageRoomWaitingCustomers, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullMessageRoomWaitingCustomers")
}

func NewUnionNullMessageRoomWaitingCustomers() *UnionNullMessageRoomWaitingCustomers {
	return &UnionNullMessageRoomWaitingCustomers{}
}

func (_ *UnionNullMessageRoomWaitingCustomers) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullMessageRoomWaitingCustomers) SetLong(v int64) { 
	r.UnionType = (UnionNullMessageRoomWaitingCustomersTypeEnum)(v)
}
func (r *UnionNullMessageRoomWaitingCustomers) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.MessageRoomWaitingCustomers = NewMessageRoomWaitingCustomers()
		
		
		return r.MessageRoomWaitingCustomers
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullMessageRoomWaitingCustomers) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomWaitingCustomers) Finalize()  { }
