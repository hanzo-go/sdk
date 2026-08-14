# RefreshReferenceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **string** | Set is the set refreshed. | [optional] 
**Stale** | Pointer to **bool** | Stale is whether it is STILL past its freshness bound after the refresh, which is what a publisher that has stopped answering looks like. | [optional] 
**Took** | Pointer to [**[]ReferenceTaken**](ReferenceTaken.md) | Took is what each publisher contributed. | [optional] 
**Version** | Pointer to **string** | Version is the set&#39;s new composed version. | [optional] 

## Methods

### NewRefreshReferenceOut

`func NewRefreshReferenceOut() *RefreshReferenceOut`

NewRefreshReferenceOut instantiates a new RefreshReferenceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshReferenceOutWithDefaults

`func NewRefreshReferenceOutWithDefaults() *RefreshReferenceOut`

NewRefreshReferenceOutWithDefaults instantiates a new RefreshReferenceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *RefreshReferenceOut) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *RefreshReferenceOut) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *RefreshReferenceOut) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *RefreshReferenceOut) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetStale

`func (o *RefreshReferenceOut) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RefreshReferenceOut) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RefreshReferenceOut) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RefreshReferenceOut) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetTook

`func (o *RefreshReferenceOut) GetTook() []ReferenceTaken`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *RefreshReferenceOut) GetTookOk() (*[]ReferenceTaken, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *RefreshReferenceOut) SetTook(v []ReferenceTaken)`

SetTook sets Took field to given value.

### HasTook

`func (o *RefreshReferenceOut) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetVersion

`func (o *RefreshReferenceOut) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RefreshReferenceOut) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RefreshReferenceOut) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RefreshReferenceOut) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


