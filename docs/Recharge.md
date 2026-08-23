# Recharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Charged** | Pointer to **int32** | Charged is how many of them were actually charged. It is at most Orgs, and the difference is orgs whose balance was already above their threshold. | [optional] 
**Orgs** | Pointer to **int32** | Orgs is how many orgs the sweep considered — every org with auto-recharge armed, whether or not it needed charging. | [optional] 
**Results** | Pointer to [**[]Recharged**](Recharged.md) | Results is one row per org considered, so a sweep that charged nobody is still explainable. Never null. | [optional] 

## Methods

### NewRecharge

`func NewRecharge() *Recharge`

NewRecharge instantiates a new Recharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechargeWithDefaults

`func NewRechargeWithDefaults() *Recharge`

NewRechargeWithDefaults instantiates a new Recharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharged

`func (o *Recharge) GetCharged() int32`

GetCharged returns the Charged field if non-nil, zero value otherwise.

### GetChargedOk

`func (o *Recharge) GetChargedOk() (*int32, bool)`

GetChargedOk returns a tuple with the Charged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharged

`func (o *Recharge) SetCharged(v int32)`

SetCharged sets Charged field to given value.

### HasCharged

`func (o *Recharge) HasCharged() bool`

HasCharged returns a boolean if a field has been set.

### GetOrgs

`func (o *Recharge) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *Recharge) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *Recharge) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *Recharge) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetResults

`func (o *Recharge) GetResults() []Recharged`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Recharge) GetResultsOk() (*[]Recharged, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Recharge) SetResults(v []Recharged)`

SetResults sets Results field to given value.

### HasResults

`func (o *Recharge) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


