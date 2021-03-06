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


type SeriesUnionTypeEnum int
const (

	 SeriesUnionTypeEnumNull SeriesUnionTypeEnum = 0

	 SeriesUnionTypeEnumString SeriesUnionTypeEnum = 1

)

type SeriesUnion struct {

	Null *types.NullVal

	String string

	UnionType SeriesUnionTypeEnum
}

func writeSeriesUnion(r *SeriesUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case SeriesUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case SeriesUnionTypeEnumString:
		return vm.WriteString(r.String, w)
        
	}
	return fmt.Errorf("invalid value for *SeriesUnion")
}

func NewSeriesUnion() *SeriesUnion {
	return &SeriesUnion{}
}

func (_ *SeriesUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *SeriesUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *SeriesUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *SeriesUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *SeriesUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *SeriesUnion) SetString(v string) { panic("Unsupported operation") }
func (r *SeriesUnion) SetLong(v int64) { 
	r.UnionType = (SeriesUnionTypeEnum)(v)
}
func (r *SeriesUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		
		return (*types.String)(&r.String)
		
	
	}
	panic("Unknown field index")
}
func (_ *SeriesUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *SeriesUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *SeriesUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *SeriesUnion) Finalize()  { }
