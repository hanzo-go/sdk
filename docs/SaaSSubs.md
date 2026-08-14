# SaaSSubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByPlan** | Pointer to [**[]SaaSPlan**](SaaSPlan.md) |  | [optional] 
**Canceled** | Pointer to **int32** |  | [optional] 
**New** | Pointer to **int32** |  | [optional] 
**Recent** | Pointer to [**[]SaaSEvent**](SaaSEvent.md) |  | [optional] 
**TrialsActive** | Pointer to **int32** |  | [optional] 

## Methods

### NewSaaSSubs

`func NewSaaSSubs() *SaaSSubs`

NewSaaSSubs instantiates a new SaaSSubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaaSSubsWithDefaults

`func NewSaaSSubsWithDefaults() *SaaSSubs`

NewSaaSSubsWithDefaults instantiates a new SaaSSubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByPlan

`func (o *SaaSSubs) GetByPlan() []SaaSPlan`

GetByPlan returns the ByPlan field if non-nil, zero value otherwise.

### GetByPlanOk

`func (o *SaaSSubs) GetByPlanOk() (*[]SaaSPlan, bool)`

GetByPlanOk returns a tuple with the ByPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByPlan

`func (o *SaaSSubs) SetByPlan(v []SaaSPlan)`

SetByPlan sets ByPlan field to given value.

### HasByPlan

`func (o *SaaSSubs) HasByPlan() bool`

HasByPlan returns a boolean if a field has been set.

### GetCanceled

`func (o *SaaSSubs) GetCanceled() int32`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *SaaSSubs) GetCanceledOk() (*int32, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *SaaSSubs) SetCanceled(v int32)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *SaaSSubs) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### GetNew

`func (o *SaaSSubs) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *SaaSSubs) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *SaaSSubs) SetNew(v int32)`

SetNew sets New field to given value.

### HasNew

`func (o *SaaSSubs) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetRecent

`func (o *SaaSSubs) GetRecent() []SaaSEvent`

GetRecent returns the Recent field if non-nil, zero value otherwise.

### GetRecentOk

`func (o *SaaSSubs) GetRecentOk() (*[]SaaSEvent, bool)`

GetRecentOk returns a tuple with the Recent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecent

`func (o *SaaSSubs) SetRecent(v []SaaSEvent)`

SetRecent sets Recent field to given value.

### HasRecent

`func (o *SaaSSubs) HasRecent() bool`

HasRecent returns a boolean if a field has been set.

### GetTrialsActive

`func (o *SaaSSubs) GetTrialsActive() int32`

GetTrialsActive returns the TrialsActive field if non-nil, zero value otherwise.

### GetTrialsActiveOk

`func (o *SaaSSubs) GetTrialsActiveOk() (*int32, bool)`

GetTrialsActiveOk returns a tuple with the TrialsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsActive

`func (o *SaaSSubs) SetTrialsActive(v int32)`

SetTrialsActive sets TrialsActive field to given value.

### HasTrialsActive

`func (o *SaaSSubs) HasTrialsActive() bool`

HasTrialsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


