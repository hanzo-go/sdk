# O11yO11yRegisterIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email is the admin&#39;s email. Required. | 
**Name** | Pointer to **string** | Name is the admin&#39;s display name. | [optional] 
**OrgDisplayName** | Pointer to **string** | OrgDisplayName is the organization&#39;s display name. | [optional] 
**OrgName** | Pointer to **string** | OrgName is the organization&#39;s name. | [optional] 
**Password** | Pointer to **string** | Password is the admin&#39;s password. | [optional] 

## Methods

### NewO11yO11yRegisterIn

`func NewO11yO11yRegisterIn(email string, ) *O11yO11yRegisterIn`

NewO11yO11yRegisterIn instantiates a new O11yO11yRegisterIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRegisterInWithDefaults

`func NewO11yO11yRegisterInWithDefaults() *O11yO11yRegisterIn`

NewO11yO11yRegisterInWithDefaults instantiates a new O11yO11yRegisterIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *O11yO11yRegisterIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yRegisterIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yRegisterIn) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *O11yO11yRegisterIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yRegisterIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yRegisterIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yRegisterIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgDisplayName

`func (o *O11yO11yRegisterIn) GetOrgDisplayName() string`

GetOrgDisplayName returns the OrgDisplayName field if non-nil, zero value otherwise.

### GetOrgDisplayNameOk

`func (o *O11yO11yRegisterIn) GetOrgDisplayNameOk() (*string, bool)`

GetOrgDisplayNameOk returns a tuple with the OrgDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDisplayName

`func (o *O11yO11yRegisterIn) SetOrgDisplayName(v string)`

SetOrgDisplayName sets OrgDisplayName field to given value.

### HasOrgDisplayName

`func (o *O11yO11yRegisterIn) HasOrgDisplayName() bool`

HasOrgDisplayName returns a boolean if a field has been set.

### GetOrgName

`func (o *O11yO11yRegisterIn) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *O11yO11yRegisterIn) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *O11yO11yRegisterIn) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *O11yO11yRegisterIn) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPassword

`func (o *O11yO11yRegisterIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *O11yO11yRegisterIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *O11yO11yRegisterIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *O11yO11yRegisterIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


