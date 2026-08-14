# O11yPostableIngestionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewO11yPostableIngestionKey

`func NewO11yPostableIngestionKey() *O11yPostableIngestionKey`

NewO11yPostableIngestionKey instantiates a new O11yPostableIngestionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPostableIngestionKeyWithDefaults

`func NewO11yPostableIngestionKeyWithDefaults() *O11yPostableIngestionKey`

NewO11yPostableIngestionKeyWithDefaults instantiates a new O11yPostableIngestionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *O11yPostableIngestionKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *O11yPostableIngestionKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *O11yPostableIngestionKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *O11yPostableIngestionKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *O11yPostableIngestionKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yPostableIngestionKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yPostableIngestionKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yPostableIngestionKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *O11yPostableIngestionKey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yPostableIngestionKey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yPostableIngestionKey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yPostableIngestionKey) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


