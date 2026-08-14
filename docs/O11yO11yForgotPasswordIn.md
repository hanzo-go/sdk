# O11yO11yForgotPasswordIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address to mail the reset link to. Required. | [optional] 
**FrontendBaseURL** | Pointer to **string** | FrontendBaseURL is the console origin the reset link is built on. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the address belongs to. Required. | [optional] 

## Methods

### NewO11yO11yForgotPasswordIn

`func NewO11yO11yForgotPasswordIn() *O11yO11yForgotPasswordIn`

NewO11yO11yForgotPasswordIn instantiates a new O11yO11yForgotPasswordIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yForgotPasswordInWithDefaults

`func NewO11yO11yForgotPasswordInWithDefaults() *O11yO11yForgotPasswordIn`

NewO11yO11yForgotPasswordInWithDefaults instantiates a new O11yO11yForgotPasswordIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *O11yO11yForgotPasswordIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yForgotPasswordIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yForgotPasswordIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yForgotPasswordIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFrontendBaseURL

`func (o *O11yO11yForgotPasswordIn) GetFrontendBaseURL() string`

GetFrontendBaseURL returns the FrontendBaseURL field if non-nil, zero value otherwise.

### GetFrontendBaseURLOk

`func (o *O11yO11yForgotPasswordIn) GetFrontendBaseURLOk() (*string, bool)`

GetFrontendBaseURLOk returns a tuple with the FrontendBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendBaseURL

`func (o *O11yO11yForgotPasswordIn) SetFrontendBaseURL(v string)`

SetFrontendBaseURL sets FrontendBaseURL field to given value.

### HasFrontendBaseURL

`func (o *O11yO11yForgotPasswordIn) HasFrontendBaseURL() bool`

HasFrontendBaseURL returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yForgotPasswordIn) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yForgotPasswordIn) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yForgotPasswordIn) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yForgotPasswordIn) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


