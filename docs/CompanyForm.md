# CompanyForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the IRS designation, e.g. \&quot;SS-4\&quot;. | [optional] 
**Name** | Pointer to **string** | Name is the form&#39;s own title, so a reader need not already know the code. | [optional] 
**Signed** | Pointer to **bool** | Signed reports whether we hold the signature. | [optional] 
**Why** | Pointer to **string** | Why states what this form is for in this application — the same form is owed for different reasons on different paths. | [optional] 

## Methods

### NewCompanyForm

`func NewCompanyForm() *CompanyForm`

NewCompanyForm instantiates a new CompanyForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFormWithDefaults

`func NewCompanyFormWithDefaults() *CompanyForm`

NewCompanyFormWithDefaults instantiates a new CompanyForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CompanyForm) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CompanyForm) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CompanyForm) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CompanyForm) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *CompanyForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyForm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyForm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSigned

`func (o *CompanyForm) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *CompanyForm) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *CompanyForm) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *CompanyForm) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetWhy

`func (o *CompanyForm) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *CompanyForm) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *CompanyForm) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *CompanyForm) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


