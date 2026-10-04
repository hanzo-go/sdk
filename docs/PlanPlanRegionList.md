# PlanPlanRegionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **[]interface{}** | Regions are the regions cloud capacity is offered in, each an opaque object exactly as the catalog emits it — typically id, name, location and flag. | [optional] 

## Methods

### NewPlanPlanRegionList

`func NewPlanPlanRegionList() *PlanPlanRegionList`

NewPlanPlanRegionList instantiates a new PlanPlanRegionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanPlanRegionListWithDefaults

`func NewPlanPlanRegionListWithDefaults() *PlanPlanRegionList`

NewPlanPlanRegionListWithDefaults instantiates a new PlanPlanRegionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *PlanPlanRegionList) GetRegions() []interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PlanPlanRegionList) GetRegionsOk() (*[]interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PlanPlanRegionList) SetRegions(v []interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PlanPlanRegionList) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


