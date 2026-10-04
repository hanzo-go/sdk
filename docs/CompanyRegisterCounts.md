# CompanyRegisterCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByStage** | Pointer to **map[string]int64** | ByStage counts formations per stage, keyed by the stage name. | [optional] 
**Total** | Pointer to **int64** | Total is every formation in the register. | [optional] 

## Methods

### NewCompanyRegisterCounts

`func NewCompanyRegisterCounts() *CompanyRegisterCounts`

NewCompanyRegisterCounts instantiates a new CompanyRegisterCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyRegisterCountsWithDefaults

`func NewCompanyRegisterCountsWithDefaults() *CompanyRegisterCounts`

NewCompanyRegisterCountsWithDefaults instantiates a new CompanyRegisterCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByStage

`func (o *CompanyRegisterCounts) GetByStage() map[string]int64`

GetByStage returns the ByStage field if non-nil, zero value otherwise.

### GetByStageOk

`func (o *CompanyRegisterCounts) GetByStageOk() (*map[string]int64, bool)`

GetByStageOk returns a tuple with the ByStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStage

`func (o *CompanyRegisterCounts) SetByStage(v map[string]int64)`

SetByStage sets ByStage field to given value.

### HasByStage

`func (o *CompanyRegisterCounts) HasByStage() bool`

HasByStage returns a boolean if a field has been set.

### GetTotal

`func (o *CompanyRegisterCounts) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CompanyRegisterCounts) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CompanyRegisterCounts) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CompanyRegisterCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


