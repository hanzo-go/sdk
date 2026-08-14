# O11yO11yAPIKeyCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **int32** | ExpiresAt is when the key stops working, as a unix timestamp in seconds. Zero means it never expires; a past timestamp is refused. | [optional] 
**Name** | Pointer to **string** | Name is the key&#39;s name: a lowercase letter followed by lowercase letters, digits or hyphens, at most 80 characters. Required. | [optional] 

## Methods

### NewO11yO11yAPIKeyCreateIn

`func NewO11yO11yAPIKeyCreateIn() *O11yO11yAPIKeyCreateIn`

NewO11yO11yAPIKeyCreateIn instantiates a new O11yO11yAPIKeyCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAPIKeyCreateInWithDefaults

`func NewO11yO11yAPIKeyCreateInWithDefaults() *O11yO11yAPIKeyCreateIn`

NewO11yO11yAPIKeyCreateInWithDefaults instantiates a new O11yO11yAPIKeyCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *O11yO11yAPIKeyCreateIn) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *O11yO11yAPIKeyCreateIn) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *O11yO11yAPIKeyCreateIn) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *O11yO11yAPIKeyCreateIn) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yAPIKeyCreateIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yAPIKeyCreateIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yAPIKeyCreateIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yAPIKeyCreateIn) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


