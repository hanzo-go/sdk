# RiskDisposeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to **string** | Before disposes of assertions WRITTEN before this instant, RFC 3339. It is measured against the server clock at the write and not against the event or observation times, both of which the asserting caller supplies — a tenant that could back-date could delete a compliance record on demand. | [optional] 

## Methods

### NewRiskDisposeIn

`func NewRiskDisposeIn() *RiskDisposeIn`

NewRiskDisposeIn instantiates a new RiskDisposeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDisposeInWithDefaults

`func NewRiskDisposeInWithDefaults() *RiskDisposeIn`

NewRiskDisposeInWithDefaults instantiates a new RiskDisposeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *RiskDisposeIn) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *RiskDisposeIn) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *RiskDisposeIn) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *RiskDisposeIn) HasBefore() bool`

HasBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


