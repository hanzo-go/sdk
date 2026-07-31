# CloudRegistryToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **int32** | Expires is the token&#39;s lifetime in seconds. | [optional] 
**Ref** | Pointer to **string** | Ref is the one repository reference the token can pull. | [optional] 
**Token** | Pointer to **string** | Token is the bearer to present on the OCI wire (&#x60;Authorization: Bearer …&#x60; against the host&#39;s /v2/ routes). | [optional] 

## Methods

### NewCloudRegistryToken

`func NewCloudRegistryToken() *CloudRegistryToken`

NewCloudRegistryToken instantiates a new CloudRegistryToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegistryTokenWithDefaults

`func NewCloudRegistryTokenWithDefaults() *CloudRegistryToken`

NewCloudRegistryTokenWithDefaults instantiates a new CloudRegistryToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *CloudRegistryToken) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CloudRegistryToken) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CloudRegistryToken) SetExpires(v int32)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CloudRegistryToken) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetRef

`func (o *CloudRegistryToken) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloudRegistryToken) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloudRegistryToken) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloudRegistryToken) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetToken

`func (o *CloudRegistryToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CloudRegistryToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CloudRegistryToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CloudRegistryToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


