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


type Message_cars_modelUnionTypeEnum int
const (

	 Message_cars_modelUnionTypeEnumNull Message_cars_modelUnionTypeEnum = 0

	 Message_cars_modelUnionTypeEnumMessageCarsModel Message_cars_modelUnionTypeEnum = 1

)

type Message_cars_modelUnion struct {

	Null *types.NullVal

	MessageCarsModel *MessageCarsModel

	UnionType Message_cars_modelUnionTypeEnum
}

func writeMessage_cars_modelUnion(r *Message_cars_modelUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case Message_cars_modelUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case Message_cars_modelUnionTypeEnumMessageCarsModel:
		return writeMessageCarsModel(r.MessageCarsModel, w)
        
	}
	return fmt.Errorf("invalid value for *Message_cars_modelUnion")
}

func NewMessage_cars_modelUnion() *Message_cars_modelUnion {
	return &Message_cars_modelUnion{}
}

func (_ *Message_cars_modelUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) SetString(v string) { panic("Unsupported operation") }
func (r *Message_cars_modelUnion) SetLong(v int64) { 
	r.UnionType = (Message_cars_modelUnionTypeEnum)(v)
}
func (r *Message_cars_modelUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.MessageCarsModel = NewMessageCarsModel()
		
		
		return r.MessageCarsModel
		
	
	}
	panic("Unknown field index")
}
func (_ *Message_cars_modelUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Message_cars_modelUnion) Finalize()  { }