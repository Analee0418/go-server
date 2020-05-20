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

  
type MessageGameConfig struct {

	
	
		GameID *GameIDUnion
	

	
	
		Config *ConfigUnion
	

}

func NewMessageGameConfig() (*MessageGameConfig) {
	return &MessageGameConfig{}
}

func DeserializeMessageGameConfig(r io.Reader) (*MessageGameConfig, error) {
	t := NewMessageGameConfig()
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

func DeserializeMessageGameConfigFromSchema(r io.Reader, schema string) (*MessageGameConfig, error) {
	t := NewMessageGameConfig()

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

func writeMessageGameConfig(r *MessageGameConfig, w io.Writer) error {
	var err error
	
	err = writeGameIDUnion( r.GameID, w)
	if err != nil {
		return err			
	}
	
	err = writeConfigUnion( r.Config, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MessageGameConfig) Serialize(w io.Writer) error {
	return writeMessageGameConfig(r, w)
}

func (r *MessageGameConfig) Schema() string {
	return "{\"fields\":[{\"name\":\"gameID\",\"type\":[\"null\",\"string\"]},{\"name\":\"config\",\"type\":[\"null\",\"string\"]}],\"name\":\"proto.MessageGameConfig\",\"type\":\"record\"}"
}

func (r *MessageGameConfig) SchemaName() string {
	return "proto.MessageGameConfig"
}

func (_ *MessageGameConfig) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MessageGameConfig) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MessageGameConfig) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MessageGameConfig) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MessageGameConfig) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MessageGameConfig) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MessageGameConfig) SetString(v string) { panic("Unsupported operation") }
func (_ *MessageGameConfig) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageGameConfig) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.GameID = NewGameIDUnion()
	
		
		
			return r.GameID
		
	
	case 1:
		
			r.Config = NewConfigUnion()
	
		
		
			return r.Config
		
	
	}
	panic("Unknown field index")
}

func (r *MessageGameConfig) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MessageGameConfig) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MessageGameConfig) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MessageGameConfig) Finalize() { }
