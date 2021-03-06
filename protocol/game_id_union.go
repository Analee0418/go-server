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


type GameIDUnionTypeEnum int
const (

	 GameIDUnionTypeEnumNull GameIDUnionTypeEnum = 0

	 GameIDUnionTypeEnumString GameIDUnionTypeEnum = 1

)

type GameIDUnion struct {

	Null *types.NullVal

	String string

	UnionType GameIDUnionTypeEnum
}

func writeGameIDUnion(r *GameIDUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case GameIDUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case GameIDUnionTypeEnumString:
		return vm.WriteString(r.String, w)
        
	}
	return fmt.Errorf("invalid value for *GameIDUnion")
}

func NewGameIDUnion() *GameIDUnion {
	return &GameIDUnion{}
}

func (_ *GameIDUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *GameIDUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *GameIDUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *GameIDUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *GameIDUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *GameIDUnion) SetString(v string) { panic("Unsupported operation") }
func (r *GameIDUnion) SetLong(v int64) { 
	r.UnionType = (GameIDUnionTypeEnum)(v)
}
func (r *GameIDUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		
		return (*types.String)(&r.String)
		
	
	}
	panic("Unknown field index")
}
func (_ *GameIDUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *GameIDUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *GameIDUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *GameIDUnion) Finalize()  { }
