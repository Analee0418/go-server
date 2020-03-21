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


type UnionNullRequestSalesAdvisorReceivingCustomersTypeEnum int
const (

	 UnionNullRequestSalesAdvisorReceivingCustomersTypeEnumNull UnionNullRequestSalesAdvisorReceivingCustomersTypeEnum = 0

	 UnionNullRequestSalesAdvisorReceivingCustomersTypeEnumRequestSalesAdvisorReceivingCustomers UnionNullRequestSalesAdvisorReceivingCustomersTypeEnum = 1

)

type UnionNullRequestSalesAdvisorReceivingCustomers struct {

	Null *types.NullVal

	RequestSalesAdvisorReceivingCustomers *RequestSalesAdvisorReceivingCustomers

	UnionType UnionNullRequestSalesAdvisorReceivingCustomersTypeEnum
}

func writeUnionNullRequestSalesAdvisorReceivingCustomers(r *UnionNullRequestSalesAdvisorReceivingCustomers, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullRequestSalesAdvisorReceivingCustomersTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullRequestSalesAdvisorReceivingCustomersTypeEnumRequestSalesAdvisorReceivingCustomers:
		return writeRequestSalesAdvisorReceivingCustomers(r.RequestSalesAdvisorReceivingCustomers, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullRequestSalesAdvisorReceivingCustomers")
}

func NewUnionNullRequestSalesAdvisorReceivingCustomers() *UnionNullRequestSalesAdvisorReceivingCustomers {
	return &UnionNullRequestSalesAdvisorReceivingCustomers{}
}

func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullRequestSalesAdvisorReceivingCustomers) SetLong(v int64) { 
	r.UnionType = (UnionNullRequestSalesAdvisorReceivingCustomersTypeEnum)(v)
}
func (r *UnionNullRequestSalesAdvisorReceivingCustomers) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.RequestSalesAdvisorReceivingCustomers = NewRequestSalesAdvisorReceivingCustomers()
		
		
		return r.RequestSalesAdvisorReceivingCustomers
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorReceivingCustomers) Finalize()  { }