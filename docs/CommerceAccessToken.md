# CommerceAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCommerceAccessToken

`func NewCommerceAccessToken() *CommerceAccessToken`

NewCommerceAccessToken instantiates a new CommerceAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceAccessTokenWithDefaults

`func NewCommerceAccessTokenWithDefaults() *CommerceAccessToken`

NewCommerceAccessTokenWithDefaults instantiates a new CommerceAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CommerceAccessToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CommerceAccessToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CommerceAccessToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CommerceAccessToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPermissions

`func (o *CommerceAccessToken) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CommerceAccessToken) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CommerceAccessToken) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CommerceAccessToken) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


