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

  
type MessageCarsModel struct {

	
	
		Brand *BrandUnion
	

	
	
		Color *ColorUnion
	

	
	
		Series *SeriesUnion
	

}

func NewMessageCarsModel() (*MessageCarsModel) {
	return &MessageCarsModel{}
}

func DeserializeMessageCarsModel(r io.Reader) (*MessageCarsModel, error) {
	t := NewMessageCarsModel()
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

func DeserializeMessageCarsModelFromSchema(r io.Reader, schema string) (*MessageCarsModel, error) {
	t := NewMessageCarsModel()

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

func writeMessageCarsModel(r *MessageCarsModel, w io.Writer) error {
	var err error
	
	err = writeBrandUnion( r.Brand, w)
	if err != nil {
		return err			
	}
	
	err = writeColorUnion( r.Color, w)
	if err != nil {
		return err			
	}
	
	err = writeSeriesUnion( r.Series, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *MessageCarsModel) Serialize(w io.Writer) error {
	return writeMessageCarsModel(r, w)
}

func (r *MessageCarsModel) Schema() string {
	return "{\"fields\":[{\"name\":\"brand\",\"type\":[\"null\",\"string\"]},{\"name\":\"color\",\"type\":[\"null\",\"string\"]},{\"name\":\"series\",\"type\":[\"null\",\"string\"]}],\"name\":\"proto.MessageCarsModel\",\"type\":\"record\"}"
}

func (r *MessageCarsModel) SchemaName() string {
	return "proto.MessageCarsModel"
}

func (_ *MessageCarsModel) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MessageCarsModel) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MessageCarsModel) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MessageCarsModel) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MessageCarsModel) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MessageCarsModel) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MessageCarsModel) SetString(v string) { panic("Unsupported operation") }
func (_ *MessageCarsModel) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MessageCarsModel) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.Brand = NewBrandUnion()
	
		
		
			return r.Brand
		
	
	case 1:
		
			r.Color = NewColorUnion()
	
		
		
			return r.Color
		
	
	case 2:
		
			r.Series = NewSeriesUnion()
	
		
		
			return r.Series
		
	
	}
	panic("Unknown field index")
}

func (r *MessageCarsModel) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *MessageCarsModel) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MessageCarsModel) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *MessageCarsModel) Finalize() { }
