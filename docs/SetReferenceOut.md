# SetReferenceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | Pointer to **int64** | Overrides is how many your org now holds in this set. | [optional] 
**Set** | Pointer to **string** | Set is the set written in. | [optional] 
**Written** | Pointer to **int64** | Written is how many entries this call wrote. | [optional] 

## Methods

### NewSetReferenceOut

`func NewSetReferenceOut() *SetReferenceOut`

NewSetReferenceOut instantiates a new SetReferenceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetReferenceOutWithDefaults

`func NewSetReferenceOutWithDefaults() *SetReferenceOut`

NewSetReferenceOutWithDefaults instantiates a new SetReferenceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *SetReferenceOut) GetOverrides() int64`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *SetReferenceOut) GetOverridesOk() (*int64, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *SetReferenceOut) SetOverrides(v int64)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *SetReferenceOut) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetSet

`func (o *SetReferenceOut) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *SetReferenceOut) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *SetReferenceOut) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *SetReferenceOut) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetWritten

`func (o *SetReferenceOut) GetWritten() int64`

GetWritten returns the Written field if non-nil, zero value otherwise.

### GetWrittenOk

`func (o *SetReferenceOut) GetWrittenOk() (*int64, bool)`

GetWrittenOk returns a tuple with the Written field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritten

`func (o *SetReferenceOut) SetWritten(v int64)`

SetWritten sets Written field to given value.

### HasWritten

`func (o *SetReferenceOut) HasWritten() bool`

HasWritten returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


