# CommerceCreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Company** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 

## Methods

### NewCommerceCreateAccountRequest

`func NewCommerceCreateAccountRequest(email string, password string, firstName string, lastName string, ) *CommerceCreateAccountRequest`

NewCommerceCreateAccountRequest instantiates a new CommerceCreateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceCreateAccountRequestWithDefaults

`func NewCommerceCreateAccountRequestWithDefaults() *CommerceCreateAccountRequest`

NewCommerceCreateAccountRequestWithDefaults instantiates a new CommerceCreateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CommerceCreateAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommerceCreateAccountRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommerceCreateAccountRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CommerceCreateAccountRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CommerceCreateAccountRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CommerceCreateAccountRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetFirstName

`func (o *CommerceCreateAccountRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CommerceCreateAccountRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CommerceCreateAccountRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CommerceCreateAccountRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CommerceCreateAccountRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CommerceCreateAccountRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetCompany

`func (o *CommerceCreateAccountRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CommerceCreateAccountRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CommerceCreateAccountRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CommerceCreateAccountRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetPhone

`func (o *CommerceCreateAccountRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CommerceCreateAccountRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CommerceCreateAccountRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CommerceCreateAccountRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


