# Form

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the IRS designation, e.g. \&quot;SS-4\&quot;. | [optional] 
**Name** | Pointer to **string** | Name is the form&#39;s own title, so a reader need not already know the code. | [optional] 
**Signed** | Pointer to **bool** | Signed reports whether we hold the signature. | [optional] 
**Why** | Pointer to **string** | Why states what this form is for in this application — the same form is owed for different reasons on different paths. | [optional] 

## Methods

### NewForm

`func NewForm() *Form`

NewForm instantiates a new Form object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormWithDefaults

`func NewFormWithDefaults() *Form`

NewFormWithDefaults instantiates a new Form object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Form) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Form) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Form) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Form) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *Form) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Form) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Form) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Form) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSigned

`func (o *Form) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *Form) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *Form) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *Form) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetWhy

`func (o *Form) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *Form) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *Form) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *Form) HasWhy() bool`

HasWhy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


