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


type Sales_advisor_receiving_customersUnionTypeEnum int
const (

	 Sales_advisor_receiving_customersUnionTypeEnumNull Sales_advisor_receiving_customersUnionTypeEnum = 0

	 Sales_advisor_receiving_customersUnionTypeEnumRequestSalesAdvisorReceivingCustomers Sales_advisor_receiving_customersUnionTypeEnum = 1

)

type Sales_advisor_receiving_customersUnion struct {

	Null *types.NullVal

	RequestSalesAdvisorReceivingCustomers *RequestSalesAdvisorReceivingCustomers

	UnionType Sales_advisor_receiving_customersUnionTypeEnum
}

func writeSales_advisor_receiving_customersUnion(r *Sales_advisor_receiving_customersUnion, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case Sales_advisor_receiving_customersUnionTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case Sales_advisor_receiving_customersUnionTypeEnumRequestSalesAdvisorReceivingCustomers:
		return writeRequestSalesAdvisorReceivingCustomers(r.RequestSalesAdvisorReceivingCustomers, w)
        
	}
	return fmt.Errorf("invalid value for *Sales_advisor_receiving_customersUnion")
}

func NewSales_advisor_receiving_customersUnion() *Sales_advisor_receiving_customersUnion {
	return &Sales_advisor_receiving_customersUnion{}
}

func (_ *Sales_advisor_receiving_customersUnion) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) SetString(v string) { panic("Unsupported operation") }
func (r *Sales_advisor_receiving_customersUnion) SetLong(v int64) { 
	r.UnionType = (Sales_advisor_receiving_customersUnionTypeEnum)(v)
}
func (r *Sales_advisor_receiving_customersUnion) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.RequestSalesAdvisorReceivingCustomers = NewRequestSalesAdvisorReceivingCustomers()
		
		
		return r.RequestSalesAdvisorReceivingCustomers
		
	
	}
	panic("Unknown field index")
}
func (_ *Sales_advisor_receiving_customersUnion) SetDefault(i int) { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Sales_advisor_receiving_customersUnion) Finalize()  { }
