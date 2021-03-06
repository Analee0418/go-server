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


type TipsUnionTypeEnum int
const (

	 TipsUnionTypeEnumNull TipsUnionTypeEnum = 0

	 TipsUnionTypeEnumString TipsUnionTypeEnum = 1

)

type TipsUnion struct {

	Null *types.NullVal

	String string

	UnionType TipsUnionTypeEnum
}

func writeTipsUnion(r *TipsUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case TipsUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case TipsUnionTypeEnumString:
		return vm.WriteString(r.String, w)
        
	}
	return fmt.Errorf("invalid value for *TipsUnion")
}

func NewTipsUnion() *TipsUnion {
	return &TipsUnion{}
}

func (_ *TipsUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *TipsUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *TipsUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *TipsUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *TipsUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *TipsUnion) SetString(v string) { panic("Unsupported operation") }
func (r *TipsUnion) SetLong(v int64) { 
	r.UnionType = (TipsUnionTypeEnum)(v)
}
func (r *TipsUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		
		return (*types.String)(&r.String)
		
	
	}
	panic("Unknown field index")
}
func (_ *TipsUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *TipsUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *TipsUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *TipsUnion) Finalize()  { }
