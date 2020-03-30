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

  
type MessageGlobalInfo struct {

	
	
		Total_online_users int32
	

	
	
		Chating_users int32
	

	
	
		Paid_users int32
	

	
	
		Broilers int32
	

	
	
		Playing int32
	

	
	
		Visitor int32
	

}

func NewMessageGlobalInfo() (*MessageGlobalInfo) {
	return &MessageGlobalInfo{}
}

func DeserializeMessageGlobalInfo(r io.Reader) (*MessageGlobalInfo, error) {
	t := NewMessageGlobalInfo()
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

func DeserializeMessageGlobalInfoFromSchema(r io.Reader, schema string) (*MessageGlobalInfo, error) {
	t := NewMessageGlobalInfo()

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

func writeMessageGlobalInfo(r *MessageGlobalInfo, w io.Writer) error {
	var err error
	
	err = vm.WriteInt( r.Total_online_users, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Chating_users, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Paid_users, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Broilers, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Playing, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteInt( r.Visitor, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MessageGlobalInfo) Serialize(w io.Writer) error {
	return writeMessageGlobalInfo(r, w)
}

func (r *MessageGlobalInfo) Schema() string {
	return "{\"fields\":[{\"name\":\"total_online_users\",\"type\":\"int\"},{\"name\":\"chating_users\",\"type\":\"int\"},{\"name\":\"paid_users\",\"type\":\"int\"},{\"name\":\"broilers\",\"type\":\"int\"},{\"name\":\"playing\",\"type\":\"int\"},{\"name\":\"visitor\",\"type\":\"int\"}],\"name\":\"proto.MessageGlobalInfo\",\"type\":\"record\"}"
}

func (r *MessageGlobalInfo) SchemaName() string {
	return "proto.MessageGlobalInfo"
}

func (_ *MessageGlobalInfo) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) SetString(v string) { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageGlobalInfo) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.Total_online_users)
		
	
	case 1:
		
		
			return (*types.Int)(&r.Chating_users)
		
	
	case 2:
		
		
			return (*types.Int)(&r.Paid_users)
		
	
	case 3:
		
		
			return (*types.Int)(&r.Broilers)
		
	
	case 4:
		
		
			return (*types.Int)(&r.Playing)
		
	
	case 5:
		
		
			return (*types.Int)(&r.Visitor)
		
	
	}
	panic("Unknown field index")
}

func (r *MessageGlobalInfo) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MessageGlobalInfo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MessageGlobalInfo) Finalize() { }