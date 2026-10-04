# ReferenceReferenceVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf is when the oldest of them was current, RFC 3339. | [optional] 
**Refusal** | Pointer to **string** | Refusal is why it could not be consulted, when it could not. | [optional] 
**Set** | Pointer to **string** | Set is the name the consulted set is addressed by. | [optional] 
**Stale** | Pointer to **bool** | Stale is whether it is past its freshness bound. | [optional] 
**Version** | Pointer to **string** | Version is every contributing publisher and its content digest. | [optional] 

## Methods

### NewReferenceReferenceVersion

`func NewReferenceReferenceVersion() *ReferenceReferenceVersion`

NewReferenceReferenceVersion instantiates a new ReferenceReferenceVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceReferenceVersionWithDefaults

`func NewReferenceReferenceVersionWithDefaults() *ReferenceReferenceVersion`

NewReferenceReferenceVersionWithDefaults instantiates a new ReferenceReferenceVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *ReferenceReferenceVersion) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ReferenceReferenceVersion) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ReferenceReferenceVersion) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ReferenceReferenceVersion) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetRefusal

`func (o *ReferenceReferenceVersion) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *ReferenceReferenceVersion) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *ReferenceReferenceVersion) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *ReferenceReferenceVersion) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetSet

`func (o *ReferenceReferenceVersion) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *ReferenceReferenceVersion) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *ReferenceReferenceVersion) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *ReferenceReferenceVersion) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetStale

`func (o *ReferenceReferenceVersion) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ReferenceReferenceVersion) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ReferenceReferenceVersion) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ReferenceReferenceVersion) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetVersion

`func (o *ReferenceReferenceVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReferenceReferenceVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReferenceReferenceVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReferenceReferenceVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


