# CloudRegisterCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByStage** | Pointer to **map[string]int32** | ByStage counts formations per stage, keyed by the stage name. | [optional] 
**Total** | Pointer to **int32** | Total is every formation in the register. | [optional] 

## Methods

### NewCloudRegisterCounts

`func NewCloudRegisterCounts() *CloudRegisterCounts`

NewCloudRegisterCounts instantiates a new CloudRegisterCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegisterCountsWithDefaults

`func NewCloudRegisterCountsWithDefaults() *CloudRegisterCounts`

NewCloudRegisterCountsWithDefaults instantiates a new CloudRegisterCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByStage

`func (o *CloudRegisterCounts) GetByStage() map[string]int32`

GetByStage returns the ByStage field if non-nil, zero value otherwise.

### GetByStageOk

`func (o *CloudRegisterCounts) GetByStageOk() (*map[string]int32, bool)`

GetByStageOk returns a tuple with the ByStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStage

`func (o *CloudRegisterCounts) SetByStage(v map[string]int32)`

SetByStage sets ByStage field to given value.

### HasByStage

`func (o *CloudRegisterCounts) HasByStage() bool`

HasByStage returns a boolean if a field has been set.

### GetTotal

`func (o *CloudRegisterCounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudRegisterCounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudRegisterCounts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudRegisterCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


