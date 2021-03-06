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

func writeMapString(r *MapString, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteString(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapString struct {
	keys []string
	values []string
	M map[string]string
}

func NewMapString() *MapString{
	return &MapString {
		keys: make([]string, 0),
		values: make([]string, 0),
		M: make(map[string]string),
	}
}

func (_ *MapString) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MapString) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MapString) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MapString) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MapString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MapString) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MapString) SetString(v string) { panic("Unsupported operation") }
func (_ *MapString) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *MapString) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapString) SetDefault(i int) { panic("Unsupported operation") }
func (r *MapString) Finalize() { 
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapString) AppendMap(key string) types.Field { 
	r.keys = append(r.keys, key)
	var v string
	
	r.values = append(r.values, v)
	
	return (*types.String)(&r.values[len(r.values)-1])
	
}

func (_ *MapString) AppendArray() types.Field { panic("Unsupported operation") }

