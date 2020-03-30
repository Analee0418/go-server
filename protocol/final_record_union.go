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


type Final_recordUnionTypeEnum int
const (

	 Final_recordUnionTypeEnumNull Final_recordUnionTypeEnum = 0

	 Final_recordUnionTypeEnumMessageAuctionRecord Final_recordUnionTypeEnum = 1

)

type Final_recordUnion struct {

	Null *types.NullVal

	MessageAuctionRecord *MessageAuctionRecord

	UnionType Final_recordUnionTypeEnum
}

func writeFinal_recordUnion(r *Final_recordUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case Final_recordUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case Final_recordUnionTypeEnumMessageAuctionRecord:
		return writeMessageAuctionRecord(r.MessageAuctionRecord, w)
        
	}
	return fmt.Errorf("invalid value for *Final_recordUnion")
}

func NewFinal_recordUnion() *Final_recordUnion {
	return &Final_recordUnion{}
}

func (_ *Final_recordUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Final_recordUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Final_recordUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Final_recordUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Final_recordUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Final_recordUnion) SetString(v string) { panic("Unsupported operation") }
func (r *Final_recordUnion) SetLong(v int64) { 
	r.UnionType = (Final_recordUnionTypeEnum)(v)
}
func (r *Final_recordUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.MessageAuctionRecord = NewMessageAuctionRecord()
		
		
		return r.MessageAuctionRecord
		
	
	}
	panic("Unknown field index")
}
func (_ *Final_recordUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *Final_recordUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Final_recordUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Final_recordUnion) Finalize()  { }