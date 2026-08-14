# O11yO11yEmailPasswordSessionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the account&#39;s address. Required. | [optional] 
**OrgId** | Pointer to **string** | OrgID picks the org to sign into when the address belongs to several. | [optional] 
**Password** | Pointer to **string** | Password is the account&#39;s password. Required. | [optional] 

## Methods

### NewO11yO11yEmailPasswordSessionIn

`func NewO11yO11yEmailPasswordSessionIn() *O11yO11yEmailPasswordSessionIn`

NewO11yO11yEmailPasswordSessionIn instantiates a new O11yO11yEmailPasswordSessionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yEmailPasswordSessionInWithDefaults

`func NewO11yO11yEmailPasswordSessionInWithDefaults() *O11yO11yEmailPasswordSessionIn`

NewO11yO11yEmailPasswordSessionInWithDefaults instantiates a new O11yO11yEmailPasswordSessionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *O11yO11yEmailPasswordSessionIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yEmailPasswordSessionIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yEmailPasswordSessionIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yEmailPasswordSessionIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yEmailPasswordSessionIn) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yEmailPasswordSessionIn) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yEmailPasswordSessionIn) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yEmailPasswordSessionIn) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPassword

`func (o *O11yO11yEmailPasswordSessionIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *O11yO11yEmailPasswordSessionIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *O11yO11yEmailPasswordSessionIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *O11yO11yEmailPasswordSessionIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


