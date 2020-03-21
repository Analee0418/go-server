// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
)

func writeArrayMessageCustomersInfo(r []*MessageCustomersInfo, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)),w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeMessageCustomersInfo(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0,w)
}



type ArrayMessageCustomersInfoWrapper []*MessageCustomersInfo

func (_ *ArrayMessageCustomersInfoWrapper) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) SetInt(v int32) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) SetLong(v int64) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) SetString(v string) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayMessageCustomersInfoWrapper) Finalize() { }
func (_ *ArrayMessageCustomersInfoWrapper) SetDefault(i int) { panic("Unsupported operation") }
func (r *ArrayMessageCustomersInfoWrapper) AppendArray() types.Field {
	var v *MessageCustomersInfo
	
	v = NewMessageCustomersInfo()

 	
	*r = append(*r, v)
        
        return (*r)[len(*r)-1]
        
}