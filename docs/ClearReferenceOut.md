# ClearReferenceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleared** | Pointer to **bool** | Cleared is false when your org held no such override — which is not an error, it is the honest answer to a removal that had nothing to remove. | [optional] 
**Key** | Pointer to **string** | Key is the entry named. | [optional] 
**Overrides** | Pointer to **int64** | Overrides is how many your org still holds in this set. | [optional] 
**Set** | Pointer to **string** | Set is the set cleared in. | [optional] 

## Methods

### NewClearReferenceOut

`func NewClearReferenceOut() *ClearReferenceOut`

NewClearReferenceOut instantiates a new ClearReferenceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearReferenceOutWithDefaults

`func NewClearReferenceOutWithDefaults() *ClearReferenceOut`

NewClearReferenceOutWithDefaults instantiates a new ClearReferenceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleared

`func (o *ClearReferenceOut) GetCleared() bool`

GetCleared returns the Cleared field if non-nil, zero value otherwise.

### GetClearedOk

`func (o *ClearReferenceOut) GetClearedOk() (*bool, bool)`

GetClearedOk returns a tuple with the Cleared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleared

`func (o *ClearReferenceOut) SetCleared(v bool)`

SetCleared sets Cleared field to given value.

### HasCleared

`func (o *ClearReferenceOut) HasCleared() bool`

HasCleared returns a boolean if a field has been set.

### GetKey

`func (o *ClearReferenceOut) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ClearReferenceOut) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ClearReferenceOut) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ClearReferenceOut) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOverrides

`func (o *ClearReferenceOut) GetOverrides() int64`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ClearReferenceOut) GetOverridesOk() (*int64, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ClearReferenceOut) SetOverrides(v int64)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *ClearReferenceOut) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetSet

`func (o *ClearReferenceOut) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *ClearReferenceOut) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *ClearReferenceOut) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *ClearReferenceOut) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


