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


type Customer_join_queueUnionTypeEnum int
const (

	 Customer_join_queueUnionTypeEnumNull Customer_join_queueUnionTypeEnum = 0

	 Customer_join_queueUnionTypeEnumRequestCustomerJoinQueue Customer_join_queueUnionTypeEnum = 1

)

type Customer_join_queueUnion struct {

	Null *types.NullVal

	RequestCustomerJoinQueue *RequestCustomerJoinQueue

	UnionType Customer_join_queueUnionTypeEnum
}

func writeCustomer_join_queueUnion(r *Customer_join_queueUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case Customer_join_queueUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case Customer_join_queueUnionTypeEnumRequestCustomerJoinQueue:
		return writeRequestCustomerJoinQueue(r.RequestCustomerJoinQueue, w)
        
	}
	return fmt.Errorf("invalid value for *Customer_join_queueUnion")
}

func NewCustomer_join_queueUnion() *Customer_join_queueUnion {
	return &Customer_join_queueUnion{}
}

func (_ *Customer_join_queueUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) SetString(v string) { panic("Unsupported operation") }
func (r *Customer_join_queueUnion) SetLong(v int64) { 
	r.UnionType = (Customer_join_queueUnionTypeEnum)(v)
}
func (r *Customer_join_queueUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.RequestCustomerJoinQueue = NewRequestCustomerJoinQueue()
		
		
		return r.RequestCustomerJoinQueue
		
	
	}
	panic("Unknown field index")
}
func (_ *Customer_join_queueUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Customer_join_queueUnion) Finalize()  { }