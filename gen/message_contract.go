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

  
type MessageContract struct {

	
	
		Contract_id int32
	

}

func NewMessageContract() (*MessageContract) {
	return &MessageContract{}
}

func DeserializeMessageContract(r io.Reader) (*MessageContract, error) {
	t := NewMessageContract()
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

func DeserializeMessageContractFromSchema(r io.Reader, schema string) (*MessageContract, error) {
	t := NewMessageContract()

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

func writeMessageContract(r *MessageContract, w io.Writer) error {
	var err error
	
	err = vm.WriteInt( r.Contract_id, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MessageContract) Serialize(w io.Writer) error {
	return writeMessageContract(r, w)
}

func (r *MessageContract) Schema() string {
	return "{\"fields\":[{\"name\":\"contract_id\",\"type\":\"int\"}],\"name\":\"proto.MessageContract\",\"type\":\"record\"}"
}

func (r *MessageContract) SchemaName() string {
	return "proto.MessageContract"
}

func (_ *MessageContract) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MessageContract) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MessageContract) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MessageContract) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MessageContract) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MessageContract) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MessageContract) SetString(v string) { panic("Unsupported operation") }
func (_ *MessageContract) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageContract) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.Contract_id)
		
	
	}
	panic("Unknown field index")
}

func (r *MessageContract) SetDefault(i int) {
	switch (i) {
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MessageContract) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MessageContract) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MessageContract) Finalize() { }