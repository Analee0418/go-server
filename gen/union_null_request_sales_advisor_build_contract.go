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


type UnionNullRequestSalesAdvisorBuildContractTypeEnum int
const (

	 UnionNullRequestSalesAdvisorBuildContractTypeEnumNull UnionNullRequestSalesAdvisorBuildContractTypeEnum = 0

	 UnionNullRequestSalesAdvisorBuildContractTypeEnumRequestSalesAdvisorBuildContract UnionNullRequestSalesAdvisorBuildContractTypeEnum = 1

)

type UnionNullRequestSalesAdvisorBuildContract struct {

	Null *types.NullVal

	RequestSalesAdvisorBuildContract *RequestSalesAdvisorBuildContract

	UnionType UnionNullRequestSalesAdvisorBuildContractTypeEnum
}

func writeUnionNullRequestSalesAdvisorBuildContract(r *UnionNullRequestSalesAdvisorBuildContract, w io.Writer) error {
	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType{
	
	case UnionNullRequestSalesAdvisorBuildContractTypeEnumNull:
		return vm.WriteNull(r.Null, w)
        
	case UnionNullRequestSalesAdvisorBuildContractTypeEnumRequestSalesAdvisorBuildContract:
		return writeRequestSalesAdvisorBuildContract(r.RequestSalesAdvisorBuildContract, w)
        
	}
	return fmt.Errorf("invalid value for *UnionNullRequestSalesAdvisorBuildContract")
}

func NewUnionNullRequestSalesAdvisorBuildContract() *UnionNullRequestSalesAdvisorBuildContract {
	return &UnionNullRequestSalesAdvisorBuildContract{}
}

func (_ *UnionNullRequestSalesAdvisorBuildContract) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) SetString(v string) { panic("Unsupported operation") }
func (r *UnionNullRequestSalesAdvisorBuildContract) SetLong(v int64) { 
	r.UnionType = (UnionNullRequestSalesAdvisorBuildContractTypeEnum)(v)
}
func (r *UnionNullRequestSalesAdvisorBuildContract) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
		return r.Null
		
	
	case 1:
		
		r.RequestSalesAdvisorBuildContract = NewRequestSalesAdvisorBuildContract()
		
		
		return r.RequestSalesAdvisorBuildContract
		
	
	}
	panic("Unknown field index")
}
func (_ *UnionNullRequestSalesAdvisorBuildContract) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullRequestSalesAdvisorBuildContract) Finalize()  { }
