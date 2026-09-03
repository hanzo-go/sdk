# RegistryToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **int64** | Expires is the token&#39;s lifetime in seconds. | [optional] 
**Ref** | Pointer to **string** | Ref is the one repository reference the token can pull. | [optional] 
**Token** | Pointer to **string** | Token is the bearer to present on the OCI wire (&#x60;Authorization: Bearer …&#x60; against the host&#39;s /v2/ routes). | [optional] 

## Methods

### NewRegistryToken

`func NewRegistryToken() *RegistryToken`

NewRegistryToken instantiates a new RegistryToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryTokenWithDefaults

`func NewRegistryTokenWithDefaults() *RegistryToken`

NewRegistryTokenWithDefaults instantiates a new RegistryToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *RegistryToken) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *RegistryToken) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *RegistryToken) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *RegistryToken) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetRef

`func (o *RegistryToken) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RegistryToken) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RegistryToken) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RegistryToken) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetToken

`func (o *RegistryToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RegistryToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RegistryToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RegistryToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


