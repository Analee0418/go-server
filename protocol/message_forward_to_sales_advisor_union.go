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


type Message_forward_to_sales_advisorUnionTypeEnum int
const (

	 Message_forward_to_sales_advisorUnionTypeEnumNull Message_forward_to_sales_advisorUnionTypeEnum = 0

	 Message_forward_to_sales_advisorUnionTypeEnumMessageForward Message_forward_to_sales_advisorUnionTypeEnum = 1

)

type Message_forward_to_sales_advisorUnion struct {

	Null *types.NullVal

	MessageForward *MessageForward

	UnionType Message_forward_to_sales_advisorUnionTypeEnum
}

func writeMessage_forward_to_sales_advisorUnion(r *Message_forward_to_sales_advisorUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case Message_forward_to_sales_advisorUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case Message_forward_to_sales_advisorUnionTypeEnumMessageForward:
		return writeMessageForward(r.MessageForward, w)
        
	}
	return fmt.Errorf("invalid value for *Message_forward_to_sales_advisorUnion")
}

func NewMessage_forward_to_sales_advisorUnion() *Message_forward_to_sales_advisorUnion {
	return &Message_forward_to_sales_advisorUnion{}
}

func (_ *Message_forward_to_sales_advisorUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) SetString(v string) { panic("Unsupported operation") }
func (r *Message_forward_to_sales_advisorUnion) SetLong(v int64) { 
	r.UnionType = (Message_forward_to_sales_advisorUnionTypeEnum)(v)
}
func (r *Message_forward_to_sales_advisorUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.MessageForward = NewMessageForward()
		
		
		return r.MessageForward
		
	
	}
	panic("Unknown field index")
}
func (_ *Message_forward_to_sales_advisorUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Message_forward_to_sales_advisorUnion) Finalize()  { }
