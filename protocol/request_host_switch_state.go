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

  
type RequestHostSwitchState struct {

	
	
		GlobalState GlobalState
	

	
	
		CountDownSeconds int32
	

	
	
		Body string
	

}

func NewRequestHostSwitchState() (*RequestHostSwitchState) {
	return &RequestHostSwitchState{}
}

func DeserializeRequestHostSwitchState(r io.Reader) (*RequestHostSwitchState, error) {
	t := NewRequestHostSwitchState()
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

func DeserializeRequestHostSwitchStateFromSchema(r io.Reader, schema string) (*RequestHostSwitchState, error) {
	t := NewRequestHostSwitchState()

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

func writeRequestHostSwitchState(r *RequestHostSwitchState, w io.Writer) error {
	var err error
	
	err = writeGlobalState( r.GlobalState, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.CountDownSeconds, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteString( r.Body, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *RequestHostSwitchState) Serialize(w io.Writer) error {
	return writeRequestHostSwitchState(r, w)
}

func (r *RequestHostSwitchState) Schema() string {
	return "{\"fields\":[{\"name\":\"globalState\",\"type\":{\"name\":\"enum.GlobalState\",\"symbols\":[\"awating_starting\",\"starting_animations\",\"speeching\",\"aution\",\"products\",\"discount_strategy\",\"chat_with_advisor\"],\"type\":\"enum\"}},{\"name\":\"countDownSeconds\",\"type\":\"int\"},{\"name\":\"body\",\"type\":\"string\"}],\"name\":\"proto.RequestHostSwitchState\",\"type\":\"record\"}"
}

func (r *RequestHostSwitchState) SchemaName() string {
	return "proto.RequestHostSwitchState"
}

func (_ *RequestHostSwitchState) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) SetInt(v int32) { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) SetLong(v int64) { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) SetString(v string) { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RequestHostSwitchState) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.GlobalState)
		
	
	case 1:
		
		
			return (*types.Int)(&r.CountDownSeconds)
		
	
	case 2:
		
		
			return (*types.String)(&r.Body)
		
	
	}
	panic("Unknown field index")
}

func (r *RequestHostSwitchState) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *RequestHostSwitchState) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *RequestHostSwitchState) Finalize() { }
