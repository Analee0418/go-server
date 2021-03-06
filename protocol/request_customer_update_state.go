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

  
type RequestCustomerUpdateState struct {

	
	
		State CustomerState
	

}

func NewRequestCustomerUpdateState() (*RequestCustomerUpdateState) {
	return &RequestCustomerUpdateState{}
}

func DeserializeRequestCustomerUpdateState(r io.Reader) (*RequestCustomerUpdateState, error) {
	t := NewRequestCustomerUpdateState()
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

func DeserializeRequestCustomerUpdateStateFromSchema(r io.Reader, schema string) (*RequestCustomerUpdateState, error) {
	t := NewRequestCustomerUpdateState()

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

func writeRequestCustomerUpdateState(r *RequestCustomerUpdateState, w io.Writer) error {
	var err error
	
	err = writeCustomerState( r.State, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *RequestCustomerUpdateState) Serialize(w io.Writer) error {
	return writeRequestCustomerUpdateState(r, w)
}

func (r *RequestCustomerUpdateState) Schema() string {
	return "{\"fields\":[{\"name\":\"state\",\"type\":{\"name\":\"enum.CustomerState\",\"symbols\":[\"idle\",\"during_chat\",\"game\",\"browse_product\",\"paying\"],\"type\":\"enum\"}}],\"name\":\"proto.RequestCustomerUpdateState\",\"type\":\"record\"}"
}

func (r *RequestCustomerUpdateState) SchemaName() string {
	return "proto.RequestCustomerUpdateState"
}

func (_ *RequestCustomerUpdateState) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) SetInt(v int32) { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) SetLong(v int64) { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) SetString(v string) { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RequestCustomerUpdateState) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.State)
		
	
	}
	panic("Unknown field index")
}

func (r *RequestCustomerUpdateState) SetDefault(i int) {
	switch (i) {
	
        
	
	}
	panic("Unknown field index")
}

func (_ *RequestCustomerUpdateState) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *RequestCustomerUpdateState) Finalize() { }
