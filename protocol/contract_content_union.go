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


type Contract_contentUnionTypeEnum int
const (

	 Contract_contentUnionTypeEnumNull Contract_contentUnionTypeEnum = 0

	 Contract_contentUnionTypeEnumString Contract_contentUnionTypeEnum = 1

)

type Contract_contentUnion struct {

	Null *types.NullVal

	String string

	UnionType Contract_contentUnionTypeEnum
}

func writeContract_contentUnion(r *Contract_contentUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case Contract_contentUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case Contract_contentUnionTypeEnumString:
		return vm.WriteString(r.String, w)
        
	}
	return fmt.Errorf("invalid value for *Contract_contentUnion")
}

func NewContract_contentUnion() *Contract_contentUnion {
	return &Contract_contentUnion{}
}

func (_ *Contract_contentUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Contract_contentUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Contract_contentUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Contract_contentUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Contract_contentUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Contract_contentUnion) SetString(v string) { panic("Unsupported operation") }
func (r *Contract_contentUnion) SetLong(v int64) { 
	r.UnionType = (Contract_contentUnionTypeEnum)(v)
}
func (r *Contract_contentUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		
		return (*types.String)(&r.String)
		
	
	}
	panic("Unknown field index")
}
func (_ *Contract_contentUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *Contract_contentUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Contract_contentUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Contract_contentUnion) Finalize()  { }