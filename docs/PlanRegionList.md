# PlanRegionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **[]interface{}** | Regions are the regions cloud capacity is offered in, each an opaque object exactly as the catalog emits it — typically id, name, location and flag. | [optional] 

## Methods

### NewPlanRegionList

`func NewPlanRegionList() *PlanRegionList`

NewPlanRegionList instantiates a new PlanRegionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanRegionListWithDefaults

`func NewPlanRegionListWithDefaults() *PlanRegionList`

NewPlanRegionListWithDefaults instantiates a new PlanRegionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *PlanRegionList) GetRegions() []interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PlanRegionList) GetRegionsOk() (*[]interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PlanRegionList) SetRegions(v []interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PlanRegionList) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


