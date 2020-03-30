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


type UsernameUnionTypeEnum int
const (

	 UsernameUnionTypeEnumNull UsernameUnionTypeEnum = 0

	 UsernameUnionTypeEnumString UsernameUnionTypeEnum = 1

)

type UsernameUnion struct {

	Null *types.NullVal

	String string

	UnionType UsernameUnionTypeEnum
}

func writeUsernameUnion(r *UsernameUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UsernameUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UsernameUnionTypeEnumString:
		return vm.WriteString(r.String, w)
        
	}
	return fmt.Errorf("invalid value for *UsernameUnion")
}

func NewUsernameUnion() *UsernameUnion {
	return &UsernameUnion{}
}

func (_ *UsernameUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UsernameUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UsernameUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UsernameUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UsernameUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UsernameUnion) SetString(v string) { panic("Unsupported operation") }
func (r *UsernameUnion) SetLong(v int64) { 
	r.UnionType = (UsernameUnionTypeEnum)(v)
}
func (r *UsernameUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		
		return (*types.String)(&r.String)
		
	
	}
	panic("Unknown field index")
}
func (_ *UsernameUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UsernameUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UsernameUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UsernameUnion) Finalize()  { }
