# IamControllersVerifyCodeForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to **string** |  | [optional] 
**Username** | **string** | Email, phone, or username; must resolve to a user | 
**Name** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**CountryCode** | Pointer to **string** | E.164 country code for phone | [optional] 

## Methods

### NewIamControllersVerifyCodeForm

`func NewIamControllersVerifyCodeForm(username string, code string, ) *IamControllersVerifyCodeForm`

NewIamControllersVerifyCodeForm instantiates a new IamControllersVerifyCodeForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamControllersVerifyCodeFormWithDefaults

`func NewIamControllersVerifyCodeFormWithDefaults() *IamControllersVerifyCodeForm`

NewIamControllersVerifyCodeFormWithDefaults instantiates a new IamControllersVerifyCodeForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *IamControllersVerifyCodeForm) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamControllersVerifyCodeForm) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamControllersVerifyCodeForm) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamControllersVerifyCodeForm) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetUsername

`func (o *IamControllersVerifyCodeForm) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamControllersVerifyCodeForm) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamControllersVerifyCodeForm) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetName

`func (o *IamControllersVerifyCodeForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamControllersVerifyCodeForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamControllersVerifyCodeForm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamControllersVerifyCodeForm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *IamControllersVerifyCodeForm) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamControllersVerifyCodeForm) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamControllersVerifyCodeForm) SetCode(v string)`

SetCode sets Code field to given value.


### GetCountryCode

`func (o *IamControllersVerifyCodeForm) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *IamControllersVerifyCodeForm) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *IamControllersVerifyCodeForm) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *IamControllersVerifyCodeForm) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


