# BackfillResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forced** | Pointer to **bool** | Forced is true when the caller overrode the already-populated guard. | [optional] 
**SeededBefore** | Pointer to **string** | SeededBefore is the RFC3339 upper bound the seed actually used. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; — a seed that did not run answered an error instead. | [optional] 

## Methods

### NewBackfillResult

`func NewBackfillResult() *BackfillResult`

NewBackfillResult instantiates a new BackfillResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackfillResultWithDefaults

`func NewBackfillResultWithDefaults() *BackfillResult`

NewBackfillResultWithDefaults instantiates a new BackfillResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForced

`func (o *BackfillResult) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *BackfillResult) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *BackfillResult) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *BackfillResult) HasForced() bool`

HasForced returns a boolean if a field has been set.

### GetSeededBefore

`func (o *BackfillResult) GetSeededBefore() string`

GetSeededBefore returns the SeededBefore field if non-nil, zero value otherwise.

### GetSeededBeforeOk

`func (o *BackfillResult) GetSeededBeforeOk() (*string, bool)`

GetSeededBeforeOk returns a tuple with the SeededBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeededBefore

`func (o *BackfillResult) SetSeededBefore(v string)`

SetSeededBefore sets SeededBefore field to given value.

### HasSeededBefore

`func (o *BackfillResult) HasSeededBefore() bool`

HasSeededBefore returns a boolean if a field has been set.

### GetStatus

`func (o *BackfillResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackfillResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackfillResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackfillResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


