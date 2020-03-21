// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package avro

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullMessageRoomInfoTypeEnum int
const (

	 UnionNullMessageRoomInfoTypeEnumNull UnionNullMessageRoomInfoTypeEnum = 0

	 UnionNullMessageRoomInfoTypeEnumMessageRoomInfo UnionNullMessageRoomInfoTypeEnum = 1

)

type UnionNullMessageRoomInfo struct {

	Null *types.NullVal

	MessageRoomInfo *MessageRoomInfo

	UnionType UnionNullMessageRoomInfoTypeEnum
}

func writeUnionNullMessageRoomInfo(r *UnionNullMessageRoomInfo, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullMessageRoomInfoTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullMessageRoomInfoTypeEnumMessageRoomInfo:
		return writeMessageRoomInfo(r.MessageRoomInfo, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullMessageRoomInfo")
}

func NewUnionNullMessageRoomInfo() *UnionNullMessageRoomInfo {
	return &UnionNullMessageRoomInfo{}
}

func (_ *UnionNullMessageRoomInfo) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullMessageRoomInfo) SetLong(v int64) { 
	r.UnionType = (UnionNullMessageRoomInfoTypeEnum)(v)
}
func (r *UnionNullMessageRoomInfo) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.MessageRoomInfo = NewMessageRoomInfo()
		
		
		return r.MessageRoomInfo
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullMessageRoomInfo) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullMessageRoomInfo) Finalize()  { }