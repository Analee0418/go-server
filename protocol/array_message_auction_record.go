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
)

func writeArrayMessageAuctionRecord(r []*MessageAuctionRecord, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)),w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeMessageAuctionRecord(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0,w)
}



type ArrayMessageAuctionRecordWrapper []*MessageAuctionRecord

func (_ *ArrayMessageAuctionRecordWrapper) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) SetInt(v int32) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) SetLong(v int64) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) SetString(v string) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayMessageAuctionRecordWrapper) Finalize() { }
func (_ *ArrayMessageAuctionRecordWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r *ArrayMessageAuctionRecordWrapper) AppendArray() types.Field {
	var v *MessageAuctionRecord
	
	v = NewMessageAuctionRecord()

 	
	*r = append(*r, v)
        
        return (*r)[len(*r)-1]
        
}