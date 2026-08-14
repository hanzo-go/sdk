# ReferenceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** | Next is the key to page from, empty when this is the last page. | [optional] 
**Overrides** | Pointer to [**[]ReferenceOverride**](ReferenceOverride.md) | Overrides is YOUR org&#39;s entries over that baseline, in key order. They are held in your organisation&#39;s own store and are not visible to any other. | [optional] 
**Set** | Pointer to [**ReferenceSet**](ReferenceSet.md) | Set is the published set: its version, its freshness and its sources. | [optional] 

## Methods

### NewReferenceOut

`func NewReferenceOut() *ReferenceOut`

NewReferenceOut instantiates a new ReferenceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceOutWithDefaults

`func NewReferenceOutWithDefaults() *ReferenceOut`

NewReferenceOutWithDefaults instantiates a new ReferenceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ReferenceOut) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ReferenceOut) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ReferenceOut) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ReferenceOut) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOverrides

`func (o *ReferenceOut) GetOverrides() []ReferenceOverride`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ReferenceOut) GetOverridesOk() (*[]ReferenceOverride, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ReferenceOut) SetOverrides(v []ReferenceOverride)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *ReferenceOut) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetSet

`func (o *ReferenceOut) GetSet() ReferenceSet`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *ReferenceOut) GetSetOk() (*ReferenceSet, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *ReferenceOut) SetSet(v ReferenceSet)`

SetSet sets Set field to given value.

### HasSet

`func (o *ReferenceOut) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


