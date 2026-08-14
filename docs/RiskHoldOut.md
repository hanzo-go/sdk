# RiskHoldOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changed** | Pointer to **int32** | Changed is how many records moved into that state. A record already in it is not counted and is not an error: the op is idempotent, so a retry after a network failure is safe. | [optional] 
**Held** | Pointer to **int32** | Held is how many records this tenant is now holding, at any age. Retention never disposes of one. | [optional] 
**Hold** | Pointer to **bool** | Hold echoes the state asked for. | [optional] 
**Missing** | Pointer to **int32** | Missing is how many of the named ids this tenant does not hold. It is reported rather than refused, so a sweep over a list that includes disposed records still places every hold it can — but it is REPORTED, because a hold that silently did nothing is a compliance control that lies. | [optional] 

## Methods

### NewRiskHoldOut

`func NewRiskHoldOut() *RiskHoldOut`

NewRiskHoldOut instantiates a new RiskHoldOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskHoldOutWithDefaults

`func NewRiskHoldOutWithDefaults() *RiskHoldOut`

NewRiskHoldOutWithDefaults instantiates a new RiskHoldOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanged

`func (o *RiskHoldOut) GetChanged() int32`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *RiskHoldOut) GetChangedOk() (*int32, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *RiskHoldOut) SetChanged(v int32)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *RiskHoldOut) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetHeld

`func (o *RiskHoldOut) GetHeld() int32`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *RiskHoldOut) GetHeldOk() (*int32, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *RiskHoldOut) SetHeld(v int32)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *RiskHoldOut) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetHold

`func (o *RiskHoldOut) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *RiskHoldOut) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *RiskHoldOut) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *RiskHoldOut) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetMissing

`func (o *RiskHoldOut) GetMissing() int32`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *RiskHoldOut) GetMissingOk() (*int32, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *RiskHoldOut) SetMissing(v int32)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *RiskHoldOut) HasMissing() bool`

HasMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


