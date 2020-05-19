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


type Message_global_stateUnionTypeEnum int
const (

	 Message_global_stateUnionTypeEnumNull Message_global_stateUnionTypeEnum = 0

	 Message_global_stateUnionTypeEnumMessageGlobalState Message_global_stateUnionTypeEnum = 1

)

type Message_global_stateUnion struct {

	Null *types.NullVal

	MessageGlobalState *MessageGlobalState

	UnionType Message_global_stateUnionTypeEnum
}

func writeMessage_global_stateUnion(r *Message_global_stateUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case Message_global_stateUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case Message_global_stateUnionTypeEnumMessageGlobalState:
		return writeMessageGlobalState(r.MessageGlobalState, w)
        
	}
	return fmt.Errorf("invalid value for *Message_global_stateUnion")
}

func NewMessage_global_stateUnion() *Message_global_stateUnion {
	return &Message_global_stateUnion{}
}

func (_ *Message_global_stateUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) SetString(v string) { panic("Unsupported operation") }
func (r *Message_global_stateUnion) SetLong(v int64) { 
	r.UnionType = (Message_global_stateUnionTypeEnum)(v)
}
func (r *Message_global_stateUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.MessageGlobalState = NewMessageGlobalState()
		
		
		return r.MessageGlobalState
		
	
	}
	panic("Unknown field index")
}
func (_ *Message_global_stateUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Message_global_stateUnion) Finalize()  { }